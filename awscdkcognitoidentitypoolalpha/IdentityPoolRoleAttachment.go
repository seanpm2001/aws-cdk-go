package awscdkcognitoidentitypoolalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcognitoidentitypoolalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcognitoidentitypoolalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Defines an Identity Pool Role Attachment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcognitoidentitypoolalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var identityPool identityPool
//   var identityPoolProviderUrl identityPoolProviderUrl
//   var role role
//
//   identityPoolRoleAttachment := cognito_identitypool_alpha.NewIdentityPoolRoleAttachment(this, jsii.String("MyIdentityPoolRoleAttachment"), &IdentityPoolRoleAttachmentProps{
//   	IdentityPool: identityPool,
//
//   	// the properties below are optional
//   	AuthenticatedRole: role,
//   	RoleMappings: []identityPoolRoleMapping{
//   		&identityPoolRoleMapping{
//   			ProviderUrl: identityPoolProviderUrl,
//
//   			// the properties below are optional
//   			MappingKey: jsii.String("mappingKey"),
//   			ResolveAmbiguousRoles: jsii.Boolean(false),
//   			Rules: []roleMappingRule{
//   				&roleMappingRule{
//   					Claim: jsii.String("claim"),
//   					ClaimValue: jsii.String("claimValue"),
//   					MappedRole: role,
//
//   					// the properties below are optional
//   					MatchType: cognito_identitypool_alpha.RoleMappingMatchType_EQUALS,
//   				},
//   			},
//   			UseToken: jsii.Boolean(false),
//   		},
//   	},
//   	UnauthenticatedRole: role,
//   })
//
// Experimental.
type IdentityPoolRoleAttachment interface {
	awscdk.Resource
	IIdentityPoolRoleAttachment
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
	// Id of the underlying identity pool.
	// Experimental.
	IdentityPoolId() *string
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
	// Experimental.
	GeneratePhysicalName() *string
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
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IdentityPoolRoleAttachment
type jsiiProxy_IdentityPoolRoleAttachment struct {
	internal.Type__awscdkResource
	jsiiProxy_IIdentityPoolRoleAttachment
}

func (j *jsiiProxy_IdentityPoolRoleAttachment) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPoolRoleAttachment) IdentityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPoolRoleAttachment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPoolRoleAttachment) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPoolRoleAttachment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewIdentityPoolRoleAttachment(scope constructs.Construct, id *string, props *IdentityPoolRoleAttachmentProps) IdentityPoolRoleAttachment {
	_init_.Initialize()

	if err := validateNewIdentityPoolRoleAttachmentParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_IdentityPoolRoleAttachment{}

	_jsii_.Create(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolRoleAttachment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewIdentityPoolRoleAttachment_Override(i IdentityPoolRoleAttachment, scope constructs.Construct, id *string, props *IdentityPoolRoleAttachmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolRoleAttachment",
		[]interface{}{scope, id, props},
		i,
	)
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
func IdentityPoolRoleAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdentityPoolRoleAttachment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolRoleAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func IdentityPoolRoleAttachment_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateIdentityPoolRoleAttachment_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolRoleAttachment",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func IdentityPoolRoleAttachment_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateIdentityPoolRoleAttachment_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-cognito-identitypool-alpha.IdentityPoolRoleAttachment",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPoolRoleAttachment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IdentityPoolRoleAttachment) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPoolRoleAttachment) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := i.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPoolRoleAttachment) GetResourceNameAttribute(nameAttr *string) *string {
	if err := i.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPoolRoleAttachment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

