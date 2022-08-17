// The CDK Construct Library for AWS::Cloud9
package awscdkcloud9alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcloud9alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodecommit"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdkcloud9alpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The class for different repository providers.
//
// Example:
//   import codecommit "github.com/aws/aws-cdk-go/awscdk"
//
//   // create a new Cloud9 environment and clone the two repositories
//   var vpc vpc
//
//
//   // create a codecommit repository to clone into the cloud9 environment
//   repoNew := codecommit.NewRepository(this, jsii.String("RepoNew"), &repositoryProps{
//   	repositoryName: jsii.String("new-repo"),
//   })
//
//   // import an existing codecommit repository to clone into the cloud9 environment
//   repoExisting := codecommit.repository.fromRepositoryName(this, jsii.String("RepoExisting"), jsii.String("existing-repo"))
//   cloud9.NewEc2Environment(this, jsii.String("C9Env"), &ec2EnvironmentProps{
//   	vpc: vpc,
//   	clonedRepositories: []cloneRepository{
//   		cloud9.*cloneRepository.fromCodeCommit(repoNew, jsii.String("/src/new-repo")),
//   		cloud9.*cloneRepository.fromCodeCommit(repoExisting, jsii.String("/src/existing-repo")),
//   	},
//   	imageId: cloud9.imageId_AMAZON_LINUX_2,
//   })
//
// Experimental.
type CloneRepository interface {
	// Experimental.
	PathComponent() *string
	// Experimental.
	RepositoryUrl() *string
}

// The jsii proxy struct for CloneRepository
type jsiiProxy_CloneRepository struct {
	_ byte // padding
}

func (j *jsiiProxy_CloneRepository) PathComponent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathComponent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloneRepository) RepositoryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUrl",
		&returns,
	)
	return returns
}


// import repository to cloud9 environment from AWS CodeCommit.
// Experimental.
func CloneRepository_FromCodeCommit(repository awscodecommit.IRepository, path *string) CloneRepository {
	_init_.Initialize()

	var returns CloneRepository

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-cloud9-alpha.CloneRepository",
		"fromCodeCommit",
		[]interface{}{repository, path},
		&returns,
	)

	return returns
}

// The connection type used for connecting to an Amazon EC2 environment.
// Experimental.
type ConnectionType string

const (
	// Connect through SSH.
	// Experimental.
	ConnectionType_CONNECT_SSH ConnectionType = "CONNECT_SSH"
	// Connect through AWS Systems Manager When using SSM, service role and instance profile aren't automatically created.
	//
	// See https://docs.aws.amazon.com/cloud9/latest/user-guide/ec2-ssm.html#service-role-ssm
	// Experimental.
	ConnectionType_CONNECT_SSM ConnectionType = "CONNECT_SSM"
)

// A Cloud9 Environment with Amazon EC2.
//
// Example:
//   import codecommit "github.com/aws/aws-cdk-go/awscdk"
//
//   // create a new Cloud9 environment and clone the two repositories
//   var vpc vpc
//
//
//   // create a codecommit repository to clone into the cloud9 environment
//   repoNew := codecommit.NewRepository(this, jsii.String("RepoNew"), &repositoryProps{
//   	repositoryName: jsii.String("new-repo"),
//   })
//
//   // import an existing codecommit repository to clone into the cloud9 environment
//   repoExisting := codecommit.repository.fromRepositoryName(this, jsii.String("RepoExisting"), jsii.String("existing-repo"))
//   cloud9.NewEc2Environment(this, jsii.String("C9Env"), &ec2EnvironmentProps{
//   	vpc: vpc,
//   	clonedRepositories: []cloneRepository{
//   		cloud9.*cloneRepository.fromCodeCommit(repoNew, jsii.String("/src/new-repo")),
//   		cloud9.*cloneRepository.fromCodeCommit(repoExisting, jsii.String("/src/existing-repo")),
//   	},
//   	imageId: cloud9.imageId_AMAZON_LINUX_2,
//   })
//
// Experimental.
type Ec2Environment interface {
	awscdk.Resource
	IEc2Environment
	// The environment ARN of this Cloud9 environment.
	// Experimental.
	Ec2EnvironmentArn() *string
	// The environment name of this Cloud9 environment.
	// Experimental.
	Ec2EnvironmentName() *string
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
	// The environment ID of this Cloud9 environment.
	// Experimental.
	EnvironmentId() *string
	// The complete IDE URL of this Cloud9 environment.
	// Experimental.
	IdeUrl() *string
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
	// VPC ID.
	// Experimental.
	Vpc() awsec2.IVpc
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

// The jsii proxy struct for Ec2Environment
type jsiiProxy_Ec2Environment struct {
	internal.Type__awscdkResource
	jsiiProxy_IEc2Environment
}

func (j *jsiiProxy_Ec2Environment) Ec2EnvironmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2EnvironmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) Ec2EnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2EnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) EnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) IdeUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ideUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Experimental.
func NewEc2Environment(scope constructs.Construct, id *string, props *Ec2EnvironmentProps) Ec2Environment {
	_init_.Initialize()

	j := jsiiProxy_Ec2Environment{}

	_jsii_.Create(
		"@aws-cdk/aws-cloud9-alpha.Ec2Environment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEc2Environment_Override(e Ec2Environment, scope constructs.Construct, id *string, props *Ec2EnvironmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-cloud9-alpha.Ec2Environment",
		[]interface{}{scope, id, props},
		e,
	)
}

