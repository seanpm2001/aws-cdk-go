package awss3

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/constructs-go/constructs/v10"
)

// An S3 bucket with associated policy objects.
//
// This bucket does not yet have all features that exposed by the underlying
// BucketResource.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   NewBucket(*scope, jsii.String("Bucket"), map[string]interface{}{
//   	"blockPublicAccess": BlockPublicAccess_BLOCK_ALL,
//   	"encryption": BucketEncryption_S3_MANAGED,
//   	"enforceSSL": jsii.Boolean(true),
//   	"versioned": jsii.Boolean(true),
//   	"removalPolicy": RemovalPolicy_RETAIN,
//   })
//
type Bucket interface {
	BucketBase
	// Indicates if a bucket resource policy should automatically created upon the first call to `addToResourcePolicy`.
	AutoCreatePolicy() *bool
	SetAutoCreatePolicy(val *bool)
	// The ARN of the bucket.
	BucketArn() *string
	// The IPv4 DNS name of the specified bucket.
	BucketDomainName() *string
	// The IPv6 DNS name of the specified bucket.
	BucketDualStackDomainName() *string
	// The name of the bucket.
	BucketName() *string
	// The regional domain name of the specified bucket.
	BucketRegionalDomainName() *string
	// The Domain name of the static website.
	BucketWebsiteDomainName() *string
	// The URL of the static website.
	BucketWebsiteUrl() *string
	// Whether to disallow public access.
	DisallowPublicAccess() *bool
	SetDisallowPublicAccess(val *bool)
	// Optional KMS encryption key associated with this bucket.
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
	// If this bucket has been configured for static website hosting.
	IsWebsite() *bool
	// The tree node.
	Node() constructs.Node
	NotificationsHandlerRole() awsiam.IRole
	SetNotificationsHandlerRole(val awsiam.IRole)
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The resource policy associated with this bucket.
	//
	// If `autoCreatePolicy` is true, a `BucketPolicy` will be created upon the
	// first call to addToResourcePolicy(s).
	Policy() BucketPolicy
	SetPolicy(val BucketPolicy)
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Adds a cross-origin access configuration for objects in an Amazon S3 bucket.
	AddCorsRule(rule *CorsRule)
	// Adds a bucket notification event destination.
	//
	// Example:
	//      declare const myLambda: lambda.Function;
	//      const bucket = new s3.Bucket(this, 'MyBucket');
	//      bucket.addEventNotification(s3.EventType.OBJECT_CREATED, new s3n.LambdaDestination(myLambda), {prefix: 'home/myusername/*'});
	//
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html
	//
	AddEventNotification(event EventType, dest IBucketNotificationDestination, filters ...*NotificationKeyFilter)
	// Add an inventory configuration.
	AddInventory(inventory *Inventory)
	// Add a lifecycle rule to the bucket.
	AddLifecycleRule(rule *LifecycleRule)
	// Adds a metrics configuration for the CloudWatch request metrics from the bucket.
	AddMetric(metric *BucketMetrics)
	// Subscribes a destination to receive notifications when an object is created in the bucket.
	//
	// This is identical to calling
	// `onEvent(EventType.OBJECT_CREATED)`.
	AddObjectCreatedNotification(dest IBucketNotificationDestination, filters ...*NotificationKeyFilter)
	// Subscribes a destination to receive notifications when an object is removed from the bucket.
	//
	// This is identical to calling
	// `onEvent(EventType.OBJECT_REMOVED)`.
	AddObjectRemovedNotification(dest IBucketNotificationDestination, filters ...*NotificationKeyFilter)
	// Adds a statement to the resource policy for a principal (i.e. account/role/service) to perform actions on this bucket and/or its contents. Use `bucketArn` and `arnForObjects(keys)` to obtain ARNs for this bucket or objects.
	//
	// Note that the policy statement may or may not be added to the policy.
	// For example, when an `IBucket` is created from an existing bucket,
	// it's not possible to tell whether the bucket already has a policy
	// attached, let alone to re-use that policy to add more statements to it.
	// So it's safest to do nothing in these cases.
	//
	// Returns: metadata about the execution of this method. If the policy
	// was not added, the value of `statementAdded` will be `false`. You
	// should always check this value to make sure that the operation was
	// actually carried out. Otherwise, synthesis and deploy will terminate
	// silently, which may be confusing.
	AddToResourcePolicy(permission awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
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
	// Returns an ARN that represents all objects within the bucket that match the key pattern specified.
	//
	// To represent all keys, specify ``"*"``.
	//
	// If you need to specify a keyPattern with multiple components, concatenate them into a single string, e.g.:
	//
	// arnForObjects(`home/${team}/${user}/*`).
	ArnForObjects(keyPattern *string) *string
	// Enables event bridge notification, causing all events below to be sent to EventBridge:.
	//
	// - Object Deleted (DeleteObject)
	// - Object Deleted (Lifecycle expiration)
	// - Object Restore Initiated
	// - Object Restore Completed
	// - Object Restore Expired
	// - Object Storage Class Changed
	// - Object Access Tier Changed
	// - Object ACL Updated
	// - Object Tags Added
	// - Object Tags Deleted.
	EnableEventBridgeNotification()
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
	// Grants s3:DeleteObject* permission to an IAM principal for objects in this bucket.
	GrantDelete(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Allows unrestricted access to objects from this bucket.
	//
	// IMPORTANT: This permission allows anyone to perform actions on S3 objects
	// in this bucket, which is useful for when you configure your bucket as a
	// website and want everyone to be able to read objects in the bucket without
	// needing to authenticate.
	//
	// Without arguments, this method will grant read ("s3:GetObject") access to
	// all objects ("*") in the bucket.
	//
	// The method returns the `iam.Grant` object, which can then be modified
	// as needed. For example, you can add a condition that will restrict access only
	// to an IPv4 range like this:
	//
	//      const grant = bucket.grantPublicAccess();
	//      grant.resourceStatement!.addCondition(‘IpAddress’, { “aws:SourceIp”: “54.240.143.0/24” });
	//
	// Note that if this `IBucket` refers to an existing bucket, possibly not
	// managed by CloudFormation, this method will have no effect, since it's
	// impossible to modify the policy of an existing bucket.
	GrantPublicAccess(keyPrefix *string, allowedActions ...*string) awsiam.Grant
	// Grants s3:PutObject* and s3:Abort* permissions for this bucket to an IAM principal.
	//
	// If encryption is used, permission to use the key to encrypt the contents
	// of written files will also be granted to the same principal.
	GrantPut(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Grant the given IAM identity permissions to modify the ACLs of objects in the given Bucket.
	//
	// If your application has the '@aws-cdk/aws-s3:grantWriteWithoutAcl' feature flag set,
	// calling {@link grantWrite} or {@link grantReadWrite} no longer grants permissions to modify the ACLs of the objects;
	// in this case, if you need to modify object ACLs, call this method explicitly.
	GrantPutAcl(identity awsiam.IGrantable, objectsKeyPattern *string) awsiam.Grant
	// Grant read permissions for this bucket and it's contents to an IAM principal (Role/Group/User).
	//
	// If encryption is used, permission to use the key to decrypt the contents
	// of the bucket will also be granted to the same principal.
	GrantRead(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Grants read/write permissions for this bucket and it's contents to an IAM principal (Role/Group/User).
	//
	// If an encryption key is used, permission to use the key for
	// encrypt/decrypt will also be granted.
	//
	// Before CDK version 1.85.0, this method granted the `s3:PutObject*` permission that included `s3:PutObjectAcl`,
	// which could be used to grant read/write object access to IAM principals in other accounts.
	// If you want to get rid of that behavior, update your CDK version to 1.85.0 or later,
	// and make sure the `@aws-cdk/aws-s3:grantWriteWithoutAcl` feature flag is set to `true`
	// in the `context` key of your cdk.json file.
	// If you've already updated, but still need the principal to have permissions to modify the ACLs,
	// use the {@link grantPutAcl} method.
	GrantReadWrite(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Grant write permissions to this bucket to an IAM principal.
	//
	// If encryption is used, permission to use the key to encrypt the contents
	// of written files will also be granted to the same principal.
	//
	// Before CDK version 1.85.0, this method granted the `s3:PutObject*` permission that included `s3:PutObjectAcl`,
	// which could be used to grant read/write object access to IAM principals in other accounts.
	// If you want to get rid of that behavior, update your CDK version to 1.85.0 or later,
	// and make sure the `@aws-cdk/aws-s3:grantWriteWithoutAcl` feature flag is set to `true`
	// in the `context` key of your cdk.json file.
	// If you've already updated, but still need the principal to have permissions to modify the ACLs,
	// use the {@link grantPutAcl} method.
	GrantWrite(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Define a CloudWatch event that triggers when something happens to this repository.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	OnCloudTrailEvent(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule
	// Defines an AWS CloudWatch event that triggers when an object is uploaded to the specified paths (keys) in this bucket using the PutObject API call.
	//
	// Note that some tools like `aws s3 cp` will automatically use either
	// PutObject or the multipart upload API depending on the file size,
	// so using `onCloudTrailWriteObject` may be preferable.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	OnCloudTrailPutObject(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule
	// Defines an AWS CloudWatch event that triggers when an object at the specified paths (keys) in this bucket are written to.
	//
	// This includes
	// the events PutObject, CopyObject, and CompleteMultipartUpload.
	//
	// Note that some tools like `aws s3 cp` will automatically use either
	// PutObject or the multipart upload API depending on the file size,
	// so using this method may be preferable to `onCloudTrailPutObject`.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	OnCloudTrailWriteObject(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule
	// The S3 URL of an S3 object. For example:.
	//
	// - `s3://onlybucket`
	// - `s3://bucket/key`.
	//
	// Returns: an ObjectS3Url token.
	S3UrlForObject(key *string) *string
	// Returns a string representation of this construct.
	ToString() *string
	// The https Transfer Acceleration URL of an S3 object.
	//
	// Specify `dualStack: true` at the options
	// for dual-stack endpoint (connect to the bucket over IPv6). For example:
	//
	// - `https://bucket.s3-accelerate.amazonaws.com`
	// - `https://bucket.s3-accelerate.amazonaws.com/key`
	//
	// Returns: an TransferAccelerationUrl token.
	TransferAccelerationUrlForObject(key *string, options *TransferAccelerationUrlOptions) *string
	// The https URL of an S3 object. Specify `regional: false` at the options for non-regional URLs. For example:.
	//
	// - `https://s3.us-west-1.amazonaws.com/onlybucket`
	// - `https://s3.us-west-1.amazonaws.com/bucket/key`
	// - `https://s3.cn-north-1.amazonaws.com.cn/china-bucket/mykey`
	//
	// Returns: an ObjectS3Url token.
	UrlForObject(key *string) *string
	// The virtual hosted-style URL of an S3 object. Specify `regional: false` at the options for non-regional URL. For example:.
	//
	// - `https://only-bucket.s3.us-west-1.amazonaws.com`
	// - `https://bucket.s3.us-west-1.amazonaws.com/key`
	// - `https://bucket.s3.amazonaws.com/key`
	// - `https://china-bucket.s3.cn-north-1.amazonaws.com.cn/mykey`
	//
	// Returns: an ObjectS3Url token.
	VirtualHostedUrlForObject(key *string, options *VirtualHostedStyleUrlOptions) *string
}

// The jsii proxy struct for Bucket
type jsiiProxy_Bucket struct {
	jsiiProxy_BucketBase
}

func (j *jsiiProxy_Bucket) AutoCreatePolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoCreatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) BucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) BucketDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) BucketDualStackDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketDualStackDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) BucketRegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegionalDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) BucketWebsiteDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketWebsiteDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) BucketWebsiteUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketWebsiteUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) DisallowPublicAccess() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"disallowPublicAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) IsWebsite() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isWebsite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) NotificationsHandlerRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"notificationsHandlerRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) Policy() BucketPolicy {
	var returns BucketPolicy
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewBucket(scope constructs.Construct, id *string, props *BucketProps) Bucket {
	_init_.Initialize()

	if err := validateNewBucketParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Bucket{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.Bucket",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewBucket_Override(b Bucket, scope constructs.Construct, id *string, props *BucketProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.Bucket",
		[]interface{}{scope, id, props},
		b,
	)
}

func (j *jsiiProxy_Bucket)SetAutoCreatePolicy(val *bool) {
	if err := j.validateSetAutoCreatePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoCreatePolicy",
		val,
	)
}

