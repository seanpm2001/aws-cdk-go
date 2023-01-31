package awsconnect


// Properties for defining a `CfnApprovedOrigin`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApprovedOriginProps := &cfnApprovedOriginProps{
//   	instanceId: jsii.String("instanceId"),
//   	origin: jsii.String("origin"),
//   }
//
type CfnApprovedOriginProps struct {
	// `AWS::Connect::ApprovedOrigin.InstanceId`.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// `AWS::Connect::ApprovedOrigin.Origin`.
	Origin *string `field:"required" json:"origin" yaml:"origin"`
}
