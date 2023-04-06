// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgameliftalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdkgameliftalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Base class for new and imported GameLift Matchmaking configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import gamelift_alpha "github.com/aws/aws-cdk-go/awscdkgameliftalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var topic topic
//
//   matchmakingConfigurationBase := gamelift_alpha.MatchmakingConfigurationBase_FromMatchmakingConfigurationAttributes(this, jsii.String("MyMatchmakingConfigurationBase"), &MatchmakingConfigurationAttributes{
//   	MatchmakingConfigurationArn: jsii.String("matchmakingConfigurationArn"),
//   	MatchmakingConfigurationName: jsii.String("matchmakingConfigurationName"),
//   	NotificationTarget: topic,
//   })
//
// Experimental.
type MatchmakingConfigurationBase interface {
	awscdk.Resource
	IMatchmakingConfiguration
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
	// The ARN of the matchmaking configuration.
	// Experimental.
	MatchmakingConfigurationArn() *string
	// The Identifier of the matchmaking configuration.
	// Experimental.
	MatchmakingConfigurationName() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The notification target for matchmaking events.
	// Experimental.
	NotificationTarget() awssns.ITopic
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
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
	// Return the given named metric for this matchmaking configuration.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Matchmaking requests currently being processed or waiting to be processed.
	// Experimental.
	MetricCurrentTickets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// For matchmaking configurations that require acceptance, the potential matches that were accepted since the last report.
	// Experimental.
	MetricMatchesAccepted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Potential matches that were created since the last report.
	// Experimental.
	MetricMatchesCreated(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Matches that were successfully placed into a game session since the last report.
	// Experimental.
	MetricMatchesPlaced(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// For matchmaking configurations that require acceptance, the potential matches that were rejected by at least one player since the last report.
	// Experimental.
	MetricMatchesRejected(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Players in matchmaking tickets that were added since the last report.
	// Experimental.
	MetricPlayersStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// For matchmaking requests that were put into a potential match before the last report, the amount of time between ticket creation and potential match creation.
	//
	// Units: seconds.
	// Experimental.
	MetricTimeToMatch(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MatchmakingConfigurationBase
type jsiiProxy_MatchmakingConfigurationBase struct {
	internal.Type__awscdkResource
	jsiiProxy_IMatchmakingConfiguration
}

func (j *jsiiProxy_MatchmakingConfigurationBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MatchmakingConfigurationBase) MatchmakingConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchmakingConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MatchmakingConfigurationBase) MatchmakingConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchmakingConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MatchmakingConfigurationBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MatchmakingConfigurationBase) NotificationTarget() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"notificationTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MatchmakingConfigurationBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MatchmakingConfigurationBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewMatchmakingConfigurationBase_Override(m MatchmakingConfigurationBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-gamelift-alpha.MatchmakingConfigurationBase",
		[]interface{}{scope, id, props},
		m,
	)
}

// Import an existing matchmaking configuration from its attributes.
// Experimental.
func MatchmakingConfigurationBase_FromMatchmakingConfigurationAttributes(scope constructs.Construct, id *string, attrs *MatchmakingConfigurationAttributes) IMatchmakingConfiguration {
	_init_.Initialize()

	if err := validateMatchmakingConfigurationBase_FromMatchmakingConfigurationAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IMatchmakingConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.MatchmakingConfigurationBase",
		"fromMatchmakingConfigurationAttributes",
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
// Experimental.
func MatchmakingConfigurationBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMatchmakingConfigurationBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.MatchmakingConfigurationBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func MatchmakingConfigurationBase_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateMatchmakingConfigurationBase_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.MatchmakingConfigurationBase",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func MatchmakingConfigurationBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateMatchmakingConfigurationBase_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.MatchmakingConfigurationBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MatchmakingConfigurationBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := m.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (m *jsiiProxy_MatchmakingConfigurationBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MatchmakingConfigurationBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (m *jsiiProxy_MatchmakingConfigurationBase) GetResourceNameAttribute(nameAttr *string) *string {
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

func (m *jsiiProxy_MatchmakingConfigurationBase) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (m *jsiiProxy_MatchmakingConfigurationBase) MetricCurrentTickets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := m.validateMetricCurrentTicketsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricCurrentTickets",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MatchmakingConfigurationBase) MetricMatchesAccepted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := m.validateMetricMatchesAcceptedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricMatchesAccepted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MatchmakingConfigurationBase) MetricMatchesCreated(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := m.validateMetricMatchesCreatedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricMatchesCreated",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MatchmakingConfigurationBase) MetricMatchesPlaced(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := m.validateMetricMatchesPlacedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricMatchesPlaced",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MatchmakingConfigurationBase) MetricMatchesRejected(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := m.validateMetricMatchesRejectedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricMatchesRejected",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MatchmakingConfigurationBase) MetricPlayersStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := m.validateMetricPlayersStartedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricPlayersStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MatchmakingConfigurationBase) MetricTimeToMatch(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := m.validateMetricTimeToMatchParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricTimeToMatch",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MatchmakingConfigurationBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

