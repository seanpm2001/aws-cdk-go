//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// CDK Integration Testing Constructs
package awscdkintegtestsalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (l *jsiiProxy_LambdaInvokeFunction) validateAssertAtPathParameters(path *string, expected ExpectedResult) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if expected == nil {
		return fmt.Errorf("parameter expected is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LambdaInvokeFunction) validateExpectParameters(expected ExpectedResult) error {
	if expected == nil {
		return fmt.Errorf("parameter expected is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LambdaInvokeFunction) validateGetAttParameters(attributeName *string) error {
	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LambdaInvokeFunction) validateGetAttStringParameters(attributeName *string) error {
	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

func validateLambdaInvokeFunction_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewLambdaInvokeFunctionParameters(scope constructs.Construct, id *string, props *LambdaInvokeFunctionProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

