//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscodebuild

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BitBucketSourceCredentials) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (b *jsiiProxy_BitBucketSourceCredentials) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (b *jsiiProxy_BitBucketSourceCredentials) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (b *jsiiProxy_BitBucketSourceCredentials) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (b *jsiiProxy_BitBucketSourceCredentials) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateBitBucketSourceCredentials_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBitBucketSourceCredentials_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewBitBucketSourceCredentialsParameters(scope constructs.Construct, id *string, props *BitBucketSourceCredentialsProps) error {
	return nil
}

