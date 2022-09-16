//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsglobalaccelerator

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Accelerator) validateAddListenerParameters(id *string, options *ListenerOptions) error {
	return nil
}

func (a *jsiiProxy_Accelerator) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_Accelerator) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_Accelerator) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_Accelerator) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (a *jsiiProxy_Accelerator) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateAccelerator_FromAcceleratorAttributesParameters(scope constructs.Construct, id *string, attrs *AcceleratorAttributes) error {
	return nil
}

func validateAccelerator_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAccelerator_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewAcceleratorParameters(scope constructs.Construct, id *string, props *AcceleratorProps) error {
	return nil
}
