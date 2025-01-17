package awscdkapigatewayv2alpha


// Properties to initialize an instance of `HttpApi`.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//   var lb applicationLoadBalancer
//
//   listener := lb.AddListener(jsii.String("listener"), &BaseApplicationListenerProps{
//   	Port: jsii.Number(80),
//   })
//   listener.AddTargets(jsii.String("target"), &AddApplicationTargetsProps{
//   	Port: jsii.Number(80),
//   })
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &HttpApiProps{
//   	DefaultIntegration: awscdkapigatewayv2integrationsalpha.NewHttpAlbIntegration(jsii.String("DefaultIntegration"), listener, &HttpAlbIntegrationProps{
//   		ParameterMapping: apigwv2.NewParameterMapping().Custom(jsii.String("myKey"), jsii.String("myValue")),
//   	}),
//   })
//
// Experimental.
type HttpApiProps struct {
	// Name for the HTTP API resource.
	// Default: - id of the HttpApi construct.
	//
	// Experimental.
	ApiName *string `field:"optional" json:"apiName" yaml:"apiName"`
	// Specifies a CORS configuration for an API.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html
	//
	// Default: - CORS disabled.
	//
	// Experimental.
	CorsPreflight *CorsPreflightOptions `field:"optional" json:"corsPreflight" yaml:"corsPreflight"`
	// Whether a default stage and deployment should be automatically created.
	// Default: true.
	//
	// Experimental.
	CreateDefaultStage *bool `field:"optional" json:"createDefaultStage" yaml:"createDefaultStage"`
	// Default OIDC scopes attached to all routes in the gateway, unless explicitly configured on the route.
	// Default: - no default authorization scopes.
	//
	// Experimental.
	DefaultAuthorizationScopes *[]*string `field:"optional" json:"defaultAuthorizationScopes" yaml:"defaultAuthorizationScopes"`
	// Default Authorizer to applied to all routes in the gateway.
	// Default: - No authorizer.
	//
	// Experimental.
	DefaultAuthorizer IHttpRouteAuthorizer `field:"optional" json:"defaultAuthorizer" yaml:"defaultAuthorizer"`
	// Configure a custom domain with the API mapping resource to the HTTP API.
	// Default: - no default domain mapping configured. meaningless if `createDefaultStage` is `false`.
	//
	// Experimental.
	DefaultDomainMapping *DomainMappingOptions `field:"optional" json:"defaultDomainMapping" yaml:"defaultDomainMapping"`
	// An integration that will be configured on the catch-all route ($default).
	// Default: - none.
	//
	// Experimental.
	DefaultIntegration HttpRouteIntegration `field:"optional" json:"defaultIntegration" yaml:"defaultIntegration"`
	// The description of the API.
	// Default: - none.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether clients can invoke your API using the default endpoint.
	//
	// By default, clients can invoke your API with the default
	// `https://{api_id}.execute-api.{region}.amazonaws.com` endpoint. Enable
	// this if you would like clients to use your custom domain name.
	// Default: false execute-api endpoint enabled.
	//
	// Experimental.
	DisableExecuteApiEndpoint *bool `field:"optional" json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
}

