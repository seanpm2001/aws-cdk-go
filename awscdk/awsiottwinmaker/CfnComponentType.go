package awsiottwinmaker

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiottwinmaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::IoTTwinMaker::ComponentType`.
//
// Use the `AWS::IoTTwinMaker::ComponentType` resource to declare a component type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataTypeProperty_ dataTypeProperty
//   var dataValueProperty_ dataValueProperty
//   var relationshipValue interface{}
//
//   cfnComponentType := awscdk.Aws_iottwinmaker.NewCfnComponentType(this, jsii.String("MyCfnComponentType"), &CfnComponentTypeProps{
//   	ComponentTypeId: jsii.String("componentTypeId"),
//   	WorkspaceId: jsii.String("workspaceId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	ExtendsFrom: []*string{
//   		jsii.String("extendsFrom"),
//   	},
//   	Functions: map[string]interface{}{
//   		"functionsKey": &FunctionProperty{
//   			"implementedBy": &DataConnectorProperty{
//   				"isNative": jsii.Boolean(false),
//   				"lambda": &LambdaFunctionProperty{
//   					"arn": jsii.String("arn"),
//   				},
//   			},
//   			"requiredProperties": []*string{
//   				jsii.String("requiredProperties"),
//   			},
//   			"scope": jsii.String("scope"),
//   		},
//   	},
//   	IsSingleton: jsii.Boolean(false),
//   	PropertyDefinitions: map[string]interface{}{
//   		"propertyDefinitionsKey": &PropertyDefinitionProperty{
//   			"configurations": map[string]*string{
//   				"configurationsKey": jsii.String("configurations"),
//   			},
//   			"dataType": &dataTypeProperty{
//   				"type": jsii.String("type"),
//
//   				// the properties below are optional
//   				"allowedValues": []interface{}{
//   					&dataValueProperty{
//   						"booleanValue": jsii.Boolean(false),
//   						"doubleValue": jsii.Number(123),
//   						"expression": jsii.String("expression"),
//   						"integerValue": jsii.Number(123),
//   						"listValue": []interface{}{
//   							dataValueProperty_,
//   						},
//   						"longValue": jsii.Number(123),
//   						"mapValue": map[string]interface{}{
//   							"mapValueKey": dataValueProperty_,
//   						},
//   						"relationshipValue": relationshipValue,
//   						"stringValue": jsii.String("stringValue"),
//   					},
//   				},
//   				"nestedType": dataTypeProperty_,
//   				"relationship": &RelationshipProperty{
//   					"relationshipType": jsii.String("relationshipType"),
//   					"targetComponentTypeId": jsii.String("targetComponentTypeId"),
//   				},
//   				"unitOfMeasure": jsii.String("unitOfMeasure"),
//   			},
//   			"defaultValue": &dataValueProperty{
//   				"booleanValue": jsii.Boolean(false),
//   				"doubleValue": jsii.Number(123),
//   				"expression": jsii.String("expression"),
//   				"integerValue": jsii.Number(123),
//   				"listValue": []interface{}{
//   					dataValueProperty_,
//   				},
//   				"longValue": jsii.Number(123),
//   				"mapValue": map[string]interface{}{
//   					"mapValueKey": dataValueProperty_,
//   				},
//   				"relationshipValue": relationshipValue,
//   				"stringValue": jsii.String("stringValue"),
//   			},
//   			"isExternalId": jsii.Boolean(false),
//   			"isRequiredInEntity": jsii.Boolean(false),
//   			"isStoredExternally": jsii.Boolean(false),
//   			"isTimeSeries": jsii.Boolean(false),
//   		},
//   	},
//   	PropertyGroups: map[string]interface{}{
//   		"propertyGroupsKey": &PropertyGroupProperty{
//   			"groupType": jsii.String("groupType"),
//   			"propertyNames": []*string{
//   				jsii.String("propertyNames"),
//   			},
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnComponentType interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ARN of the component type.
	AttrArn() *string
	// The date and time when the component type was created.
	AttrCreationDateTime() *string
	// A boolean value that specifies whether the component type is abstract.
	AttrIsAbstract() awscdk.IResolvable
	// A boolean value that specifies whether the component type has a schema initializer and that the schema initializer has run.
	AttrIsSchemaInitialized() awscdk.IResolvable
	AttrStatusErrorCode() *string
	AttrStatusErrorMessage() *string
	AttrStatusState() *string
	// The component type the update time.
	AttrUpdateDateTime() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The ID of the component type.
	ComponentTypeId() *string
	SetComponentTypeId(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the component type.
	Description() *string
	SetDescription(val *string)
	// The name of the parent component type that this component type extends.
	ExtendsFrom() *[]*string
	SetExtendsFrom(val *[]*string)
	// An object that maps strings to the functions in the component type.
	//
	// Each string in the mapping must be unique to this object.
	//
	// For information on the FunctionResponse object see the [FunctionResponse](https://docs.aws.amazon.com//iot-twinmaker/latest/apireference/API_FunctionResponse.html) API reference.
	Functions() interface{}
	SetFunctions(val interface{})
	// A boolean value that specifies whether an entity can have more than one component of this type.
	IsSingleton() interface{}
	SetIsSingleton(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// An object that maps strings to the property definitions in the component type.
	//
	// Each string in the mapping must be unique to this object.
	//
	// For information about the PropertyDefinitionResponse object, see the [PropertyDefinitionResponse](https://docs.aws.amazon.com//iot-twinmaker/latest/apireference/API_PropertyDefinitionResponse.html) API reference.
	PropertyDefinitions() interface{}
	SetPropertyDefinitions(val interface{})
	// An object that maps strings to the property groups in the component type.
	//
	// Each string in the mapping must be unique to this object.
	PropertyGroups() interface{}
	SetPropertyGroups(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The ComponentType tags.
	Tags() awscdk.TagManager
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// The ID of the workspace.
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnComponentType
type jsiiProxy_CfnComponentType struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnComponentType) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) AttrCreationDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreationDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) AttrIsAbstract() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrIsAbstract",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) AttrIsSchemaInitialized() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrIsSchemaInitialized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) AttrStatusErrorCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatusErrorCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) AttrStatusErrorMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatusErrorMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) AttrStatusState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatusState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) AttrUpdateDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUpdateDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) ComponentTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"componentTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) ExtendsFrom() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"extendsFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) Functions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"functions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) IsSingleton() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSingleton",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) PropertyDefinitions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"propertyDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) PropertyGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"propertyGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentType) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}


// Create a new `AWS::IoTTwinMaker::ComponentType`.
func NewCfnComponentType(scope constructs.Construct, id *string, props *CfnComponentTypeProps) CfnComponentType {
	_init_.Initialize()

	if err := validateNewCfnComponentTypeParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnComponentType{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iottwinmaker.CfnComponentType",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::IoTTwinMaker::ComponentType`.
