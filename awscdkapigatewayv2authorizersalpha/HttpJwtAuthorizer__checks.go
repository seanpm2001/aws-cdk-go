//go:build !no_runtime_type_checking

package awscdkapigatewayv2authorizersalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
)

func (h *jsiiProxy_HttpJwtAuthorizer) validateBindParameters(options *awscdkapigatewayv2alpha.HttpRouteAuthorizerBindOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewHttpJwtAuthorizerParameters(id *string, jwtIssuer *string, props *HttpJwtAuthorizerProps) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if jwtIssuer == nil {
		return fmt.Errorf("parameter jwtIssuer is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

