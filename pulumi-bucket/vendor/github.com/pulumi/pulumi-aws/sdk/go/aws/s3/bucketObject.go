// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package s3

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a S3 bucket object resource.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/s3_bucket_object.html.markdown.
type BucketObject struct {
	pulumi.CustomResourceState

	// The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".
	Acl pulumi.StringPtrOutput `pulumi:"acl"`
	// The name of the bucket to put the file in.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Specifies caching behavior along the request/reply chain Read [w3c cacheControl](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
	CacheControl pulumi.StringPtrOutput `pulumi:"cacheControl"`
	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	Content pulumi.StringPtrOutput `pulumi:"content"`
	// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
	ContentBase64 pulumi.StringPtrOutput `pulumi:"contentBase64"`
	// Specifies presentational information for the object. Read [w3c contentDisposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
	ContentDisposition pulumi.StringPtrOutput `pulumi:"contentDisposition"`
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
	ContentEncoding pulumi.StringPtrOutput `pulumi:"contentEncoding"`
	// The language the content is in e.g. en-US or en-GB.
	ContentLanguage pulumi.StringPtrOutput `pulumi:"contentLanguage"`
	// A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType pulumi.StringOutput `pulumi:"contentType"`
	// Used to trigger updates. The only meaningful value is `${filemd5("path/to/file")}` (this provider 0.11.12 or later) or `${md5(file("path/to/file"))}` (this provider 0.11.11 or earlier).
	// This attribute is not compatible with KMS encryption, `kmsKeyId` or `serverSideEncryption = "aws:kms"`.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Allow the object to be deleted by removing any legal hold on any object version.
	// Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// The name of the object once it is in the bucket.
	Key pulumi.StringOutput `pulumi:"key"`
	// Specifies the AWS KMS Key ARN to use for object encryption.
	// This value is a fully qualified **ARN** of the KMS Key. If using `kms.Key`,
	// use the exported `arn` attribute:
	// `kmsKeyId = "${aws_kms_key.foo.arn}"`
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// A mapping of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The [legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds) status that you want to apply to the specified object. Valid values are `ON` and `OFF`.
	ObjectLockLegalHoldStatus pulumi.StringPtrOutput `pulumi:"objectLockLegalHoldStatus"`
	// The object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) that you want to apply to this object. Valid values are `GOVERNANCE` and `COMPLIANCE`.
	ObjectLockMode pulumi.StringPtrOutput `pulumi:"objectLockMode"`
	// The date and time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8), when this object's object lock will [expire](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-periods).
	ObjectLockRetainUntilDate pulumi.StringPtrOutput `pulumi:"objectLockRetainUntilDate"`
	// Specifies server-side encryption of the object in S3. Valid values are "`AES256`" and "`aws:kms`".
	ServerSideEncryption pulumi.StringOutput `pulumi:"serverSideEncryption"`
	// The path to a file that will be read and uploaded as raw bytes for the object content.
	Source pulumi.AssetOrArchiveOutput `pulumi:"source"`
	// Specifies the desired [Storage Class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html)
	// for the object. Can be either "`STANDARD`", "`REDUCED_REDUNDANCY`", "`ONEZONE_IA`", "`INTELLIGENT_TIERING`", "`GLACIER`", "`DEEP_ARCHIVE`", or "`STANDARD_IA`". Defaults to "`STANDARD`".
	StorageClass pulumi.StringOutput `pulumi:"storageClass"`
	// A mapping of tags to assign to the object.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// A unique version ID value for the object, if bucket versioning
	// is enabled.
	VersionId pulumi.StringOutput `pulumi:"versionId"`
	// Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
	WebsiteRedirect pulumi.StringPtrOutput `pulumi:"websiteRedirect"`
}

