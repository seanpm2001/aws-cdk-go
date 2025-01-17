package awsquicksight


// The conditional formatting for a `PivotTableVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pivotTableConditionalFormattingProperty := &PivotTableConditionalFormattingProperty{
//   	ConditionalFormattingOptions: []interface{}{
//   		&PivotTableConditionalFormattingOptionProperty{
//   			Cell: &PivotTableCellConditionalFormattingProperty{
//   				FieldId: jsii.String("fieldId"),
//
//   				// the properties below are optional
//   				Scope: &PivotTableConditionalFormattingScopeProperty{
//   					Role: jsii.String("role"),
//   				},
//   				Scopes: []interface{}{
//   					&PivotTableConditionalFormattingScopeProperty{
//   						Role: jsii.String("role"),
//   					},
//   				},
//   				TextFormat: &TextConditionalFormatProperty{
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
//   					Icon: &ConditionalFormattingIconProperty{
//   						CustomCondition: &ConditionalFormattingCustomIconConditionProperty{
//   							Expression: jsii.String("expression"),
//   							IconOptions: &ConditionalFormattingCustomIconOptionsProperty{
//   								Icon: jsii.String("icon"),
//   								UnicodeIcon: jsii.String("unicodeIcon"),
//   							},
//
//   							// the properties below are optional
//   							Color: jsii.String("color"),
//   							DisplayConfiguration: &ConditionalFormattingIconDisplayConfigurationProperty{
//   								IconDisplayOption: jsii.String("iconDisplayOption"),
//   							},
//   						},
//   						IconSet: &ConditionalFormattingIconSetProperty{
//   							Expression: jsii.String("expression"),
//
//   							// the properties below are optional
//   							IconSetType: jsii.String("iconSetType"),
//   						},
//   					},
//   					TextColor: &ConditionalFormattingColorProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableconditionalformatting.html
//
type CfnDashboard_PivotTableConditionalFormattingProperty struct {
	// Conditional formatting options for a `PivotTableVisual` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottableconditionalformatting.html#cfn-quicksight-dashboard-pivottableconditionalformatting-conditionalformattingoptions
	//
	ConditionalFormattingOptions interface{} `field:"optional" json:"conditionalFormattingOptions" yaml:"conditionalFormattingOptions"`
}

