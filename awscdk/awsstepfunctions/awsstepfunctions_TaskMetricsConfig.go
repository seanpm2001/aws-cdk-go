package awsstepfunctions


// Task Metrics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var metricDimensions interface{}
//
//   taskMetricsConfig := &taskMetricsConfig{
//   	metricDimensions: map[string]interface{}{
//   		"metricDimensionsKey": metricDimensions,
//   	},
//   	metricPrefixPlural: jsii.String("metricPrefixPlural"),
//   	metricPrefixSingular: jsii.String("metricPrefixSingular"),
//   }
//
// Experimental.
type TaskMetricsConfig struct {
	// The dimensions to attach to metrics.
	// Experimental.
	MetricDimensions *map[string]interface{} `field:"optional" json:"metricDimensions" yaml:"metricDimensions"`
	// Prefix for plural metric names of activity actions.
	// Experimental.
	MetricPrefixPlural *string `field:"optional" json:"metricPrefixPlural" yaml:"metricPrefixPlural"`
	// Prefix for singular metric names of activity actions.
	// Experimental.
	MetricPrefixSingular *string `field:"optional" json:"metricPrefixSingular" yaml:"metricPrefixSingular"`
}
