package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   decimalDatasetParameterProperty := &DecimalDatasetParameterProperty{
//   	Id: jsii.String("id"),
//   	Name: jsii.String("name"),
//   	ValueType: jsii.String("valueType"),
//
//   	// the properties below are optional
//   	DefaultValues: &DecimalDatasetParameterDefaultValuesProperty{
//   		StaticValues: []interface{}{
//   			jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnDataSet_DecimalDatasetParameterProperty struct {
	// `CfnDataSet.DecimalDatasetParameterProperty.Id`.
	Id *string `field:"required" json:"id" yaml:"id"`
	// `CfnDataSet.DecimalDatasetParameterProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnDataSet.DecimalDatasetParameterProperty.ValueType`.
	ValueType *string `field:"required" json:"valueType" yaml:"valueType"`
	// `CfnDataSet.DecimalDatasetParameterProperty.DefaultValues`.
	DefaultValues interface{} `field:"optional" json:"defaultValues" yaml:"defaultValues"`
}
