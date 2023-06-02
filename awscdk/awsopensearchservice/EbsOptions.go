package awsopensearchservice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to data nodes in the Amazon OpenSearch Service domain.
//
// For more information, see
// [Amazon EBS]
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AmazonEBS.html)
// in the Amazon Elastic Compute Cloud Developer Guide.
//
// Example:
//   prodDomain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
//   	Capacity: &CapacityConfig{
//   		MasterNodes: jsii.Number(5),
//   		DataNodes: jsii.Number(20),
//   	},
//   	Ebs: &EbsOptions{
//   		VolumeSize: jsii.Number(20),
//   	},
//   	ZoneAwareness: &ZoneAwarenessConfig{
//   		AvailabilityZoneCount: jsii.Number(3),
//   	},
//   	Logging: &LoggingOptions{
//   		SlowSearchLogEnabled: jsii.Boolean(true),
//   		AppLogEnabled: jsii.Boolean(true),
//   		SlowIndexLogEnabled: jsii.Boolean(true),
//   	},
//   })
//
type EbsOptions struct {
	// Specifies whether Amazon EBS volumes are attached to data nodes in the Amazon OpenSearch Service domain.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The number of I/O operations per second (IOPS) that the volume supports.
	//
	// This property applies only to the Provisioned IOPS (SSD) EBS
	// volume type.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The size (in GiB) of the EBS volume for each data node.
	//
	// The minimum and
	// maximum size of an EBS volume depends on the EBS volume type and the
	// instance type to which it is attached.  For  valid values, see
	// [EBS volume size limits]
	// (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/limits.html#ebsresource)
	// in the Amazon OpenSearch Service Developer Guide.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// The EBS volume type to use with the Amazon OpenSearch Service domain, such as standard, gp2, io1.
	VolumeType awsec2.EbsDeviceVolumeType `field:"optional" json:"volumeType" yaml:"volumeType"`
}