func NewCfnComponentType_Override(c CfnComponentType, scope constructs.Construct, id *string, props *CfnComponentTypeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iottwinmaker.CfnComponentType",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnComponentType)SetComponentTypeId(val *string) {
	if err := j.validateSetComponentTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"componentTypeId",
		val,
	)
}

func (j *jsiiProxy_CfnComponentType)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnComponentType)SetExtendsFrom(val *[]*string) {
	_jsii_.Set(
		j,
		"extendsFrom",
		val,
	)
}

func (j *jsiiProxy_CfnComponentType)SetFunctions(val interface{}) {
	if err := j.validateSetFunctionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functions",
		val,
	)
}

func (j *jsiiProxy_CfnComponentType)SetIsSingleton(val interface{}) {
	if err := j.validateSetIsSingletonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSingleton",
		val,
	)
}

func (j *jsiiProxy_CfnComponentType)SetPropertyDefinitions(val interface{}) {
	if err := j.validateSetPropertyDefinitionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propertyDefinitions",
		val,
	)
}

func (j *jsiiProxy_CfnComponentType)SetPropertyGroups(val interface{}) {
	if err := j.validateSetPropertyGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propertyGroups",
		val,
	)
}

func (j *jsiiProxy_CfnComponentType)SetWorkspaceId(val *string) {
	if err := j.validateSetWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnComponentType_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnComponentType_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iottwinmaker.CfnComponentType",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnComponentType_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnComponentType_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iottwinmaker.CfnComponentType",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnComponentType_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnComponentType_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iottwinmaker.CfnComponentType",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnComponentType_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iottwinmaker.CfnComponentType",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnComponentType) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnComponentType) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnComponentType) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnComponentType) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnComponentType) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnComponentType) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnComponentType) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnComponentType) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnComponentType) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponentType) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponentType) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnComponentType) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponentType) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponentType) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnComponentType) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnComponentType) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponentType) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnComponentType) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponentType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponentType) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}
