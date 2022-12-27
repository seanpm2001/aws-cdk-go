package awspipes


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeSourceKinesisStreamParametersProperty := &pipeSourceKinesisStreamParametersProperty{
//   	startingPosition: jsii.String("startingPosition"),
//
//   	// the properties below are optional
//   	batchSize: jsii.Number(123),
//   	deadLetterConfig: &deadLetterConfigProperty{
//   		arn: jsii.String("arn"),
//   	},
//   	maximumBatchingWindowInSeconds: jsii.Number(123),
//   	maximumRecordAgeInSeconds: jsii.Number(123),
//   	maximumRetryAttempts: jsii.Number(123),
//   	onPartialBatchItemFailure: jsii.String("onPartialBatchItemFailure"),
//   	parallelizationFactor: jsii.Number(123),
//   	startingPositionTimestamp: jsii.String("startingPositionTimestamp"),
//   }
//
type CfnPipe_PipeSourceKinesisStreamParametersProperty struct {
	// `CfnPipe.PipeSourceKinesisStreamParametersProperty.StartingPosition`.
	StartingPosition *string `field:"required" json:"startingPosition" yaml:"startingPosition"`
	// `CfnPipe.PipeSourceKinesisStreamParametersProperty.BatchSize`.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// `CfnPipe.PipeSourceKinesisStreamParametersProperty.DeadLetterConfig`.
	DeadLetterConfig interface{} `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// `CfnPipe.PipeSourceKinesisStreamParametersProperty.MaximumBatchingWindowInSeconds`.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// `CfnPipe.PipeSourceKinesisStreamParametersProperty.MaximumRecordAgeInSeconds`.
	MaximumRecordAgeInSeconds *float64 `field:"optional" json:"maximumRecordAgeInSeconds" yaml:"maximumRecordAgeInSeconds"`
	// `CfnPipe.PipeSourceKinesisStreamParametersProperty.MaximumRetryAttempts`.
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
	// `CfnPipe.PipeSourceKinesisStreamParametersProperty.OnPartialBatchItemFailure`.
	OnPartialBatchItemFailure *string `field:"optional" json:"onPartialBatchItemFailure" yaml:"onPartialBatchItemFailure"`
	// `CfnPipe.PipeSourceKinesisStreamParametersProperty.ParallelizationFactor`.
	ParallelizationFactor *float64 `field:"optional" json:"parallelizationFactor" yaml:"parallelizationFactor"`
	// `CfnPipe.PipeSourceKinesisStreamParametersProperty.StartingPositionTimestamp`.
	StartingPositionTimestamp *string `field:"optional" json:"startingPositionTimestamp" yaml:"startingPositionTimestamp"`
}

