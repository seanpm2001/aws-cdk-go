package awsapigatewayv2


// Options for the Route with Integration resource.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // This function handles your auth logic
//   var authHandler function
//
//
//   authorizer := awscdk.NewHttpLambdaAuthorizer(jsii.String("BooksAuthorizer"), authHandler, &HttpLambdaAuthorizerProps{
//   	ResponseTypes: []httpLambdaResponseType{
//   		awscdk.HttpLambdaResponseType_SIMPLE,
//   	},
//   })
//
//   api := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   api.AddRoutes(&AddRoutesOptions{
//   	Integration: awscdk.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
//   	Path: jsii.String("/books"),
//   	Authorizer: Authorizer,
//   })
//
// Experimental.
type AddRoutesOptions struct {
	// The integration to be configured on this route.
	// Experimental.
	Integration HttpRouteIntegration `field:"required" json:"integration" yaml:"integration"`
	// The path at which all of these routes are configured.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// The list of OIDC scopes to include in the authorization.
	//
	// These scopes will override the default authorization scopes on the gateway.
	// Set to [] to remove default scopes.
	// Experimental.
	AuthorizationScopes *[]*string `field:"optional" json:"authorizationScopes" yaml:"authorizationScopes"`
	// Authorizer to be associated to these routes.
	//
	// Use NoneAuthorizer to remove the default authorizer for the api.
	// Experimental.
	Authorizer IHttpRouteAuthorizer `field:"optional" json:"authorizer" yaml:"authorizer"`
	// The HTTP methods to be configured.
	// Experimental.
	Methods *[]HttpMethod `field:"optional" json:"methods" yaml:"methods"`
}