func (j *jsiiProxy_Bucket)SetDisallowPublicAccess(val *bool) {
	_jsii_.Set(
		j,
		"disallowPublicAccess",
		val,
	)
}

func (j *jsiiProxy_Bucket)SetNotificationsHandlerRole(val awsiam.IRole) {
	_jsii_.Set(
		j,
		"notificationsHandlerRole",
		val,
	)
}

func (j *jsiiProxy_Bucket)SetPolicy(val BucketPolicy) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func Bucket_FromBucketArn(scope constructs.Construct, id *string, bucketArn *string) IBucket {
	_init_.Initialize()

	if err := validateBucket_FromBucketArnParameters(scope, id, bucketArn); err != nil {
		panic(err)
	}
	var returns IBucket

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.Bucket",
		"fromBucketArn",
		[]interface{}{scope, id, bucketArn},
		&returns,
	)

	return returns
}

// Creates a Bucket construct that represents an external bucket.
func Bucket_FromBucketAttributes(scope constructs.Construct, id *string, attrs *BucketAttributes) IBucket {
	_init_.Initialize()

	if err := validateBucket_FromBucketAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IBucket

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.Bucket",
		"fromBucketAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

func Bucket_FromBucketName(scope constructs.Construct, id *string, bucketName *string) IBucket {
	_init_.Initialize()

	if err := validateBucket_FromBucketNameParameters(scope, id, bucketName); err != nil {
		panic(err)
	}
	var returns IBucket

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.Bucket",
		"fromBucketName",
		[]interface{}{scope, id, bucketName},
		&returns,
	)

	return returns
}

