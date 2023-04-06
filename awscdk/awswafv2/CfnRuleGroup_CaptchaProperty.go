package awswafv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captchaProperty := &CaptchaProperty{
//   	CustomRequestHandling: &CustomRequestHandlingProperty{
//   		InsertHeaders: []interface{}{
//   			&CustomHTTPHeaderProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnRuleGroup_CaptchaProperty struct {
	// `CfnRuleGroup.CaptchaProperty.CustomRequestHandling`.
	CustomRequestHandling interface{} `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}

