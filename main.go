package main

import (
	"io/ioutil"

	"github.com/pulumi/pulumi-aws/sdk/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

const size = "t2.micro"

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		err := pulumiForAutomation(ctx)

		if err != nil {
			return err
		}

		return createJenkinsVM(ctx)
	})
}

func getCloudInitYaml(fileName string) (*string, error) {
	b, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		return nil, err
	}
	yaml := string(b)
	return &yaml, nil
}

func createJenkinsVM(ctx *pulumi.Context) error {
	group, err := ec2.NewSecurityGroup(ctx, "web-secgrp-2", &ec2.SecurityGroupArgs{
		Description: pulumi.String("Enable HTTP access"),
		Ingress: ec2.SecurityGroupIngressArray{
			ec2.SecurityGroupIngressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(80),
				ToPort:     pulumi.Int(80),
				CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
			},
			ec2.SecurityGroupIngressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(22),
				ToPort:     pulumi.Int(22),
				CidrBlocks: pulumi.StringArray{pulumi.String("95.90.242.194/32")},
			},
		},
		Egress: ec2.SecurityGroupEgressArray{
			ec2.SecurityGroupEgressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(80),
				ToPort:     pulumi.Int(80),
				CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
			},
			ec2.SecurityGroupEgressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(443),
				ToPort:     pulumi.Int(443),
				CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
			},
			//github ssh
			//140.82.118.3
			ec2.SecurityGroupEgressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(22),
				ToPort:     pulumi.Int(22),
				CidrBlocks: pulumi.StringArray{pulumi.String("192.30.252.0/22")},
			},
			ec2.SecurityGroupEgressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(22),
				ToPort:     pulumi.Int(22),
				CidrBlocks: pulumi.StringArray{pulumi.String("140.82.118.0/24")},
			},
			ec2.SecurityGroupEgressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(22),
				ToPort:     pulumi.Int(22),
				CidrBlocks: pulumi.StringArray{pulumi.String("204.232.175.90/32")},
			},
			ec2.SecurityGroupEgressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(22),
				ToPort:     pulumi.Int(22),
				CidrBlocks: pulumi.StringArray{pulumi.String("207.97.227.239/32")},
			},
		},
	})
	if err != nil {
		return err
	}

	mostRecent := true
	ami, err := aws.GetAmi(ctx, &aws.GetAmiArgs{
		Filters: []aws.GetAmiFilter{
			{
				Name:   "name",
				Values: []string{"ubuntu/images/hvm-ssd/ubuntu-xenial-16.04-amd64-server-20200129"},
			},
		},
		Owners:     []string{"099720109477"},
		MostRecent: &mostRecent,
	})

	if err != nil {
		return err
	}

	yaml, err := getCloudInitYaml("cloud-init/cloud-init.yaml")

	if err != nil {
		return err
	}

	server, err := ec2.NewInstance(ctx, "web-server-www", &ec2.InstanceArgs{
		InstanceType: pulumi.String(size),
		SecurityGroups: pulumi.StringArray{
			group.Name,
		},
		KeyName:  pulumi.String("test"), //create the keypair with pulumi
		Ami:      pulumi.String(ami.Id),
		UserData: pulumi.String(*yaml),
	})

	ctx.Export("publicIp", server.PublicIp)
	ctx.Export("publicDns", server.PublicDns)

	return err
}

func pulumiForAutomation(ctx *pulumi.Context) error {
	const bucketName = "s3-pulumi-state"
	_, err := s3.NewAccountPublicAccessBlock(ctx, bucketName+"-acl", &s3.AccountPublicAccessBlockArgs{
		BlockPublicAcls: pulumi.Bool(true),
	})
	if err != nil {
		return err
	}

	stateBucket, err := s3.NewBucket(ctx, bucketName, &s3.BucketArgs{})
	if err != nil {
		return err
	}

	// Stack exports
	ctx.Export("bucketName", stateBucket.ID())
	ctx.Export("s3Urn", stateBucket.URN())

	// Create the bucket to store the pulumi state
	const username = "pulumi-automation"

	iamUser, err := iam.NewUser(ctx, username+"-user", &iam.UserArgs{
		Tags: pulumi.Map{"Creator": pulumi.String("jenkins-aws-pulumi")},
	})

	if err != nil {
		return err
	}

	var s3PolicyContent = `{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Action": "s3:*",
				"Resource": "*"
			}
		]
	}`

	var ec2PolicyContent = `{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Action": "ec2:*",
				"Effect": "Allow",
				"Resource": "*"
			}
		]
	}
	`

	s3IamPolicy, err := iam.NewPolicy(ctx, username+"-user-policy-s3", &iam.PolicyArgs{
		Policy: pulumi.String(s3PolicyContent),
	})

	if err != nil {
		return err
	}

	ec2IamPolicy, err := iam.NewPolicy(ctx, username+"-user-policy-ec2", &iam.PolicyArgs{
		Policy: pulumi.String(ec2PolicyContent),
	})

	if err != nil {
		return err
	}

	iam.NewUserPolicyAttachment(ctx, username+"-user-policy-attachment-s3", &iam.UserPolicyAttachmentArgs{
		User:      iamUser.ID(),
		PolicyArn: s3IamPolicy.ID(),
	})

	iam.NewUserPolicyAttachment(ctx, username+"-user-policy-attachment-ec2", &iam.UserPolicyAttachmentArgs{
		User:      iamUser.ID(),
		PolicyArn: ec2IamPolicy.ID(),
	})

	iamKeys, err := iam.NewAccessKey(ctx, username+"-user-keys", &iam.AccessKeyArgs{
		User: iamUser.ID(),
	})

	if err != nil {
		return err
	}

	ctx.Export("AccessKeysSecret", iamKeys.Secret)
	ctx.Export("AccessKeys", iamKeys.ID())

	return nil
}
