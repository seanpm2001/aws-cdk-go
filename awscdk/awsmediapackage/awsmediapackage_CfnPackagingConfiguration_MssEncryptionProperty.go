package awsmediapackage


// Holds encryption information so that access to the content can be controlled by a DRM solution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mssEncryptionProperty := &mssEncryptionProperty{
//   	spekeKeyProvider: &spekeKeyProviderProperty{
//   		roleArn: jsii.String("roleArn"),
//   		systemIds: []*string{
//   			jsii.String("systemIds"),
//   		},
//   		url: jsii.String("url"),
//   	},
//   }
//
type CfnPackagingConfiguration_MssEncryptionProperty struct {
	// Parameters for the SPEKE key provider.
	SpekeKeyProvider interface{} `field:"required" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
}
