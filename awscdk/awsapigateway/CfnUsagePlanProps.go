package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnUsagePlan`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUsagePlanProps := &CfnUsagePlanProps{
//   	ApiStages: []interface{}{
//   		&ApiStageProperty{
//   			ApiId: jsii.String("apiId"),
//   			Stage: jsii.String("stage"),
//   			Throttle: map[string]interface{}{
//   				"throttleKey": &ThrottleSettingsProperty{
//   					"burstLimit": jsii.Number(123),
//   					"rateLimit": jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Quota: &QuotaSettingsProperty{
//   		Limit: jsii.Number(123),
//   		Offset: jsii.Number(123),
//   		Period: jsii.String("period"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Throttle: &ThrottleSettingsProperty{
//   		BurstLimit: jsii.Number(123),
//   		RateLimit: jsii.Number(123),
//   	},
//   	UsagePlanName: jsii.String("usagePlanName"),
//   }
//
type CfnUsagePlanProps struct {
	// The associated API stages of a usage plan.
	ApiStages interface{} `field:"optional" json:"apiStages" yaml:"apiStages"`
	// The description of a usage plan.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The target maximum number of permitted requests per a given unit time interval.
	Quota interface{} `field:"optional" json:"quota" yaml:"quota"`
	// The collection of tags.
	//
	// Each tag element is associated with a given resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A map containing method level throttling information for API stage in a usage plan.
	Throttle interface{} `field:"optional" json:"throttle" yaml:"throttle"`
	// The name of a usage plan.
	UsagePlanName *string `field:"optional" json:"usagePlanName" yaml:"usagePlanName"`
}

