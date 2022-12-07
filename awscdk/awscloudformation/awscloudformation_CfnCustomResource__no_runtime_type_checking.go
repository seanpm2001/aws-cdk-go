//go:build no_runtime_type_checking

package awscloudformation

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCustomResource) validateAddDeletionOverrideParameters(path *string) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateAddDependsOnParameters(target awscdk.CfnResource) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateAddMetadataParameters(key *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateApplyRemovalPolicyParameters(options *awscdk.RemovalPolicyOptions) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateGetAttParameters(attributeName *string) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateGetMetadataParameters(key *string) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateInspectParameters(inspector awscdk.TreeInspector) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnCustomResource) validateValidatePropertiesParameters(_properties interface{}) error {
	return nil
}

func validateCfnCustomResource_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnCustomResource_IsCfnResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnCustomResource_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_CfnCustomResource) validateSetServiceTokenParameters(val *string) error {
	return nil
}

func validateNewCfnCustomResourceParameters(scope constructs.Construct, id *string, props *CfnCustomResourceProps) error {
	return nil
}

