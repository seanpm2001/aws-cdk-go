//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awssns

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Subscription) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_Subscription) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_Subscription) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_Subscription) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (s *jsiiProxy_Subscription) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateSubscription_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSubscription_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewSubscriptionParameters(scope constructs.Construct, id *string, props *SubscriptionProps) error {
	return nil
}

