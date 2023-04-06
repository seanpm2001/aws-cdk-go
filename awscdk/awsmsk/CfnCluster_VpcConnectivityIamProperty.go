package awsmsk


// Not currently supported by AWS CloudFormation .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConnectivityIamProperty := &VpcConnectivityIamProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
type CfnCluster_VpcConnectivityIamProperty struct {
	// `CfnCluster.VpcConnectivityIamProperty.Enabled`.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

