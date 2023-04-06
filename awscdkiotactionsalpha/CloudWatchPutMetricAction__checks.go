//go:build !no_runtime_type_checking

// Receipt rule actions for AWS IoT
package awscdkiotactionsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateNewCloudWatchPutMetricActionParameters(props *CloudWatchPutMetricActionProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}
