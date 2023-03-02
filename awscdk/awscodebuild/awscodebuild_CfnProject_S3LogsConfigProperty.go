package awscodebuild


// `S3Logs` is a property of the [AWS CodeBuild Project LogsConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-logsconfig.html) property type that specifies settings for logs generated by an AWS CodeBuild build in an S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LogsConfigProperty := &s3LogsConfigProperty{
//   	status: jsii.String("status"),
//
//   	// the properties below are optional
//   	encryptionDisabled: jsii.Boolean(false),
//   	location: jsii.String("location"),
//   }
//
type CfnProject_S3LogsConfigProperty struct {
	// The current status of the S3 build logs. Valid values are:.
	//
	// - `ENABLED` : S3 build logs are enabled for this build project.
	// - `DISABLED` : S3 build logs are not enabled for this build project.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Set to true if you do not want your S3 build log output encrypted.
	//
	// By default S3 build logs are encrypted.
	EncryptionDisabled interface{} `field:"optional" json:"encryptionDisabled" yaml:"encryptionDisabled"`
	// The ARN of an S3 bucket and the path prefix for S3 logs.
	//
	// If your Amazon S3 bucket name is `my-bucket` , and your path prefix is `build-log` , then acceptable formats are `my-bucket/build-log` or `arn:aws:s3:::my-bucket/build-log` .
	Location *string `field:"optional" json:"location" yaml:"location"`
}

