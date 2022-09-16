//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsecs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

func (e *jsiiProxy_EcsOptimizedImage) validateGetImageParameters(scope awscdk.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateEcsOptimizedImage_AmazonLinuxParameters(options *EcsOptimizedImageOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateEcsOptimizedImage_AmazonLinux2Parameters(options *EcsOptimizedImageOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateEcsOptimizedImage_WindowsParameters(windowsVersion WindowsOptimizedVersion, options *EcsOptimizedImageOptions) error {
	if windowsVersion == "" {
		return fmt.Errorf("parameter windowsVersion is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}
