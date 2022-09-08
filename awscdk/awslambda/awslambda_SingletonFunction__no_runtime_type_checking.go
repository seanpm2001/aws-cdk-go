//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SingletonFunction) validateAddEnvironmentParameters(key *string, value *string, options *EnvironmentOptions) error {
	return nil
}

func (s *jsiiProxy_SingletonFunction) validateAddEventSourceParameters(source IEventSource) error {
	return nil
}

func (s *jsiiProxy_SingletonFunction) validateAddEventSourceMappingParameters(id *string, options *EventSourceMappingOptions) error {
	return nil
}

func (s *jsiiProxy_SingletonFunction) validateAddFunctionUrlParameters(options *FunctionUrlOptions) error {
	return nil
}

func (s *jsiiProxy_SingletonFunction) validateAddPermissionParameters(name *string, permission *Permission) error {
	return nil
}

func (s *jsiiProxy_SingletonFunction) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_SingletonFunction) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SingletonFunction) validateConfigureAsyncInvokeParameters(options *EventInvokeConfigOptions) error {
	return nil
}

func (s *jsiiProxy_SingletonFunction) validateConsiderWarningOnInvokeFunctionPermissionsParameters(scope awscdk.Construct, action *string) error {
	return nil
}

func (s *jsiiProxy_SingletonFunction) validateDependOnParameters(down awscdk.IConstruct) error {
	return nil
}

func (s *jsiiProxy_SingletonFunction) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SingletonFunction) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_SingletonFunction) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_SingletonFunction) validateGrantInvokeUrlParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_SingletonFunction) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SingletonFunction) validateMetricDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SingletonFunction) validateMetricErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SingletonFunction) validateMetricInvocationsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SingletonFunction) validateMetricThrottlesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SingletonFunction) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (s *jsiiProxy_SingletonFunction) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func (s *jsiiProxy_SingletonFunction) validateWarnInvokeFunctionPermissionsParameters(scope awscdk.Construct) error {
	return nil
}

func validateSingletonFunction_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSingletonFunction_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewSingletonFunctionParameters(scope constructs.Construct, id *string, props *SingletonFunctionProps) error {
	return nil
}

