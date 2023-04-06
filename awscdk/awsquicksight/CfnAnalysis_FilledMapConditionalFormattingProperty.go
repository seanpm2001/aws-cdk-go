package awsquicksight


// The conditional formatting of a `FilledMapVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filledMapConditionalFormattingProperty := &FilledMapConditionalFormattingProperty{
//   	ConditionalFormattingOptions: []interface{}{
//   		&FilledMapConditionalFormattingOptionProperty{
//   			Shape: &FilledMapShapeConditionalFormattingProperty{
//   				FieldId: jsii.String("fieldId"),
//
//   				// the properties below are optional
//   				Format: &ShapeConditionalFormatProperty{
//   					BackgroundColor: &ConditionalFormattingColorProperty{
//   						Gradient: &ConditionalFormattingGradientColorProperty{
//   							Color: &GradientColorProperty{
//   								Stops: []interface{}{
//   									&GradientStopProperty{
//   										GradientOffset: jsii.Number(123),
//
//   										// the properties below are optional
//   										Color: jsii.String("color"),
//   										DataValue: jsii.Number(123),
//   									},
//   								},
//   							},
//   							Expression: jsii.String("expression"),
//   						},
//   						Solid: &ConditionalFormattingSolidColorProperty{
//   							Expression: jsii.String("expression"),
//
//   							// the properties below are optional
//   							Color: jsii.String("color"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnAnalysis_FilledMapConditionalFormattingProperty struct {
	// Conditional formatting options of a `FilledMapVisual` .
	ConditionalFormattingOptions interface{} `field:"required" json:"conditionalFormattingOptions" yaml:"conditionalFormattingOptions"`
}
