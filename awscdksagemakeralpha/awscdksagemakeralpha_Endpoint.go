// The CDK Construct Library for AWS::SageMaker
package awscdksagemakeralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdksagemakeralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdksagemakeralpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Defines a SageMaker endpoint.
//
// Example:
//   import sagemaker "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//
//   var endpointConfig endpointConfig
//
//
//   endpoint := sagemaker.NewEndpoint(this, jsii.String("Endpoint"), &endpointProps{
//   	endpointConfig: endpointConfig,
//   })
//   productionVariant := endpoint.findInstanceProductionVariant(jsii.String("my-variant"))
//   productionVariant.metricModelLatency().createAlarm(this, jsii.String("ModelLatencyAlarm"), &createAlarmOptions{
//   	threshold: jsii.Number(100000),
//   	evaluationPeriods: jsii.Number(3),
//   })
//
// Experimental.
type Endpoint interface {
	awscdk.Resource
	IEndpoint
	// The ARN of the endpoint.
	// Experimental.
	EndpointArn() *string
	// The name of the endpoint.
	// Experimental.
	EndpointName() *string
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
	// Get instance production variants associated with endpoint.
	// Experimental.
	InstanceProductionVariants() *[]IEndpointInstanceProductionVariant
	// The tree node.
	// Experimental.
	Node() constructs.Node
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
	// Find instance production variant based on variant name.
	// Experimental.
	FindInstanceProductionVariant(name *string) IEndpointInstanceProductionVariant
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
	// Permits an IAM principal to invoke this endpoint.
	// Experimental.
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Endpoint
type jsiiProxy_Endpoint struct {
	internal.Type__awscdkResource
	jsiiProxy_IEndpoint
}

func (j *jsiiProxy_Endpoint) EndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Endpoint) EndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Endpoint) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Endpoint) InstanceProductionVariants() *[]IEndpointInstanceProductionVariant {
	var returns *[]IEndpointInstanceProductionVariant
	_jsii_.Get(
		j,
		"instanceProductionVariants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Endpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Endpoint) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Endpoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewEndpoint(scope constructs.Construct, id *string, props *EndpointProps) Endpoint {
	_init_.Initialize()

	if err := validateNewEndpointParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Endpoint{}

	_jsii_.Create(
		"@aws-cdk/aws-sagemaker-alpha.Endpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEndpoint_Override(e Endpoint, scope constructs.Construct, id *string, props *EndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-sagemaker-alpha.Endpoint",
		[]interface{}{scope, id, props},
		e,
	)
}

// Imports an Endpoint defined either outside the CDK or in a different CDK stack.
// Experimental.
func Endpoint_FromEndpointArn(scope constructs.Construct, id *string, endpointArn *string) IEndpoint {
	_init_.Initialize()

	if err := validateEndpoint_FromEndpointArnParameters(scope, id, endpointArn); err != nil {
		panic(err)
	}
	var returns IEndpoint

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-sagemaker-alpha.Endpoint",
		"fromEndpointArn",
		[]interface{}{scope, id, endpointArn},
		&returns,
	)

	return returns
}

// Imports an Endpoint defined either outside the CDK or in a different CDK stack.
// Experimental.
func Endpoint_FromEndpointAttributes(scope constructs.Construct, id *string, attrs *EndpointAttributes) IEndpoint {
	_init_.Initialize()

	if err := validateEndpoint_FromEndpointAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IEndpoint

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-sagemaker-alpha.Endpoint",
		"fromEndpointAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports an Endpoint defined either outside the CDK or in a different CDK stack.
// Experimental.
func Endpoint_FromEndpointName(scope constructs.Construct, id *string, endpointName *string) IEndpoint {
	_init_.Initialize()

	if err := validateEndpoint_FromEndpointNameParameters(scope, id, endpointName); err != nil {
		panic(err)
	}
	var returns IEndpoint

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-sagemaker-alpha.Endpoint",
		"fromEndpointName",
		[]interface{}{scope, id, endpointName},
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
// Experimental.
func Endpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-sagemaker-alpha.Endpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func Endpoint_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateEndpoint_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-sagemaker-alpha.Endpoint",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Endpoint_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateEndpoint_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-sagemaker-alpha.Endpoint",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Endpoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := e.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (e *jsiiProxy_Endpoint) FindInstanceProductionVariant(name *string) IEndpointInstanceProductionVariant {
	if err := e.validateFindInstanceProductionVariantParameters(name); err != nil {
		panic(err)
	}
	var returns IEndpointInstanceProductionVariant

	_jsii_.Invoke(
		e,
		"findInstanceProductionVariant",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Endpoint) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Endpoint) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := e.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Endpoint) GetResourceNameAttribute(nameAttr *string) *string {
	if err := e.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Endpoint) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := e.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		e,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Endpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

