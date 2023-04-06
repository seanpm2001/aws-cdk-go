package awskafkaconnect


// Specifies how the connector scales.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoScalingProperty := &AutoScalingProperty{
//   	MaxWorkerCount: jsii.Number(123),
//   	McuCount: jsii.Number(123),
//   	MinWorkerCount: jsii.Number(123),
//   	ScaleInPolicy: &ScaleInPolicyProperty{
//   		CpuUtilizationPercentage: jsii.Number(123),
//   	},
//   	ScaleOutPolicy: &ScaleOutPolicyProperty{
//   		CpuUtilizationPercentage: jsii.Number(123),
//   	},
//   }
//
type CfnConnector_AutoScalingProperty struct {
	// The maximum number of workers allocated to the connector.
	MaxWorkerCount *float64 `field:"required" json:"maxWorkerCount" yaml:"maxWorkerCount"`
	// The number of microcontroller units (MCUs) allocated to each connector worker.
	//
	// The valid values are 1,2,4,8.
	McuCount *float64 `field:"required" json:"mcuCount" yaml:"mcuCount"`
	// The minimum number of workers allocated to the connector.
	MinWorkerCount *float64 `field:"required" json:"minWorkerCount" yaml:"minWorkerCount"`
	// The sacle-in policy for the connector.
	ScaleInPolicy interface{} `field:"required" json:"scaleInPolicy" yaml:"scaleInPolicy"`
	// The sacle-out policy for the connector.
	ScaleOutPolicy interface{} `field:"required" json:"scaleOutPolicy" yaml:"scaleOutPolicy"`
}
