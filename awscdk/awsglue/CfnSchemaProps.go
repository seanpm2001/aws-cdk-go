package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSchema`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSchemaProps := &CfnSchemaProps{
//   	Compatibility: jsii.String("compatibility"),
//   	DataFormat: jsii.String("dataFormat"),
//   	Name: jsii.String("name"),
//   	SchemaDefinition: jsii.String("schemaDefinition"),
//
//   	// the properties below are optional
//   	CheckpointVersion: &SchemaVersionProperty{
//   		IsLatest: jsii.Boolean(false),
//   		VersionNumber: jsii.Number(123),
//   	},
//   	Description: jsii.String("description"),
//   	Registry: &RegistryProperty{
//   		Arn: jsii.String("arn"),
//   		Name: jsii.String("name"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnSchemaProps struct {
	// The compatibility mode of the schema.
	Compatibility *string `field:"required" json:"compatibility" yaml:"compatibility"`
	// The data format of the schema definition.
	//
	// Currently only `AVRO` is supported.
	DataFormat *string `field:"required" json:"dataFormat" yaml:"dataFormat"`
	// Name of the schema to be created of max length of 255, and may only contain letters, numbers, hyphen, underscore, dollar sign, or hash mark.
	//
	// No whitespace.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema definition using the `DataFormat` setting for `SchemaName` .
	SchemaDefinition *string `field:"required" json:"schemaDefinition" yaml:"schemaDefinition"`
	// Specify the `VersionNumber` or the `IsLatest` for setting the checkpoint for the schema.
	//
	// This is only required for updating a checkpoint.
	CheckpointVersion interface{} `field:"optional" json:"checkpointVersion" yaml:"checkpointVersion"`
	// A description of the schema if specified when created.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The registry where a schema is stored.
	Registry interface{} `field:"optional" json:"registry" yaml:"registry"`
	// AWS tags that contain a key value pair and may be searched by console, command line, or API.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}
