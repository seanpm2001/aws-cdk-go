package awsapigatewayv2authorizers


// Properties to initialize HttpJwtAuthorizer.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   issuer := "https://test.us.auth0.com"
//   authorizer := awscdk.NewHttpJwtAuthorizer(jsii.String("BooksAuthorizer"), issuer, &httpJwtAuthorizerProps{
//   	jwtAudience: []*string{
//   		jsii.String("3131231"),
//   	},
//   })
//
//   api := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   api.addRoutes(&addRoutesOptions{
//   	integration: awscdk.NewHttpUrlIntegration(jsii.String("BooksIntegration"), jsii.String("https://get-books-proxy.myproxy.internal")),
//   	path: jsii.String("/books"),
//   	authorizer: authorizer,
//   })
//
// Experimental.
type HttpJwtAuthorizerProps struct {
	// A list of the intended recipients of the JWT.
	//
	// A valid JWT must provide an aud that matches at least one entry in this list.
	// Experimental.
	JwtAudience *[]*string `field:"required" json:"jwtAudience" yaml:"jwtAudience"`
	// The name of the authorizer.
	// Experimental.
	AuthorizerName *string `field:"optional" json:"authorizerName" yaml:"authorizerName"`
	// The identity source for which authorization is requested.
	// Experimental.
	IdentitySource *[]*string `field:"optional" json:"identitySource" yaml:"identitySource"`
}

