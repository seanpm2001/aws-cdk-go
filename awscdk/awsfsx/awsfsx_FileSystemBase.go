package awsfsx

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsfsx/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A new or imported FSx file system.
// Experimental.
type FileSystemBase interface {
	awscdk.Resource
	IFileSystem
	// The security groups/rules used to allow network connections to the file system.
	// Experimental.
	Connections() awsec2.Connections
	// The DNS name assigned to this file system.
	// Experimental.
	DnsName() *string
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
	// The ID of the file system, assigned by Amazon FSx.
	// Experimental.
	FileSystemId() *string
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

// The jsii proxy struct for FileSystemBase
type jsiiProxy_FileSystemBase struct {
	internal.Type__awscdkResource
	jsiiProxy_IFileSystem
}

func (j *jsiiProxy_FileSystemBase) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystemBase) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystemBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystemBase) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystemBase) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystemBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystemBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewFileSystemBase_Override(f FileSystemBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_fsx.FileSystemBase",
		[]interface{}{scope, id, props},
		f,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func FileSystemBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFileSystemBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_fsx.FileSystemBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func FileSystemBase_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateFileSystemBase_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_fsx.FileSystemBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystemBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := f.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (f *jsiiProxy_FileSystemBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystemBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := f.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystemBase) GetResourceNameAttribute(nameAttr *string) *string {
	if err := f.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystemBase) OnPrepare() {
	_jsii_.InvokeVoid(
		f,
		"onPrepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileSystemBase) OnSynthesize(session constructs.ISynthesisSession) {
	if err := f.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (f *jsiiProxy_FileSystemBase) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystemBase) Prepare() {
	_jsii_.InvokeVoid(
		f,
		"prepare",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FileSystemBase) Synthesize(session awscdk.ISynthesisSession) {
	if err := f.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"synthesize",
		[]interface{}{session},
	)
}

func (f *jsiiProxy_FileSystemBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystemBase) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