// Create a mutable {@link IBucket} based on a low-level {@link CfnBucket}.
func Bucket_FromCfnBucket(cfnBucket CfnBucket) IBucket {
	_init_.Initialize()

	if err := validateBucket_FromCfnBucketParameters(cfnBucket); err != nil {
		panic(err)
	}
	var returns IBucket

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.Bucket",
		"fromCfnBucket",
		[]interface{}{cfnBucket},
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
func Bucket_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBucket_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.Bucket",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Bucket_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateBucket_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.Bucket",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Bucket_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateBucket_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.Bucket",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Thrown an exception if the given bucket name is not valid.
func Bucket_ValidateBucketName(physicalName *string) {
	_init_.Initialize()

	if err := validateBucket_ValidateBucketNameParameters(physicalName); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_s3.Bucket",
		"validateBucketName",
		[]interface{}{physicalName},
	)
}

func (b *jsiiProxy_Bucket) AddCorsRule(rule *CorsRule) {
	if err := b.validateAddCorsRuleParameters(rule); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addCorsRule",
		[]interface{}{rule},
	)
}

func (b *jsiiProxy_Bucket) AddEventNotification(event EventType, dest IBucketNotificationDestination, filters ...*NotificationKeyFilter) {
	if err := b.validateAddEventNotificationParameters(event, dest, &filters); err != nil {
		panic(err)
	}
	args := []interface{}{event, dest}
	for _, a := range filters {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		b,
		"addEventNotification",
		args,
	)
}

