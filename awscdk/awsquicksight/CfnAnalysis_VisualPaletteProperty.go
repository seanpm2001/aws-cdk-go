package awsquicksight


// The visual display options for the visual palette.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   visualPaletteProperty := &VisualPaletteProperty{
//   	ChartColor: jsii.String("chartColor"),
//   	ColorMap: []interface{}{
//   		&DataPathColorProperty{
//   			Color: jsii.String("color"),
//   			Element: &DataPathValueProperty{
//   				FieldId: jsii.String("fieldId"),
//   				FieldValue: jsii.String("fieldValue"),
//   			},
//
//   			// the properties below are optional
//   			TimeGranularity: jsii.String("timeGranularity"),
//   		},
//   	},
//   }
//
type CfnAnalysis_VisualPaletteProperty struct {
	// The chart color options for the visual palette.
	ChartColor *string `field:"optional" json:"chartColor" yaml:"chartColor"`
	// The color map options for the visual palette.
	ColorMap interface{} `field:"optional" json:"colorMap" yaml:"colorMap"`
}

