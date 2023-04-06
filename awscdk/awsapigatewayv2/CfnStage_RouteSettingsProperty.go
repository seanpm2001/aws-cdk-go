package awsapigatewayv2


// Represents a collection of route settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routeSettingsProperty := &RouteSettingsProperty{
//   	DataTraceEnabled: jsii.Boolean(false),
//   	DetailedMetricsEnabled: jsii.Boolean(false),
//   	LoggingLevel: jsii.String("loggingLevel"),
//   	ThrottlingBurstLimit: jsii.Number(123),
//   	ThrottlingRateLimit: jsii.Number(123),
//   }
//
type CfnStage_RouteSettingsProperty struct {
	// Specifies whether ( `true` ) or not ( `false` ) data trace logging is enabled for this route.
	//
	// This property affects the log entries pushed to Amazon CloudWatch Logs. Supported only for WebSocket APIs.
	DataTraceEnabled interface{} `field:"optional" json:"dataTraceEnabled" yaml:"dataTraceEnabled"`
	// Specifies whether detailed metrics are enabled.
	DetailedMetricsEnabled interface{} `field:"optional" json:"detailedMetricsEnabled" yaml:"detailedMetricsEnabled"`
	// Specifies the logging level for this route: `INFO` , `ERROR` , or `OFF` .
	//
	// This property affects the log entries pushed to Amazon CloudWatch Logs. Supported only for WebSocket APIs.
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// Specifies the throttling burst limit.
	ThrottlingBurstLimit *float64 `field:"optional" json:"throttlingBurstLimit" yaml:"throttlingBurstLimit"`
	// Specifies the throttling rate limit.
	ThrottlingRateLimit *float64 `field:"optional" json:"throttlingRateLimit" yaml:"throttlingRateLimit"`
}
