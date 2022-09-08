//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnPublisher) validateAddDeletionOverrideParameters(path *string) error {
	return nil
}

func (c *jsiiProxy_CfnPublisher) validateAddDependsOnParameters(target CfnResource) error {
	return nil
}

func (c *jsiiProxy_CfnPublisher) validateAddMetadataParameters(key *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnPublisher) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnPublisher) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	return nil
}

func (c *jsiiProxy_CfnPublisher) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnPublisher) validateApplyRemovalPolicyParameters(options *RemovalPolicyOptions) error {
	return nil
}

func (c *jsiiProxy_CfnPublisher) validateGetAttParameters(attributeName *string) error {
	return nil
}

func (c *jsiiProxy_CfnPublisher) validateGetMetadataParameters(key *string) error {
	return nil
}

func (c *jsiiProxy_CfnPublisher) validateInspectParameters(inspector TreeInspector) error {
	return nil
}

func (c *jsiiProxy_CfnPublisher) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnPublisher) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_CfnPublisher) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	return nil
}

func (c *jsiiProxy_CfnPublisher) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_CfnPublisher) validateValidatePropertiesParameters(_properties interface{}) error {
	return nil
}

func validateCfnPublisher_IsCfnElementParameters(x interface{}) error {
	return nil
}

func validateCfnPublisher_IsCfnResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnPublisher_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_CfnPublisher) validateSetAcceptTermsAndConditionsParameters(val interface{}) error {
	return nil
}

func validateNewCfnPublisherParameters(scope Construct, id *string, props *CfnPublisherProps) error {
	return nil
}

