/*
Copyright 2024.

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
	context "context"
	json "encoding/json"
	fmt "fmt"

	v1alpha1 "inference.networking.x-k8s.io/llm-instance-gateway/api/v1alpha1"
	apiv1alpha1 "inference.networking.x-k8s.io/llm-instance-gateway/client-go/applyconfiguration/api/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLLMServerPools implements LLMServerPoolInterface
type FakeLLMServerPools struct {
	Fake *FakeApiV1alpha1
	ns   string
}

var llmserverpoolsResource = v1alpha1.SchemeGroupVersion.WithResource("llmserverpools")

var llmserverpoolsKind = v1alpha1.SchemeGroupVersion.WithKind("LLMServerPool")

// Get takes name of the lLMServerPool, and returns the corresponding lLMServerPool object, and an error if there is any.
func (c *FakeLLMServerPools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LLMServerPool, err error) {
	emptyResult := &v1alpha1.LLMServerPool{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(llmserverpoolsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.LLMServerPool), err
}

// List takes label and field selectors, and returns the list of LLMServerPools that match those selectors.
func (c *FakeLLMServerPools) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LLMServerPoolList, err error) {
	emptyResult := &v1alpha1.LLMServerPoolList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(llmserverpoolsResource, llmserverpoolsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LLMServerPoolList{ListMeta: obj.(*v1alpha1.LLMServerPoolList).ListMeta}
	for _, item := range obj.(*v1alpha1.LLMServerPoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lLMServerPools.
func (c *FakeLLMServerPools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(llmserverpoolsResource, c.ns, opts))

}

// Create takes the representation of a lLMServerPool and creates it.  Returns the server's representation of the lLMServerPool, and an error, if there is any.
func (c *FakeLLMServerPools) Create(ctx context.Context, lLMServerPool *v1alpha1.LLMServerPool, opts v1.CreateOptions) (result *v1alpha1.LLMServerPool, err error) {
	emptyResult := &v1alpha1.LLMServerPool{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(llmserverpoolsResource, c.ns, lLMServerPool, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.LLMServerPool), err
}

// Update takes the representation of a lLMServerPool and updates it. Returns the server's representation of the lLMServerPool, and an error, if there is any.
func (c *FakeLLMServerPools) Update(ctx context.Context, lLMServerPool *v1alpha1.LLMServerPool, opts v1.UpdateOptions) (result *v1alpha1.LLMServerPool, err error) {
	emptyResult := &v1alpha1.LLMServerPool{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(llmserverpoolsResource, c.ns, lLMServerPool, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.LLMServerPool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLLMServerPools) UpdateStatus(ctx context.Context, lLMServerPool *v1alpha1.LLMServerPool, opts v1.UpdateOptions) (result *v1alpha1.LLMServerPool, err error) {
	emptyResult := &v1alpha1.LLMServerPool{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(llmserverpoolsResource, "status", c.ns, lLMServerPool, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.LLMServerPool), err
}

// Delete takes name of the lLMServerPool and deletes it. Returns an error if one occurs.
func (c *FakeLLMServerPools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(llmserverpoolsResource, c.ns, name, opts), &v1alpha1.LLMServerPool{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLLMServerPools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(llmserverpoolsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LLMServerPoolList{})
	return err
}

// Patch applies the patch and returns the patched lLMServerPool.
func (c *FakeLLMServerPools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LLMServerPool, err error) {
	emptyResult := &v1alpha1.LLMServerPool{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(llmserverpoolsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.LLMServerPool), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied lLMServerPool.
func (c *FakeLLMServerPools) Apply(ctx context.Context, lLMServerPool *apiv1alpha1.LLMServerPoolApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.LLMServerPool, err error) {
	if lLMServerPool == nil {
		return nil, fmt.Errorf("lLMServerPool provided to Apply must not be nil")
	}
	data, err := json.Marshal(lLMServerPool)
	if err != nil {
		return nil, err
	}
	name := lLMServerPool.Name
	if name == nil {
		return nil, fmt.Errorf("lLMServerPool.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.LLMServerPool{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(llmserverpoolsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.LLMServerPool), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeLLMServerPools) ApplyStatus(ctx context.Context, lLMServerPool *apiv1alpha1.LLMServerPoolApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.LLMServerPool, err error) {
	if lLMServerPool == nil {
		return nil, fmt.Errorf("lLMServerPool provided to Apply must not be nil")
	}
	data, err := json.Marshal(lLMServerPool)
	if err != nil {
		return nil, err
	}
	name := lLMServerPool.Name
	if name == nil {
		return nil, fmt.Errorf("lLMServerPool.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.LLMServerPool{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(llmserverpoolsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.LLMServerPool), err
}
