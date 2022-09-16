//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsservicecatalog

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProductStackHistory) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (p *jsiiProxy_ProductStackHistory) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func (p *jsiiProxy_ProductStackHistory) validateVersionFromSnapshotParameters(productVersionName *string) error {
	return nil
}

func validateProductStackHistory_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewProductStackHistoryParameters(scope constructs.Construct, id *string, props *ProductStackHistoryProps) error {
	return nil
}
