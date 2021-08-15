// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an [EC2 key pair](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html) resource. A key pair is used to control login access to EC2 instances.
//
// Currently this resource requires an existing user-supplied key pair. This key pair's public key will be registered with AWS to allow logging-in to EC2 instances.
//
// When importing an existing key pair the public key material may be in any format supported by AWS. Supported formats (per the [AWS documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html#how-to-generate-your-own-key-and-import-it-to-aws)) are:
//
// * OpenSSH public key format (the format in ~/.ssh/authorized_keys)
// * Base64 encoded DER format
// * SSH public key file format as specified in RFC4716
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
// 		_, err := ec2.NewKeyPair(ctx, "deployer", &ec2.KeyPairArgs{
// 			PublicKey: pulumi.String("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQD3F6tyPEFEzV0LX3X8BsXdMsQz1x2cEikKDEY0aIj41qgxMCP/iteneqXSIFZBp5vizPvaoIR3Um9xK7PGoW8giupGn+EPuxIA4cDM4vzOqOkiMPhz5XK0whEjkVzTo4+S0puvDZuwIsdiW9mxhJc7tgBNL0cYlWSYVkz4G/fslNfRPW5mYAM49f4fhtxPb5ok4Q2Lg9dPKVHO/Bgeu5woMc7RY0p1ej6D4CKFE6lymSDJpW0YHX/wqE9+cfEauh7xZcG0q9t2ta6F6fmX0agvpFyZo8aFbXeUBr7osSCJNgvavWbM/06niWrOvYX2xwWdhXmXSrbX8ZbabVohBK41 email@example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Key Pairs can be imported using the `key_name`, e.g.
//
// ```sh
//  $ pulumi import aws:ec2/keyPair:KeyPair deployer deployer-key
// ```
type KeyPair struct {
	pulumi.CustomResourceState

	// The key pair ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The MD5 public key fingerprint as specified in section 4 of RFC 4716.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The name for the key pair.
	KeyName pulumi.StringOutput `pulumi:"keyName"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `keyName`.
	KeyNamePrefix pulumi.StringPtrOutput `pulumi:"keyNamePrefix"`
	// The key pair ID.
	KeyPairId pulumi.StringOutput `pulumi:"keyPairId"`
	// The public key material.
	PublicKey pulumi.StringOutput    `pulumi:"publicKey"`
	Tags      pulumi.StringMapOutput `pulumi:"tags"`
	TagsAll   pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewKeyPair registers a new resource with the given unique name, arguments, and options.
func NewKeyPair(ctx *pulumi.Context,
	name string, args *KeyPairArgs, opts ...pulumi.ResourceOption) (*KeyPair, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PublicKey == nil {
		return nil, errors.New("invalid value for required argument 'PublicKey'")
	}
	var resource KeyPair
	err := ctx.RegisterResource("aws:ec2/keyPair:KeyPair", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyPair gets an existing KeyPair resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyPair(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyPairState, opts ...pulumi.ResourceOption) (*KeyPair, error) {
	var resource KeyPair
	err := ctx.ReadResource("aws:ec2/keyPair:KeyPair", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyPair resources.
type keyPairState struct {
	// The key pair ARN.
	Arn *string `pulumi:"arn"`
	// The MD5 public key fingerprint as specified in section 4 of RFC 4716.
	Fingerprint *string `pulumi:"fingerprint"`
	// The name for the key pair.
	KeyName *string `pulumi:"keyName"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `keyName`.
	KeyNamePrefix *string `pulumi:"keyNamePrefix"`
	// The key pair ID.
	KeyPairId *string `pulumi:"keyPairId"`
	// The public key material.
	PublicKey *string           `pulumi:"publicKey"`
	Tags      map[string]string `pulumi:"tags"`
	TagsAll   map[string]string `pulumi:"tagsAll"`
}

type KeyPairState struct {
	// The key pair ARN.
	Arn pulumi.StringPtrInput
	// The MD5 public key fingerprint as specified in section 4 of RFC 4716.
	Fingerprint pulumi.StringPtrInput
	// The name for the key pair.
	KeyName pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `keyName`.
	KeyNamePrefix pulumi.StringPtrInput
	// The key pair ID.
	KeyPairId pulumi.StringPtrInput
	// The public key material.
	PublicKey pulumi.StringPtrInput
	Tags      pulumi.StringMapInput
	TagsAll   pulumi.StringMapInput
}

func (KeyPairState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyPairState)(nil)).Elem()
}

