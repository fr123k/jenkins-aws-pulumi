// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Adds launch permission to Amazon Machine Image (AMI) from another AWS account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2.NewAmiLaunchPermission(ctx, "example", &ec2.AmiLaunchPermissionArgs{
// 			AccountId: pulumi.String("123456789012"),
// 			ImageId:   pulumi.String("ami-12345678"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// AWS AMI Launch Permission can be imported using the `ACCOUNT-ID/IMAGE-ID`, e.g.
//
// ```sh
//  $ pulumi import aws:ec2/amiLaunchPermission:AmiLaunchPermission example 123456789012/ami-12345678
// ```
type AmiLaunchPermission struct {
	pulumi.CustomResourceState

	// An AWS Account ID to add launch permissions.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// A region-unique name for the AMI.
	ImageId pulumi.StringOutput `pulumi:"imageId"`
}

// NewAmiLaunchPermission registers a new resource with the given unique name, arguments, and options.
func NewAmiLaunchPermission(ctx *pulumi.Context,
	name string, args *AmiLaunchPermissionArgs, opts ...pulumi.ResourceOption) (*AmiLaunchPermission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountId == nil {
		return nil, errors.New("invalid value for required argument 'AccountId'")
	}
	if args.ImageId == nil {
		return nil, errors.New("invalid value for required argument 'ImageId'")
	}
	var resource AmiLaunchPermission
	err := ctx.RegisterResource("aws:ec2/amiLaunchPermission:AmiLaunchPermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAmiLaunchPermission gets an existing AmiLaunchPermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAmiLaunchPermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AmiLaunchPermissionState, opts ...pulumi.ResourceOption) (*AmiLaunchPermission, error) {
	var resource AmiLaunchPermission
	err := ctx.ReadResource("aws:ec2/amiLaunchPermission:AmiLaunchPermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AmiLaunchPermission resources.
type amiLaunchPermissionState struct {
	// An AWS Account ID to add launch permissions.
	AccountId *string `pulumi:"accountId"`
	// A region-unique name for the AMI.
	ImageId *string `pulumi:"imageId"`
}

type AmiLaunchPermissionState struct {
	// An AWS Account ID to add launch permissions.
	AccountId pulumi.StringPtrInput
	// A region-unique name for the AMI.
	ImageId pulumi.StringPtrInput
}

func (AmiLaunchPermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*amiLaunchPermissionState)(nil)).Elem()
}

type amiLaunchPermissionArgs struct {
	// An AWS Account ID to add launch permissions.
	AccountId string `pulumi:"accountId"`
	// A region-unique name for the AMI.
	ImageId string `pulumi:"imageId"`
}

// The set of arguments for constructing a AmiLaunchPermission resource.
type AmiLaunchPermissionArgs struct {
	// An AWS Account ID to add launch permissions.
	AccountId pulumi.StringInput
	// A region-unique name for the AMI.
	ImageId pulumi.StringInput
}

func (AmiLaunchPermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*amiLaunchPermissionArgs)(nil)).Elem()
}

type AmiLaunchPermissionInput interface {
	pulumi.Input

	ToAmiLaunchPermissionOutput() AmiLaunchPermissionOutput
	ToAmiLaunchPermissionOutputWithContext(ctx context.Context) AmiLaunchPermissionOutput
}

func (*AmiLaunchPermission) ElementType() reflect.Type {
	return reflect.TypeOf((*AmiLaunchPermission)(nil))
}

func (i *AmiLaunchPermission) ToAmiLaunchPermissionOutput() AmiLaunchPermissionOutput {
	return i.ToAmiLaunchPermissionOutputWithContext(context.Background())
}

func (i *AmiLaunchPermission) ToAmiLaunchPermissionOutputWithContext(ctx context.Context) AmiLaunchPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmiLaunchPermissionOutput)
}

func (i *AmiLaunchPermission) ToAmiLaunchPermissionPtrOutput() AmiLaunchPermissionPtrOutput {
	return i.ToAmiLaunchPermissionPtrOutputWithContext(context.Background())
}

func (i *AmiLaunchPermission) ToAmiLaunchPermissionPtrOutputWithContext(ctx context.Context) AmiLaunchPermissionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmiLaunchPermissionPtrOutput)
}

type AmiLaunchPermissionPtrInput interface {
	pulumi.Input

	ToAmiLaunchPermissionPtrOutput() AmiLaunchPermissionPtrOutput
	ToAmiLaunchPermissionPtrOutputWithContext(ctx context.Context) AmiLaunchPermissionPtrOutput
}

type amiLaunchPermissionPtrType AmiLaunchPermissionArgs

func (*amiLaunchPermissionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AmiLaunchPermission)(nil))
}

func (i *amiLaunchPermissionPtrType) ToAmiLaunchPermissionPtrOutput() AmiLaunchPermissionPtrOutput {
	return i.ToAmiLaunchPermissionPtrOutputWithContext(context.Background())
}

func (i *amiLaunchPermissionPtrType) ToAmiLaunchPermissionPtrOutputWithContext(ctx context.Context) AmiLaunchPermissionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmiLaunchPermissionPtrOutput)
}

// AmiLaunchPermissionArrayInput is an input type that accepts AmiLaunchPermissionArray and AmiLaunchPermissionArrayOutput values.
// You can construct a concrete instance of `AmiLaunchPermissionArrayInput` via:
//
//          AmiLaunchPermissionArray{ AmiLaunchPermissionArgs{...} }
type AmiLaunchPermissionArrayInput interface {
	pulumi.Input

	ToAmiLaunchPermissionArrayOutput() AmiLaunchPermissionArrayOutput
	ToAmiLaunchPermissionArrayOutputWithContext(context.Context) AmiLaunchPermissionArrayOutput
}

