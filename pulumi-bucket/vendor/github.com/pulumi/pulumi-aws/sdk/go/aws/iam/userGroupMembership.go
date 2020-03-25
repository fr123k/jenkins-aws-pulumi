// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package iam

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource for adding an [IAM User][2] to [IAM Groups][1]. This
// resource can be used multiple times with the same user for non-overlapping
// groups.
// 
// To exclusively manage the users in a group, see the
// [`iam.GroupMembership` resource][3].
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iam_user_group_membership.html.markdown.
type UserGroupMembership struct {
	pulumi.CustomResourceState

	// A list of [IAM Groups][1] to add the user to
	Groups pulumi.StringArrayOutput `pulumi:"groups"`
	// The name of the [IAM User][2] to add to groups
	User pulumi.StringOutput `pulumi:"user"`
}

// NewUserGroupMembership registers a new resource with the given unique name, arguments, and options.
func NewUserGroupMembership(ctx *pulumi.Context,
	name string, args *UserGroupMembershipArgs, opts ...pulumi.ResourceOption) (*UserGroupMembership, error) {
	if args == nil || args.Groups == nil {
		return nil, errors.New("missing required argument 'Groups'")
	}
	if args == nil || args.User == nil {
		return nil, errors.New("missing required argument 'User'")
	}
	if args == nil {
		args = &UserGroupMembershipArgs{}
	}
	var resource UserGroupMembership
	err := ctx.RegisterResource("aws:iam/userGroupMembership:UserGroupMembership", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserGroupMembership gets an existing UserGroupMembership resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserGroupMembership(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserGroupMembershipState, opts ...pulumi.ResourceOption) (*UserGroupMembership, error) {
	var resource UserGroupMembership
	err := ctx.ReadResource("aws:iam/userGroupMembership:UserGroupMembership", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserGroupMembership resources.
type userGroupMembershipState struct {
	// A list of [IAM Groups][1] to add the user to
	Groups []string `pulumi:"groups"`
	// The name of the [IAM User][2] to add to groups
	User *string `pulumi:"user"`
}

type UserGroupMembershipState struct {
	// A list of [IAM Groups][1] to add the user to
	Groups pulumi.StringArrayInput
	// The name of the [IAM User][2] to add to groups
	User pulumi.StringPtrInput
}

func (UserGroupMembershipState) ElementType() reflect.Type {
	return reflect.TypeOf((*userGroupMembershipState)(nil)).Elem()
}

type userGroupMembershipArgs struct {
	// A list of [IAM Groups][1] to add the user to
	Groups []string `pulumi:"groups"`
	// The name of the [IAM User][2] to add to groups
	User string `pulumi:"user"`
}

// The set of arguments for constructing a UserGroupMembership resource.
type UserGroupMembershipArgs struct {
	// A list of [IAM Groups][1] to add the user to
	Groups pulumi.StringArrayInput
	// The name of the [IAM User][2] to add to groups
	User pulumi.StringInput
}

func (UserGroupMembershipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userGroupMembershipArgs)(nil)).Elem()
}

