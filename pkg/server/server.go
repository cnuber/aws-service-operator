package server

import (
	"context"
	"fmt"
	"net/http"

	awsscheme "github.com/awslabs/aws-service-operator/pkg/client/clientset/versioned/scheme"
	"github.com/awslabs/aws-service-operator/pkg/config"
	opBase "github.com/awslabs/aws-service-operator/pkg/operators/base"
	"github.com/awslabs/aws-service-operator/pkg/queue"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/record"
)

const controllerName = "aws-service-operator"

// New creates a new server from a config
func New(config *config.Config) *Server {
	return &Server{
		Config: config,
	}
}

func (c *Server) exposeMetrics(errChan chan error, ctx context.Context) {
	c.Handle("/metrics", promhttp.Handler())
	server := http.Server{
		Addr:    ":9090",
		Handler: c,
	}
	defer server.Shutdown(ctx)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			errChan <- fmt.Errorf("unable to expose metrics: %v", err)
		}
	}()

	c.Config.Logger.Info("metrics server started")
	<-ctx.Done()
	c.Config.Logger.Info("metrics server stopped")
}

func (c *Server) watchOperatorResources(errChan chan error, ctx context.Context) {
	logger := c.Config.Logger

	logger.Info("getting kubernetes context")
	awsscheme.AddToScheme(scheme.Scheme)
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(logger.Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: c.Config.KubeClientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerName})
	c.Config.Recorder = recorder

	queueURL, queueARN, queueManager, err := queue.RegisterQueue(c.Config, "cloudformation")
	if err != nil {
		logger.WithError(err).Error("error reqistering queue")
	}
	c.Config.QueueURL = queueURL
	c.Config.QueueARN = queueARN

	operators := opBase.New(c.Config, queueManager)

	err = queue.SetQueuePolicy(c.Config, queueManager)
	if err != nil {
		logger.WithError(err).Error("error setting queue policy")
	}

	logger.WithFields(logrus.Fields{"resources": c.Config.Resources}).Info("Watching")
	go operators.Watch(ctx, corev1.NamespaceAll)
	go queue.Subscribe(c.Config, queueManager, ctx)
	<-ctx.Done()
	c.Config.Logger.Info("operators stopped")
}

// Run starts the server to listen to Kubernetes
func (c *Server) Run(ctx context.Context) {
	config := c.Config
	logger := config.Logger
	errChan := make(chan error, 1)

	logger.Info("starting metrics server")
	go c.exposeMetrics(errChan, ctx)

	logger.Info("starting resource watcher")
	go c.watchOperatorResources(errChan, ctx)

	for {
		select {
		case err := <-errChan:
			c.Config.Logger.WithError(err).Fatal(err)
		case <-ctx.Done():
			c.Config.Logger.Info("stop signal received. waiting for operators to stop")
			return
		}
	}
}
