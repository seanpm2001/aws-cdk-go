//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func validateAddressRecordTarget_FromAliasParameters(aliasTarget IAliasRecordTarget) error {
	return nil
}

