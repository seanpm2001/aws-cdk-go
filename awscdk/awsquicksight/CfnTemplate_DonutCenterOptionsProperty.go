package awsquicksight


// The label options of the label that is displayed in the center of a donut chart.
//
// This option isn't available for pie charts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   donutCenterOptionsProperty := &DonutCenterOptionsProperty{
//   	LabelVisibility: jsii.String("labelVisibility"),
//   }
//
type CfnTemplate_DonutCenterOptionsProperty struct {
	// Determines the visibility of the label in a donut chart.
	//
	// In the Amazon QuickSight console, this option is called `'Show total'` .
	LabelVisibility *string `field:"optional" json:"labelVisibility" yaml:"labelVisibility"`
}

