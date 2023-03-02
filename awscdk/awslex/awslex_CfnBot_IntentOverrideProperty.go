package awslex


// Override settings to configure the intent state.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var slotValueOverrideProperty_ slotValueOverrideProperty
//
//   intentOverrideProperty := &intentOverrideProperty{
//   	name: jsii.String("name"),
//   	slots: []interface{}{
//   		&slotValueOverrideMapProperty{
//   			slotName: jsii.String("slotName"),
//   			slotValueOverride: &slotValueOverrideProperty{
//   				shape: jsii.String("shape"),
//   				value: &slotValueProperty{
//   					interpretedValue: jsii.String("interpretedValue"),
//   				},
//   				values: []interface{}{
//   					slotValueOverrideProperty_,
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnBot_IntentOverrideProperty struct {
	// The name of the intent.
	//
	// Only required when you're switching intents.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A map of all of the slot value overrides for the intent.
	//
	// The name of the slot maps to the value of the slot. Slots that are not included in the map aren't overridden.
	Slots interface{} `field:"optional" json:"slots" yaml:"slots"`
}

