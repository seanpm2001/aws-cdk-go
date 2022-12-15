package awsfis


// Specifies the configuration for experiment logging.
//
// For more information, see [Experiment logging](https://docs.aws.amazon.com/fis/latest/userguide/monitoring-logging.html) in the *AWS Fault Injection Simulator User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   experimentTemplateLogConfigurationProperty := &experimentTemplateLogConfigurationProperty{
//   	logSchemaVersion: jsii.Number(123),
//
//   	// the properties below are optional
//   	cloudWatchLogsConfiguration: &cloudWatchLogsConfigurationProperty{
//   		logGroupArn: jsii.String("logGroupArn"),
//   	},
//   	s3Configuration: &s3ConfigurationProperty{
//   		bucketName: jsii.String("bucketName"),
//
//   		// the properties below are optional
//   		prefix: jsii.String("prefix"),
//   	},
//   }
//
type CfnExperimentTemplate_ExperimentTemplateLogConfigurationProperty struct {
	// The schema version.
	//
	// The supported value is 1.
	LogSchemaVersion *float64 `field:"required" json:"logSchemaVersion" yaml:"logSchemaVersion"`
	// The configuration for experiment logging to Amazon CloudWatch Logs. The supported field is `LogGroupArn` . For example:.
	//
	// `{"LogGroupArn": "aws:arn:logs: *region_name* : *account_id* :log-group: *log_group_name* "}`.
	CloudWatchLogsConfiguration interface{} `field:"optional" json:"cloudWatchLogsConfiguration" yaml:"cloudWatchLogsConfiguration"`
	// The configuration for experiment logging to Amazon S3. The following fields are supported:.
	//
	// - `bucketName` - The name of the destination bucket.
	// - `prefix` - An optional bucket prefix.
	//
	// For example:
	//
	// `{"BucketName": " *my-s3-bucket* ", "Prefix": " *log-folder* "}`.
	S3Configuration interface{} `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

