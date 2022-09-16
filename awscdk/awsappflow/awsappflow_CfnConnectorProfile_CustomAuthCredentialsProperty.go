package awsappflow


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var credentialsMap interface{}
//
//   customAuthCredentialsProperty := &customAuthCredentialsProperty{
//   	customAuthenticationType: jsii.String("customAuthenticationType"),
//
//   	// the properties below are optional
//   	credentialsMap: credentialsMap,
//   }
//
type CfnConnectorProfile_CustomAuthCredentialsProperty struct {
	// `CfnConnectorProfile.CustomAuthCredentialsProperty.CustomAuthenticationType`.
	CustomAuthenticationType *string `field:"required" json:"customAuthenticationType" yaml:"customAuthenticationType"`
	// `CfnConnectorProfile.CustomAuthCredentialsProperty.CredentialsMap`.
	CredentialsMap interface{} `field:"optional" json:"credentialsMap" yaml:"credentialsMap"`
}
