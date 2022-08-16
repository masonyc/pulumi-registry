// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package registry

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the registry package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	AzureClientId     pulumi.StringOutput `pulumi:"azureClientId"`
	AzureClientSecret pulumi.StringOutput `pulumi:"azureClientSecret"`
	AzureTenantId     pulumi.StringOutput `pulumi:"azureTenantId"`
	RegistryBaseUrl   pulumi.StringOutput `pulumi:"registryBaseUrl"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AzureClientId == nil {
		return nil, errors.New("invalid value for required argument 'AzureClientId'")
	}
	if args.AzureClientSecret == nil {
		return nil, errors.New("invalid value for required argument 'AzureClientSecret'")
	}
	if args.AzureTenantId == nil {
		return nil, errors.New("invalid value for required argument 'AzureTenantId'")
	}
	if args.RegistryBaseUrl == nil {
		return nil, errors.New("invalid value for required argument 'RegistryBaseUrl'")
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:registry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	AzureClientId     string `pulumi:"azureClientId"`
	AzureClientSecret string `pulumi:"azureClientSecret"`
	AzureTenantId     string `pulumi:"azureTenantId"`
	RegistryBaseUrl   string `pulumi:"registryBaseUrl"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	AzureClientId     pulumi.StringInput
	AzureClientSecret pulumi.StringInput
	AzureTenantId     pulumi.StringInput
	RegistryBaseUrl   pulumi.StringInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func (o ProviderOutput) AzureClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.AzureClientId }).(pulumi.StringOutput)
}

func (o ProviderOutput) AzureClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.AzureClientSecret }).(pulumi.StringOutput)
}

func (o ProviderOutput) AzureTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.AzureTenantId }).(pulumi.StringOutput)
}

func (o ProviderOutput) RegistryBaseUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.RegistryBaseUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}