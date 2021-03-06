/*
Copyright 2020 The KubeEdge Authors.

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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/kubeedge/kubeedge/cloud/pkg/apis/devices/v1alpha1"
	scheme "github.com/kubeedge/kubeedge/cloud/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DeviceModelsGetter has a method to return a DeviceModelInterface.
// A group's client should implement this interface.
type DeviceModelsGetter interface {
	DeviceModels(namespace string) DeviceModelInterface
}

// DeviceModelInterface has methods to work with DeviceModel resources.
type DeviceModelInterface interface {
	Create(*v1alpha1.DeviceModel) (*v1alpha1.DeviceModel, error)
	Update(*v1alpha1.DeviceModel) (*v1alpha1.DeviceModel, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DeviceModel, error)
	List(opts v1.ListOptions) (*v1alpha1.DeviceModelList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DeviceModel, err error)
	DeviceModelExpansion
}

// deviceModels implements DeviceModelInterface
type deviceModels struct {
	client rest.Interface
	ns     string
}

// newDeviceModels returns a DeviceModels
func newDeviceModels(c *DevicesV1alpha1Client, namespace string) *deviceModels {
	return &deviceModels{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the deviceModel, and returns the corresponding deviceModel object, and an error if there is any.
func (c *deviceModels) Get(name string, options v1.GetOptions) (result *v1alpha1.DeviceModel, err error) {
	result = &v1alpha1.DeviceModel{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("devicemodels").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DeviceModels that match those selectors.
func (c *deviceModels) List(opts v1.ListOptions) (result *v1alpha1.DeviceModelList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DeviceModelList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("devicemodels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested deviceModels.
func (c *deviceModels) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("devicemodels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a deviceModel and creates it.  Returns the server's representation of the deviceModel, and an error, if there is any.
func (c *deviceModels) Create(deviceModel *v1alpha1.DeviceModel) (result *v1alpha1.DeviceModel, err error) {
	result = &v1alpha1.DeviceModel{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("devicemodels").
		Body(deviceModel).
		Do().
		Into(result)
	return
}

// Update takes the representation of a deviceModel and updates it. Returns the server's representation of the deviceModel, and an error, if there is any.
func (c *deviceModels) Update(deviceModel *v1alpha1.DeviceModel) (result *v1alpha1.DeviceModel, err error) {
	result = &v1alpha1.DeviceModel{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("devicemodels").
		Name(deviceModel.Name).
		Body(deviceModel).
		Do().
		Into(result)
	return
}

// Delete takes name of the deviceModel and deletes it. Returns an error if one occurs.
func (c *deviceModels) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("devicemodels").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *deviceModels) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("devicemodels").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched deviceModel.
func (c *deviceModels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DeviceModel, err error) {
	result = &v1alpha1.DeviceModel{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("devicemodels").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
