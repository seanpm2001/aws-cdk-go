//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awseks

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KubernetesPatch) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (k *jsiiProxy_KubernetesPatch) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateKubernetesPatch_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewKubernetesPatchParameters(scope constructs.Construct, id *string, props *KubernetesPatchProps) error {
	return nil
}

