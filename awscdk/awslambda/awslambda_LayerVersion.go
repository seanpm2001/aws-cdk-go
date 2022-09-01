package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Defines a new Lambda Layer version.
//
// Example:
//   lambda.NewLayerVersion(this, jsii.String("MyLayer"), &layerVersionProps{
//   	removalPolicy: awscdk.RemovalPolicy_RETAIN,
//   	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	compatibleArchitectures: []architecture{
//   		lambda.*architecture_X86_64(),
//   		lambda.*architecture_ARM_64(),
//   	},
//   })
//
type LayerVersion interface {
	awscdk.Resource
	ILayerVersion
	// The runtimes compatible with this Layer.
	CompatibleRuntimes() *[]Runtime
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The ARN of the Lambda Layer version that this Layer defines.
	LayerVersionArn() *string
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
	// Add permission for this layer version to specific entities.
	//
	// Usage within
	// the same account where the layer is defined is always allowed and does not
	// require calling this method. Note that the principal that creates the
	// Lambda function using the layer (for example, a CloudFormation changeset
	// execution role) also needs to have the ``lambda:GetLayerVersion``
	// permission on the layer version.
	AddPermission(id *string, permission *LayerVersionPermission)
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

// The jsii proxy struct for LayerVersion
type jsiiProxy_LayerVersion struct {
	internal.Type__awscdkResource
	jsiiProxy_ILayerVersion
}

func (j *jsiiProxy_LayerVersion) CompatibleRuntimes() *[]Runtime {
	var returns *[]Runtime
	_jsii_.Get(
		j,
		"compatibleRuntimes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LayerVersion) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LayerVersion) LayerVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LayerVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LayerVersion) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LayerVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewLayerVersion(scope constructs.Construct, id *string, props *LayerVersionProps) LayerVersion {
	_init_.Initialize()

	if err := validateNewLayerVersionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LayerVersion{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.LayerVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLayerVersion_Override(l LayerVersion, scope constructs.Construct, id *string, props *LayerVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.LayerVersion",
		[]interface{}{scope, id, props},
		l,
	)
}

// Imports a layer version by ARN.
//
// Assumes it is compatible with all Lambda runtimes.
func LayerVersion_FromLayerVersionArn(scope constructs.Construct, id *string, layerVersionArn *string) ILayerVersion {
	_init_.Initialize()

	if err := validateLayerVersion_FromLayerVersionArnParameters(scope, id, layerVersionArn); err != nil {
		panic(err)
	}
	var returns ILayerVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.LayerVersion",
		"fromLayerVersionArn",
		[]interface{}{scope, id, layerVersionArn},
		&returns,
	)

	return returns
}

// Imports a Layer that has been defined externally.
func LayerVersion_FromLayerVersionAttributes(scope constructs.Construct, id *string, attrs *LayerVersionAttributes) ILayerVersion {
	_init_.Initialize()

	if err := validateLayerVersion_FromLayerVersionAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns ILayerVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.LayerVersion",
		"fromLayerVersionAttributes",
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
func LayerVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLayerVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.LayerVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func LayerVersion_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateLayerVersion_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.LayerVersion",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func LayerVersion_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateLayerVersion_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.LayerVersion",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LayerVersion) AddPermission(id *string, permission *LayerVersionPermission) {
	if err := l.validateAddPermissionParameters(id, permission); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addPermission",
		[]interface{}{id, permission},
	)
}

func (l *jsiiProxy_LayerVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := l.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (l *jsiiProxy_LayerVersion) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LayerVersion) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := l.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LayerVersion) GetResourceNameAttribute(nameAttr *string) *string {
	if err := l.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LayerVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

