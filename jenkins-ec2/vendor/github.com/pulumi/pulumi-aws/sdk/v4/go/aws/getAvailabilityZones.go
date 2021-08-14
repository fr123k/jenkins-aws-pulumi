// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Availability Zones data source allows access to the list of AWS
// Availability Zones which can be accessed by an AWS account within the region
// configured in the provider.
//
// This is different from the `getAvailabilityZone` (singular) data source,
// which provides some details about a specific availability zone.
//
// > When [Local Zones](https://aws.amazon.com/about-aws/global-infrastructure/localzones/) are enabled in a region, by default the API and this data source include both Local Zones and Availability Zones. To return only Availability Zones, see the example section below.
//
// ## Example Usage
// ### By State
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "available"
// 		available, err := aws.GetAvailabilityZones(ctx, &GetAvailabilityZonesArgs{
// 			State: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewSubnet(ctx, "primary", &ec2.SubnetArgs{
// 			AvailabilityZone: pulumi.String(available.Names[0]),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewSubnet(ctx, "secondary", &ec2.SubnetArgs{
// 			AvailabilityZone: pulumi.String(available.Names[1]),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### By Filter
//
// All Local Zones (regardless of opt-in status):
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := true
// 		_, err := aws.GetAvailabilityZones(ctx, &GetAvailabilityZonesArgs{
// 			AllAvailabilityZones: &opt0,
// 			Filters: []GetAvailabilityZonesFilter{
// 				GetAvailabilityZonesFilter{
// 					Name: "opt-in-status",
// 					Values: []string{
// 						"not-opted-in",
// 						"opted-in",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Only Availability Zones (no Local Zones):
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := aws.GetAvailabilityZones(ctx, &GetAvailabilityZonesArgs{
// 			Filters: []GetAvailabilityZonesFilter{
// 				GetAvailabilityZonesFilter{
// 					Name: "opt-in-status",
// 					Values: []string{
// 						"opt-in-not-required",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetAvailabilityZones(ctx *pulumi.Context, args *GetAvailabilityZonesArgs, opts ...pulumi.InvokeOption) (*GetAvailabilityZonesResult, error) {
	var rv GetAvailabilityZonesResult
	err := ctx.Invoke("aws:index/getAvailabilityZones:getAvailabilityZones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAvailabilityZones.
type GetAvailabilityZonesArgs struct {
	// Set to `true` to include all Availability Zones and Local Zones regardless of your opt in status.
	AllAvailabilityZones *bool `pulumi:"allAvailabilityZones"`
	// List of Availability Zone names to exclude.
	ExcludeNames []string `pulumi:"excludeNames"`
	// List of Availability Zone IDs to exclude.
	ExcludeZoneIds []string `pulumi:"excludeZoneIds"`
	// Configuration block(s) for filtering. Detailed below.
	Filters []GetAvailabilityZonesFilter `pulumi:"filters"`
	// Allows to filter list of Availability Zones based on their
	// current state. Can be either `"available"`, `"information"`, `"impaired"` or
	// `"unavailable"`. By default the list includes a complete set of Availability Zones
	// to which the underlying AWS account has access, regardless of their state.
	State *string `pulumi:"state"`
}

// A collection of values returned by getAvailabilityZones.
type GetAvailabilityZonesResult struct {
	AllAvailabilityZones *bool                        `pulumi:"allAvailabilityZones"`
	ExcludeNames         []string                     `pulumi:"excludeNames"`
	ExcludeZoneIds       []string                     `pulumi:"excludeZoneIds"`
	Filters              []GetAvailabilityZonesFilter `pulumi:"filters"`
	GroupNames           []string                     `pulumi:"groupNames"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the Availability Zone names available to the account.
	Names []string `pulumi:"names"`
	State *string  `pulumi:"state"`
	// A list of the Availability Zone IDs available to the account.
	ZoneIds []string `pulumi:"zoneIds"`
}