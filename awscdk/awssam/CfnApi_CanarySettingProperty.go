package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   canarySettingProperty := &CanarySettingProperty{
//   	DeploymentId: jsii.String("deploymentId"),
//   	PercentTraffic: jsii.Number(123),
//   	StageVariableOverrides: map[string]*string{
//   		"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   	},
//   	UseStageCache: jsii.Boolean(false),
//   }
//
type CfnApi_CanarySettingProperty struct {
	// `CfnApi.CanarySettingProperty.DeploymentId`.
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// `CfnApi.CanarySettingProperty.PercentTraffic`.
	PercentTraffic *float64 `field:"optional" json:"percentTraffic" yaml:"percentTraffic"`
	// `CfnApi.CanarySettingProperty.StageVariableOverrides`.
	StageVariableOverrides interface{} `field:"optional" json:"stageVariableOverrides" yaml:"stageVariableOverrides"`
	// `CfnApi.CanarySettingProperty.UseStageCache`.
	UseStageCache interface{} `field:"optional" json:"useStageCache" yaml:"useStageCache"`
}

