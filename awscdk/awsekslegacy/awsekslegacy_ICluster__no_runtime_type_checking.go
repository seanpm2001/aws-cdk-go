//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsekslegacy

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ICluster) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

