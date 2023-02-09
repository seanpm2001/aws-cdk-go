package awscloudformation

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Deprecated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var customResourceProvider customResourceProvider
//   var properties interface{}
//
//   customResource := awscdk.Aws_cloudformation.NewCustomResource(this, jsii.String("MyCustomResource"), &customResourceProps{
//   	provider: customResourceProvider,
//
//   	// the properties below are optional
//   	properties: map[string]interface{}{
//   		"propertiesKey": properties,
//   	},
//   	removalPolicy: monocdk.removalPolicy_DESTROY,
//   	resourceType: jsii.String("resourceType"),
//   })
//
// Deprecated: use `core.CustomResource`
type CustomResource interface {
	awscdk.CustomResource
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Deprecated: use `core.CustomResource`
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Deprecated: use `core.CustomResource`
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Deprecated: use `core.CustomResource`
	PhysicalName() *string
	// The physical name of this custom resource.
	// Deprecated: use `core.CustomResource`
	Ref() *string
	// The stack in which this resource is defined.
	// Deprecated: use `core.CustomResource`
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
	// Deprecated: use `core.CustomResource`
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Deprecated: use `core.CustomResource`
	GeneratePhysicalName() *string
	// Returns the value of an attribute of the custom resource of an arbitrary type.
	//
	// Attributes are returned from the custom resource provider through the
	// `Data` map where the key is the attribute name.
	//
	// Returns: a token for `Fn::GetAtt`. Use `Token.asXxx` to encode the returned `Reference` as a specific type or
	// use the convenience `getAttString` for string attributes.
	// Deprecated: use `core.CustomResource`
	GetAtt(attributeName *string) awscdk.Reference
	// Returns the value of an attribute of the custom resource of type string.
	//
	// Attributes are returned from the custom resource provider through the
	// `Data` map where the key is the attribute name.
	//
	// Returns: a token for `Fn::GetAtt` encoded as a string.
	// Deprecated: use `core.CustomResource`
	GetAttString(attributeName *string) *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Deprecated: use `core.CustomResource`
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Deprecated: use `core.CustomResource`
	GetResourceNameAttribute(nameAttr *string) *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: use `core.CustomResource`
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: use `core.CustomResource`
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Deprecated: use `core.CustomResource`
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: use `core.CustomResource`
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: use `core.CustomResource`
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Deprecated: use `core.CustomResource`
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Deprecated: use `core.CustomResource`
	Validate() *[]*string
}

// The jsii proxy struct for CustomResource
type jsiiProxy_CustomResource struct {
	internal.Type__awscdkCustomResource
}

func (j *jsiiProxy_CustomResource) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomResource) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomResource) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomResource) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomResource) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Deprecated: use `core.CustomResource`
func NewCustomResource(scope awscdk.Construct, id *string, props *CustomResourceProps) CustomResource {
	_init_.Initialize()

	if err := validateNewCustomResourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomResource{}

	_jsii_.Create(
		"monocdk.aws_cloudformation.CustomResource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Deprecated: use `core.CustomResource`
func NewCustomResource_Override(c CustomResource, scope awscdk.Construct, id *string, props *CustomResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudformation.CustomResource",
		[]interface{}{scope, id, props},
		c,
	)
}

// Return whether the given object is a Construct.
// Deprecated: use `core.CustomResource`
func CustomResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCustomResource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudformation.CustomResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Deprecated: use `core.CustomResource`
func CustomResource_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCustomResource_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudformation.CustomResource",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomResource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := c.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_CustomResource) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomResource) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CustomResource) GetAttString(attributeName *string) *string {
	if err := c.validateGetAttStringParameters(attributeName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getAttString",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomResource) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := c.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomResource) GetResourceNameAttribute(nameAttr *string) *string {
	if err := c.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomResource) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomResource) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CustomResource) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomResource) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomResource) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CustomResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomResource) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

