package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctionstasks/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Start a Job run as a Task.
//
// Example:
//   tasks.NewGlueDataBrewStartJobRun(this, jsii.String("Task"), &glueDataBrewStartJobRunProps{
//   	name: jsii.String("databrew-job"),
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-databrew.html
//
// Experimental.
type GlueDataBrewStartJobRun interface {
	awsstepfunctions.TaskStateBase
	// Experimental.
	Branches() *[]awsstepfunctions.StateGraph
	// Experimental.
	Comment() *string
	// Experimental.
	DefaultChoice() awsstepfunctions.State
	// Experimental.
	SetDefaultChoice(val awsstepfunctions.State)
	// Continuable states of this Chainable.
	// Experimental.
	EndStates() *[]awsstepfunctions.INextable
	// Descriptive identifier for this chainable.
	// Experimental.
	Id() *string
	// Experimental.
	InputPath() *string
	// Experimental.
	Iteration() awsstepfunctions.StateGraph
	// Experimental.
	SetIteration(val awsstepfunctions.StateGraph)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Experimental.
	OutputPath() *string
	// Experimental.
	Parameters() *map[string]interface{}
	// Experimental.
	ResultPath() *string
	// Experimental.
	ResultSelector() *map[string]interface{}
	// First state of this Chainable.
	// Experimental.
	StartState() awsstepfunctions.State
	// Tokenized string that evaluates to the state's ID.
	// Experimental.
	StateId() *string
	// Experimental.
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	// Experimental.
	TaskPolicies() *[]awsiam.PolicyStatement
	// Add a paralle branch to this state.
	// Experimental.
	AddBranch(branch awsstepfunctions.StateGraph)
	// Add a recovery handler for this state.
	//
	// When a particular error occurs, execution will continue at the error
	// handler instead of failing the state machine execution.
	// Experimental.
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	// Add a choice branch to this state.
	// Experimental.
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	// Add a map iterator to this state.
	// Experimental.
	AddIterator(iteration awsstepfunctions.StateGraph)
	// Add a prefix to the stateId of this state.
	// Experimental.
	AddPrefix(x *string)
	// Add retry configuration for this state.
	//
	// This controls if and how the execution will be retried if a particular
	// error occurs.
	// Experimental.
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	// Register this state as part of the given graph.
	//
	// Don't call this. It will be called automatically when you work
	// with states normally.
	// Experimental.
	BindToGraph(graph awsstepfunctions.StateGraph)
	// Make the indicated state the default choice transition of this state.
	// Experimental.
	MakeDefault(def awsstepfunctions.State)
	// Make the indicated state the default transition of this state.
	// Experimental.
	MakeNext(next awsstepfunctions.State)
	// Return the given named metric for this Task.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity fails.
	// Experimental.
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times the heartbeat times out for this activity.
	// Experimental.
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The interval, in milliseconds, between the time the Task starts and the time it closes.
	// Experimental.
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity is scheduled.
	// Experimental.
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The interval, in milliseconds, for which the activity stays in the schedule state.
	// Experimental.
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity is started.
	// Experimental.
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity succeeds.
	// Experimental.
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
	// Experimental.
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity times out.
	// Experimental.
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Continue normal execution with the given state.
	// Experimental.
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
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
	// Render parallel branches in ASL JSON format.
	// Experimental.
	RenderBranches() interface{}
	// Render the choices in ASL JSON format.
	// Experimental.
	RenderChoices() interface{}
	// Render InputPath/Parameters/OutputPath in ASL JSON format.
	// Experimental.
	RenderInputOutput() interface{}
	// Render map iterator in ASL JSON format.
	// Experimental.
	RenderIterator() interface{}
	// Render the default next state in ASL JSON format.
	// Experimental.
	RenderNextEnd() interface{}
	// Render ResultSelector in ASL JSON format.
	// Experimental.
	RenderResultSelector() interface{}
	// Render error recovery options in ASL JSON format.
	// Experimental.
	RenderRetryCatch() interface{}
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Return the Amazon States Language object for this state.
	// Experimental.
	ToStateJson() *map[string]interface{}
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
	// Called whenever this state is bound to a graph.
	//
	// Can be overridden by subclasses.
	// Experimental.
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for GlueDataBrewStartJobRun
type jsiiProxy_GlueDataBrewStartJobRun struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewGlueDataBrewStartJobRun(scope constructs.Construct, id *string, props *GlueDataBrewStartJobRunProps) GlueDataBrewStartJobRun {
	_init_.Initialize()

	if err := validateNewGlueDataBrewStartJobRunParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueDataBrewStartJobRun{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewGlueDataBrewStartJobRun_Override(g GlueDataBrewStartJobRun, scope constructs.Construct, id *string, props *GlueDataBrewStartJobRunProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
		[]interface{}{scope, id, props},
		g,
	)
}

func (j *jsiiProxy_GlueDataBrewStartJobRun)SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_GlueDataBrewStartJobRun)SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func GlueDataBrewStartJobRun_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	if err := validateGlueDataBrewStartJobRun_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func GlueDataBrewStartJobRun_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateGlueDataBrewStartJobRun_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
// Experimental.
func GlueDataBrewStartJobRun_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateGlueDataBrewStartJobRun_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func GlueDataBrewStartJobRun_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGlueDataBrewStartJobRun_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func GlueDataBrewStartJobRun_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateGlueDataBrewStartJobRun_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) AddBranch(branch awsstepfunctions.StateGraph) {
	if err := g.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addBranch",
		[]interface{}{branch},
	)
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	if err := g.validateAddCatchParameters(handler, props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		g,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	if err := g.validateAddChoiceParameters(condition, next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addChoice",
		[]interface{}{condition, next},
	)
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) AddIterator(iteration awsstepfunctions.StateGraph) {
	if err := g.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) AddPrefix(x *string) {
	if err := g.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addPrefix",
		[]interface{}{x},
	)
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	if err := g.validateAddRetryParameters(props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		g,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) BindToGraph(graph awsstepfunctions.StateGraph) {
	if err := g.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) MakeDefault(def awsstepfunctions.State) {
	if err := g.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"makeDefault",
		[]interface{}{def},
	)
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) MakeNext(next awsstepfunctions.State) {
	if err := g.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"makeNext",
		[]interface{}{next},
	)
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricFailedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricHeartbeatTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricRunTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricScheduledParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricScheduleTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricStartedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricSucceededParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	if err := g.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		g,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) OnPrepare() {
	_jsii_.InvokeVoid(
		g,
		"onPrepare",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) OnSynthesize(session constructs.ISynthesisSession) {
	if err := g.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) Prepare() {
	_jsii_.InvokeVoid(
		g,
		"prepare",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) Synthesize(session awscdk.ISynthesisSession) {
	if err := g.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"synthesize",
		[]interface{}{session},
	)
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	if err := g.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}
