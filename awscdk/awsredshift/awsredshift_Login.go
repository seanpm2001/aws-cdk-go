package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// Username and password combination.
//
// Example:
//   user := awscdk.NewUser(this, jsii.String("User"), &userProps{
//   	cluster: cluster,
//   	databaseName: jsii.String("databaseName"),
//   })
//   cluster.addRotationMultiUser(jsii.String("MultiUserRotation"), &rotationMultiUserOptions{
//   	secret: user.secret,
//   })
//
// Experimental.
type Login struct {
	// Username.
	// Experimental.
	MasterUsername *string `field:"required" json:"masterUsername" yaml:"masterUsername"`
	// KMS encryption key to encrypt the generated secret.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Password.
	//
	// Do not put passwords in your CDK code directly.
	// Experimental.
	MasterPassword awscdk.SecretValue `field:"optional" json:"masterPassword" yaml:"masterPassword"`
}
