package awslambdapython

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awslambdapython/internal"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
	"github.com/aws/constructs-go/constructs/v3"
)

// A Python Lambda function.
//
// Example:
//   entry := "/path/to/function"
//   image := awscdk.DockerImage.fromBuild(entry)
//
//   lambda.NewPythonFunction(this, jsii.String("function"), &pythonFunctionProps{
//   	entry: jsii.String(entry),
//   	runtime: awscdk.Runtime_PYTHON_3_8(),
//   	bundling: &bundlingOptions{
//   		buildArgs: map[string]*string{
//   			"PIP_INDEX_URL": jsii.String("https://your.index.url/simple/"),
//   			"PIP_EXTRA_INDEX_URL": jsii.String("https://your.extra-index.url/simple/"),
//   		},
//   	},
//   })
//
// Experimental.
type PythonFunction interface {
	awslambda.Function
	// The architecture of this Lambda Function (this is an optional attribute and defaults to X86_64).
	// Experimental.
	Architecture() awslambda.Architecture
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
	// Returns a `lambda.Version` which represents the current version of this Lambda function. A new version will be created every time the function's configuration changes.
	//
	// You can specify options for this version using the `currentVersionOptions`
	// prop when initializing the `lambda.Function`.
	// Experimental.
	CurrentVersion() awslambda.Version
	// The DLQ (as queue) associated with this Lambda Function (this is an optional attribute).
	// Experimental.
	DeadLetterQueue() awssqs.IQueue
	// The DLQ (as topic) associated with this Lambda Function (this is an optional attribute).
	// Experimental.
	DeadLetterTopic() awssns.ITopic
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
	// ARN of this function.
	// Experimental.
	FunctionArn() *string
	// Name of this function.
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
	LatestVersion() awslambda.IVersion
	// The LogGroup where the Lambda function's logs are made available.
	//
	// If either `logRetention` is set or this property is called, a CloudFormation custom resource is added to the stack that
	// pre-creates the log group as part of the stack deployment, if it already doesn't exist, and sets the correct log retention
	// period (never expire, by default).
	//
	// Further, if the log group already exists and the `logRetention` is not set, the custom resource will reset the log retention
	// to never expire even if it was configured with a different value.
	// Experimental.
	LogGroup() awslogs.ILogGroup
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
	// Execution role associated with this function.
	// Experimental.
	Role() awsiam.IRole
	// The runtime configured for this lambda.
	// Experimental.
	Runtime() awslambda.Runtime
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// The timeout configured for this lambda.
	// Experimental.
	Timeout() awscdk.Duration
	// Defines an alias for this function.
	//
	// The alias will automatically be updated to point to the latest version of
	// the function as it is being updated during a deployment.
	//
	// ```ts
	// declare const fn: lambda.Function;
	//
	// fn.addAlias('Live');
	//
	// // Is equivalent to
	//
	// new lambda.Alias(this, 'AliasLive', {
	//    aliasName: 'Live',
	//    version: fn.currentVersion,
	// });.
	// Experimental.
	AddAlias(aliasName *string, options *awslambda.AliasOptions) awslambda.Alias
	// Adds an environment variable to this Lambda function.
	//
	// If this is a ref to a Lambda function, this operation results in a no-op.
	// Experimental.
	AddEnvironment(key *string, value *string, options *awslambda.EnvironmentOptions) awslambda.Function
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
	AddEventSource(source awslambda.IEventSource)
	// Adds an event source that maps to this AWS Lambda function.
	// Experimental.
	AddEventSourceMapping(id *string, options *awslambda.EventSourceMappingOptions) awslambda.EventSourceMapping
	// Adds a url to this lambda function.
	// Experimental.
	AddFunctionUrl(options *awslambda.FunctionUrlOptions) awslambda.FunctionUrl
	// Adds one or more Lambda Layers to this Lambda function.
	// Experimental.
	AddLayers(layers ...awslambda.ILayerVersion)
	// Adds a permission to the Lambda resource policy.
	// See: Permission for details.
	//
	// Experimental.
	AddPermission(id *string, permission *awslambda.Permission)
	// Adds a statement to the IAM role assumed by the instance.
	// Experimental.
	AddToRolePolicy(statement awsiam.PolicyStatement)
	// Add a new version for this Lambda.
	//
	// If you want to deploy through CloudFormation and use aliases, you need to
	// add a new version (with a new name) to your Lambda every time you want to
	// deploy an update. An alias can then refer to the newly created Version.
	//
	// All versions should have distinct names, and you should not delete versions
	// as long as your Alias needs to refer to them.
	//
	// Returns: A new Version object.
	// Deprecated: This method will create an AWS::Lambda::Version resource which
	// snapshots the AWS Lambda function *at the time of its creation* and it
	// won't get updated when the function changes. Instead, use
	// `this.currentVersion` to obtain a reference to a version resource that gets
	// automatically recreated when the function configuration (or code) changes.
	AddVersion(name *string, codeSha256 *string, description *string, provisionedExecutions *float64, asyncInvokeConfig *awslambda.EventInvokeConfigOptions) awslambda.Version
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
	ConfigureAsyncInvoke(options *awslambda.EventInvokeConfigOptions)
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

// The jsii proxy struct for PythonFunction
type jsiiProxy_PythonFunction struct {
	internal.Type__awslambdaFunction
}

func (j *jsiiProxy_PythonFunction) Architecture() awslambda.Architecture {
	var returns awslambda.Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) CanCreatePermissions() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canCreatePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) CurrentVersion() awslambda.Version {
	var returns awslambda.Version
	_jsii_.Get(
		j,
		"currentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) DeadLetterQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) DeadLetterTopic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"deadLetterTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) IsBoundToVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isBoundToVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) LatestVersion() awslambda.IVersion {
	var returns awslambda.IVersion
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) PermissionsNode() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"permissionsNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) ResourceArnsForGrantInvoke() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceArnsForGrantInvoke",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) Runtime() awslambda.Runtime {
	var returns awslambda.Runtime
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) Timeout() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}


