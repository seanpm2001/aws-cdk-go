package awscdkapigatewayv2alpha


// Properties to initialize an instance of `WebSocketStage`.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//   var messageHandler function
//
//
//   webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"))
//   apigwv2.NewWebSocketStage(this, jsii.String("mystage"), &WebSocketStageProps{
//   	WebSocketApi: WebSocketApi,
//   	StageName: jsii.String("dev"),
//   	AutoDeploy: jsii.Boolean(true),
//   })
//   webSocketApi.AddRoute(jsii.String("sendmessage"), &WebSocketRouteOptions{
//   	Integration: awscdkapigatewayv2integrationsalpha.NewWebSocketLambdaIntegration(jsii.String("SendMessageIntegration"), messageHandler),
//   })
//
// Experimental.
type WebSocketStageProps struct {
	// Whether updates to an API automatically trigger a new deployment.
	// Default: false.
	//
	// Experimental.
	AutoDeploy *bool `field:"optional" json:"autoDeploy" yaml:"autoDeploy"`
	// The options for custom domain and api mapping.
	// Default: - no custom domain and api mapping configuration.
	//
	// Experimental.
	DomainMapping *DomainMappingOptions `field:"optional" json:"domainMapping" yaml:"domainMapping"`
	// Throttle settings for the routes of this stage.
	// Default: - no throttling configuration.
	//
	// Experimental.
	Throttle *ThrottleSettings `field:"optional" json:"throttle" yaml:"throttle"`
	// The name of the stage.
	// Experimental.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
	// The WebSocket API to which this stage is associated.
	// Experimental.
	WebSocketApi IWebSocketApi `field:"required" json:"webSocketApi" yaml:"webSocketApi"`
}

