// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "istio.io/client-go/pkg/apis/authentication/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMeshPolicies implements MeshPolicyInterface
type FakeMeshPolicies struct {
	Fake *FakeAuthenticationV1alpha1
}

var meshpoliciesResource = schema.GroupVersionResource{Group: "authentication", Version: "v1alpha1", Resource: "meshpolicies"}

var meshpoliciesKind = schema.GroupVersionKind{Group: "authentication", Version: "v1alpha1", Kind: "MeshPolicy"}

// Get takes name of the meshPolicy, and returns the corresponding meshPolicy object, and an error if there is any.
func (c *FakeMeshPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.MeshPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(meshpoliciesResource, name), &v1alpha1.MeshPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MeshPolicy), err
}

// List takes label and field selectors, and returns the list of MeshPolicies that match those selectors.
func (c *FakeMeshPolicies) List(opts v1.ListOptions) (result *v1alpha1.MeshPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(meshpoliciesResource, meshpoliciesKind, opts), &v1alpha1.MeshPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MeshPolicyList{ListMeta: obj.(*v1alpha1.MeshPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.MeshPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested meshPolicies.
func (c *FakeMeshPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(meshpoliciesResource, opts))
}

// Create takes the representation of a meshPolicy and creates it.  Returns the server's representation of the meshPolicy, and an error, if there is any.
func (c *FakeMeshPolicies) Create(meshPolicy *v1alpha1.MeshPolicy) (result *v1alpha1.MeshPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(meshpoliciesResource, meshPolicy), &v1alpha1.MeshPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MeshPolicy), err
}

// Update takes the representation of a meshPolicy and updates it. Returns the server's representation of the meshPolicy, and an error, if there is any.
func (c *FakeMeshPolicies) Update(meshPolicy *v1alpha1.MeshPolicy) (result *v1alpha1.MeshPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(meshpoliciesResource, meshPolicy), &v1alpha1.MeshPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MeshPolicy), err
}

// Delete takes name of the meshPolicy and deletes it. Returns an error if one occurs.
func (c *FakeMeshPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(meshpoliciesResource, name), &v1alpha1.MeshPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMeshPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(meshpoliciesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MeshPolicyList{})
	return err
}

// Patch applies the patch and returns the patched meshPolicy.
func (c *FakeMeshPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MeshPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(meshpoliciesResource, name, pt, data, subresources...), &v1alpha1.MeshPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MeshPolicy), err
}