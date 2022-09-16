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

// A StepFunctions task to call DynamoDeleteItem.
//
// Example:
//   var myTable table
//
//   tasks.NewDynamoDeleteItem(this, jsii.String("DeleteItem"), &dynamoDeleteItemProps{
//   	key: map[string]dynamoAttributeValue{
//   		"MessageId": tasks.*dynamoAttributeValue.fromString(jsii.String("message-007")),
//   	},
//   	table: myTable,
//   	resultPath: sfn.jsonPath_DISCARD(),
//   })
//
// Experimental.
type DynamoDeleteItem interface {
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

// The jsii proxy struct for DynamoDeleteItem
type jsiiProxy_DynamoDeleteItem struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_DynamoDeleteItem) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewDynamoDeleteItem(scope constructs.Construct, id *string, props *DynamoDeleteItemProps) DynamoDeleteItem {
	_init_.Initialize()

	if err := validateNewDynamoDeleteItemParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamoDeleteItem{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.DynamoDeleteItem",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDynamoDeleteItem_Override(d DynamoDeleteItem, scope constructs.Construct, id *string, props *DynamoDeleteItemProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.DynamoDeleteItem",
		[]interface{}{scope, id, props},
		d,
	)
}

func (j *jsiiProxy_DynamoDeleteItem)SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_DynamoDeleteItem)SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func DynamoDeleteItem_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	if err := validateDynamoDeleteItem_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoDeleteItem",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func DynamoDeleteItem_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateDynamoDeleteItem_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoDeleteItem",
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
func DynamoDeleteItem_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateDynamoDeleteItem_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoDeleteItem",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func DynamoDeleteItem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamoDeleteItem_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoDeleteItem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func DynamoDeleteItem_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateDynamoDeleteItem_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.DynamoDeleteItem",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (d *jsiiProxy_DynamoDeleteItem) AddBranch(branch awsstepfunctions.StateGraph) {
	if err := d.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addBranch",
		[]interface{}{branch},
	)
}

func (d *jsiiProxy_DynamoDeleteItem) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	if err := d.validateAddCatchParameters(handler, props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		d,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	if err := d.validateAddChoiceParameters(condition, next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addChoice",
		[]interface{}{condition, next},
	)
}

func (d *jsiiProxy_DynamoDeleteItem) AddIterator(iteration awsstepfunctions.StateGraph) {
	if err := d.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (d *jsiiProxy_DynamoDeleteItem) AddPrefix(x *string) {
	if err := d.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addPrefix",
		[]interface{}{x},
	)
}

func (d *jsiiProxy_DynamoDeleteItem) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	if err := d.validateAddRetryParameters(props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		d,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) BindToGraph(graph awsstepfunctions.StateGraph) {
	if err := d.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (d *jsiiProxy_DynamoDeleteItem) MakeDefault(def awsstepfunctions.State) {
	if err := d.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"makeDefault",
		[]interface{}{def},
	)
}

func (d *jsiiProxy_DynamoDeleteItem) MakeNext(next awsstepfunctions.State) {
	if err := d.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"makeNext",
		[]interface{}{next},
	)
}

func (d *jsiiProxy_DynamoDeleteItem) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricFailedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricHeartbeatTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricRunTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricScheduledParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricScheduleTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricStartedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricSucceededParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	if err := d.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		d,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamoDeleteItem) OnSynthesize(session constructs.ISynthesisSession) {
	if err := d.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (d *jsiiProxy_DynamoDeleteItem) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) Prepare() {
	_jsii_.InvokeVoid(
		d,
		"prepare",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamoDeleteItem) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) Synthesize(session awscdk.ISynthesisSession) {
	if err := d.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		[]interface{}{session},
	)
}

func (d *jsiiProxy_DynamoDeleteItem) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoDeleteItem) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	if err := d.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}
