package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodestarnotifications"
)

// Additional options to pass to the notification rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   projectNotifyOnOptions := &projectNotifyOnOptions{
//   	events: []projectNotificationEvents{
//   		awscdk.Aws_codebuild.*projectNotificationEvents_BUILD_FAILED,
//   	},
//
//   	// the properties below are optional
//   	detailType: awscdk.Aws_codestarnotifications.detailType_BASIC,
//   	enabled: jsii.Boolean(false),
//   	notificationRuleName: jsii.String("notificationRuleName"),
//   }
//
// Experimental.
type ProjectNotifyOnOptions struct {
	// The level of detail to include in the notifications for this resource.
	//
	// BASIC will include only the contents of the event as it would appear in AWS CloudWatch.
	// FULL will include any supplemental information provided by AWS CodeStar Notifications and/or the service for the resource for which the notification is created.
	// Experimental.
	DetailType awscodestarnotifications.DetailType `field:"optional" json:"detailType" yaml:"detailType"`
	// The status of the notification rule.
	//
	// If the enabled is set to DISABLED, notifications aren't sent for the notification rule.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The name for the notification rule.
	//
	// Notification rule names must be unique in your AWS account.
	// Experimental.
	NotificationRuleName *string `field:"optional" json:"notificationRuleName" yaml:"notificationRuleName"`
	// A list of event types associated with this notification rule for CodeBuild Project.
	//
	// For a complete list of event types and IDs, see Notification concepts in the Developer Tools Console User Guide.
	// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#concepts-api
	//
	// Experimental.
	Events *[]ProjectNotificationEvents `field:"required" json:"events" yaml:"events"`
}
