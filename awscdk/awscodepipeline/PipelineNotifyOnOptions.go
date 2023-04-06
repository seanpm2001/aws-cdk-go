package awscodepipeline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarnotifications"
)

// Additional options to pass to the notification rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipelineNotifyOnOptions := &PipelineNotifyOnOptions{
//   	Events: []pipelineNotificationEvents{
//   		awscdk.Aws_codepipeline.*pipelineNotificationEvents_PIPELINE_EXECUTION_FAILED,
//   	},
//
//   	// the properties below are optional
//   	DetailType: awscdk.Aws_codestarnotifications.DetailType_BASIC,
//   	Enabled: jsii.Boolean(false),
//   	NotificationRuleName: jsii.String("notificationRuleName"),
//   }
//
type PipelineNotifyOnOptions struct {
	// The level of detail to include in the notifications for this resource.
	//
	// BASIC will include only the contents of the event as it would appear in AWS CloudWatch.
	// FULL will include any supplemental information provided by AWS CodeStar Notifications and/or the service for the resource for which the notification is created.
	DetailType awscodestarnotifications.DetailType `field:"optional" json:"detailType" yaml:"detailType"`
	// The status of the notification rule.
	//
	// If the enabled is set to DISABLED, notifications aren't sent for the notification rule.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The name for the notification rule.
	//
	// Notification rule names must be unique in your AWS account.
	NotificationRuleName *string `field:"optional" json:"notificationRuleName" yaml:"notificationRuleName"`
	// A list of event types associated with this notification rule for CodePipeline Pipeline.
	//
	// For a complete list of event types and IDs, see Notification concepts in the Developer Tools Console User Guide.
	// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#concepts-api
	//
	Events *[]PipelineNotificationEvents `field:"required" json:"events" yaml:"events"`
}
