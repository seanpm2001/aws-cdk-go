// The CDK Construct Library for AWS::Glue
package awscdkgluealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// CloudWatch Logs encryption configuration.
//
// Example:
//   glue.NewSecurityConfiguration(this, jsii.String("MySecurityConfiguration"), &securityConfigurationProps{
//   	cloudWatchEncryption: &cloudWatchEncryption{
//   		mode: glue.cloudWatchEncryptionMode_KMS,
//   	},
//   	jobBookmarksEncryption: &jobBookmarksEncryption{
//   		mode: glue.jobBookmarksEncryptionMode_CLIENT_SIDE_KMS,
//   	},
//   	s3Encryption: &s3Encryption{
//   		mode: glue.s3EncryptionMode_KMS,
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
