//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsneptune

// Building without runtime type checking enabled, so all the below just return nil

func validateNewEndpointParameters(address *string, port *float64) error {
	return nil
}

