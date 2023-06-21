package awsosis


// Options that specify the subnets and security groups for an OpenSearch Ingestion VPC endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcOptionsProperty := &VpcOptionsProperty{
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
type CfnPipeline_VpcOptionsProperty struct {
	// A list of security groups associated with the VPC endpoint.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of subnet IDs associated with the VPC endpoint.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}
