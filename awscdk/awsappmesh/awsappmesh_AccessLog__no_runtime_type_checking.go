//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessLog) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateAccessLog_FromFilePathParameters(filePath *string) error {
	return nil
}

