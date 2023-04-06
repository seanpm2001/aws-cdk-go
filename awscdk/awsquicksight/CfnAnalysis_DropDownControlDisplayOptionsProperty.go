package awsquicksight


// The display options of a control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dropDownControlDisplayOptionsProperty := &DropDownControlDisplayOptionsProperty{
//   	SelectAllOptions: &ListControlSelectAllOptionsProperty{
//   		Visibility: jsii.String("visibility"),
//   	},
//   	TitleOptions: &LabelOptionsProperty{
//   		CustomLabel: jsii.String("customLabel"),
//   		FontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontSize: &FontSizeProperty{
//   				Relative: jsii.String("relative"),
//   			},
//   			FontStyle: jsii.String("fontStyle"),
//   			FontWeight: &FontWeightProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		Visibility: jsii.String("visibility"),
//   	},
//   }
//
type CfnAnalysis_DropDownControlDisplayOptionsProperty struct {
	// The configuration of the `Select all` options in a dropdown control.
	SelectAllOptions interface{} `field:"optional" json:"selectAllOptions" yaml:"selectAllOptions"`
	// The options to configure the title visibility, name, and font size.
	TitleOptions interface{} `field:"optional" json:"titleOptions" yaml:"titleOptions"`
}

