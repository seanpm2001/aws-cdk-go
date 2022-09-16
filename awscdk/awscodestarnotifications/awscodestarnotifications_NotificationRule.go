package awscodestarnotifications

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodestarnotifications/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A new notification rule.
//
// Example:
//   import notifications "github.com/aws/aws-cdk-go/awscdk"
//   import codebuild "github.com/aws/aws-cdk-go/awscdk"
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//   import chatbot "github.com/aws/aws-cdk-go/awscdk"
//
//
//   project := codebuild.NewPipelineProject(this, jsii.String("MyProject"))
//
//   topic := sns.NewTopic(this, jsii.String("MyTopic1"))
//
//   slack := chatbot.NewSlackChannelConfiguration(this, jsii.String("MySlackChannel"), &slackChannelConfigurationProps{
//   	slackChannelConfigurationName: jsii.String("YOUR_CHANNEL_NAME"),
//   	slackWorkspaceId: jsii.String("YOUR_SLACK_WORKSPACE_ID"),
//   	slackChannelId: jsii.String("YOUR_SLACK_CHANNEL_ID"),
//   })
//
//   rule := notifications.NewNotificationRule(this, jsii.String("NotificationRule"), &notificationRuleProps{
//   	source: project,
//   	events: []*string{
//   		jsii.String("codebuild-project-build-state-succeeded"),
//   		jsii.String("codebuild-project-build-state-failed"),
//   	},
//   	targets: []iNotificationRuleTarget{
//   		topic,
//   	},
//   })
//   rule.addTarget(slack)
//
// Experimental.
type NotificationRule interface {
	awscdk.Resource
	INotificationRule
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
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The ARN of the notification rule (i.e. arn:aws:codestar-notifications:::notificationrule/01234abcde).
	// Experimental.
	NotificationRuleArn() *string
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
	// Adds target to notification rule.
	// Experimental.
	AddTarget(target INotificationRuleTarget) *bool
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

// The jsii proxy struct for NotificationRule
type jsiiProxy_NotificationRule struct {
	internal.Type__awscdkResource
	jsiiProxy_INotificationRule
}

func (j *jsiiProxy_NotificationRule) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationRule) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationRule) NotificationRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationRule) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotificationRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewNotificationRule(scope constructs.Construct, id *string, props *NotificationRuleProps) NotificationRule {
	_init_.Initialize()

	if err := validateNewNotificationRuleParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NotificationRule{}

	_jsii_.Create(
		"monocdk.aws_codestarnotifications.NotificationRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewNotificationRule_Override(n NotificationRule, scope constructs.Construct, id *string, props *NotificationRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_codestarnotifications.NotificationRule",
		[]interface{}{scope, id, props},
		n,
	)
}

// Import an existing notification rule provided an ARN.
// Experimental.
func NotificationRule_FromNotificationRuleArn(scope constructs.Construct, id *string, notificationRuleArn *string) INotificationRule {
	_init_.Initialize()

	if err := validateNotificationRule_FromNotificationRuleArnParameters(scope, id, notificationRuleArn); err != nil {
		panic(err)
	}
	var returns INotificationRule

	_jsii_.StaticInvoke(
		"monocdk.aws_codestarnotifications.NotificationRule",
		"fromNotificationRuleArn",
		[]interface{}{scope, id, notificationRuleArn},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func NotificationRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNotificationRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codestarnotifications.NotificationRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func NotificationRule_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateNotificationRule_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_codestarnotifications.NotificationRule",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationRule) AddTarget(target INotificationRuleTarget) *bool {
	if err := n.validateAddTargetParameters(target); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		n,
		"addTarget",
		[]interface{}{target},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := n.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (n *jsiiProxy_NotificationRule) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationRule) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := n.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationRule) GetResourceNameAttribute(nameAttr *string) *string {
	if err := n.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationRule) OnPrepare() {
	_jsii_.InvokeVoid(
		n,
		"onPrepare",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationRule) OnSynthesize(session constructs.ISynthesisSession) {
	if err := n.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (n *jsiiProxy_NotificationRule) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationRule) Prepare() {
	_jsii_.InvokeVoid(
		n,
		"prepare",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotificationRule) Synthesize(session awscdk.ISynthesisSession) {
	if err := n.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		[]interface{}{session},
	)
}

func (n *jsiiProxy_NotificationRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotificationRule) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}
