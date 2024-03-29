// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ec2.SecurityGroup` provides details about a specific Security Group.
//
// This resource can prove useful when a module accepts a Security Group id as
// an input variable and needs to, for example, determine the id of the
// VPC that the security group belongs to.
//
// ## Example Usage
//
// The following example shows how one might accept a Security Group id as a variable
// and use this data source to obtain the data necessary to create a subnet.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		securityGroupId := cfg.RequireObject("securityGroupId")
// 		opt0 := securityGroupId
// 		selected, err := ec2.LookupSecurityGroup(ctx, &ec2.LookupSecurityGroupArgs{
// 			Id: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewSubnet(ctx, "subnet", &ec2.SubnetArgs{
// 			VpcId:     pulumi.String(selected.VpcId),
// 			CidrBlock: pulumi.String("10.0.1.0/24"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupSecurityGroup(ctx *pulumi.Context, args *LookupSecurityGroupArgs, opts ...pulumi.InvokeOption) (*LookupSecurityGroupResult, error) {
	var rv LookupSecurityGroupResult
	err := ctx.Invoke("aws:ec2/getSecurityGroup:getSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecurityGroup.
type LookupSecurityGroupArgs struct {
	// Custom filter block as described below.
	Filters []GetSecurityGroupFilter `pulumi:"filters"`
	// The id of the specific security group to retrieve.
	Id *string `pulumi:"id"`
	// The name of the field to filter by, as defined by
	// [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSecurityGroups.html).
	Name *string `pulumi:"name"`
	// A map of tags, each pair of which must exactly match
	// a pair on the desired security group.
	Tags map[string]string `pulumi:"tags"`
	// The id of the VPC that the desired security group belongs to.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getSecurityGroup.
type LookupSecurityGroupResult struct {
	// The computed ARN of the security group.
	Arn string `pulumi:"arn"`
	// The description of the security group.
	Description string                   `pulumi:"description"`
	Filters     []GetSecurityGroupFilter `pulumi:"filters"`
	Id          string                   `pulumi:"id"`
	Name        string                   `pulumi:"name"`
	Tags        map[string]string        `pulumi:"tags"`
	VpcId       string                   `pulumi:"vpcId"`
}