type AmiLaunchPermissionArray []AmiLaunchPermissionInput

func (AmiLaunchPermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AmiLaunchPermission)(nil)).Elem()
}

func (i AmiLaunchPermissionArray) ToAmiLaunchPermissionArrayOutput() AmiLaunchPermissionArrayOutput {
	return i.ToAmiLaunchPermissionArrayOutputWithContext(context.Background())
}

func (i AmiLaunchPermissionArray) ToAmiLaunchPermissionArrayOutputWithContext(ctx context.Context) AmiLaunchPermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmiLaunchPermissionArrayOutput)
}

// AmiLaunchPermissionMapInput is an input type that accepts AmiLaunchPermissionMap and AmiLaunchPermissionMapOutput values.
// You can construct a concrete instance of `AmiLaunchPermissionMapInput` via:
//
//          AmiLaunchPermissionMap{ "key": AmiLaunchPermissionArgs{...} }
type AmiLaunchPermissionMapInput interface {
	pulumi.Input

	ToAmiLaunchPermissionMapOutput() AmiLaunchPermissionMapOutput
	ToAmiLaunchPermissionMapOutputWithContext(context.Context) AmiLaunchPermissionMapOutput
}

type AmiLaunchPermissionMap map[string]AmiLaunchPermissionInput

func (AmiLaunchPermissionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AmiLaunchPermission)(nil)).Elem()
}

func (i AmiLaunchPermissionMap) ToAmiLaunchPermissionMapOutput() AmiLaunchPermissionMapOutput {
	return i.ToAmiLaunchPermissionMapOutputWithContext(context.Background())
}

func (i AmiLaunchPermissionMap) ToAmiLaunchPermissionMapOutputWithContext(ctx context.Context) AmiLaunchPermissionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmiLaunchPermissionMapOutput)
}

type AmiLaunchPermissionOutput struct{ *pulumi.OutputState }

func (AmiLaunchPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AmiLaunchPermission)(nil))
}

func (o AmiLaunchPermissionOutput) ToAmiLaunchPermissionOutput() AmiLaunchPermissionOutput {
	return o
}

func (o AmiLaunchPermissionOutput) ToAmiLaunchPermissionOutputWithContext(ctx context.Context) AmiLaunchPermissionOutput {
	return o
}

func (o AmiLaunchPermissionOutput) ToAmiLaunchPermissionPtrOutput() AmiLaunchPermissionPtrOutput {
	return o.ToAmiLaunchPermissionPtrOutputWithContext(context.Background())
}

func (o AmiLaunchPermissionOutput) ToAmiLaunchPermissionPtrOutputWithContext(ctx context.Context) AmiLaunchPermissionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AmiLaunchPermission) *AmiLaunchPermission {
		return &v
	}).(AmiLaunchPermissionPtrOutput)
}

type AmiLaunchPermissionPtrOutput struct{ *pulumi.OutputState }

func (AmiLaunchPermissionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AmiLaunchPermission)(nil))
}

func (o AmiLaunchPermissionPtrOutput) ToAmiLaunchPermissionPtrOutput() AmiLaunchPermissionPtrOutput {
	return o
}

func (o AmiLaunchPermissionPtrOutput) ToAmiLaunchPermissionPtrOutputWithContext(ctx context.Context) AmiLaunchPermissionPtrOutput {
	return o
}

func (o AmiLaunchPermissionPtrOutput) Elem() AmiLaunchPermissionOutput {
	return o.ApplyT(func(v *AmiLaunchPermission) AmiLaunchPermission {
		if v != nil {
			return *v
		}
		var ret AmiLaunchPermission
		return ret
	}).(AmiLaunchPermissionOutput)
}

type AmiLaunchPermissionArrayOutput struct{ *pulumi.OutputState }

func (AmiLaunchPermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AmiLaunchPermission)(nil))
}

func (o AmiLaunchPermissionArrayOutput) ToAmiLaunchPermissionArrayOutput() AmiLaunchPermissionArrayOutput {
	return o
}

func (o AmiLaunchPermissionArrayOutput) ToAmiLaunchPermissionArrayOutputWithContext(ctx context.Context) AmiLaunchPermissionArrayOutput {
	return o
}

func (o AmiLaunchPermissionArrayOutput) Index(i pulumi.IntInput) AmiLaunchPermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AmiLaunchPermission {
		return vs[0].([]AmiLaunchPermission)[vs[1].(int)]
	}).(AmiLaunchPermissionOutput)
}

type AmiLaunchPermissionMapOutput struct{ *pulumi.OutputState }

func (AmiLaunchPermissionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AmiLaunchPermission)(nil))
}

func (o AmiLaunchPermissionMapOutput) ToAmiLaunchPermissionMapOutput() AmiLaunchPermissionMapOutput {
	return o
}

func (o AmiLaunchPermissionMapOutput) ToAmiLaunchPermissionMapOutputWithContext(ctx context.Context) AmiLaunchPermissionMapOutput {
	return o
}

func (o AmiLaunchPermissionMapOutput) MapIndex(k pulumi.StringInput) AmiLaunchPermissionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AmiLaunchPermission {
		return vs[0].(map[string]AmiLaunchPermission)[vs[1].(string)]
	}).(AmiLaunchPermissionOutput)
}

func init() {
	pulumi.RegisterOutputType(AmiLaunchPermissionOutput{})
	pulumi.RegisterOutputType(AmiLaunchPermissionPtrOutput{})
	pulumi.RegisterOutputType(AmiLaunchPermissionArrayOutput{})
	pulumi.RegisterOutputType(AmiLaunchPermissionMapOutput{})
}
