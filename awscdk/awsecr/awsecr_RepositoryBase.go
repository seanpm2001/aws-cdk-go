package awsecr

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Base class for ECR repository.
//
// Reused between imported repositories and owned repositories.
type RepositoryBase interface {
	awscdk.Resource
	IRepository
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The ARN of the repository.
	RepositoryArn() *string
	// The name of the repository.
	RepositoryName() *string
	// The URI of this repository (represents the latest image):.
	//
	// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY
	RepositoryUri() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Add a policy statement to the repository's resource policy.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
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
	// Grant the given principal identity permissions to perform the actions on this repository.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity permissions to use the images in this repository.
	GrantPull(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to pull and push images to this repository.
	GrantPullPush(grantee awsiam.IGrantable) awsiam.Grant
	// Define a CloudWatch event that triggers when something happens to this repository.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	OnCloudTrailEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an AWS CloudWatch event rule that can trigger a target when an image is pushed to this repository.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	OnCloudTrailImagePushed(id *string, options *OnCloudTrailImagePushedOptions) awsevents.Rule
	// Defines a CloudWatch event rule which triggers for repository events.
	//
	// Use
	// `rule.addEventPattern(pattern)` to specify a filter.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an AWS CloudWatch event rule that can trigger a target when an image scan is completed.
	OnImageScanCompleted(id *string, options *OnImageScanCompletedOptions) awsevents.Rule
	// Returns the URL of the repository. Can be used in `docker push/pull`.
	//
	// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY[@DIGEST]
	RepositoryUriForDigest(digest *string) *string
	// Returns the URL of the repository. Can be used in `docker push/pull`.
	//
	// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY[:TAG]
	RepositoryUriForTag(tag *string) *string
	// Returns the URL of the repository. Can be used in `docker push/pull`.
	//
	// ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY[:TAG]
	//     ACCOUNT.dkr.ecr.REGION.amazonaws.com/REPOSITORY[@DIGEST]
	RepositoryUriForTagOrDigest(tagOrDigest *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for RepositoryBase
type jsiiProxy_RepositoryBase struct {
	internal.Type__awscdkResource
	jsiiProxy_IRepository
}

func (j *jsiiProxy_RepositoryBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryBase) RepositoryArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryBase) RepositoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryBase) RepositoryUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewRepositoryBase_Override(r RepositoryBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr.RepositoryBase",
		[]interface{}{scope, id, props},
		r,
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
func RepositoryBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRepositoryBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.RepositoryBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func RepositoryBase_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateRepositoryBase_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.RepositoryBase",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func RepositoryBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateRepositoryBase_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr.RepositoryBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryBase) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	if err := r.validateAddToResourcePolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		r,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := r.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_RepositoryBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := r.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryBase) GetResourceNameAttribute(nameAttr *string) *string {
	if err := r.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryBase) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := r.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		r,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryBase) GrantPull(grantee awsiam.IGrantable) awsiam.Grant {
	if err := r.validateGrantPullParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		r,
		"grantPull",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryBase) GrantPullPush(grantee awsiam.IGrantable) awsiam.Grant {
	if err := r.validateGrantPullPushParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		r,
		"grantPullPush",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryBase) OnCloudTrailEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := r.validateOnCloudTrailEventParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		r,
		"onCloudTrailEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryBase) OnCloudTrailImagePushed(id *string, options *OnCloudTrailImagePushedOptions) awsevents.Rule {
	if err := r.validateOnCloudTrailImagePushedParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		r,
		"onCloudTrailImagePushed",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryBase) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := r.validateOnEventParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		r,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryBase) OnImageScanCompleted(id *string, options *OnImageScanCompletedOptions) awsevents.Rule {
	if err := r.validateOnImageScanCompletedParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		r,
		"onImageScanCompleted",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryBase) RepositoryUriForDigest(digest *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"repositoryUriForDigest",
		[]interface{}{digest},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryBase) RepositoryUriForTag(tag *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"repositoryUriForTag",
		[]interface{}{tag},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryBase) RepositoryUriForTagOrDigest(tagOrDigest *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"repositoryUriForTagOrDigest",
		[]interface{}{tagOrDigest},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

