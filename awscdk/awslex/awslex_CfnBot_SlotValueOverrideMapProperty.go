package awslex


// Maps a slot name to the [SlotValueOverride](https://docs.aws.amazon.com/lexv2/latest/APIReference/API_SlotValueOverride.html) object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var slotValueOverrideProperty_ slotValueOverrideProperty
//
//   slotValueOverrideMapProperty := &slotValueOverrideMapProperty{
//   	slotName: jsii.String("slotName"),
//   	slotValueOverride: &slotValueOverrideProperty{
//   		shape: jsii.String("shape"),
//   		value: &slotValueProperty{
//   			interpretedValue: jsii.String("interpretedValue"),
//   		},
//   		values: []interface{}{
//   			slotValueOverrideProperty_,
//   		},
//   	},
//   }
//
type CfnBot_SlotValueOverrideMapProperty struct {
	// The name of the slot.
	SlotName *string `field:"optional" json:"slotName" yaml:"slotName"`
	// The SlotValueOverride object to which the slot name will be mapped.
	SlotValueOverride interface{} `field:"optional" json:"slotValueOverride" yaml:"slotValueOverride"`
}

