// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgameliftalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Amazon GameLift alias is used to abstract a fleet designation.
//
// Fleet designations tell GameLift where to search for available resources when creating new game sessions for players.
// Use aliases instead of specific fleet IDs to seamlessly switch player traffic from one fleet to another by changing the alias's target location.
//
// Aliases are useful in games that don't use queues.
// Switching fleets in a queue is a simple matter of creating a new fleet, adding it to the queue, and removing the old fleet, none of which is visible to players.
// In contrast, game clients that don't use queues must specify which fleet to use when communicating with the GameLift service.
// Without aliases, a fleet switch requires updates to your game code and possibly distribution of an updated game clients to players.
//
// When updating the fleet-id an alias points to, there is a transition period of up to 2 minutes where game sessions on the alias may end up on the old fleet.
//
// Example:
//   var fleet buildFleet
//
//
//   // Add an alias to an existing fleet using a dedicated fleet method
//   liveAlias := fleet.addAlias(jsii.String("live"))
//
//   // You can also create a standalone alias
//   // You can also create a standalone alias
//   gamelift.NewAlias(this, jsii.String("TerminalAlias"), &aliasProps{
//   	aliasName: jsii.String("terminal-alias"),
//   	terminalMessage: jsii.String("A terminal message"),
//   })
//
// See: https://docs.aws.amazon.com/gamelift/latest/developerguide/aliases-creating.html
//
// Experimental.
type Alias interface {
	AliasBase
	// The ARN of the alias.
	// Experimental.
	AliasArn() *string
	// The Identifier of the alias.
	// Experimental.
	AliasId() *string
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
	// A fleet that the alias points to.
	// Experimental.
	Fleet() IFleet
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The ARN to put into the destination field of a game session queue.
	// Experimental.
	ResourceArnForDestination() *string
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

// The jsii proxy struct for Alias
type jsiiProxy_Alias struct {
	jsiiProxy_AliasBase
}

func (j *jsiiProxy_Alias) AliasArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) AliasId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) Fleet() IFleet {
	var returns IFleet
	_jsii_.Get(
		j,
		"fleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) ResourceArnForDestination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArnForDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewAlias(scope constructs.Construct, id *string, props *AliasProps) Alias {
	_init_.Initialize()

	if err := validateNewAliasParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Alias{}

	_jsii_.Create(
		"@aws-cdk/aws-gamelift-alpha.Alias",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAlias_Override(a Alias, scope constructs.Construct, id *string, props *AliasProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-gamelift-alpha.Alias",
		[]interface{}{scope, id, props},
		a,
	)
}

// Import an existing alias from its ARN.
// Experimental.
func Alias_FromAliasArn(scope constructs.Construct, id *string, aliasArn *string) IAlias {
	_init_.Initialize()

	if err := validateAlias_FromAliasArnParameters(scope, id, aliasArn); err != nil {
		panic(err)
	}
	var returns IAlias

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.Alias",
		"fromAliasArn",
		[]interface{}{scope, id, aliasArn},
		&returns,
	)

	return returns
}

// Import an existing alias from its attributes.
// Experimental.
func Alias_FromAliasAttributes(scope constructs.Construct, id *string, attrs *AliasAttributes) IAlias {
	_init_.Initialize()

	if err := validateAlias_FromAliasAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IAlias

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.Alias",
		"fromAliasAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import an existing alias from its identifier.
// Experimental.
func Alias_FromAliasId(scope constructs.Construct, id *string, aliasId *string) IAlias {
	_init_.Initialize()

	if err := validateAlias_FromAliasIdParameters(scope, id, aliasId); err != nil {
		panic(err)
	}
	var returns IAlias

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.Alias",
		"fromAliasId",
		[]interface{}{scope, id, aliasId},
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
func Alias_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlias_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.Alias",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func Alias_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateAlias_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.Alias",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Alias_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateAlias_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-gamelift-alpha.Alias",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Alias) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := a.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_Alias) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Alias) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (a *jsiiProxy_Alias) GetResourceNameAttribute(nameAttr *string) *string {
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

func (a *jsiiProxy_Alias) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
