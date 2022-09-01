package awssqs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/constructs-go/constructs/v10"
)

// A new Amazon SQS queue.
//
// Example:
//   // An sqs queue for unsuccessful invocations of a lambda function
//   import sqs "github.com/aws/aws-cdk-go/awscdk"
//
//
//   deadLetterQueue := sqs.NewQueue(this, jsii.String("DeadLetterQueue"))
//
//   myFn := lambda.NewFunction(this, jsii.String("Fn"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromInline(jsii.String("// your code")),
//   	// sqs queue for unsuccessful invocations
//   	onFailure: destinations.NewSqsDestination(deadLetterQueue),
//   })
//
type Queue interface {
	QueueBase
	// Controls automatic creation of policy objects.
	//
	// Set by subclasses.
	AutoCreatePolicy() *bool
	// If this queue is configured with a dead-letter queue, this is the dead-letter queue settings.
	DeadLetterQueue() *DeadLetterQueue
	// If this queue is encrypted, this is the KMS key.
	EncryptionMasterKey() awskms.IKey
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// Whether this queue is an Amazon SQS FIFO queue.
	//
	// If false, this is a standard queue.
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
	// The ARN of this queue.
	QueueArn() *string
	// The name of this queue.
	QueueName() *string
	// The URL of this queue.
	QueueUrl() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Adds a statement to the IAM resource policy associated with this queue.
	//
	// If this queue was created in this stack (`new Queue`), a queue policy
	// will be automatically created upon the first call to `addToPolicy`. If
	// the queue is imported (`Queue.import`), then this is a no-op.
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
	// Grant the actions defined in queueActions to the identity Principal given on this SQS queue resource.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant permissions to consume messages from a queue.
	//
	// This will grant the following permissions:
	//
	//    - sqs:ChangeMessageVisibility
	//    - sqs:DeleteMessage
	//    - sqs:ReceiveMessage
	//    - sqs:GetQueueAttributes
	// - sqs:GetQueueUrl.
	GrantConsumeMessages(grantee awsiam.IGrantable) awsiam.Grant
	// Grant an IAM principal permissions to purge all messages from the queue.
	//
	// This will grant the following permissions:
	//
	//   - sqs:PurgeQueue
	//   - sqs:GetQueueAttributes
	// - sqs:GetQueueUrl.
	GrantPurge(grantee awsiam.IGrantable) awsiam.Grant
	// Grant access to send messages to a queue to the given identity.
	//
	// This will grant the following permissions:
	//
	//   - sqs:SendMessage
	//   - sqs:GetQueueAttributes
	// - sqs:GetQueueUrl.
	GrantSendMessages(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this Queue.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The approximate age of the oldest non-deleted message in the queue.
	//
	// Maximum over 5 minutes.
	MetricApproximateAgeOfOldestMessage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages in the queue that are delayed and not available for reading immediately.
	//
	// Maximum over 5 minutes.
	MetricApproximateNumberOfMessagesDelayed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages that are in flight.
	//
	// Maximum over 5 minutes.
	MetricApproximateNumberOfMessagesNotVisible(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages available for retrieval from the queue.
	//
	// Maximum over 5 minutes.
	MetricApproximateNumberOfMessagesVisible(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of ReceiveMessage API calls that did not return a message.
	//
	// Sum over 5 minutes.
	MetricNumberOfEmptyReceives(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages deleted from the queue.
	//
	// Sum over 5 minutes.
	MetricNumberOfMessagesDeleted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages returned by calls to the ReceiveMessage action.
	//
	// Sum over 5 minutes.
	MetricNumberOfMessagesReceived(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of messages added to a queue.
	//
	// Sum over 5 minutes.
	MetricNumberOfMessagesSent(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The size of messages added to a queue.
	//
	// Average over 5 minutes.
	MetricSentMessageSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Queue
type jsiiProxy_Queue struct {
	jsiiProxy_QueueBase
}

func (j *jsiiProxy_Queue) AutoCreatePolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoCreatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Queue) DeadLetterQueue() *DeadLetterQueue {
	var returns *DeadLetterQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Queue) EncryptionMasterKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionMasterKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Queue) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Queue) Fifo() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"fifo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Queue) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Queue) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Queue) QueueArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Queue) QueueName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Queue) QueueUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Queue) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewQueue(scope constructs.Construct, id *string, props *QueueProps) Queue {
	_init_.Initialize()

	if err := validateNewQueueParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Queue{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sqs.Queue",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewQueue_Override(q Queue, scope constructs.Construct, id *string, props *QueueProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sqs.Queue",
		[]interface{}{scope, id, props},
		q,
	)
}

// Import an existing SQS queue provided an ARN.
func Queue_FromQueueArn(scope constructs.Construct, id *string, queueArn *string) IQueue {
	_init_.Initialize()

	if err := validateQueue_FromQueueArnParameters(scope, id, queueArn); err != nil {
		panic(err)
	}
	var returns IQueue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sqs.Queue",
		"fromQueueArn",
		[]interface{}{scope, id, queueArn},
		&returns,
	)

	return returns
}

// Import an existing queue.
func Queue_FromQueueAttributes(scope constructs.Construct, id *string, attrs *QueueAttributes) IQueue {
	_init_.Initialize()

	if err := validateQueue_FromQueueAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IQueue

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sqs.Queue",
		"fromQueueAttributes",
		[]interface{}{scope, id, attrs},
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
func Queue_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQueue_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sqs.Queue",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Queue_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateQueue_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sqs.Queue",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Queue_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateQueue_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sqs.Queue",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_Queue) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	if err := q.validateAddToResourcePolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		q,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_Queue) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := q.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (q *jsiiProxy_Queue) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_Queue) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := q.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_Queue) GetResourceNameAttribute(nameAttr *string) *string {
	if err := q.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_Queue) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := q.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		q,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (q *jsiiProxy_Queue) GrantConsumeMessages(grantee awsiam.IGrantable) awsiam.Grant {
	if err := q.validateGrantConsumeMessagesParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		q,
		"grantConsumeMessages",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_Queue) GrantPurge(grantee awsiam.IGrantable) awsiam.Grant {
	if err := q.validateGrantPurgeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		q,
		"grantPurge",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_Queue) GrantSendMessages(grantee awsiam.IGrantable) awsiam.Grant {
	if err := q.validateGrantSendMessagesParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		q,
		"grantSendMessages",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_Queue) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_Queue) MetricApproximateAgeOfOldestMessage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricApproximateAgeOfOldestMessageParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricApproximateAgeOfOldestMessage",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_Queue) MetricApproximateNumberOfMessagesDelayed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricApproximateNumberOfMessagesDelayedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricApproximateNumberOfMessagesDelayed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_Queue) MetricApproximateNumberOfMessagesNotVisible(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricApproximateNumberOfMessagesNotVisibleParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricApproximateNumberOfMessagesNotVisible",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_Queue) MetricApproximateNumberOfMessagesVisible(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricApproximateNumberOfMessagesVisibleParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricApproximateNumberOfMessagesVisible",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_Queue) MetricNumberOfEmptyReceives(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricNumberOfEmptyReceivesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricNumberOfEmptyReceives",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_Queue) MetricNumberOfMessagesDeleted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricNumberOfMessagesDeletedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricNumberOfMessagesDeleted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_Queue) MetricNumberOfMessagesReceived(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricNumberOfMessagesReceivedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricNumberOfMessagesReceived",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_Queue) MetricNumberOfMessagesSent(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricNumberOfMessagesSentParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricNumberOfMessagesSent",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_Queue) MetricSentMessageSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := q.validateMetricSentMessageSizeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricSentMessageSize",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_Queue) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

