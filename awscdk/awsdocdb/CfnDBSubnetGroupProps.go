package awsdocdb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDBSubnetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBSubnetGroupProps := &CfnDBSubnetGroupProps{
//   	DbSubnetGroupDescription: jsii.String("dbSubnetGroupDescription"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	DbSubnetGroupName: jsii.String("dbSubnetGroupName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDBSubnetGroupProps struct {
	// The description for the subnet group.
	DbSubnetGroupDescription *string `field:"required" json:"dbSubnetGroupDescription" yaml:"dbSubnetGroupDescription"`
	// The Amazon EC2 subnet IDs for the subnet group.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The name for the subnet group. This value is stored as a lowercase string.
	//
	// Constraints: Must contain no more than 255 letters, numbers, periods, underscores, spaces, or hyphens. Must not be default.
	//
	// Example: `mySubnetgroup`.
	DbSubnetGroupName *string `field:"optional" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// The tags to be assigned to the subnet group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

