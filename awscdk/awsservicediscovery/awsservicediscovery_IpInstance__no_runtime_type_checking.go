//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsservicediscovery

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IpInstance) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IpInstance) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_IpInstance) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (i *jsiiProxy_IpInstance) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (i *jsiiProxy_IpInstance) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateIpInstance_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIpInstance_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewIpInstanceParameters(scope constructs.Construct, id *string, props *IpInstanceProps) error {
	return nil
}
