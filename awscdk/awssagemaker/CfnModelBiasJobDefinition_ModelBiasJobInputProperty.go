package awssagemaker


// Inputs for the model bias job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var json interface{}
//
//   modelBiasJobInputProperty := &ModelBiasJobInputProperty{
//   	GroundTruthS3Input: &MonitoringGroundTruthS3InputProperty{
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//
//   	// the properties below are optional
//   	BatchTransformInput: &BatchTransformInputProperty{
//   		DataCapturedDestinationS3Uri: jsii.String("dataCapturedDestinationS3Uri"),
//   		DatasetFormat: &DatasetFormatProperty{
//   			Csv: &CsvProperty{
//   				Header: jsii.Boolean(false),
//   			},
//   			Json: json,
//   			Parquet: jsii.Boolean(false),
//   		},
//   		LocalPath: jsii.String("localPath"),
//
//   		// the properties below are optional
//   		EndTimeOffset: jsii.String("endTimeOffset"),
//   		FeaturesAttribute: jsii.String("featuresAttribute"),
//   		InferenceAttribute: jsii.String("inferenceAttribute"),
//   		ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   		ProbabilityThresholdAttribute: jsii.Number(123),
//   		S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		S3InputMode: jsii.String("s3InputMode"),
//   		StartTimeOffset: jsii.String("startTimeOffset"),
//   	},
//   	EndpointInput: &EndpointInputProperty{
//   		EndpointName: jsii.String("endpointName"),
//   		LocalPath: jsii.String("localPath"),
//
//   		// the properties below are optional
//   		EndTimeOffset: jsii.String("endTimeOffset"),
//   		FeaturesAttribute: jsii.String("featuresAttribute"),
//   		InferenceAttribute: jsii.String("inferenceAttribute"),
//   		ProbabilityAttribute: jsii.String("probabilityAttribute"),
//   		ProbabilityThresholdAttribute: jsii.Number(123),
//   		S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   		S3InputMode: jsii.String("s3InputMode"),
//   		StartTimeOffset: jsii.String("startTimeOffset"),
//   	},
//   }
//
type CfnModelBiasJobDefinition_ModelBiasJobInputProperty struct {
	// Location of ground truth labels to use in model bias job.
	GroundTruthS3Input interface{} `field:"required" json:"groundTruthS3Input" yaml:"groundTruthS3Input"`
	// `CfnModelBiasJobDefinition.ModelBiasJobInputProperty.BatchTransformInput`.
	BatchTransformInput interface{} `field:"optional" json:"batchTransformInput" yaml:"batchTransformInput"`
	// Input object for the endpoint.
	EndpointInput interface{} `field:"optional" json:"endpointInput" yaml:"endpointInput"`
}
