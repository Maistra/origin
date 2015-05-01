package client

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/fields"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/labels"

	authorizationapi "github.com/openshift/origin/pkg/authorization/api"
)

type ClusterRoleBindingsInterface interface {
	ClusterRoleBindings() ClusterRoleBindingInterface
}

type ClusterRoleBindingInterface interface {
	List(label labels.Selector, field fields.Selector) (*authorizationapi.ClusterRoleBindingList, error)
	Get(name string) (*authorizationapi.ClusterRoleBinding, error)
	Create(roleBinding *authorizationapi.ClusterRoleBinding) (*authorizationapi.ClusterRoleBinding, error)
	Delete(name string) error
}

type clusterRoleBindings struct {
	r *Client
}

// newClusterRoleBindings returns a clusterRoleBindings
func newClusterRoleBindings(c *Client) *clusterRoleBindings {
	return &clusterRoleBindings{
		r: c,
	}
}

// List returns a list of clusterRoleBindings that match the label and field selectors.
func (c *clusterRoleBindings) List(label labels.Selector, field fields.Selector) (result *authorizationapi.ClusterRoleBindingList, err error) {
	result = &authorizationapi.ClusterRoleBindingList{}
	err = c.r.Get().Resource("clusterRoleBindings").LabelsSelectorParam(label).FieldsSelectorParam(field).Do().Into(result)
	return
}

// Get returns information about a particular roleBinding and error if one occurs.
func (c *clusterRoleBindings) Get(name string) (result *authorizationapi.ClusterRoleBinding, err error) {
	result = &authorizationapi.ClusterRoleBinding{}
	err = c.r.Get().Resource("clusterRoleBindings").Name(name).Do().Into(result)
	return
}

// Create creates new roleBinding. Returns the server's representation of the roleBinding and error if one occurs.
func (c *clusterRoleBindings) Create(roleBinding *authorizationapi.ClusterRoleBinding) (result *authorizationapi.ClusterRoleBinding, err error) {
	result = &authorizationapi.ClusterRoleBinding{}
	err = c.r.Post().Resource("clusterRoleBindings").Body(roleBinding).Do().Into(result)
	return
}

// Delete deletes a roleBinding, returns error if one occurs.
func (c *clusterRoleBindings) Delete(name string) (err error) {
	err = c.r.Delete().Resource("clusterRoleBindings").Name(name).Do().Error()
	return
}
