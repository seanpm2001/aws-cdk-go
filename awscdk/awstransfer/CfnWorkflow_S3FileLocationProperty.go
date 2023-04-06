package awstransfer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3FileLocationProperty := &S3FileLocationProperty{
//   	S3FileLocation: &S3InputFileLocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   	},
//   }
//
type CfnWorkflow_S3FileLocationProperty struct {
	// `CfnWorkflow.S3FileLocationProperty.S3FileLocation`.
	S3FileLocation interface{} `field:"optional" json:"s3FileLocation" yaml:"s3FileLocation"`
}
