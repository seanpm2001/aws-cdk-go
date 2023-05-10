package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnRegistry`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRegistryProps := &CfnRegistryProps{
//   	Name: jsii.String("name"),
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
type CfnRegistryProps struct {
	// The name of the registry.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the registry.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// AWS tags that contain a key value pair and may be searched by console, command line, or API.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

