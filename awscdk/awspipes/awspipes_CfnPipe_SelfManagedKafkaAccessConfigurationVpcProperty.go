package awspipes


// This structure specifies the VPC subnets and security groups for the stream, and whether a public IP address is to be used.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selfManagedKafkaAccessConfigurationVpcProperty := &selfManagedKafkaAccessConfigurationVpcProperty{
//   	securityGroup: []*string{
//   		jsii.String("securityGroup"),
//   	},
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   }
//
type CfnPipe_SelfManagedKafkaAccessConfigurationVpcProperty struct {
	// Specifies the security groups associated with the stream.
	//
	// These security groups must all be in the same VPC. You can specify as many as five security groups. If you do not specify a security group, the default security group for the VPC is used.
	SecurityGroup *[]*string `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Specifies the subnets associated with the stream.
	//
	// These subnets must all be in the same VPC. You can specify as many as 16 subnets.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}
