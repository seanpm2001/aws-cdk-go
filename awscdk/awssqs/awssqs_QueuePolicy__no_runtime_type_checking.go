//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awssqs

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QueuePolicy) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (q *jsiiProxy_QueuePolicy) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (q *jsiiProxy_QueuePolicy) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (q *jsiiProxy_QueuePolicy) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (q *jsiiProxy_QueuePolicy) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateQueuePolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateQueuePolicy_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewQueuePolicyParameters(scope constructs.Construct, id *string, props *QueuePolicyProps) error {
	return nil
}

