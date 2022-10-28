//go:build no_runtime_type_checking

// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Content) validateBindParameters(scope constructs.Construct, grantable awsiam.IGrantable) error {
	return nil
}

func validateContent_FromAssetParameters(path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateContent_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}
