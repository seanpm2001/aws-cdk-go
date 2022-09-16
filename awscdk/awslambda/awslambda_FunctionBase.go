package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Experimental.
type FunctionBase interface {
	awscdk.Resource
	awsec2.IClientVpnConnectionHandler
	IFunction
	// The architecture of this Lambda Function.
	// Experimental.
	Architecture() Architecture
	// Whether the addPermission() call adds any permissions.
	//
	// True for new Lambdas, false for version $LATEST and imported Lambdas
	// from different accounts.
	// Experimental.
	CanCreatePermissions() *bool
	// Access the Connections object.
	//
	// Will fail if not a VPC-enabled Lambda Function.
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
	// The ARN fo the function.
	// Experimental.
	FunctionArn() *string
	// The name of the function.
	// Experimental.
	FunctionName() *string
	// The principal this Lambda Function is running as.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// Whether or not this Lambda function was bound to a VPC.
	//
	// If this is is `false`, trying to access the `connections` object will fail.
	// Experimental.
	IsBoundToVpc() *bool
	// The `$LATEST` version of this function.
	//
	// Note that this is reference to a non-specific AWS Lambda version, which
	// means the function this version refers to can return different results in
	// different invocations.
	//
	// To obtain a reference to an explicit version which references the current
	// function configuration, use `lambdaFunction.currentVersion` instead.
	// Experimental.
	LatestVersion() IVersion
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The construct node where permissions are attached.
	// Experimental.
	PermissionsNode() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The ARN(s) to put into the resource field of the generated IAM policy for grantInvoke().
	// Experimental.
	ResourceArnsForGrantInvoke() *[]*string
	// The IAM role associated with this function.
	//
	// Undefined if the function was imported without a role.
	// Experimental.
	Role() awsiam.IRole
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Adds an event source to this function.
	//
	// Event sources are implemented in the @aws-cdk/aws-lambda-event-sources module.
	//
	// The following example adds an SQS Queue as an event source:
	// ```
	// import { SqsEventSource } from '@aws-cdk/aws-lambda-event-sources';
	// myFunction.addEventSource(new SqsEventSource(myQueue));
	// ```.
	// Experimental.
	AddEventSource(source IEventSource)
	// Adds an event source that maps to this AWS Lambda function.
	// Experimental.
	AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping
	// Adds a url to this lambda function.
	// Experimental.
	AddFunctionUrl(options *FunctionUrlOptions) FunctionUrl
	// Adds a permission to the Lambda resource policy.
	// See: Permission for details.
	//
	// Experimental.
	AddPermission(id *string, permission *Permission)
	// Adds a statement to the IAM role assumed by the instance.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Configures options for asynchronous invocation.
	// Experimental.
	ConfigureAsyncInvoke(options *EventInvokeConfigOptions)
	// A warning will be added to functions under the following conditions: - permissions that include `lambda:InvokeFunction` are added to the unqualified function.
	//
	// - function.currentVersion is invoked before or after the permission is created.
	//
	// This applies only to permissions on Lambda functions, not versions or aliases.
	// This function is overridden as a noOp for QualifiedFunctionBase.
	// Experimental.
	ConsiderWarningOnInvokeFunctionPermissions(scope awscdk.Construct, action *string)
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
	// Grant the given identity permissions to invoke this Lambda.
	// Experimental.
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to invoke this Lambda Function URL.
	// Experimental.
	GrantInvokeUrl(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this Function.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// How long execution of this Lambda takes.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// How many invocations of this Lambda fail.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// How often this Lambda is invoked.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// How often this Lambda is throttled.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
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
	// Experimental.
	WarnInvokeFunctionPermissions(scope awscdk.Construct)
}

// The jsii proxy struct for FunctionBase
type jsiiProxy_FunctionBase struct {
	internal.Type__awscdkResource
	internal.Type__awsec2IClientVpnConnectionHandler
	jsiiProxy_IFunction
}

func (j *jsiiProxy_FunctionBase) Architecture() Architecture {
	var returns Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) CanCreatePermissions() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canCreatePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) IsBoundToVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isBoundToVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) LatestVersion() IVersion {
	var returns IVersion
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) PermissionsNode() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"permissionsNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) ResourceArnsForGrantInvoke() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceArnsForGrantInvoke",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewFunctionBase_Override(f FunctionBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lambda.FunctionBase",
		[]interface{}{scope, id, props},
		f,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func FunctionBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFunctionBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.FunctionBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func FunctionBase_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateFunctionBase_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.FunctionBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionBase) AddEventSource(source IEventSource) {
	if err := f.validateAddEventSourceParameters(source); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addEventSource",
		[]interface{}{source},
	)
}

func (f *jsiiProxy_FunctionBase) AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping {
	if err := f.validateAddEventSourceMappingParameters(id, options); err != nil {
		panic(err)
	}
	var returns EventSourceMapping

	_jsii_.Invoke(
		f,
		"addEventSourceMapping",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionBase) AddFunctionUrl(options *FunctionUrlOptions) FunctionUrl {
	if err := f.validateAddFunctionUrlParameters(options); err != nil {
		panic(err)
	}
	var returns FunctionUrl

	_jsii_.Invoke(
		f,
		"addFunctionUrl",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionBase) AddPermission(id *string, permission *Permission) {
	if err := f.validateAddPermissionParameters(id, permission); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addPermission",
		[]interface{}{id, permission},
	)
}

func (f *jsiiProxy_FunctionBase) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := f.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (f *jsiiProxy_FunctionBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := f.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (f *jsiiProxy_FunctionBase) ConfigureAsyncInvoke(options *EventInvokeConfigOptions) {
	if err := f.validateConfigureAsyncInvokeParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

func (f *jsiiProxy_FunctionBase) ConsiderWarningOnInvokeFunctionPermissions(scope awscdk.Construct, action *string) {
	if err := f.validateConsiderWarningOnInvokeFunctionPermissionsParameters(scope, action); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"considerWarningOnInvokeFunctionPermissions",
		[]interface{}{scope, action},
	)
}

func (f *jsiiProxy_FunctionBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := f.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionBase) GetResourceNameAttribute(nameAttr *string) *string {
	if err := f.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionBase) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := f.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		f,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionBase) GrantInvokeUrl(grantee awsiam.IGrantable) awsiam.Grant {
	if err := f.validateGrantInvokeUrlParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		f,
		"grantInvokeUrl",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionBase) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := f.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		f,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionBase) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := f.validateMetricDurationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		f,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionBase) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := f.validateMetricErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		f,
		"metricErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionBase) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := f.validateMetricInvocationsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		f,
		"metricInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionBase) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := f.validateMetricThrottlesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		f,
		"metricThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionBase) OnPrepare() {
	_jsii_.InvokeVoid(
		f,
		"onPrepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionBase) OnSynthesize(session constructs.ISynthesisSession) {
	if err := f.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (f *jsiiProxy_FunctionBase) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionBase) Prepare() {
	_jsii_.InvokeVoid(
		f,
		"prepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionBase) Synthesize(session awscdk.ISynthesisSession) {
	if err := f.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"synthesize",
		[]interface{}{session},
	)
}

func (f *jsiiProxy_FunctionBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionBase) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionBase) WarnInvokeFunctionPermissions(scope awscdk.Construct) {
	if err := f.validateWarnInvokeFunctionPermissionsParameters(scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"warnInvokeFunctionPermissions",
		[]interface{}{scope},
	)
}
