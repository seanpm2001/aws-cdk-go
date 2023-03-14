package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnQuickConnect`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnQuickConnectProps := &CfnQuickConnectProps{
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//   	QuickConnectConfig: &QuickConnectConfigProperty{
//   		QuickConnectType: jsii.String("quickConnectType"),
//
//   		// the properties below are optional
//   		PhoneConfig: &PhoneNumberQuickConnectConfigProperty{
//   			PhoneNumber: jsii.String("phoneNumber"),
//   		},
//   		QueueConfig: &QueueQuickConnectConfigProperty{
//   			ContactFlowArn: jsii.String("contactFlowArn"),
//   			QueueArn: jsii.String("queueArn"),
//   		},
//   		UserConfig: &UserQuickConnectConfigProperty{
//   			ContactFlowArn: jsii.String("contactFlowArn"),
//   			UserArn: jsii.String("userArn"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnQuickConnectProps struct {
	// The Amazon Resource Name (ARN) of the instance.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the quick connect.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Contains information about the quick connect.
	QuickConnectConfig interface{} `field:"required" json:"quickConnectConfig" yaml:"quickConnectConfig"`
	// The description of the quick connect.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags used to organize, track, or control access for this resource.
	//
	// For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

