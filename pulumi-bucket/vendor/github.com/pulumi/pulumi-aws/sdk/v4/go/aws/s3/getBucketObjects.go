// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// > **NOTE on `maxKeys`:** Retrieving very large numbers of keys can adversely affect this provider's performance.
//
// The bucket-objects data source returns keys (i.e., file names) and other metadata about objects in an S3 bucket.
func GetBucketObjects(ctx *pulumi.Context, args *GetBucketObjectsArgs, opts ...pulumi.InvokeOption) (*GetBucketObjectsResult, error) {
	var rv GetBucketObjectsResult
	err := ctx.Invoke("aws:s3/getBucketObjects:getBucketObjects", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBucketObjects.
type GetBucketObjectsArgs struct {
	// Lists object keys in this S3 bucket. Alternatively, an [S3 access point](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) ARN can be specified
	Bucket string `pulumi:"bucket"`
	// A character used to group keys (Default: none)
	Delimiter *string `pulumi:"delimiter"`
	// Encodes keys using this method (Default: none; besides none, only "url" can be used)
	EncodingType *string `pulumi:"encodingType"`
	// Boolean specifying whether to populate the owner list (Default: false)
	FetchOwner *bool `pulumi:"fetchOwner"`
	// Maximum object keys to return (Default: 1000)
	MaxKeys *int `pulumi:"maxKeys"`
	// Limits results to object keys with this prefix (Default: none)
	Prefix *string `pulumi:"prefix"`
	// Returns key names lexicographically after a specific object key in your bucket (Default: none; S3 lists object keys in UTF-8 character encoding in lexicographical order)
	StartAfter *string `pulumi:"startAfter"`
}

// A collection of values returned by getBucketObjects.
type GetBucketObjectsResult struct {
	Bucket string `pulumi:"bucket"`
	// List of any keys between `prefix` and the next occurrence of `delimiter` (i.e., similar to subdirectories of the `prefix` "directory"); the list is only returned when you specify `delimiter`
	CommonPrefixes []string `pulumi:"commonPrefixes"`
	Delimiter      *string  `pulumi:"delimiter"`
	EncodingType   *string  `pulumi:"encodingType"`
	FetchOwner     *bool    `pulumi:"fetchOwner"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of strings representing object keys
	Keys    []string `pulumi:"keys"`
	MaxKeys *int     `pulumi:"maxKeys"`
	// List of strings representing object owner IDs (see `fetchOwner` above)
	Owners     []string `pulumi:"owners"`
	Prefix     *string  `pulumi:"prefix"`
	StartAfter *string  `pulumi:"startAfter"`
}