// Experimental.
func NewPythonFunction(scope awscdk.Construct, id *string, props *PythonFunctionProps) PythonFunction {
	_init_.Initialize()

	if err := validateNewPythonFunctionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PythonFunction{}

	_jsii_.Create(
		"monocdk.aws_lambda_python.PythonFunction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewPythonFunction_Override(p PythonFunction, scope awscdk.Construct, id *string, props *PythonFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lambda_python.PythonFunction",
		[]interface{}{scope, id, props},
		p,
	)
}

// Record whether specific properties in the `AWS::Lambda::Function` resource should also be associated to the Version resource.
//
// See 'currentVersion' section in the module README for more details.
// Experimental.
func PythonFunction_ClassifyVersionProperty(propertyName *string, locked *bool) {
	_init_.Initialize()

	if err := validatePythonFunction_ClassifyVersionPropertyParameters(propertyName, locked); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"monocdk.aws_lambda_python.PythonFunction",
		"classifyVersionProperty",
		[]interface{}{propertyName, locked},
	)
}

// Import a lambda function into the CDK using its ARN.
// Experimental.
func PythonFunction_FromFunctionArn(scope constructs.Construct, id *string, functionArn *string) awslambda.IFunction {
	_init_.Initialize()

	if err := validatePythonFunction_FromFunctionArnParameters(scope, id, functionArn); err != nil {
		panic(err)
	}
	var returns awslambda.IFunction

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
		"fromFunctionArn",
		[]interface{}{scope, id, functionArn},
		&returns,
	)

	return returns
}

