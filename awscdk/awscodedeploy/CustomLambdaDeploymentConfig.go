package awscodedeploy

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A custom Deployment Configuration for a Lambda Deployment Group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customLambdaDeploymentConfig := awscdk.Aws_codedeploy.NewCustomLambdaDeploymentConfig(this, jsii.String("MyCustomLambdaDeploymentConfig"), &CustomLambdaDeploymentConfigProps{
//   	Interval: cdk.Duration_Minutes(jsii.Number(30)),
//   	Percentage: jsii.Number(123),
//   	Type: awscdk.*Aws_codedeploy.CustomLambdaDeploymentConfigType_CANARY,
//
//   	// the properties below are optional
//   	DeploymentConfigName: jsii.String("deploymentConfigName"),
//   })
//
// Deprecated: CloudFormation now supports Lambda deployment configurations without custom resources. Use `LambdaDeploymentConfig`.
type CustomLambdaDeploymentConfig interface {
	awscdk.Resource
	ILambdaDeploymentConfig
	// The arn of the deployment config.
	// Deprecated: Use `LambdaDeploymentConfig`.
	DeploymentConfigArn() *string
	// The name of the deployment config.
	// Deprecated: Use `LambdaDeploymentConfig`.
	DeploymentConfigName() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Deprecated: CloudFormation now supports Lambda deployment configurations without custom resources. Use `LambdaDeploymentConfig`.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	// Deprecated: CloudFormation now supports Lambda deployment configurations without custom resources. Use `LambdaDeploymentConfig`.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	// Deprecated: CloudFormation now supports Lambda deployment configurations without custom resources. Use `LambdaDeploymentConfig`.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Deprecated: CloudFormation now supports Lambda deployment configurations without custom resources. Use `LambdaDeploymentConfig`.
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
	// Deprecated: CloudFormation now supports Lambda deployment configurations without custom resources. Use `LambdaDeploymentConfig`.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Deprecated: CloudFormation now supports Lambda deployment configurations without custom resources. Use `LambdaDeploymentConfig`.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Deprecated: CloudFormation now supports Lambda deployment configurations without custom resources. Use `LambdaDeploymentConfig`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Deprecated: CloudFormation now supports Lambda deployment configurations without custom resources. Use `LambdaDeploymentConfig`.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	// Deprecated: CloudFormation now supports Lambda deployment configurations without custom resources. Use `LambdaDeploymentConfig`.
	ToString() *string
}

// The jsii proxy struct for CustomLambdaDeploymentConfig
type jsiiProxy_CustomLambdaDeploymentConfig struct {
	internal.Type__awscdkResource
	jsiiProxy_ILambdaDeploymentConfig
}

func (j *jsiiProxy_CustomLambdaDeploymentConfig) DeploymentConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomLambdaDeploymentConfig) DeploymentConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomLambdaDeploymentConfig) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomLambdaDeploymentConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomLambdaDeploymentConfig) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomLambdaDeploymentConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Deprecated: CloudFormation now supports Lambda deployment configurations without custom resources. Use `LambdaDeploymentConfig`.
func NewCustomLambdaDeploymentConfig(scope constructs.Construct, id *string, props *CustomLambdaDeploymentConfigProps) CustomLambdaDeploymentConfig {
	_init_.Initialize()

	if err := validateNewCustomLambdaDeploymentConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomLambdaDeploymentConfig{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.CustomLambdaDeploymentConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Deprecated: CloudFormation now supports Lambda deployment configurations without custom resources. Use `LambdaDeploymentConfig`.
func NewCustomLambdaDeploymentConfig_Override(c CustomLambdaDeploymentConfig, scope constructs.Construct, id *string, props *CustomLambdaDeploymentConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.CustomLambdaDeploymentConfig",
		[]interface{}{scope, id, props},
		c,
	)
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
// Deprecated: CloudFormation now supports Lambda deployment configurations without custom resources. Use `LambdaDeploymentConfig`.
func CustomLambdaDeploymentConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCustomLambdaDeploymentConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.CustomLambdaDeploymentConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Deprecated: CloudFormation now supports Lambda deployment configurations without custom resources. Use `LambdaDeploymentConfig`.
func CustomLambdaDeploymentConfig_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCustomLambdaDeploymentConfig_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.CustomLambdaDeploymentConfig",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Deprecated: CloudFormation now supports Lambda deployment configurations without custom resources. Use `LambdaDeploymentConfig`.
func CustomLambdaDeploymentConfig_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCustomLambdaDeploymentConfig_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.CustomLambdaDeploymentConfig",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomLambdaDeploymentConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := c.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_CustomLambdaDeploymentConfig) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomLambdaDeploymentConfig) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (c *jsiiProxy_CustomLambdaDeploymentConfig) GetResourceNameAttribute(nameAttr *string) *string {
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

func (c *jsiiProxy_CustomLambdaDeploymentConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
