package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var json interface{}
//
//   datasetFormatProperty := &DatasetFormatProperty{
//   	Csv: &CsvProperty{
//   		Header: jsii.Boolean(false),
//   	},
//   	Json: json,
//   	Parquet: jsii.Boolean(false),
//   }
//
type CfnMonitoringSchedule_DatasetFormatProperty struct {
	// `CfnMonitoringSchedule.DatasetFormatProperty.Csv`.
	Csv interface{} `field:"optional" json:"csv" yaml:"csv"`
	// `CfnMonitoringSchedule.DatasetFormatProperty.Json`.
	Json interface{} `field:"optional" json:"json" yaml:"json"`
	// `CfnMonitoringSchedule.DatasetFormatProperty.Parquet`.
	Parquet interface{} `field:"optional" json:"parquet" yaml:"parquet"`
}
