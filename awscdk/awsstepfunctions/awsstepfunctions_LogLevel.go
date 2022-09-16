package awsstepfunctions


// Defines which category of execution history events are logged.
//
// Example:
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//
//
//   logGroup := logs.NewLogGroup(this, jsii.String("MyLogGroup"))
//
//   sfn.NewStateMachine(this, jsii.String("MyStateMachine"), &stateMachineProps{
//   	definition: sfn.chain.start(sfn.NewPass(this, jsii.String("Pass"))),
//   	logs: &logOptions{
//   		destination: logGroup,
//   		level: sfn.logLevel_ALL,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html
//
// Experimental.
type LogLevel string

const (
	// No Logging.
	// Experimental.
	LogLevel_OFF LogLevel = "OFF"
	// Log everything.
	// Experimental.
	LogLevel_ALL LogLevel = "ALL"
	// Log all errors.
	// Experimental.
	LogLevel_ERROR LogLevel = "ERROR"
	// Log fatal errors.
	// Experimental.
	LogLevel_FATAL LogLevel = "FATAL"
)
