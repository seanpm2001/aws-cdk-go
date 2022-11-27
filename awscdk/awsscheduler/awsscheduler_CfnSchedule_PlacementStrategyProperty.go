package awsscheduler


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   placementStrategyProperty := &placementStrategyProperty{
//   	field: jsii.String("field"),
//   	type: jsii.String("type"),
//   }
//
type CfnSchedule_PlacementStrategyProperty struct {
	// `CfnSchedule.PlacementStrategyProperty.Field`.
	Field *string `field:"optional" json:"field" yaml:"field"`
	// `CfnSchedule.PlacementStrategyProperty.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}
