package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnContactList`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContactListProps := &CfnContactListProps{
//   	ContactListName: jsii.String("contactListName"),
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Topics: []interface{}{
//   		&TopicProperty{
//   			DefaultSubscriptionStatus: jsii.String("defaultSubscriptionStatus"),
//   			DisplayName: jsii.String("displayName"),
//   			TopicName: jsii.String("topicName"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   		},
//   	},
//   }
//
type CfnContactListProps struct {
	// The name of the contact list.
	ContactListName *string `field:"optional" json:"contactListName" yaml:"contactListName"`
	// A description of what the contact list is about.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags associated with a contact list.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// An interest group, theme, or label within a list.
	//
	// A contact list can have multiple topics.
	Topics interface{} `field:"optional" json:"topics" yaml:"topics"`
}

