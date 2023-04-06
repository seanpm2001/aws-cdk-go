package awsconfig


// Properties for defining a `CfnRemediationConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   cfnRemediationConfigurationProps := &CfnRemediationConfigurationProps{
//   	ConfigRuleName: jsii.String("configRuleName"),
//   	TargetId: jsii.String("targetId"),
//   	TargetType: jsii.String("targetType"),
//
//   	// the properties below are optional
//   	Automatic: jsii.Boolean(false),
//   	ExecutionControls: &ExecutionControlsProperty{
//   		SsmControls: &SsmControlsProperty{
//   			ConcurrentExecutionRatePercentage: jsii.Number(123),
//   			ErrorPercentage: jsii.Number(123),
//   		},
//   	},
//   	MaximumAutomaticAttempts: jsii.Number(123),
//   	Parameters: parameters,
//   	ResourceType: jsii.String("resourceType"),
//   	RetryAttemptSeconds: jsii.Number(123),
//   	TargetVersion: jsii.String("targetVersion"),
//   }
//
type CfnRemediationConfigurationProps struct {
	// The name of the AWS Config rule.
	ConfigRuleName *string `field:"required" json:"configRuleName" yaml:"configRuleName"`
	// Target ID is the name of the SSM document.
	TargetId *string `field:"required" json:"targetId" yaml:"targetId"`
	// The type of the target.
	//
	// Target executes remediation. For example, SSM document.
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
	// The remediation is triggered automatically.
	Automatic interface{} `field:"optional" json:"automatic" yaml:"automatic"`
	// An ExecutionControls object.
	ExecutionControls interface{} `field:"optional" json:"executionControls" yaml:"executionControls"`
	// The maximum number of failed attempts for auto-remediation. If you do not select a number, the default is 5.
	//
	// For example, if you specify MaximumAutomaticAttempts as 5 with RetryAttemptSeconds as 50 seconds, AWS Config will put a RemediationException on your behalf for the failing resource after the 5th failed attempt within 50 seconds.
	MaximumAutomaticAttempts *float64 `field:"optional" json:"maximumAutomaticAttempts" yaml:"maximumAutomaticAttempts"`
	// An object of the RemediationParameterValue. For more information, see [RemediationParameterValue](https://docs.aws.amazon.com/config/latest/APIReference/API_RemediationParameterValue.html) .
	//
	// > The type is a map of strings to RemediationParameterValue.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The type of a resource.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Maximum time in seconds that AWS Config runs auto-remediation.
	//
	// If you do not select a number, the default is 60 seconds.
	//
	// For example, if you specify RetryAttemptSeconds as 50 seconds and MaximumAutomaticAttempts as 5, AWS Config will run auto-remediations 5 times within 50 seconds before throwing an exception.
	RetryAttemptSeconds *float64 `field:"optional" json:"retryAttemptSeconds" yaml:"retryAttemptSeconds"`
	// Version of the target. For example, version of the SSM document.
	//
	// > If you make backward incompatible changes to the SSM document, you must call PutRemediationConfiguration API again to ensure the remediations can run.
	TargetVersion *string `field:"optional" json:"targetVersion" yaml:"targetVersion"`
}

