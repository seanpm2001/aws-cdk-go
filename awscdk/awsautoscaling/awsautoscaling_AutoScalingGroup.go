package awsautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsautoscaling/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancing"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/constructs-go/constructs/v3"
)

// A Fleet represents a managed set of EC2 instances.
//
// The Fleet models a number of AutoScalingGroups, a launch configuration, a
// security group and an instance role.
//
// It allows adding arbitrary commands to the startup scripts of the instances
// in the fleet.
//
// The ASG spans the availability zones specified by vpcSubnets, falling back to
// the Vpc default strategy if not specified.
//
// Example:
//   var vpc vpc
//
//
//   mySecurityGroup := ec2.NewSecurityGroup(this, jsii.String("SecurityGroup"), &securityGroupProps{
//   	vpc: vpc,
//   })
//   autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
//   	vpc: vpc,
//   	instanceType: ec2.instanceType.of(ec2.instanceClass_BURSTABLE2, ec2.instanceSize_MICRO),
//   	machineImage: ec2.NewAmazonLinuxImage(),
//   	securityGroup: mySecurityGroup,
//   })
//
// Experimental.
type AutoScalingGroup interface {
	awscdk.Resource
	IAutoScalingGroup
	awsec2.IConnectable
	awselasticloadbalancing.ILoadBalancerTarget
	awselasticloadbalancingv2.IApplicationLoadBalancerTarget
	awselasticloadbalancingv2.INetworkLoadBalancerTarget
	// Experimental.
	AlbTargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup
	// Experimental.
	SetAlbTargetGroup(val awselasticloadbalancingv2.ApplicationTargetGroup)
	// Arn of the AutoScalingGroup.
	// Experimental.
	AutoScalingGroupArn() *string
	// Name of the AutoScalingGroup.
	// Experimental.
	AutoScalingGroupName() *string
	// The network connections associated with this resource.
	// Experimental.
	Connections() awsec2.Connections
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
	// The principal to grant permissions to.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// The maximum amount of time that an instance can be in service.
	// Experimental.
	MaxInstanceLifetime() awscdk.Duration
	// Experimental.
	NewInstancesProtectedFromScaleIn() *bool
	// Experimental.
	SetNewInstancesProtectedFromScaleIn(val *bool)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The type of OS instances of this fleet are running.
	// Experimental.
	OsType() awsec2.OperatingSystemType
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The IAM Role in the instance profile.
	// Experimental.
	Role() awsiam.IRole
	// The maximum spot price configured for the autoscaling group.
	//
	// `undefined`
	// indicates that this group uses on-demand capacity.
	// Experimental.
	SpotPrice() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// The Base64-encoded user data to make available to the launched EC2 instances.
	// Experimental.
	UserData() awsec2.UserData
	// Send a message to either an SQS queue or SNS topic when instances launch or terminate.
	// Experimental.
	AddLifecycleHook(id *string, props *BasicLifecycleHookProps) LifecycleHook
	// Add the security group to all instances via the launch configuration security groups array.
	// Experimental.
	AddSecurityGroup(securityGroup awsec2.ISecurityGroup)
	// Adds a statement to the IAM role assumed by instances of this fleet.
	// Experimental.
	AddToRolePolicy(statement awsiam.PolicyStatement)
	// Add command to the startup script of fleet instances.
	//
	// The command must be in the scripting language supported by the fleet's OS (i.e. Linux/Windows).
	// Does nothing for imported ASGs.
	// Experimental.
	AddUserData(commands ...*string)
	// Add a pool of pre-initialized EC2 instances that sits alongside an Auto Scaling group.
	// Experimental.
	AddWarmPool(options *WarmPoolOptions) WarmPool
	// Use a CloudFormation Init configuration at instance startup.
	//
	// This does the following:
	//
	// - Attaches the CloudFormation Init metadata to the AutoScalingGroup resource.
	// - Add commands to the UserData to run `cfn-init` and `cfn-signal`.
	// - Update the instance's CreationPolicy to wait for `cfn-init` to finish
	//    before reporting success.
	// Experimental.
	ApplyCloudFormationInit(init awsec2.CloudFormationInit, options *ApplyCloudFormationInitOptions)
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
	// Returns `true` if newly-launched instances are protected from scale-in.
	// Experimental.
	AreNewInstancesProtectedFromScaleIn() *bool
	// Attach to ELBv2 Application Target Group.
	// Experimental.
	AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	// Attach to a classic load balancer.
	// Experimental.
	AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer)
	// Attach to ELBv2 Application Target Group.
	// Experimental.
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
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
	// Ensures newly-launched instances are protected from scale-in.
	// Experimental.
	ProtectNewInstancesFromScaleIn()
	// Scale out or in to achieve a target CPU utilization.
	// Experimental.
	ScaleOnCpuUtilization(id *string, props *CpuUtilizationScalingProps) TargetTrackingScalingPolicy
	// Scale out or in to achieve a target network ingress rate.
	// Experimental.
	ScaleOnIncomingBytes(id *string, props *NetworkUtilizationScalingProps) TargetTrackingScalingPolicy
	// Scale out or in, in response to a metric.
	// Experimental.
	ScaleOnMetric(id *string, props *BasicStepScalingPolicyProps) StepScalingPolicy
	// Scale out or in to achieve a target network egress rate.
	// Experimental.
	ScaleOnOutgoingBytes(id *string, props *NetworkUtilizationScalingProps) TargetTrackingScalingPolicy
	// Scale out or in to achieve a target request handling rate.
	//
	// The AutoScalingGroup must have been attached to an Application Load Balancer
	// in order to be able to call this.
	// Experimental.
	ScaleOnRequestCount(id *string, props *RequestCountScalingProps) TargetTrackingScalingPolicy
	// Scale out or in based on time.
	// Experimental.
	ScaleOnSchedule(id *string, props *BasicScheduledActionProps) ScheduledAction
	// Scale out or in in order to keep a metric around a target value.
	// Experimental.
	ScaleToTrackMetric(id *string, props *MetricTargetTrackingProps) TargetTrackingScalingPolicy
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

