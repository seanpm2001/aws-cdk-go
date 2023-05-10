package cxapi

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/cloudassemblyschema"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cloudAssembly cloudAssembly
//
//   cloudFormationStackArtifact := awscdk.Cx_api.NewCloudFormationStackArtifact(cloudAssembly, jsii.String("artifactId"), &ArtifactManifest{
//   	Type: awscdk.Cloud_assembly_schema.ArtifactType_NONE,
//
//   	// the properties below are optional
//   	Dependencies: []*string{
//   		jsii.String("dependencies"),
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	Environment: jsii.String("environment"),
//   	Metadata: map[string][]metadataEntry{
//   		"metadataKey": []*metadataEntry{
//   			&metadataEntry{
//   				"type": jsii.String("type"),
//
//   				// the properties below are optional
//   				"data": jsii.String("data"),
//   				"trace": []*string{
//   					jsii.String("trace"),
//   				},
//   			},
//   		},
//   	},
//   	Properties: &AwsCloudFormationStackProperties{
//   		TemplateFile: jsii.String("templateFile"),
//
//   		// the properties below are optional
//   		AssumeRoleArn: jsii.String("assumeRoleArn"),
//   		AssumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   		BootstrapStackVersionSsmParameter: jsii.String("bootstrapStackVersionSsmParameter"),
//   		CloudFormationExecutionRoleArn: jsii.String("cloudFormationExecutionRoleArn"),
//   		LookupRole: &BootstrapRole{
//   			Arn: jsii.String("arn"),
//
//   			// the properties below are optional
//   			AssumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   			BootstrapStackVersionSsmParameter: jsii.String("bootstrapStackVersionSsmParameter"),
//   			RequiresBootstrapStackVersion: jsii.Number(123),
//   		},
//   		Parameters: map[string]*string{
//   			"parametersKey": jsii.String("parameters"),
//   		},
//   		RequiresBootstrapStackVersion: jsii.Number(123),
//   		StackName: jsii.String("stackName"),
//   		StackTemplateAssetObjectUrl: jsii.String("stackTemplateAssetObjectUrl"),
//   		Tags: map[string]*string{
//   			"tagsKey": jsii.String("tags"),
//   		},
//   		TerminationProtection: jsii.Boolean(false),
//   		ValidateOnSynth: jsii.Boolean(false),
//   	},
//   })
//
// Experimental.
type CloudFormationStackArtifact interface {
	CloudArtifact
	// Experimental.
	Assembly() CloudAssembly
	// Any assets associated with this stack.
	// Experimental.
	Assets() *[]interface{}
	// The role that needs to be assumed to deploy the stack.
	// Experimental.
	AssumeRoleArn() *string
	// External ID to use when assuming role for cloudformation deployments.
	// Experimental.
	AssumeRoleExternalId() *string
	// Name of SSM parameter with bootstrap stack version.
	// Experimental.
	BootstrapStackVersionSsmParameter() *string
	// The role that is passed to CloudFormation to execute the change set.
	// Experimental.
	CloudFormationExecutionRoleArn() *string
	// Returns all the artifacts that this artifact depends on.
	// Experimental.
	Dependencies() *[]CloudArtifact
	// A string that represents this stack.
	//
	// Should only be used in user
	// interfaces. If the stackName has not been set explicitly, or has been set
	// to artifactId, it will return the hierarchicalId of the stack. Otherwise,
	// it will return something like "<hierarchicalId> (<stackName>)".
	// Experimental.
	DisplayName() *string
	// The environment into which to deploy this artifact.
	// Experimental.
	Environment() *Environment
	// An identifier that shows where this artifact is located in the tree of nested assemblies, based on their manifests.
	//
	// Defaults to the normal
	// id. Should only be used in user interfaces.
	// Experimental.
	HierarchicalId() *string
	// Experimental.
	Id() *string
	// The role to use to look up values from the target AWS account.
	// Experimental.
	LookupRole() *cloudassemblyschema.BootstrapRole
	// The artifact's manifest.
	// Experimental.
	Manifest() *cloudassemblyschema.ArtifactManifest
	// The set of messages extracted from the artifact's metadata.
	// Experimental.
	Messages() *[]*SynthesisMessage
	// The physical name of this stack.
	// Deprecated: renamed to `stackName`.
	Name() *string
	// The original name as defined in the CDK app.
	// Experimental.
	OriginalName() *string
	// CloudFormation parameters to pass to the stack.
	// Experimental.
	Parameters() *map[string]*string
	// Version of bootstrap stack required to deploy this stack.
	// Experimental.
	RequiresBootstrapStackVersion() *float64
	// The physical name of this stack.
	// Experimental.
	StackName() *string
	// If the stack template has already been included in the asset manifest, its asset URL.
	// Experimental.
	StackTemplateAssetObjectUrl() *string
	// CloudFormation tags to pass to the stack.
	// Experimental.
	Tags() *map[string]*string
	// The CloudFormation template for this stack.
	// Experimental.
	Template() interface{}
	// The file name of the template.
	// Experimental.
	TemplateFile() *string
	// Full path to the template file.
	// Experimental.
	TemplateFullPath() *string
	// Whether termination protection is enabled for this stack.
	// Experimental.
	TerminationProtection() *bool
	// Whether this stack should be validated by the CLI after synthesis.
	// Experimental.
	ValidateOnSynth() *bool
	// Returns: all the metadata entries of a specific type in this artifact.
	// Experimental.
	FindMetadataByType(type_ *string) *[]*MetadataEntryResult
}

