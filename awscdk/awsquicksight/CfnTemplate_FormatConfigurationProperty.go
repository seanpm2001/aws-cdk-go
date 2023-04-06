package awsquicksight


// The formatting configuration for all types of field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   formatConfigurationProperty := &FormatConfigurationProperty{
//   	DateTimeFormatConfiguration: &DateTimeFormatConfigurationProperty{
//   		DateTimeFormat: jsii.String("dateTimeFormat"),
//   		NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   			NullString: jsii.String("nullString"),
//   		},
//   		NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   			CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   				DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   					DecimalPlaces: jsii.Number(123),
//   				},
//   				NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   					DisplayMode: jsii.String("displayMode"),
//   				},
//   				NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   					NullString: jsii.String("nullString"),
//   				},
//   				NumberScale: jsii.String("numberScale"),
//   				Prefix: jsii.String("prefix"),
//   				SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   					DecimalSeparator: jsii.String("decimalSeparator"),
//   					ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   						Symbol: jsii.String("symbol"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   				Suffix: jsii.String("suffix"),
//   				Symbol: jsii.String("symbol"),
//   			},
//   			NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   				DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   					DecimalPlaces: jsii.Number(123),
//   				},
//   				NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   					DisplayMode: jsii.String("displayMode"),
//   				},
//   				NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   					NullString: jsii.String("nullString"),
//   				},
//   				NumberScale: jsii.String("numberScale"),
//   				Prefix: jsii.String("prefix"),
//   				SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   					DecimalSeparator: jsii.String("decimalSeparator"),
//   					ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   						Symbol: jsii.String("symbol"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   				Suffix: jsii.String("suffix"),
//   			},
//   			PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   				DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   					DecimalPlaces: jsii.Number(123),
//   				},
//   				NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   					DisplayMode: jsii.String("displayMode"),
//   				},
//   				NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   					NullString: jsii.String("nullString"),
//   				},
//   				Prefix: jsii.String("prefix"),
//   				SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   					DecimalSeparator: jsii.String("decimalSeparator"),
//   					ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   						Symbol: jsii.String("symbol"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   				Suffix: jsii.String("suffix"),
//   			},
//   		},
//   	},
//   	NumberFormatConfiguration: &NumberFormatConfigurationProperty{
//   		FormatConfiguration: &NumericFormatConfigurationProperty{
//   			CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   				DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   					DecimalPlaces: jsii.Number(123),
//   				},
//   				NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   					DisplayMode: jsii.String("displayMode"),
//   				},
//   				NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   					NullString: jsii.String("nullString"),
//   				},
//   				NumberScale: jsii.String("numberScale"),
//   				Prefix: jsii.String("prefix"),
//   				SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   					DecimalSeparator: jsii.String("decimalSeparator"),
//   					ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   						Symbol: jsii.String("symbol"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   				Suffix: jsii.String("suffix"),
//   				Symbol: jsii.String("symbol"),
//   			},
//   			NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   				DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   					DecimalPlaces: jsii.Number(123),
//   				},
//   				NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   					DisplayMode: jsii.String("displayMode"),
//   				},
//   				NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   					NullString: jsii.String("nullString"),
//   				},
//   				NumberScale: jsii.String("numberScale"),
//   				Prefix: jsii.String("prefix"),
//   				SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   					DecimalSeparator: jsii.String("decimalSeparator"),
//   					ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   						Symbol: jsii.String("symbol"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   				Suffix: jsii.String("suffix"),
//   			},
//   			PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   				DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   					DecimalPlaces: jsii.Number(123),
//   				},
//   				NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   					DisplayMode: jsii.String("displayMode"),
//   				},
//   				NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   					NullString: jsii.String("nullString"),
//   				},
//   				Prefix: jsii.String("prefix"),
//   				SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   					DecimalSeparator: jsii.String("decimalSeparator"),
//   					ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   						Symbol: jsii.String("symbol"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   				Suffix: jsii.String("suffix"),
//   			},
//   		},
//   	},
//   	StringFormatConfiguration: &StringFormatConfigurationProperty{
//   		NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   			NullString: jsii.String("nullString"),
//   		},
//   		NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   			CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   				DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   					DecimalPlaces: jsii.Number(123),
//   				},
//   				NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   					DisplayMode: jsii.String("displayMode"),
//   				},
//   				NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   					NullString: jsii.String("nullString"),
//   				},
//   				NumberScale: jsii.String("numberScale"),
//   				Prefix: jsii.String("prefix"),
//   				SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   					DecimalSeparator: jsii.String("decimalSeparator"),
//   					ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   						Symbol: jsii.String("symbol"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   				Suffix: jsii.String("suffix"),
//   				Symbol: jsii.String("symbol"),
//   			},
//   			NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   				DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   					DecimalPlaces: jsii.Number(123),
//   				},
//   				NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   					DisplayMode: jsii.String("displayMode"),
//   				},
//   				NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   					NullString: jsii.String("nullString"),
//   				},
//   				NumberScale: jsii.String("numberScale"),
//   				Prefix: jsii.String("prefix"),
//   				SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   					DecimalSeparator: jsii.String("decimalSeparator"),
//   					ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   						Symbol: jsii.String("symbol"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   				Suffix: jsii.String("suffix"),
//   			},
//   			PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   				DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   					DecimalPlaces: jsii.Number(123),
//   				},
//   				NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   					DisplayMode: jsii.String("displayMode"),
//   				},
//   				NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   					NullString: jsii.String("nullString"),
//   				},
//   				Prefix: jsii.String("prefix"),
//   				SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   					DecimalSeparator: jsii.String("decimalSeparator"),
//   					ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   						Symbol: jsii.String("symbol"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   				Suffix: jsii.String("suffix"),
//   			},
//   		},
//   	},
//   }
//
type CfnTemplate_FormatConfigurationProperty struct {
	// Formatting configuration for `DateTime` fields.
	DateTimeFormatConfiguration interface{} `field:"optional" json:"dateTimeFormatConfiguration" yaml:"dateTimeFormatConfiguration"`
	// Formatting configuration for number fields.
	NumberFormatConfiguration interface{} `field:"optional" json:"numberFormatConfiguration" yaml:"numberFormatConfiguration"`
	// Formatting configuration for string fields.
	StringFormatConfiguration interface{} `field:"optional" json:"stringFormatConfiguration" yaml:"stringFormatConfiguration"`
}

