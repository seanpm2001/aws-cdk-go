package awsnimblestudio


// Properties for defining a `CfnStudioComponent`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStudioComponentProps := &CfnStudioComponentProps{
//   	Name: jsii.String("name"),
//   	StudioId: jsii.String("studioId"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Configuration: &StudioComponentConfigurationProperty{
//   		ActiveDirectoryConfiguration: &ActiveDirectoryConfigurationProperty{
//   			ComputerAttributes: []interface{}{
//   				&ActiveDirectoryComputerAttributeProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			DirectoryId: jsii.String("directoryId"),
//   			OrganizationalUnitDistinguishedName: jsii.String("organizationalUnitDistinguishedName"),
//   		},
//   		ComputeFarmConfiguration: &ComputeFarmConfigurationProperty{
//   			ActiveDirectoryUser: jsii.String("activeDirectoryUser"),
//   			Endpoint: jsii.String("endpoint"),
//   		},
//   		LicenseServiceConfiguration: &LicenseServiceConfigurationProperty{
//   			Endpoint: jsii.String("endpoint"),
//   		},
//   		SharedFileSystemConfiguration: &SharedFileSystemConfigurationProperty{
//   			Endpoint: jsii.String("endpoint"),
//   			FileSystemId: jsii.String("fileSystemId"),
//   			LinuxMountPoint: jsii.String("linuxMountPoint"),
//   			ShareName: jsii.String("shareName"),
//   			WindowsMountDrive: jsii.String("windowsMountDrive"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Ec2SecurityGroupIds: []*string{
//   		jsii.String("ec2SecurityGroupIds"),
//   	},
//   	InitializationScripts: []interface{}{
//   		&StudioComponentInitializationScriptProperty{
//   			LaunchProfileProtocolVersion: jsii.String("launchProfileProtocolVersion"),
//   			Platform: jsii.String("platform"),
//   			RunContext: jsii.String("runContext"),
//   			Script: jsii.String("script"),
//   		},
//   	},
//   	ScriptParameters: []interface{}{
//   		&ScriptParameterKeyValueProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Subtype: jsii.String("subtype"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnStudioComponentProps struct {
	// A friendly name for the studio component resource.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The unique identifier for a studio resource.
	//
	// In Nimble Studio , all other resources are contained in a studio resource.
	StudioId *string `field:"required" json:"studioId" yaml:"studioId"`
	// The type of the studio component.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The configuration of the studio component, based on component type.
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// A human-readable description for the studio component resource.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The EC2 security groups that control access to the studio component.
	Ec2SecurityGroupIds *[]*string `field:"optional" json:"ec2SecurityGroupIds" yaml:"ec2SecurityGroupIds"`
	// Initialization scripts for studio components.
	InitializationScripts interface{} `field:"optional" json:"initializationScripts" yaml:"initializationScripts"`
	// Parameters for the studio component scripts.
	ScriptParameters interface{} `field:"optional" json:"scriptParameters" yaml:"scriptParameters"`
	// The specific subtype of a studio component.
	Subtype *string `field:"optional" json:"subtype" yaml:"subtype"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

