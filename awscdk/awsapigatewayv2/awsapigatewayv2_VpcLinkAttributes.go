package awsapigatewayv2

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// Attributes when importing a new VpcLink.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//
//   awesomeLink := apigwv2.vpcLink.fromVpcLinkAttributes(this, jsii.String("awesome-vpc-link"), &vpcLinkAttributes{
//   	vpcLinkId: jsii.String("us-east-1_oiuR12Abd"),
//   	vpc: vpc,
//   })
//
// Experimental.
type VpcLinkAttributes struct {
	// The VPC to which this VPC link is associated with.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// The VPC Link id.
	// Experimental.
	VpcLinkId *string `field:"required" json:"vpcLinkId" yaml:"vpcLinkId"`
}

