//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// The CDK Construct Library for AWS Lambda in Golang
package awscdklambdagoalpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ICommandHooks) validateAfterBundlingParameters(inputDir *string, outputDir *string) error {
	return nil
}

func (i *jsiiProxy_ICommandHooks) validateBeforeBundlingParameters(inputDir *string, outputDir *string) error {
	return nil
}

