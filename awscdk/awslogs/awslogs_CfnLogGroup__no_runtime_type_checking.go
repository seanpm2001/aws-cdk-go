//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLogGroup) validateAddDeletionOverrideParameters(path *string) error {
	return nil
}

func (c *jsiiProxy_CfnLogGroup) validateAddDependsOnParameters(target awscdk.CfnResource) error {
	return nil
}

func (c *jsiiProxy_CfnLogGroup) validateAddMetadataParameters(key *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnLogGroup) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnLogGroup) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	return nil
}

func (c *jsiiProxy_CfnLogGroup) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnLogGroup) validateApplyRemovalPolicyParameters(options *awscdk.RemovalPolicyOptions) error {
	return nil
}

func (c *jsiiProxy_CfnLogGroup) validateGetAttParameters(attributeName *string) error {
	return nil
}

func (c *jsiiProxy_CfnLogGroup) validateGetMetadataParameters(key *string) error {
	return nil
}

func (c *jsiiProxy_CfnLogGroup) validateInspectParameters(inspector awscdk.TreeInspector) error {
	return nil
}

func (c *jsiiProxy_CfnLogGroup) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnLogGroup) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnLogGroup) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnLogGroup) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnLogGroup) validateValidatePropertiesParameters(_properties interface{}) error {
	return nil
}

func validateCfnLogGroup_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnLogGroup_IsCfnResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnLogGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCfnLogGroupParameters(scope awscdk.Construct, id *string, props *CfnLogGroupProps) error {
	return nil
}

