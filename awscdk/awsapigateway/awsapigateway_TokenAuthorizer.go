package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/constructs-go/constructs/v3"
)

// Token based lambda authorizer that recognizes the caller's identity as a bearer token, such as a JSON Web Token (JWT) or an OAuth token.
//
// Based on the token, authorization is performed by a lambda function.
//
// Example:
//   var authFn function
//   var books resource
//
//
//   auth := apigateway.NewTokenAuthorizer(this, jsii.String("booksAuthorizer"), &tokenAuthorizerProps{
//   	handler: authFn,
//   })
//
//   books.addMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &methodOptions{
//   	authorizer: auth,
//   })
//
// Experimental.
type TokenAuthorizer interface {
	Authorizer
	IAuthorizer
	// The authorization type of this authorizer.
	// Experimental.
	AuthorizationType() AuthorizationType
	// The ARN of the authorizer to be used in permission policies, such as IAM and resource-based grants.
	// Experimental.
	AuthorizerArn() *string
	// The id of the authorizer.
	// Experimental.
	AuthorizerId() *string
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
	// The Lambda function handler that this authorizer uses.
	// Experimental.
	Handler() awslambda.IFunction
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// Experimental.
	RestApiId() *string
	// Experimental.
	SetRestApiId(val *string)
	// The IAM role that the API Gateway service assumes while invoking the Lambda function.
	// Experimental.
	Role() awsiam.IRole
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
	// Returns a token that resolves to the Rest Api Id at the time of synthesis.
	//
	// Throws an error, during token resolution, if no RestApi is attached to this authorizer.
	// Experimental.
	LazyRestApiId() *string
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
	// Sets up the permissions necessary for the API Gateway service to invoke the Lambda function.
	// Experimental.
	SetupPermissions()
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

// The jsii proxy struct for TokenAuthorizer
type jsiiProxy_TokenAuthorizer struct {
	jsiiProxy_Authorizer
	jsiiProxy_IAuthorizer
}

func (j *jsiiProxy_TokenAuthorizer) AuthorizationType() AuthorizationType {
	var returns AuthorizationType
	_jsii_.Get(
		j,
		"authorizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthorizer) AuthorizerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthorizer) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthorizer) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthorizer) Handler() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthorizer) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthorizer) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthorizer) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthorizer) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthorizer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewTokenAuthorizer(scope constructs.Construct, id *string, props *TokenAuthorizerProps) TokenAuthorizer {
	_init_.Initialize()

	if err := validateNewTokenAuthorizerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_TokenAuthorizer{}

	_jsii_.Create(
		"monocdk.aws_apigateway.TokenAuthorizer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewTokenAuthorizer_Override(t TokenAuthorizer, scope constructs.Construct, id *string, props *TokenAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigateway.TokenAuthorizer",
		[]interface{}{scope, id, props},
		t,
	)
}

func (j *jsiiProxy_TokenAuthorizer)SetRestApiId(val *string) {
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

// Return whether the given object is an Authorizer.
// Experimental.
func TokenAuthorizer_IsAuthorizer(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTokenAuthorizer_IsAuthorizerParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apigateway.TokenAuthorizer",
		"isAuthorizer",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func TokenAuthorizer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTokenAuthorizer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apigateway.TokenAuthorizer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func TokenAuthorizer_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateTokenAuthorizer_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apigateway.TokenAuthorizer",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthorizer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := t.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (t *jsiiProxy_TokenAuthorizer) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthorizer) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := t.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthorizer) GetResourceNameAttribute(nameAttr *string) *string {
	if err := t.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthorizer) LazyRestApiId() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"lazyRestApiId",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthorizer) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthorizer) OnSynthesize(session constructs.ISynthesisSession) {
	if err := t.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (t *jsiiProxy_TokenAuthorizer) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthorizer) Prepare() {
	_jsii_.InvokeVoid(
		t,
		"prepare",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthorizer) SetupPermissions() {
	_jsii_.InvokeVoid(
		t,
		"setupPermissions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthorizer) Synthesize(session awscdk.ISynthesisSession) {
	if err := t.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		[]interface{}{session},
	)
}

func (t *jsiiProxy_TokenAuthorizer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthorizer) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}