func (b *jsiiProxy_Bucket) AddInventory(inventory *Inventory) {
	if err := b.validateAddInventoryParameters(inventory); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addInventory",
		[]interface{}{inventory},
	)
}

func (b *jsiiProxy_Bucket) AddLifecycleRule(rule *LifecycleRule) {
	if err := b.validateAddLifecycleRuleParameters(rule); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addLifecycleRule",
		[]interface{}{rule},
	)
}

func (b *jsiiProxy_Bucket) AddMetric(metric *BucketMetrics) {
	if err := b.validateAddMetricParameters(metric); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMetric",
		[]interface{}{metric},
	)
}

func (b *jsiiProxy_Bucket) AddObjectCreatedNotification(dest IBucketNotificationDestination, filters ...*NotificationKeyFilter) {
	if err := b.validateAddObjectCreatedNotificationParameters(dest, &filters); err != nil {
		panic(err)
	}
	args := []interface{}{dest}
	for _, a := range filters {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		b,
		"addObjectCreatedNotification",
		args,
	)
}

func (b *jsiiProxy_Bucket) AddObjectRemovedNotification(dest IBucketNotificationDestination, filters ...*NotificationKeyFilter) {
	if err := b.validateAddObjectRemovedNotificationParameters(dest, &filters); err != nil {
		panic(err)
	}
	args := []interface{}{dest}
	for _, a := range filters {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		b,
		"addObjectRemovedNotification",
		args,
	)
}

