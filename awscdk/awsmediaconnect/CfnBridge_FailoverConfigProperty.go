package awsmediaconnect


// The settings for source failover.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   failoverConfigProperty := &FailoverConfigProperty{
//   	FailoverMode: jsii.String("failoverMode"),
//
//   	// the properties below are optional
//   	SourcePriority: &SourcePriorityProperty{
//   		PrimarySource: jsii.String("primarySource"),
//   	},
//   	State: jsii.String("state"),
//   }
//
type CfnBridge_FailoverConfigProperty struct {
	// The type of failover you choose for this flow.
	//
	// MERGE combines the source streams into a single stream, allowing graceful recovery from any single-source loss. FAILOVER allows switching between different streams.
	FailoverMode *string `field:"required" json:"failoverMode" yaml:"failoverMode"`
	// The priority you want to assign to a source.
	//
	// You can have a primary stream and a backup stream or two equally prioritized streams. This setting only applies when Failover Mode is set to FAILOVER.
	SourcePriority interface{} `field:"optional" json:"sourcePriority" yaml:"sourcePriority"`
	// The state of source failover on the flow.
	//
	// If the state is inactive, the flow can have only one source. If the state is active, the flow can have one or two sources.
	State *string `field:"optional" json:"state" yaml:"state"`
}
