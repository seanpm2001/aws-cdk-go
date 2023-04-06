package awslex


// Defines the Amazon CloudWatch Logs destination log group for conversation text logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   textLogDestinationProperty := &TextLogDestinationProperty{
//   	CloudWatch: &CloudWatchLogGroupLogDestinationProperty{
//   		CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   		LogPrefix: jsii.String("logPrefix"),
//   	},
//   }
//
type CfnBot_TextLogDestinationProperty struct {
	// Defines the Amazon CloudWatch Logs log group where text and metadata logs are delivered.
	CloudWatch interface{} `field:"required" json:"cloudWatch" yaml:"cloudWatch"`
}

