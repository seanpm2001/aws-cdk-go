package awssagemaker


// The Amazon Simple Storage (Amazon S3) location and and security configuration for `OfflineStore` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3StorageConfigProperty := &s3StorageConfigProperty{
//   	s3Uri: jsii.String("s3Uri"),
//
//   	// the properties below are optional
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnFeatureGroup_S3StorageConfigProperty struct {
	// The S3 URI, or location in Amazon S3, of `OfflineStore` .
	//
	// S3 URIs have a format similar to the following: `s3://example-bucket/prefix/` .
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// The AWS Key Management Service (KMS) key ID of the key used to encrypt any objects written into the `OfflineStore` S3 location.
	//
	// The IAM `roleARN` that is passed as a parameter to `CreateFeatureGroup` must have below permissions to the `KmsKeyId` :
	//
	// - `"kms:GenerateDataKey"`.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}
