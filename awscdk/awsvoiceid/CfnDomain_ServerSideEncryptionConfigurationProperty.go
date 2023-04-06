package awsvoiceid


// The configuration containing information about the customer managed key used for encrypting customer data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverSideEncryptionConfigurationProperty := &ServerSideEncryptionConfigurationProperty{
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnDomain_ServerSideEncryptionConfigurationProperty struct {
	// The identifier of the KMS key to use to encrypt data stored by Voice ID.
	//
	// Voice ID doesn't support asymmetric customer managed keys .
	KmsKeyId *string `field:"required" json:"kmsKeyId" yaml:"kmsKeyId"`
}

