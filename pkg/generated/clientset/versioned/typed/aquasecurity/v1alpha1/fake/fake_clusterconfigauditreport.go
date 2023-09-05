// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/danielpacak/kube-security-manager/pkg/apis/aquasecurity/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterConfigAuditReports implements ClusterConfigAuditReportInterface
type FakeClusterConfigAuditReports struct {
	Fake *FakeAquasecurityV1alpha1
}

var clusterconfigauditreportsResource = schema.GroupVersionResource{Group: "aquasecurity.github.io", Version: "v1alpha1", Resource: "clusterconfigauditreports"}

var clusterconfigauditreportsKind = schema.GroupVersionKind{Group: "aquasecurity.github.io", Version: "v1alpha1", Kind: "ClusterConfigAuditReport"}

// Get takes name of the clusterConfigAuditReport, and returns the corresponding clusterConfigAuditReport object, and an error if there is any.
func (c *FakeClusterConfigAuditReports) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterConfigAuditReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterconfigauditreportsResource, name), &v1alpha1.ClusterConfigAuditReport{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterConfigAuditReport), err
}

// List takes label and field selectors, and returns the list of ClusterConfigAuditReports that match those selectors.
func (c *FakeClusterConfigAuditReports) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterConfigAuditReportList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterconfigauditreportsResource, clusterconfigauditreportsKind, opts), &v1alpha1.ClusterConfigAuditReportList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterConfigAuditReportList{ListMeta: obj.(*v1alpha1.ClusterConfigAuditReportList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterConfigAuditReportList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterConfigAuditReports.
func (c *FakeClusterConfigAuditReports) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterconfigauditreportsResource, opts))
}

// Create takes the representation of a clusterConfigAuditReport and creates it.  Returns the server's representation of the clusterConfigAuditReport, and an error, if there is any.
func (c *FakeClusterConfigAuditReports) Create(ctx context.Context, clusterConfigAuditReport *v1alpha1.ClusterConfigAuditReport, opts v1.CreateOptions) (result *v1alpha1.ClusterConfigAuditReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterconfigauditreportsResource, clusterConfigAuditReport), &v1alpha1.ClusterConfigAuditReport{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterConfigAuditReport), err
}

// Update takes the representation of a clusterConfigAuditReport and updates it. Returns the server's representation of the clusterConfigAuditReport, and an error, if there is any.
func (c *FakeClusterConfigAuditReports) Update(ctx context.Context, clusterConfigAuditReport *v1alpha1.ClusterConfigAuditReport, opts v1.UpdateOptions) (result *v1alpha1.ClusterConfigAuditReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterconfigauditreportsResource, clusterConfigAuditReport), &v1alpha1.ClusterConfigAuditReport{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterConfigAuditReport), err
}

// Delete takes name of the clusterConfigAuditReport and deletes it. Returns an error if one occurs.
func (c *FakeClusterConfigAuditReports) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterconfigauditreportsResource, name, opts), &v1alpha1.ClusterConfigAuditReport{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterConfigAuditReports) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterconfigauditreportsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterConfigAuditReportList{})
	return err
}

// Patch applies the patch and returns the patched clusterConfigAuditReport.
func (c *FakeClusterConfigAuditReports) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterConfigAuditReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterconfigauditreportsResource, name, pt, data, subresources...), &v1alpha1.ClusterConfigAuditReport{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterConfigAuditReport), err
}
