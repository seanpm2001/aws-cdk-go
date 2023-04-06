package awsquicksight


// The range options for the data zoom scroll bar.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   visibleRangeOptionsProperty := &VisibleRangeOptionsProperty{
//   	PercentRange: &PercentVisibleRangeProperty{
//   		From: jsii.Number(123),
//   		To: jsii.Number(123),
//   	},
//   }
//
type CfnTemplate_VisibleRangeOptionsProperty struct {
	// The percent range in the visible range.
	PercentRange interface{} `field:"optional" json:"percentRange" yaml:"percentRange"`
}
