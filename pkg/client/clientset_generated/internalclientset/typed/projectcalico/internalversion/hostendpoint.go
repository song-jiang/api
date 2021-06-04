// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	projectcalico "github.com/projectcalico/api/pkg/apis/projectcalico"
	scheme "github.com/projectcalico/api/pkg/client/clientset_generated/internalclientset/scheme"
)

// HostEndpointsGetter has a method to return a HostEndpointInterface.
// A group's client should implement this interface.
type HostEndpointsGetter interface {
	HostEndpoints() HostEndpointInterface
}

// HostEndpointInterface has methods to work with HostEndpoint resources.
type HostEndpointInterface interface {
	Create(ctx context.Context, hostEndpoint *projectcalico.HostEndpoint, opts v1.CreateOptions) (*projectcalico.HostEndpoint, error)
	Update(ctx context.Context, hostEndpoint *projectcalico.HostEndpoint, opts v1.UpdateOptions) (*projectcalico.HostEndpoint, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*projectcalico.HostEndpoint, error)
	List(ctx context.Context, opts v1.ListOptions) (*projectcalico.HostEndpointList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *projectcalico.HostEndpoint, err error)
	HostEndpointExpansion
}

// hostEndpoints implements HostEndpointInterface
type hostEndpoints struct {
	client rest.Interface
}

// newHostEndpoints returns a HostEndpoints
func newHostEndpoints(c *ProjectcalicoClient) *hostEndpoints {
	return &hostEndpoints{
		client: c.RESTClient(),
	}
}

// Get takes name of the hostEndpoint, and returns the corresponding hostEndpoint object, and an error if there is any.
func (c *hostEndpoints) Get(ctx context.Context, name string, options v1.GetOptions) (result *projectcalico.HostEndpoint, err error) {
	result = &projectcalico.HostEndpoint{}
	err = c.client.Get().
		Resource("hostendpoints").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HostEndpoints that match those selectors.
func (c *hostEndpoints) List(ctx context.Context, opts v1.ListOptions) (result *projectcalico.HostEndpointList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &projectcalico.HostEndpointList{}
	err = c.client.Get().
		Resource("hostendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hostEndpoints.
func (c *hostEndpoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("hostendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a hostEndpoint and creates it.  Returns the server's representation of the hostEndpoint, and an error, if there is any.
func (c *hostEndpoints) Create(ctx context.Context, hostEndpoint *projectcalico.HostEndpoint, opts v1.CreateOptions) (result *projectcalico.HostEndpoint, err error) {
	result = &projectcalico.HostEndpoint{}
	err = c.client.Post().
		Resource("hostendpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hostEndpoint).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a hostEndpoint and updates it. Returns the server's representation of the hostEndpoint, and an error, if there is any.
func (c *hostEndpoints) Update(ctx context.Context, hostEndpoint *projectcalico.HostEndpoint, opts v1.UpdateOptions) (result *projectcalico.HostEndpoint, err error) {
	result = &projectcalico.HostEndpoint{}
	err = c.client.Put().
		Resource("hostendpoints").
		Name(hostEndpoint.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hostEndpoint).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the hostEndpoint and deletes it. Returns an error if one occurs.
func (c *hostEndpoints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("hostendpoints").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hostEndpoints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("hostendpoints").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched hostEndpoint.
func (c *hostEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *projectcalico.HostEndpoint, err error) {
	result = &projectcalico.HostEndpoint{}
	err = c.client.Patch(pt).
		Resource("hostendpoints").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}