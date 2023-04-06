package awsiot


// A Device Defender security profile behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   behaviorProperty := &BehaviorProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Criteria: &BehaviorCriteriaProperty{
//   		ComparisonOperator: jsii.String("comparisonOperator"),
//   		ConsecutiveDatapointsToAlarm: jsii.Number(123),
//   		ConsecutiveDatapointsToClear: jsii.Number(123),
//   		DurationSeconds: jsii.Number(123),
//   		MlDetectionConfig: &MachineLearningDetectionConfigProperty{
//   			ConfidenceLevel: jsii.String("confidenceLevel"),
//   		},
//   		StatisticalThreshold: &StatisticalThresholdProperty{
//   			Statistic: jsii.String("statistic"),
//   		},
//   		Value: &MetricValueProperty{
//   			Cidrs: []*string{
//   				jsii.String("cidrs"),
//   			},
//   			Count: jsii.String("count"),
//   			Number: jsii.Number(123),
//   			Numbers: []interface{}{
//   				jsii.Number(123),
//   			},
//   			Ports: []interface{}{
//   				jsii.Number(123),
//   			},
//   			Strings: []*string{
//   				jsii.String("strings"),
//   			},
//   		},
//   	},
//   	Metric: jsii.String("metric"),
//   	MetricDimension: &MetricDimensionProperty{
//   		DimensionName: jsii.String("dimensionName"),
//
//   		// the properties below are optional
//   		Operator: jsii.String("operator"),
//   	},
//   	SuppressAlerts: jsii.Boolean(false),
//   }
//
type CfnSecurityProfile_BehaviorProperty struct {
	// The name you've given to the behavior.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The criteria that determine if a device is behaving normally in regard to the `metric` .
	Criteria interface{} `field:"optional" json:"criteria" yaml:"criteria"`
	// What is measured by the behavior.
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
	// The dimension of the metric.
	MetricDimension interface{} `field:"optional" json:"metricDimension" yaml:"metricDimension"`
	// The alert status.
	//
	// If you set the value to `true` , alerts will be suppressed.
	SuppressAlerts interface{} `field:"optional" json:"suppressAlerts" yaml:"suppressAlerts"`
}