type keyPairArgs struct {
	// The name for the key pair.
	KeyName *string `pulumi:"keyName"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `keyName`.
	KeyNamePrefix *string `pulumi:"keyNamePrefix"`
	// The public key material.
	PublicKey string            `pulumi:"publicKey"`
	Tags      map[string]string `pulumi:"tags"`
	TagsAll   map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a KeyPair resource.
type KeyPairArgs struct {
	// The name for the key pair.
	KeyName pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `keyName`.
	KeyNamePrefix pulumi.StringPtrInput
	// The public key material.
	PublicKey pulumi.StringInput
	Tags      pulumi.StringMapInput
	TagsAll   pulumi.StringMapInput
}

func (KeyPairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyPairArgs)(nil)).Elem()
}

type KeyPairInput interface {
	pulumi.Input

	ToKeyPairOutput() KeyPairOutput
	ToKeyPairOutputWithContext(ctx context.Context) KeyPairOutput
}

func (*KeyPair) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyPair)(nil))
}

func (i *KeyPair) ToKeyPairOutput() KeyPairOutput {
	return i.ToKeyPairOutputWithContext(context.Background())
}

func (i *KeyPair) ToKeyPairOutputWithContext(ctx context.Context) KeyPairOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairOutput)
}

func (i *KeyPair) ToKeyPairPtrOutput() KeyPairPtrOutput {
	return i.ToKeyPairPtrOutputWithContext(context.Background())
}

func (i *KeyPair) ToKeyPairPtrOutputWithContext(ctx context.Context) KeyPairPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairPtrOutput)
}

type KeyPairPtrInput interface {
	pulumi.Input

	ToKeyPairPtrOutput() KeyPairPtrOutput
	ToKeyPairPtrOutputWithContext(ctx context.Context) KeyPairPtrOutput
}

type keyPairPtrType KeyPairArgs

func (*keyPairPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyPair)(nil))
}

func (i *keyPairPtrType) ToKeyPairPtrOutput() KeyPairPtrOutput {
	return i.ToKeyPairPtrOutputWithContext(context.Background())
}

func (i *keyPairPtrType) ToKeyPairPtrOutputWithContext(ctx context.Context) KeyPairPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairPtrOutput)
}

// KeyPairArrayInput is an input type that accepts KeyPairArray and KeyPairArrayOutput values.
// You can construct a concrete instance of `KeyPairArrayInput` via:
//
//          KeyPairArray{ KeyPairArgs{...} }
type KeyPairArrayInput interface {
	pulumi.Input

	ToKeyPairArrayOutput() KeyPairArrayOutput
	ToKeyPairArrayOutputWithContext(context.Context) KeyPairArrayOutput
}

type KeyPairArray []KeyPairInput

func (KeyPairArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KeyPair)(nil)).Elem()
}

func (i KeyPairArray) ToKeyPairArrayOutput() KeyPairArrayOutput {
	return i.ToKeyPairArrayOutputWithContext(context.Background())
}

func (i KeyPairArray) ToKeyPairArrayOutputWithContext(ctx context.Context) KeyPairArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairArrayOutput)
}

// KeyPairMapInput is an input type that accepts KeyPairMap and KeyPairMapOutput values.
// You can construct a concrete instance of `KeyPairMapInput` via:
//
//          KeyPairMap{ "key": KeyPairArgs{...} }
type KeyPairMapInput interface {
	pulumi.Input

	ToKeyPairMapOutput() KeyPairMapOutput
	ToKeyPairMapOutputWithContext(context.Context) KeyPairMapOutput
}

type KeyPairMap map[string]KeyPairInput

func (KeyPairMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KeyPair)(nil)).Elem()
}

func (i KeyPairMap) ToKeyPairMapOutput() KeyPairMapOutput {
	return i.ToKeyPairMapOutputWithContext(context.Background())
}

func (i KeyPairMap) ToKeyPairMapOutputWithContext(ctx context.Context) KeyPairMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairMapOutput)
}

type KeyPairOutput struct{ *pulumi.OutputState }

func (KeyPairOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyPair)(nil))
}

func (o KeyPairOutput) ToKeyPairOutput() KeyPairOutput {
	return o
}

func (o KeyPairOutput) ToKeyPairOutputWithContext(ctx context.Context) KeyPairOutput {
	return o
}

func (o KeyPairOutput) ToKeyPairPtrOutput() KeyPairPtrOutput {
	return o.ToKeyPairPtrOutputWithContext(context.Background())
}

func (o KeyPairOutput) ToKeyPairPtrOutputWithContext(ctx context.Context) KeyPairPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyPair) *KeyPair {
		return &v
	}).(KeyPairPtrOutput)
}

type KeyPairPtrOutput struct{ *pulumi.OutputState }

func (KeyPairPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyPair)(nil))
}

func (o KeyPairPtrOutput) ToKeyPairPtrOutput() KeyPairPtrOutput {
	return o
}

func (o KeyPairPtrOutput) ToKeyPairPtrOutputWithContext(ctx context.Context) KeyPairPtrOutput {
	return o
}

func (o KeyPairPtrOutput) Elem() KeyPairOutput {
	return o.ApplyT(func(v *KeyPair) KeyPair {
		if v != nil {
			return *v
		}
		var ret KeyPair
		return ret
	}).(KeyPairOutput)
}

type KeyPairArrayOutput struct{ *pulumi.OutputState }

func (KeyPairArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyPair)(nil))
}

func (o KeyPairArrayOutput) ToKeyPairArrayOutput() KeyPairArrayOutput {
	return o
}

func (o KeyPairArrayOutput) ToKeyPairArrayOutputWithContext(ctx context.Context) KeyPairArrayOutput {
	return o
}

func (o KeyPairArrayOutput) Index(i pulumi.IntInput) KeyPairOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyPair {
		return vs[0].([]KeyPair)[vs[1].(int)]
	}).(KeyPairOutput)
}

type KeyPairMapOutput struct{ *pulumi.OutputState }

func (KeyPairMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KeyPair)(nil))
}

func (o KeyPairMapOutput) ToKeyPairMapOutput() KeyPairMapOutput {
	return o
}

func (o KeyPairMapOutput) ToKeyPairMapOutputWithContext(ctx context.Context) KeyPairMapOutput {
	return o
}

func (o KeyPairMapOutput) MapIndex(k pulumi.StringInput) KeyPairOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) KeyPair {
		return vs[0].(map[string]KeyPair)[vs[1].(string)]
	}).(KeyPairOutput)
}

func init() {
	pulumi.RegisterOutputType(KeyPairOutput{})
	pulumi.RegisterOutputType(KeyPairPtrOutput{})
	pulumi.RegisterOutputType(KeyPairArrayOutput{})
	pulumi.RegisterOutputType(KeyPairMapOutput{})
}
