// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/contiv/vpp/plugins/crd/pkg/apis/bgpconfig/v1"
	scheme "github.com/contiv/vpp/plugins/crd/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BgpConfigsGetter has a method to return a BgpConfigInterface.
// A group's client should implement this interface.
type BgpConfigsGetter interface {
	BgpConfigs(namespace string) BgpConfigInterface
}

// BgpConfigInterface has methods to work with BgpConfig resources.
type BgpConfigInterface interface {
	Create(*v1.BgpConfig) (*v1.BgpConfig, error)
	Update(*v1.BgpConfig) (*v1.BgpConfig, error)
	UpdateStatus(*v1.BgpConfig) (*v1.BgpConfig, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.BgpConfig, error)
	List(opts metav1.ListOptions) (*v1.BgpConfigList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.BgpConfig, err error)
	BgpConfigExpansion
}

// bgpConfigs implements BgpConfigInterface
type bgpConfigs struct {
	client rest.Interface
	ns     string
}

// newBgpConfigs returns a BgpConfigs
func newBgpConfigs(c *BgpconfigV1Client, namespace string) *bgpConfigs {
	return &bgpConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bgpConfig, and returns the corresponding bgpConfig object, and an error if there is any.
func (c *bgpConfigs) Get(name string, options metav1.GetOptions) (result *v1.BgpConfig, err error) {
	result = &v1.BgpConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bgpconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BgpConfigs that match those selectors.
func (c *bgpConfigs) List(opts metav1.ListOptions) (result *v1.BgpConfigList, err error) {
	result = &v1.BgpConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bgpconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bgpConfigs.
func (c *bgpConfigs) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bgpconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a bgpConfig and creates it.  Returns the server's representation of the bgpConfig, and an error, if there is any.
func (c *bgpConfigs) Create(bgpConfig *v1.BgpConfig) (result *v1.BgpConfig, err error) {
	result = &v1.BgpConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bgpconfigs").
		Body(bgpConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a bgpConfig and updates it. Returns the server's representation of the bgpConfig, and an error, if there is any.
func (c *bgpConfigs) Update(bgpConfig *v1.BgpConfig) (result *v1.BgpConfig, err error) {
	result = &v1.BgpConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bgpconfigs").
		Name(bgpConfig.Name).
		Body(bgpConfig).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *bgpConfigs) UpdateStatus(bgpConfig *v1.BgpConfig) (result *v1.BgpConfig, err error) {
	result = &v1.BgpConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bgpconfigs").
		Name(bgpConfig.Name).
		SubResource("status").
		Body(bgpConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the bgpConfig and deletes it. Returns an error if one occurs.
func (c *bgpConfigs) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bgpconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bgpConfigs) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bgpconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched bgpConfig.
func (c *bgpConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.BgpConfig, err error) {
	result = &v1.BgpConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bgpconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
