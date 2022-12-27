// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties to initialize an instance of `HttpAuthorizer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var httpApi httpApi
//
//   httpAuthorizerProps := &httpAuthorizerProps{
//   	httpApi: httpApi,
//   	identitySource: []*string{
//   		jsii.String("identitySource"),
//   	},
//   	type: apigatewayv2_alpha.httpAuthorizerType_IAM,
//
//   	// the properties below are optional
//   	authorizerName: jsii.String("authorizerName"),
//   	authorizerUri: jsii.String("authorizerUri"),
//   	enableSimpleResponses: jsii.Boolean(false),
//   	jwtAudience: []*string{
//   		jsii.String("jwtAudience"),
//   	},
//   	jwtIssuer: jsii.String("jwtIssuer"),
//   	payloadFormatVersion: apigatewayv2_alpha.authorizerPayloadVersion_VERSION_1_0,
//   	resultsCacheTtl: cdk.duration.minutes(jsii.Number(30)),
//   }
//
// Experimental.
type HttpAuthorizerProps struct {
	// HTTP Api to attach the authorizer to.
	// Experimental.
	HttpApi IHttpApi `field:"required" json:"httpApi" yaml:"httpApi"`
	// The identity source for which authorization is requested.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-identitysource
	//
	// Experimental.
	IdentitySource *[]*string `field:"required" json:"identitySource" yaml:"identitySource"`
	// The type of authorizer.
	// Experimental.
	Type HttpAuthorizerType `field:"required" json:"type" yaml:"type"`
	// Name of the authorizer.
	// Experimental.
	AuthorizerName *string `field:"optional" json:"authorizerName" yaml:"authorizerName"`
	// The authorizer's Uniform Resource Identifier (URI).
	//
	// For REQUEST authorizers, this must be a well-formed Lambda function URI.
	// Experimental.
	AuthorizerUri *string `field:"optional" json:"authorizerUri" yaml:"authorizerUri"`
	// Specifies whether a Lambda authorizer returns a response in a simple format.
	//
	// If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy.
	// Experimental.
	EnableSimpleResponses *bool `field:"optional" json:"enableSimpleResponses" yaml:"enableSimpleResponses"`
	// A list of the intended recipients of the JWT.
	//
	// A valid JWT must provide an aud that matches at least one entry in this list.
	// Experimental.
	JwtAudience *[]*string `field:"optional" json:"jwtAudience" yaml:"jwtAudience"`
	// The base domain of the identity provider that issues JWT.
	// Experimental.
	JwtIssuer *string `field:"optional" json:"jwtIssuer" yaml:"jwtIssuer"`
	// Specifies the format of the payload sent to an HTTP API Lambda authorizer.
	// Experimental.
	PayloadFormatVersion AuthorizerPayloadVersion `field:"optional" json:"payloadFormatVersion" yaml:"payloadFormatVersion"`
	// How long APIGateway should cache the results.
	//
	// Max 1 hour.
	// Experimental.
	ResultsCacheTtl awscdk.Duration `field:"optional" json:"resultsCacheTtl" yaml:"resultsCacheTtl"`
}

