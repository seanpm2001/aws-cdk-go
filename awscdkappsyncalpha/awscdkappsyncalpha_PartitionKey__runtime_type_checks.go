//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	"fmt"
)

func (p *jsiiProxy_PartitionKey) validateSortParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validatePartitionKey_PartitionParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateNewPartitionKeyParameters(pkey Assign) error {
	if pkey == nil {
		return fmt.Errorf("parameter pkey is required, but nil was provided")
	}

	return nil
}