// import from EnvironmentEc2Name.
// Experimental.
func Ec2Environment_FromEc2EnvironmentName(scope constructs.Construct, id *string, ec2EnvironmentName *string) IEc2Environment {
	_init_.Initialize()

	var returns IEc2Environment

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-cloud9-alpha.Ec2Environment",
		"fromEc2EnvironmentName",
		[]interface{}{scope, id, ec2EnvironmentName},
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
func Ec2Environment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-cloud9-alpha.Ec2Environment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func Ec2Environment_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-cloud9-alpha.Ec2Environment",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Ec2Environment_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-cloud9-alpha.Ec2Environment",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Environment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (e *jsiiProxy_Ec2Environment) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Environment) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Environment) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Environment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for Ec2Environment.
//
// Example:
//   import codecommit "github.com/aws/aws-cdk-go/awscdk"
//
//   // create a new Cloud9 environment and clone the two repositories
//   var vpc vpc
//
//
//   // create a codecommit repository to clone into the cloud9 environment
//   repoNew := codecommit.NewRepository(this, jsii.String("RepoNew"), &repositoryProps{
//   	repositoryName: jsii.String("new-repo"),
//   })
//
//   // import an existing codecommit repository to clone into the cloud9 environment
//   repoExisting := codecommit.repository.fromRepositoryName(this, jsii.String("RepoExisting"), jsii.String("existing-repo"))
//   cloud9.NewEc2Environment(this, jsii.String("C9Env"), &ec2EnvironmentProps{
//   	vpc: vpc,
//   	clonedRepositories: []cloneRepository{
//   		cloud9.*cloneRepository.fromCodeCommit(repoNew, jsii.String("/src/new-repo")),
//   		cloud9.*cloneRepository.fromCodeCommit(repoExisting, jsii.String("/src/existing-repo")),
//   	},
//   	imageId: cloud9.imageId_AMAZON_LINUX_2,
//   })
//
// Experimental.
type Ec2EnvironmentProps struct {
	// The image ID used for creating an Amazon EC2 environment.
	// Experimental.
	ImageId ImageId `field:"required" json:"imageId" yaml:"imageId"`
	// The VPC that AWS Cloud9 will use to communicate with the Amazon Elastic Compute Cloud (Amazon EC2) instance.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// The AWS CodeCommit repository to be cloned.
	// Experimental.
	ClonedRepositories *[]CloneRepository `field:"optional" json:"clonedRepositories" yaml:"clonedRepositories"`
	// The connection type used for connecting to an Amazon EC2 environment.
	//
	// Valid values are: CONNECT_SSH (default) and CONNECT_SSM (connected through AWS Systems Manager).
	// Experimental.
	ConnectionType ConnectionType `field:"optional" json:"connectionType" yaml:"connectionType"`
	// Description of the environment.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Name of the environment.
	// Experimental.
	Ec2EnvironmentName *string `field:"optional" json:"ec2EnvironmentName" yaml:"ec2EnvironmentName"`
	// The type of instance to connect to the environment.
	// Experimental.
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The subnetSelection of the VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
}

// A Cloud9 Environment.
// Experimental.
type IEc2Environment interface {
	awscdk.IResource
	// The arn of the EnvironmentEc2.
	// Experimental.
	Ec2EnvironmentArn() *string
	// The name of the EnvironmentEc2.
	// Experimental.
	Ec2EnvironmentName() *string
}

// The jsii proxy for IEc2Environment
type jsiiProxy_IEc2Environment struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IEc2Environment) Ec2EnvironmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2EnvironmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEc2Environment) Ec2EnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2EnvironmentName",
		&returns,
	)
	return returns
}

// The image ID used for creating an Amazon EC2 environment.
//
// Example:
//   import codecommit "github.com/aws/aws-cdk-go/awscdk"
//
//   // create a new Cloud9 environment and clone the two repositories
//   var vpc vpc
//
//
//   // create a codecommit repository to clone into the cloud9 environment
//   repoNew := codecommit.NewRepository(this, jsii.String("RepoNew"), &repositoryProps{
//   	repositoryName: jsii.String("new-repo"),
//   })
//
//   // import an existing codecommit repository to clone into the cloud9 environment
//   repoExisting := codecommit.repository.fromRepositoryName(this, jsii.String("RepoExisting"), jsii.String("existing-repo"))
//   cloud9.NewEc2Environment(this, jsii.String("C9Env"), &ec2EnvironmentProps{
//   	vpc: vpc,
//   	clonedRepositories: []cloneRepository{
//   		cloud9.*cloneRepository.fromCodeCommit(repoNew, jsii.String("/src/new-repo")),
//   		cloud9.*cloneRepository.fromCodeCommit(repoExisting, jsii.String("/src/existing-repo")),
//   	},
//   	imageId: cloud9.imageId_AMAZON_LINUX_2,
//   })
//
// Experimental.
type ImageId string

const (
	// Create using Amazon Linux 2.
	// Experimental.
	ImageId_AMAZON_LINUX_2 ImageId = "AMAZON_LINUX_2"
	// Create using Ubunut 18.04.
	// Experimental.
	ImageId_UBUNTU_18_04 ImageId = "UBUNTU_18_04"
)
