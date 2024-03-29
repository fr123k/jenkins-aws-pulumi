// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about the host when allocating an EC2 Dedicated Host.
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
// 		test, err := ec2.NewDedicatedHost(ctx, "test", &ec2.DedicatedHostArgs{
// 			AutoPlacement:    pulumi.String("on"),
// 			AvailabilityZone: pulumi.String("us-west-1a"),
// 			HostRecovery:     pulumi.String("on"),
// 			InstanceType:     pulumi.String("c5.18xlarge"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupDedicatedHost(ctx *pulumi.Context, args *LookupDedicatedHostArgs, opts ...pulumi.InvokeOption) (*LookupDedicatedHostResult, error) {
	var rv LookupDedicatedHostResult
	err := ctx.Invoke("aws:ec2/getDedicatedHost:getDedicatedHost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDedicatedHost.
type LookupDedicatedHostArgs struct {
	// The host ID.
	HostId string            `pulumi:"hostId"`
	Tags   map[string]string `pulumi:"tags"`
}

// A collection of values returned by getDedicatedHost.
type LookupDedicatedHostResult struct {
	AutoPlacement    string `pulumi:"autoPlacement"`
	AvailabilityZone string `pulumi:"availabilityZone"`
	// The number of cores on the Dedicated Host.
	Cores int `pulumi:"cores"`
	// The host ID.
	HostId       string `pulumi:"hostId"`
	HostRecovery string `pulumi:"hostRecovery"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The instance family supported by the Dedicated Host. For example, m5.
	// * `instanceType` -The instance type supported by the Dedicated Host. For example, m5.large. If the host supports multiple instance types, no instanceType is returned.
	InstanceFamily string `pulumi:"instanceFamily"`
	InstanceState  string `pulumi:"instanceState"`
	InstanceType   string `pulumi:"instanceType"`
	// The instance family supported by the Dedicated Host. For example, m5.
	Sockets int               `pulumi:"sockets"`
	Tags    map[string]string `pulumi:"tags"`
	// The total number of vCPUs on the Dedicated Host.
	TotalVcpus int `pulumi:"totalVcpus"`
}
