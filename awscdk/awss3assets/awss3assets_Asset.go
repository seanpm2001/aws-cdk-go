package awss3assets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// An asset represents a local file or directory, which is automatically uploaded to S3 and then can be referenced within a CDK application.
//
// Example:
//   import s3Assets "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//
//   chartAsset := s3Assets.NewAsset(this, jsii.String("ChartAsset"), &AssetProps{
//   	Path: jsii.String("/path/to/asset"),
//   })
//
//   cluster.addHelmChart(jsii.String("test-chart"), &HelmChartOptions{
//   	ChartAsset: chartAsset,
//   })
//
// Experimental.
type Asset interface {
	awscdk.Construct
	awscdk.IAsset
	// A hash of this asset, which is available at construction time.
	//
	// As this is a plain string, it
	// can be used in construct IDs in order to enforce creation of a new resource when the content
	// hash has changed.
	// Experimental.
	AssetHash() *string
	// The path to the asset, relative to the current Cloud Assembly.
	//
	// If asset staging is disabled, this will just be the original path.
	// If asset staging is enabled it will be the staged path.
	// Experimental.
	AssetPath() *string
	// The S3 bucket in which this asset resides.
	// Experimental.
	Bucket() awss3.IBucket
	// Attribute which represents the S3 HTTP URL of this asset.
	//
	// For example, `https://s3.us-west-1.amazonaws.com/bucket/key`
	// Experimental.
	HttpUrl() *string
	// Indicates if this asset is a single file.
	//
	// Allows constructs to ensure that the
	// correct file type was used.
	// Experimental.
	IsFile() *bool
	// Indicates if this asset is a zip archive.
	//
	// Allows constructs to ensure that the
	// correct file type was used.
	// Experimental.
	IsZipArchive() *bool
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Attribute that represents the name of the bucket this asset exists in.
	// Experimental.
	S3BucketName() *string
	// Attribute which represents the S3 object key of this asset.
	// Experimental.
	S3ObjectKey() *string
	// Attribute which represents the S3 URL of this asset.
	//
	// For example, `s3://bucket/key`.
	// Experimental.
	S3ObjectUrl() *string
	// Attribute which represents the S3 URL of this asset.
	// Deprecated: use `httpUrl`.
	S3Url() *string
	// A cryptographic hash of the asset.
	// Deprecated: see `assetHash`.
	SourceHash() *string
	// Adds CloudFormation template metadata to the specified resource with information that indicates which resource property is mapped to this local asset.
	//
	// This can be used by tools such as SAM CLI to provide local
	// experience such as local invocation and debugging of Lambda functions.
	//
	// Asset metadata will only be included if the stack is synthesized with the
	// "aws:cdk:enable-asset-metadata" context key defined, which is the default
	// behavior when synthesizing via the CDK Toolkit.
	// See: https://github.com/aws/aws-cdk/issues/1432
	//
	// Experimental.
	AddResourceMetadata(resource awscdk.CfnResource, resourceProperty *string)
	// Grants read permissions to the principal on the assets bucket.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for Asset
type jsiiProxy_Asset struct {
	internal.Type__awscdkConstruct
	internal.Type__awscdkIAsset
}

func (j *jsiiProxy_Asset) AssetHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) AssetPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) Bucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) HttpUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) IsFile() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) IsZipArchive() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isZipArchive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) S3ObjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) S3ObjectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) S3Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) SourceHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceHash",
		&returns,
	)
	return returns
}


// Experimental.
func NewAsset(scope constructs.Construct, id *string, props *AssetProps) Asset {
	_init_.Initialize()

	if err := validateNewAssetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Asset{}

	_jsii_.Create(
		"monocdk.aws_s3_assets.Asset",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAsset_Override(a Asset, scope constructs.Construct, id *string, props *AssetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_s3_assets.Asset",
		[]interface{}{scope, id, props},
		a,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func Asset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAsset_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_s3_assets.Asset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Asset) AddResourceMetadata(resource awscdk.CfnResource, resourceProperty *string) {
	if err := a.validateAddResourceMetadataParameters(resource, resourceProperty); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addResourceMetadata",
		[]interface{}{resource, resourceProperty},
	)
}

func (a *jsiiProxy_Asset) GrantRead(grantee awsiam.IGrantable) {
	if err := a.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"grantRead",
		[]interface{}{grantee},
	)
}

func (a *jsiiProxy_Asset) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Asset) OnSynthesize(session constructs.ISynthesisSession) {
	if err := a.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (a *jsiiProxy_Asset) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Asset) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Asset) Synthesize(session awscdk.ISynthesisSession) {
	if err := a.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
}

func (a *jsiiProxy_Asset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Asset) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

