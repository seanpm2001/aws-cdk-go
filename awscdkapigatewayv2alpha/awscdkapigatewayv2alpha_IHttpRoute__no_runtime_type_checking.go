//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IHttpRoute) validateGrantInvokeParameters(grantee awsiam.IGrantable, options *GrantInvokeOptions) error {
	return nil
}

