//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IDestination) validateBindParameters(scope awscdk.Construct, fn IFunction, options *DestinationOptions) error {
	return nil
}