func (b *jsiiProxy_Bucket) AddToResourcePolicy(permission awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	if err := b.validateAddToResourcePolicyParameters(permission); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		b,
		"addToResourcePolicy",
		[]interface{}{permission},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := b.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (b *jsiiProxy_Bucket) ArnForObjects(keyPattern *string) *string {
	if err := b.validateArnForObjectsParameters(keyPattern); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"arnForObjects",
		[]interface{}{keyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) EnableEventBridgeNotification() {
	_jsii_.InvokeVoid(
		b,
		"enableEventBridgeNotification",
		nil, // no parameters
	)
}

func (b *jsiiProxy_Bucket) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (b *jsiiProxy_Bucket) GetResourceNameAttribute(nameAttr *string) *string {
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

func (b *jsiiProxy_Bucket) GrantDelete(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	if err := b.validateGrantDeleteParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantDelete",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) GrantPublicAccess(keyPrefix *string, allowedActions ...*string) awsiam.Grant {
	args := []interface{}{keyPrefix}
	for _, a := range allowedActions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantPublicAccess",
		args,
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) GrantPut(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	if err := b.validateGrantPutParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantPut",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) GrantPutAcl(identity awsiam.IGrantable, objectsKeyPattern *string) awsiam.Grant {
	if err := b.validateGrantPutAclParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantPutAcl",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) GrantRead(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	if err := b.validateGrantReadParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantRead",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) GrantReadWrite(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	if err := b.validateGrantReadWriteParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantReadWrite",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) GrantWrite(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	if err := b.validateGrantWriteParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantWrite",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) OnCloudTrailEvent(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule {
	if err := b.validateOnCloudTrailEventParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		b,
		"onCloudTrailEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) OnCloudTrailPutObject(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule {
	if err := b.validateOnCloudTrailPutObjectParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		b,
		"onCloudTrailPutObject",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) OnCloudTrailWriteObject(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule {
	if err := b.validateOnCloudTrailWriteObjectParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		b,
		"onCloudTrailWriteObject",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) S3UrlForObject(key *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"s3UrlForObject",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) TransferAccelerationUrlForObject(key *string, options *TransferAccelerationUrlOptions) *string {
	if err := b.validateTransferAccelerationUrlForObjectParameters(options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"transferAccelerationUrlForObject",
		[]interface{}{key, options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) UrlForObject(key *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"urlForObject",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) VirtualHostedUrlForObject(key *string, options *VirtualHostedStyleUrlOptions) *string {
	if err := b.validateVirtualHostedUrlForObjectParameters(options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"virtualHostedUrlForObject",
		[]interface{}{key, options},
		&returns,
	)

	return returns
}

