//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnHookTypeConfig) validateAddDeletionOverrideParameters(path *string) error {
	return nil
}

func (c *jsiiProxy_CfnHookTypeConfig) validateAddDependsOnParameters(target CfnResource) error {
	return nil
}

func (c *jsiiProxy_CfnHookTypeConfig) validateAddMetadataParameters(key *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnHookTypeConfig) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnHookTypeConfig) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	return nil
}

func (c *jsiiProxy_CfnHookTypeConfig) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnHookTypeConfig) validateApplyRemovalPolicyParameters(options *RemovalPolicyOptions) error {
	return nil
}

func (c *jsiiProxy_CfnHookTypeConfig) validateGetAttParameters(attributeName *string) error {
	return nil
}

func (c *jsiiProxy_CfnHookTypeConfig) validateGetMetadataParameters(key *string) error {
	return nil
}

func (c *jsiiProxy_CfnHookTypeConfig) validateInspectParameters(inspector TreeInspector) error {
	return nil
}

func (c *jsiiProxy_CfnHookTypeConfig) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnHookTypeConfig) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnHookTypeConfig) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnHookTypeConfig) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnHookTypeConfig) validateValidatePropertiesParameters(_properties interface{}) error {
	return nil
}

func validateCfnHookTypeConfig_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnHookTypeConfig_IsCfnResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnHookTypeConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_CfnHookTypeConfig) validateSetConfigurationParameters(val *string) error {
	return nil
}

func validateNewCfnHookTypeConfigParameters(scope Construct, id *string, props *CfnHookTypeConfigProps) error {
	return nil
}

