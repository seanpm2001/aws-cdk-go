package awsstepfunctions


// Defines which category of execution history events are logged.
//
// Example:
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//
//
//   logGroup := logs.NewLogGroup(this, jsii.String("MyLogGroup"))
//
//   sfn.NewStateMachine(this, jsii.String("MyStateMachine"), &StateMachineProps{
//   	Definition: sfn.Chain_Start(sfn.NewPass(this, jsii.String("Pass"))),
//   	Logs: &LogOptions{
//   		Destination: logGroup,
//   		Level: sfn.LogLevel_ALL,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html
//
type LogLevel string

const (
	// No Logging.
	LogLevel_OFF LogLevel = "OFF"
	// Log everything.
	LogLevel_ALL LogLevel = "ALL"
	// Log all errors.
	LogLevel_ERROR LogLevel = "ERROR"
	// Log fatal errors.
	LogLevel_FATAL LogLevel = "FATAL"
)

