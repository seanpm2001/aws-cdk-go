package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A Origin Request Policy configuration.
//
// Example:
//   // Using an existing origin request policy for a Distribution
//   var bucketOrigin s3Origin
//
//   cloudfront.NewDistribution(this, jsii.String("myDistManagedPolicy"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: bucketOrigin,
//   		originRequestPolicy: cloudfront.originRequestPolicy_CORS_S3_ORIGIN(),
//   	},
//   })
//
// Experimental.
type OriginRequestPolicy interface {
	awscdk.Resource
	IOriginRequestPolicy
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The ID of the origin request policy.
	// Experimental.
	OriginRequestPolicyId() *string
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
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

// The jsii proxy struct for OriginRequestPolicy
type jsiiProxy_OriginRequestPolicy struct {
	internal.Type__awscdkResource
	jsiiProxy_IOriginRequestPolicy
}

func (j *jsiiProxy_OriginRequestPolicy) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginRequestPolicy) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginRequestPolicy) OriginRequestPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originRequestPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginRequestPolicy) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginRequestPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewOriginRequestPolicy(scope constructs.Construct, id *string, props *OriginRequestPolicyProps) OriginRequestPolicy {
	_init_.Initialize()

	if err := validateNewOriginRequestPolicyParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_OriginRequestPolicy{}

	_jsii_.Create(
		"monocdk.aws_cloudfront.OriginRequestPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewOriginRequestPolicy_Override(o OriginRequestPolicy, scope constructs.Construct, id *string, props *OriginRequestPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudfront.OriginRequestPolicy",
		[]interface{}{scope, id, props},
		o,
	)
}

// Imports a Origin Request Policy from its id.
// Experimental.
func OriginRequestPolicy_FromOriginRequestPolicyId(scope constructs.Construct, id *string, originRequestPolicyId *string) IOriginRequestPolicy {
	_init_.Initialize()

	if err := validateOriginRequestPolicy_FromOriginRequestPolicyIdParameters(scope, id, originRequestPolicyId); err != nil {
		panic(err)
	}
	var returns IOriginRequestPolicy

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudfront.OriginRequestPolicy",
		"fromOriginRequestPolicyId",
		[]interface{}{scope, id, originRequestPolicyId},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func OriginRequestPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOriginRequestPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudfront.OriginRequestPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func OriginRequestPolicy_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateOriginRequestPolicy_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudfront.OriginRequestPolicy",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func OriginRequestPolicy_ALL_VIEWER() IOriginRequestPolicy {
	_init_.Initialize()
	var returns IOriginRequestPolicy
	_jsii_.StaticGet(
		"monocdk.aws_cloudfront.OriginRequestPolicy",
		"ALL_VIEWER",
		&returns,
	)
	return returns
}

func OriginRequestPolicy_CORS_CUSTOM_ORIGIN() IOriginRequestPolicy {
	_init_.Initialize()
	var returns IOriginRequestPolicy
	_jsii_.StaticGet(
		"monocdk.aws_cloudfront.OriginRequestPolicy",
		"CORS_CUSTOM_ORIGIN",
		&returns,
	)
	return returns
}

func OriginRequestPolicy_CORS_S3_ORIGIN() IOriginRequestPolicy {
	_init_.Initialize()
	var returns IOriginRequestPolicy
	_jsii_.StaticGet(
		"monocdk.aws_cloudfront.OriginRequestPolicy",
		"CORS_S3_ORIGIN",
		&returns,
	)
	return returns
}

func OriginRequestPolicy_ELEMENTAL_MEDIA_TAILOR() IOriginRequestPolicy {
	_init_.Initialize()
	var returns IOriginRequestPolicy
	_jsii_.StaticGet(
		"monocdk.aws_cloudfront.OriginRequestPolicy",
		"ELEMENTAL_MEDIA_TAILOR",
		&returns,
	)
	return returns
}

func OriginRequestPolicy_USER_AGENT_REFERER_HEADERS() IOriginRequestPolicy {
	_init_.Initialize()
	var returns IOriginRequestPolicy
	_jsii_.StaticGet(
		"monocdk.aws_cloudfront.OriginRequestPolicy",
		"USER_AGENT_REFERER_HEADERS",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OriginRequestPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := o.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (o *jsiiProxy_OriginRequestPolicy) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginRequestPolicy) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := o.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginRequestPolicy) GetResourceNameAttribute(nameAttr *string) *string {
	if err := o.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginRequestPolicy) OnPrepare() {
	_jsii_.InvokeVoid(
		o,
		"onPrepare",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OriginRequestPolicy) OnSynthesize(session constructs.ISynthesisSession) {
	if err := o.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (o *jsiiProxy_OriginRequestPolicy) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginRequestPolicy) Prepare() {
	_jsii_.InvokeVoid(
		o,
		"prepare",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OriginRequestPolicy) Synthesize(session awscdk.ISynthesisSession) {
	if err := o.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"synthesize",
		[]interface{}{session},
	)
}

func (o *jsiiProxy_OriginRequestPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OriginRequestPolicy) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}
