package awsmemorydb

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnParameterGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   cfnParameterGroupProps := &CfnParameterGroupProps{
//   	Family: jsii.String("family"),
//   	ParameterGroupName: jsii.String("parameterGroupName"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Parameters: parameters,
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnParameterGroupProps struct {
	// The name of the parameter group family that this parameter group is compatible with.
	Family *string `field:"required" json:"family" yaml:"family"`
	// The name of the parameter group.
	ParameterGroupName *string `field:"required" json:"parameterGroupName" yaml:"parameterGroupName"`
	// A description of the parameter group.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Returns the detailed parameter list for the parameter group.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