// The jsii proxy struct for CloudFormationStackArtifact
type jsiiProxy_CloudFormationStackArtifact struct {
	jsiiProxy_CloudArtifact
}

func (j *jsiiProxy_CloudFormationStackArtifact) Assembly() CloudAssembly {
	var returns CloudAssembly
	_jsii_.Get(
		j,
		"assembly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) Assets() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"assets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) AssumeRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) AssumeRoleExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleExternalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) BootstrapStackVersionSsmParameter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapStackVersionSsmParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) CloudFormationExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudFormationExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) Dependencies() *[]CloudArtifact {
	var returns *[]CloudArtifact
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) Environment() *Environment {
	var returns *Environment
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) HierarchicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) LookupRole() *cloudassemblyschema.BootstrapRole {
	var returns *cloudassemblyschema.BootstrapRole
	_jsii_.Get(
		j,
		"lookupRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) Manifest() *cloudassemblyschema.ArtifactManifest {
	var returns *cloudassemblyschema.ArtifactManifest
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) Messages() *[]*SynthesisMessage {
	var returns *[]*SynthesisMessage
	_jsii_.Get(
		j,
		"messages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) OriginalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) RequiresBootstrapStackVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiresBootstrapStackVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) StackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) StackTemplateAssetObjectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackTemplateAssetObjectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) Template() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"template",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) TemplateFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) TemplateFullPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateFullPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) TerminationProtection() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"terminationProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) ValidateOnSynth() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"validateOnSynth",
		&returns,
	)
	return returns
}


// Experimental.
func NewCloudFormationStackArtifact(assembly CloudAssembly, artifactId *string, artifact *cloudassemblyschema.ArtifactManifest) CloudFormationStackArtifact {
	_init_.Initialize()

	if err := validateNewCloudFormationStackArtifactParameters(assembly, artifactId, artifact); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudFormationStackArtifact{}

	_jsii_.Create(
		"monocdk.cx_api.CloudFormationStackArtifact",
		[]interface{}{assembly, artifactId, artifact},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudFormationStackArtifact_Override(c CloudFormationStackArtifact, assembly CloudAssembly, artifactId *string, artifact *cloudassemblyschema.ArtifactManifest) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.cx_api.CloudFormationStackArtifact",
		[]interface{}{assembly, artifactId, artifact},
		c,
	)
}

// Returns a subclass of `CloudArtifact` based on the artifact type defined in the artifact manifest.
//
// Returns: the `CloudArtifact` that matches the artifact type or `undefined` if it's an artifact type that is unrecognized by this module.
// Experimental.
func CloudFormationStackArtifact_FromManifest(assembly CloudAssembly, id *string, artifact *cloudassemblyschema.ArtifactManifest) CloudArtifact {
	_init_.Initialize()

	if err := validateCloudFormationStackArtifact_FromManifestParameters(assembly, id, artifact); err != nil {
		panic(err)
	}
	var returns CloudArtifact

	_jsii_.StaticInvoke(
		"monocdk.cx_api.CloudFormationStackArtifact",
		"fromManifest",
		[]interface{}{assembly, id, artifact},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationStackArtifact) FindMetadataByType(type_ *string) *[]*MetadataEntryResult {
	if err := c.validateFindMetadataByTypeParameters(type_); err != nil {
		panic(err)
	}
	var returns *[]*MetadataEntryResult

	_jsii_.Invoke(
		c,
		"findMetadataByType",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

