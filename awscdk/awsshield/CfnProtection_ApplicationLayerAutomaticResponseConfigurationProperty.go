package awsshield


// The automatic application layer DDoS mitigation settings for a `Protection` .
//
// This configuration determines whether Shield Advanced automatically manages rules in the web ACL in order to respond to application layer events that Shield Advanced determines to be DDoS attacks.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var block interface{}
//   var count interface{}
//
//   applicationLayerAutomaticResponseConfigurationProperty := &ApplicationLayerAutomaticResponseConfigurationProperty{
//   	Action: &ActionProperty{
//   		Block: block,
//   		Count: count,
//   	},
//   	Status: jsii.String("status"),
//   }
//
type CfnProtection_ApplicationLayerAutomaticResponseConfigurationProperty struct {
	// Specifies the action setting that Shield Advanced should use in the AWS WAF rules that it creates on behalf of the protected resource in response to DDoS attacks.
	//
	// You specify this as part of the configuration for the automatic application layer DDoS mitigation feature, when you enable or update automatic mitigation. Shield Advanced creates the AWS WAF rules in a Shield Advanced-managed rule group, inside the web ACL that you have associated with the resource.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// Indicates whether automatic application layer DDoS mitigation is enabled for the protection.
	Status *string `field:"required" json:"status" yaml:"status"`
}
