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

// Deletes an EMR Containers virtual cluster as a Task.
//
// Example:
//   tasks.NewEmrContainersDeleteVirtualCluster(this, jsii.String("Delete a Virtual Cluster"), &emrContainersDeleteVirtualClusterProps{
//   	virtualClusterId: sfn.taskInput.fromJsonPathAt(jsii.String("$.virtualCluster")),
//   })
//
// See: https://docs.amazonaws.cn/en_us/step-functions/latest/dg/connect-emr-eks.html
//
// Experimental.
type EmrContainersDeleteVirtualCluster interface {
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

// The jsii proxy struct for EmrContainersDeleteVirtualCluster
type jsiiProxy_EmrContainersDeleteVirtualCluster struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_EmrContainersDeleteVirtualCluster) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrContainersDeleteVirtualCluster) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrContainersDeleteVirtualCluster) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrContainersDeleteVirtualCluster) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrContainersDeleteVirtualCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrContainersDeleteVirtualCluster) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrContainersDeleteVirtualCluster) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrContainersDeleteVirtualCluster) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrContainersDeleteVirtualCluster) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrContainersDeleteVirtualCluster) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrContainersDeleteVirtualCluster) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrContainersDeleteVirtualCluster) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrContainersDeleteVirtualCluster) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrContainersDeleteVirtualCluster) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrContainersDeleteVirtualCluster) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrContainersDeleteVirtualCluster) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewEmrContainersDeleteVirtualCluster(scope constructs.Construct, id *string, props *EmrContainersDeleteVirtualClusterProps) EmrContainersDeleteVirtualCluster {
	_init_.Initialize()

	if err := validateNewEmrContainersDeleteVirtualClusterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmrContainersDeleteVirtualCluster{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EmrContainersDeleteVirtualCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEmrContainersDeleteVirtualCluster_Override(e EmrContainersDeleteVirtualCluster, scope constructs.Construct, id *string, props *EmrContainersDeleteVirtualClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EmrContainersDeleteVirtualCluster",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_EmrContainersDeleteVirtualCluster)SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_EmrContainersDeleteVirtualCluster)SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func EmrContainersDeleteVirtualCluster_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	if err := validateEmrContainersDeleteVirtualCluster_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrContainersDeleteVirtualCluster",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func EmrContainersDeleteVirtualCluster_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateEmrContainersDeleteVirtualCluster_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrContainersDeleteVirtualCluster",
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
func EmrContainersDeleteVirtualCluster_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateEmrContainersDeleteVirtualCluster_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrContainersDeleteVirtualCluster",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func EmrContainersDeleteVirtualCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEmrContainersDeleteVirtualCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrContainersDeleteVirtualCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func EmrContainersDeleteVirtualCluster_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateEmrContainersDeleteVirtualCluster_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.EmrContainersDeleteVirtualCluster",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) AddBranch(branch awsstepfunctions.StateGraph) {
	if err := e.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addBranch",
		[]interface{}{branch},
	)
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
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

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	if err := e.validateAddChoiceParameters(condition, next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addChoice",
		[]interface{}{condition, next},
	)
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) AddIterator(iteration awsstepfunctions.StateGraph) {
	if err := e.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) AddPrefix(x *string) {
	if err := e.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addPrefix",
		[]interface{}{x},
	)
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
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

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) BindToGraph(graph awsstepfunctions.StateGraph) {
	if err := e.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) MakeDefault(def awsstepfunctions.State) {
	if err := e.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"makeDefault",
		[]interface{}{def},
	)
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) MakeNext(next awsstepfunctions.State) {
	if err := e.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"makeNext",
		[]interface{}{next},
	)
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
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

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) OnSynthesize(session constructs.ISynthesisSession) {
	if err := e.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) Prepare() {
	_jsii_.InvokeVoid(
		e,
		"prepare",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) Synthesize(session awscdk.ISynthesisSession) {
	if err := e.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		[]interface{}{session},
	)
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrContainersDeleteVirtualCluster) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	if err := e.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}
