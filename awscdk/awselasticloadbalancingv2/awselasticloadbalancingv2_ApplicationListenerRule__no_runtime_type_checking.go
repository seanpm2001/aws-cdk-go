//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awselasticloadbalancingv2

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationListenerRule) validateAddConditionParameters(condition ListenerCondition) error {
	return nil
}

func (a *jsiiProxy_ApplicationListenerRule) validateAddFixedResponseParameters(fixedResponse *FixedResponse) error {
	return nil
}

func (a *jsiiProxy_ApplicationListenerRule) validateAddRedirectResponseParameters(redirectResponse *RedirectResponse) error {
	return nil
}

func (a *jsiiProxy_ApplicationListenerRule) validateAddTargetGroupParameters(targetGroup IApplicationTargetGroup) error {
	return nil
}

func (a *jsiiProxy_ApplicationListenerRule) validateConfigureActionParameters(action ListenerAction) error {
	return nil
}

func (a *jsiiProxy_ApplicationListenerRule) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (a *jsiiProxy_ApplicationListenerRule) validateSetConditionParameters(field *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationListenerRule) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateApplicationListenerRule_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewApplicationListenerRuleParameters(scope constructs.Construct, id *string, props *ApplicationListenerRuleProps) error {
	return nil
}

