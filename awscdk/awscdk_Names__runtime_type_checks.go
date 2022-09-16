//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v3"
)

func validateNames_NodeUniqueIdParameters(node ConstructNode) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

func validateNames_UniqueIdParameters(construct constructs.Construct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}
