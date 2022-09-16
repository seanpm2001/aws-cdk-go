//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsredshift

import (
	"fmt"
)

func (i *jsiiProxy_IUser) validateAddTablePrivilegesParameters(table ITable) error {
	if table == nil {
		return fmt.Errorf("parameter table is required, but nil was provided")
	}

	return nil
}
