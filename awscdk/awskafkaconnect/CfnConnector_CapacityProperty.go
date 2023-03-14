package awskafkaconnect


// Information about the capacity of the connector, whether it is auto scaled or provisioned.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityProperty := &CapacityProperty{
//   	AutoScaling: &AutoScalingProperty{
//   		MaxWorkerCount: jsii.Number(123),
//   		McuCount: jsii.Number(123),
//   		MinWorkerCount: jsii.Number(123),
//   		ScaleInPolicy: &ScaleInPolicyProperty{
//   			CpuUtilizationPercentage: jsii.Number(123),
//   		},
//   		ScaleOutPolicy: &ScaleOutPolicyProperty{
//   			CpuUtilizationPercentage: jsii.Number(123),
//   		},
//   	},
//   	ProvisionedCapacity: &ProvisionedCapacityProperty{
//   		WorkerCount: jsii.Number(123),
//
//   		// the properties below are optional
//   		McuCount: jsii.Number(123),
//   	},
//   }
//
type CfnConnector_CapacityProperty struct {
	// Information about the auto scaling parameters for the connector.
	AutoScaling interface{} `field:"optional" json:"autoScaling" yaml:"autoScaling"`
	// Details about a fixed capacity allocated to a connector.
	ProvisionedCapacity interface{} `field:"optional" json:"provisionedCapacity" yaml:"provisionedCapacity"`
}

