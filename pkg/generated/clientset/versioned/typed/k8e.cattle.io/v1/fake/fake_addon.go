/*
Copyright The Kubernetes Authors.

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
// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	k8ecattleiov1 "github.com/xiaods/k8e/pkg/apis/k8e.cattle.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAddons implements AddonInterface
type FakeAddons struct {
	Fake *FakeK8eV1
	ns   string
}

var addonsResource = schema.GroupVersionResource{Group: "k8e.cattle.io", Version: "v1", Resource: "addons"}

var addonsKind = schema.GroupVersionKind{Group: "k8e.cattle.io", Version: "v1", Kind: "Addon"}

// Get takes name of the addon, and returns the corresponding addon object, and an error if there is any.
func (c *FakeAddons) Get(ctx context.Context, name string, options v1.GetOptions) (result *k8ecattleiov1.Addon, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(addonsResource, c.ns, name), &k8ecattleiov1.Addon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k8ecattleiov1.Addon), err
}

// List takes label and field selectors, and returns the list of Addons that match those selectors.
func (c *FakeAddons) List(ctx context.Context, opts v1.ListOptions) (result *k8ecattleiov1.AddonList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(addonsResource, addonsKind, c.ns, opts), &k8ecattleiov1.AddonList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &k8ecattleiov1.AddonList{ListMeta: obj.(*k8ecattleiov1.AddonList).ListMeta}
	for _, item := range obj.(*k8ecattleiov1.AddonList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested addons.
func (c *FakeAddons) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(addonsResource, c.ns, opts))

}

// Create takes the representation of a addon and creates it.  Returns the server's representation of the addon, and an error, if there is any.
func (c *FakeAddons) Create(ctx context.Context, addon *k8ecattleiov1.Addon, opts v1.CreateOptions) (result *k8ecattleiov1.Addon, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(addonsResource, c.ns, addon), &k8ecattleiov1.Addon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k8ecattleiov1.Addon), err
}

// Update takes the representation of a addon and updates it. Returns the server's representation of the addon, and an error, if there is any.
func (c *FakeAddons) Update(ctx context.Context, addon *k8ecattleiov1.Addon, opts v1.UpdateOptions) (result *k8ecattleiov1.Addon, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(addonsResource, c.ns, addon), &k8ecattleiov1.Addon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k8ecattleiov1.Addon), err
}

// Delete takes name of the addon and deletes it. Returns an error if one occurs.
func (c *FakeAddons) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(addonsResource, c.ns, name, opts), &k8ecattleiov1.Addon{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAddons) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(addonsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &k8ecattleiov1.AddonList{})
	return err
}

// Patch applies the patch and returns the patched addon.
func (c *FakeAddons) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *k8ecattleiov1.Addon, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(addonsResource, c.ns, name, pt, data, subresources...), &k8ecattleiov1.Addon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k8ecattleiov1.Addon), err
}
