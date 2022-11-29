//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

func (i *jsiiProxy_InstanceRequireImdsv2Aspect) validateVisitParameters(node awscdk.IConstruct) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_InstanceRequireImdsv2Aspect) validateWarnParameters(node awscdk.IConstruct, message *string) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func validateNewInstanceRequireImdsv2AspectParameters(props *InstanceRequireImdsv2AspectProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

