package awsquicksight


// The conditional formatting for the primary value of a KPI visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kPIPrimaryValueConditionalFormattingProperty := &KPIPrimaryValueConditionalFormattingProperty{
//   	Icon: &ConditionalFormattingIconProperty{
//   		CustomCondition: &ConditionalFormattingCustomIconConditionProperty{
//   			Expression: jsii.String("expression"),
//   			IconOptions: &ConditionalFormattingCustomIconOptionsProperty{
//   				Icon: jsii.String("icon"),
//   				UnicodeIcon: jsii.String("unicodeIcon"),
//   			},
//
//   			// the properties below are optional
//   			Color: jsii.String("color"),
//   			DisplayConfiguration: &ConditionalFormattingIconDisplayConfigurationProperty{
//   				IconDisplayOption: jsii.String("iconDisplayOption"),
//   			},
//   		},
//   		IconSet: &ConditionalFormattingIconSetProperty{
//   			Expression: jsii.String("expression"),
//
//   			// the properties below are optional
//   			IconSetType: jsii.String("iconSetType"),
//   		},
//   	},
//   	TextColor: &ConditionalFormattingColorProperty{
//   		Gradient: &ConditionalFormattingGradientColorProperty{
//   			Color: &GradientColorProperty{
//   				Stops: []interface{}{
//   					&GradientStopProperty{
//   						GradientOffset: jsii.Number(123),
//
//   						// the properties below are optional
//   						Color: jsii.String("color"),
//   						DataValue: jsii.Number(123),
//   					},
//   				},
//   			},
//   			Expression: jsii.String("expression"),
//   		},
//   		Solid: &ConditionalFormattingSolidColorProperty{
//   			Expression: jsii.String("expression"),
//
//   			// the properties below are optional
//   			Color: jsii.String("color"),
//   		},
//   	},
//   }
//
type CfnAnalysis_KPIPrimaryValueConditionalFormattingProperty struct {
	// The conditional formatting of the primary value's icon.
	Icon interface{} `field:"optional" json:"icon" yaml:"icon"`
	// The conditional formatting of the primary value's text color.
	TextColor interface{} `field:"optional" json:"textColor" yaml:"textColor"`
}

