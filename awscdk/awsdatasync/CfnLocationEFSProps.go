package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLocationEFS`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocationEFSProps := &CfnLocationEFSProps{
//   	Ec2Config: &Ec2ConfigProperty{
//   		SecurityGroupArns: []*string{
//   			jsii.String("securityGroupArns"),
//   		},
//   		SubnetArn: jsii.String("subnetArn"),
//   	},
//
//   	// the properties below are optional
//   	AccessPointArn: jsii.String("accessPointArn"),
//   	EfsFilesystemArn: jsii.String("efsFilesystemArn"),
//   	FileSystemAccessRoleArn: jsii.String("fileSystemAccessRoleArn"),
//   	InTransitEncryption: jsii.String("inTransitEncryption"),
//   	Subdirectory: jsii.String("subdirectory"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLocationEFSProps struct {
	// Specifies the subnet and security groups DataSync uses to access your Amazon EFS file system.
	Ec2Config interface{} `field:"required" json:"ec2Config" yaml:"ec2Config"`
	// Specifies the Amazon Resource Name (ARN) of the access point that DataSync uses to access the Amazon EFS file system.
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// Specifies the ARN for the Amazon EFS file system.
	EfsFilesystemArn *string `field:"optional" json:"efsFilesystemArn" yaml:"efsFilesystemArn"`
	// Specifies an AWS Identity and Access Management (IAM) role that DataSync assumes when mounting the Amazon EFS file system.
	FileSystemAccessRoleArn *string `field:"optional" json:"fileSystemAccessRoleArn" yaml:"fileSystemAccessRoleArn"`
	// Specifies whether you want DataSync to use Transport Layer Security (TLS) 1.2 encryption when it copies data to or from the Amazon EFS file system.
	//
	// If you specify an access point using `AccessPointArn` or an IAM role using `FileSystemAccessRoleArn` , you must set this parameter to `TLS1_2` .
	InTransitEncryption *string `field:"optional" json:"inTransitEncryption" yaml:"inTransitEncryption"`
	// Specifies a mount path for your Amazon EFS file system.
	//
	// This is where DataSync reads or writes data (depending on if this is a source or destination location). By default, DataSync uses the root directory, but you can also include subdirectories.
	//
	// > You must specify a value with forward slashes (for example, `/path/to/folder` ).
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// Specifies the key-value pair that represents a tag that you want to add to the resource.
	//
	// The value can be an empty string. This value helps you manage, filter, and search for your resources. We recommend that you create a name tag for your location.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}
