//go:build no_runtime_type_checking

package awsiotevents

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IInput) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IInput) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

