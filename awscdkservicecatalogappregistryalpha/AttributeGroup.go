// The CDK Construct Library for AWS::ServiceCatalogAppRegistry
package awscdkservicecatalogappregistryalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkservicecatalogappregistryalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkservicecatalogappregistryalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Service Catalog AppRegistry Attribute Group.
//
// Example:
//   attributeGroup := appreg.NewAttributeGroup(this, jsii.String("MyFirstAttributeGroup"), &AttributeGroupProps{
//   	AttributeGroupName: jsii.String("MyFirstAttributeGroupName"),
//   	Description: jsii.String("description for my attribute group"),
//   	 // the description is optional,
//   	Attributes: map[string]interface{}{
//   		"project": jsii.String("foo"),
//   		"team": []interface{}{
//   			jsii.String("member1"),
//   			jsii.String("member2"),
//   			jsii.String("member3"),
//   		},
//   		"public": jsii.Boolean(false),
//   		"stages": map[string]*string{
//   			"alpha": jsii.String("complete"),
//   			"beta": jsii.String("incomplete"),
//   			"release": jsii.String("not started"),
//   		},
//   	},
//   })
//
// Experimental.
type AttributeGroup interface {
	awscdk.Resource
	IAttributeGroup
	// The ARN of the attribute group.
	// Experimental.
	AttributeGroupArn() *string
	// The ID of the attribute group.
	// Experimental.
	AttributeGroupId() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Associate an application with attribute group If the attribute group is already associated, it will ignore duplicate request.
	// Experimental.
	AssociateWith(application IApplication)
	// Experimental.
	GeneratePhysicalName() *string
	// Create a unique hash.
	// Experimental.
	GenerateUniqueHash(resourceAddress *string) *string
	// Get the correct permission ARN based on the SharePermission.
	// Experimental.
	GetAttributeGroupSharePermissionARN(shareOptions *ShareOptions) *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Share the attribute group resource with other IAM entities, accounts, or OUs.
	// Experimental.
	ShareAttributeGroup(id *string, shareOptions *ShareOptions)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AttributeGroup
type jsiiProxy_AttributeGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_IAttributeGroup
}

func (j *jsiiProxy_AttributeGroup) AttributeGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttributeGroup) AttributeGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttributeGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttributeGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttributeGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttributeGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewAttributeGroup(scope constructs.Construct, id *string, props *AttributeGroupProps) AttributeGroup {
	_init_.Initialize()

	if err := validateNewAttributeGroupParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AttributeGroup{}

	_jsii_.Create(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.AttributeGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAttributeGroup_Override(a AttributeGroup, scope constructs.Construct, id *string, props *AttributeGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.AttributeGroup",
		[]interface{}{scope, id, props},
		a,
	)
}

// Imports an attribute group construct that represents an external attribute group.
// Experimental.
func AttributeGroup_FromAttributeGroupArn(scope constructs.Construct, id *string, attributeGroupArn *string) IAttributeGroup {
	_init_.Initialize()

	if err := validateAttributeGroup_FromAttributeGroupArnParameters(scope, id, attributeGroupArn); err != nil {
		panic(err)
	}
	var returns IAttributeGroup

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.AttributeGroup",
		"fromAttributeGroupArn",
		[]interface{}{scope, id, attributeGroupArn},
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
// Experimental.
func AttributeGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAttributeGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.AttributeGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func AttributeGroup_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateAttributeGroup_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.AttributeGroup",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func AttributeGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateAttributeGroup_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.AttributeGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttributeGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := a.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_AttributeGroup) AssociateWith(application IApplication) {
	if err := a.validateAssociateWithParameters(application); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"associateWith",
		[]interface{}{application},
	)
}

func (a *jsiiProxy_AttributeGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttributeGroup) GenerateUniqueHash(resourceAddress *string) *string {
	if err := a.validateGenerateUniqueHashParameters(resourceAddress); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"generateUniqueHash",
		[]interface{}{resourceAddress},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttributeGroup) GetAttributeGroupSharePermissionARN(shareOptions *ShareOptions) *string {
	if err := a.validateGetAttributeGroupSharePermissionARNParameters(shareOptions); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getAttributeGroupSharePermissionARN",
		[]interface{}{shareOptions},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttributeGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := a.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttributeGroup) GetResourceNameAttribute(nameAttr *string) *string {
	if err := a.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttributeGroup) ShareAttributeGroup(id *string, shareOptions *ShareOptions) {
	if err := a.validateShareAttributeGroupParameters(id, shareOptions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"shareAttributeGroup",
		[]interface{}{id, shareOptions},
	)
}

func (a *jsiiProxy_AttributeGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

