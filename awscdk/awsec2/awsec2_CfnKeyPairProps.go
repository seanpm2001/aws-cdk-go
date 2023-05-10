package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnKeyPair`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnKeyPairProps := &CfnKeyPairProps{
//   	KeyName: jsii.String("keyName"),
//
//   	// the properties below are optional
//   	KeyType: jsii.String("keyType"),
//   	PublicKeyMaterial: jsii.String("publicKeyMaterial"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnKeyPairProps struct {
	// A unique name for the key pair.
	//
	// Constraints: Up to 255 ASCII characters.
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
	// The type of key pair. Note that ED25519 keys are not supported for Windows instances.
	//
	// If the `PublicKeyMaterial` property is specified, the `KeyType` property is ignored, and the key type is inferred from the `PublicKeyMaterial` value.
	//
	// Default: `rsa`.
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
	// The public key material.
	//
	// The `PublicKeyMaterial` property is used to import a key pair. If this property is not specified, then a new key pair will be created.
	PublicKeyMaterial *string `field:"optional" json:"publicKeyMaterial" yaml:"publicKeyMaterial"`
	// The tags to apply to the key pair.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

