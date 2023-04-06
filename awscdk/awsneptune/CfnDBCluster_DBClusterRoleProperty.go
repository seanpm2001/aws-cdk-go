package awsneptune


// Describes an Amazon Identity and Access Management (IAM) role that is associated with a DB cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBClusterRoleProperty := &DBClusterRoleProperty{
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	FeatureName: jsii.String("featureName"),
//   }
//
type CfnDBCluster_DBClusterRoleProperty struct {
	// The Amazon Resource Name (ARN) of the IAM role that is associated with the DB cluster.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the feature associated with the Amazon Identity and Access Management (IAM) role.
	//
	// For the list of supported feature names, see [DescribeDBEngineVersions](https://docs.aws.amazon.com/neptune/latest/userguide/api-other-apis.html#DescribeDBEngineVersions) .
	FeatureName *string `field:"optional" json:"featureName" yaml:"featureName"`
}
