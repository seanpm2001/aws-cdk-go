package awscloudwatch


// This structure is used in both `GetMetricData` and `PutMetricAlarm` .
//
// The supported use of this structure is different for those two operations.
//
// When used in `GetMetricData` , it indicates the metric data to return, and whether this call is just retrieving a batch set of data for one metric, or is performing a Metrics Insights query or a math expression. A single `GetMetricData` call can include up to 500 `MetricDataQuery` structures.
//
// When used in `PutMetricAlarm` , it enables you to create an alarm based on a metric math expression. Each `MetricDataQuery` in the array specifies either a metric to retrieve, or a math expression to be performed on retrieved metrics. A single `PutMetricAlarm` call can include up to 20 `MetricDataQuery` structures in the array. The 20 structures can include as many as 10 structures that contain a `MetricStat` parameter to retrieve a metric, and as many as 10 structures that contain the `Expression` parameter to perform a math expression. Of those `Expression` structures, one must have `True` as the value for `ReturnData` . The result of this expression is the value the alarm watches.
//
// Any expression used in a `PutMetricAlarm` operation must return a single time series. For more information, see [Metric Math Syntax and Functions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-metric-math.html#metric-math-syntax) in the *Amazon CloudWatch User Guide* .
//
// Some of the parameters of this structure also have different uses whether you are using this structure in a `GetMetricData` operation or a `PutMetricAlarm` operation. These differences are explained in the following parameter list.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDataQueryProperty := &metricDataQueryProperty{
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	accountId: jsii.String("accountId"),
//   	expression: jsii.String("expression"),
//   	label: jsii.String("label"),
//   	metricStat: &metricStatProperty{
//   		metric: &metricProperty{
//   			metricName: jsii.String("metricName"),
//   			namespace: jsii.String("namespace"),
//
//   			// the properties below are optional
//   			dimensions: []interface{}{
//   				&dimensionProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		period: jsii.Number(123),
//   		stat: jsii.String("stat"),
//
//   		// the properties below are optional
//   		unit: jsii.String("unit"),
//   	},
//   	period: jsii.Number(123),
//   	returnData: jsii.Boolean(false),
//   }
//
type CfnAnomalyDetector_MetricDataQueryProperty struct {
	// A short name used to tie this object to the results in the response.
	//
	// This name must be unique within a single call to `GetMetricData` . If you are performing math expressions on this set of data, this name represents that data and can serve as a variable in the mathematical expression. The valid characters are letters, numbers, and underscore. The first character must be a lowercase letter.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The ID of the account where the metrics are located, if this is a cross-account alarm.
	//
	// Use this field only for `PutMetricAlarm` operations. It is not used in `GetMetricData` operations.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// This field can contain either a Metrics Insights query, or a metric math expression to be performed on the returned data.
	//
	// For more information about Metrics Insights queries, see [Metrics Insights query components and syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch-metrics-insights-querylanguage) in the *Amazon CloudWatch User Guide* .
	//
	// A math expression can use the `Id` of the other metrics or queries to refer to those metrics, and can also use the `Id` of other expressions to use the result of those expressions. For more information about metric math expressions, see [Metric Math Syntax and Functions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-metric-math.html#metric-math-syntax) in the *Amazon CloudWatch User Guide* .
	//
	// Within each MetricDataQuery object, you must specify either `Expression` or `MetricStat` but not both.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// A human-readable label for this metric or expression.
	//
	// This is especially useful if this is an expression, so that you know what the value represents. If the metric or expression is shown in a CloudWatch dashboard widget, the label is shown. If Label is omitted, CloudWatch generates a default.
	//
	// You can put dynamic expressions into a label, so that it is more descriptive. For more information, see [Using Dynamic Labels](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/graph-dynamic-labels.html) .
	Label *string `field:"optional" json:"label" yaml:"label"`
	// The metric to be returned, along with statistics, period, and units.
	//
	// Use this parameter only if this object is retrieving a metric and not performing a math expression on returned data.
	//
	// Within one MetricDataQuery object, you must specify either `Expression` or `MetricStat` but not both.
	MetricStat interface{} `field:"optional" json:"metricStat" yaml:"metricStat"`
	// The granularity, in seconds, of the returned data points.
	//
	// For metrics with regular resolution, a period can be as short as one minute (60 seconds) and must be a multiple of 60. For high-resolution metrics that are collected at intervals of less than one minute, the period can be 1, 5, 10, 30, 60, or any multiple of 60. High-resolution metrics are those metrics stored by a `PutMetricData` operation that includes a `StorageResolution of 1 second` .
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// When used in `GetMetricData` , this option indicates whether to return the timestamps and raw data values of this metric.
	//
	// If you are performing this call just to do math expressions and do not also need the raw data returned, you can specify `False` . If you omit this, the default of `True` is used.
	//
	// When used in `PutMetricAlarm` , specify `True` for the one expression result to use as the alarm. For all other metrics and expressions in the same `PutMetricAlarm` operation, specify `ReturnData` as False.
	ReturnData interface{} `field:"optional" json:"returnData" yaml:"returnData"`
}

