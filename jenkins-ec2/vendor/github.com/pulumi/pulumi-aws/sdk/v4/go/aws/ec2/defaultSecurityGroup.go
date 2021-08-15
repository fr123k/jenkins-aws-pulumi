// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage a default security group. This resource can manage the default security group of the default or a non-default VPC.
//
// > **NOTE:** This is an advanced resource with special caveats. Please read this document in its entirety before using this resource. The `ec2.DefaultSecurityGroup` resource behaves differently from normal resources. This provider does not _create_ this resource but instead attempts to "adopt" it into management.
//
// For EC2 Classic accounts, each region comes with a default security group. Additionally, each VPC created in AWS comes with a default security group that can be managed but not destroyed.
//
// When the provider first adopts the default security group, it **immediately removes all ingress and egress rules in the Security Group**. It then creates any rules specified in the configuration. This way only the rules specified in the configuration are created.
//
// This resource treats its inline rules as absolute; only the rules defined inline are created, and any additions/removals external to this resource will result in diff shown. For these reasons, this resource is incompatible with the `ec2.SecurityGroupRule` resource.
//
// For more information about default security groups, see the AWS documentation on [Default Security Groups][aws-default-security-groups]. To manage normal security groups, see the `ec2.SecurityGroup` resource.
//
// ## Example Usage
//
// The following config gives the default security group the same rules that AWS provides by default but under management by this provider. This means that any ingress or egress rules added or changed will be detected as drift.
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
// 		mainvpc, err := ec2.NewVpc(ctx, "mainvpc", &ec2.VpcArgs{
// 			CidrBlock: pulumi.String("10.1.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewDefaultSecurityGroup(ctx, "_default", &ec2.DefaultSecurityGroupArgs{
// 			VpcId: mainvpc.ID(),
// 			Ingress: ec2.DefaultSecurityGroupIngressArray{
// 				&ec2.DefaultSecurityGroupIngressArgs{
// 					Protocol: pulumi.String("-1"),
// 					Self:     pulumi.Bool(true),
// 					FromPort: pulumi.Int(0),
// 					ToPort:   pulumi.Int(0),
// 				},
// 			},
// 			Egress: ec2.DefaultSecurityGroupEgressArray{
// 				&ec2.DefaultSecurityGroupEgressArgs{
// 					FromPort: pulumi.Int(0),
// 					ToPort:   pulumi.Int(0),
// 					Protocol: pulumi.String("-1"),
// 					CidrBlocks: pulumi.StringArray{
// 						pulumi.String("0.0.0.0/0"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Example Config To Deny All Egress Traffic, Allowing Ingress
//
// The following denies all Egress traffic by omitting any `egress` rules, while including the default `ingress` rule to allow all traffic.
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
// 		mainvpc, err := ec2.NewVpc(ctx, "mainvpc", &ec2.VpcArgs{
// 			CidrBlock: pulumi.String("10.1.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewDefaultSecurityGroup(ctx, "_default", &ec2.DefaultSecurityGroupArgs{
// 			VpcId: mainvpc.ID(),
// 			Ingress: ec2.DefaultSecurityGroupIngressArray{
// 				&ec2.DefaultSecurityGroupIngressArgs{
// 					Protocol: pulumi.String("-1"),
// 					Self:     pulumi.Bool(true),
// 					FromPort: pulumi.Int(0),
// 					ToPort:   pulumi.Int(0),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Removing `ec2.DefaultSecurityGroup` From Your Configuration
//
// Removing this resource from your configuration will remove it from your statefile and management, but will not destroy the Security Group. All ingress or egress rules will be left as they are at the time of removal. You can resume managing them via the AWS Console.
//
// ## Import
//
// Security Groups can be imported using the `security group id`, e.g.
//
// ```sh
//  $ pulumi import aws:ec2/defaultSecurityGroup:DefaultSecurityGroup default_sg sg-903004f8
// ```
type DefaultSecurityGroup struct {
	pulumi.CustomResourceState

	// ARN of the security group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of this rule.
	Description pulumi.StringOutput `pulumi:"description"`
	// Configuration block. Detailed below.
	Egress DefaultSecurityGroupEgressArrayOutput `pulumi:"egress"`
	// Configuration block. Detailed below.
	Ingress DefaultSecurityGroupIngressArrayOutput `pulumi:"ingress"`
	// Name of the security group.
	Name pulumi.StringOutput `pulumi:"name"`
	// Owner ID.
	OwnerId             pulumi.StringOutput  `pulumi:"ownerId"`
	RevokeRulesOnDelete pulumi.BoolPtrOutput `pulumi:"revokeRulesOnDelete"`
	// Map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// VPC ID. **Note that changing the `vpcId` will _not_ restore any default security group rules that were modified, added, or removed.** It will be left in its current state.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewDefaultSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewDefaultSecurityGroup(ctx *pulumi.Context,
	name string, args *DefaultSecurityGroupArgs, opts ...pulumi.ResourceOption) (*DefaultSecurityGroup, error) {
	if args == nil {
		args = &DefaultSecurityGroupArgs{}
	}

	var resource DefaultSecurityGroup
	err := ctx.RegisterResource("aws:ec2/defaultSecurityGroup:DefaultSecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultSecurityGroup gets an existing DefaultSecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultSecurityGroupState, opts ...pulumi.ResourceOption) (*DefaultSecurityGroup, error) {
	var resource DefaultSecurityGroup
	err := ctx.ReadResource("aws:ec2/defaultSecurityGroup:DefaultSecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultSecurityGroup resources.
type defaultSecurityGroupState struct {
	// ARN of the security group.
	Arn *string `pulumi:"arn"`
	// Description of this rule.
	Description *string `pulumi:"description"`
	// Configuration block. Detailed below.
	Egress []DefaultSecurityGroupEgress `pulumi:"egress"`
	// Configuration block. Detailed below.
	Ingress []DefaultSecurityGroupIngress `pulumi:"ingress"`
	// Name of the security group.
	Name *string `pulumi:"name"`
	// Owner ID.
	OwnerId             *string `pulumi:"ownerId"`
	RevokeRulesOnDelete *bool   `pulumi:"revokeRulesOnDelete"`
	// Map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// VPC ID. **Note that changing the `vpcId` will _not_ restore any default security group rules that were modified, added, or removed.** It will be left in its current state.
	VpcId *string `pulumi:"vpcId"`
}

type DefaultSecurityGroupState struct {
	// ARN of the security group.
	Arn pulumi.StringPtrInput
	// Description of this rule.
	Description pulumi.StringPtrInput
	// Configuration block. Detailed below.
	Egress DefaultSecurityGroupEgressArrayInput
	// Configuration block. Detailed below.
	Ingress DefaultSecurityGroupIngressArrayInput
	// Name of the security group.
	Name pulumi.StringPtrInput
	// Owner ID.
	OwnerId             pulumi.StringPtrInput
	RevokeRulesOnDelete pulumi.BoolPtrInput
	// Map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// VPC ID. **Note that changing the `vpcId` will _not_ restore any default security group rules that were modified, added, or removed.** It will be left in its current state.
	VpcId pulumi.StringPtrInput
}

func (DefaultSecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultSecurityGroupState)(nil)).Elem()
}

type defaultSecurityGroupArgs struct {
	// Configuration block. Detailed below.
	Egress []DefaultSecurityGroupEgress `pulumi:"egress"`
	// Configuration block. Detailed below.
	Ingress             []DefaultSecurityGroupIngress `pulumi:"ingress"`
	RevokeRulesOnDelete *bool                         `pulumi:"revokeRulesOnDelete"`
	// Map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// VPC ID. **Note that changing the `vpcId` will _not_ restore any default security group rules that were modified, added, or removed.** It will be left in its current state.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a DefaultSecurityGroup resource.
type DefaultSecurityGroupArgs struct {
	// Configuration block. Detailed below.
	Egress DefaultSecurityGroupEgressArrayInput
	// Configuration block. Detailed below.
	Ingress             DefaultSecurityGroupIngressArrayInput
	RevokeRulesOnDelete pulumi.BoolPtrInput
	// Map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// VPC ID. **Note that changing the `vpcId` will _not_ restore any default security group rules that were modified, added, or removed.** It will be left in its current state.
	VpcId pulumi.StringPtrInput
}

func (DefaultSecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultSecurityGroupArgs)(nil)).Elem()
}

