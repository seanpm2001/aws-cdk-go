//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BastionHostLinux) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (b *jsiiProxy_BastionHostLinux) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (b *jsiiProxy_BastionHostLinux) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (b *jsiiProxy_BastionHostLinux) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (b *jsiiProxy_BastionHostLinux) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateBastionHostLinux_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBastionHostLinux_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewBastionHostLinuxParameters(scope constructs.Construct, id *string, props *BastionHostLinuxProps) error {
	return nil
}

