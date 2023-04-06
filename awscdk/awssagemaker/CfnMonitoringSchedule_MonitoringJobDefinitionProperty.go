package awssagemaker


// Defines the monitoring job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var json interface{}
//
//   monitoringJobDefinitionProperty := &MonitoringJobDefinitionProperty{
//   	MonitoringAppSpecification: &MonitoringAppSpecificationProperty{
//   		ImageUri: jsii.String("imageUri"),
//
//   		// the properties below are optional
//   		ContainerArguments: []*string{
//   			jsii.String("containerArguments"),
//   		},
//   		ContainerEntrypoint: []*string{
//   			jsii.String("containerEntrypoint"),
//   		},
//   		PostAnalyticsProcessorSourceUri: jsii.String("postAnalyticsProcessorSourceUri"),
//   		RecordPreprocessorSourceUri: jsii.String("recordPreprocessorSourceUri"),
//   	},
//   	MonitoringInputs: []interface{}{
//   		&MonitoringInputProperty{
//   			BatchTransformInput: &BatchTransformInputProperty{
//   				DataCapturedDestinationS3Uri: jsii.String("dataCapturedDestinationS3Uri"),
//   				DatasetFormat: &DatasetFormatProperty{
//   					Csv: &CsvProperty{
//   						Header: jsii.Boolean(false),
//   					},
//   					Json: json,
//   					Parquet: jsii.Boolean(false),
//   				},
//   				LocalPath: jsii.String("localPath"),
//
//   				// the properties below are optional
//   				S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   				S3InputMode: jsii.String("s3InputMode"),
//   			},
//   			EndpointInput: &EndpointInputProperty{
//   				EndpointName: jsii.String("endpointName"),
//   				LocalPath: jsii.String("localPath"),
//
//   				// the properties below are optional
//   				S3DataDistributionType: jsii.String("s3DataDistributionType"),
//   				S3InputMode: jsii.String("s3InputMode"),
//   			},
//   		},
//   	},
//   	MonitoringOutputConfig: &MonitoringOutputConfigProperty{
//   		MonitoringOutputs: []interface{}{
//   			&MonitoringOutputProperty{
//   				S3Output: &S3OutputProperty{
//   					LocalPath: jsii.String("localPath"),
//   					S3Uri: jsii.String("s3Uri"),
//
//   					// the properties below are optional
//   					S3UploadMode: jsii.String("s3UploadMode"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	MonitoringResources: &MonitoringResourcesProperty{
//   		ClusterConfig: &ClusterConfigProperty{
//   			InstanceCount: jsii.Number(123),
//   			InstanceType: jsii.String("instanceType"),
//   			VolumeSizeInGb: jsii.Number(123),
//
//   			// the properties below are optional
//   			VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	BaselineConfig: &BaselineConfigProperty{
//   		ConstraintsResource: &ConstraintsResourceProperty{
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   		StatisticsResource: &StatisticsResourceProperty{
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	Environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	NetworkConfig: &NetworkConfigProperty{
//   		EnableInterContainerTrafficEncryption: jsii.Boolean(false),
//   		EnableNetworkIsolation: jsii.Boolean(false),
//   		VpcConfig: &VpcConfigProperty{
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	StoppingCondition: &StoppingConditionProperty{
//   		MaxRuntimeInSeconds: jsii.Number(123),
//   	},
//   }
//
type CfnMonitoringSchedule_MonitoringJobDefinitionProperty struct {
	// Configures the monitoring job to run a specified Docker container image.
	MonitoringAppSpecification interface{} `field:"required" json:"monitoringAppSpecification" yaml:"monitoringAppSpecification"`
	// The array of inputs for the monitoring job.
	//
	// Currently we support monitoring an Amazon SageMaker Endpoint.
	MonitoringInputs interface{} `field:"required" json:"monitoringInputs" yaml:"monitoringInputs"`
	// The array of outputs from the monitoring job to be uploaded to Amazon Simple Storage Service (Amazon S3).
	MonitoringOutputConfig interface{} `field:"required" json:"monitoringOutputConfig" yaml:"monitoringOutputConfig"`
	// Identifies the resources, ML compute instances, and ML storage volumes to deploy for a monitoring job.
	//
	// In distributed processing, you specify more than one instance.
	MonitoringResources interface{} `field:"required" json:"monitoringResources" yaml:"monitoringResources"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Baseline configuration used to validate that the data conforms to the specified constraints and statistics.
	BaselineConfig interface{} `field:"optional" json:"baselineConfig" yaml:"baselineConfig"`
	// Sets the environment variables in the Docker container.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// Specifies networking options for an monitoring job.
	NetworkConfig interface{} `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// Specifies a time limit for how long the monitoring job is allowed to run.
	StoppingCondition interface{} `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
}

