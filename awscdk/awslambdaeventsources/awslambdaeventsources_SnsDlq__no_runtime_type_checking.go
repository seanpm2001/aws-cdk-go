//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awslambdaeventsources

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SnsDlq) validateBindParameters(_target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) error {
	return nil
}

func validateNewSnsDlqParameters(topic awssns.ITopic) error {
	return nil
}
