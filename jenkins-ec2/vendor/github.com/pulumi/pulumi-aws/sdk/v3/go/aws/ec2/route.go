// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to create a routing table entry (a route) in a VPC routing table.
//
// > **NOTE on Route Tables and Routes:** This provider currently
// provides both a standalone Route resource and a Route Table resource with routes
// defined in-line. At this time you cannot use a Route Table with in-line routes
// in conjunction with any Route resources. Doing so will cause
// a conflict of rule settings and will overwrite rules.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2.NewRoute(ctx, "route", &ec2.RouteArgs{
// 			RouteTableId:           pulumi.String("rtb-4fbb3ac4"),
// 			DestinationCidrBlock:   pulumi.String("10.0.1.0/22"),
// 			VpcPeeringConnectionId: pulumi.String("pcx-45ff3dc1"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			aws_route_table.Testing,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Example IPv6 Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		vpc, err := ec2.NewVpc(ctx, "vpc", &ec2.VpcArgs{
// 			CidrBlock:                    pulumi.String("10.1.0.0/16"),
// 			AssignGeneratedIpv6CidrBlock: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		egress, err := ec2.NewEgressOnlyInternetGateway(ctx, "egress", &ec2.EgressOnlyInternetGatewayArgs{
// 			VpcId: vpc.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewRoute(ctx, "route", &ec2.RouteArgs{
// 			RouteTableId:             pulumi.String("rtb-4fbb3ac4"),
// 			DestinationIpv6CidrBlock: pulumi.String("::/0"),
// 			EgressOnlyGatewayId:      egress.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Route struct {
	pulumi.CustomResourceState

	// The destination CIDR block.
	DestinationCidrBlock pulumi.StringPtrOutput `pulumi:"destinationCidrBlock"`
	// The destination IPv6 CIDR block.
	DestinationIpv6CidrBlock pulumi.StringPtrOutput `pulumi:"destinationIpv6CidrBlock"`
	DestinationPrefixListId  pulumi.StringOutput    `pulumi:"destinationPrefixListId"`
	// Identifier of a VPC Egress Only Internet Gateway.
	EgressOnlyGatewayId pulumi.StringOutput `pulumi:"egressOnlyGatewayId"`
	// Identifier of a VPC internet gateway or a virtual private gateway.
	GatewayId pulumi.StringOutput `pulumi:"gatewayId"`
	// Identifier of an EC2 instance.
	InstanceId      pulumi.StringOutput `pulumi:"instanceId"`
	InstanceOwnerId pulumi.StringOutput `pulumi:"instanceOwnerId"`
	// Identifier of a Outpost local gateway.
	LocalGatewayId pulumi.StringOutput `pulumi:"localGatewayId"`
	// Identifier of a VPC NAT gateway.
	NatGatewayId pulumi.StringOutput `pulumi:"natGatewayId"`
	// Identifier of an EC2 network interface.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	Origin             pulumi.StringOutput `pulumi:"origin"`
	// The ID of the routing table.
	RouteTableId pulumi.StringOutput `pulumi:"routeTableId"`
	State        pulumi.StringOutput `pulumi:"state"`
	// Identifier of an EC2 Transit Gateway.
	TransitGatewayId pulumi.StringPtrOutput `pulumi:"transitGatewayId"`
	// Identifier of a VPC peering connection.
	VpcPeeringConnectionId pulumi.StringPtrOutput `pulumi:"vpcPeeringConnectionId"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil || args.RouteTableId == nil {
		return nil, errors.New("missing required argument 'RouteTableId'")
	}
	if args == nil {
		args = &RouteArgs{}
	}
	var resource Route
	err := ctx.RegisterResource("aws:ec2/route:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("aws:ec2/route:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
	// The destination CIDR block.
	DestinationCidrBlock *string `pulumi:"destinationCidrBlock"`
	// The destination IPv6 CIDR block.
	DestinationIpv6CidrBlock *string `pulumi:"destinationIpv6CidrBlock"`
	DestinationPrefixListId  *string `pulumi:"destinationPrefixListId"`
	// Identifier of a VPC Egress Only Internet Gateway.
	EgressOnlyGatewayId *string `pulumi:"egressOnlyGatewayId"`
	// Identifier of a VPC internet gateway or a virtual private gateway.
	GatewayId *string `pulumi:"gatewayId"`
	// Identifier of an EC2 instance.
	InstanceId      *string `pulumi:"instanceId"`
	InstanceOwnerId *string `pulumi:"instanceOwnerId"`
	// Identifier of a Outpost local gateway.
	LocalGatewayId *string `pulumi:"localGatewayId"`
	// Identifier of a VPC NAT gateway.
	NatGatewayId *string `pulumi:"natGatewayId"`
	// Identifier of an EC2 network interface.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	Origin             *string `pulumi:"origin"`
	// The ID of the routing table.
	RouteTableId *string `pulumi:"routeTableId"`
	State        *string `pulumi:"state"`
	// Identifier of an EC2 Transit Gateway.
	TransitGatewayId *string `pulumi:"transitGatewayId"`
	// Identifier of a VPC peering connection.
	VpcPeeringConnectionId *string `pulumi:"vpcPeeringConnectionId"`
}

type RouteState struct {
	// The destination CIDR block.
	DestinationCidrBlock pulumi.StringPtrInput
	// The destination IPv6 CIDR block.
	DestinationIpv6CidrBlock pulumi.StringPtrInput
	DestinationPrefixListId  pulumi.StringPtrInput
	// Identifier of a VPC Egress Only Internet Gateway.
	EgressOnlyGatewayId pulumi.StringPtrInput
	// Identifier of a VPC internet gateway or a virtual private gateway.
	GatewayId pulumi.StringPtrInput
	// Identifier of an EC2 instance.
	InstanceId      pulumi.StringPtrInput
	InstanceOwnerId pulumi.StringPtrInput
	// Identifier of a Outpost local gateway.
	LocalGatewayId pulumi.StringPtrInput
	// Identifier of a VPC NAT gateway.
	NatGatewayId pulumi.StringPtrInput
	// Identifier of an EC2 network interface.
	NetworkInterfaceId pulumi.StringPtrInput
	Origin             pulumi.StringPtrInput
	// The ID of the routing table.
	RouteTableId pulumi.StringPtrInput
	State        pulumi.StringPtrInput
	// Identifier of an EC2 Transit Gateway.
	TransitGatewayId pulumi.StringPtrInput
	// Identifier of a VPC peering connection.
	VpcPeeringConnectionId pulumi.StringPtrInput
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	// The destination CIDR block.
	DestinationCidrBlock *string `pulumi:"destinationCidrBlock"`
	// The destination IPv6 CIDR block.
	DestinationIpv6CidrBlock *string `pulumi:"destinationIpv6CidrBlock"`
	// Identifier of a VPC Egress Only Internet Gateway.
	EgressOnlyGatewayId *string `pulumi:"egressOnlyGatewayId"`
	// Identifier of a VPC internet gateway or a virtual private gateway.
	GatewayId *string `pulumi:"gatewayId"`
	// Identifier of an EC2 instance.
	InstanceId *string `pulumi:"instanceId"`
	// Identifier of a Outpost local gateway.
	LocalGatewayId *string `pulumi:"localGatewayId"`
	// Identifier of a VPC NAT gateway.
	NatGatewayId *string `pulumi:"natGatewayId"`
	// Identifier of an EC2 network interface.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The ID of the routing table.
	RouteTableId string `pulumi:"routeTableId"`
	// Identifier of an EC2 Transit Gateway.
	TransitGatewayId *string `pulumi:"transitGatewayId"`
	// Identifier of a VPC peering connection.
	VpcPeeringConnectionId *string `pulumi:"vpcPeeringConnectionId"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// The destination CIDR block.
	DestinationCidrBlock pulumi.StringPtrInput
	// The destination IPv6 CIDR block.
	DestinationIpv6CidrBlock pulumi.StringPtrInput
	// Identifier of a VPC Egress Only Internet Gateway.
	EgressOnlyGatewayId pulumi.StringPtrInput
	// Identifier of a VPC internet gateway or a virtual private gateway.
	GatewayId pulumi.StringPtrInput
	// Identifier of an EC2 instance.
	InstanceId pulumi.StringPtrInput
	// Identifier of a Outpost local gateway.
	LocalGatewayId pulumi.StringPtrInput
	// Identifier of a VPC NAT gateway.
	NatGatewayId pulumi.StringPtrInput
	// Identifier of an EC2 network interface.
	NetworkInterfaceId pulumi.StringPtrInput
	// The ID of the routing table.
	RouteTableId pulumi.StringInput
	// Identifier of an EC2 Transit Gateway.
	TransitGatewayId pulumi.StringPtrInput
	// Identifier of a VPC peering connection.
	VpcPeeringConnectionId pulumi.StringPtrInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}