// The jsii proxy struct for AutoScalingGroup
type jsiiProxy_AutoScalingGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_IAutoScalingGroup
	internal.Type__awsec2IConnectable
	internal.Type__awselasticloadbalancingILoadBalancerTarget
	internal.Type__awselasticloadbalancingv2IApplicationLoadBalancerTarget
	internal.Type__awselasticloadbalancingv2INetworkLoadBalancerTarget
}

func (j *jsiiProxy_AutoScalingGroup) AlbTargetGroup() awselasticloadbalancingv2.ApplicationTargetGroup {
	var returns awselasticloadbalancingv2.ApplicationTargetGroup
	_jsii_.Get(
		j,
		"albTargetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) AutoScalingGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) AutoScalingGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) MaxInstanceLifetime() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"maxInstanceLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) NewInstancesProtectedFromScaleIn() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"newInstancesProtectedFromScaleIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) OsType() awsec2.OperatingSystemType {
	var returns awsec2.OperatingSystemType
	_jsii_.Get(
		j,
		"osType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) SpotPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroup) UserData() awsec2.UserData {
	var returns awsec2.UserData
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}


// Experimental.
func NewAutoScalingGroup(scope constructs.Construct, id *string, props *AutoScalingGroupProps) AutoScalingGroup {
	_init_.Initialize()

	if err := validateNewAutoScalingGroupParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoScalingGroup{}

	_jsii_.Create(
		"monocdk.aws_autoscaling.AutoScalingGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAutoScalingGroup_Override(a AutoScalingGroup, scope constructs.Construct, id *string, props *AutoScalingGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_autoscaling.AutoScalingGroup",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_AutoScalingGroup)SetAlbTargetGroup(val awselasticloadbalancingv2.ApplicationTargetGroup) {
	_jsii_.Set(
		j,
		"albTargetGroup",
		val,
	)
}

func (j *jsiiProxy_AutoScalingGroup)SetNewInstancesProtectedFromScaleIn(val *bool) {
	_jsii_.Set(
		j,
		"newInstancesProtectedFromScaleIn",
		val,
	)
}

// Experimental.
func AutoScalingGroup_FromAutoScalingGroupName(scope constructs.Construct, id *string, autoScalingGroupName *string) IAutoScalingGroup {
	_init_.Initialize()

	if err := validateAutoScalingGroup_FromAutoScalingGroupNameParameters(scope, id, autoScalingGroupName); err != nil {
		panic(err)
	}
	var returns IAutoScalingGroup

	_jsii_.StaticInvoke(
		"monocdk.aws_autoscaling.AutoScalingGroup",
		"fromAutoScalingGroupName",
		[]interface{}{scope, id, autoScalingGroupName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func AutoScalingGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutoScalingGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_autoscaling.AutoScalingGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func AutoScalingGroup_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateAutoScalingGroup_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_autoscaling.AutoScalingGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) AddLifecycleHook(id *string, props *BasicLifecycleHookProps) LifecycleHook {
	if err := a.validateAddLifecycleHookParameters(id, props); err != nil {
		panic(err)
	}
	var returns LifecycleHook

	_jsii_.Invoke(
		a,
		"addLifecycleHook",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) AddSecurityGroup(securityGroup awsec2.ISecurityGroup) {
	if err := a.validateAddSecurityGroupParameters(securityGroup); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addSecurityGroup",
		[]interface{}{securityGroup},
	)
}

func (a *jsiiProxy_AutoScalingGroup) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := a.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (a *jsiiProxy_AutoScalingGroup) AddUserData(commands ...*string) {
	args := []interface{}{}
	for _, a := range commands {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addUserData",
		args,
	)
}

func (a *jsiiProxy_AutoScalingGroup) AddWarmPool(options *WarmPoolOptions) WarmPool {
	if err := a.validateAddWarmPoolParameters(options); err != nil {
		panic(err)
	}
	var returns WarmPool

	_jsii_.Invoke(
		a,
		"addWarmPool",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) ApplyCloudFormationInit(init awsec2.CloudFormationInit, options *ApplyCloudFormationInitOptions) {
	if err := a.validateApplyCloudFormationInitParameters(init, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"applyCloudFormationInit",
		[]interface{}{init, options},
	)
}

func (a *jsiiProxy_AutoScalingGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := a.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_AutoScalingGroup) AreNewInstancesProtectedFromScaleIn() *bool {
	var returns *bool

	_jsii_.Invoke(
		a,
		"areNewInstancesProtectedFromScaleIn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	if err := a.validateAttachToApplicationTargetGroupParameters(targetGroup); err != nil {
		panic(err)
	}
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		a,
		"attachToApplicationTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer) {
	if err := a.validateAttachToClassicLBParameters(loadBalancer); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"attachToClassicLB",
		[]interface{}{loadBalancer},
	)
}

func (a *jsiiProxy_AutoScalingGroup) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	if err := a.validateAttachToNetworkTargetGroupParameters(targetGroup); err != nil {
		panic(err)
	}
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		a,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := a.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) GetResourceNameAttribute(nameAttr *string) *string {
	if err := a.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoScalingGroup) OnSynthesize(session constructs.ISynthesisSession) {
	if err := a.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (a *jsiiProxy_AutoScalingGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoScalingGroup) ProtectNewInstancesFromScaleIn() {
	_jsii_.InvokeVoid(
		a,
		"protectNewInstancesFromScaleIn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoScalingGroup) ScaleOnCpuUtilization(id *string, props *CpuUtilizationScalingProps) TargetTrackingScalingPolicy {
	if err := a.validateScaleOnCpuUtilizationParameters(id, props); err != nil {
		panic(err)
	}
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		a,
		"scaleOnCpuUtilization",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) ScaleOnIncomingBytes(id *string, props *NetworkUtilizationScalingProps) TargetTrackingScalingPolicy {
	if err := a.validateScaleOnIncomingBytesParameters(id, props); err != nil {
		panic(err)
	}
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		a,
		"scaleOnIncomingBytes",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) ScaleOnMetric(id *string, props *BasicStepScalingPolicyProps) StepScalingPolicy {
	if err := a.validateScaleOnMetricParameters(id, props); err != nil {
		panic(err)
	}
	var returns StepScalingPolicy

	_jsii_.Invoke(
		a,
		"scaleOnMetric",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) ScaleOnOutgoingBytes(id *string, props *NetworkUtilizationScalingProps) TargetTrackingScalingPolicy {
	if err := a.validateScaleOnOutgoingBytesParameters(id, props); err != nil {
		panic(err)
	}
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		a,
		"scaleOnOutgoingBytes",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) ScaleOnRequestCount(id *string, props *RequestCountScalingProps) TargetTrackingScalingPolicy {
	if err := a.validateScaleOnRequestCountParameters(id, props); err != nil {
		panic(err)
	}
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		a,
		"scaleOnRequestCount",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) ScaleOnSchedule(id *string, props *BasicScheduledActionProps) ScheduledAction {
	if err := a.validateScaleOnScheduleParameters(id, props); err != nil {
		panic(err)
	}
	var returns ScheduledAction

	_jsii_.Invoke(
		a,
		"scaleOnSchedule",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) ScaleToTrackMetric(id *string, props *MetricTargetTrackingProps) TargetTrackingScalingPolicy {
	if err := a.validateScaleToTrackMetricParameters(id, props); err != nil {
		panic(err)
	}
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		a,
		"scaleToTrackMetric",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) Synthesize(session awscdk.ISynthesisSession) {
	if err := a.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
}

func (a *jsiiProxy_AutoScalingGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}
