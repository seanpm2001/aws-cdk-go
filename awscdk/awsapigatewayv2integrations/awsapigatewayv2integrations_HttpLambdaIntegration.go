package awsapigatewayv2integrations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2integrations/internal"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
)

// The Lambda Proxy integration resource for HTTP API.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var booksDefaultFn function
//
//   booksIntegration := awscdk.NewHttpLambdaIntegration(jsii.String("BooksIntegration"), booksDefaultFn)
//
//   httpApi := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   httpApi.addRoutes(&addRoutesOptions{
//   	path: jsii.String("/books"),
//   	methods: []httpMethod{
//   		apigwv2.*httpMethod_GET,
//   	},
//   	integration: booksIntegration,
//   })
//
// Experimental.
type HttpLambdaIntegration interface {
	awsapigatewayv2.HttpRouteIntegration
	// Bind this integration to the route.
	// Experimental.
	Bind(_arg *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig
	// Complete the binding of the integration to the route.
	//
	// In some cases, there is
	// some additional work to do, such as adding permissions for the API to access
	// the target. This work is necessary whether the integration has just been
	// created for this route or it is an existing one, previously created for other
	// routes. In most cases, however, concrete implementations do not need to
	// override this method.
	// Experimental.
	CompleteBind(options *awsapigatewayv2.HttpRouteIntegrationBindOptions)
}

// The jsii proxy struct for HttpLambdaIntegration
type jsiiProxy_HttpLambdaIntegration struct {
	internal.Type__awsapigatewayv2HttpRouteIntegration
}

// Experimental.
func NewHttpLambdaIntegration(id *string, handler awslambda.IFunction, props *HttpLambdaIntegrationProps) HttpLambdaIntegration {
	_init_.Initialize()

	if err := validateNewHttpLambdaIntegrationParameters(id, handler, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpLambdaIntegration{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.HttpLambdaIntegration",
		[]interface{}{id, handler, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpLambdaIntegration_Override(h HttpLambdaIntegration, id *string, handler awslambda.IFunction, props *HttpLambdaIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.HttpLambdaIntegration",
		[]interface{}{id, handler, props},
		h,
	)
}

func (h *jsiiProxy_HttpLambdaIntegration) Bind(_arg *awsapigatewayv2.HttpRouteIntegrationBindOptions) *awsapigatewayv2.HttpRouteIntegrationConfig {
	if err := h.validateBindParameters(_arg); err != nil {
		panic(err)
	}
	var returns *awsapigatewayv2.HttpRouteIntegrationConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpLambdaIntegration) CompleteBind(options *awsapigatewayv2.HttpRouteIntegrationBindOptions) {
	if err := h.validateCompleteBindParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"completeBind",
		[]interface{}{options},
	)
}
