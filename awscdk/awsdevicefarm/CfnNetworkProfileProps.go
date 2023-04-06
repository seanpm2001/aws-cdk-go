package awsdevicefarm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnNetworkProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNetworkProfileProps := &CfnNetworkProfileProps{
//   	Name: jsii.String("name"),
//   	ProjectArn: jsii.String("projectArn"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	DownlinkBandwidthBits: jsii.Number(123),
//   	DownlinkDelayMs: jsii.Number(123),
//   	DownlinkJitterMs: jsii.Number(123),
//   	DownlinkLossPercent: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UplinkBandwidthBits: jsii.Number(123),
//   	UplinkDelayMs: jsii.Number(123),
//   	UplinkJitterMs: jsii.Number(123),
//   	UplinkLossPercent: jsii.Number(123),
//   }
//
type CfnNetworkProfileProps struct {
	// The name of the network profile.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the specified project.
	ProjectArn *string `field:"required" json:"projectArn" yaml:"projectArn"`
	// The description of the network profile.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The data throughput rate in bits per second, as an integer from 0 to 104857600.
	DownlinkBandwidthBits *float64 `field:"optional" json:"downlinkBandwidthBits" yaml:"downlinkBandwidthBits"`
	// Delay time for all packets to destination in milliseconds as an integer from 0 to 2000.
	DownlinkDelayMs *float64 `field:"optional" json:"downlinkDelayMs" yaml:"downlinkDelayMs"`
	// Time variation in the delay of received packets in milliseconds as an integer from 0 to 2000.
	DownlinkJitterMs *float64 `field:"optional" json:"downlinkJitterMs" yaml:"downlinkJitterMs"`
	// Proportion of received packets that fail to arrive from 0 to 100 percent.
	DownlinkLossPercent *float64 `field:"optional" json:"downlinkLossPercent" yaml:"downlinkLossPercent"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The data throughput rate in bits per second, as an integer from 0 to 104857600.
	UplinkBandwidthBits *float64 `field:"optional" json:"uplinkBandwidthBits" yaml:"uplinkBandwidthBits"`
	// Delay time for all packets to destination in milliseconds as an integer from 0 to 2000.
	UplinkDelayMs *float64 `field:"optional" json:"uplinkDelayMs" yaml:"uplinkDelayMs"`
	// Time variation in the delay of received packets in milliseconds as an integer from 0 to 2000.
	UplinkJitterMs *float64 `field:"optional" json:"uplinkJitterMs" yaml:"uplinkJitterMs"`
	// Proportion of transmitted packets that fail to arrive from 0 to 100 percent.
	UplinkLossPercent *float64 `field:"optional" json:"uplinkLossPercent" yaml:"uplinkLossPercent"`
}

