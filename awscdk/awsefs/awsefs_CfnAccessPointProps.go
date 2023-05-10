package awsefs


// Properties for defining a `CfnAccessPoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccessPointProps := &CfnAccessPointProps{
//   	FileSystemId: jsii.String("fileSystemId"),
//
//   	// the properties below are optional
//   	AccessPointTags: []interface{}{
//   		&AccessPointTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ClientToken: jsii.String("clientToken"),
//   	PosixUser: &PosixUserProperty{
//   		Gid: jsii.String("gid"),
//   		Uid: jsii.String("uid"),
//
//   		// the properties below are optional
//   		SecondaryGids: []*string{
//   			jsii.String("secondaryGids"),
//   		},
//   	},
//   	RootDirectory: &RootDirectoryProperty{
//   		CreationInfo: &CreationInfoProperty{
//   			OwnerGid: jsii.String("ownerGid"),
//   			OwnerUid: jsii.String("ownerUid"),
//   			Permissions: jsii.String("permissions"),
//   		},
//   		Path: jsii.String("path"),
//   	},
//   }
//
type CfnAccessPointProps struct {
	// The ID of the EFS file system that the access point applies to.
	//
	// Accepts only the ID format for input when specifying a file system, for example `fs-0123456789abcedf2` .
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	AccessPointTags interface{} `field:"optional" json:"accessPointTags" yaml:"accessPointTags"`
	// The opaque string specified in the request to ensure idempotent creation.
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// The full POSIX identity, including the user ID, group ID, and secondary group IDs on the access point that is used for all file operations by NFS clients using the access point.
	PosixUser interface{} `field:"optional" json:"posixUser" yaml:"posixUser"`
	// The directory on the Amazon EFS file system that the access point exposes as the root directory to NFS clients using the access point.
	RootDirectory interface{} `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
}

