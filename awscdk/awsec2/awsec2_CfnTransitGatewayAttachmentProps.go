package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnTransitGatewayAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var options interface{}
//
//   cfnTransitGatewayAttachmentProps := &CfnTransitGatewayAttachmentProps{
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	TransitGatewayId: jsii.String("transitGatewayId"),
//   	VpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	Options: options,
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnTransitGatewayAttachmentProps struct {
	// The IDs of one or more subnets.
	//
	// You can specify only one subnet per Availability Zone. You must specify at least one subnet, but we recommend that you specify two subnets for better availability. The transit gateway uses one IP address from each specified subnet.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the transit gateway.
	TransitGatewayId *string `field:"required" json:"transitGatewayId" yaml:"transitGatewayId"`
	// The ID of the VPC.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The VPC attachment options, in JSON or YAML.
	//
	// - `ApplianceModeSupport` - Set to `enable` or `disable` . The default is `disable` .
	// - `DnsSupport` - Set to `enable` or `disable` . The default is `enable` .
	// - `Ipv6Support` - Set to `enable` or `disable` . The default is `disable` .
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// The tags for the attachment.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

