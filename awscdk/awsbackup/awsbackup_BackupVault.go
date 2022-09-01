package awsbackup

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbackup/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// A backup vault.
//
// Example:
//   importedVault := backup.backupVault.fromBackupVaultName(this, jsii.String("Vault"), jsii.String("myVaultName"))
//
//   role := iam.NewRole(this, jsii.String("Access Role"), &roleProps{
//   	assumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
//   })
//
//   importedVault.grant(role, jsii.String("backup:StartBackupJob"))
//
type BackupVault interface {
	awscdk.Resource
	IBackupVault
	// The ARN of the backup vault.
	BackupVaultArn() *string
	// The name of a logical container where backups are stored.
	BackupVaultName() *string
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
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Adds a statement to the vault access policy.
	AddToAccessPolicy(statement awsiam.PolicyStatement)
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
	// Adds a statement to the vault access policy that prevents anyone from deleting a recovery point.
	BlockRecoveryPointDeletion()
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
	// Grant the actions defined in actions to the given grantee on this Backup Vault resource.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for BackupVault
type jsiiProxy_BackupVault struct {
	internal.Type__awscdkResource
	jsiiProxy_IBackupVault
}

func (j *jsiiProxy_BackupVault) BackupVaultArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVaultArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupVault) BackupVaultName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVaultName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupVault) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupVault) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupVault) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupVault) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewBackupVault(scope constructs.Construct, id *string, props *BackupVaultProps) BackupVault {
	_init_.Initialize()

	if err := validateNewBackupVaultParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupVault{}

	_jsii_.Create(
		"aws-cdk-lib.aws_backup.BackupVault",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewBackupVault_Override(b BackupVault, scope constructs.Construct, id *string, props *BackupVaultProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_backup.BackupVault",
		[]interface{}{scope, id, props},
		b,
	)
}

// Import an existing backup vault by arn.
func BackupVault_FromBackupVaultArn(scope constructs.Construct, id *string, backupVaultArn *string) IBackupVault {
	_init_.Initialize()

	if err := validateBackupVault_FromBackupVaultArnParameters(scope, id, backupVaultArn); err != nil {
		panic(err)
	}
	var returns IBackupVault

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.BackupVault",
		"fromBackupVaultArn",
		[]interface{}{scope, id, backupVaultArn},
		&returns,
	)

	return returns
}

// Import an existing backup vault by name.
func BackupVault_FromBackupVaultName(scope constructs.Construct, id *string, backupVaultName *string) IBackupVault {
	_init_.Initialize()

	if err := validateBackupVault_FromBackupVaultNameParameters(scope, id, backupVaultName); err != nil {
		panic(err)
	}
	var returns IBackupVault

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.BackupVault",
		"fromBackupVaultName",
		[]interface{}{scope, id, backupVaultName},
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
func BackupVault_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBackupVault_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.BackupVault",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func BackupVault_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateBackupVault_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.BackupVault",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func BackupVault_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateBackupVault_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.BackupVault",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupVault) AddToAccessPolicy(statement awsiam.PolicyStatement) {
	if err := b.validateAddToAccessPolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addToAccessPolicy",
		[]interface{}{statement},
	)
}

func (b *jsiiProxy_BackupVault) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := b.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (b *jsiiProxy_BackupVault) BlockRecoveryPointDeletion() {
	_jsii_.InvokeVoid(
		b,
		"blockRecoveryPointDeletion",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupVault) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupVault) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := b.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupVault) GetResourceNameAttribute(nameAttr *string) *string {
	if err := b.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupVault) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := b.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupVault) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

