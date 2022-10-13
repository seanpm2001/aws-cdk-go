package awsses

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::SES::EmailIdentity`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEmailIdentity := awscdk.Aws_ses.NewCfnEmailIdentity(this, jsii.String("MyCfnEmailIdentity"), &cfnEmailIdentityProps{
//   	emailIdentity: jsii.String("emailIdentity"),
//
//   	// the properties below are optional
//   	configurationSetAttributes: &configurationSetAttributesProperty{
//   		configurationSetName: jsii.String("configurationSetName"),
//   	},
//   	dkimAttributes: &dkimAttributesProperty{
//   		signingEnabled: jsii.Boolean(false),
//   	},
//   	dkimSigningAttributes: &dkimSigningAttributesProperty{
//   		domainSigningPrivateKey: jsii.String("domainSigningPrivateKey"),
//   		domainSigningSelector: jsii.String("domainSigningSelector"),
//   		nextSigningKeyLength: jsii.String("nextSigningKeyLength"),
//   	},
//   	feedbackAttributes: &feedbackAttributesProperty{
//   		emailForwardingEnabled: jsii.Boolean(false),
//   	},
//   	mailFromAttributes: &mailFromAttributesProperty{
//   		behaviorOnMxFailure: jsii.String("behaviorOnMxFailure"),
//   		mailFromDomain: jsii.String("mailFromDomain"),
//   	},
//   })
//
type CfnEmailIdentity interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrDkimDnsTokenName1() *string
	AttrDkimDnsTokenName2() *string
	AttrDkimDnsTokenName3() *string
	AttrDkimDnsTokenValue1() *string
	AttrDkimDnsTokenValue2() *string
	AttrDkimDnsTokenValue3() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// `AWS::SES::EmailIdentity.ConfigurationSetAttributes`.
	ConfigurationSetAttributes() interface{}
	SetConfigurationSetAttributes(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// `AWS::SES::EmailIdentity.DkimAttributes`.
	DkimAttributes() interface{}
	SetDkimAttributes(val interface{})
	// `AWS::SES::EmailIdentity.DkimSigningAttributes`.
	DkimSigningAttributes() interface{}
	SetDkimSigningAttributes(val interface{})
	// `AWS::SES::EmailIdentity.EmailIdentity`.
	EmailIdentity() *string
	SetEmailIdentity(val *string)
	// `AWS::SES::EmailIdentity.FeedbackAttributes`.
	FeedbackAttributes() interface{}
	SetFeedbackAttributes(val interface{})
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
	// `AWS::SES::EmailIdentity.MailFromAttributes`.
	MailFromAttributes() interface{}
	SetMailFromAttributes(val interface{})
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnEmailIdentity
type jsiiProxy_CfnEmailIdentity struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEmailIdentity) AttrDkimDnsTokenName1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDkimDnsTokenName1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) AttrDkimDnsTokenName2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDkimDnsTokenName2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) AttrDkimDnsTokenName3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDkimDnsTokenName3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) AttrDkimDnsTokenValue1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDkimDnsTokenValue1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) AttrDkimDnsTokenValue2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDkimDnsTokenValue2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) AttrDkimDnsTokenValue3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDkimDnsTokenValue3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) ConfigurationSetAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurationSetAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) DkimAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dkimAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) DkimSigningAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dkimSigningAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) EmailIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) FeedbackAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"feedbackAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) MailFromAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mailFromAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentity) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::SES::EmailIdentity`.
func NewCfnEmailIdentity(scope constructs.Construct, id *string, props *CfnEmailIdentityProps) CfnEmailIdentity {
	_init_.Initialize()

	if err := validateNewCfnEmailIdentityParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEmailIdentity{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.CfnEmailIdentity",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::SES::EmailIdentity`.
func NewCfnEmailIdentity_Override(c CfnEmailIdentity, scope constructs.Construct, id *string, props *CfnEmailIdentityProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.CfnEmailIdentity",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEmailIdentity)SetConfigurationSetAttributes(val interface{}) {
	if err := j.validateSetConfigurationSetAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationSetAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnEmailIdentity)SetDkimAttributes(val interface{}) {
	if err := j.validateSetDkimAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dkimAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnEmailIdentity)SetDkimSigningAttributes(val interface{}) {
	if err := j.validateSetDkimSigningAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dkimSigningAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnEmailIdentity)SetEmailIdentity(val *string) {
	if err := j.validateSetEmailIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailIdentity",
		val,
	)
}

func (j *jsiiProxy_CfnEmailIdentity)SetFeedbackAttributes(val interface{}) {
	if err := j.validateSetFeedbackAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"feedbackAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnEmailIdentity)SetMailFromAttributes(val interface{}) {
	if err := j.validateSetMailFromAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mailFromAttributes",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnEmailIdentity_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEmailIdentity_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnEmailIdentity",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnEmailIdentity_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnEmailIdentity_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnEmailIdentity",
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
func CfnEmailIdentity_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEmailIdentity_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.CfnEmailIdentity",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEmailIdentity_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ses.CfnEmailIdentity",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEmailIdentity) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEmailIdentity) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEmailIdentity) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEmailIdentity) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEmailIdentity) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEmailIdentity) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEmailIdentity) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnEmailIdentity) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnEmailIdentity) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnEmailIdentity) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEmailIdentity) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEmailIdentity) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnEmailIdentity) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEmailIdentity) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEmailIdentity) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

