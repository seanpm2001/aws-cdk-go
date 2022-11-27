// Authorizers for AWS APIGateway V2
package awscdkapigatewayv2authorizersalpha


// Properties to initialize WebSocketTokenAuthorizer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_authorizers_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2authorizersalpha"
//
//   webSocketLambdaAuthorizerProps := &webSocketLambdaAuthorizerProps{
//   	authorizerName: jsii.String("authorizerName"),
//   	identitySource: []*string{
//   		jsii.String("identitySource"),
//   	},
//   }
//
// Experimental.
type WebSocketLambdaAuthorizerProps struct {
	// The name of the authorizer.
	// Experimental.
	AuthorizerName *string `field:"optional" json:"authorizerName" yaml:"authorizerName"`
	// The identity source for which authorization is requested.
	//
	// Request parameter match `'route.request.querystring|header.[a-zA-z0-9._-]+'`.
	// Staged variable match `'stageVariables.[a-zA-Z0-9._-]+'`.
	// Context parameter match `'context.[a-zA-Z0-9._-]+'`.
	// Experimental.
	IdentitySource *[]*string `field:"optional" json:"identitySource" yaml:"identitySource"`
}
