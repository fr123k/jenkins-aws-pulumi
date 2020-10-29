// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `getAvailabilityZone` provides details about a specific availability zone (AZ)
// in the current region.
//
// This can be used both to validate an availability zone given in a variable
// and to split the AZ name into its component parts of an AWS region and an
// AZ identifier letter. The latter may be useful e.g. for implementing a
// consistent subnet numbering scheme across several regions by mapping both
// the region and the subnet letter to network numbers.
//
// This is different from the `getAvailabilityZones` (plural) data source,
// which provides a list of the available zones.
func GetAvailabilityZone(ctx *pulumi.Context, args *GetAvailabilityZoneArgs, opts ...pulumi.InvokeOption) (*GetAvailabilityZoneResult, error) {
	var rv GetAvailabilityZoneResult
	err := ctx.Invoke("aws:index/getAvailabilityZone:getAvailabilityZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAvailabilityZone.
type GetAvailabilityZoneArgs struct {
	// Set to `true` to include all Availability Zones and Local Zones regardless of your opt in status.
	AllAvailabilityZones *bool `pulumi:"allAvailabilityZones"`
	// Configuration block(s) for filtering. Detailed below.
	Filters []GetAvailabilityZoneFilter `pulumi:"filters"`
	// The name of the filter field. Valid values can be found in the [EC2 DescribeAvailabilityZones API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAvailabilityZones.html).
	Name *string `pulumi:"name"`
	// A specific availability zone state to require. May be any of `"available"`, `"information"` or `"impaired"`.
	State *string `pulumi:"state"`
	// The zone ID of the availability zone to select.
	ZoneId *string `pulumi:"zoneId"`
}

// A collection of values returned by getAvailabilityZone.
type GetAvailabilityZoneResult struct {
	AllAvailabilityZones *bool                       `pulumi:"allAvailabilityZones"`
	Filters              []GetAvailabilityZoneFilter `pulumi:"filters"`
	// For Availability Zones, this is the same value as the Region name. For Local Zones, the name of the associated group, for example `us-west-2-lax-1`.
	GroupName string `pulumi:"groupName"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The part of the AZ name that appears after the region name, uniquely identifying the AZ within its region.
	NameSuffix string `pulumi:"nameSuffix"`
	// The name of the location from which the address is advertised.
	NetworkBorderGroup string `pulumi:"networkBorderGroup"`
	// For Availability Zones, this always has the value of `opt-in-not-required`. For Local Zones, this is the opt in status. The possible values are `opted-in` and `not-opted-in`.
	OptInStatus string `pulumi:"optInStatus"`
	// The region where the selected availability zone resides. This is always the region selected on the provider, since this data source searches only within that region.
	Region string `pulumi:"region"`
	State  string `pulumi:"state"`
	ZoneId string `pulumi:"zoneId"`
}
