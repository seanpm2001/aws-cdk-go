package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var mesh mesh
//
//   router := mesh.addVirtualRouter(jsii.String("router"), &virtualRouterBaseProps{
//   	listeners: []virtualRouterListener{
//   		appmesh.*virtualRouterListener.http(jsii.Number(8080)),
//   	},
//   })
//
type VirtualRouter interface {
	awscdk.Resource
	IVirtualRouter
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The Mesh which the VirtualRouter belongs to.
	Mesh() IMesh
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The Amazon Resource Name (ARN) for the VirtualRouter.
	VirtualRouterArn() *string
	// The name of the VirtualRouter.
	VirtualRouterName() *string
	// Add a single route to the router.
	AddRoute(id *string, props *RouteBaseProps) Route
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for VirtualRouter
type jsiiProxy_VirtualRouter struct {
	internal.Type__awscdkResource
	jsiiProxy_IVirtualRouter
}

func (j *jsiiProxy_VirtualRouter) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualRouter) Mesh() IMesh {
	var returns IMesh
	_jsii_.Get(
		j,
		"mesh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualRouter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualRouter) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualRouter) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualRouter) VirtualRouterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualRouterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualRouter) VirtualRouterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualRouterName",
		&returns,
	)
	return returns
}


func NewVirtualRouter(scope constructs.Construct, id *string, props *VirtualRouterProps) VirtualRouter {
	_init_.Initialize()

	if err := validateNewVirtualRouterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualRouter{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.VirtualRouter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewVirtualRouter_Override(v VirtualRouter, scope constructs.Construct, id *string, props *VirtualRouterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.VirtualRouter",
		[]interface{}{scope, id, props},
		v,
	)
}

// Import an existing VirtualRouter given an ARN.
func VirtualRouter_FromVirtualRouterArn(scope constructs.Construct, id *string, virtualRouterArn *string) IVirtualRouter {
	_init_.Initialize()

	if err := validateVirtualRouter_FromVirtualRouterArnParameters(scope, id, virtualRouterArn); err != nil {
		panic(err)
	}
	var returns IVirtualRouter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualRouter",
		"fromVirtualRouterArn",
		[]interface{}{scope, id, virtualRouterArn},
		&returns,
	)

	return returns
}

// Import an existing VirtualRouter given attributes.
func VirtualRouter_FromVirtualRouterAttributes(scope constructs.Construct, id *string, attrs *VirtualRouterAttributes) IVirtualRouter {
	_init_.Initialize()

	if err := validateVirtualRouter_FromVirtualRouterAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IVirtualRouter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualRouter",
		"fromVirtualRouterAttributes",
		[]interface{}{scope, id, attrs},
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
func VirtualRouter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualRouter_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualRouter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func VirtualRouter_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateVirtualRouter_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualRouter",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func VirtualRouter_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateVirtualRouter_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualRouter",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualRouter) AddRoute(id *string, props *RouteBaseProps) Route {
	if err := v.validateAddRouteParameters(id, props); err != nil {
		panic(err)
	}
	var returns Route

	_jsii_.Invoke(
		v,
		"addRoute",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualRouter) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := v.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (v *jsiiProxy_VirtualRouter) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualRouter) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := v.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualRouter) GetResourceNameAttribute(nameAttr *string) *string {
	if err := v.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualRouter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

