// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a VPC VPN Gateway.
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
// 		_, err := ec2.NewVpnGateway(ctx, "vpnGw", &ec2.VpnGatewayArgs{
// 			VpcId: pulumi.Any(aws_vpc.Main.Id),
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("main"),
// 			},
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
// VPN Gateways can be imported using the `vpn gateway id`, e.g.
//
// ```sh
//  $ pulumi import aws:ec2/vpnGateway:VpnGateway testvpngateway vgw-9a4cacf3
// ```
type VpnGateway struct {
	pulumi.CustomResourceState

	// The Autonomous System Number (ASN) for the Amazon side of the gateway. If you don't specify an ASN, the virtual private gateway is created with the default ASN.
	AmazonSideAsn pulumi.StringOutput `pulumi:"amazonSideAsn"`
	// Amazon Resource Name (ARN) of the VPN Gateway.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Availability Zone for the virtual private gateway.
	AvailabilityZone pulumi.StringPtrOutput `pulumi:"availabilityZone"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The VPC ID to create in.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewVpnGateway registers a new resource with the given unique name, arguments, and options.
func NewVpnGateway(ctx *pulumi.Context,
	name string, args *VpnGatewayArgs, opts ...pulumi.ResourceOption) (*VpnGateway, error) {
	if args == nil {
		args = &VpnGatewayArgs{}
	}

	var resource VpnGateway
	err := ctx.RegisterResource("aws:ec2/vpnGateway:VpnGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnGateway gets an existing VpnGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnGatewayState, opts ...pulumi.ResourceOption) (*VpnGateway, error) {
	var resource VpnGateway
	err := ctx.ReadResource("aws:ec2/vpnGateway:VpnGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnGateway resources.
type vpnGatewayState struct {
	// The Autonomous System Number (ASN) for the Amazon side of the gateway. If you don't specify an ASN, the virtual private gateway is created with the default ASN.
	AmazonSideAsn *string `pulumi:"amazonSideAsn"`
	// Amazon Resource Name (ARN) of the VPN Gateway.
	Arn *string `pulumi:"arn"`
	// The Availability Zone for the virtual private gateway.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The VPC ID to create in.
	VpcId *string `pulumi:"vpcId"`
}

type VpnGatewayState struct {
	// The Autonomous System Number (ASN) for the Amazon side of the gateway. If you don't specify an ASN, the virtual private gateway is created with the default ASN.
	AmazonSideAsn pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the VPN Gateway.
	Arn pulumi.StringPtrInput
	// The Availability Zone for the virtual private gateway.
	AvailabilityZone pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// The VPC ID to create in.
	VpcId pulumi.StringPtrInput
}

func (VpnGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnGatewayState)(nil)).Elem()
}

type vpnGatewayArgs struct {
	// The Autonomous System Number (ASN) for the Amazon side of the gateway. If you don't specify an ASN, the virtual private gateway is created with the default ASN.
	AmazonSideAsn *string `pulumi:"amazonSideAsn"`
	// The Availability Zone for the virtual private gateway.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The VPC ID to create in.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a VpnGateway resource.
type VpnGatewayArgs struct {
	// The Autonomous System Number (ASN) for the Amazon side of the gateway. If you don't specify an ASN, the virtual private gateway is created with the default ASN.
	AmazonSideAsn pulumi.StringPtrInput
	// The Availability Zone for the virtual private gateway.
	AvailabilityZone pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// The VPC ID to create in.
	VpcId pulumi.StringPtrInput
}

func (VpnGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnGatewayArgs)(nil)).Elem()
}

type VpnGatewayInput interface {
	pulumi.Input

	ToVpnGatewayOutput() VpnGatewayOutput
	ToVpnGatewayOutputWithContext(ctx context.Context) VpnGatewayOutput
}

func (*VpnGateway) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnGateway)(nil))
}

func (i *VpnGateway) ToVpnGatewayOutput() VpnGatewayOutput {
	return i.ToVpnGatewayOutputWithContext(context.Background())
}

func (i *VpnGateway) ToVpnGatewayOutputWithContext(ctx context.Context) VpnGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnGatewayOutput)
}

func (i *VpnGateway) ToVpnGatewayPtrOutput() VpnGatewayPtrOutput {
	return i.ToVpnGatewayPtrOutputWithContext(context.Background())
}

func (i *VpnGateway) ToVpnGatewayPtrOutputWithContext(ctx context.Context) VpnGatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnGatewayPtrOutput)
}

type VpnGatewayPtrInput interface {
	pulumi.Input

	ToVpnGatewayPtrOutput() VpnGatewayPtrOutput
	ToVpnGatewayPtrOutputWithContext(ctx context.Context) VpnGatewayPtrOutput
}

type vpnGatewayPtrType VpnGatewayArgs

func (*vpnGatewayPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnGateway)(nil))
}

func (i *vpnGatewayPtrType) ToVpnGatewayPtrOutput() VpnGatewayPtrOutput {
	return i.ToVpnGatewayPtrOutputWithContext(context.Background())
}

func (i *vpnGatewayPtrType) ToVpnGatewayPtrOutputWithContext(ctx context.Context) VpnGatewayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnGatewayPtrOutput)
}

// VpnGatewayArrayInput is an input type that accepts VpnGatewayArray and VpnGatewayArrayOutput values.
// You can construct a concrete instance of `VpnGatewayArrayInput` via:
//
//          VpnGatewayArray{ VpnGatewayArgs{...} }
type VpnGatewayArrayInput interface {
	pulumi.Input

	ToVpnGatewayArrayOutput() VpnGatewayArrayOutput
	ToVpnGatewayArrayOutputWithContext(context.Context) VpnGatewayArrayOutput
}

type VpnGatewayArray []VpnGatewayInput

func (VpnGatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*VpnGateway)(nil))
}

func (i VpnGatewayArray) ToVpnGatewayArrayOutput() VpnGatewayArrayOutput {
	return i.ToVpnGatewayArrayOutputWithContext(context.Background())
}

func (i VpnGatewayArray) ToVpnGatewayArrayOutputWithContext(ctx context.Context) VpnGatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnGatewayArrayOutput)
}

// VpnGatewayMapInput is an input type that accepts VpnGatewayMap and VpnGatewayMapOutput values.
// You can construct a concrete instance of `VpnGatewayMapInput` via:
//
//          VpnGatewayMap{ "key": VpnGatewayArgs{...} }
type VpnGatewayMapInput interface {
	pulumi.Input

	ToVpnGatewayMapOutput() VpnGatewayMapOutput
	ToVpnGatewayMapOutputWithContext(context.Context) VpnGatewayMapOutput
}

type VpnGatewayMap map[string]VpnGatewayInput

func (VpnGatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*VpnGateway)(nil))
}

func (i VpnGatewayMap) ToVpnGatewayMapOutput() VpnGatewayMapOutput {
	return i.ToVpnGatewayMapOutputWithContext(context.Background())
}

func (i VpnGatewayMap) ToVpnGatewayMapOutputWithContext(ctx context.Context) VpnGatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnGatewayMapOutput)
}

type VpnGatewayOutput struct {
	*pulumi.OutputState
}

func (VpnGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnGateway)(nil))
}

func (o VpnGatewayOutput) ToVpnGatewayOutput() VpnGatewayOutput {
	return o
}

func (o VpnGatewayOutput) ToVpnGatewayOutputWithContext(ctx context.Context) VpnGatewayOutput {
	return o
}

func (o VpnGatewayOutput) ToVpnGatewayPtrOutput() VpnGatewayPtrOutput {
	return o.ToVpnGatewayPtrOutputWithContext(context.Background())
}

func (o VpnGatewayOutput) ToVpnGatewayPtrOutputWithContext(ctx context.Context) VpnGatewayPtrOutput {
	return o.ApplyT(func(v VpnGateway) *VpnGateway {
		return &v
	}).(VpnGatewayPtrOutput)
}

type VpnGatewayPtrOutput struct {
	*pulumi.OutputState
}

func (VpnGatewayPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnGateway)(nil))
}

func (o VpnGatewayPtrOutput) ToVpnGatewayPtrOutput() VpnGatewayPtrOutput {
	return o
}

func (o VpnGatewayPtrOutput) ToVpnGatewayPtrOutputWithContext(ctx context.Context) VpnGatewayPtrOutput {
	return o
}

type VpnGatewayArrayOutput struct{ *pulumi.OutputState }

func (VpnGatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnGateway)(nil))
}

func (o VpnGatewayArrayOutput) ToVpnGatewayArrayOutput() VpnGatewayArrayOutput {
	return o
}

func (o VpnGatewayArrayOutput) ToVpnGatewayArrayOutputWithContext(ctx context.Context) VpnGatewayArrayOutput {
	return o
}

func (o VpnGatewayArrayOutput) Index(i pulumi.IntInput) VpnGatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnGateway {
		return vs[0].([]VpnGateway)[vs[1].(int)]
	}).(VpnGatewayOutput)
}

type VpnGatewayMapOutput struct{ *pulumi.OutputState }

func (VpnGatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VpnGateway)(nil))
}

func (o VpnGatewayMapOutput) ToVpnGatewayMapOutput() VpnGatewayMapOutput {
	return o
}

func (o VpnGatewayMapOutput) ToVpnGatewayMapOutputWithContext(ctx context.Context) VpnGatewayMapOutput {
	return o
}

func (o VpnGatewayMapOutput) MapIndex(k pulumi.StringInput) VpnGatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VpnGateway {
		return vs[0].(map[string]VpnGateway)[vs[1].(string)]
	}).(VpnGatewayOutput)
}

func init() {
	pulumi.RegisterOutputType(VpnGatewayOutput{})
	pulumi.RegisterOutputType(VpnGatewayPtrOutput{})
	pulumi.RegisterOutputType(VpnGatewayArrayOutput{})
	pulumi.RegisterOutputType(VpnGatewayMapOutput{})
}
