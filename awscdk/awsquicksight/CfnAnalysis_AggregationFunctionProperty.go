package awsquicksight


// An aggregation function aggregates values from a dimension or measure.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aggregationFunctionProperty := &AggregationFunctionProperty{
//   	CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   	DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   	NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   		PercentileAggregation: &PercentileAggregationProperty{
//   			PercentileValue: jsii.Number(123),
//   		},
//   		SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   	},
//   }
//
type CfnAnalysis_AggregationFunctionProperty struct {
	// Aggregation for categorical values.
	//
	// - `COUNT` : Aggregate by the total number of values, including duplicates.
	// - `DISTINCT_COUNT` : Aggregate by the total number of distinct values.
	CategoricalAggregationFunction *string `field:"optional" json:"categoricalAggregationFunction" yaml:"categoricalAggregationFunction"`
	// Aggregation for date values.
	//
	// - `COUNT` : Aggregate by the total number of values, including duplicates.
	// - `DISTINCT_COUNT` : Aggregate by the total number of distinct values.
	// - `MIN` : Select the smallest date value.
	// - `MAX` : Select the largest date value.
	DateAggregationFunction *string `field:"optional" json:"dateAggregationFunction" yaml:"dateAggregationFunction"`
	// Aggregation for numerical values.
	NumericalAggregationFunction interface{} `field:"optional" json:"numericalAggregationFunction" yaml:"numericalAggregationFunction"`
}

