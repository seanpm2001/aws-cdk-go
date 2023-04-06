package awslex


// The value to set in a slot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slotValueProperty := &SlotValueProperty{
//   	InterpretedValue: jsii.String("interpretedValue"),
//   }
//
type CfnBot_SlotValueProperty struct {
	// The value that Amazon Lex determines for the slot.
	//
	// The actual value depends on the setting of the value selection strategy for the bot. You can choose to use the value entered by the user, or you can have Amazon Lex choose the first value in the `resolvedValues` list.
	InterpretedValue *string `field:"optional" json:"interpretedValue" yaml:"interpretedValue"`
}

