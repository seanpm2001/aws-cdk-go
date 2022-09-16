//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsredshift

import (
	"fmt"
)

func (i *jsiiProxy_ITable) validateGrantParameters(user IUser) error {
	if user == nil {
		return fmt.Errorf("parameter user is required, but nil was provided")
	}

	return nil
}
