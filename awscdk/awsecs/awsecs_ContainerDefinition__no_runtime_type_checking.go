//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerDefinition) validateAddContainerDependenciesParameters(containerDependencies *[]*ContainerDependency) error {
	return nil
}

func (c *jsiiProxy_ContainerDefinition) validateAddEnvironmentParameters(name *string, value *string) error {
	return nil
}

func (c *jsiiProxy_ContainerDefinition) validateAddLinkParameters(container ContainerDefinition) error {
	return nil
}

func (c *jsiiProxy_ContainerDefinition) validateAddMountPointsParameters(mountPoints *[]*MountPoint) error {
	return nil
}

func (c *jsiiProxy_ContainerDefinition) validateAddPortMappingsParameters(portMappings *[]*PortMapping) error {
	return nil
}

func (c *jsiiProxy_ContainerDefinition) validateAddScratchParameters(scratch *ScratchSpace) error {
	return nil
}

func (c *jsiiProxy_ContainerDefinition) validateAddToExecutionPolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (c *jsiiProxy_ContainerDefinition) validateAddUlimitsParameters(ulimits *[]*Ulimit) error {
	return nil
}

func (c *jsiiProxy_ContainerDefinition) validateAddVolumesFromParameters(volumesFrom *[]*VolumeFrom) error {
	return nil
}

func (c *jsiiProxy_ContainerDefinition) validateFindPortMappingParameters(containerPort *float64, protocol Protocol) error {
	return nil
}

func (c *jsiiProxy_ContainerDefinition) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_ContainerDefinition) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateContainerDefinition_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewContainerDefinitionParameters(scope constructs.Construct, id *string, props *ContainerDefinitionProps) error {
	return nil
}
