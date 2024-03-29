// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can be used to fetch information about a specific
// IAM user. By using this data source, you can reference IAM user
// properties without having to hard code ARNs or unique IDs as input.
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
// 		_, err := iam.LookupUser(ctx, &iam.LookupUserArgs{
// 			UserName: "an_example_user_name",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	var rv LookupUserResult
	err := ctx.Invoke("aws:iam/getUser:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUser.
type LookupUserArgs struct {
	// Map of key-value pairs associated with the user.
	Tags map[string]string `pulumi:"tags"`
	// The friendly IAM user name to match.
	UserName string `pulumi:"userName"`
}

// A collection of values returned by getUser.
type LookupUserResult struct {
	// The Amazon Resource Name (ARN) assigned by AWS for this user.
	Arn string `pulumi:"arn"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Path in which this user was created.
	Path string `pulumi:"path"`
	// The ARN of the policy that is used to set the permissions boundary for the user.
	PermissionsBoundary string `pulumi:"permissionsBoundary"`
	// Map of key-value pairs associated with the user.
	Tags map[string]string `pulumi:"tags"`
	// The unique ID assigned by AWS for this user.
	UserId string `pulumi:"userId"`
	// The name associated to this User
	UserName string `pulumi:"userName"`
}
