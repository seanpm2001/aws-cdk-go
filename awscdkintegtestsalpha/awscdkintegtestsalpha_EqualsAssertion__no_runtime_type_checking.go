//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// CDK Integration Testing Constructs
package awscdkintegtestsalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateEqualsAssertion_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewEqualsAssertionParameters(scope constructs.Construct, id *string, props *EqualsAssertionProps) error {
	return nil
}

