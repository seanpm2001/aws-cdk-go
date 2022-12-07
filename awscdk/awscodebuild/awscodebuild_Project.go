package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarnotifications"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// A representation of a CodeBuild Project.
//
// Example:
//   var bucket bucket
//
//
//   project := codebuild.NewProject(this, jsii.String("MyProject"), &projectProps{
//   	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   	}),
//   	artifacts: codebuild.artifacts.s3(&s3ArtifactsProps{
//   		bucket: bucket,
//   		includeBuildId: jsii.Boolean(false),
//   		packageZip: jsii.Boolean(true),
//   		path: jsii.String("another/path"),
//   		identifier: jsii.String("AddArtifact1"),
//   	}),
//   })
//
type Project interface {
	awscdk.Resource
	IProject
	// Access the Connections object.
	//
	// Will fail if this Project does not have a VPC set.
	Connections() awsec2.Connections
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The principal to grant permissions to.
	GrantPrincipal() awsiam.IPrincipal
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
	// The ARN of the project.
	ProjectArn() *string
	// The name of the project.
	ProjectName() *string
	// The IAM role for this project.
	Role() awsiam.IRole
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Adds a fileSystemLocation to the Project.
	AddFileSystemLocation(fileSystemLocation IFileSystemLocation)
	// Adds a secondary artifact to the Project.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-multi-in-out.html
	//
	AddSecondaryArtifact(secondaryArtifact IArtifacts)
	// Adds a secondary source to the Project.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-multi-in-out.html
	//
	AddSecondarySource(secondarySource ISource)
	// Add a permission only if there's a policy attached.
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
	// Returns a source configuration for notification rule.
	BindAsNotificationRuleSource(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleSourceConfig
	// A callback invoked when the given project is added to a CodePipeline.
	BindToCodePipeline(_scope constructs.Construct, options *BindToCodePipelineOptions)
	// Enable batch builds.
	//
	// Returns an object contining the batch service role if batch builds
	// could be enabled.
	EnableBatchBuilds() *BatchBuildConfig
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
	// Returns: a CloudWatch metric associated with this build project.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Measures the number of builds triggered.
	//
	// Units: Count
	//
	// Valid CloudWatch statistics: Sum.
	MetricBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Measures the duration of all builds over time.
	//
	// Units: Seconds
	//
	// Valid CloudWatch statistics: Average (recommended), Maximum, Minimum.
	MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Measures the number of builds that failed because of client error or because of a timeout.
	//
	// Units: Count
	//
	// Valid CloudWatch statistics: Sum.
	MetricFailedBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Measures the number of successful builds.
	//
	// Units: Count
	//
	// Valid CloudWatch statistics: Sum.
	MetricSucceededBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Defines a CodeStar Notification rule triggered when the project events emitted by you specified, it very similar to `onEvent` API.
	//
	// You can also use the methods `notifyOnBuildSucceeded` and
	// `notifyOnBuildFailed` to define rules for these specific event emitted.
	NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *ProjectNotifyOnOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar notification rule which triggers when a build fails.
	NotifyOnBuildFailed(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar notification rule which triggers when a build completes successfully.
	NotifyOnBuildSucceeded(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines an event rule which triggers when a build fails.
	//
	// To access fields from the event in the event target input,
	// use the static fields on the `StateChangeEvent` class.
	OnBuildFailed(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an event rule which triggers when a build starts.
	//
	// To access fields from the event in the event target input,
	// use the static fields on the `StateChangeEvent` class.
	OnBuildStarted(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an event rule which triggers when a build completes successfully.
	//
	// To access fields from the event in the event target input,
	// use the static fields on the `StateChangeEvent` class.
	OnBuildSucceeded(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule triggered when something happens with this project.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-notifications.html
	//
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule that triggers upon phase change of this build project.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-notifications.html
	//
	OnPhaseChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule triggered when the build project state changes.
	//
	// You can filter specific build status events using an event
	// pattern filter on the `build-status` detail field:
	//
	//     const rule = project.onStateChange('OnBuildStarted', { target });
	//     rule.addEventPattern({
	//       detail: {
	//         'build-status': [
	//           "IN_PROGRESS",
	//           "SUCCEEDED",
	//           "FAILED",
	//           "STOPPED"
	//         ]
	//       }
	//     });
	//
	// You can also use the methods `onBuildFailed` and `onBuildSucceeded` to define rules for
	// these specific state changes.
	//
	// To access fields from the event in the event target input,
	// use the static fields on the `StateChangeEvent` class.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-notifications.html
	//
	OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Project
type jsiiProxy_Project struct {
	internal.Type__awscdkResource
	jsiiProxy_IProject
}

func (j *jsiiProxy_Project) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ProjectArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ProjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewProject(scope constructs.Construct, id *string, props *ProjectProps) Project {
	_init_.Initialize()

	if err := validateNewProjectParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Project{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.Project",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewProject_Override(p Project, scope constructs.Construct, id *string, props *ProjectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.Project",
		[]interface{}{scope, id, props},
		p,
	)
}

func Project_FromProjectArn(scope constructs.Construct, id *string, projectArn *string) IProject {
	_init_.Initialize()

	if err := validateProject_FromProjectArnParameters(scope, id, projectArn); err != nil {
		panic(err)
	}
	var returns IProject

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Project",
		"fromProjectArn",
		[]interface{}{scope, id, projectArn},
		&returns,
	)

	return returns
}

// Import a Project defined either outside the CDK, or in a different CDK Stack (and exported using the {@link export} method).
//
// Returns: a reference to the existing Project.
func Project_FromProjectName(scope constructs.Construct, id *string, projectName *string) IProject {
	_init_.Initialize()

	if err := validateProject_FromProjectNameParameters(scope, id, projectName); err != nil {
		panic(err)
	}
	var returns IProject

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Project",
		"fromProjectName",
		[]interface{}{scope, id, projectName},
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
func Project_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProject_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Project",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Project_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateProject_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Project",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Project_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateProject_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Project",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Convert the environment variables map of string to {@link BuildEnvironmentVariable}, which is the customer-facing type, to a list of {@link CfnProject.EnvironmentVariableProperty}, which is the representation of environment variables in CloudFormation.
//
// Returns: an array of {@link CfnProject.EnvironmentVariableProperty} instances
func Project_SerializeEnvVariables(environmentVariables *map[string]*BuildEnvironmentVariable, validateNoPlainTextSecrets *bool, principal awsiam.IGrantable) *[]*CfnProject_EnvironmentVariableProperty {
	_init_.Initialize()

	if err := validateProject_SerializeEnvVariablesParameters(environmentVariables); err != nil {
		panic(err)
	}
	var returns *[]*CfnProject_EnvironmentVariableProperty

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Project",
		"serializeEnvVariables",
		[]interface{}{environmentVariables, validateNoPlainTextSecrets, principal},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) AddFileSystemLocation(fileSystemLocation IFileSystemLocation) {
	if err := p.validateAddFileSystemLocationParameters(fileSystemLocation); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addFileSystemLocation",
		[]interface{}{fileSystemLocation},
	)
}

func (p *jsiiProxy_Project) AddSecondaryArtifact(secondaryArtifact IArtifacts) {
	if err := p.validateAddSecondaryArtifactParameters(secondaryArtifact); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addSecondaryArtifact",
		[]interface{}{secondaryArtifact},
	)
}

func (p *jsiiProxy_Project) AddSecondarySource(secondarySource ISource) {
	if err := p.validateAddSecondarySourceParameters(secondarySource); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addSecondarySource",
		[]interface{}{secondarySource},
	)
}

func (p *jsiiProxy_Project) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := p.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (p *jsiiProxy_Project) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := p.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (p *jsiiProxy_Project) BindAsNotificationRuleSource(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleSourceConfig {
	if err := p.validateBindAsNotificationRuleSourceParameters(_scope); err != nil {
		panic(err)
	}
	var returns *awscodestarnotifications.NotificationRuleSourceConfig

	_jsii_.Invoke(
		p,
		"bindAsNotificationRuleSource",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) BindToCodePipeline(_scope constructs.Construct, options *BindToCodePipelineOptions) {
	if err := p.validateBindToCodePipelineParameters(_scope, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"bindToCodePipeline",
		[]interface{}{_scope, options},
	)
}

func (p *jsiiProxy_Project) EnableBatchBuilds() *BatchBuildConfig {
	var returns *BatchBuildConfig

	_jsii_.Invoke(
		p,
		"enableBatchBuilds",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (p *jsiiProxy_Project) GetResourceNameAttribute(nameAttr *string) *string {
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

func (p *jsiiProxy_Project) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (p *jsiiProxy_Project) MetricBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := p.validateMetricBuildsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
		"metricBuilds",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (p *jsiiProxy_Project) MetricFailedBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := p.validateMetricFailedBuildsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
		"metricFailedBuilds",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) MetricSucceededBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := p.validateMetricSucceededBuildsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
		"metricSucceededBuilds",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *ProjectNotifyOnOptions) awscodestarnotifications.INotificationRule {
	if err := p.validateNotifyOnParameters(id, target, options); err != nil {
		panic(err)
	}
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		p,
		"notifyOn",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) NotifyOnBuildFailed(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	if err := p.validateNotifyOnBuildFailedParameters(id, target, options); err != nil {
		panic(err)
	}
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		p,
		"notifyOnBuildFailed",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) NotifyOnBuildSucceeded(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	if err := p.validateNotifyOnBuildSucceededParameters(id, target, options); err != nil {
		panic(err)
	}
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		p,
		"notifyOnBuildSucceeded",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) OnBuildFailed(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := p.validateOnBuildFailedParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onBuildFailed",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) OnBuildStarted(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := p.validateOnBuildStartedParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onBuildStarted",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) OnBuildSucceeded(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := p.validateOnBuildSucceededParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onBuildSucceeded",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := p.validateOnEventParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) OnPhaseChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := p.validateOnPhaseChangeParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onPhaseChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := p.validateOnStateChangeParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onStateChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

