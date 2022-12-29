// The CDK Construct Library for AWS Lambda in Golang
package awscdklambdagoalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Golang Lambda function.
//
// Example:
//   go.NewGoFunction(this, jsii.String("handler"), &goFunctionProps{
//   	entry: jsii.String("app/cmd/api"),
//   	bundling: &bundlingOptions{
//   		environment: map[string]*string{
//   			"HELLO": jsii.String("WORLD"),
//   		},
//   	},
//   })
//
// Experimental.
type GoFunction interface {
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
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The construct node where permissions are attached.
	// Experimental.
	PermissionsNode() constructs.Node
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
	// });
	// ```.
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
	ConsiderWarningOnInvokeFunctionPermissions(scope constructs.Construct, action *string)
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
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Experimental.
	WarnInvokeFunctionPermissions(scope constructs.Construct)
}

// The jsii proxy struct for GoFunction
type jsiiProxy_GoFunction struct {
	internal.Type__awslambdaFunction
}

func (j *jsiiProxy_GoFunction) Architecture() awslambda.Architecture {
	var returns awslambda.Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) CanCreatePermissions() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canCreatePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) CurrentVersion() awslambda.Version {
	var returns awslambda.Version
	_jsii_.Get(
		j,
		"currentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) DeadLetterQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) DeadLetterTopic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"deadLetterTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) IsBoundToVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isBoundToVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) LatestVersion() awslambda.IVersion {
	var returns awslambda.IVersion
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) PermissionsNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"permissionsNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) ResourceArnsForGrantInvoke() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceArnsForGrantInvoke",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) Runtime() awslambda.Runtime {
	var returns awslambda.Runtime
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) Timeout() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}


// Experimental.
func NewGoFunction(scope constructs.Construct, id *string, props *GoFunctionProps) GoFunction {
	_init_.Initialize()

	if err := validateNewGoFunctionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoFunction{}

	_jsii_.Create(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewGoFunction_Override(g GoFunction, scope constructs.Construct, id *string, props *GoFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		[]interface{}{scope, id, props},
		g,
	)
}

// Record whether specific properties in the `AWS::Lambda::Function` resource should also be associated to the Version resource.
//
// See 'currentVersion' section in the module README for more details.
// Experimental.
func GoFunction_ClassifyVersionProperty(propertyName *string, locked *bool) {
	_init_.Initialize()

	if err := validateGoFunction_ClassifyVersionPropertyParameters(propertyName, locked); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"classifyVersionProperty",
		[]interface{}{propertyName, locked},
	)
}

// Import a lambda function into the CDK using its ARN.
// Experimental.
func GoFunction_FromFunctionArn(scope constructs.Construct, id *string, functionArn *string) awslambda.IFunction {
	_init_.Initialize()

	if err := validateGoFunction_FromFunctionArnParameters(scope, id, functionArn); err != nil {
		panic(err)
	}
	var returns awslambda.IFunction

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"fromFunctionArn",
		[]interface{}{scope, id, functionArn},
		&returns,
	)

	return returns
}

