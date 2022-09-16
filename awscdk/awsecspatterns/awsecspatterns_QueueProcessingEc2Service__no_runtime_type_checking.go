//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsecspatterns

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QueueProcessingEc2Service) validateConfigureAutoscalingForServiceParameters(service awsecs.BaseService) error {
	return nil
}

func (q *jsiiProxy_QueueProcessingEc2Service) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	return nil
}

func (q *jsiiProxy_QueueProcessingEc2Service) validateGrantPermissionsToServiceParameters(service awsecs.BaseService) error {
	return nil
}

func (q *jsiiProxy_QueueProcessingEc2Service) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (q *jsiiProxy_QueueProcessingEc2Service) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateQueueProcessingEc2Service_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewQueueProcessingEc2ServiceParameters(scope constructs.Construct, id *string, props *QueueProcessingEc2ServiceProps) error {
	return nil
}
