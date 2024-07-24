/*
Copyright 2021 The Knative Authors

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

package v1

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1 "knative.dev/eventing/pkg/apis/messaging/v1"
	scheme "knative.dev/eventing/pkg/client/clientset/versioned/scheme"
)

// ChannelsGetter has a method to return a ChannelInterface.
// A group's client should implement this interface.
type ChannelsGetter interface {
	Channels(namespace string) ChannelInterface
}

// ChannelInterface has methods to work with Channel resources.
type ChannelInterface interface {
	Create(ctx context.Context, channel *v1.Channel, opts metav1.CreateOptions) (*v1.Channel, error)
	Update(ctx context.Context, channel *v1.Channel, opts metav1.UpdateOptions) (*v1.Channel, error)
	UpdateStatus(ctx context.Context, channel *v1.Channel, opts metav1.UpdateOptions) (*v1.Channel, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Channel, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ChannelList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Channel, err error)
	ChannelExpansion
}

// channels implements ChannelInterface
type channels struct {
	client rest.Interface
	ns     string
}

// newChannels returns a Channels
func newChannels(c *MessagingV1Client, namespace string) *channels {
	return &channels{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the channel, and returns the corresponding channel object, and an error if there is any.
func (c *channels) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Channel, err error) {
	result = &v1.Channel{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("channels").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Channels that match those selectors.
func (c *channels) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ChannelList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ChannelList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("channels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested channels.
func (c *channels) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("channels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a channel and creates it.  Returns the server's representation of the channel, and an error, if there is any.
func (c *channels) Create(ctx context.Context, channel *v1.Channel, opts metav1.CreateOptions) (result *v1.Channel, err error) {
	result = &v1.Channel{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("channels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(channel).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a channel and updates it. Returns the server's representation of the channel, and an error, if there is any.
func (c *channels) Update(ctx context.Context, channel *v1.Channel, opts metav1.UpdateOptions) (result *v1.Channel, err error) {
	result = &v1.Channel{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("channels").
		Name(channel.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(channel).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *channels) UpdateStatus(ctx context.Context, channel *v1.Channel, opts metav1.UpdateOptions) (result *v1.Channel, err error) {
	result = &v1.Channel{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("channels").
		Name(channel.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(channel).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the channel and deletes it. Returns an error if one occurs.
func (c *channels) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("channels").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *channels) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("channels").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched channel.
func (c *channels) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Channel, err error) {
	result = &v1.Channel{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("channels").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
