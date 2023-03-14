package awsopsworks


// Describes a layer's load-based auto scaling configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBasedAutoScalingProperty := &LoadBasedAutoScalingProperty{
//   	DownScaling: &AutoScalingThresholdsProperty{
//   		CpuThreshold: jsii.Number(123),
//   		IgnoreMetricsTime: jsii.Number(123),
//   		InstanceCount: jsii.Number(123),
//   		LoadThreshold: jsii.Number(123),
//   		MemoryThreshold: jsii.Number(123),
//   		ThresholdsWaitTime: jsii.Number(123),
//   	},
//   	Enable: jsii.Boolean(false),
//   	UpScaling: &AutoScalingThresholdsProperty{
//   		CpuThreshold: jsii.Number(123),
//   		IgnoreMetricsTime: jsii.Number(123),
//   		InstanceCount: jsii.Number(123),
//   		LoadThreshold: jsii.Number(123),
//   		MemoryThreshold: jsii.Number(123),
//   		ThresholdsWaitTime: jsii.Number(123),
//   	},
//   }
//
type CfnLayer_LoadBasedAutoScalingProperty struct {
	// An `AutoScalingThresholds` object that describes the downscaling configuration, which defines how and when AWS OpsWorks Stacks reduces the number of instances.
	DownScaling interface{} `field:"optional" json:"downScaling" yaml:"downScaling"`
	// Whether load-based auto scaling is enabled for the layer.
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// An `AutoScalingThresholds` object that describes the upscaling configuration, which defines how and when AWS OpsWorks Stacks increases the number of instances.
	UpScaling interface{} `field:"optional" json:"upScaling" yaml:"upScaling"`
}

