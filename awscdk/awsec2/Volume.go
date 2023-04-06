package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new EBS Volume in AWS EC2.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var instance instance
//   var role role
//
//
//   volume := ec2.NewVolume(this, jsii.String("Volume"), &VolumeProps{
//   	AvailabilityZone: jsii.String("us-west-2a"),
//   	Size: awscdk.Size_Gibibytes(jsii.Number(500)),
//   	Encrypted: jsii.Boolean(true),
//   })
//
//   volume.grantAttachVolume(role, []iInstance{
//   	instance,
//   })
//
type Volume interface {
	awscdk.Resource
	IVolume
	// The availability zone that the EBS Volume is contained within (ex: us-west-2a).
	AvailabilityZone() *string
	// The customer-managed encryption key that is used to encrypt the Volume.
	EncryptionKey() awskms.IKey
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
	//   cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The EBS Volume's ID.
	VolumeId() *string
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
	// Grants permission to attach this Volume to an instance.
	//
	// CAUTION: Granting an instance permission to attach to itself using this method will lead to
	// an unresolvable circular reference between the instance role and the instance.
	// Use `IVolume.grantAttachVolumeToSelf` to grant an instance permission to attach this
	// volume to itself.
	GrantAttachVolume(grantee awsiam.IGrantable, instances *[]IInstance) awsiam.Grant
	// Grants permission to attach the Volume by a ResourceTag condition.
	//
	// If you are looking to
	// grant an Instance, AutoScalingGroup, EC2-Fleet, SpotFleet, ECS host, etc the ability to attach
	// this volume to **itself** then this is the method you want to use.
	//
	// This is implemented by adding a Tag with key `VolumeGrantAttach-<suffix>` to the given
	// constructs and this Volume, and then conditioning the Grant such that the grantee is only
	// given the ability to AttachVolume if both the Volume and the destination Instance have that
	// tag applied to them.
	GrantAttachVolumeByResourceTag(grantee awsiam.IGrantable, constructs *[]constructs.Construct, tagKeySuffix *string) awsiam.Grant
	// Grants permission to detach this Volume from an instance CAUTION: Granting an instance permission to detach from itself using this method will lead to an unresolvable circular reference between the instance role and the instance.
	//
	// Use `IVolume.grantDetachVolumeFromSelf` to grant an instance permission to detach this
	// volume from itself.
	GrantDetachVolume(grantee awsiam.IGrantable, instances *[]IInstance) awsiam.Grant
	// Grants permission to detach the Volume by a ResourceTag condition.
	//
	// This is implemented via the same mechanism as `IVolume.grantAttachVolumeByResourceTag`,
	// and is subject to the same conditions.
	GrantDetachVolumeByResourceTag(grantee awsiam.IGrantable, constructs *[]constructs.Construct, tagKeySuffix *string) awsiam.Grant
	// Returns a string representation of this construct.
	ToString() *string
	ValidateProps(props *VolumeProps)
}

// The jsii proxy struct for Volume
type jsiiProxy_Volume struct {
	internal.Type__awscdkResource
	jsiiProxy_IVolume
}

func (j *jsiiProxy_Volume) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Volume) VolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeId",
		&returns,
	)
	return returns
}


func NewVolume(scope constructs.Construct, id *string, props *VolumeProps) Volume {
	_init_.Initialize()

	if err := validateNewVolumeParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Volume{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.Volume",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewVolume_Override(v Volume, scope constructs.Construct, id *string, props *VolumeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.Volume",
		[]interface{}{scope, id, props},
		v,
	)
}

// Import an existing EBS Volume into the Stack.
func Volume_FromVolumeAttributes(scope constructs.Construct, id *string, attrs *VolumeAttributes) IVolume {
	_init_.Initialize()

	if err := validateVolume_FromVolumeAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IVolume

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Volume",
		"fromVolumeAttributes",
		[]interface{}{scope, id, attrs},
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
func Volume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVolume_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Volume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Volume_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateVolume_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Volume",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Volume_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateVolume_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Volume",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := v.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (v *jsiiProxy_Volume) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := v.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) GetResourceNameAttribute(nameAttr *string) *string {
	if err := v.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) GrantAttachVolume(grantee awsiam.IGrantable, instances *[]IInstance) awsiam.Grant {
	if err := v.validateGrantAttachVolumeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		v,
		"grantAttachVolume",
		[]interface{}{grantee, instances},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) GrantAttachVolumeByResourceTag(grantee awsiam.IGrantable, constructs *[]constructs.Construct, tagKeySuffix *string) awsiam.Grant {
	if err := v.validateGrantAttachVolumeByResourceTagParameters(grantee, constructs); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		v,
		"grantAttachVolumeByResourceTag",
		[]interface{}{grantee, constructs, tagKeySuffix},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) GrantDetachVolume(grantee awsiam.IGrantable, instances *[]IInstance) awsiam.Grant {
	if err := v.validateGrantDetachVolumeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		v,
		"grantDetachVolume",
		[]interface{}{grantee, instances},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) GrantDetachVolumeByResourceTag(grantee awsiam.IGrantable, constructs *[]constructs.Construct, tagKeySuffix *string) awsiam.Grant {
	if err := v.validateGrantDetachVolumeByResourceTagParameters(grantee, constructs); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		v,
		"grantDetachVolumeByResourceTag",
		[]interface{}{grantee, constructs, tagKeySuffix},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_Volume) ValidateProps(props *VolumeProps) {
	if err := v.validateValidatePropsParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"validateProps",
		[]interface{}{props},
	)
}

