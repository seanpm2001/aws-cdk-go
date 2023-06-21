package awsses

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A configuration set event destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationSet configurationSet
//   var eventDestination eventDestination
//
//   configurationSetEventDestination := awscdk.Aws_ses.NewConfigurationSetEventDestination(this, jsii.String("MyConfigurationSetEventDestination"), &ConfigurationSetEventDestinationProps{
//   	ConfigurationSet: configurationSet,
//   	Destination: eventDestination,
//
//   	// the properties below are optional
//   	ConfigurationSetEventDestinationName: jsii.String("configurationSetEventDestinationName"),
//   	Enabled: jsii.Boolean(false),
//   	Events: []emailSendingEvent{
//   		awscdk.*Aws_ses.*emailSendingEvent_SEND,
//   	},
//   })
//
type ConfigurationSetEventDestination interface {
	awscdk.Resource
	IConfigurationSetEventDestination
	// The ID of the configuration set event destination.
	ConfigurationSetEventDestinationId() *string
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

// The jsii proxy struct for ConfigurationSetEventDestination
type jsiiProxy_ConfigurationSetEventDestination struct {
	internal.Type__awscdkResource
	jsiiProxy_IConfigurationSetEventDestination
}

func (j *jsiiProxy_ConfigurationSetEventDestination) ConfigurationSetEventDestinationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSetEventDestinationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigurationSetEventDestination) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigurationSetEventDestination) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigurationSetEventDestination) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigurationSetEventDestination) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewConfigurationSetEventDestination(scope constructs.Construct, id *string, props *ConfigurationSetEventDestinationProps) ConfigurationSetEventDestination {
	_init_.Initialize()

	if err := validateNewConfigurationSetEventDestinationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigurationSetEventDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.ConfigurationSetEventDestination",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewConfigurationSetEventDestination_Override(c ConfigurationSetEventDestination, scope constructs.Construct, id *string, props *ConfigurationSetEventDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.ConfigurationSetEventDestination",
		[]interface{}{scope, id, props},
		c,
	)
}

// Use an existing configuration set.
func ConfigurationSetEventDestination_FromConfigurationSetEventDestinationId(scope constructs.Construct, id *string, configurationSetEventDestinationId *string) IConfigurationSetEventDestination {
	_init_.Initialize()

	if err := validateConfigurationSetEventDestination_FromConfigurationSetEventDestinationIdParameters(scope, id, configurationSetEventDestinationId); err != nil {
		panic(err)
	}
	var returns IConfigurationSetEventDestination

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.ConfigurationSetEventDestination",
		"fromConfigurationSetEventDestinationId",
		[]interface{}{scope, id, configurationSetEventDestinationId},
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
func ConfigurationSetEventDestination_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateConfigurationSetEventDestination_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.ConfigurationSetEventDestination",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func ConfigurationSetEventDestination_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateConfigurationSetEventDestination_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.ConfigurationSetEventDestination",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ConfigurationSetEventDestination_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateConfigurationSetEventDestination_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.ConfigurationSetEventDestination",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigurationSetEventDestination) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := c.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_ConfigurationSetEventDestination) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigurationSetEventDestination) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (c *jsiiProxy_ConfigurationSetEventDestination) GetResourceNameAttribute(nameAttr *string) *string {
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

func (c *jsiiProxy_ConfigurationSetEventDestination) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
