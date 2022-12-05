// Code generated by internal/generate/servicepackagedata/main.go; DO NOT EDIT.

package ecrpublic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/experimental/intf"
)

var spd = &servicePackageData{}

func registerFrameworkDataSourceFactory(factory func(context.Context) (datasource.DataSourceWithConfigure, error)) {
	spd.frameworkDataSourceFactories = append(spd.frameworkDataSourceFactories, factory)
}

func registerFrameworkResourceFactory(factory func(context.Context) (resource.ResourceWithConfigure, error)) {
	spd.frameworkResourceFactories = append(spd.frameworkResourceFactories, factory)
}

func registerSDKDataSourceFactory(factory func() *schema.Resource) {
	spd.sdkResourceFactories = append(spd.sdkResourceFactories, factory)
}

func registerSDKResourceFactory(factory func() *schema.Resource) {
	spd.sdkDataSourceFactories = append(spd.sdkDataSourceFactories, factory)
}

type servicePackageData struct {
	frameworkDataSourceFactories []func(context.Context) (datasource.DataSourceWithConfigure, error)
	frameworkResourceFactories   []func(context.Context) (resource.ResourceWithConfigure, error)
	sdkDataSourceFactories       []func() *schema.Resource
	sdkResourceFactories         []func() *schema.Resource
}

func (d *servicePackageData) Configure(ctx context.Context, meta any) error {
	return nil
}

func (d *servicePackageData) FrameworkDataSources(ctx context.Context) []func(context.Context) (datasource.DataSourceWithConfigure, error) {
	return d.frameworkDataSourceFactories
}

func (d *servicePackageData) FrameworkResources(ctx context.Context) []func(context.Context) (resource.ResourceWithConfigure, error) {
	return d.frameworkResourceFactories
}

func (d *servicePackageData) SDKDataSources(ctx context.Context) []func() *schema.Resource {
	return d.sdkDataSourceFactories
}

func (d *servicePackageData) SDKResources(ctx context.Context) []func() *schema.Resource {
	return d.sdkResourceFactories
}

func (d *servicePackageData) ServicePackageName() string {
	return "ecrpublic"
}

var ServicePackageData intf.ServicePackageData = spd
