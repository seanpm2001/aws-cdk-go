package awsiotfleethub

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &CfnApplicationProps{
//   	ApplicationName: jsii.String("applicationName"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	ApplicationDescription: jsii.String("applicationDescription"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnApplicationProps struct {
	// The name of the web application.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// The ARN of the role that the web application assumes when it interacts with AWS IoT Core .
	//
	// > The name of the role must be in the form `FleetHub_random_string` .
	//
	// Pattern: `^arn:[!-~]+$`
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// An optional description of the web application.
	ApplicationDescription *string `field:"optional" json:"applicationDescription" yaml:"applicationDescription"`
	// A set of key/value pairs that you can use to manage the web application resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}
