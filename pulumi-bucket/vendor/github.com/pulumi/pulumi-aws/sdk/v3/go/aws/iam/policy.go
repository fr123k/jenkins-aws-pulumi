// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an IAM policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iam.NewPolicy(ctx, "policy", &iam.PolicyArgs{
// 			Description: pulumi.String("My test policy"),
// 			Path:        pulumi.String("/"),
// 			Policy:      pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": [\n", "        \"ec2:Describe*\"\n", "      ],\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": \"*\"\n", "    }\n", "  ]\n", "}\n", "\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Policy struct {
	pulumi.CustomResourceState

	// The ARN assigned by AWS to this policy.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of the IAM policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the policy. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// Path in which to create the policy.
	// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// The policy document. This is a JSON formatted string.
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil || args.Policy == nil {
		return nil, errors.New("missing required argument 'Policy'")
	}
	if args == nil {
		args = &PolicyArgs{}
	}
	var resource Policy
	err := ctx.RegisterResource("aws:iam/policy:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("aws:iam/policy:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// The ARN assigned by AWS to this policy.
	Arn *string `pulumi:"arn"`
	// Description of the IAM policy.
	Description *string `pulumi:"description"`
	// The name of the policy. If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Path in which to create the policy.
	// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path *string `pulumi:"path"`
	// The policy document. This is a JSON formatted string.
	Policy *string `pulumi:"policy"`
}

type PolicyState struct {
	// The ARN assigned by AWS to this policy.
	Arn pulumi.StringPtrInput
	// Description of the IAM policy.
	Description pulumi.StringPtrInput
	// The name of the policy. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Path in which to create the policy.
	// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path pulumi.StringPtrInput
	// The policy document. This is a JSON formatted string.
	Policy pulumi.StringPtrInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// Description of the IAM policy.
	Description *string `pulumi:"description"`
	// The name of the policy. If omitted, this provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Path in which to create the policy.
	// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path *string `pulumi:"path"`
	// The policy document. This is a JSON formatted string.
	Policy interface{} `pulumi:"policy"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// Description of the IAM policy.
	Description pulumi.StringPtrInput
	// The name of the policy. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Path in which to create the policy.
	// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path pulumi.StringPtrInput
	// The policy document. This is a JSON formatted string.
	Policy pulumi.Input
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}