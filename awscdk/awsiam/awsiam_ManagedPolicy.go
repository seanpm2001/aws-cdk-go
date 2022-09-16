package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Managed policy.
//
// Example:
//   myRole := iam.NewRole(this, jsii.String("My Role"), &roleProps{
//   	assumedBy: iam.NewServicePrincipal(jsii.String("sns.amazonaws.com")),
//   })
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_16_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	role: myRole,
//   })
//
//   myRole.addManagedPolicy(iam.managedPolicy.fromAwsManagedPolicyName(jsii.String("service-role/AWSLambdaBasicExecutionRole")))
//   myRole.addManagedPolicy(iam.managedPolicy.fromAwsManagedPolicyName(jsii.String("service-role/AWSLambdaVPCAccessExecutionRole")))
//
// Experimental.
type ManagedPolicy interface {
	awscdk.Resource
	IManagedPolicy
	// The description of this policy.
	// Experimental.
	Description() *string
	// The policy document.
	// Experimental.
	Document() PolicyDocument
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
	// Returns the ARN of this managed policy.
	// Experimental.
	ManagedPolicyArn() *string
	// The name of this policy.
	// Experimental.
	ManagedPolicyName() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The path of this policy.
	// Experimental.
	Path() *string
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
	// Adds a statement to the policy document.
	// Experimental.
	AddStatements(statement ...PolicyStatement)
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
	// Attaches this policy to a group.
	// Experimental.
	AttachToGroup(group IGroup)
	// Attaches this policy to a role.
	// Experimental.
	AttachToRole(role IRole)
	// Attaches this policy to a user.
	// Experimental.
	AttachToUser(user IUser)
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
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for ManagedPolicy
type jsiiProxy_ManagedPolicy struct {
	internal.Type__awscdkResource
	jsiiProxy_IManagedPolicy
}

func (j *jsiiProxy_ManagedPolicy) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedPolicy) Document() PolicyDocument {
	var returns PolicyDocument
	_jsii_.Get(
		j,
		"document",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedPolicy) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedPolicy) ManagedPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPolicyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedPolicy) ManagedPolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPolicyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedPolicy) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedPolicy) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedPolicy) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewManagedPolicy(scope constructs.Construct, id *string, props *ManagedPolicyProps) ManagedPolicy {
	_init_.Initialize()

	if err := validateNewManagedPolicyParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedPolicy{}

	_jsii_.Create(
		"monocdk.aws_iam.ManagedPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewManagedPolicy_Override(m ManagedPolicy, scope constructs.Construct, id *string, props *ManagedPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iam.ManagedPolicy",
		[]interface{}{scope, id, props},
		m,
	)
}

// Import a managed policy from one of the policies that AWS manages.
//
// For this managed policy, you only need to know the name to be able to use it.
//
// Some managed policy names start with "service-role/", some start with
// "job-function/", and some don't start with anything. Include the
// prefix when constructing this object.
// Experimental.
func ManagedPolicy_FromAwsManagedPolicyName(managedPolicyName *string) IManagedPolicy {
	_init_.Initialize()

	if err := validateManagedPolicy_FromAwsManagedPolicyNameParameters(managedPolicyName); err != nil {
		panic(err)
	}
	var returns IManagedPolicy

	_jsii_.StaticInvoke(
		"monocdk.aws_iam.ManagedPolicy",
		"fromAwsManagedPolicyName",
		[]interface{}{managedPolicyName},
		&returns,
	)

	return returns
}

// Import an external managed policy by ARN.
//
// For this managed policy, you only need to know the ARN to be able to use it.
// This can be useful if you got the ARN from a CloudFormation Export.
//
// If the imported Managed Policy ARN is a Token (such as a
// `CfnParameter.valueAsString` or a `Fn.importValue()`) *and* the referenced
// managed policy has a `path` (like `arn:...:policy/AdminPolicy/AdminAllow`), the
// `managedPolicyName` property will not resolve to the correct value. Instead it
// will resolve to the first path component. We unfortunately cannot express
// the correct calculation of the full path name as a CloudFormation
// expression. In this scenario the Managed Policy ARN should be supplied without the
// `path` in order to resolve the correct managed policy resource.
// Experimental.
func ManagedPolicy_FromManagedPolicyArn(scope constructs.Construct, id *string, managedPolicyArn *string) IManagedPolicy {
	_init_.Initialize()

	if err := validateManagedPolicy_FromManagedPolicyArnParameters(scope, id, managedPolicyArn); err != nil {
		panic(err)
	}
	var returns IManagedPolicy

	_jsii_.StaticInvoke(
		"monocdk.aws_iam.ManagedPolicy",
		"fromManagedPolicyArn",
		[]interface{}{scope, id, managedPolicyArn},
		&returns,
	)

	return returns
}

// Import a customer managed policy from the managedPolicyName.
//
// For this managed policy, you only need to know the name to be able to use it.
// Experimental.
func ManagedPolicy_FromManagedPolicyName(scope constructs.Construct, id *string, managedPolicyName *string) IManagedPolicy {
	_init_.Initialize()

	if err := validateManagedPolicy_FromManagedPolicyNameParameters(scope, id, managedPolicyName); err != nil {
		panic(err)
	}
	var returns IManagedPolicy

	_jsii_.StaticInvoke(
		"monocdk.aws_iam.ManagedPolicy",
		"fromManagedPolicyName",
		[]interface{}{scope, id, managedPolicyName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func ManagedPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagedPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iam.ManagedPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func ManagedPolicy_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateManagedPolicy_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_iam.ManagedPolicy",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedPolicy) AddStatements(statement ...PolicyStatement) {
	args := []interface{}{}
	for _, a := range statement {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		m,
		"addStatements",
		args,
	)
}

func (m *jsiiProxy_ManagedPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := m.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (m *jsiiProxy_ManagedPolicy) AttachToGroup(group IGroup) {
	if err := m.validateAttachToGroupParameters(group); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"attachToGroup",
		[]interface{}{group},
	)
}

func (m *jsiiProxy_ManagedPolicy) AttachToRole(role IRole) {
	if err := m.validateAttachToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"attachToRole",
		[]interface{}{role},
	)
}

func (m *jsiiProxy_ManagedPolicy) AttachToUser(user IUser) {
	if err := m.validateAttachToUserParameters(user); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"attachToUser",
		[]interface{}{user},
	)
}

func (m *jsiiProxy_ManagedPolicy) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedPolicy) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := m.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedPolicy) GetResourceNameAttribute(nameAttr *string) *string {
	if err := m.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedPolicy) OnPrepare() {
	_jsii_.InvokeVoid(
		m,
		"onPrepare",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedPolicy) OnSynthesize(session constructs.ISynthesisSession) {
	if err := m.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (m *jsiiProxy_ManagedPolicy) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedPolicy) Prepare() {
	_jsii_.InvokeVoid(
		m,
		"prepare",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedPolicy) Synthesize(session awscdk.ISynthesisSession) {
	if err := m.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"synthesize",
		[]interface{}{session},
	)
}

func (m *jsiiProxy_ManagedPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedPolicy) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}
