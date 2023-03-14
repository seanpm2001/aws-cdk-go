package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Route represents a new or existing route attached to a VirtualRouter and Mesh.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var mesh mesh
//   var routeSpec routeSpec
//   var virtualRouter virtualRouter
//
//   route := awscdk.Aws_appmesh.NewRoute(this, jsii.String("MyRoute"), &RouteProps{
//   	Mesh: mesh,
//   	RouteSpec: routeSpec,
//   	VirtualRouter: virtualRouter,
//
//   	// the properties below are optional
//   	RouteName: jsii.String("routeName"),
//   })
//
// See: https://docs.aws.amazon.com/app-mesh/latest/userguide/routes.html
//
type Route interface {
	awscdk.Resource
	IRoute
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
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
	// The Amazon Resource Name (ARN) for the route.
	RouteArn() *string
	// The name of the Route.
	RouteName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The VirtualRouter the Route belongs to.
	VirtualRouter() IVirtualRouter
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

// The jsii proxy struct for Route
type jsiiProxy_Route struct {
	internal.Type__awscdkResource
	jsiiProxy_IRoute
}

func (j *jsiiProxy_Route) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route) RouteArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route) RouteName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route) VirtualRouter() IVirtualRouter {
	var returns IVirtualRouter
	_jsii_.Get(
		j,
		"virtualRouter",
		&returns,
	)
	return returns
}


func NewRoute(scope constructs.Construct, id *string, props *RouteProps) Route {
	_init_.Initialize()

	if err := validateNewRouteParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Route{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.Route",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewRoute_Override(r Route, scope constructs.Construct, id *string, props *RouteProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.Route",
		[]interface{}{scope, id, props},
		r,
	)
}

// Import an existing Route given an ARN.
func Route_FromRouteArn(scope constructs.Construct, id *string, routeArn *string) IRoute {
	_init_.Initialize()

	if err := validateRoute_FromRouteArnParameters(scope, id, routeArn); err != nil {
		panic(err)
	}
	var returns IRoute

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.Route",
		"fromRouteArn",
		[]interface{}{scope, id, routeArn},
		&returns,
	)

	return returns
}

// Import an existing Route given attributes.
func Route_FromRouteAttributes(scope constructs.Construct, id *string, attrs *RouteAttributes) IRoute {
	_init_.Initialize()

	if err := validateRoute_FromRouteAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IRoute

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.Route",
		"fromRouteAttributes",
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
func Route_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoute_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.Route",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Route_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateRoute_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.Route",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Route_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateRoute_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.Route",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := r.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_Route) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := r.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route) GetResourceNameAttribute(nameAttr *string) *string {
	if err := r.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

