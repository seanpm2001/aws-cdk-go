package awsquicksight


// Determines the row alternate color options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rowAlternateColorOptionsProperty := &RowAlternateColorOptionsProperty{
//   	RowAlternateColors: []*string{
//   		jsii.String("rowAlternateColors"),
//   	},
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-rowalternatecoloroptions.html
//
type CfnDashboard_RowAlternateColorOptionsProperty struct {
	// Determines the list of row alternate colors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-rowalternatecoloroptions.html#cfn-quicksight-dashboard-rowalternatecoloroptions-rowalternatecolors
	//
	RowAlternateColors *[]*string `field:"optional" json:"rowAlternateColors" yaml:"rowAlternateColors"`
	// Determines the widget status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-rowalternatecoloroptions.html#cfn-quicksight-dashboard-rowalternatecoloroptions-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

