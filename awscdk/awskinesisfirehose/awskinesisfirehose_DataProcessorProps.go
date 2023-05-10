package awskinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Configure the data processor.
//
// Example:
//   import path "github.com/aws-samples/dummy/path"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import lambdanodejs "github.com/aws/aws-cdk-go/awscdk"
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   app := cdk.NewApp()
//
//   stack := cdk.NewStack(app, jsii.String("aws-cdk-firehose-delivery-stream-s3-all-properties"))
//
//   bucket := s3.NewBucket(stack, jsii.String("Bucket"), &BucketProps{
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   	AutoDeleteObjects: jsii.Boolean(true),
//   })
//
//   backupBucket := s3.NewBucket(stack, jsii.String("BackupBucket"), &BucketProps{
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   	AutoDeleteObjects: jsii.Boolean(true),
//   })
//   logGroup := logs.NewLogGroup(stack, jsii.String("LogGroup"), &LogGroupProps{
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   })
//
//   dataProcessorFunction := lambdanodejs.NewNodejsFunction(stack, jsii.String("DataProcessorFunction"), &NodejsFunctionProps{
//   	Entry: path.join(__dirname, jsii.String("lambda-data-processor.js")),
//   	Timeout: cdk.Duration_Minutes(jsii.Number(1)),
//   })
//
//   processor := firehose.NewLambdaFunctionProcessor(dataProcessorFunction, &DataProcessorProps{
//   	BufferInterval: cdk.Duration_Seconds(jsii.Number(60)),
//   	BufferSize: cdk.Size_Mebibytes(jsii.Number(1)),
//   	Retries: jsii.Number(1),
//   })
//
//   key := kms.NewKey(stack, jsii.String("Key"), &KeyProps{
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   })
//
//   backupKey := kms.NewKey(stack, jsii.String("BackupKey"), &KeyProps{
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   })
//
//   firehose.NewDeliveryStream(stack, jsii.String("Delivery Stream"), &DeliveryStreamProps{
//   	Destinations: []iDestination{
//   		destinations.NewS3Bucket(bucket, &S3BucketProps{
//   			Logging: jsii.Boolean(true),
//   			LogGroup: logGroup,
//   			Processor: processor,
//   			Compression: destinations.Compression_GZIP(),
//   			DataOutputPrefix: jsii.String("regularPrefix"),
//   			ErrorOutputPrefix: jsii.String("errorPrefix"),
//   			BufferingInterval: cdk.Duration_*Seconds(jsii.Number(60)),
//   			BufferingSize: cdk.Size_*Mebibytes(jsii.Number(1)),
//   			EncryptionKey: key,
//   			S3Backup: &DestinationS3BackupProps{
//   				Mode: destinations.BackupMode_ALL,
//   				Bucket: backupBucket,
//   				Compression: destinations.Compression_ZIP(),
//   				DataOutputPrefix: jsii.String("backupPrefix"),
//   				ErrorOutputPrefix: jsii.String("backupErrorPrefix"),
//   				BufferingInterval: cdk.Duration_*Seconds(jsii.Number(60)),
//   				BufferingSize: cdk.Size_*Mebibytes(jsii.Number(1)),
//   				EncryptionKey: backupKey,
//   			},
//   		}),
//   	},
//   })
//
//   app.Synth()
//
// Experimental.
type DataProcessorProps struct {
	// The length of time Kinesis Data Firehose will buffer incoming data before calling the processor.
	//
	// s.
	// Experimental.
	BufferInterval awscdk.Duration `field:"optional" json:"bufferInterval" yaml:"bufferInterval"`
	// The amount of incoming data Kinesis Data Firehose will buffer before calling the processor.
	// Experimental.
	BufferSize awscdk.Size `field:"optional" json:"bufferSize" yaml:"bufferSize"`
	// The number of times Kinesis Data Firehose will retry the processor invocation after a failure due to network timeout or invocation limits.
	// Experimental.
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
}

