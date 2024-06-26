package client

import (
	"context"

	"github.com/glasskube/glasskube/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

var packageGVR = v1alpha1.GroupVersion.WithResource("packages")

type packageClient struct {
	restClient rest.Interface
}

func (c *packageClient) Create(ctx context.Context, pkg *v1alpha1.Package, opts metav1.CreateOptions) error {
	return c.restClient.Post().
		Resource(packageGVR.Resource).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(pkg).Do(ctx).Into(pkg)
}

// Update implements PackageInterface.
func (c *packageClient) Update(ctx context.Context, p *v1alpha1.Package) error {
	return c.restClient.Put().
		Resource(packageGVR.Resource).
		Name(p.GetName()).
		Body(p).
		Do(ctx).
		Into(p)
}

func (c *packageClient) Watch(ctx context.Context) (watch.Interface, error) {
	opts := metav1.ListOptions{Watch: true}
	return c.restClient.Get().
		Resource(packageGVR.Resource).
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(ctx)
}

func (c *packageClient) Get(ctx context.Context, pkgName string, result *v1alpha1.Package) error {
	return c.restClient.Get().
		Resource(packageGVR.Resource).
		Name(pkgName).
		Do(ctx).Into(result)
}

func (c *packageClient) GetAll(ctx context.Context, result *v1alpha1.PackageList) error {
	return c.restClient.Get().
		Resource(packageGVR.Resource).
		Do(ctx).Into(result)
}

func (c *packageClient) Delete(ctx context.Context, pkg *v1alpha1.Package, options metav1.DeleteOptions) error {
	return c.restClient.Delete().
		Resource(packageGVR.Resource).
		Name(pkg.Name).
		Body(&options).
		Do(ctx).Into(nil)
}

// NewPackage instantiates a new v1alpha1.Package struct with the given package name
func NewPackage(packageName, version string) *v1alpha1.Package {
	return &v1alpha1.Package{
		ObjectMeta: metav1.ObjectMeta{
			Name: packageName,
		},
		Spec: v1alpha1.PackageSpec{
			PackageInfo: v1alpha1.PackageInfoTemplate{
				Name:    packageName,
				Version: version,
			},
		},
	}
}
