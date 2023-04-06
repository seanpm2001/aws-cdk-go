package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnProvisioningTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProvisioningTemplateProps := &CfnProvisioningTemplateProps{
//   	ProvisioningRoleArn: jsii.String("provisioningRoleArn"),
//   	TemplateBody: jsii.String("templateBody"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Enabled: jsii.Boolean(false),
//   	PreProvisioningHook: &ProvisioningHookProperty{
//   		PayloadVersion: jsii.String("payloadVersion"),
//   		TargetArn: jsii.String("targetArn"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TemplateName: jsii.String("templateName"),
//   	TemplateType: jsii.String("templateType"),
//   }
//
type CfnProvisioningTemplateProps struct {
	// The role ARN for the role associated with the fleet provisioning template.
	//
	// This IoT role grants permission to provision a device.
	ProvisioningRoleArn *string `field:"required" json:"provisioningRoleArn" yaml:"provisioningRoleArn"`
	// The JSON formatted contents of the fleet provisioning template version.
	TemplateBody *string `field:"required" json:"templateBody" yaml:"templateBody"`
	// The description of the fleet provisioning template.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// True to enable the fleet provisioning template, otherwise false.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Creates a pre-provisioning hook template.
	PreProvisioningHook interface{} `field:"optional" json:"preProvisioningHook" yaml:"preProvisioningHook"`
	// Metadata that can be used to manage the fleet provisioning template.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the fleet provisioning template.
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
	// The type of the provisioning template.
	TemplateType *string `field:"optional" json:"templateType" yaml:"templateType"`
}
