package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Display the metric associated with an alarm, including the alarm line.
//
// Example:
//   var dashboard dashboard
//   var errorAlarm alarm
//
//
//   dashboard.addWidgets(cloudwatch.NewAlarmWidget(&alarmWidgetProps{
//   	title: jsii.String("Errors"),
//   	alarm: errorAlarm,
//   }))
//
// Experimental.
type AlarmWidget interface {
	ConcreteWidget
	// The amount of vertical grid units the widget will take up.
	// Experimental.
	Height() *float64
	// Any warnings that are produced as a result of putting together this widget.
	// Experimental.
	Warnings() *[]*string
	// The amount of horizontal grid units the widget will take up.
	// Experimental.
	Width() *float64
	// Experimental.
	X() *float64
	// Experimental.
	SetX(val *float64)
	// Experimental.
	Y() *float64
	// Experimental.
	SetY(val *float64)
	// Copy the warnings from the given metric.
	// Experimental.
	CopyMetricWarnings(ms ...IMetric)
	// Place the widget at a given position.
	// Experimental.
	Position(x *float64, y *float64)
	// Return the widget JSON for use in the dashboard.
	// Experimental.
	ToJson() *[]interface{}
}

// The jsii proxy struct for AlarmWidget
type jsiiProxy_AlarmWidget struct {
	jsiiProxy_ConcreteWidget
}

func (j *jsiiProxy_AlarmWidget) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlarmWidget) Warnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlarmWidget) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlarmWidget) X() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"x",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlarmWidget) Y() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"y",
		&returns,
	)
	return returns
}


// Experimental.
func NewAlarmWidget(props *AlarmWidgetProps) AlarmWidget {
	_init_.Initialize()

	if err := validateNewAlarmWidgetParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlarmWidget{}

	_jsii_.Create(
		"monocdk.aws_cloudwatch.AlarmWidget",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewAlarmWidget_Override(a AlarmWidget, props *AlarmWidgetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudwatch.AlarmWidget",
		[]interface{}{props},
		a,
	)
}

func (j *jsiiProxy_AlarmWidget)SetX(val *float64) {
	_jsii_.Set(
		j,
		"x",
		val,
	)
}

func (j *jsiiProxy_AlarmWidget)SetY(val *float64) {
	_jsii_.Set(
		j,
		"y",
		val,
	)
}

func (a *jsiiProxy_AlarmWidget) CopyMetricWarnings(ms ...IMetric) {
	args := []interface{}{}
	for _, a := range ms {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"copyMetricWarnings",
		args,
	)
}

func (a *jsiiProxy_AlarmWidget) Position(x *float64, y *float64) {
	if err := a.validatePositionParameters(x, y); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"position",
		[]interface{}{x, y},
	)
}

func (a *jsiiProxy_AlarmWidget) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		a,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}
