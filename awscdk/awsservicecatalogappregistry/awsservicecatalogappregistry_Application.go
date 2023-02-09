package awsservicecatalogappregistry

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsservicecatalogappregistry/internal"
	"github.com/aws/constructs-go/constructs/v3"
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
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
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
	// Associate an attribute group with application If the attribute group is already associated, it will ignore duplicate request.
	// Experimental.
	AssociateAttributeGroup(attributeGroup IAttributeGroup)
	// Associate a stack with the application If the resource is already associated, it will ignore duplicate request.
	//
	// A stack can only be associated with one application.
	// Experimental.
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
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
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

func (j *jsiiProxy_Application) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
		"monocdk.aws_servicecatalogappregistry.Application",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewApplication_Override(a Application, scope constructs.Construct, id *string, props *ApplicationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_servicecatalogappregistry.Application",
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
		"monocdk.aws_servicecatalogappregistry.Application",
		"fromApplicationArn",
		[]interface{}{scope, id, applicationArn},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Application_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplication_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_servicecatalogappregistry.Application",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Application_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateApplication_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_servicecatalogappregistry.Application",
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

func (a *jsiiProxy_Application) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) OnSynthesize(session constructs.ISynthesisSession) {
	if err := a.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (a *jsiiProxy_Application) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Application) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_Application) Synthesize(session awscdk.ISynthesisSession) {
	if err := a.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
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

func (a *jsiiProxy_Application) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

