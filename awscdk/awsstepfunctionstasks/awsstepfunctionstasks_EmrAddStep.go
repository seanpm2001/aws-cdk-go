package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Step Functions Task to add a Step to an EMR Cluster.
//
// The StepConfiguration is defined as Parameters in the state machine definition.
//
// OUTPUT: the StepId.
//
// Example:
//   tasks.NewEmrAddStep(this, jsii.String("Task"), &emrAddStepProps{
//   	clusterId: jsii.String("ClusterId"),
//   	name: jsii.String("StepName"),
//   	jar: jsii.String("Jar"),
//   	actionOnFailure: tasks.actionOnFailure_CONTINUE,
//   })
//
type EmrAddStep interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	// Continuable states of this Chainable.
	EndStates() *[]awsstepfunctions.INextable
	// Descriptive identifier for this chainable.
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	// The tree node.
	Node() constructs.Node
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	// First state of this Chainable.
	StartState() awsstepfunctions.State
	// Tokenized string that evaluates to the state's ID.
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	// Add a paralle branch to this state.
	AddBranch(branch awsstepfunctions.StateGraph)
	// Add a recovery handler for this state.
	//
	// When a particular error occurs, execution will continue at the error
	// handler instead of failing the state machine execution.
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	// Add a choice branch to this state.
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	// Add a map iterator to this state.
	AddIterator(iteration awsstepfunctions.StateGraph)
	// Add a prefix to the stateId of this state.
	AddPrefix(x *string)
	// Add retry configuration for this state.
	//
	// This controls if and how the execution will be retried if a particular
	// error occurs.
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	// Register this state as part of the given graph.
	//
	// Don't call this. It will be called automatically when you work
	// with states normally.
	BindToGraph(graph awsstepfunctions.StateGraph)
	// Make the indicated state the default choice transition of this state.
	MakeDefault(def awsstepfunctions.State)
	// Make the indicated state the default transition of this state.
	MakeNext(next awsstepfunctions.State)
	// Return the given named metric for this Task.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity fails.
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times the heartbeat times out for this activity.
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The interval, in milliseconds, between the time the Task starts and the time it closes.
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity is scheduled.
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The interval, in milliseconds, for which the activity stays in the schedule state.
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity is started.
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity succeeds.
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity times out.
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Continue normal execution with the given state.
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	// Render parallel branches in ASL JSON format.
	RenderBranches() interface{}
	// Render the choices in ASL JSON format.
	RenderChoices() interface{}
	// Render InputPath/Parameters/OutputPath in ASL JSON format.
	RenderInputOutput() interface{}
	// Render map iterator in ASL JSON format.
	RenderIterator() interface{}
	// Render the default next state in ASL JSON format.
	RenderNextEnd() interface{}
	// Render ResultSelector in ASL JSON format.
	RenderResultSelector() interface{}
	// Render error recovery options in ASL JSON format.
	RenderRetryCatch() interface{}
	// Return the Amazon States Language object for this state.
	ToStateJson() *map[string]interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Allows the state to validate itself.
	ValidateState() *[]*string
	// Called whenever this state is bound to a graph.
	//
	// Can be overridden by subclasses.
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for EmrAddStep
type jsiiProxy_EmrAddStep struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_EmrAddStep) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


func NewEmrAddStep(scope constructs.Construct, id *string, props *EmrAddStepProps) EmrAddStep {
	_init_.Initialize()

	if err := validateNewEmrAddStepParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrAddStep{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrAddStep",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewEmrAddStep_Override(e EmrAddStep, scope constructs.Construct, id *string, props *EmrAddStepProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrAddStep",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_EmrAddStep)SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_EmrAddStep)SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func EmrAddStep_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	if err := validateEmrAddStep_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrAddStep",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func EmrAddStep_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateEmrAddStep_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrAddStep",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func EmrAddStep_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateEmrAddStep_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrAddStep",
		"findReachableStates",
		[]interface{}{start, options},
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
func EmrAddStep_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmrAddStep_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrAddStep",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func EmrAddStep_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateEmrAddStep_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrAddStep",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (e *jsiiProxy_EmrAddStep) AddBranch(branch awsstepfunctions.StateGraph) {
	if err := e.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addBranch",
		[]interface{}{branch},
	)
}

func (e *jsiiProxy_EmrAddStep) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	if err := e.validateAddCatchParameters(handler, props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	if err := e.validateAddChoiceParameters(condition, next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addChoice",
		[]interface{}{condition, next},
	)
}

func (e *jsiiProxy_EmrAddStep) AddIterator(iteration awsstepfunctions.StateGraph) {
	if err := e.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (e *jsiiProxy_EmrAddStep) AddPrefix(x *string) {
	if err := e.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addPrefix",
		[]interface{}{x},
	)
}

func (e *jsiiProxy_EmrAddStep) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	if err := e.validateAddRetryParameters(props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) BindToGraph(graph awsstepfunctions.StateGraph) {
	if err := e.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (e *jsiiProxy_EmrAddStep) MakeDefault(def awsstepfunctions.State) {
	if err := e.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"makeDefault",
		[]interface{}{def},
	)
}

func (e *jsiiProxy_EmrAddStep) MakeNext(next awsstepfunctions.State) {
	if err := e.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"makeNext",
		[]interface{}{next},
	)
}

func (e *jsiiProxy_EmrAddStep) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := e.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := e.validateMetricFailedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := e.validateMetricHeartbeatTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := e.validateMetricRunTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := e.validateMetricScheduledParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := e.validateMetricScheduleTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := e.validateMetricStartedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := e.validateMetricSucceededParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := e.validateMetricTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := e.validateMetricTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	if err := e.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		e,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrAddStep) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	if err := e.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

