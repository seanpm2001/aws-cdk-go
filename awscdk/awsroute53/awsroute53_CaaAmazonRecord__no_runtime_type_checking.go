//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CaaAmazonRecord) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CaaAmazonRecord) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CaaAmazonRecord) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_CaaAmazonRecord) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CaaAmazonRecord) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateCaaAmazonRecord_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCaaAmazonRecord_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewCaaAmazonRecordParameters(scope constructs.Construct, id *string, props *CaaAmazonRecordProps) error {
	return nil
}
