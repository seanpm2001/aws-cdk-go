package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awsrds/internal"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v3"
)

// Create an Aurora Serverless Cluster.
//
// Example:
//   var vpc vpc
//
//   var code code
//
//
//   cluster := rds.NewServerlessCluster(this, jsii.String("AnotherCluster"), &serverlessClusterProps{
//   	engine: rds.databaseClusterEngine_AURORA_MYSQL(),
//   	vpc: vpc,
//   	 // this parameter is optional for serverless Clusters
//   	enableDataApi: jsii.Boolean(true),
//   })
//   fn := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	handler: jsii.String("index.handler"),
//   	code: code,
//   	environment: map[string]*string{
//   		"CLUSTER_ARN": cluster.clusterArn,
//   		"SECRET_ARN": cluster.secret.secretArn,
//   	},
//   })
//   cluster.grantDataApiAccess(fn)
//
// Experimental.
type ServerlessCluster interface {
	awscdk.Resource
	IServerlessCluster
	// The ARN of the cluster.
	// Experimental.
	ClusterArn() *string
	// The endpoint to use for read/write operations.
	// Experimental.
	ClusterEndpoint() Endpoint
	// Identifier of the cluster.
	// Experimental.
	ClusterIdentifier() *string
	// The endpoint to use for read/write operations.
	// Experimental.
	ClusterReadEndpoint() Endpoint
	// Access to the network connections.
	// Experimental.
	Connections() awsec2.Connections
	// Experimental.
	EnableDataApi() *bool
	// Experimental.
	SetEnableDataApi(val *bool)
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
	// Experimental.
	NewCfnProps() *CfnDBClusterProps
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
	// The secret attached to this cluster.
	// Experimental.
	Secret() awssecretsmanager.ISecret
	// Experimental.
	SecurityGroups() *[]awsec2.ISecurityGroup
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Adds the multi user rotation to this cluster.
	// Experimental.
	AddRotationMultiUser(id *string, options *RotationMultiUserOptions) awssecretsmanager.SecretRotation
	// Adds the single user rotation of the master password to this cluster.
	// Experimental.
	AddRotationSingleUser(options *RotationSingleUserOptions) awssecretsmanager.SecretRotation
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
	// Renders the secret attachment target specifications.
	// Experimental.
	AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps
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
	// Grant the given identity to access to the Data API, including read access to the secret attached to the cluster if present.
	// Experimental.
	GrantDataApiAccess(grantee awsiam.IGrantable) awsiam.Grant
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

// The jsii proxy struct for ServerlessCluster
type jsiiProxy_ServerlessCluster struct {
	internal.Type__awscdkResource
	jsiiProxy_IServerlessCluster
}

func (j *jsiiProxy_ServerlessCluster) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) ClusterEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) ClusterReadEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterReadEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) EnableDataApi() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableDataApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) NewCfnProps() *CfnDBClusterProps {
	var returns *CfnDBClusterProps
	_jsii_.Get(
		j,
		"newCfnProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) Secret() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) SecurityGroups() *[]awsec2.ISecurityGroup {
	var returns *[]awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewServerlessCluster(scope constructs.Construct, id *string, props *ServerlessClusterProps) ServerlessCluster {
	_init_.Initialize()

	if err := validateNewServerlessClusterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServerlessCluster{}

	_jsii_.Create(
		"monocdk.aws_rds.ServerlessCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewServerlessCluster_Override(s ServerlessCluster, scope constructs.Construct, id *string, props *ServerlessClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_rds.ServerlessCluster",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_ServerlessCluster)SetEnableDataApi(val *bool) {
	_jsii_.Set(
		j,
		"enableDataApi",
		val,
	)
}

// Import an existing DatabaseCluster from properties.
// Experimental.
func ServerlessCluster_FromServerlessClusterAttributes(scope constructs.Construct, id *string, attrs *ServerlessClusterAttributes) IServerlessCluster {
	_init_.Initialize()

	if err := validateServerlessCluster_FromServerlessClusterAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IServerlessCluster

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.ServerlessCluster",
		"fromServerlessClusterAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func ServerlessCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServerlessCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.ServerlessCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func ServerlessCluster_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateServerlessCluster_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.ServerlessCluster",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessCluster) AddRotationMultiUser(id *string, options *RotationMultiUserOptions) awssecretsmanager.SecretRotation {
	if err := s.validateAddRotationMultiUserParameters(id, options); err != nil {
		panic(err)
	}
	var returns awssecretsmanager.SecretRotation

	_jsii_.Invoke(
		s,
		"addRotationMultiUser",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessCluster) AddRotationSingleUser(options *RotationSingleUserOptions) awssecretsmanager.SecretRotation {
	if err := s.validateAddRotationSingleUserParameters(options); err != nil {
		panic(err)
	}
	var returns awssecretsmanager.SecretRotation

	_jsii_.Invoke(
		s,
		"addRotationSingleUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := s.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_ServerlessCluster) AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps {
	var returns *awssecretsmanager.SecretAttachmentTargetProps

	_jsii_.Invoke(
		s,
		"asSecretAttachmentTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessCluster) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessCluster) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := s.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessCluster) GetResourceNameAttribute(nameAttr *string) *string {
	if err := s.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessCluster) GrantDataApiAccess(grantee awsiam.IGrantable) awsiam.Grant {
	if err := s.validateGrantDataApiAccessParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantDataApiAccess",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessCluster) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerlessCluster) OnSynthesize(session constructs.ISynthesisSession) {
	if err := s.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_ServerlessCluster) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessCluster) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerlessCluster) Synthesize(session awscdk.ISynthesisSession) {
	if err := s.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_ServerlessCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessCluster) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

