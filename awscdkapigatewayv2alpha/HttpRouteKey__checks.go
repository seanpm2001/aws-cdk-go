//go:build !no_runtime_type_checking

package awscdkapigatewayv2alpha

import (
	"fmt"
)

func validateHttpRouteKey_WithParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

