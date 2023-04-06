package awschatbot

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awschatbot/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarnotifications"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
)

// A new Slack channel configuration.
//
// Example:
//   import chatbot "github.com/aws/aws-cdk-go/awscdk"
//
//   var project project
//
//
//   target := chatbot.NewSlackChannelConfiguration(this, jsii.String("MySlackChannel"), &SlackChannelConfigurationProps{
//   	SlackChannelConfigurationName: jsii.String("YOUR_CHANNEL_NAME"),
//   	SlackWorkspaceId: jsii.String("YOUR_SLACK_WORKSPACE_ID"),
//   	SlackChannelId: jsii.String("YOUR_SLACK_CHANNEL_ID"),
//   })
//
//   rule := project.notifyOnBuildSucceeded(jsii.String("NotifyOnBuildSucceeded"), target)
//
type SlackChannelConfiguration interface {
	awscdk.Resource
	ISlackChannelConfiguration
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
	//   cross-environment scenarios.
	PhysicalName() *string
	// The permission role of Slack channel configuration.
	Role() awsiam.IRole
	// The ARN of the Slack channel configuration In the form of arn:aws:chatbot:{region}:{account}:chat-configuration/slack-channel/{slackChannelName}.
	SlackChannelConfigurationArn() *string
	// The name of Slack channel configuration.
	SlackChannelConfigurationName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Adds a SNS topic that deliver notifications to AWS Chatbot.
	AddNotificationTopic(notificationTopic awssns.ITopic)
	// Adds extra permission to iam-role of Slack channel configuration.
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
	// Returns a target configuration for notification rule.
	BindAsNotificationRuleTarget(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleTargetConfig
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
	// Return the given named metric for this SlackChannelConfiguration.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for SlackChannelConfiguration
type jsiiProxy_SlackChannelConfiguration struct {
	internal.Type__awscdkResource
	jsiiProxy_ISlackChannelConfiguration
}

func (j *jsiiProxy_SlackChannelConfiguration) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackChannelConfiguration) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackChannelConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackChannelConfiguration) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackChannelConfiguration) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackChannelConfiguration) SlackChannelConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackChannelConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackChannelConfiguration) SlackChannelConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackChannelConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SlackChannelConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewSlackChannelConfiguration(scope constructs.Construct, id *string, props *SlackChannelConfigurationProps) SlackChannelConfiguration {
	_init_.Initialize()

	if err := validateNewSlackChannelConfigurationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SlackChannelConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_chatbot.SlackChannelConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSlackChannelConfiguration_Override(s SlackChannelConfiguration, scope constructs.Construct, id *string, props *SlackChannelConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_chatbot.SlackChannelConfiguration",
		[]interface{}{scope, id, props},
		s,
	)
}

// Import an existing Slack channel configuration provided an ARN.
//
// Returns: a reference to the existing Slack channel configuration.
func SlackChannelConfiguration_FromSlackChannelConfigurationArn(scope constructs.Construct, id *string, slackChannelConfigurationArn *string) ISlackChannelConfiguration {
	_init_.Initialize()

	if err := validateSlackChannelConfiguration_FromSlackChannelConfigurationArnParameters(scope, id, slackChannelConfigurationArn); err != nil {
		panic(err)
	}
	var returns ISlackChannelConfiguration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_chatbot.SlackChannelConfiguration",
		"fromSlackChannelConfigurationArn",
		[]interface{}{scope, id, slackChannelConfigurationArn},
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
func SlackChannelConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSlackChannelConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_chatbot.SlackChannelConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func SlackChannelConfiguration_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateSlackChannelConfiguration_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_chatbot.SlackChannelConfiguration",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func SlackChannelConfiguration_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateSlackChannelConfiguration_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_chatbot.SlackChannelConfiguration",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return the given named metric for All SlackChannelConfigurations.
func SlackChannelConfiguration_MetricAll(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateSlackChannelConfiguration_MetricAllParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_chatbot.SlackChannelConfiguration",
		"metricAll",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackChannelConfiguration) AddNotificationTopic(notificationTopic awssns.ITopic) {
	if err := s.validateAddNotificationTopicParameters(notificationTopic); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addNotificationTopic",
		[]interface{}{notificationTopic},
	)
}

func (s *jsiiProxy_SlackChannelConfiguration) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := s.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (s *jsiiProxy_SlackChannelConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := s.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_SlackChannelConfiguration) BindAsNotificationRuleTarget(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleTargetConfig {
	if err := s.validateBindAsNotificationRuleTargetParameters(_scope); err != nil {
		panic(err)
	}
	var returns *awscodestarnotifications.NotificationRuleTargetConfig

	_jsii_.Invoke(
		s,
		"bindAsNotificationRuleTarget",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackChannelConfiguration) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackChannelConfiguration) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (s *jsiiProxy_SlackChannelConfiguration) GetResourceNameAttribute(nameAttr *string) *string {
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

func (s *jsiiProxy_SlackChannelConfiguration) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SlackChannelConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

