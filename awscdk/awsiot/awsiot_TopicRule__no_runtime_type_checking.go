//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsiot

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TopicRule) validateAddActionParameters(action IAction) error {
	return nil
}

func (t *jsiiProxy_TopicRule) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TopicRule) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TopicRule) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (t *jsiiProxy_TopicRule) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (t *jsiiProxy_TopicRule) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateTopicRule_FromTopicRuleArnParameters(scope constructs.Construct, id *string, topicRuleArn *string) error {
	return nil
}

func validateTopicRule_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTopicRule_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewTopicRuleParameters(scope constructs.Construct, id *string, props *TopicRuleProps) error {
	return nil
}

