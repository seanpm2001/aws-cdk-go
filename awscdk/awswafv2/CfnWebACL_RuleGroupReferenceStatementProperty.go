package awswafv2


// A rule statement used to run the rules that are defined in a `RuleGroup` .
//
// To use this, create a rule group with your rules, then provide the ARN of the rule group in this statement.
//
// You cannot nest a `RuleGroupReferenceStatement` , for example for use inside a `NotStatement` or `OrStatement` . You can only use a rule group reference statement at the top level inside a web ACL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleGroupReferenceStatementProperty := &RuleGroupReferenceStatementProperty{
//   	Arn: jsii.String("arn"),
//
//   	// the properties below are optional
//   	ExcludedRules: []interface{}{
//   		&ExcludedRuleProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	RuleActionOverrides: []interface{}{
//   		&RuleActionOverrideProperty{
//   			ActionToUse: &RuleActionProperty{
//   				Allow: &AllowActionProperty{
//   					CustomRequestHandling: &CustomRequestHandlingProperty{
//   						InsertHeaders: []interface{}{
//   							&CustomHTTPHeaderProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				Block: &BlockActionProperty{
//   					CustomResponse: &CustomResponseProperty{
//   						ResponseCode: jsii.Number(123),
//
//   						// the properties below are optional
//   						CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   						ResponseHeaders: []interface{}{
//   							&CustomHTTPHeaderProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				Captcha: &CaptchaActionProperty{
//   					CustomRequestHandling: &CustomRequestHandlingProperty{
//   						InsertHeaders: []interface{}{
//   							&CustomHTTPHeaderProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				Challenge: &ChallengeActionProperty{
//   					CustomRequestHandling: &CustomRequestHandlingProperty{
//   						InsertHeaders: []interface{}{
//   							&CustomHTTPHeaderProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				Count: &CountActionProperty{
//   					CustomRequestHandling: &CustomRequestHandlingProperty{
//   						InsertHeaders: []interface{}{
//   							&CustomHTTPHeaderProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
type CfnWebACL_RuleGroupReferenceStatementProperty struct {
	// The Amazon Resource Name (ARN) of the entity.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Rules in the referenced rule group whose actions are set to `Count` .
	//
	// > Instead of this option, use `RuleActionOverrides` . It accepts any valid action setting, including `Count` .
	ExcludedRules interface{} `field:"optional" json:"excludedRules" yaml:"excludedRules"`
	// Action settings to use in the place of the rule actions that are configured inside the rule group.
	//
	// You specify one override for each rule whose action you want to change.
	//
	// You can use overrides for testing, for example you can override all of rule actions to `Count` and then monitor the resulting count metrics to understand how the rule group would handle your web traffic. You can also permanently override some or all actions, to modify how the rule group manages your web traffic.
	RuleActionOverrides interface{} `field:"optional" json:"ruleActionOverrides" yaml:"ruleActionOverrides"`
}

