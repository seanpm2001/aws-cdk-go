//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsrds

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/constructs-go/constructs/v3"
)

func (s *jsiiProxy_ServerlessCluster) validateAddRotationMultiUserParameters(id *string, options *RotationMultiUserOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_ServerlessCluster) validateAddRotationSingleUserParameters(options *RotationSingleUserOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_ServerlessCluster) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ServerlessCluster) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	if arnAttr == nil {
		return fmt.Errorf("parameter arnAttr is required, but nil was provided")
	}

	if arnComponents == nil {
		return fmt.Errorf("parameter arnComponents is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(arnComponents, func() string { return "parameter arnComponents" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_ServerlessCluster) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ServerlessCluster) validateGrantDataApiAccessParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ServerlessCluster) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ServerlessCluster) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func validateServerlessCluster_FromServerlessClusterAttributesParameters(scope constructs.Construct, id *string, attrs *ServerlessClusterAttributes) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if attrs == nil {
		return fmt.Errorf("parameter attrs is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(attrs, func() string { return "parameter attrs" }); err != nil {
		return err
	}

	return nil
}

func validateServerlessCluster_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateServerlessCluster_IsResourceParameters(construct awscdk.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewServerlessClusterParameters(scope constructs.Construct, id *string, props *ServerlessClusterProps) error {
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