type DefaultSecurityGroupInput interface {
	pulumi.Input

	ToDefaultSecurityGroupOutput() DefaultSecurityGroupOutput
	ToDefaultSecurityGroupOutputWithContext(ctx context.Context) DefaultSecurityGroupOutput
}

func (*DefaultSecurityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultSecurityGroup)(nil))
}

func (i *DefaultSecurityGroup) ToDefaultSecurityGroupOutput() DefaultSecurityGroupOutput {
	return i.ToDefaultSecurityGroupOutputWithContext(context.Background())
}

func (i *DefaultSecurityGroup) ToDefaultSecurityGroupOutputWithContext(ctx context.Context) DefaultSecurityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultSecurityGroupOutput)
}

func (i *DefaultSecurityGroup) ToDefaultSecurityGroupPtrOutput() DefaultSecurityGroupPtrOutput {
	return i.ToDefaultSecurityGroupPtrOutputWithContext(context.Background())
}

func (i *DefaultSecurityGroup) ToDefaultSecurityGroupPtrOutputWithContext(ctx context.Context) DefaultSecurityGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultSecurityGroupPtrOutput)
}

type DefaultSecurityGroupPtrInput interface {
	pulumi.Input

	ToDefaultSecurityGroupPtrOutput() DefaultSecurityGroupPtrOutput
	ToDefaultSecurityGroupPtrOutputWithContext(ctx context.Context) DefaultSecurityGroupPtrOutput
}

