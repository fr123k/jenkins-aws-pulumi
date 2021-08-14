// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// > **WARNING:** Multiple iam.GroupMembership resources with the same group name will produce inconsistent behavior!
//
// Provides a top level resource to manage IAM Group membership for IAM Users. For
// more information on managing IAM Groups or IAM Users, see IAM Groups or
// IAM Users
//
// > **Note:** `iam.GroupMembership` will conflict with itself if used more than once with the same group. To non-exclusively manage the users in a group, see the
// `iam.UserGroupMembership` resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		group, err := iam.NewGroup(ctx, "group", nil)
// 		if err != nil {
// 			return err
// 		}
// 		userOne, err := iam.NewUser(ctx, "userOne", nil)
// 		if err != nil {
// 			return err
// 		}
// 		userTwo, err := iam.NewUser(ctx, "userTwo", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewGroupMembership(ctx, "team", &iam.GroupMembershipArgs{
// 			Users: pulumi.StringArray{
// 				userOne.Name,
// 				userTwo.Name,
// 			},
// 			Group: group.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type GroupMembership struct {
	pulumi.CustomResourceState

	// The IAM Group name to attach the list of `users` to
	Group pulumi.StringOutput `pulumi:"group"`
	// The name to identify the Group Membership
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of IAM User names to associate with the Group
	Users pulumi.StringArrayOutput `pulumi:"users"`
}

// NewGroupMembership registers a new resource with the given unique name, arguments, and options.
func NewGroupMembership(ctx *pulumi.Context,
	name string, args *GroupMembershipArgs, opts ...pulumi.ResourceOption) (*GroupMembership, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	if args.Users == nil {
		return nil, errors.New("invalid value for required argument 'Users'")
	}
	var resource GroupMembership
	err := ctx.RegisterResource("aws:iam/groupMembership:GroupMembership", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupMembership gets an existing GroupMembership resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupMembership(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupMembershipState, opts ...pulumi.ResourceOption) (*GroupMembership, error) {
	var resource GroupMembership
	err := ctx.ReadResource("aws:iam/groupMembership:GroupMembership", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupMembership resources.
type groupMembershipState struct {
	// The IAM Group name to attach the list of `users` to
	Group *string `pulumi:"group"`
	// The name to identify the Group Membership
	Name *string `pulumi:"name"`
	// A list of IAM User names to associate with the Group
	Users []string `pulumi:"users"`
}

type GroupMembershipState struct {
	// The IAM Group name to attach the list of `users` to
	Group pulumi.StringPtrInput
	// The name to identify the Group Membership
	Name pulumi.StringPtrInput
	// A list of IAM User names to associate with the Group
	Users pulumi.StringArrayInput
}

func (GroupMembershipState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembershipState)(nil)).Elem()
}

type groupMembershipArgs struct {
	// The IAM Group name to attach the list of `users` to
	Group string `pulumi:"group"`
	// The name to identify the Group Membership
	Name *string `pulumi:"name"`
	// A list of IAM User names to associate with the Group
	Users []string `pulumi:"users"`
}

// The set of arguments for constructing a GroupMembership resource.
type GroupMembershipArgs struct {
	// The IAM Group name to attach the list of `users` to
	Group pulumi.StringInput
	// The name to identify the Group Membership
	Name pulumi.StringPtrInput
	// A list of IAM User names to associate with the Group
	Users pulumi.StringArrayInput
}

func (GroupMembershipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembershipArgs)(nil)).Elem()
}

type GroupMembershipInput interface {
	pulumi.Input

	ToGroupMembershipOutput() GroupMembershipOutput
	ToGroupMembershipOutputWithContext(ctx context.Context) GroupMembershipOutput
}

func (*GroupMembership) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupMembership)(nil))
}

func (i *GroupMembership) ToGroupMembershipOutput() GroupMembershipOutput {
	return i.ToGroupMembershipOutputWithContext(context.Background())
}

func (i *GroupMembership) ToGroupMembershipOutputWithContext(ctx context.Context) GroupMembershipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipOutput)
}

func (i *GroupMembership) ToGroupMembershipPtrOutput() GroupMembershipPtrOutput {
	return i.ToGroupMembershipPtrOutputWithContext(context.Background())
}

func (i *GroupMembership) ToGroupMembershipPtrOutputWithContext(ctx context.Context) GroupMembershipPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipPtrOutput)
}

type GroupMembershipPtrInput interface {
	pulumi.Input

	ToGroupMembershipPtrOutput() GroupMembershipPtrOutput
	ToGroupMembershipPtrOutputWithContext(ctx context.Context) GroupMembershipPtrOutput
}

type groupMembershipPtrType GroupMembershipArgs

func (*groupMembershipPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMembership)(nil))
}

func (i *groupMembershipPtrType) ToGroupMembershipPtrOutput() GroupMembershipPtrOutput {
	return i.ToGroupMembershipPtrOutputWithContext(context.Background())
}

