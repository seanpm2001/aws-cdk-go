package awsomics


// Properties for defining a `CfnWorkflow`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkflowProps := &CfnWorkflowProps{
//   	DefinitionUri: jsii.String("definitionUri"),
//   	Description: jsii.String("description"),
//   	Engine: jsii.String("engine"),
//   	Main: jsii.String("main"),
//   	Name: jsii.String("name"),
//   	ParameterTemplate: map[string]interface{}{
//   		"parameterTemplateKey": &WorkflowParameterProperty{
//   			"description": jsii.String("description"),
//   			"optional": jsii.Boolean(false),
//   		},
//   	},
//   	StorageCapacity: jsii.Number(123),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnWorkflowProps struct {
	// The URI of a definition for the workflow.
	DefinitionUri *string `field:"optional" json:"definitionUri" yaml:"definitionUri"`
	// The parameter's description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An engine for the workflow.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The path of the main definition file for the workflow.
	Main *string `field:"optional" json:"main" yaml:"main"`
	// The workflow's name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The workflow's parameter template.
	ParameterTemplate interface{} `field:"optional" json:"parameterTemplate" yaml:"parameterTemplate"`
	// A storage capacity for the workflow in gigabytes.
	StorageCapacity *float64 `field:"optional" json:"storageCapacity" yaml:"storageCapacity"`
	// Tags for the workflow.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