type defaultSecurityGroupPtrType DefaultSecurityGroupArgs

func (*defaultSecurityGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultSecurityGroup)(nil))
}

func (i *defaultSecurityGroupPtrType) ToDefaultSecurityGroupPtrOutput() DefaultSecurityGroupPtrOutput {
	return i.ToDefaultSecurityGroupPtrOutputWithContext(context.Background())
}

func (i *defaultSecurityGroupPtrType) ToDefaultSecurityGroupPtrOutputWithContext(ctx context.Context) DefaultSecurityGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultSecurityGroupPtrOutput)
}

// DefaultSecurityGroupArrayInput is an input type that accepts DefaultSecurityGroupArray and DefaultSecurityGroupArrayOutput values.
// You can construct a concrete instance of `DefaultSecurityGroupArrayInput` via:
//
//          DefaultSecurityGroupArray{ DefaultSecurityGroupArgs{...} }
type DefaultSecurityGroupArrayInput interface {
	pulumi.Input

	ToDefaultSecurityGroupArrayOutput() DefaultSecurityGroupArrayOutput
	ToDefaultSecurityGroupArrayOutputWithContext(context.Context) DefaultSecurityGroupArrayOutput
}

type DefaultSecurityGroupArray []DefaultSecurityGroupInput

func (DefaultSecurityGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultSecurityGroup)(nil)).Elem()
}

func (i DefaultSecurityGroupArray) ToDefaultSecurityGroupArrayOutput() DefaultSecurityGroupArrayOutput {
	return i.ToDefaultSecurityGroupArrayOutputWithContext(context.Background())
}

func (i DefaultSecurityGroupArray) ToDefaultSecurityGroupArrayOutputWithContext(ctx context.Context) DefaultSecurityGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultSecurityGroupArrayOutput)
}

// DefaultSecurityGroupMapInput is an input type that accepts DefaultSecurityGroupMap and DefaultSecurityGroupMapOutput values.
// You can construct a concrete instance of `DefaultSecurityGroupMapInput` via:
//
//          DefaultSecurityGroupMap{ "key": DefaultSecurityGroupArgs{...} }
type DefaultSecurityGroupMapInput interface {
	pulumi.Input

	ToDefaultSecurityGroupMapOutput() DefaultSecurityGroupMapOutput
	ToDefaultSecurityGroupMapOutputWithContext(context.Context) DefaultSecurityGroupMapOutput
}

type DefaultSecurityGroupMap map[string]DefaultSecurityGroupInput

func (DefaultSecurityGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultSecurityGroup)(nil)).Elem()
}

