//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsautoscaling

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AutoScalingGroupRequireImdsv2Aspect) validateVisitParameters(node awscdk.IConstruct) error {
	return nil
}

func (a *jsiiProxy_AutoScalingGroupRequireImdsv2Aspect) validateWarnParameters(node awscdk.IConstruct, message *string) error {
	return nil
}
