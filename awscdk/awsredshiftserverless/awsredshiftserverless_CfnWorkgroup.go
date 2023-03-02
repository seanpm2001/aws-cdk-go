package awsredshiftserverless

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsredshiftserverless/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::RedshiftServerless::Workgroup`.
//
// The collection of compute resources in Amazon Redshift Serverless.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkgroup := awscdk.Aws_redshiftserverless.NewCfnWorkgroup(this, jsii.String("MyCfnWorkgroup"), &cfnWorkgroupProps{
//   	workgroupName: jsii.String("workgroupName"),
//
//   	// the properties below are optional
//   	baseCapacity: jsii.Number(123),
//   	configParameters: []interface{}{
//   		&configParameterProperty{
//   			parameterKey: jsii.String("parameterKey"),
//   			parameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	enhancedVpcRouting: jsii.Boolean(false),
//   	namespaceName: jsii.String("namespaceName"),
//   	publiclyAccessible: jsii.Boolean(false),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	workgroup: &workgroupProperty{
//   		baseCapacity: jsii.Number(123),
//   		configParameters: []interface{}{
//   			&configParameterProperty{
//   				parameterKey: jsii.String("parameterKey"),
//   				parameterValue: jsii.String("parameterValue"),
//   			},
//   		},
//   		creationDate: jsii.String("creationDate"),
//   		endpoint: &endpointProperty{
//   			address: jsii.String("address"),
//   			port: jsii.Number(123),
//   			vpcEndpoints: []interface{}{
//   				&vpcEndpointProperty{
//   					networkInterfaces: []interface{}{
//   						&networkInterfaceProperty{
//   							availabilityZone: jsii.String("availabilityZone"),
//   							networkInterfaceId: jsii.String("networkInterfaceId"),
//   							privateIpAddress: jsii.String("privateIpAddress"),
//   							subnetId: jsii.String("subnetId"),
//   						},
//   					},
//   					vpcEndpointId: jsii.String("vpcEndpointId"),
//   					vpcId: jsii.String("vpcId"),
//   				},
//   			},
//   		},
//   		enhancedVpcRouting: jsii.Boolean(false),
//   		namespaceName: jsii.String("namespaceName"),
//   		publiclyAccessible: jsii.Boolean(false),
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		status: jsii.String("status"),
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		workgroupArn: jsii.String("workgroupArn"),
//   		workgroupId: jsii.String("workgroupId"),
//   		workgroupName: jsii.String("workgroupName"),
//   	},
//   })
//
type CfnWorkgroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrWorkgroupBaseCapacity() *float64
	AttrWorkgroupCreationDate() *string
	AttrWorkgroupEndpointAddress() *string
	AttrWorkgroupEndpointPort() *float64
	AttrWorkgroupEnhancedVpcRouting() awscdk.IResolvable
	AttrWorkgroupNamespaceName() *string
	AttrWorkgroupPubliclyAccessible() awscdk.IResolvable
	AttrWorkgroupSecurityGroupIds() *[]*string
	AttrWorkgroupStatus() *string
	AttrWorkgroupSubnetIds() *[]*string
	AttrWorkgroupWorkgroupArn() *string
	AttrWorkgroupWorkgroupId() *string
	AttrWorkgroupWorkgroupName() *string
	// The base compute capacity of the workgroup in Redshift Processing Units (RPUs).
	BaseCapacity() *float64
	SetBaseCapacity(val *float64)
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// A list of parameters to set for finer control over a database.
	//
	// Available options are `datestyle` , `enable_user_activity_logging` , `query_group` , `search_path` , and `max_query_execution_time` .
	ConfigParameters() interface{}
	SetConfigParameters(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The value that specifies whether to enable enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC.
	EnhancedVpcRouting() interface{}
	SetEnhancedVpcRouting(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The namespace the workgroup is associated with.
	NamespaceName() *string
	SetNamespaceName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// A value that specifies whether the workgroup can be accessible from a public network.
	PubliclyAccessible() interface{}
	SetPubliclyAccessible(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// A list of security group IDs to associate with the workgroup.
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A list of subnet IDs the workgroup is associated with.
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	// The map of the key-value pairs used to tag the workgroup.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// `AWS::RedshiftServerless::Workgroup.Workgroup`.
	Workgroup() interface{}
	SetWorkgroup(val interface{})
	// The name of the workgroup.
	WorkgroupName() *string
	SetWorkgroupName(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnWorkgroup
type jsiiProxy_CfnWorkgroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnWorkgroup) AttrWorkgroupBaseCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrWorkgroupBaseCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) AttrWorkgroupCreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrWorkgroupCreationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) AttrWorkgroupEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrWorkgroupEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) AttrWorkgroupEndpointPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrWorkgroupEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) AttrWorkgroupEnhancedVpcRouting() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrWorkgroupEnhancedVpcRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) AttrWorkgroupNamespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrWorkgroupNamespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) AttrWorkgroupPubliclyAccessible() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrWorkgroupPubliclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) AttrWorkgroupSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrWorkgroupSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) AttrWorkgroupStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrWorkgroupStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) AttrWorkgroupSubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrWorkgroupSubnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) AttrWorkgroupWorkgroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrWorkgroupWorkgroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) AttrWorkgroupWorkgroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrWorkgroupWorkgroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) AttrWorkgroupWorkgroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrWorkgroupWorkgroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) BaseCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"baseCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) ConfigParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) EnhancedVpcRouting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedVpcRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) NamespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) Workgroup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workgroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkgroup) WorkgroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workgroupName",
		&returns,
	)
	return returns
}


// Create a new `AWS::RedshiftServerless::Workgroup`.
func NewCfnWorkgroup(scope awscdk.Construct, id *string, props *CfnWorkgroupProps) CfnWorkgroup {
	_init_.Initialize()

	if err := validateNewCfnWorkgroupParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWorkgroup{}

	_jsii_.Create(
		"monocdk.aws_redshiftserverless.CfnWorkgroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::RedshiftServerless::Workgroup`.
