package awsinternetmonitor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMonitor`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMonitorProps := &CfnMonitorProps{
//   	MonitorName: jsii.String("monitorName"),
//
//   	// the properties below are optional
//   	MaxCityNetworksToMonitor: jsii.Number(123),
//   	Resources: []*string{
//   		jsii.String("resources"),
//   	},
//   	ResourcesToAdd: []*string{
//   		jsii.String("resourcesToAdd"),
//   	},
//   	ResourcesToRemove: []*string{
//   		jsii.String("resourcesToRemove"),
//   	},
//   	Status: jsii.String("status"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnMonitorProps struct {
	// The name of the monitor.
	//
	// A monitor name can contain only alphanumeric characters, dashes (-), periods (.), and underscores (_).
	MonitorName *string `field:"required" json:"monitorName" yaml:"monitorName"`
	// The maximum number of city-networks to monitor for your resources.
	//
	// A city-network is the location (city) where clients access your application resources from and the network, such as an internet service provider, that clients access the resources through.
	//
	// For more information, see [Choosing a city-network maximum value](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/IMCityNetworksMaximum.html) in *Using Amazon CloudWatch Internet Monitor* .
	MaxCityNetworksToMonitor *float64 `field:"optional" json:"maxCityNetworksToMonitor" yaml:"maxCityNetworksToMonitor"`
	// The resources that have been added for the monitor, listed by their Amazon Resource Names (ARNs).
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// The resources to add to a monitor, which you provide as a set of Amazon Resource Names (ARNs).
	//
	// You can add a combination of Virtual Private Clouds (VPCs) and Amazon CloudFront distributions, or you can add WorkSpaces directories. You can't add all three types of resources.
	//
	// > If you add only VPC resources, at least one VPC must have an Internet Gateway attached to it, to make sure that it has internet connectivity.
	ResourcesToAdd *[]*string `field:"optional" json:"resourcesToAdd" yaml:"resourcesToAdd"`
	// The resources to remove from a monitor, which you provide as a set of Amazon Resource Names (ARNs).
	ResourcesToRemove *[]*string `field:"optional" json:"resourcesToRemove" yaml:"resourcesToRemove"`
	// The status of a monitor.
	//
	// The accepted values that you can specify for `Status` are `ACTIVE` and `INACTIVE` .
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The tags for a monitor, listed as a set of *key:value* pairs.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

