package awscdkapigatewayv2alpha


// Props for WebSocket API.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2authorizersalpha"
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//   // This function handles your auth logic
//   var authHandler function
//
//   // This function handles your WebSocket requests
//   var handler function
//
//
//   authorizer := awscdkapigatewayv2authorizersalpha.NewWebSocketLambdaAuthorizer(jsii.String("Authorizer"), authHandler)
//
//   integration := awscdkapigatewayv2integrationsalpha.NewWebSocketLambdaIntegration(jsii.String("Integration"), handler)
//
//   apigwv2.NewWebSocketApi(this, jsii.String("WebSocketApi"), &WebSocketApiProps{
//   	ConnectRouteOptions: &WebSocketRouteOptions{
//   		Integration: *Integration,
//   		Authorizer: *Authorizer,
//   	},
//   })
//
// Experimental.
type WebSocketApiProps struct {
	// An API key selection expression.
	//
	// Providing this option will require an API Key be provided to access the API.
	// Default: - Key is not required to access these APIs.
	//
	// Experimental.
	ApiKeySelectionExpression WebSocketApiKeySelectionExpression `field:"optional" json:"apiKeySelectionExpression" yaml:"apiKeySelectionExpression"`
	// Name for the WebSocket API resource.
	// Default: - id of the WebSocketApi construct.
	//
	// Experimental.
	ApiName *string `field:"optional" json:"apiName" yaml:"apiName"`
	// Options to configure a '$connect' route.
	// Default: - no '$connect' route configured.
	//
	// Experimental.
	ConnectRouteOptions *WebSocketRouteOptions `field:"optional" json:"connectRouteOptions" yaml:"connectRouteOptions"`
	// Options to configure a '$default' route.
	// Default: - no '$default' route configured.
	//
	// Experimental.
	DefaultRouteOptions *WebSocketRouteOptions `field:"optional" json:"defaultRouteOptions" yaml:"defaultRouteOptions"`
	// The description of the API.
	// Default: - none.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Options to configure a '$disconnect' route.
	// Default: - no '$disconnect' route configured.
	//
	// Experimental.
	DisconnectRouteOptions *WebSocketRouteOptions `field:"optional" json:"disconnectRouteOptions" yaml:"disconnectRouteOptions"`
	// The route selection expression for the API.
	// Default: '$request.body.action'
	//
	// Experimental.
	RouteSelectionExpression *string `field:"optional" json:"routeSelectionExpression" yaml:"routeSelectionExpression"`
}

