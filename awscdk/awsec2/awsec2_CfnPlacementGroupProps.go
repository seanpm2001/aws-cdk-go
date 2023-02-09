package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnPlacementGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPlacementGroupProps := &cfnPlacementGroupProps{
//   	partitionCount: jsii.Number(123),
//   	spreadLevel: jsii.String("spreadLevel"),
//   	strategy: jsii.String("strategy"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPlacementGroupProps struct {
	// The number of partitions.
	//
	// Valid only when *Strategy* is set to `partition` .
	PartitionCount *float64 `field:"optional" json:"partitionCount" yaml:"partitionCount"`
	// Determines how placement groups spread instances.
	//
	// - Host – You can use `host` only with Outpost placement groups.
	// - Rack – No usage restrictions.
	SpreadLevel *string `field:"optional" json:"spreadLevel" yaml:"spreadLevel"`
	// The placement strategy.
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
	// The tags to apply to the new placement group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