func (i *groupMembershipPtrType) ToGroupMembershipPtrOutputWithContext(ctx context.Context) GroupMembershipPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipPtrOutput)
}

// GroupMembershipArrayInput is an input type that accepts GroupMembershipArray and GroupMembershipArrayOutput values.
// You can construct a concrete instance of `GroupMembershipArrayInput` via:
//
//          GroupMembershipArray{ GroupMembershipArgs{...} }
type GroupMembershipArrayInput interface {
	pulumi.Input

	ToGroupMembershipArrayOutput() GroupMembershipArrayOutput
	ToGroupMembershipArrayOutputWithContext(context.Context) GroupMembershipArrayOutput
}

type GroupMembershipArray []GroupMembershipInput

func (GroupMembershipArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*GroupMembership)(nil))
}

func (i GroupMembershipArray) ToGroupMembershipArrayOutput() GroupMembershipArrayOutput {
	return i.ToGroupMembershipArrayOutputWithContext(context.Background())
}

func (i GroupMembershipArray) ToGroupMembershipArrayOutputWithContext(ctx context.Context) GroupMembershipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipArrayOutput)
}

// GroupMembershipMapInput is an input type that accepts GroupMembershipMap and GroupMembershipMapOutput values.
// You can construct a concrete instance of `GroupMembershipMapInput` via:
//
//          GroupMembershipMap{ "key": GroupMembershipArgs{...} }
type GroupMembershipMapInput interface {
	pulumi.Input

	ToGroupMembershipMapOutput() GroupMembershipMapOutput
	ToGroupMembershipMapOutputWithContext(context.Context) GroupMembershipMapOutput
}

type GroupMembershipMap map[string]GroupMembershipInput

func (GroupMembershipMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*GroupMembership)(nil))
}

func (i GroupMembershipMap) ToGroupMembershipMapOutput() GroupMembershipMapOutput {
	return i.ToGroupMembershipMapOutputWithContext(context.Background())
}

func (i GroupMembershipMap) ToGroupMembershipMapOutputWithContext(ctx context.Context) GroupMembershipMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMembershipMapOutput)
}

type GroupMembershipOutput struct {
	*pulumi.OutputState
}

func (GroupMembershipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupMembership)(nil))
}

func (o GroupMembershipOutput) ToGroupMembershipOutput() GroupMembershipOutput {
	return o
}

func (o GroupMembershipOutput) ToGroupMembershipOutputWithContext(ctx context.Context) GroupMembershipOutput {
	return o
}

func (o GroupMembershipOutput) ToGroupMembershipPtrOutput() GroupMembershipPtrOutput {
	return o.ToGroupMembershipPtrOutputWithContext(context.Background())
}

func (o GroupMembershipOutput) ToGroupMembershipPtrOutputWithContext(ctx context.Context) GroupMembershipPtrOutput {
	return o.ApplyT(func(v GroupMembership) *GroupMembership {
		return &v
	}).(GroupMembershipPtrOutput)
}

type GroupMembershipPtrOutput struct {
	*pulumi.OutputState
}

func (GroupMembershipPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupMembership)(nil))
}

func (o GroupMembershipPtrOutput) ToGroupMembershipPtrOutput() GroupMembershipPtrOutput {
	return o
}

func (o GroupMembershipPtrOutput) ToGroupMembershipPtrOutputWithContext(ctx context.Context) GroupMembershipPtrOutput {
	return o
}

type GroupMembershipArrayOutput struct{ *pulumi.OutputState }

func (GroupMembershipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupMembership)(nil))
}

func (o GroupMembershipArrayOutput) ToGroupMembershipArrayOutput() GroupMembershipArrayOutput {
	return o
}

func (o GroupMembershipArrayOutput) ToGroupMembershipArrayOutputWithContext(ctx context.Context) GroupMembershipArrayOutput {
	return o
}

func (o GroupMembershipArrayOutput) Index(i pulumi.IntInput) GroupMembershipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GroupMembership {
		return vs[0].([]GroupMembership)[vs[1].(int)]
	}).(GroupMembershipOutput)
}

type GroupMembershipMapOutput struct{ *pulumi.OutputState }

func (GroupMembershipMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GroupMembership)(nil))
}

func (o GroupMembershipMapOutput) ToGroupMembershipMapOutput() GroupMembershipMapOutput {
	return o
}

func (o GroupMembershipMapOutput) ToGroupMembershipMapOutputWithContext(ctx context.Context) GroupMembershipMapOutput {
	return o
}

func (o GroupMembershipMapOutput) MapIndex(k pulumi.StringInput) GroupMembershipOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GroupMembership {
		return vs[0].(map[string]GroupMembership)[vs[1].(string)]
	}).(GroupMembershipOutput)
}

func init() {
	pulumi.RegisterOutputType(GroupMembershipOutput{})
	pulumi.RegisterOutputType(GroupMembershipPtrOutput{})
	pulumi.RegisterOutputType(GroupMembershipArrayOutput{})
	pulumi.RegisterOutputType(GroupMembershipMapOutput{})
}
