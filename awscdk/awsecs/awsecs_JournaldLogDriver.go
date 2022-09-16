package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// A log driver that sends log information to journald Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   journaldLogDriver := awscdk.Aws_ecs.NewJournaldLogDriver(&journaldLogDriverProps{
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	tag: jsii.String("tag"),
//   })
//
// Experimental.
type JournaldLogDriver interface {
	LogDriver
	// Called when the log driver is configured on a container.
	// Experimental.
	Bind(_scope awscdk.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for JournaldLogDriver
type jsiiProxy_JournaldLogDriver struct {
	jsiiProxy_LogDriver
}

// Constructs a new instance of the JournaldLogDriver class.
// Experimental.
func NewJournaldLogDriver(props *JournaldLogDriverProps) JournaldLogDriver {
	_init_.Initialize()

	if err := validateNewJournaldLogDriverParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_JournaldLogDriver{}

	_jsii_.Create(
		"monocdk.aws_ecs.JournaldLogDriver",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the JournaldLogDriver class.
// Experimental.
func NewJournaldLogDriver_Override(j JournaldLogDriver, props *JournaldLogDriverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs.JournaldLogDriver",
		[]interface{}{props},
		j,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
// Experimental.
func JournaldLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	if err := validateJournaldLogDriver_AwsLogsParameters(props); err != nil {
		panic(err)
	}
	var returns LogDriver

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.JournaldLogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JournaldLogDriver) Bind(_scope awscdk.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig {
	if err := j.validateBindParameters(_scope, _containerDefinition); err != nil {
		panic(err)
	}
	var returns *LogDriverConfig

	_jsii_.Invoke(
		j,
		"bind",
		[]interface{}{_scope, _containerDefinition},
		&returns,
	)

	return returns
}
