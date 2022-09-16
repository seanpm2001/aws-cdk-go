//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnHook) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnHook) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnHook) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func validateCfnHook_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnHook_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCfnHookParameters(scope constructs.Construct, id *string, props *CfnHookProps) error {
	return nil
}
