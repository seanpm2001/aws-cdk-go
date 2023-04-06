package awsquicksight


// The configuration of custom values for the destination parameter in `DestinationParameterValueConfiguration` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customValuesConfigurationProperty := &CustomValuesConfigurationProperty{
//   	CustomValues: &CustomParameterValuesProperty{
//   		DateTimeValues: []*string{
//   			jsii.String("dateTimeValues"),
//   		},
//   		DecimalValues: []interface{}{
//   			jsii.Number(123),
//   		},
//   		IntegerValues: []interface{}{
//   			jsii.Number(123),
//   		},
//   		StringValues: []*string{
//   			jsii.String("stringValues"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	IncludeNullValue: jsii.Boolean(false),
//   }
//
type CfnAnalysis_CustomValuesConfigurationProperty struct {
	// `CfnAnalysis.CustomValuesConfigurationProperty.CustomValues`.
	CustomValues interface{} `field:"required" json:"customValues" yaml:"customValues"`
	// Includes the null value in custom action parameter values.
	IncludeNullValue interface{} `field:"optional" json:"includeNullValue" yaml:"includeNullValue"`
}
