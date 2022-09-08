package awsservicediscovery

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsservicediscovery/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Define a Public DNS Namespace.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import servicediscovery "github.com/aws/aws-cdk-go/awscdk"
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("aws-servicediscovery-integ"))
//
//   namespace := servicediscovery.NewPublicDnsNamespace(stack, jsii.String("Namespace"), &publicDnsNamespaceProps{
//   	name: jsii.String("foobar.com"),
//   })
//
//   service := namespace.createService(jsii.String("Service"), &dnsServiceProps{
//   	name: jsii.String("foo"),
//   	dnsRecordType: servicediscovery.dnsRecordType_A,
//   	dnsTtl: cdk.duration.seconds(jsii.Number(30)),
//   	healthCheck: &healthCheckConfig{
//   		type: servicediscovery.healthCheckType_HTTPS,
//   		resourcePath: jsii.String("/healthcheck"),
//   		failureThreshold: jsii.Number(2),
//   	},
//   })
//
//   service.registerIpInstance(jsii.String("IpInstance"), &ipInstanceBaseProps{
//   	ipv4: jsii.String("54.239.25.192"),
//   	port: jsii.Number(443),
//   })
//
//   app.synth()
//
// Experimental.
type PublicDnsNamespace interface {
	awscdk.Resource
	IPublicDnsNamespace
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
	// Namespace Arn for the namespace.
	// Experimental.
	NamespaceArn() *string
	// Namespace Id for the namespace.
	// Experimental.
	NamespaceId() *string
	// A name for the namespace.
	// Experimental.
	NamespaceName() *string
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
	// Experimental.
	PublicDnsNamespaceArn() *string
	// Experimental.
	PublicDnsNamespaceId() *string
	// Experimental.
	PublicDnsNamespaceName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Type of the namespace.
	// Experimental.
	Type() NamespaceType
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
	// Creates a service within the namespace.
	// Experimental.
	CreateService(id *string, props *DnsServiceProps) Service
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

// The jsii proxy struct for PublicDnsNamespace
type jsiiProxy_PublicDnsNamespace struct {
	internal.Type__awscdkResource
	jsiiProxy_IPublicDnsNamespace
}

func (j *jsiiProxy_PublicDnsNamespace) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) NamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) NamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) NamespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) PublicDnsNamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicDnsNamespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) PublicDnsNamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicDnsNamespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) PublicDnsNamespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicDnsNamespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublicDnsNamespace) Type() NamespaceType {
	var returns NamespaceType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Experimental.
func NewPublicDnsNamespace(scope constructs.Construct, id *string, props *PublicDnsNamespaceProps) PublicDnsNamespace {
	_init_.Initialize()

	if err := validateNewPublicDnsNamespaceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PublicDnsNamespace{}

	_jsii_.Create(
		"monocdk.aws_servicediscovery.PublicDnsNamespace",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewPublicDnsNamespace_Override(p PublicDnsNamespace, scope constructs.Construct, id *string, props *PublicDnsNamespaceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_servicediscovery.PublicDnsNamespace",
		[]interface{}{scope, id, props},
		p,
	)
}

// Experimental.
func PublicDnsNamespace_FromPublicDnsNamespaceAttributes(scope constructs.Construct, id *string, attrs *PublicDnsNamespaceAttributes) IPublicDnsNamespace {
	_init_.Initialize()

	if err := validatePublicDnsNamespace_FromPublicDnsNamespaceAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IPublicDnsNamespace

	_jsii_.StaticInvoke(
		"monocdk.aws_servicediscovery.PublicDnsNamespace",
		"fromPublicDnsNamespaceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func PublicDnsNamespace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePublicDnsNamespace_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_servicediscovery.PublicDnsNamespace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func PublicDnsNamespace_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validatePublicDnsNamespace_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_servicediscovery.PublicDnsNamespace",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicDnsNamespace) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := p.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (p *jsiiProxy_PublicDnsNamespace) CreateService(id *string, props *DnsServiceProps) Service {
	if err := p.validateCreateServiceParameters(id, props); err != nil {
		panic(err)
	}
	var returns Service

	_jsii_.Invoke(
		p,
		"createService",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicDnsNamespace) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicDnsNamespace) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := p.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicDnsNamespace) GetResourceNameAttribute(nameAttr *string) *string {
	if err := p.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicDnsNamespace) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PublicDnsNamespace) OnSynthesize(session constructs.ISynthesisSession) {
	if err := p.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (p *jsiiProxy_PublicDnsNamespace) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicDnsNamespace) Prepare() {
	_jsii_.InvokeVoid(
		p,
		"prepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PublicDnsNamespace) Synthesize(session awscdk.ISynthesisSession) {
	if err := p.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		[]interface{}{session},
	)
}

func (p *jsiiProxy_PublicDnsNamespace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PublicDnsNamespace) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

