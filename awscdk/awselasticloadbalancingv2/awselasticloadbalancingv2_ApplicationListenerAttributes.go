package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties to reference an existing listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup securityGroup
//
//   applicationListenerAttributes := &applicationListenerAttributes{
//   	listenerArn: jsii.String("listenerArn"),
//
//   	// the properties below are optional
//   	defaultPort: jsii.Number(123),
//   	securityGroup: securityGroup,
//   }
//
type ApplicationListenerAttributes struct {
	// ARN of the listener.
	ListenerArn *string `field:"required" json:"listenerArn" yaml:"listenerArn"`
	// The default port on which this listener is listening.
	DefaultPort *float64 `field:"optional" json:"defaultPort" yaml:"defaultPort"`
	// Security group of the load balancer this listener is associated with.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
}

