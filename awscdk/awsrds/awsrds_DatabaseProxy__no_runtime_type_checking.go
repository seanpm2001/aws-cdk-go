//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseProxy) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DatabaseProxy) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DatabaseProxy) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseProxy) validateGrantConnectParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (d *jsiiProxy_DatabaseProxy) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (d *jsiiProxy_DatabaseProxy) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateDatabaseProxy_FromDatabaseProxyAttributesParameters(scope constructs.Construct, id *string, attrs *DatabaseProxyAttributes) error {
	return nil
}

func validateDatabaseProxy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDatabaseProxy_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewDatabaseProxyParameters(scope constructs.Construct, id *string, props *DatabaseProxyProps) error {
	return nil
}

