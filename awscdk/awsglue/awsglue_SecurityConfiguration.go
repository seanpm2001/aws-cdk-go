package awsglue

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsglue/internal"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/constructs-go/constructs/v3"
)

// A security configuration is a set of security properties that can be used by AWS Glue to encrypt data at rest.
//
// The following scenarios show some of the ways that you can use a security configuration.
// - Attach a security configuration to an AWS Glue crawler to write encrypted Amazon CloudWatch Logs.
// - Attach a security configuration to an extract, transform, and load (ETL) job to write encrypted Amazon Simple Storage Service (Amazon S3) targets and encrypted CloudWatch Logs.
// - Attach a security configuration to an ETL job to write its jobs bookmarks as encrypted Amazon S3 data.
// - Attach a security configuration to a development endpoint to write encrypted Amazon S3 targets.
//
// Example:
//   glue.NewSecurityConfiguration(this, jsii.String("MySecurityConfiguration"), &securityConfigurationProps{
//   	securityConfigurationName: jsii.String("name"),
//   	cloudWatchEncryption: &cloudWatchEncryption{
//   		mode: glue.cloudWatchEncryptionMode_KMS,
//   	},
//   	jobBookmarksEncryption: &jobBookmarksEncryption{
//   		mode: glue.jobBookmarksEncryptionMode_CLIENT_SIDE_KMS,
//   	},
//   	s3Encryption: &s3Encryption{
//   		mode: glue.s3EncryptionMode_KMS,
//   	},
//   })
//
// Experimental.
type SecurityConfiguration interface {
	awscdk.Resource
	ISecurityConfiguration
	// The KMS key used in CloudWatch encryption if it requires a kms key.
	// Experimental.
	CloudWatchEncryptionKey() awskms.IKey
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
	// The KMS key used in job bookmarks encryption if it requires a kms key.
	// Experimental.
	JobBookmarksEncryptionKey() awskms.IKey
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
	// The KMS key used in S3 encryption if it requires a kms key.
	// Experimental.
	S3EncryptionKey() awskms.IKey
	// The name of the security configuration.
	// Experimental.
	SecurityConfigurationName() *string
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

// The jsii proxy struct for SecurityConfiguration
type jsiiProxy_SecurityConfiguration struct {
	internal.Type__awscdkResource
	jsiiProxy_ISecurityConfiguration
}

func (j *jsiiProxy_SecurityConfiguration) CloudWatchEncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"cloudWatchEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityConfiguration) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityConfiguration) JobBookmarksEncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"jobBookmarksEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityConfiguration) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityConfiguration) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityConfiguration) S3EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"s3EncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityConfiguration) SecurityConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewSecurityConfiguration(scope constructs.Construct, id *string, props *SecurityConfigurationProps) SecurityConfiguration {
	_init_.Initialize()

	if err := validateNewSecurityConfigurationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityConfiguration{}

	_jsii_.Create(
		"monocdk.aws_glue.SecurityConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSecurityConfiguration_Override(s SecurityConfiguration, scope constructs.Construct, id *string, props *SecurityConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_glue.SecurityConfiguration",
		[]interface{}{scope, id, props},
		s,
	)
}

// Creates a Connection construct that represents an external security configuration.
// Experimental.
func SecurityConfiguration_FromSecurityConfigurationName(scope constructs.Construct, id *string, securityConfigurationName *string) ISecurityConfiguration {
	_init_.Initialize()

	if err := validateSecurityConfiguration_FromSecurityConfigurationNameParameters(scope, id, securityConfigurationName); err != nil {
		panic(err)
	}
	var returns ISecurityConfiguration

	_jsii_.StaticInvoke(
		"monocdk.aws_glue.SecurityConfiguration",
		"fromSecurityConfigurationName",
		[]interface{}{scope, id, securityConfigurationName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func SecurityConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecurityConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_glue.SecurityConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func SecurityConfiguration_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateSecurityConfiguration_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_glue.SecurityConfiguration",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := s.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_SecurityConfiguration) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityConfiguration) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (s *jsiiProxy_SecurityConfiguration) GetResourceNameAttribute(nameAttr *string) *string {
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

func (s *jsiiProxy_SecurityConfiguration) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityConfiguration) OnSynthesize(session constructs.ISynthesisSession) {
	if err := s.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_SecurityConfiguration) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityConfiguration) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityConfiguration) Synthesize(session awscdk.ISynthesisSession) {
	if err := s.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_SecurityConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityConfiguration) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

