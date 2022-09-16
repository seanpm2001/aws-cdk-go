//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsroute53

import (
	"fmt"
)

func validateAddressRecordTarget_FromAliasParameters(aliasTarget IAliasRecordTarget) error {
	if aliasTarget == nil {
		return fmt.Errorf("parameter aliasTarget is required, but nil was provided")
	}

	return nil
}
