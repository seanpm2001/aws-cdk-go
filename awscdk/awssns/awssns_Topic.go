package awssns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarnotifications"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// A new SNS topic.
//
// Example:
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//
//
//   dlt := sns.NewTopic(this, jsii.String("DLQ"))
//   fn := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_18_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromInline(jsii.String("// your code here")),
//   	deadLetterTopic: dlt,
//   })
//
type Topic interface {
	TopicBase
	// Controls automatic creation of policy objects.
	//
	// Set by subclasses.
	AutoCreatePolicy() *bool
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// Whether this topic is an Amazon SNS FIFO queue.
	//
	// If false, this is a standard topic.
	Fifo() *bool
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
	// The ARN of the topic.
	TopicArn() *string
	// The name of the topic.
	TopicName() *string
	// Subscribe some endpoint to this topic.
	AddSubscription(topicSubscription ITopicSubscription) Subscription
	// Adds a statement to the IAM resource policy associated with this topic.
	//
	// If this topic was created in this stack (`new Topic`), a topic policy
	// will be automatically created upon the first call to `addToPolicy`. If
	// the topic is imported (`Topic.import`), then this is a no-op.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
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
	// Represents a notification target That allows SNS topic to associate with this rule target.
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
	// Grant topic publishing permissions to the given identity.
	GrantPublish(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this Topic.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages published to your Amazon SNS topics.
	//
	// Sum over 5 minutes.
	MetricNumberOfMessagesPublished(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages successfully delivered from your Amazon SNS topics to subscribing endpoints.
	//
	// Sum over 5 minutes.
	MetricNumberOfNotificationsDelivered(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that Amazon SNS failed to deliver.
	//
	// Sum over 5 minutes.
	MetricNumberOfNotificationsFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that were rejected by subscription filter policies.
	//
	// Sum over 5 minutes.
	MetricNumberOfNotificationsFilteredOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that were rejected by subscription filter policies because the messages' attributes are invalid.
	//
	// Sum over 5 minutes.
	MetricNumberOfNotificationsFilteredOutInvalidAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that were rejected by subscription filter policies because the messages have no attributes.
	//
	// Sum over 5 minutes.
	MetricNumberOfNotificationsFilteredOutNoMessageAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the size of messages published through this topic.
	//
	// Average over 5 minutes.
	MetricPublishSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The charges you have accrued since the start of the current calendar month for sending SMS messages.
	//
	// Maximum over 5 minutes.
	MetricSMSMonthToDateSpentUSD(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The rate of successful SMS message deliveries.
	//
	// Sum over 5 minutes.
	MetricSMSSuccessRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Topic
type jsiiProxy_Topic struct {
	jsiiProxy_TopicBase
}

func (j *jsiiProxy_Topic) AutoCreatePolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoCreatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Topic) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Topic) Fifo() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"fifo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Topic) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Topic) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Topic) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Topic) TopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Topic) TopicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicName",
		&returns,
	)
	return returns
}


func NewTopic(scope constructs.Construct, id *string, props *TopicProps) Topic {
	_init_.Initialize()

	if err := validateNewTopicParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Topic{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sns.Topic",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewTopic_Override(t Topic, scope constructs.Construct, id *string, props *TopicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sns.Topic",
		[]interface{}{scope, id, props},
		t,
	)
}

// Import an existing SNS topic provided an ARN.
func Topic_FromTopicArn(scope constructs.Construct, id *string, topicArn *string) ITopic {
	_init_.Initialize()

	if err := validateTopic_FromTopicArnParameters(scope, id, topicArn); err != nil {
		panic(err)
	}
	var returns ITopic

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sns.Topic",
		"fromTopicArn",
		[]interface{}{scope, id, topicArn},
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
func Topic_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTopic_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sns.Topic",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Topic_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateTopic_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sns.Topic",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Topic_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateTopic_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sns.Topic",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) AddSubscription(topicSubscription ITopicSubscription) Subscription {
	if err := t.validateAddSubscriptionParameters(topicSubscription); err != nil {
		panic(err)
	}
	var returns Subscription

	_jsii_.Invoke(
		t,
		"addSubscription",
		[]interface{}{topicSubscription},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	if err := t.validateAddToResourcePolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		t,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := t.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (t *jsiiProxy_Topic) BindAsNotificationRuleTarget(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleTargetConfig {
	if err := t.validateBindAsNotificationRuleTargetParameters(_scope); err != nil {
		panic(err)
	}
	var returns *awscodestarnotifications.NotificationRuleTargetConfig

	_jsii_.Invoke(
		t,
		"bindAsNotificationRuleTarget",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := t.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) GetResourceNameAttribute(nameAttr *string) *string {
	if err := t.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) GrantPublish(grantee awsiam.IGrantable) awsiam.Grant {
	if err := t.validateGrantPublishParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"grantPublish",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) MetricNumberOfMessagesPublished(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricNumberOfMessagesPublishedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricNumberOfMessagesPublished",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) MetricNumberOfNotificationsDelivered(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricNumberOfNotificationsDeliveredParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricNumberOfNotificationsDelivered",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) MetricNumberOfNotificationsFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricNumberOfNotificationsFailedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricNumberOfNotificationsFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) MetricNumberOfNotificationsFilteredOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricNumberOfNotificationsFilteredOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricNumberOfNotificationsFilteredOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) MetricNumberOfNotificationsFilteredOutInvalidAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricNumberOfNotificationsFilteredOutInvalidAttributesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricNumberOfNotificationsFilteredOutInvalidAttributes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) MetricNumberOfNotificationsFilteredOutNoMessageAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricNumberOfNotificationsFilteredOutNoMessageAttributesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricNumberOfNotificationsFilteredOutNoMessageAttributes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) MetricPublishSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricPublishSizeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricPublishSize",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) MetricSMSMonthToDateSpentUSD(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricSMSMonthToDateSpentUSDParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricSMSMonthToDateSpentUSD",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) MetricSMSSuccessRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricSMSSuccessRateParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricSMSSuccessRate",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

