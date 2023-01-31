//go:build !no_runtime_type_checking

// Authorizers for AWS APIGateway V2
package awscdkapigatewayv2authorizersalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
)

func (w *jsiiProxy_WebSocketIamAuthorizer) validateBindParameters(_options *awscdkapigatewayv2alpha.WebSocketRouteAuthorizerBindOptions) error {
	if _options == nil {
		return fmt.Errorf("parameter _options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_options, func() string { return "parameter _options" }); err != nil {
		return err
	}

	return nil
}
