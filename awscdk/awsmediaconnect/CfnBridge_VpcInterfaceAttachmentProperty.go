package awsmediaconnect


// The VPC interface that you want to send your output to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcInterfaceAttachmentProperty := &VpcInterfaceAttachmentProperty{
//   	VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   }
//
type CfnBridge_VpcInterfaceAttachmentProperty struct {
	// The name of the VPC interface that you want to send your output to.
	VpcInterfaceName *string `field:"optional" json:"vpcInterfaceName" yaml:"vpcInterfaceName"`
}
