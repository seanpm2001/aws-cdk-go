package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Amazon CloudFront is a global content delivery network (CDN) service that securely delivers data, videos, applications, and APIs to your viewers with low latency and high transfer speeds.
//
// CloudFront fronts user provided content and caches it at edge locations across the world.
//
// Here's how you can use this construct:
//
// ```ts
// const sourceBucket = new s3.Bucket(this, 'Bucket');
//
// const distribution = new cloudfront.CloudFrontWebDistribution(this, 'MyDistribution', {
//    originConfigs: [
//      {
//        s3OriginSource: {
//        s3BucketSource: sourceBucket,
//        },
//        behaviors : [ {isDefaultBehavior: true}],
//      },
//    ],
// });
// ```
//
// This will create a CloudFront distribution that uses your S3Bucket as it's origin.
//
// You can customize the distribution using additional properties from the CloudFrontWebDistributionProps interface.
//
// Example:
//   var sourceBucket bucket
//
//   viewerCertificate := cloudfront.viewerCertificate.fromIamCertificate(jsii.String("MYIAMROLEIDENTIFIER"), &viewerCertificateOptions{
//   	aliases: []*string{
//   		jsii.String("MYALIAS"),
//   	},
//   })
//
//   cloudfront.NewCloudFrontWebDistribution(this, jsii.String("MyCfWebDistribution"), &cloudFrontWebDistributionProps{
//   	originConfigs: []sourceConfiguration{
//   		&sourceConfiguration{
//   			s3OriginSource: &s3OriginConfig{
//   				s3BucketSource: sourceBucket,
//   			},
//   			behaviors: []behavior{
//   				&behavior{
//   					isDefaultBehavior: jsii.Boolean(true),
//   				},
//   			},
//   		},
//   	},
//   	viewerCertificate: viewerCertificate,
//   })
//
type CloudFrontWebDistribution interface {
	awscdk.Resource
	IDistribution
	// The domain name created by CloudFront for this distribution.
	//
	// If you are using aliases for your distribution, this is the domainName your DNS records should point to.
	// (In Route53, you could create an ALIAS record to this value, for example.)
	DistributionDomainName() *string
	// The distribution ID for this distribution.
	DistributionId() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The logging bucket for this CloudFront distribution.
	//
	// If logging is not enabled for this distribution - this property will be undefined.
	LoggingBucket() awss3.IBucket
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Adds an IAM policy statement associated with this distribution to an IAM principal's policy.
	Grant(identity awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant to create invalidations for this bucket to an IAM principal (Role/Group/User).
	GrantCreateInvalidation(identity awsiam.IGrantable) awsiam.Grant
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CloudFrontWebDistribution
type jsiiProxy_CloudFrontWebDistribution struct {
	internal.Type__awscdkResource
	jsiiProxy_IDistribution
}

func (j *jsiiProxy_CloudFrontWebDistribution) DistributionDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontWebDistribution) DistributionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontWebDistribution) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontWebDistribution) LoggingBucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"loggingBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontWebDistribution) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontWebDistribution) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontWebDistribution) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewCloudFrontWebDistribution(scope constructs.Construct, id *string, props *CloudFrontWebDistributionProps) CloudFrontWebDistribution {
	_init_.Initialize()

	if err := validateNewCloudFrontWebDistributionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudFrontWebDistribution{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CloudFrontWebDistribution",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCloudFrontWebDistribution_Override(c CloudFrontWebDistribution, scope constructs.Construct, id *string, props *CloudFrontWebDistributionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.CloudFrontWebDistribution",
		[]interface{}{scope, id, props},
		c,
	)
}

// Creates a construct that represents an external (imported) distribution.
func CloudFrontWebDistribution_FromDistributionAttributes(scope constructs.Construct, id *string, attrs *CloudFrontWebDistributionAttributes) IDistribution {
	_init_.Initialize()

	if err := validateCloudFrontWebDistribution_FromDistributionAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IDistribution

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CloudFrontWebDistribution",
		"fromDistributionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CloudFrontWebDistribution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudFrontWebDistribution_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CloudFrontWebDistribution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func CloudFrontWebDistribution_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCloudFrontWebDistribution_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CloudFrontWebDistribution",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func CloudFrontWebDistribution_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCloudFrontWebDistribution_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.CloudFrontWebDistribution",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontWebDistribution) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := c.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_CloudFrontWebDistribution) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontWebDistribution) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := c.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontWebDistribution) GetResourceNameAttribute(nameAttr *string) *string {
	if err := c.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontWebDistribution) Grant(identity awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := c.validateGrantParameters(identity); err != nil {
		panic(err)
	}
	args := []interface{}{identity}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		c,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontWebDistribution) GrantCreateInvalidation(identity awsiam.IGrantable) awsiam.Grant {
	if err := c.validateGrantCreateInvalidationParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		c,
		"grantCreateInvalidation",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontWebDistribution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

