package awsquicksight


// A list of selectable values that are used in a control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterSelectableValuesProperty := &FilterSelectableValuesProperty{
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnDashboard_FilterSelectableValuesProperty struct {
	// The values that are used in the `FilterSelectableValues` .
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

