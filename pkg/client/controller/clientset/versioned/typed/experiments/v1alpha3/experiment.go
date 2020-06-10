/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha3

import (
	v1alpha3 "github.com/kubeflow/katib/pkg/apis/controller/experiments/v1alpha3"
	scheme "github.com/kubeflow/katib/pkg/client/controller/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ExperimentsGetter has a method to return a ExperimentInterface.
// A group's client should implement this interface.
type ExperimentsGetter interface {
	Experiments(namespace string) ExperimentInterface
}

// ExperimentInterface has methods to work with Experiment resources.
type ExperimentInterface interface {
	Create(*v1alpha3.Experiment) (*v1alpha3.Experiment, error)
	Update(*v1alpha3.Experiment) (*v1alpha3.Experiment, error)
	UpdateStatus(*v1alpha3.Experiment) (*v1alpha3.Experiment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha3.Experiment, error)
	List(opts v1.ListOptions) (*v1alpha3.ExperimentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha3.Experiment, err error)
	ExperimentExpansion
}

// experiments implements ExperimentInterface
type experiments struct {
	client rest.Interface
	ns     string
}

// newExperiments returns a Experiments
func newExperiments(c *ExperimentV1alpha3Client, namespace string) *experiments {
	return &experiments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the experiment, and returns the corresponding experiment object, and an error if there is any.
func (c *experiments) Get(name string, options v1.GetOptions) (result *v1alpha3.Experiment, err error) {
	result = &v1alpha3.Experiment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("experiments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Experiments that match those selectors.
func (c *experiments) List(opts v1.ListOptions) (result *v1alpha3.ExperimentList, err error) {
	result = &v1alpha3.ExperimentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("experiments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested experiments.
func (c *experiments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("experiments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a experiment and creates it.  Returns the server's representation of the experiment, and an error, if there is any.
func (c *experiments) Create(experiment *v1alpha3.Experiment) (result *v1alpha3.Experiment, err error) {
	result = &v1alpha3.Experiment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("experiments").
		Body(experiment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a experiment and updates it. Returns the server's representation of the experiment, and an error, if there is any.
func (c *experiments) Update(experiment *v1alpha3.Experiment) (result *v1alpha3.Experiment, err error) {
	result = &v1alpha3.Experiment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("experiments").
		Name(experiment.Name).
		Body(experiment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *experiments) UpdateStatus(experiment *v1alpha3.Experiment) (result *v1alpha3.Experiment, err error) {
	result = &v1alpha3.Experiment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("experiments").
		Name(experiment.Name).
		SubResource("status").
		Body(experiment).
		Do().
		Into(result)
	return
}

// Delete takes name of the experiment and deletes it. Returns an error if one occurs.
func (c *experiments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("experiments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *experiments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("experiments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched experiment.
func (c *experiments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha3.Experiment, err error) {
	result = &v1alpha3.Experiment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("experiments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