func (i DefaultSecurityGroupMap) ToDefaultSecurityGroupMapOutput() DefaultSecurityGroupMapOutput {
	return i.ToDefaultSecurityGroupMapOutputWithContext(context.Background())
}

func (i DefaultSecurityGroupMap) ToDefaultSecurityGroupMapOutputWithContext(ctx context.Context) DefaultSecurityGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultSecurityGroupMapOutput)
}

type DefaultSecurityGroupOutput struct{ *pulumi.OutputState }

func (DefaultSecurityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultSecurityGroup)(nil))
}

func (o DefaultSecurityGroupOutput) ToDefaultSecurityGroupOutput() DefaultSecurityGroupOutput {
	return o
}

func (o DefaultSecurityGroupOutput) ToDefaultSecurityGroupOutputWithContext(ctx context.Context) DefaultSecurityGroupOutput {
	return o
}

func (o DefaultSecurityGroupOutput) ToDefaultSecurityGroupPtrOutput() DefaultSecurityGroupPtrOutput {
	return o.ToDefaultSecurityGroupPtrOutputWithContext(context.Background())
}

func (o DefaultSecurityGroupOutput) ToDefaultSecurityGroupPtrOutputWithContext(ctx context.Context) DefaultSecurityGroupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefaultSecurityGroup) *DefaultSecurityGroup {
		return &v
	}).(DefaultSecurityGroupPtrOutput)
}

type DefaultSecurityGroupPtrOutput struct{ *pulumi.OutputState }

func (DefaultSecurityGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultSecurityGroup)(nil))
}

func (o DefaultSecurityGroupPtrOutput) ToDefaultSecurityGroupPtrOutput() DefaultSecurityGroupPtrOutput {
	return o
}

func (o DefaultSecurityGroupPtrOutput) ToDefaultSecurityGroupPtrOutputWithContext(ctx context.Context) DefaultSecurityGroupPtrOutput {
	return o
}

func (o DefaultSecurityGroupPtrOutput) Elem() DefaultSecurityGroupOutput {
	return o.ApplyT(func(v *DefaultSecurityGroup) DefaultSecurityGroup {
		if v != nil {
			return *v
		}
		var ret DefaultSecurityGroup
		return ret
	}).(DefaultSecurityGroupOutput)
}

type DefaultSecurityGroupArrayOutput struct{ *pulumi.OutputState }

func (DefaultSecurityGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DefaultSecurityGroup)(nil))
}

func (o DefaultSecurityGroupArrayOutput) ToDefaultSecurityGroupArrayOutput() DefaultSecurityGroupArrayOutput {
	return o
}

func (o DefaultSecurityGroupArrayOutput) ToDefaultSecurityGroupArrayOutputWithContext(ctx context.Context) DefaultSecurityGroupArrayOutput {
	return o
}

func (o DefaultSecurityGroupArrayOutput) Index(i pulumi.IntInput) DefaultSecurityGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DefaultSecurityGroup {
		return vs[0].([]DefaultSecurityGroup)[vs[1].(int)]
	}).(DefaultSecurityGroupOutput)
}

type DefaultSecurityGroupMapOutput struct{ *pulumi.OutputState }

func (DefaultSecurityGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DefaultSecurityGroup)(nil))
}

func (o DefaultSecurityGroupMapOutput) ToDefaultSecurityGroupMapOutput() DefaultSecurityGroupMapOutput {
	return o
}

func (o DefaultSecurityGroupMapOutput) ToDefaultSecurityGroupMapOutputWithContext(ctx context.Context) DefaultSecurityGroupMapOutput {
	return o
}

func (o DefaultSecurityGroupMapOutput) MapIndex(k pulumi.StringInput) DefaultSecurityGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DefaultSecurityGroup {
		return vs[0].(map[string]DefaultSecurityGroup)[vs[1].(string)]
	}).(DefaultSecurityGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(DefaultSecurityGroupOutput{})
	pulumi.RegisterOutputType(DefaultSecurityGroupPtrOutput{})
	pulumi.RegisterOutputType(DefaultSecurityGroupArrayOutput{})
	pulumi.RegisterOutputType(DefaultSecurityGroupMapOutput{})
}
