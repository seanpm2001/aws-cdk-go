package awscognito

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Define a UserPool App Client.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//   provider := cognito.NewUserPoolIdentityProviderAmazon(this, jsii.String("Amazon"), &UserPoolIdentityProviderAmazonProps{
//   	UserPool: pool,
//   	ClientId: jsii.String("amzn-client-id"),
//   	ClientSecret: jsii.String("amzn-client-secret"),
//   })
//
//   client := pool.addClient(jsii.String("app-client"), &UserPoolClientOptions{
//   	// ...
//   	SupportedIdentityProviders: []userPoolClientIdentityProvider{
//   		cognito.*userPoolClientIdentityProvider_AMAZON(),
//   	},
//   })
//
//   client.Node.AddDependency(provider)
//
type UserPoolClient interface {
	awscdk.Resource
	IUserPoolClient
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// The OAuth flows enabled for this client.
	OAuthFlows() *OAuthFlows
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Name of the application client.
	UserPoolClientId() *string
	// The client name that was specified via the `userPoolClientName` property during initialization, throws an error otherwise.
	UserPoolClientName() *string
	// The generated client secret.
	//
	// Only available if the "generateSecret" props is set to true.
	UserPoolClientSecret() awscdk.SecretValue
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
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for UserPoolClient
type jsiiProxy_UserPoolClient struct {
	internal.Type__awscdkResource
	jsiiProxy_IUserPoolClient
}

func (j *jsiiProxy_UserPoolClient) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolClient) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolClient) OAuthFlows() *OAuthFlows {
	var returns *OAuthFlows
	_jsii_.Get(
		j,
		"oAuthFlows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolClient) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolClient) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolClient) UserPoolClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolClient) UserPoolClientName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolClientName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserPoolClient) UserPoolClientSecret() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"userPoolClientSecret",
		&returns,
	)
	return returns
}


func NewUserPoolClient(scope constructs.Construct, id *string, props *UserPoolClientProps) UserPoolClient {
	_init_.Initialize()

	if err := validateNewUserPoolClientParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_UserPoolClient{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.UserPoolClient",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewUserPoolClient_Override(u UserPoolClient, scope constructs.Construct, id *string, props *UserPoolClientProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cognito.UserPoolClient",
		[]interface{}{scope, id, props},
		u,
	)
}

// Import a user pool client given its id.
func UserPoolClient_FromUserPoolClientId(scope constructs.Construct, id *string, userPoolClientId *string) IUserPoolClient {
	_init_.Initialize()

	if err := validateUserPoolClient_FromUserPoolClientIdParameters(scope, id, userPoolClientId); err != nil {
		panic(err)
	}
	var returns IUserPoolClient

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.UserPoolClient",
		"fromUserPoolClientId",
		[]interface{}{scope, id, userPoolClientId},
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
func UserPoolClient_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUserPoolClient_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.UserPoolClient",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func UserPoolClient_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateUserPoolClient_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.UserPoolClient",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func UserPoolClient_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateUserPoolClient_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cognito.UserPoolClient",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolClient) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := u.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (u *jsiiProxy_UserPoolClient) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolClient) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := u.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolClient) GetResourceNameAttribute(nameAttr *string) *string {
	if err := u.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserPoolClient) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

