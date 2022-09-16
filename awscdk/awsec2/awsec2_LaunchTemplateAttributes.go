package awsec2


// Attributes for an imported LaunchTemplate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateAttributes := &launchTemplateAttributes{
//   	launchTemplateId: jsii.String("launchTemplateId"),
//   	launchTemplateName: jsii.String("launchTemplateName"),
//   	versionNumber: jsii.String("versionNumber"),
//   }
//
// Experimental.
type LaunchTemplateAttributes struct {
	// The identifier of the Launch Template.
	//
	// Exactly one of `launchTemplateId` and `launchTemplateName` may be set.
	// Experimental.
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// The name of the Launch Template.
	//
	// Exactly one of `launchTemplateId` and `launchTemplateName` may be set.
	// Experimental.
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// The version number of this launch template to use.
	// Experimental.
	VersionNumber *string `field:"optional" json:"versionNumber" yaml:"versionNumber"`
}
