/*
Copyright The Searchlight Authors.

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
	"context"

	v1alpha1 "go.searchlight.dev/grafana-operator/apis/grafana/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDashboardTemplates implements DashboardTemplateInterface
type FakeDashboardTemplates struct {
	Fake *FakeGrafanaV1alpha1
	ns   string
}

var dashboardtemplatesResource = schema.GroupVersionResource{Group: "grafana.searchlight.dev", Version: "v1alpha1", Resource: "dashboardtemplates"}

var dashboardtemplatesKind = schema.GroupVersionKind{Group: "grafana.searchlight.dev", Version: "v1alpha1", Kind: "DashboardTemplate"}

// Get takes name of the dashboardTemplate, and returns the corresponding dashboardTemplate object, and an error if there is any.
func (c *FakeDashboardTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DashboardTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dashboardtemplatesResource, c.ns, name), &v1alpha1.DashboardTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DashboardTemplate), err
}

// List takes label and field selectors, and returns the list of DashboardTemplates that match those selectors.
func (c *FakeDashboardTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DashboardTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dashboardtemplatesResource, dashboardtemplatesKind, c.ns, opts), &v1alpha1.DashboardTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DashboardTemplateList{ListMeta: obj.(*v1alpha1.DashboardTemplateList).ListMeta}
	for _, item := range obj.(*v1alpha1.DashboardTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dashboardTemplates.
func (c *FakeDashboardTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dashboardtemplatesResource, c.ns, opts))

}

// Create takes the representation of a dashboardTemplate and creates it.  Returns the server's representation of the dashboardTemplate, and an error, if there is any.
func (c *FakeDashboardTemplates) Create(ctx context.Context, dashboardTemplate *v1alpha1.DashboardTemplate, opts v1.CreateOptions) (result *v1alpha1.DashboardTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dashboardtemplatesResource, c.ns, dashboardTemplate), &v1alpha1.DashboardTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DashboardTemplate), err
}

// Update takes the representation of a dashboardTemplate and updates it. Returns the server's representation of the dashboardTemplate, and an error, if there is any.
func (c *FakeDashboardTemplates) Update(ctx context.Context, dashboardTemplate *v1alpha1.DashboardTemplate, opts v1.UpdateOptions) (result *v1alpha1.DashboardTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dashboardtemplatesResource, c.ns, dashboardTemplate), &v1alpha1.DashboardTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DashboardTemplate), err
}

// Delete takes name of the dashboardTemplate and deletes it. Returns an error if one occurs.
func (c *FakeDashboardTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dashboardtemplatesResource, c.ns, name), &v1alpha1.DashboardTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDashboardTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dashboardtemplatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DashboardTemplateList{})
	return err
}

// Patch applies the patch and returns the patched dashboardTemplate.
func (c *FakeDashboardTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DashboardTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dashboardtemplatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DashboardTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DashboardTemplate), err
}
