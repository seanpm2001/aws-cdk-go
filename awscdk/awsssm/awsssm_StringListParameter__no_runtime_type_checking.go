//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsssm

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StringListParameter) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_StringListParameter) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_StringListParameter) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_StringListParameter) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_StringListParameter) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_StringListParameter) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (s *jsiiProxy_StringListParameter) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateStringListParameter_FromStringListParameterNameParameters(scope constructs.Construct, id *string, stringListParameterName *string) error {
	return nil
}

func validateStringListParameter_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStringListParameter_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewStringListParameterParameters(scope constructs.Construct, id *string, props *StringListParameterProps) error {
	return nil
}
