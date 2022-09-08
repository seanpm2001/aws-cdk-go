//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscognito

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserPoolIdentityProviderApple) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (u *jsiiProxy_UserPoolIdentityProviderApple) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (u *jsiiProxy_UserPoolIdentityProviderApple) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (u *jsiiProxy_UserPoolIdentityProviderApple) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (u *jsiiProxy_UserPoolIdentityProviderApple) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateUserPoolIdentityProviderApple_IsConstructParameters(x interface{}) error {
	return nil
}

func validateUserPoolIdentityProviderApple_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewUserPoolIdentityProviderAppleParameters(scope constructs.Construct, id *string, props *UserPoolIdentityProviderAppleProps) error {
	return nil
}

