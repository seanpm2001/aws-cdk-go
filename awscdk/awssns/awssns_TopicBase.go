package awssns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awscodestarnotifications"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awssns/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Either a new or imported Topic.
// Experimental.
type TopicBase interface {
	awscdk.Resource
	ITopic
	// Controls automatic creation of policy objects.
	//
	// Set by subclasses.
	// Experimental.
	AutoCreatePolicy() *bool
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
	// Whether this topic is an Amazon SNS FIFO queue.
	//
	// If false, this is a standard topic.
	// Experimental.
	Fifo() *bool
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
	// The ARN of the topic.
	// Experimental.
	TopicArn() *string
	// The name of the topic.
	// Experimental.
	TopicName() *string
	// Subscribe some endpoint to this topic.
	// Experimental.
	AddSubscription(subscription ITopicSubscription)
	// Adds a statement to the IAM resource policy associated with this topic.
	//
	// If this topic was created in this stack (`new Topic`), a topic policy
	// will be automatically created upon the first call to `addToPolicy`. If
	// the topic is imported (`Topic.import`), then this is a no-op.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Represents a notification target That allows SNS topic to associate with this rule target.
	// Experimental.
	BindAsNotificationRuleTarget(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleTargetConfig
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
	// Grant topic publishing permissions to the given identity.
	// Experimental.
	GrantPublish(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this Topic.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages published to your Amazon SNS topics.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfMessagesPublished(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages successfully delivered from your Amazon SNS topics to subscribing endpoints.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfNotificationsDelivered(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that Amazon SNS failed to deliver.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfNotificationsFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that were rejected by subscription filter policies.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfNotificationsFilteredOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that were rejected by subscription filter policies because the messages' attributes are invalid.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfNotificationsFilteredOutInvalidAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that were rejected by subscription filter policies because the messages have no attributes.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricNumberOfNotificationsFilteredOutNoMessageAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the size of messages published through this topic.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricPublishSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The charges you have accrued since the start of the current calendar month for sending SMS messages.
	//
	// Maximum over 5 minutes.
	// Experimental.
	MetricSMSMonthToDateSpentUSD(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The rate of successful SMS message deliveries.
	//
	// Sum over 5 minutes.
	// Experimental.
	MetricSMSSuccessRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
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
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for TopicBase
type jsiiProxy_TopicBase struct {
	internal.Type__awscdkResource
	jsiiProxy_ITopic
}

func (j *jsiiProxy_TopicBase) AutoCreatePolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoCreatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicBase) Fifo() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"fifo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicBase) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicBase) TopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TopicBase) TopicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicName",
		&returns,
	)
	return returns
}


// Experimental.
func NewTopicBase_Override(t TopicBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_sns.TopicBase",
		[]interface{}{scope, id, props},
		t,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func TopicBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTopicBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.TopicBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func TopicBase_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateTopicBase_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_sns.TopicBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) AddSubscription(subscription ITopicSubscription) {
	if err := t.validateAddSubscriptionParameters(subscription); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addSubscription",
		[]interface{}{subscription},
	)
}

func (t *jsiiProxy_TopicBase) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
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

func (t *jsiiProxy_TopicBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := t.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (t *jsiiProxy_TopicBase) BindAsNotificationRuleTarget(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleTargetConfig {
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

func (t *jsiiProxy_TopicBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (t *jsiiProxy_TopicBase) GetResourceNameAttribute(nameAttr *string) *string {
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

func (t *jsiiProxy_TopicBase) GrantPublish(grantee awsiam.IGrantable) awsiam.Grant {
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

func (t *jsiiProxy_TopicBase) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (t *jsiiProxy_TopicBase) MetricNumberOfMessagesPublished(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (t *jsiiProxy_TopicBase) MetricNumberOfNotificationsDelivered(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (t *jsiiProxy_TopicBase) MetricNumberOfNotificationsFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (t *jsiiProxy_TopicBase) MetricNumberOfNotificationsFilteredOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (t *jsiiProxy_TopicBase) MetricNumberOfNotificationsFilteredOutInvalidAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (t *jsiiProxy_TopicBase) MetricNumberOfNotificationsFilteredOutNoMessageAttributes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (t *jsiiProxy_TopicBase) MetricPublishSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (t *jsiiProxy_TopicBase) MetricSMSMonthToDateSpentUSD(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (t *jsiiProxy_TopicBase) MetricSMSSuccessRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (t *jsiiProxy_TopicBase) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TopicBase) OnSynthesize(session constructs.ISynthesisSession) {
	if err := t.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (t *jsiiProxy_TopicBase) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) Prepare() {
	_jsii_.InvokeVoid(
		t,
		"prepare",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TopicBase) Synthesize(session awscdk.ISynthesisSession) {
	if err := t.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		[]interface{}{session},
	)
}

func (t *jsiiProxy_TopicBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicBase) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}