// NewBucketObject registers a new resource with the given unique name, arguments, and options.
func NewBucketObject(ctx *pulumi.Context,
	name string, args *BucketObjectArgs, opts ...pulumi.ResourceOption) (*BucketObject, error) {
	if args == nil || args.Bucket == nil {
		return nil, errors.New("missing required argument 'Bucket'")
	}
	if args == nil {
		args = &BucketObjectArgs{}
	}
	var resource BucketObject
	err := ctx.RegisterResource("aws:s3/bucketObject:BucketObject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketObject gets an existing BucketObject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketObject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketObjectState, opts ...pulumi.ResourceOption) (*BucketObject, error) {
	var resource BucketObject
	err := ctx.ReadResource("aws:s3/bucketObject:BucketObject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketObject resources.
type bucketObjectState struct {
	// The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".
	Acl *string `pulumi:"acl"`
	// The name of the bucket to put the file in.
	Bucket *string `pulumi:"bucket"`
	// Specifies caching behavior along the request/reply chain Read [w3c cacheControl](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
	CacheControl *string `pulumi:"cacheControl"`
	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	Content *string `pulumi:"content"`
	// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
	ContentBase64 *string `pulumi:"contentBase64"`
	// Specifies presentational information for the object. Read [w3c contentDisposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
	ContentDisposition *string `pulumi:"contentDisposition"`
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
	ContentEncoding *string `pulumi:"contentEncoding"`
	// The language the content is in e.g. en-US or en-GB.
	ContentLanguage *string `pulumi:"contentLanguage"`
	// A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType *string `pulumi:"contentType"`
	// Used to trigger updates. The only meaningful value is `${filemd5("path/to/file")}` (this provider 0.11.12 or later) or `${md5(file("path/to/file"))}` (this provider 0.11.11 or earlier).
	// This attribute is not compatible with KMS encryption, `kmsKeyId` or `serverSideEncryption = "aws:kms"`.
	Etag *string `pulumi:"etag"`
	// Allow the object to be deleted by removing any legal hold on any object version.
	// Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// The name of the object once it is in the bucket.
	Key *string `pulumi:"key"`
	// Specifies the AWS KMS Key ARN to use for object encryption.
	// This value is a fully qualified **ARN** of the KMS Key. If using `kms.Key`,
	// use the exported `arn` attribute:
	// `kmsKeyId = "${aws_kms_key.foo.arn}"`
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// A mapping of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
	Metadata map[string]string `pulumi:"metadata"`
	// The [legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds) status that you want to apply to the specified object. Valid values are `ON` and `OFF`.
	ObjectLockLegalHoldStatus *string `pulumi:"objectLockLegalHoldStatus"`
	// The object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) that you want to apply to this object. Valid values are `GOVERNANCE` and `COMPLIANCE`.
	ObjectLockMode *string `pulumi:"objectLockMode"`
	// The date and time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8), when this object's object lock will [expire](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-periods).
	ObjectLockRetainUntilDate *string `pulumi:"objectLockRetainUntilDate"`
	// Specifies server-side encryption of the object in S3. Valid values are "`AES256`" and "`aws:kms`".
	ServerSideEncryption *string `pulumi:"serverSideEncryption"`
	// The path to a file that will be read and uploaded as raw bytes for the object content.
	Source pulumi.AssetOrArchive `pulumi:"source"`
	// Specifies the desired [Storage Class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html)
	// for the object. Can be either "`STANDARD`", "`REDUCED_REDUNDANCY`", "`ONEZONE_IA`", "`INTELLIGENT_TIERING`", "`GLACIER`", "`DEEP_ARCHIVE`", or "`STANDARD_IA`". Defaults to "`STANDARD`".
	StorageClass *string `pulumi:"storageClass"`
	// A mapping of tags to assign to the object.
	Tags map[string]interface{} `pulumi:"tags"`
	// A unique version ID value for the object, if bucket versioning
	// is enabled.
	VersionId *string `pulumi:"versionId"`
	// Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
	WebsiteRedirect *string `pulumi:"websiteRedirect"`
}

type BucketObjectState struct {
	// The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".
	Acl pulumi.StringPtrInput
	// The name of the bucket to put the file in.
	Bucket pulumi.StringPtrInput
	// Specifies caching behavior along the request/reply chain Read [w3c cacheControl](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
	CacheControl pulumi.StringPtrInput
	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	Content pulumi.StringPtrInput
	// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
	ContentBase64 pulumi.StringPtrInput
	// Specifies presentational information for the object. Read [w3c contentDisposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
	ContentDisposition pulumi.StringPtrInput
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
	ContentEncoding pulumi.StringPtrInput
	// The language the content is in e.g. en-US or en-GB.
	ContentLanguage pulumi.StringPtrInput
	// A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType pulumi.StringPtrInput
	// Used to trigger updates. The only meaningful value is `${filemd5("path/to/file")}` (this provider 0.11.12 or later) or `${md5(file("path/to/file"))}` (this provider 0.11.11 or earlier).
	// This attribute is not compatible with KMS encryption, `kmsKeyId` or `serverSideEncryption = "aws:kms"`.
	Etag pulumi.StringPtrInput
	// Allow the object to be deleted by removing any legal hold on any object version.
	// Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
	ForceDestroy pulumi.BoolPtrInput
	// The name of the object once it is in the bucket.
	Key pulumi.StringPtrInput
	// Specifies the AWS KMS Key ARN to use for object encryption.
	// This value is a fully qualified **ARN** of the KMS Key. If using `kms.Key`,
	// use the exported `arn` attribute:
	// `kmsKeyId = "${aws_kms_key.foo.arn}"`
	KmsKeyId pulumi.StringPtrInput
	// A mapping of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
	Metadata pulumi.StringMapInput
	// The [legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds) status that you want to apply to the specified object. Valid values are `ON` and `OFF`.
	ObjectLockLegalHoldStatus pulumi.StringPtrInput
	// The object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) that you want to apply to this object. Valid values are `GOVERNANCE` and `COMPLIANCE`.
	ObjectLockMode pulumi.StringPtrInput
	// The date and time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8), when this object's object lock will [expire](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-periods).
	ObjectLockRetainUntilDate pulumi.StringPtrInput
	// Specifies server-side encryption of the object in S3. Valid values are "`AES256`" and "`aws:kms`".
	ServerSideEncryption pulumi.StringPtrInput
	// The path to a file that will be read and uploaded as raw bytes for the object content.
	Source pulumi.AssetOrArchiveInput
	// Specifies the desired [Storage Class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html)
	// for the object. Can be either "`STANDARD`", "`REDUCED_REDUNDANCY`", "`ONEZONE_IA`", "`INTELLIGENT_TIERING`", "`GLACIER`", "`DEEP_ARCHIVE`", or "`STANDARD_IA`". Defaults to "`STANDARD`".
	StorageClass pulumi.StringPtrInput
	// A mapping of tags to assign to the object.
	Tags pulumi.MapInput
	// A unique version ID value for the object, if bucket versioning
	// is enabled.
	VersionId pulumi.StringPtrInput
	// Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
	WebsiteRedirect pulumi.StringPtrInput
}

func (BucketObjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketObjectState)(nil)).Elem()
}

type bucketObjectArgs struct {
	// The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".
	Acl *string `pulumi:"acl"`
	// The name of the bucket to put the file in.
	Bucket interface{} `pulumi:"bucket"`
	// Specifies caching behavior along the request/reply chain Read [w3c cacheControl](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
	CacheControl *string `pulumi:"cacheControl"`
	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	Content *string `pulumi:"content"`
	// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
	ContentBase64 *string `pulumi:"contentBase64"`
	// Specifies presentational information for the object. Read [w3c contentDisposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
	ContentDisposition *string `pulumi:"contentDisposition"`
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
	ContentEncoding *string `pulumi:"contentEncoding"`
	// The language the content is in e.g. en-US or en-GB.
	ContentLanguage *string `pulumi:"contentLanguage"`
	// A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType *string `pulumi:"contentType"`
	// Used to trigger updates. The only meaningful value is `${filemd5("path/to/file")}` (this provider 0.11.12 or later) or `${md5(file("path/to/file"))}` (this provider 0.11.11 or earlier).
	// This attribute is not compatible with KMS encryption, `kmsKeyId` or `serverSideEncryption = "aws:kms"`.
	Etag *string `pulumi:"etag"`
	// Allow the object to be deleted by removing any legal hold on any object version.
	// Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// The name of the object once it is in the bucket.
	Key *string `pulumi:"key"`
	// Specifies the AWS KMS Key ARN to use for object encryption.
	// This value is a fully qualified **ARN** of the KMS Key. If using `kms.Key`,
	// use the exported `arn` attribute:
	// `kmsKeyId = "${aws_kms_key.foo.arn}"`
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// A mapping of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
	Metadata map[string]string `pulumi:"metadata"`
	// The [legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds) status that you want to apply to the specified object. Valid values are `ON` and `OFF`.
	ObjectLockLegalHoldStatus *string `pulumi:"objectLockLegalHoldStatus"`
	// The object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) that you want to apply to this object. Valid values are `GOVERNANCE` and `COMPLIANCE`.
	ObjectLockMode *string `pulumi:"objectLockMode"`
	// The date and time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8), when this object's object lock will [expire](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-periods).
	ObjectLockRetainUntilDate *string `pulumi:"objectLockRetainUntilDate"`
	// Specifies server-side encryption of the object in S3. Valid values are "`AES256`" and "`aws:kms`".
	ServerSideEncryption *string `pulumi:"serverSideEncryption"`
	// The path to a file that will be read and uploaded as raw bytes for the object content.
	Source pulumi.AssetOrArchive `pulumi:"source"`
	// Specifies the desired [Storage Class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html)
	// for the object. Can be either "`STANDARD`", "`REDUCED_REDUNDANCY`", "`ONEZONE_IA`", "`INTELLIGENT_TIERING`", "`GLACIER`", "`DEEP_ARCHIVE`", or "`STANDARD_IA`". Defaults to "`STANDARD`".
	StorageClass *string `pulumi:"storageClass"`
	// A mapping of tags to assign to the object.
	Tags map[string]interface{} `pulumi:"tags"`
	// Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
	WebsiteRedirect *string `pulumi:"websiteRedirect"`
}

// The set of arguments for constructing a BucketObject resource.
type BucketObjectArgs struct {
	// The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".
	Acl pulumi.StringPtrInput
	// The name of the bucket to put the file in.
	Bucket pulumi.Input
	// Specifies caching behavior along the request/reply chain Read [w3c cacheControl](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
	CacheControl pulumi.StringPtrInput
	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	Content pulumi.StringPtrInput
	// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
	ContentBase64 pulumi.StringPtrInput
	// Specifies presentational information for the object. Read [w3c contentDisposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
	ContentDisposition pulumi.StringPtrInput
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
	ContentEncoding pulumi.StringPtrInput
	// The language the content is in e.g. en-US or en-GB.
	ContentLanguage pulumi.StringPtrInput
	// A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
	ContentType pulumi.StringPtrInput
	// Used to trigger updates. The only meaningful value is `${filemd5("path/to/file")}` (this provider 0.11.12 or later) or `${md5(file("path/to/file"))}` (this provider 0.11.11 or earlier).
	// This attribute is not compatible with KMS encryption, `kmsKeyId` or `serverSideEncryption = "aws:kms"`.
	Etag pulumi.StringPtrInput
	// Allow the object to be deleted by removing any legal hold on any object version.
	// Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
	ForceDestroy pulumi.BoolPtrInput
	// The name of the object once it is in the bucket.
	Key pulumi.StringPtrInput
	// Specifies the AWS KMS Key ARN to use for object encryption.
	// This value is a fully qualified **ARN** of the KMS Key. If using `kms.Key`,
	// use the exported `arn` attribute:
	// `kmsKeyId = "${aws_kms_key.foo.arn}"`
	KmsKeyId pulumi.StringPtrInput
	// A mapping of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
	Metadata pulumi.StringMapInput
	// The [legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds) status that you want to apply to the specified object. Valid values are `ON` and `OFF`.
	ObjectLockLegalHoldStatus pulumi.StringPtrInput
	// The object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) that you want to apply to this object. Valid values are `GOVERNANCE` and `COMPLIANCE`.
	ObjectLockMode pulumi.StringPtrInput
	// The date and time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8), when this object's object lock will [expire](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-periods).
	ObjectLockRetainUntilDate pulumi.StringPtrInput
	// Specifies server-side encryption of the object in S3. Valid values are "`AES256`" and "`aws:kms`".
	ServerSideEncryption pulumi.StringPtrInput
	// The path to a file that will be read and uploaded as raw bytes for the object content.
	Source pulumi.AssetOrArchiveInput
	// Specifies the desired [Storage Class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html)
	// for the object. Can be either "`STANDARD`", "`REDUCED_REDUNDANCY`", "`ONEZONE_IA`", "`INTELLIGENT_TIERING`", "`GLACIER`", "`DEEP_ARCHIVE`", or "`STANDARD_IA`". Defaults to "`STANDARD`".
	StorageClass pulumi.StringPtrInput
	// A mapping of tags to assign to the object.
	Tags pulumi.MapInput
	// Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
	WebsiteRedirect pulumi.StringPtrInput
}

func (BucketObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketObjectArgs)(nil)).Elem()
}
