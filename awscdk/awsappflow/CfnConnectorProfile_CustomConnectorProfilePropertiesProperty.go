package awsappflow


// The profile properties required by the custom connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customConnectorProfilePropertiesProperty := &CustomConnectorProfilePropertiesProperty{
//   	OAuth2Properties: &OAuth2PropertiesProperty{
//   		OAuth2GrantType: jsii.String("oAuth2GrantType"),
//   		TokenUrl: jsii.String("tokenUrl"),
//   		TokenUrlCustomProperties: map[string]*string{
//   			"tokenUrlCustomPropertiesKey": jsii.String("tokenUrlCustomProperties"),
//   		},
//   	},
//   	ProfileProperties: map[string]*string{
//   		"profilePropertiesKey": jsii.String("profileProperties"),
//   	},
//   }
//
type CfnConnectorProfile_CustomConnectorProfilePropertiesProperty struct {
	// The OAuth 2.0 properties required for OAuth 2.0 authentication.
	OAuth2Properties interface{} `field:"optional" json:"oAuth2Properties" yaml:"oAuth2Properties"`
	// A map of properties that are required to create a profile for the custom connector.
	ProfileProperties interface{} `field:"optional" json:"profileProperties" yaml:"profileProperties"`
}
