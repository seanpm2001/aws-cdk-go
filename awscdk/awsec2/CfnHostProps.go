package awsec2


// Properties for defining a `CfnHost`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnHostProps := &CfnHostProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//
//   	// the properties below are optional
//   	AutoPlacement: jsii.String("autoPlacement"),
//   	HostRecovery: jsii.String("hostRecovery"),
//   	InstanceFamily: jsii.String("instanceFamily"),
//   	InstanceType: jsii.String("instanceType"),
//   	OutpostArn: jsii.String("outpostArn"),
//   }
//
type CfnHostProps struct {
	// The Availability Zone in which to allocate the Dedicated Host.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.
	//
	// For more information, see [Understanding auto-placement and affinity](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/how-dedicated-hosts-work.html#dedicated-hosts-understanding) in the *Amazon EC2 User Guide* .
	//
	// Default: `on`.
	AutoPlacement *string `field:"optional" json:"autoPlacement" yaml:"autoPlacement"`
	// Indicates whether to enable or disable host recovery for the Dedicated Host.
	//
	// Host recovery is disabled by default. For more information, see [Host recovery](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-recovery.html) in the *Amazon EC2 User Guide* .
	//
	// Default: `off`.
	HostRecovery *string `field:"optional" json:"hostRecovery" yaml:"hostRecovery"`
	// The instance family supported by the Dedicated Host.
	//
	// For example, `m5` .
	InstanceFamily *string `field:"optional" json:"instanceFamily" yaml:"instanceFamily"`
	// Specifies the instance type to be supported by the Dedicated Hosts.
	//
	// If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The Amazon Resource Name (ARN) of the AWS Outpost on which the Dedicated Host is allocated.
	OutpostArn *string `field:"optional" json:"outpostArn" yaml:"outpostArn"`
}
