package awsdlm


// *[Event-based policies only]* Specifies an action for an event-based policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &ActionProperty{
//   	CrossRegionCopy: []interface{}{
//   		&CrossRegionCopyActionProperty{
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				Encrypted: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				CmkArn: jsii.String("cmkArn"),
//   			},
//   			Target: jsii.String("target"),
//
//   			// the properties below are optional
//   			RetainRule: &CrossRegionCopyRetainRuleProperty{
//   				Interval: jsii.Number(123),
//   				IntervalUnit: jsii.String("intervalUnit"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
type CfnLifecyclePolicy_ActionProperty struct {
	// The rule for copying shared snapshots across Regions.
	CrossRegionCopy interface{} `field:"required" json:"crossRegionCopy" yaml:"crossRegionCopy"`
	// A descriptive name for the action.
	Name *string `field:"required" json:"name" yaml:"name"`
}
