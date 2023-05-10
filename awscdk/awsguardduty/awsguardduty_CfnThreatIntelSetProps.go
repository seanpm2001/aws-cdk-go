package awsguardduty

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnThreatIntelSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnThreatIntelSetProps := &CfnThreatIntelSetProps{
//   	Activate: jsii.Boolean(false),
//   	DetectorId: jsii.String("detectorId"),
//   	Format: jsii.String("format"),
//   	Location: jsii.String("location"),
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnThreatIntelSetProps struct {
	// A Boolean value that indicates whether GuardDuty is to start using the uploaded ThreatIntelSet.
	Activate interface{} `field:"required" json:"activate" yaml:"activate"`
	// The unique ID of the detector of the GuardDuty account that you want to create a threatIntelSet for.
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// The format of the file that contains the ThreatIntelSet.
	Format *string `field:"required" json:"format" yaml:"format"`
	// The URI of the file that contains the ThreatIntelSet.
	Location *string `field:"required" json:"location" yaml:"location"`
	// A user-friendly ThreatIntelSet name displayed in all findings that are generated by activity that involves IP addresses included in this ThreatIntelSet.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags to be added to a new threat list resource.
	//
	// Each tag consists of a key and an optional value, both of which you define.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

