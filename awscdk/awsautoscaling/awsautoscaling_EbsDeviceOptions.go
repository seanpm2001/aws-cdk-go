package awsautoscaling


// Block device options for an EBS volume.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ebsDeviceOptions := &ebsDeviceOptions{
//   	deleteOnTermination: jsii.Boolean(false),
//   	encrypted: jsii.Boolean(false),
//   	iops: jsii.Number(123),
//   	volumeType: awscdk.Aws_autoscaling.ebsDeviceVolumeType_STANDARD,
//   }
//
// Experimental.
type EbsDeviceOptions struct {
	// Indicates whether to delete the volume when the instance is terminated.
	// Experimental.
	DeleteOnTermination *bool `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// The number of I/O operations per second (IOPS) to provision for the volume.
	//
	// Must only be set for {@link volumeType}: {@link EbsDeviceVolumeType.IO1}
	//
	// The maximum ratio of IOPS to volume size (in GiB) is 50:1, so for 5,000 provisioned IOPS,
	// you need at least 100 GiB storage on the volume.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	// Experimental.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The EBS volume type.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html
	//
	// Experimental.
	VolumeType EbsDeviceVolumeType `field:"optional" json:"volumeType" yaml:"volumeType"`
	// Specifies whether the EBS volume is encrypted.
	//
	// Encrypted EBS volumes can only be attached to instances that support Amazon EBS encryption.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#EBSEncryption_supported_instances
	//
	// Experimental.
	Encrypted *bool `field:"optional" json:"encrypted" yaml:"encrypted"`
}
