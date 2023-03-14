package awspersonalize


// Describes the properties for hyperparameter optimization (HPO).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hpoConfigProperty := &HpoConfigProperty{
//   	AlgorithmHyperParameterRanges: &AlgorithmHyperParameterRangesProperty{
//   		CategoricalHyperParameterRanges: []interface{}{
//   			&CategoricalHyperParameterRangeProperty{
//   				Name: jsii.String("name"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   		ContinuousHyperParameterRanges: []interface{}{
//   			&ContinuousHyperParameterRangeProperty{
//   				MaxValue: jsii.Number(123),
//   				MinValue: jsii.Number(123),
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		IntegerHyperParameterRanges: []interface{}{
//   			&IntegerHyperParameterRangeProperty{
//   				MaxValue: jsii.Number(123),
//   				MinValue: jsii.Number(123),
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	HpoObjective: &HpoObjectiveProperty{
//   		MetricName: jsii.String("metricName"),
//   		MetricRegex: jsii.String("metricRegex"),
//   		Type: jsii.String("type"),
//   	},
//   	HpoResourceConfig: &HpoResourceConfigProperty{
//   		MaxNumberOfTrainingJobs: jsii.String("maxNumberOfTrainingJobs"),
//   		MaxParallelTrainingJobs: jsii.String("maxParallelTrainingJobs"),
//   	},
//   }
//
type CfnSolution_HpoConfigProperty struct {
	// The hyperparameters and their allowable ranges.
	AlgorithmHyperParameterRanges interface{} `field:"optional" json:"algorithmHyperParameterRanges" yaml:"algorithmHyperParameterRanges"`
	// The metric to optimize during HPO.
	//
	// > Amazon Personalize doesn't support configuring the `hpoObjective` at this time.
	HpoObjective interface{} `field:"optional" json:"hpoObjective" yaml:"hpoObjective"`
	// Describes the resource configuration for HPO.
	HpoResourceConfig interface{} `field:"optional" json:"hpoResourceConfig" yaml:"hpoResourceConfig"`
}

