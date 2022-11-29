//go:build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnElement) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnElement) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnElement) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func validateCfnElement_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnElement_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCfnElementParameters(scope constructs.Construct, id *string) error {
	return nil
}

