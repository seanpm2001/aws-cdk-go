// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgameliftalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new rule set for FlexMatch matchmaking.
//
// The rule set determines the two key elements of a match: your game's team structure and size, and how to group players together for the best possible match.
//
// For example, a rule set might describe a match like this:
//   - Create a match with two teams of five players each, one team is the defenders and the other team the invaders.
//   - A team can have novice and experienced players, but the average skill of the two teams must be within 10 points of each other.
//   - If no match is made after 30 seconds, gradually relax the skill requirements.
//
// Rule sets must be defined in the same Region as the matchmaking configuration they are used with.
//
// Example:
//   gamelift.NewMatchmakingRuleSet(this, jsii.String("RuleSet"), &MatchmakingRuleSetProps{
//   	MatchmakingRuleSetName: jsii.String("my-test-ruleset"),
//   	Content: gamelift.RuleSetContent_FromJsonFile(path.join(__dirname, jsii.String("my-ruleset/ruleset.json"))),
//   })
//
// See: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-rulesets.html
//
// Experimental.
type MatchmakingRuleSet interface {
	MatchmakingRuleSetBase
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
	// The ARN of the ruleSet.
	// Experimental.
	MatchmakingRuleSetArn() *string
	// The unique name of the ruleSet.
	// Experimental.
	MatchmakingRuleSetName() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
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
	// Return the given named metric for this matchmaking ruleSet.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Rule evaluations during matchmaking that failed since the last report.
	//
	// This metric is limited to the top 50 rules.
	// Experimental.
	MetricRuleEvaluationsFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Rule evaluations during the matchmaking process that passed since the last report.
	//
	// This metric is limited to the top 50 rules.
	// Experimental.
	MetricRuleEvaluationsPassed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MatchmakingRuleSet
type jsiiProxy_MatchmakingRuleSet struct {
	jsiiProxy_MatchmakingRuleSetBase
}

func (j *jsiiProxy_MatchmakingRuleSet) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MatchmakingRuleSet) MatchmakingRuleSetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchmakingRuleSetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MatchmakingRuleSet) MatchmakingRuleSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchmakingRuleSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MatchmakingRuleSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MatchmakingRuleSet) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MatchmakingRuleSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewMatchmakingRuleSet(scope constructs.Construct, id *string, props *MatchmakingRuleSetProps) MatchmakingRuleSet {
	_init_.Initialize()

	if err := validateNewMatchmakingRuleSetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_MatchmakingRuleSet{}

	_jsii_.Create(
		"@aws-cdk/aws-gamelift-alpha.MatchmakingRuleSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewMatchmakingRuleSet_Override(m MatchmakingRuleSet, scope constructs.Construct, id *string, props *MatchmakingRuleSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-gamelift-alpha.MatchmakingRuleSet",
		[]interface{}{scope, id, props},
		m,
	)
}

// Import a ruleSet into CDK using its ARN.
// Experimental.
func MatchmakingRuleSet_FromMatchmakingRuleSetArn(scope constructs.Construct, id *string, matchmakingRuleSetArn *string) IMatchmakingRuleSet {
	_init_.Initialize()

	if err := validateMatchmakingRuleSet_FromMatchmakingRuleSetArnParameters(scope, id, matchmakingRuleSetArn); err != nil {
		panic(err)
	}
	var returns IMatchmakingRuleSet

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.MatchmakingRuleSet",
		"fromMatchmakingRuleSetArn",
		[]interface{}{scope, id, matchmakingRuleSetArn},
		&returns,
	)

	return returns
}

// Import an existing matchmaking ruleSet from its attributes.
// Experimental.
func MatchmakingRuleSet_FromMatchmakingRuleSetAttributes(scope constructs.Construct, id *string, attrs *MatchmakingRuleSetAttributes) IMatchmakingRuleSet {
	_init_.Initialize()

	if err := validateMatchmakingRuleSet_FromMatchmakingRuleSetAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IMatchmakingRuleSet

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.MatchmakingRuleSet",
		"fromMatchmakingRuleSetAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import a ruleSet into CDK using its name.
// Experimental.
func MatchmakingRuleSet_FromMatchmakingRuleSetName(scope constructs.Construct, id *string, matchmakingRuleSetName *string) IMatchmakingRuleSet {
	_init_.Initialize()

	if err := validateMatchmakingRuleSet_FromMatchmakingRuleSetNameParameters(scope, id, matchmakingRuleSetName); err != nil {
		panic(err)
	}
	var returns IMatchmakingRuleSet

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.MatchmakingRuleSet",
		"fromMatchmakingRuleSetName",
		[]interface{}{scope, id, matchmakingRuleSetName},
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
func MatchmakingRuleSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMatchmakingRuleSet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.MatchmakingRuleSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func MatchmakingRuleSet_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateMatchmakingRuleSet_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.MatchmakingRuleSet",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func MatchmakingRuleSet_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateMatchmakingRuleSet_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.MatchmakingRuleSet",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MatchmakingRuleSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := m.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (m *jsiiProxy_MatchmakingRuleSet) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MatchmakingRuleSet) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := m.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MatchmakingRuleSet) GetResourceNameAttribute(nameAttr *string) *string {
	if err := m.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MatchmakingRuleSet) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := m.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MatchmakingRuleSet) MetricRuleEvaluationsFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := m.validateMetricRuleEvaluationsFailedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricRuleEvaluationsFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MatchmakingRuleSet) MetricRuleEvaluationsPassed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := m.validateMetricRuleEvaluationsPassedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricRuleEvaluationsPassed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MatchmakingRuleSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

