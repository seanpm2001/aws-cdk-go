//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// The CDK Construct Library for AWS::Glue
package awscdkgluealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateJobExecutable_OfParameters(config *JobExecutableConfig) error {
	return nil
}

func validateJobExecutable_PythonEtlParameters(props *PythonSparkJobExecutableProps) error {
	return nil
}

func validateJobExecutable_PythonShellParameters(props *PythonShellExecutableProps) error {
	return nil
}

func validateJobExecutable_PythonStreamingParameters(props *PythonSparkJobExecutableProps) error {
	return nil
}

func validateJobExecutable_ScalaEtlParameters(props *ScalaJobExecutableProps) error {
	return nil
}

func validateJobExecutable_ScalaStreamingParameters(props *ScalaJobExecutableProps) error {
	return nil
}

