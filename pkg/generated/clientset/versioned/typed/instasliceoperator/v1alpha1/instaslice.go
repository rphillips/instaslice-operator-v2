/*
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

package v1alpha1

import (
	context "context"

	instasliceoperatorv1alpha1 "github.com/openshift/instaslice-operator/pkg/apis/instasliceoperator/v1alpha1"
	applyconfigurationinstasliceoperatorv1alpha1 "github.com/openshift/instaslice-operator/pkg/generated/applyconfiguration/instasliceoperator/v1alpha1"
	scheme "github.com/openshift/instaslice-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// InstaslicesGetter has a method to return a InstasliceInterface.
// A group's client should implement this interface.
type InstaslicesGetter interface {
	Instaslices(namespace string) InstasliceInterface
}

// InstasliceInterface has methods to work with Instaslice resources.
type InstasliceInterface interface {
	Create(ctx context.Context, instaslice *instasliceoperatorv1alpha1.Instaslice, opts v1.CreateOptions) (*instasliceoperatorv1alpha1.Instaslice, error)
	Update(ctx context.Context, instaslice *instasliceoperatorv1alpha1.Instaslice, opts v1.UpdateOptions) (*instasliceoperatorv1alpha1.Instaslice, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, instaslice *instasliceoperatorv1alpha1.Instaslice, opts v1.UpdateOptions) (*instasliceoperatorv1alpha1.Instaslice, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*instasliceoperatorv1alpha1.Instaslice, error)
	List(ctx context.Context, opts v1.ListOptions) (*instasliceoperatorv1alpha1.InstasliceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *instasliceoperatorv1alpha1.Instaslice, err error)
	Apply(ctx context.Context, instaslice *applyconfigurationinstasliceoperatorv1alpha1.InstasliceApplyConfiguration, opts v1.ApplyOptions) (result *instasliceoperatorv1alpha1.Instaslice, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, instaslice *applyconfigurationinstasliceoperatorv1alpha1.InstasliceApplyConfiguration, opts v1.ApplyOptions) (result *instasliceoperatorv1alpha1.Instaslice, err error)
	InstasliceExpansion
}

// instaslices implements InstasliceInterface
type instaslices struct {
	*gentype.ClientWithListAndApply[*instasliceoperatorv1alpha1.Instaslice, *instasliceoperatorv1alpha1.InstasliceList, *applyconfigurationinstasliceoperatorv1alpha1.InstasliceApplyConfiguration]
}

// newInstaslices returns a Instaslices
func newInstaslices(c *OpenShiftOperatorV1alpha1Client, namespace string) *instaslices {
	return &instaslices{
		gentype.NewClientWithListAndApply[*instasliceoperatorv1alpha1.Instaslice, *instasliceoperatorv1alpha1.InstasliceList, *applyconfigurationinstasliceoperatorv1alpha1.InstasliceApplyConfiguration](
			"instaslices",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *instasliceoperatorv1alpha1.Instaslice { return &instasliceoperatorv1alpha1.Instaslice{} },
			func() *instasliceoperatorv1alpha1.InstasliceList { return &instasliceoperatorv1alpha1.InstasliceList{} },
		),
	}
}
