/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/christopherhein/aws-operator/pkg/apis/operator.aws/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSNSTopics implements SNSTopicInterface
type FakeSNSTopics struct {
	Fake *FakeOperatorV1alpha1
	ns   string
}

var snstopicsResource = schema.GroupVersionResource{Group: "operator.aws", Version: "v1alpha1", Resource: "snstopics"}

var snstopicsKind = schema.GroupVersionKind{Group: "operator.aws", Version: "v1alpha1", Kind: "SNSTopic"}

// Get takes name of the sNSTopic, and returns the corresponding sNSTopic object, and an error if there is any.
func (c *FakeSNSTopics) Get(name string, options v1.GetOptions) (result *v1alpha1.SNSTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(snstopicsResource, c.ns, name), &v1alpha1.SNSTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SNSTopic), err
}

// List takes label and field selectors, and returns the list of SNSTopics that match those selectors.
func (c *FakeSNSTopics) List(opts v1.ListOptions) (result *v1alpha1.SNSTopicList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(snstopicsResource, snstopicsKind, c.ns, opts), &v1alpha1.SNSTopicList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SNSTopicList{}
	for _, item := range obj.(*v1alpha1.SNSTopicList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sNSTopics.
func (c *FakeSNSTopics) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(snstopicsResource, c.ns, opts))

}

// Create takes the representation of a sNSTopic and creates it.  Returns the server's representation of the sNSTopic, and an error, if there is any.
func (c *FakeSNSTopics) Create(sNSTopic *v1alpha1.SNSTopic) (result *v1alpha1.SNSTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(snstopicsResource, c.ns, sNSTopic), &v1alpha1.SNSTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SNSTopic), err
}

// Update takes the representation of a sNSTopic and updates it. Returns the server's representation of the sNSTopic, and an error, if there is any.
func (c *FakeSNSTopics) Update(sNSTopic *v1alpha1.SNSTopic) (result *v1alpha1.SNSTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(snstopicsResource, c.ns, sNSTopic), &v1alpha1.SNSTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SNSTopic), err
}

// Delete takes name of the sNSTopic and deletes it. Returns an error if one occurs.
func (c *FakeSNSTopics) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(snstopicsResource, c.ns, name), &v1alpha1.SNSTopic{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSNSTopics) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(snstopicsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SNSTopicList{})
	return err
}

// Patch applies the patch and returns the patched sNSTopic.
func (c *FakeSNSTopics) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SNSTopic, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(snstopicsResource, c.ns, name, data, subresources...), &v1alpha1.SNSTopic{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SNSTopic), err
}
