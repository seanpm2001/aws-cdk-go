package awsbillingconductor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPricingPlan`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPricingPlanProps := &CfnPricingPlanProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	PricingRuleArns: []*string{
//   		jsii.String("pricingRuleArns"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPricingPlanProps struct {
	// The name of a pricing plan.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The pricing plan description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The `PricingRuleArns` that are associated with the Pricing Plan.
	PricingRuleArns *[]*string `field:"optional" json:"pricingRuleArns" yaml:"pricingRuleArns"`
	// A map that contains tag keys and tag values that are attached to a pricing plan.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}
