//go:build no_runtime_type_checking

package awscognito

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserPoolIdentityProviderGoogle) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (u *jsiiProxy_UserPoolIdentityProviderGoogle) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (u *jsiiProxy_UserPoolIdentityProviderGoogle) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (u *jsiiProxy_UserPoolIdentityProviderGoogle) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (u *jsiiProxy_UserPoolIdentityProviderGoogle) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateUserPoolIdentityProviderGoogle_IsConstructParameters(x interface{}) error {
	return nil
}

func validateUserPoolIdentityProviderGoogle_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewUserPoolIdentityProviderGoogleParameters(scope constructs.Construct, id *string, props *UserPoolIdentityProviderGoogleProps) error {
	return nil
}

