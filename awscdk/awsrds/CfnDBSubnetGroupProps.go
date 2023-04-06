package awsrds

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
	// The description for the DB subnet group.
	DbSubnetGroupDescription *string `field:"required" json:"dbSubnetGroupDescription" yaml:"dbSubnetGroupDescription"`
	// The EC2 Subnet IDs for the DB subnet group.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The name for the DB subnet group. This value is stored as a lowercase string.
	//
	// Constraints: Must contain no more than 255 lowercase alphanumeric characters or hyphens. Must not be "Default".
	//
	// Example: `mysubnetgroup`.
	DbSubnetGroupName *string `field:"optional" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// An optional array of key-value pairs to apply to this DB subnet group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

