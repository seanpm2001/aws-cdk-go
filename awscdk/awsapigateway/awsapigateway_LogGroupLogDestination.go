package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awslogs"
)

// Use CloudWatch Logs as a custom access log destination for API Gateway.
//
// Example:
//   logGroup := logs.NewLogGroup(this, jsii.String("ApiGatewayAccessLogs"))
//   apigateway.NewRestApi(this, jsii.String("books"), &restApiProps{
//   	deployOptions: &stageOptions{
//   		accessLogDestination: apigateway.NewLogGroupLogDestination(logGroup),
//   		accessLogFormat: apigateway.accessLogFormat.custom(
//   		fmt.Sprintf("%v %v %v", apigateway.accessLogField.contextRequestId(), apigateway.*accessLogField.contextErrorMessage(), apigateway.*accessLogField.contextErrorMessageString())),
//   	},
//   })
//
// Experimental.
type LogGroupLogDestination interface {
	IAccessLogDestination
	// Binds this destination to the CloudWatch Logs.
	// Experimental.
	Bind(_stage IStage) *AccessLogDestinationConfig
}

// The jsii proxy struct for LogGroupLogDestination
type jsiiProxy_LogGroupLogDestination struct {
	jsiiProxy_IAccessLogDestination
}

// Experimental.
func NewLogGroupLogDestination(logGroup awslogs.ILogGroup) LogGroupLogDestination {
	_init_.Initialize()

	if err := validateNewLogGroupLogDestinationParameters(logGroup); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogGroupLogDestination{}

	_jsii_.Create(
		"monocdk.aws_apigateway.LogGroupLogDestination",
		[]interface{}{logGroup},
		&j,
	)

	return &j
}

// Experimental.
func NewLogGroupLogDestination_Override(l LogGroupLogDestination, logGroup awslogs.ILogGroup) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigateway.LogGroupLogDestination",
		[]interface{}{logGroup},
		l,
	)
}

func (l *jsiiProxy_LogGroupLogDestination) Bind(_stage IStage) *AccessLogDestinationConfig {
	if err := l.validateBindParameters(_stage); err != nil {
		panic(err)
	}
	var returns *AccessLogDestinationConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{_stage},
		&returns,
	)

	return returns
}

