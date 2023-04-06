package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnServiceNetwork`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServiceNetworkProps := &CfnServiceNetworkProps{
//   	AuthType: jsii.String("authType"),
//   	Name: jsii.String("name"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnServiceNetworkProps struct {
	// The type of IAM policy.
	//
	// - `NONE` : The resource does not use an IAM policy. This is the default.
	// - `AWS_IAM` : The resource uses an IAM policy. When this type is used, auth is enabled and an auth policy is required.
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// The name of the service network.
	//
	// The name must be unique to the account. The valid characters are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or immediately after another hyphen.
	//
	// If you don't specify a name, CloudFormation generates one. However, if you specify a name, and later want to replace the resource, you must specify a new name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags for the service network.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}
