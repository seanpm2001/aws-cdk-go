package awsefs

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// Properties that describe an existing EFS file system.
//
// Example:
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//
//
//   importedFileSystem := efs.FileSystem_FromFileSystemAttributes(this, jsii.String("existingFS"), &FileSystemAttributes{
//   	FileSystemId: jsii.String("fs-12345678"),
//   	 // You can also use fileSystemArn instead of fileSystemId.
//   	SecurityGroup: ec2.SecurityGroup_FromSecurityGroupId(this, jsii.String("SG"), jsii.String("sg-123456789"), &SecurityGroupImportOptions{
//   		AllowAllOutbound: jsii.Boolean(false),
//   	}),
//   })
//
// Experimental.
type FileSystemAttributes struct {
	// The security group of the file system.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"required" json:"securityGroup" yaml:"securityGroup"`
	// The File System's Arn.
	// Experimental.
	FileSystemArn *string `field:"optional" json:"fileSystemArn" yaml:"fileSystemArn"`
	// The File System's ID.
	// Experimental.
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
}

