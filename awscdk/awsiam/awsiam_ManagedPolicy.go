package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Managed policy.
//
// Example:
//   myRole := iam.NewRole(this, jsii.String("My Role"), &roleProps{
//   	assumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
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
type ManagedPolicy interface {
	awscdk.Resource
	IManagedPolicy
	// The description of this policy.
	Description() *string
	// The policy document.
	Document() PolicyDocument
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// Returns the ARN of this managed policy.
	ManagedPolicyArn() *string
	// The name of this policy.
	ManagedPolicyName() *string
	// The tree node.
	Node() constructs.Node
	// The path of this policy.
	Path() *string
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
	// Adds a statement to the policy document.
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
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Attaches this policy to a group.
	AttachToGroup(group IGroup)
	// Attaches this policy to a role.
	AttachToRole(role IRole)
	// Attaches this policy to a user.
	AttachToUser(user IUser)
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

func (j *jsiiProxy_ManagedPolicy) Node() constructs.Node {
	var returns constructs.Node
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


func NewManagedPolicy(scope constructs.Construct, id *string, props *ManagedPolicyProps) ManagedPolicy {
	_init_.Initialize()

	if err := validateNewManagedPolicyParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.ManagedPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewManagedPolicy_Override(m ManagedPolicy, scope constructs.Construct, id *string, props *ManagedPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.ManagedPolicy",
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
func ManagedPolicy_FromAwsManagedPolicyName(managedPolicyName *string) IManagedPolicy {
	_init_.Initialize()

	if err := validateManagedPolicy_FromAwsManagedPolicyNameParameters(managedPolicyName); err != nil {
		panic(err)
	}
	var returns IManagedPolicy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.ManagedPolicy",
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
func ManagedPolicy_FromManagedPolicyArn(scope constructs.Construct, id *string, managedPolicyArn *string) IManagedPolicy {
	_init_.Initialize()

	if err := validateManagedPolicy_FromManagedPolicyArnParameters(scope, id, managedPolicyArn); err != nil {
		panic(err)
	}
	var returns IManagedPolicy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.ManagedPolicy",
		"fromManagedPolicyArn",
		[]interface{}{scope, id, managedPolicyArn},
		&returns,
	)

	return returns
}

// Import a customer managed policy from the managedPolicyName.
//
// For this managed policy, you only need to know the name to be able to use it.
func ManagedPolicy_FromManagedPolicyName(scope constructs.Construct, id *string, managedPolicyName *string) IManagedPolicy {
	_init_.Initialize()

	if err := validateManagedPolicy_FromManagedPolicyNameParameters(scope, id, managedPolicyName); err != nil {
		panic(err)
	}
	var returns IManagedPolicy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.ManagedPolicy",
		"fromManagedPolicyName",
		[]interface{}{scope, id, managedPolicyName},
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
func ManagedPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagedPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.ManagedPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func ManagedPolicy_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateManagedPolicy_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.ManagedPolicy",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ManagedPolicy_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateManagedPolicy_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.ManagedPolicy",
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

