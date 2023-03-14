package awsemr


// `InstanceFleetProvisioningSpecification` is a subproperty of `InstanceFleetConfig` .
//
// `InstanceFleetProvisioningSpecification` defines the launch specification for Spot instances in an instance fleet, which determines the defined duration and provisioning timeout behavior for Spot instances.
//
// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceFleetProvisioningSpecificationsProperty := &InstanceFleetProvisioningSpecificationsProperty{
//   	OnDemandSpecification: &OnDemandProvisioningSpecificationProperty{
//   		AllocationStrategy: jsii.String("allocationStrategy"),
//   	},
//   	SpotSpecification: &SpotProvisioningSpecificationProperty{
//   		TimeoutAction: jsii.String("timeoutAction"),
//   		TimeoutDurationMinutes: jsii.Number(123),
//
//   		// the properties below are optional
//   		AllocationStrategy: jsii.String("allocationStrategy"),
//   		BlockDurationMinutes: jsii.Number(123),
//   	},
//   }
//
type CfnCluster_InstanceFleetProvisioningSpecificationsProperty struct {
	// The launch specification for On-Demand Instances in the instance fleet, which determines the allocation strategy.
	//
	// > The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions. On-Demand Instances allocation strategy is available in Amazon EMR version 5.12.1 and later.
	OnDemandSpecification interface{} `field:"optional" json:"onDemandSpecification" yaml:"onDemandSpecification"`
	// The launch specification for Spot instances in the fleet, which determines the defined duration, provisioning timeout behavior, and allocation strategy.
	SpotSpecification interface{} `field:"optional" json:"spotSpecification" yaml:"spotSpecification"`
}

