package main

import (
	"github.com/pulumi/pulumi-aws/sdk/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

const size = "t2.micro"

func main() {

	pulumi.Run(func(ctx *pulumi.Context) error {
		return pulumiForAutomation(ctx)
	})
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
	ctx.Export("s3Domain", stateBucket.BucketDomainName)

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

	var iamPolicyContent = `{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Action": "iam:*",
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

	iamIamPolicy, err := iam.NewPolicy(ctx, username+"-user-policy-iam", &iam.PolicyArgs{
		Policy: pulumi.String(iamPolicyContent),
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

	iam.NewUserPolicyAttachment(ctx, username+"-user-policy-attachment-iam", &iam.UserPolicyAttachmentArgs{
		User:      iamUser.ID(),
		PolicyArn: iamIamPolicy.ID(),
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
