//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VirtualRouterListener) validateBindParameters(scope awscdk.Construct) error {
	return nil
}
