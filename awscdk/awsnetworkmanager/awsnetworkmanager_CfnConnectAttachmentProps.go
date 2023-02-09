package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnConnectAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConnectAttachmentProps := &cfnConnectAttachmentProps{
//   	coreNetworkId: jsii.String("coreNetworkId"),
//   	edgeLocation: jsii.String("edgeLocation"),
//   	options: &connectAttachmentOptionsProperty{
//   		protocol: jsii.String("protocol"),
//   	},
//   	transportAttachmentId: jsii.String("transportAttachmentId"),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnConnectAttachmentProps struct {
	// The ID of the core network where the Connect attachment is located.
	CoreNetworkId *string `field:"required" json:"coreNetworkId" yaml:"coreNetworkId"`
	// The Region where the edge is located.
	EdgeLocation *string `field:"required" json:"edgeLocation" yaml:"edgeLocation"`
	// Options for connecting an attachment.
	Options interface{} `field:"required" json:"options" yaml:"options"`
	// The ID of the transport attachment.
	TransportAttachmentId *string `field:"required" json:"transportAttachmentId" yaml:"transportAttachmentId"`
	// `AWS::NetworkManager::ConnectAttachment.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

