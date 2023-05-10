package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Example:
//   var acm interface{}
//
//
//   apigateway.NewDomainName(this, jsii.String("domain-name"), &DomainNameProps{
//   	DomainName: jsii.String("example.com"),
//   	Certificate: acm.certificate_FromCertificateArn(this, jsii.String("cert"), jsii.String("arn:aws:acm:us-east-1:1111111:certificate/11-3336f1-44483d-adc7-9cd375c5169d")),
//   	Mtls: &MTLSConfig{
//   		Bucket: s3.NewBucket(this, jsii.String("bucket")),
//   		Key: jsii.String("truststore.pem"),
//   		Version: jsii.String("version"),
//   	},
//   })
//
// Experimental.
type DomainName interface {
	awscdk.Resource
	IDomainName
	// The domain name (e.g. `example.com`).
	// Experimental.
	DomainName() *string
	// The Route53 alias target to use in order to connect a record set to this domain through an alias.
	// Experimental.
	DomainNameAliasDomainName() *string
	// The Route53 hosted zone ID to use in order to connect a record set to this domain through an alias.
	// Experimental.
	DomainNameAliasHostedZoneId() *string
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
	// Maps this domain to an API endpoint.
	// Experimental.
	AddBasePathMapping(targetApi IRestApi, options *BasePathMappingOptions) BasePathMapping
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

// The jsii proxy struct for DomainName
type jsiiProxy_DomainName struct {
	internal.Type__awscdkResource
	jsiiProxy_IDomainName
}

func (j *jsiiProxy_DomainName) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DomainName) DomainNameAliasDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameAliasDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DomainName) DomainNameAliasHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameAliasHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DomainName) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DomainName) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DomainName) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DomainName) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewDomainName(scope constructs.Construct, id *string, props *DomainNameProps) DomainName {
	_init_.Initialize()

	if err := validateNewDomainNameParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DomainName{}

	_jsii_.Create(
		"monocdk.aws_apigateway.DomainName",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDomainName_Override(d DomainName, scope constructs.Construct, id *string, props *DomainNameProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigateway.DomainName",
		[]interface{}{scope, id, props},
		d,
	)
}

// Imports an existing domain name.
// Experimental.
func DomainName_FromDomainNameAttributes(scope constructs.Construct, id *string, attrs *DomainNameAttributes) IDomainName {
	_init_.Initialize()

	if err := validateDomainName_FromDomainNameAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IDomainName

	_jsii_.StaticInvoke(
		"monocdk.aws_apigateway.DomainName",
		"fromDomainNameAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func DomainName_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDomainName_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apigateway.DomainName",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func DomainName_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateDomainName_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_apigateway.DomainName",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainName) AddBasePathMapping(targetApi IRestApi, options *BasePathMappingOptions) BasePathMapping {
	if err := d.validateAddBasePathMappingParameters(targetApi, options); err != nil {
		panic(err)
	}
	var returns BasePathMapping

	_jsii_.Invoke(
		d,
		"addBasePathMapping",
		[]interface{}{targetApi, options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainName) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := d.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (d *jsiiProxy_DomainName) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainName) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := d.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainName) GetResourceNameAttribute(nameAttr *string) *string {
	if err := d.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainName) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DomainName) OnSynthesize(session constructs.ISynthesisSession) {
	if err := d.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (d *jsiiProxy_DomainName) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainName) Prepare() {
	_jsii_.InvokeVoid(
		d,
		"prepare",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DomainName) Synthesize(session awscdk.ISynthesisSession) {
	if err := d.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		[]interface{}{session},
	)
}

func (d *jsiiProxy_DomainName) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DomainName) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

