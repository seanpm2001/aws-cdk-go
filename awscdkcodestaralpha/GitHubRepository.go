// The CDK Construct Library for AWS::CodeStar
package awscdkcodestaralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcodestaralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcodestaralpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The GitHubRepository resource.
//
// Example:
//   import codestar "github.com/aws/aws-cdk-go/awscdkcodestaralpha"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   codestar.NewGitHubRepository(this, jsii.String("GitHubRepo"), &GitHubRepositoryProps{
//   	Owner: jsii.String("aws"),
//   	RepositoryName: jsii.String("aws-cdk"),
//   	AccessToken: awscdk.SecretValue_SecretsManager(jsii.String("my-github-token"), &SecretsManagerSecretOptions{
//   		JsonField: jsii.String("token"),
//   	}),
//   	ContentsBucket: s3.Bucket_FromBucketName(this, jsii.String("Bucket"), jsii.String("bucket-name")),
//   	ContentsKey: jsii.String("import.zip"),
//   })
//
// Experimental.
type GitHubRepository interface {
	awscdk.Resource
	IGitHubRepository
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
	// the repository owner.
	// Experimental.
	Owner() *string
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// the repository name.
	// Experimental.
	Repo() *string
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
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GitHubRepository
type jsiiProxy_GitHubRepository struct {
	internal.Type__awscdkResource
	jsiiProxy_IGitHubRepository
}

func (j *jsiiProxy_GitHubRepository) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubRepository) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubRepository) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubRepository) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubRepository) Repo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubRepository) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewGitHubRepository(scope constructs.Construct, id *string, props *GitHubRepositoryProps) GitHubRepository {
	_init_.Initialize()

	if err := validateNewGitHubRepositoryParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GitHubRepository{}

	_jsii_.Create(
		"@aws-cdk/aws-codestar-alpha.GitHubRepository",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewGitHubRepository_Override(g GitHubRepository, scope constructs.Construct, id *string, props *GitHubRepositoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-codestar-alpha.GitHubRepository",
		[]interface{}{scope, id, props},
		g,
	)
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
func GitHubRepository_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGitHubRepository_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-codestar-alpha.GitHubRepository",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func GitHubRepository_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateGitHubRepository_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-codestar-alpha.GitHubRepository",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func GitHubRepository_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateGitHubRepository_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-codestar-alpha.GitHubRepository",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubRepository) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := g.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (g *jsiiProxy_GitHubRepository) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubRepository) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := g.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubRepository) GetResourceNameAttribute(nameAttr *string) *string {
	if err := g.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GitHubRepository) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

