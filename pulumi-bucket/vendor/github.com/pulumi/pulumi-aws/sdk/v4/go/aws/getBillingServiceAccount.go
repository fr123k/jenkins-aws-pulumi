// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the Account ID of the [AWS Billing and Cost Management Service Account](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/billing-getting-started.html#step-2) for the purpose of permitting in S3 bucket policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		main, err := aws.GetBillingServiceAccount(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewBucket(ctx, "billingLogs", &s3.BucketArgs{
// 			Acl:    pulumi.String("private"),
// 			Policy: pulumi.Any(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Id\": \"Policy\",\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": [\n", "        \"s3:GetBucketAcl\", \"s3:GetBucketPolicy\"\n", "      ],\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": \"arn:aws:s3:::my-billing-tf-test-bucket\",\n", "      \"Principal\": {\n", "        \"AWS\": [\n", "          \"", main.Arn, "\"\n", "        ]\n", "      }\n", "    },\n", "    {\n", "      \"Action\": [\n", "        \"s3:PutObject\"\n", "      ],\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": \"arn:aws:s3:::my-billing-tf-test-bucket/*\",\n", "      \"Principal\": {\n", "        \"AWS\": [\n", "          \"", main.Arn, "\"\n", "        ]\n", "      }\n", "    }\n", "  ]\n", "}\n", "\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetBillingServiceAccount(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetBillingServiceAccountResult, error) {
	var rv GetBillingServiceAccountResult
	err := ctx.Invoke("aws:index/getBillingServiceAccount:getBillingServiceAccount", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getBillingServiceAccount.
type GetBillingServiceAccountResult struct {
	// The ARN of the AWS billing service account.
	Arn string `pulumi:"arn"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}