// Creates a Lambda function object which represents a function not defined within this stack.
// Experimental.
func GoFunction_FromFunctionAttributes(scope constructs.Construct, id *string, attrs *awslambda.FunctionAttributes) awslambda.IFunction {
	_init_.Initialize()

	if err := validateGoFunction_FromFunctionAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns awslambda.IFunction

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"fromFunctionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import a lambda function into the CDK using its name.
// Experimental.
func GoFunction_FromFunctionName(scope constructs.Construct, id *string, functionName *string) awslambda.IFunction {
	_init_.Initialize()

	if err := validateGoFunction_FromFunctionNameParameters(scope, id, functionName); err != nil {
		panic(err)
	}
	var returns awslambda.IFunction

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"fromFunctionName",
		[]interface{}{scope, id, functionName},
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
func GoFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoFunction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func GoFunction_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateGoFunction_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func GoFunction_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateGoFunction_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return the given named metric for this Lambda.
// Experimental.
func GoFunction_MetricAll(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateGoFunction_MetricAllParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"metricAll",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of concurrent executions across all Lambdas.
// Experimental.
func GoFunction_MetricAllConcurrentExecutions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateGoFunction_MetricAllConcurrentExecutionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"metricAllConcurrentExecutions",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the Duration executing all Lambdas.
// Experimental.
func GoFunction_MetricAllDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateGoFunction_MetricAllDurationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"metricAllDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of Errors executing all Lambdas.
// Experimental.
func GoFunction_MetricAllErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateGoFunction_MetricAllErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"metricAllErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of invocations of all Lambdas.
// Experimental.
func GoFunction_MetricAllInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateGoFunction_MetricAllInvocationsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"metricAllInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of throttled invocations of all Lambdas.
// Experimental.
func GoFunction_MetricAllThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateGoFunction_MetricAllThrottlesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"metricAllThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of unreserved concurrent executions across all Lambdas.
// Experimental.
func GoFunction_MetricAllUnreservedConcurrentExecutions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateGoFunction_MetricAllUnreservedConcurrentExecutionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"metricAllUnreservedConcurrentExecutions",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func GoFunction_GOOGLE_GOPROXY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"GOOGLE_GOPROXY",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoFunction) AddAlias(aliasName *string, options *awslambda.AliasOptions) awslambda.Alias {
	if err := g.validateAddAliasParameters(aliasName, options); err != nil {
		panic(err)
	}
	var returns awslambda.Alias

	_jsii_.Invoke(
		g,
		"addAlias",
		[]interface{}{aliasName, options},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoFunction) AddEnvironment(key *string, value *string, options *awslambda.EnvironmentOptions) awslambda.Function {
	if err := g.validateAddEnvironmentParameters(key, value, options); err != nil {
		panic(err)
	}
	var returns awslambda.Function

	_jsii_.Invoke(
		g,
		"addEnvironment",
		[]interface{}{key, value, options},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoFunction) AddEventSource(source awslambda.IEventSource) {
	if err := g.validateAddEventSourceParameters(source); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addEventSource",
		[]interface{}{source},
	)
}

func (g *jsiiProxy_GoFunction) AddEventSourceMapping(id *string, options *awslambda.EventSourceMappingOptions) awslambda.EventSourceMapping {
	if err := g.validateAddEventSourceMappingParameters(id, options); err != nil {
		panic(err)
	}
	var returns awslambda.EventSourceMapping

	_jsii_.Invoke(
		g,
		"addEventSourceMapping",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoFunction) AddFunctionUrl(options *awslambda.FunctionUrlOptions) awslambda.FunctionUrl {
	if err := g.validateAddFunctionUrlParameters(options); err != nil {
		panic(err)
	}
	var returns awslambda.FunctionUrl

	_jsii_.Invoke(
		g,
		"addFunctionUrl",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoFunction) AddLayers(layers ...awslambda.ILayerVersion) {
	args := []interface{}{}
	for _, a := range layers {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		g,
		"addLayers",
		args,
	)
}

func (g *jsiiProxy_GoFunction) AddPermission(id *string, permission *awslambda.Permission) {
	if err := g.validateAddPermissionParameters(id, permission); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addPermission",
		[]interface{}{id, permission},
	)
}

func (g *jsiiProxy_GoFunction) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := g.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (g *jsiiProxy_GoFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := g.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (g *jsiiProxy_GoFunction) ConfigureAsyncInvoke(options *awslambda.EventInvokeConfigOptions) {
	if err := g.validateConfigureAsyncInvokeParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

func (g *jsiiProxy_GoFunction) ConsiderWarningOnInvokeFunctionPermissions(scope constructs.Construct, action *string) {
	if err := g.validateConsiderWarningOnInvokeFunctionPermissionsParameters(scope, action); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"considerWarningOnInvokeFunctionPermissions",
		[]interface{}{scope, action},
	)
}

func (g *jsiiProxy_GoFunction) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoFunction) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := g.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoFunction) GetResourceNameAttribute(nameAttr *string) *string {
	if err := g.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoFunction) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := g.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		g,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoFunction) GrantInvokeUrl(grantee awsiam.IGrantable) awsiam.Grant {
	if err := g.validateGrantInvokeUrlParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		g,
		"grantInvokeUrl",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoFunction) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoFunction) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricDurationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoFunction) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoFunction) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricInvocationsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoFunction) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricThrottlesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoFunction) WarnInvokeFunctionPermissions(scope constructs.Construct) {
	if err := g.validateWarnInvokeFunctionPermissionsParameters(scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"warnInvokeFunctionPermissions",
		[]interface{}{scope},
	)
}