// Creates a Lambda function object which represents a function not defined within this stack.
// Experimental.
func PythonFunction_FromFunctionAttributes(scope constructs.Construct, id *string, attrs *awslambda.FunctionAttributes) awslambda.IFunction {
	_init_.Initialize()

	if err := validatePythonFunction_FromFunctionAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns awslambda.IFunction

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
		"fromFunctionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import a lambda function into the CDK using its name.
// Experimental.
func PythonFunction_FromFunctionName(scope constructs.Construct, id *string, functionName *string) awslambda.IFunction {
	_init_.Initialize()

	if err := validatePythonFunction_FromFunctionNameParameters(scope, id, functionName); err != nil {
		panic(err)
	}
	var returns awslambda.IFunction

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
		"fromFunctionName",
		[]interface{}{scope, id, functionName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func PythonFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePythonFunction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func PythonFunction_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validatePythonFunction_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return the given named metric for this Lambda.
// Experimental.
func PythonFunction_MetricAll(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validatePythonFunction_MetricAllParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
		"metricAll",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of concurrent executions across all Lambdas.
// Experimental.
func PythonFunction_MetricAllConcurrentExecutions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validatePythonFunction_MetricAllConcurrentExecutionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
		"metricAllConcurrentExecutions",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the Duration executing all Lambdas.
// Experimental.
func PythonFunction_MetricAllDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validatePythonFunction_MetricAllDurationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
		"metricAllDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of Errors executing all Lambdas.
// Experimental.
func PythonFunction_MetricAllErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validatePythonFunction_MetricAllErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
		"metricAllErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of invocations of all Lambdas.
// Experimental.
func PythonFunction_MetricAllInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validatePythonFunction_MetricAllInvocationsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
		"metricAllInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of throttled invocations of all Lambdas.
// Experimental.
func PythonFunction_MetricAllThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validatePythonFunction_MetricAllThrottlesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
		"metricAllThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of unreserved concurrent executions across all Lambdas.
// Experimental.
func PythonFunction_MetricAllUnreservedConcurrentExecutions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validatePythonFunction_MetricAllUnreservedConcurrentExecutionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
		"metricAllUnreservedConcurrentExecutions",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonFunction) AddAlias(aliasName *string, options *awslambda.AliasOptions) awslambda.Alias {
	if err := p.validateAddAliasParameters(aliasName, options); err != nil {
		panic(err)
	}
	var returns awslambda.Alias

	_jsii_.Invoke(
		p,
		"addAlias",
		[]interface{}{aliasName, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonFunction) AddEnvironment(key *string, value *string, options *awslambda.EnvironmentOptions) awslambda.Function {
	if err := p.validateAddEnvironmentParameters(key, value, options); err != nil {
		panic(err)
	}
	var returns awslambda.Function

	_jsii_.Invoke(
		p,
		"addEnvironment",
		[]interface{}{key, value, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonFunction) AddEventSource(source awslambda.IEventSource) {
	if err := p.validateAddEventSourceParameters(source); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addEventSource",
		[]interface{}{source},
	)
}

func (p *jsiiProxy_PythonFunction) AddEventSourceMapping(id *string, options *awslambda.EventSourceMappingOptions) awslambda.EventSourceMapping {
	if err := p.validateAddEventSourceMappingParameters(id, options); err != nil {
		panic(err)
	}
	var returns awslambda.EventSourceMapping

	_jsii_.Invoke(
		p,
		"addEventSourceMapping",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonFunction) AddFunctionUrl(options *awslambda.FunctionUrlOptions) awslambda.FunctionUrl {
	if err := p.validateAddFunctionUrlParameters(options); err != nil {
		panic(err)
	}
	var returns awslambda.FunctionUrl

	_jsii_.Invoke(
		p,
		"addFunctionUrl",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonFunction) AddLayers(layers ...awslambda.ILayerVersion) {
	args := []interface{}{}
	for _, a := range layers {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addLayers",
		args,
	)
}

func (p *jsiiProxy_PythonFunction) AddPermission(id *string, permission *awslambda.Permission) {
	if err := p.validateAddPermissionParameters(id, permission); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addPermission",
		[]interface{}{id, permission},
	)
}

func (p *jsiiProxy_PythonFunction) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := p.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (p *jsiiProxy_PythonFunction) AddVersion(name *string, codeSha256 *string, description *string, provisionedExecutions *float64, asyncInvokeConfig *awslambda.EventInvokeConfigOptions) awslambda.Version {
	if err := p.validateAddVersionParameters(name, asyncInvokeConfig); err != nil {
		panic(err)
	}
	var returns awslambda.Version

	_jsii_.Invoke(
		p,
		"addVersion",
		[]interface{}{name, codeSha256, description, provisionedExecutions, asyncInvokeConfig},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := p.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (p *jsiiProxy_PythonFunction) ConfigureAsyncInvoke(options *awslambda.EventInvokeConfigOptions) {
	if err := p.validateConfigureAsyncInvokeParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

func (p *jsiiProxy_PythonFunction) ConsiderWarningOnInvokeFunctionPermissions(scope awscdk.Construct, action *string) {
	if err := p.validateConsiderWarningOnInvokeFunctionPermissionsParameters(scope, action); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"considerWarningOnInvokeFunctionPermissions",
		[]interface{}{scope, action},
	)
}

func (p *jsiiProxy_PythonFunction) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonFunction) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := p.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonFunction) GetResourceNameAttribute(nameAttr *string) *string {
	if err := p.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonFunction) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := p.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		p,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonFunction) GrantInvokeUrl(grantee awsiam.IGrantable) awsiam.Grant {
	if err := p.validateGrantInvokeUrlParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		p,
		"grantInvokeUrl",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonFunction) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := p.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonFunction) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := p.validateMetricDurationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonFunction) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := p.validateMetricErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
		"metricErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonFunction) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := p.validateMetricInvocationsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
		"metricInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonFunction) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := p.validateMetricThrottlesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
		"metricThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonFunction) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PythonFunction) OnSynthesize(session constructs.ISynthesisSession) {
	if err := p.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (p *jsiiProxy_PythonFunction) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonFunction) Prepare() {
	_jsii_.InvokeVoid(
		p,
		"prepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PythonFunction) Synthesize(session awscdk.ISynthesisSession) {
	if err := p.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		[]interface{}{session},
	)
}

func (p *jsiiProxy_PythonFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonFunction) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonFunction) WarnInvokeFunctionPermissions(scope awscdk.Construct) {
	if err := p.validateWarnInvokeFunctionPermissionsParameters(scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"warnInvokeFunctionPermissions",
		[]interface{}{scope},
	)
}