func NewCfnWorkgroup_Override(c CfnWorkgroup, scope awscdk.Construct, id *string, props *CfnWorkgroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_redshiftserverless.CfnWorkgroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnWorkgroup)SetBaseCapacity(val *float64) {
	_jsii_.Set(
		j,
		"baseCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnWorkgroup)SetConfigParameters(val interface{}) {
	if err := j.validateSetConfigParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configParameters",
		val,
	)
}

func (j *jsiiProxy_CfnWorkgroup)SetEnhancedVpcRouting(val interface{}) {
	if err := j.validateSetEnhancedVpcRoutingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enhancedVpcRouting",
		val,
	)
}

func (j *jsiiProxy_CfnWorkgroup)SetNamespaceName(val *string) {
	_jsii_.Set(
		j,
		"namespaceName",
		val,
	)
}

func (j *jsiiProxy_CfnWorkgroup)SetPubliclyAccessible(val interface{}) {
	if err := j.validateSetPubliclyAccessibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_CfnWorkgroup)SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_CfnWorkgroup)SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_CfnWorkgroup)SetWorkgroup(val interface{}) {
	if err := j.validateSetWorkgroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workgroup",
		val,
	)
}

func (j *jsiiProxy_CfnWorkgroup)SetWorkgroupName(val *string) {
	if err := j.validateSetWorkgroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workgroupName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnWorkgroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkgroup_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_redshiftserverless.CfnWorkgroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnWorkgroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnWorkgroup_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_redshiftserverless.CfnWorkgroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnWorkgroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkgroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_redshiftserverless.CfnWorkgroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkgroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_redshiftserverless.CfnWorkgroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkgroup) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnWorkgroup) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnWorkgroup) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnWorkgroup) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnWorkgroup) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnWorkgroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnWorkgroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnWorkgroup) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkgroup) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnWorkgroup) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnWorkgroup) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnWorkgroup) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnWorkgroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkgroup) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnWorkgroup) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnWorkgroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnWorkgroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkgroup) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnWorkgroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkgroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkgroup) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

