//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LaunchTemplateRequireImdsv2Aspect) validateVisitParameters(node awscdk.IConstruct) error {
	return nil
}

func (l *jsiiProxy_LaunchTemplateRequireImdsv2Aspect) validateWarnParameters(node awscdk.IConstruct, message *string) error {
	return nil
}

func validateNewLaunchTemplateRequireImdsv2AspectParameters(props *LaunchTemplateRequireImdsv2AspectProps) error {
	return nil
}

