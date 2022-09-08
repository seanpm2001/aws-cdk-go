//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsiotevents

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DetectorModel) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DetectorModel) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DetectorModel) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_DetectorModel) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (d *jsiiProxy_DetectorModel) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateDetectorModel_FromDetectorModelNameParameters(scope constructs.Construct, id *string, detectorModelName *string) error {
	return nil
}

func validateDetectorModel_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDetectorModel_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewDetectorModelParameters(scope constructs.Construct, id *string, props *DetectorModelProps) error {
	return nil
}

