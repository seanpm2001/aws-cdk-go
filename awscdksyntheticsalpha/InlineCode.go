package awscdksyntheticsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdksyntheticsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Canary code from an inline string.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import synthetics_alpha "github.com/aws/aws-cdk-go/awscdksyntheticsalpha"
//
//   inlineCode := synthetics_alpha.NewInlineCode(jsii.String("code"))
//
// Experimental.
type InlineCode interface {
	Code
	// Called when the canary is initialized to allow this object to bind to the stack, add resources and have fun.
	// Experimental.
	Bind(_scope constructs.Construct, handler *string, _family RuntimeFamily) *CodeConfig
}

// The jsii proxy struct for InlineCode
type jsiiProxy_InlineCode struct {
	jsiiProxy_Code
}

// Experimental.
func NewInlineCode(code *string) InlineCode {
	_init_.Initialize()

	if err := validateNewInlineCodeParameters(code); err != nil {
		panic(err)
	}
	j := jsiiProxy_InlineCode{}

	_jsii_.Create(
		"@aws-cdk/aws-synthetics-alpha.InlineCode",
		[]interface{}{code},
		&j,
	)

	return &j
}

// Experimental.
func NewInlineCode_Override(i InlineCode, code *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-synthetics-alpha.InlineCode",
		[]interface{}{code},
		i,
	)
}

// Specify code from a local path.
//
// Path must include the folder structure `nodejs/node_modules/myCanaryFilename.js`.
//
// Returns: `AssetCode` associated with the specified path.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_WritingCanary.html#CloudWatch_Synthetics_Canaries_write_from_scratch
//
// Experimental.
func InlineCode_FromAsset(assetPath *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateInlineCode_FromAssetParameters(assetPath, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.InlineCode",
		"fromAsset",
		[]interface{}{assetPath, options},
		&returns,
	)

	return returns
}

// Specify code from an s3 bucket.
//
// The object in the s3 bucket must be a .zip file that contains
// the structure `nodejs/node_modules/myCanaryFilename.js`.
//
// Returns: `S3Code` associated with the specified S3 object.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_WritingCanary.html#CloudWatch_Synthetics_Canaries_write_from_scratch
//
// Experimental.
func InlineCode_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	if err := validateInlineCode_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3Code

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.InlineCode",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Specify code inline.
//
// Returns: `InlineCode` with inline code.
// Experimental.
func InlineCode_FromInline(code *string) InlineCode {
	_init_.Initialize()

	if err := validateInlineCode_FromInlineParameters(code); err != nil {
		panic(err)
	}
	var returns InlineCode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.InlineCode",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InlineCode) Bind(_scope constructs.Construct, handler *string, _family RuntimeFamily) *CodeConfig {
	if err := i.validateBindParameters(_scope, handler, _family); err != nil {
		panic(err)
	}
	var returns *CodeConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{_scope, handler, _family},
		&returns,
	)

	return returns
}

