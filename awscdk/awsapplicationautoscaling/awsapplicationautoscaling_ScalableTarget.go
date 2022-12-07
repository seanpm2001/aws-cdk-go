package awsapplicationautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Define a scalable target.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var code code
//
//
//   handler := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
//   	runtime: lambda.runtime_PYTHON_3_7(),
//   	handler: jsii.String("index.handler"),
//   	code: code,
//
//   	reservedConcurrentExecutions: jsii.Number(2),
//   })
//
//   fnVer := handler.currentVersion
//
//   target := appscaling.NewScalableTarget(this, jsii.String("ScalableTarget"), &scalableTargetProps{
//   	serviceNamespace: appscaling.serviceNamespace_LAMBDA,
//   	maxCapacity: jsii.Number(100),
//   	minCapacity: jsii.Number(10),
//   	resourceId: fmt.Sprintf("function:%v:%v", handler.functionName, fnVer.version),
//   	scalableDimension: jsii.String("lambda:function:ProvisionedConcurrency"),
//   })
//
//   target.scaleToTrackMetric(jsii.String("PceTracking"), &basicTargetTrackingScalingPolicyProps{
//   	targetValue: jsii.Number(0.9),
//   	predefinedMetric: appscaling.predefinedMetric_LAMBDA_PROVISIONED_CONCURRENCY_UTILIZATION,
//   })
//
type ScalableTarget interface {
	awscdk.Resource
	IScalableTarget
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
	//    cross-environment scenarios.
	PhysicalName() *string
	// The role used to give AutoScaling permissions to your resource.
	Role() awsiam.IRole
	// ID of the Scalable Target.
	//
	// Example value: `service/ecsStack-MyECSCluster-AB12CDE3F4GH/ecsStack-MyECSService-AB12CDE3F4GH|ecs:service:DesiredCount|ecs`.
	ScalableTargetId() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Add a policy statement to the role's policy.
	AddToRolePolicy(statement awsiam.PolicyStatement)
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
	// Scale out or in, in response to a metric.
	ScaleOnMetric(id *string, props *BasicStepScalingPolicyProps) StepScalingPolicy
	// Scale out or in based on time.
	ScaleOnSchedule(id *string, action *ScalingSchedule)
	// Scale out or in in order to keep a metric around a target value.
	ScaleToTrackMetric(id *string, props *BasicTargetTrackingScalingPolicyProps) TargetTrackingScalingPolicy
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ScalableTarget
type jsiiProxy_ScalableTarget struct {
	internal.Type__awscdkResource
	jsiiProxy_IScalableTarget
}

func (j *jsiiProxy_ScalableTarget) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScalableTarget) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScalableTarget) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScalableTarget) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScalableTarget) ScalableTargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalableTargetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScalableTarget) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewScalableTarget(scope constructs.Construct, id *string, props *ScalableTargetProps) ScalableTarget {
	_init_.Initialize()

	if err := validateNewScalableTargetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ScalableTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_applicationautoscaling.ScalableTarget",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewScalableTarget_Override(s ScalableTarget, scope constructs.Construct, id *string, props *ScalableTargetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_applicationautoscaling.ScalableTarget",
		[]interface{}{scope, id, props},
		s,
	)
}

func ScalableTarget_FromScalableTargetId(scope constructs.Construct, id *string, scalableTargetId *string) IScalableTarget {
	_init_.Initialize()

	if err := validateScalableTarget_FromScalableTargetIdParameters(scope, id, scalableTargetId); err != nil {
		panic(err)
	}
	var returns IScalableTarget

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationautoscaling.ScalableTarget",
		"fromScalableTargetId",
		[]interface{}{scope, id, scalableTargetId},
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
func ScalableTarget_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateScalableTarget_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationautoscaling.ScalableTarget",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func ScalableTarget_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateScalableTarget_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationautoscaling.ScalableTarget",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ScalableTarget_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateScalableTarget_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationautoscaling.ScalableTarget",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalableTarget) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := s.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (s *jsiiProxy_ScalableTarget) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := s.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_ScalableTarget) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalableTarget) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := s.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalableTarget) GetResourceNameAttribute(nameAttr *string) *string {
	if err := s.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalableTarget) ScaleOnMetric(id *string, props *BasicStepScalingPolicyProps) StepScalingPolicy {
	if err := s.validateScaleOnMetricParameters(id, props); err != nil {
		panic(err)
	}
	var returns StepScalingPolicy

	_jsii_.Invoke(
		s,
		"scaleOnMetric",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalableTarget) ScaleOnSchedule(id *string, action *ScalingSchedule) {
	if err := s.validateScaleOnScheduleParameters(id, action); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"scaleOnSchedule",
		[]interface{}{id, action},
	)
}

func (s *jsiiProxy_ScalableTarget) ScaleToTrackMetric(id *string, props *BasicTargetTrackingScalingPolicyProps) TargetTrackingScalingPolicy {
	if err := s.validateScaleToTrackMetricParameters(id, props); err != nil {
		panic(err)
	}
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		s,
		"scaleToTrackMetric",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalableTarget) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

