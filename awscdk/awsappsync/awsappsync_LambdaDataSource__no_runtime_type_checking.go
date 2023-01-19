//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaDataSource) validateCreateFunctionParameters(props *BaseAppsyncFunctionProps) error {
	return nil
}

func (l *jsiiProxy_LambdaDataSource) validateCreateResolverParameters(props *BaseResolverProps) error {
	return nil
}

func (l *jsiiProxy_LambdaDataSource) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (l *jsiiProxy_LambdaDataSource) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateLambdaDataSource_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_LambdaDataSource) validateSetApiParameters(val IGraphqlApi) error {
	return nil
}

func validateNewLambdaDataSourceParameters(scope constructs.Construct, id *string, props *LambdaDataSourceProps) error {
	return nil
}
