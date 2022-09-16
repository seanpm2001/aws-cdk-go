//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscloudformation

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnResourceDefaultVersion) validateAddDeletionOverrideParameters(path *string) error {
	return nil
}

func (c *jsiiProxy_CfnResourceDefaultVersion) validateAddDependsOnParameters(target awscdk.CfnResource) error {
	return nil
}

func (c *jsiiProxy_CfnResourceDefaultVersion) validateAddMetadataParameters(key *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnResourceDefaultVersion) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnResourceDefaultVersion) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	return nil
}

func (c *jsiiProxy_CfnResourceDefaultVersion) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnResourceDefaultVersion) validateApplyRemovalPolicyParameters(options *awscdk.RemovalPolicyOptions) error {
	return nil
}

func (c *jsiiProxy_CfnResourceDefaultVersion) validateGetAttParameters(attributeName *string) error {
	return nil
}

func (c *jsiiProxy_CfnResourceDefaultVersion) validateGetMetadataParameters(key *string) error {
	return nil
}

func (c *jsiiProxy_CfnResourceDefaultVersion) validateInspectParameters(inspector awscdk.TreeInspector) error {
	return nil
}

func (c *jsiiProxy_CfnResourceDefaultVersion) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnResourceDefaultVersion) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnResourceDefaultVersion) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnResourceDefaultVersion) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnResourceDefaultVersion) validateValidatePropertiesParameters(_properties interface{}) error {
	return nil
}

func validateCfnResourceDefaultVersion_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnResourceDefaultVersion_IsCfnResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnResourceDefaultVersion_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCfnResourceDefaultVersionParameters(scope awscdk.Construct, id *string, props *CfnResourceDefaultVersionProps) error {
	return nil
}
