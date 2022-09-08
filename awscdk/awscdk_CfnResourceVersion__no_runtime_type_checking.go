//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnResourceVersion) validateAddDeletionOverrideParameters(path *string) error {
	return nil
}

func (c *jsiiProxy_CfnResourceVersion) validateAddDependsOnParameters(target CfnResource) error {
	return nil
}

func (c *jsiiProxy_CfnResourceVersion) validateAddMetadataParameters(key *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnResourceVersion) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnResourceVersion) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	return nil
}

func (c *jsiiProxy_CfnResourceVersion) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnResourceVersion) validateApplyRemovalPolicyParameters(options *RemovalPolicyOptions) error {
	return nil
}

func (c *jsiiProxy_CfnResourceVersion) validateGetAttParameters(attributeName *string) error {
	return nil
}

func (c *jsiiProxy_CfnResourceVersion) validateGetMetadataParameters(key *string) error {
	return nil
}

func (c *jsiiProxy_CfnResourceVersion) validateInspectParameters(inspector TreeInspector) error {
	return nil
}

func (c *jsiiProxy_CfnResourceVersion) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnResourceVersion) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnResourceVersion) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnResourceVersion) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnResourceVersion) validateValidatePropertiesParameters(_properties interface{}) error {
	return nil
}

func validateCfnResourceVersion_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnResourceVersion_IsCfnResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnResourceVersion_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_CfnResourceVersion) validateSetLoggingConfigParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CfnResourceVersion) validateSetSchemaHandlerPackageParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CfnResourceVersion) validateSetTypeNameParameters(val *string) error {
	return nil
}

func validateNewCfnResourceVersionParameters(scope Construct, id *string, props *CfnResourceVersionProps) error {
	return nil
}

