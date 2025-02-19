// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package railway

import (
	"context"
	"reflect"

	"errors"
	"example.com/pulumi-railway/sdk/go/railway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RandomComponent struct {
	pulumi.ResourceState

	Length   pulumi.IntOutput    `pulumi:"length"`
	Password pulumi.StringOutput `pulumi:"password"`
}

// NewRandomComponent registers a new resource with the given unique name, arguments, and options.
func NewRandomComponent(ctx *pulumi.Context,
	name string, args *RandomComponentArgs, opts ...pulumi.ResourceOption) (*RandomComponent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Length == nil {
		return nil, errors.New("invalid value for required argument 'Length'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RandomComponent
	err := ctx.RegisterRemoteComponentResource("railway:index:RandomComponent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type randomComponentArgs struct {
	Length int `pulumi:"length"`
}

// The set of arguments for constructing a RandomComponent resource.
type RandomComponentArgs struct {
	Length pulumi.IntInput
}

func (RandomComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomComponentArgs)(nil)).Elem()
}

type RandomComponentInput interface {
	pulumi.Input

	ToRandomComponentOutput() RandomComponentOutput
	ToRandomComponentOutputWithContext(ctx context.Context) RandomComponentOutput
}

func (*RandomComponent) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomComponent)(nil)).Elem()
}

func (i *RandomComponent) ToRandomComponentOutput() RandomComponentOutput {
	return i.ToRandomComponentOutputWithContext(context.Background())
}

func (i *RandomComponent) ToRandomComponentOutputWithContext(ctx context.Context) RandomComponentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomComponentOutput)
}

type RandomComponentOutput struct{ *pulumi.OutputState }

func (RandomComponentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomComponent)(nil)).Elem()
}

func (o RandomComponentOutput) ToRandomComponentOutput() RandomComponentOutput {
	return o
}

func (o RandomComponentOutput) ToRandomComponentOutputWithContext(ctx context.Context) RandomComponentOutput {
	return o
}

func (o RandomComponentOutput) Length() pulumi.IntOutput {
	return o.ApplyT(func(v *RandomComponent) pulumi.IntOutput { return v.Length }).(pulumi.IntOutput)
}

func (o RandomComponentOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *RandomComponent) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RandomComponentInput)(nil)).Elem(), &RandomComponent{})
	pulumi.RegisterOutputType(RandomComponentOutput{})
}
