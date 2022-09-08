//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTypeActivation) validateAddDeletionOverrideParameters(path *string) error {
	return nil
}

func (c *jsiiProxy_CfnTypeActivation) validateAddDependsOnParameters(target CfnResource) error {
	return nil
}

func (c *jsiiProxy_CfnTypeActivation) validateAddMetadataParameters(key *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnTypeActivation) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnTypeActivation) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	return nil
}

func (c *jsiiProxy_CfnTypeActivation) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnTypeActivation) validateApplyRemovalPolicyParameters(options *RemovalPolicyOptions) error {
	return nil
}

func (c *jsiiProxy_CfnTypeActivation) validateGetAttParameters(attributeName *string) error {
	return nil
}

func (c *jsiiProxy_CfnTypeActivation) validateGetMetadataParameters(key *string) error {
	return nil
}

func (c *jsiiProxy_CfnTypeActivation) validateInspectParameters(inspector TreeInspector) error {
	return nil
}

func (c *jsiiProxy_CfnTypeActivation) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnTypeActivation) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnTypeActivation) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnTypeActivation) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnTypeActivation) validateValidatePropertiesParameters(_properties interface{}) error {
	return nil
}

func validateCfnTypeActivation_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnTypeActivation_IsCfnResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnTypeActivation_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_CfnTypeActivation) validateSetAutoUpdateParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CfnTypeActivation) validateSetLoggingConfigParameters(val interface{}) error {
	return nil
}

func validateNewCfnTypeActivationParameters(scope Construct, id *string, props *CfnTypeActivationProps) error {
	return nil
}

