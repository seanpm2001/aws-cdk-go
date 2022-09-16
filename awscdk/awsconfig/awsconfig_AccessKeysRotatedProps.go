package awsconfig

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Construction properties for a AccessKeysRotated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var inputParameters interface{}
//   var ruleScope ruleScope
//
//   accessKeysRotatedProps := &accessKeysRotatedProps{
//   	configRuleName: jsii.String("configRuleName"),
//   	description: jsii.String("description"),
//   	inputParameters: map[string]interface{}{
//   		"inputParametersKey": inputParameters,
//   	},
//   	maxAge: duration,
//   	maximumExecutionFrequency: awscdk.Aws_config.maximumExecutionFrequency_ONE_HOUR,
//   	ruleScope: ruleScope,
//   }
//
// Experimental.
type AccessKeysRotatedProps struct {
	// A name for the AWS Config rule.
	// Experimental.
	ConfigRuleName *string `field:"optional" json:"configRuleName" yaml:"configRuleName"`
	// A description about this AWS Config rule.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Input parameter values that are passed to the AWS Config rule.
	// Experimental.
	InputParameters *map[string]interface{} `field:"optional" json:"inputParameters" yaml:"inputParameters"`
	// The maximum frequency at which the AWS Config rule runs evaluations.
	// Experimental.
	MaximumExecutionFrequency MaximumExecutionFrequency `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
	// Defines which resources trigger an evaluation for an AWS Config rule.
	// Experimental.
	RuleScope RuleScope `field:"optional" json:"ruleScope" yaml:"ruleScope"`
	// The maximum number of days within which the access keys must be rotated.
	// Experimental.
	MaxAge awscdk.Duration `field:"optional" json:"maxAge" yaml:"maxAge"`
}
