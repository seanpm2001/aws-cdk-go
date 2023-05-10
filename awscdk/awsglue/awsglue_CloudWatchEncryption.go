package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// CloudWatch Logs encryption configuration.
//
// Example:
//   glue.NewSecurityConfiguration(this, jsii.String("MySecurityConfiguration"), &SecurityConfigurationProps{
//   	SecurityConfigurationName: jsii.String("name"),
//   	CloudWatchEncryption: &CloudWatchEncryption{
//   		Mode: glue.CloudWatchEncryptionMode_KMS,
//   	},
//   	JobBookmarksEncryption: &JobBookmarksEncryption{
//   		Mode: glue.JobBookmarksEncryptionMode_CLIENT_SIDE_KMS,
//   	},
//   	S3Encryption: &S3Encryption{
//   		Mode: glue.S3EncryptionMode_KMS,
//   	},
//   })
//
// Experimental.
type CloudWatchEncryption struct {
	// Encryption mode.
	// Experimental.
	Mode CloudWatchEncryptionMode `field:"required" json:"mode" yaml:"mode"`
	// The KMS key to be used to encrypt the data.
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

