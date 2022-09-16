//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsecspatterns

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationLoadBalancedServiceBase) validateAddServiceAsTargetParameters(service awsecs.BaseService) error {
	return nil
}

func (a *jsiiProxy_ApplicationLoadBalancedServiceBase) validateCreateAWSLogDriverParameters(prefix *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationLoadBalancedServiceBase) validateGetDefaultClusterParameters(scope awscdk.Construct) error {
	return nil
}

func (a *jsiiProxy_ApplicationLoadBalancedServiceBase) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (a *jsiiProxy_ApplicationLoadBalancedServiceBase) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateApplicationLoadBalancedServiceBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewApplicationLoadBalancedServiceBaseParameters(scope constructs.Construct, id *string, props *ApplicationLoadBalancedServiceBaseProps) error {
	return nil
}
