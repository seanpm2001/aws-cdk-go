package awsec2


// Specify configuration parameters for a single subnet group in a VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subnetConfiguration := &SubnetConfiguration{
//   	Name: jsii.String("name"),
//   	SubnetType: awscdk.Aws_ec2.SubnetType_ISOLATED,
//
//   	// the properties below are optional
//   	CidrMask: jsii.Number(123),
//   	MapPublicIpOnLaunch: jsii.Boolean(false),
//   	Reserved: jsii.Boolean(false),
//   }
//
// Experimental.
type SubnetConfiguration struct {
	// Logical name for the subnet group.
	//
	// This name can be used when selecting VPC subnets to distinguish
	// between different subnet groups of the same type.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of Subnet to configure.
	//
	// The Subnet type will control the ability to route and connect to the
	// Internet.
	// Experimental.
	SubnetType SubnetType `field:"required" json:"subnetType" yaml:"subnetType"`
	// The number of leading 1 bits in the routing mask.
	//
	// The number of available IP addresses in each subnet of this group
	// will be equal to `2^(32 - cidrMask) - 2`.
	//
	// Valid values are `16--28`.
	// Experimental.
	CidrMask *float64 `field:"optional" json:"cidrMask" yaml:"cidrMask"`
	// Controls if a public IP is associated to an instance at launch.
	// Experimental.
	MapPublicIpOnLaunch *bool `field:"optional" json:"mapPublicIpOnLaunch" yaml:"mapPublicIpOnLaunch"`
	// Controls if subnet IP space needs to be reserved.
	//
	// When true, the IP space for the subnet is reserved but no actual
	// resources are provisioned. This space is only dependent on the
	// number of availability zones and on `cidrMask` - all other subnet
	// properties are ignored.
	// Experimental.
	Reserved *bool `field:"optional" json:"reserved" yaml:"reserved"`
}

