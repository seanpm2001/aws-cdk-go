package awsautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Define a warm pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var autoScalingGroup autoScalingGroup
//
//   warmPool := awscdk.Aws_autoscaling.NewWarmPool(this, jsii.String("MyWarmPool"), &warmPoolProps{
//   	autoScalingGroup: autoScalingGroup,
//
//   	// the properties below are optional
//   	maxGroupPreparedCapacity: jsii.Number(123),
//   	minSize: jsii.Number(123),
//   	poolState: awscdk.*Aws_autoscaling.poolState_HIBERNATED,
//   	reuseOnScaleIn: jsii.Boolean(false),
//   })
//
// Experimental.
type WarmPool interface {
	awscdk.Resource
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

// The jsii proxy struct for WarmPool
type jsiiProxy_WarmPool struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_WarmPool) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarmPool) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarmPool) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WarmPool) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewWarmPool(scope constructs.Construct, id *string, props *WarmPoolProps) WarmPool {
	_init_.Initialize()

	if err := validateNewWarmPoolParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_WarmPool{}

	_jsii_.Create(
		"monocdk.aws_autoscaling.WarmPool",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewWarmPool_Override(w WarmPool, scope constructs.Construct, id *string, props *WarmPoolProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_autoscaling.WarmPool",
		[]interface{}{scope, id, props},
		w,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func WarmPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWarmPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_autoscaling.WarmPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func WarmPool_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateWarmPool_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_autoscaling.WarmPool",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarmPool) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := w.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (w *jsiiProxy_WarmPool) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarmPool) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := w.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarmPool) GetResourceNameAttribute(nameAttr *string) *string {
	if err := w.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarmPool) OnPrepare() {
	_jsii_.InvokeVoid(
		w,
		"onPrepare",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WarmPool) OnSynthesize(session constructs.ISynthesisSession) {
	if err := w.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (w *jsiiProxy_WarmPool) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarmPool) Prepare() {
	_jsii_.InvokeVoid(
		w,
		"prepare",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WarmPool) Synthesize(session awscdk.ISynthesisSession) {
	if err := w.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"synthesize",
		[]interface{}{session},
	)
}

func (w *jsiiProxy_WarmPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WarmPool) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}
