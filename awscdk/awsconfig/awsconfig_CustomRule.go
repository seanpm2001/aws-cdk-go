package awsconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsconfig/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/constructs-go/constructs/v3"
)

// A new custom rule.
//
// Example:
//   // Lambda function containing logic that evaluates compliance with the rule.
//   evalComplianceFn := lambda.NewFunction(this, jsii.String("CustomFunction"), &FunctionProps{
//   	Code: lambda.AssetCode_FromInline(jsii.String("exports.handler = (event) => console.log(event);")),
//   	Handler: jsii.String("index.handler"),
//   	Runtime: lambda.Runtime_NODEJS_14_X(),
//   })
//
//   // A custom rule that runs on configuration changes of EC2 instances
//   customRule := config.NewCustomRule(this, jsii.String("Custom"), &CustomRuleProps{
//   	ConfigurationChanges: jsii.Boolean(true),
//   	LambdaFunction: evalComplianceFn,
//   	RuleScope: config.RuleScope_FromResource(config.ResourceType_EC2_INSTANCE()),
//   })
//
//   // A rule to detect stack drifts
//   driftRule := config.NewCloudFormationStackDriftDetectionCheck(this, jsii.String("Drift"))
//
//   // Topic to which compliance notification events will be published
//   complianceTopic := sns.NewTopic(this, jsii.String("ComplianceTopic"))
//
//   // Send notification on compliance change events
//   driftRule.onComplianceChange(jsii.String("ComplianceChange"), &OnEventOptions{
//   	Target: targets.NewSnsTopic(complianceTopic),
//   })
//
// Experimental.
type CustomRule interface {
	awscdk.Resource
	IRule
	// The arn of the rule.
	// Experimental.
	ConfigRuleArn() *string
	// The compliance status of the rule.
	// Experimental.
	ConfigRuleComplianceType() *string
	// The id of the rule.
	// Experimental.
	ConfigRuleId() *string
	// The name of the rule.
	// Experimental.
	ConfigRuleName() *string
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
	// Experimental.
	IsCustomWithChanges() *bool
	// Experimental.
	SetIsCustomWithChanges(val *bool)
	// Experimental.
	IsManaged() *bool
	// Experimental.
	SetIsManaged(val *bool)
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
	// Experimental.
	RuleScope() RuleScope
	// Experimental.
	SetRuleScope(val RuleScope)
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
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
	// Defines an EventBridge event rule which triggers for rule compliance events.
	// Experimental.
	OnComplianceChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an EventBridge event rule which triggers for rule events.
	//
	// Use
	// `rule.addEventPattern(pattern)` to specify a filter.
	// Experimental.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
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
	// Defines an EventBridge event rule which triggers for rule re-evaluation status events.
	// Experimental.
	OnReEvaluationStatus(id *string, options *awsevents.OnEventOptions) awsevents.Rule
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

// The jsii proxy struct for CustomRule
type jsiiProxy_CustomRule struct {
	internal.Type__awscdkResource
	jsiiProxy_IRule
}

func (j *jsiiProxy_CustomRule) ConfigRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) ConfigRuleComplianceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleComplianceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) ConfigRuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) ConfigRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) IsCustomWithChanges() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isCustomWithChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) IsManaged() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isManaged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) RuleScope() RuleScope {
	var returns RuleScope
	_jsii_.Get(
		j,
		"ruleScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewCustomRule(scope constructs.Construct, id *string, props *CustomRuleProps) CustomRule {
	_init_.Initialize()

	if err := validateNewCustomRuleParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomRule{}

	_jsii_.Create(
		"monocdk.aws_config.CustomRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCustomRule_Override(c CustomRule, scope constructs.Construct, id *string, props *CustomRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_config.CustomRule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CustomRule)SetIsCustomWithChanges(val *bool) {
	_jsii_.Set(
		j,
		"isCustomWithChanges",
		val,
	)
}

func (j *jsiiProxy_CustomRule)SetIsManaged(val *bool) {
	_jsii_.Set(
		j,
		"isManaged",
		val,
	)
}

func (j *jsiiProxy_CustomRule)SetRuleScope(val RuleScope) {
	_jsii_.Set(
		j,
		"ruleScope",
		val,
	)
}

// Imports an existing rule.
// Experimental.
func CustomRule_FromConfigRuleName(scope constructs.Construct, id *string, configRuleName *string) IRule {
	_init_.Initialize()

	if err := validateCustomRule_FromConfigRuleNameParameters(scope, id, configRuleName); err != nil {
		panic(err)
	}
	var returns IRule

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CustomRule",
		"fromConfigRuleName",
		[]interface{}{scope, id, configRuleName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CustomRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCustomRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CustomRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func CustomRule_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCustomRule_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_config.CustomRule",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := c.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_CustomRule) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomRule) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := c.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomRule) GetResourceNameAttribute(nameAttr *string) *string {
	if err := c.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomRule) OnComplianceChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := c.validateOnComplianceChangeParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onComplianceChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomRule) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := c.validateOnEventParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomRule) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomRule) OnReEvaluationStatus(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := c.validateOnReEvaluationStatusParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onReEvaluationStatus",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomRule) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CustomRule) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomRule) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomRule) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CustomRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomRule) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

