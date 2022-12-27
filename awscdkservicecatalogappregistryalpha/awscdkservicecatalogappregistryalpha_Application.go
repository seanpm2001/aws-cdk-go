// The CDK Construct Library for AWS::ServiceCatalogAppRegistry
package awscdkservicecatalogappregistryalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkservicecatalogappregistryalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkservicecatalogappregistryalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Service Catalog AppRegistry Application.
//
// Example:
//   application := appreg.NewApplication(this, jsii.String("MyFirstApplication"), &applicationProps{
//   	applicationName: jsii.String("MyFirstApplicationName"),
//   	description: jsii.String("description for my application"),
//   })
//
// Experimental.
type Application interface {
	awscdk.Resource
	IApplication
	// The ARN of the application.
	// Experimental.
	ApplicationArn() *string
	// The ID of the application.
	// Experimental.
	ApplicationId() *string
	// Application manager URL for the Application.
	// Experimental.
	ApplicationManagerUrl() awscdk.CfnOutput
	// The name of the application.
	// Experimental.
	ApplicationName() *string
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
	//    cross-environment scenarios.
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
	// Associate all stacks present in construct's aspect with application.
	//
	// NOTE: This method won't automatically register stacks under pipeline stages,
	// and requires association of each pipeline stage by calling this method with stage Construct.
	// Experimental.
	AssociateAllStacksInScope(scope constructs.Construct)
	// Associate stack with the application in the stack passed as parameter.
	//
	// A stack can only be associated with one application.
	// Experimental.
	AssociateApplicationWithStack(stack awscdk.Stack)
	// Associate an attribute group with application If the attribute group is already associated, it will ignore duplicate request.
	// Experimental.
	AssociateAttributeGroup(attributeGroup IAttributeGroup)
	// Associate a stack with the application If the resource is already associated, it will ignore duplicate request.
	//
	// A stack can only be associated with one application.
	// Deprecated: Use `associateApplicationWithStack` instead.
	AssociateStack(stack awscdk.Stack)
	// Experimental.
	GeneratePhysicalName() *string
	// Create a unique id.
	// Experimental.
	GenerateUniqueHash(resourceAddress *string) *string
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
	// Share an application with accounts, organizations and OUs, and IAM roles and users.
	//
	// The application will become available to end users within those principals.
	// Experimental.
	ShareApplication(shareOptions *ShareOptions)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Application
type jsiiProxy_Application struct {
	internal.Type__awscdkResource
	jsiiProxy_IApplication
}

func (j *jsiiProxy_Application) ApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) ApplicationManagerUrl() awscdk.CfnOutput {
	var returns awscdk.CfnOutput
	_jsii_.Get(
		j,
		"applicationManagerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewApplication(scope constructs.Construct, id *string, props *ApplicationProps) Application {
	_init_.Initialize()

	if err := validateNewApplicationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Application{}

	_jsii_.Create(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.Application",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewApplication_Override(a Application, scope constructs.Construct, id *string, props *ApplicationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.Application",
		[]interface{}{scope, id, props},
		a,
	)
}

// Imports an Application construct that represents an external application.
// Experimental.
func Application_FromApplicationArn(scope constructs.Construct, id *string, applicationArn *string) IApplication {
	_init_.Initialize()

	if err := validateApplication_FromApplicationArnParameters(scope, id, applicationArn); err != nil {
		panic(err)
	}
	var returns IApplication

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.Application",
		"fromApplicationArn",
		[]interface{}{scope, id, applicationArn},
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
func Application_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplication_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.Application",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func Application_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateApplication_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.Application",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Application_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateApplication_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.Application",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Application) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := a.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_Application) AssociateAllStacksInScope(scope constructs.Construct) {
	if err := a.validateAssociateAllStacksInScopeParameters(scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"associateAllStacksInScope",
		[]interface{}{scope},
	)
}

func (a *jsiiProxy_Application) AssociateApplicationWithStack(stack awscdk.Stack) {
	if err := a.validateAssociateApplicationWithStackParameters(stack); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"associateApplicationWithStack",
		[]interface{}{stack},
	)
}

func (a *jsiiProxy_Application) AssociateAttributeGroup(attributeGroup IAttributeGroup) {
	if err := a.validateAssociateAttributeGroupParameters(attributeGroup); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"associateAttributeGroup",
		[]interface{}{attributeGroup},
	)
}

func (a *jsiiProxy_Application) AssociateStack(stack awscdk.Stack) {
	if err := a.validateAssociateStackParameters(stack); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"associateStack",
		[]interface{}{stack},
	)
}

func (a *jsiiProxy_Application) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Application) GenerateUniqueHash(resourceAddress *string) *string {
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

func (a *jsiiProxy_Application) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (a *jsiiProxy_Application) GetResourceNameAttribute(nameAttr *string) *string {
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

func (a *jsiiProxy_Application) ShareApplication(shareOptions *ShareOptions) {
	if err := a.validateShareApplicationParameters(shareOptions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"shareApplication",
		[]interface{}{shareOptions},
	)
}

func (a *jsiiProxy_Application) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

