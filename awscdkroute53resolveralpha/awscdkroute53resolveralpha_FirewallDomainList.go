// The CDK Construct Library for AWS::Route53Resolver
package awscdkroute53resolveralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkroute53resolveralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkroute53resolveralpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Firewall Domain List.
//
// Example:
//   blockList := route53resolver.NewFirewallDomainList(this, jsii.String("BlockList"), &firewallDomainListProps{
//   	domains: route53resolver.firewallDomains.fromList([]*string{
//   		jsii.String("bad-domain.com"),
//   		jsii.String("bot-domain.net"),
//   	}),
//   })
//
//   s3List := route53resolver.NewFirewallDomainList(this, jsii.String("S3List"), &firewallDomainListProps{
//   	domains: route53resolver.*firewallDomains.fromS3Url(jsii.String("s3://bucket/prefix/object")),
//   })
//
//   assetList := route53resolver.NewFirewallDomainList(this, jsii.String("AssetList"), &firewallDomainListProps{
//   	domains: route53resolver.*firewallDomains.fromAsset(jsii.String("/path/to/domains.txt")),
//   })
//
// Experimental.
type FirewallDomainList interface {
	awscdk.Resource
	IFirewallDomainList
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
	// The ARN (Amazon Resource Name) of the domain list.
	// Experimental.
	FirewallDomainListArn() *string
	// The date and time that the domain list was created.
	// Experimental.
	FirewallDomainListCreationTime() *string
	// The creator request ID.
	// Experimental.
	FirewallDomainListCreatorRequestId() *string
	// The number of domains in the list.
	// Experimental.
	FirewallDomainListDomainCount() *float64
	// The ID of the domain list.
	// Experimental.
	FirewallDomainListId() *string
	// The owner of the list, used only for lists that are not managed by you.
	//
	// For example, the managed domain list `AWSManagedDomainsMalwareDomainList`
	// has the managed owner name `Route 53 Resolver DNS Firewall`.
	// Experimental.
	FirewallDomainListManagedOwnerName() *string
	// The date and time that the domain list was last modified.
	// Experimental.
	FirewallDomainListModificationTime() *string
	// The status of the domain list.
	// Experimental.
	FirewallDomainListStatus() *string
	// Additional information about the status of the rule group.
	// Experimental.
	FirewallDomainListStatusMessage() *string
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

// The jsii proxy struct for FirewallDomainList
type jsiiProxy_FirewallDomainList struct {
	internal.Type__awscdkResource
	jsiiProxy_IFirewallDomainList
}

func (j *jsiiProxy_FirewallDomainList) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListCreatorRequestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListCreatorRequestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListDomainCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firewallDomainListDomainCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListManagedOwnerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListManagedOwnerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListModificationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListModificationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) FirewallDomainListStatusMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallDomainListStatusMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallDomainList) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewFirewallDomainList(scope constructs.Construct, id *string, props *FirewallDomainListProps) FirewallDomainList {
	_init_.Initialize()

	if err := validateNewFirewallDomainListParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirewallDomainList{}

	_jsii_.Create(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomainList",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewFirewallDomainList_Override(f FirewallDomainList, scope constructs.Construct, id *string, props *FirewallDomainListProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomainList",
		[]interface{}{scope, id, props},
		f,
	)
}

// Import an existing Firewall Rule Group.
// Experimental.
func FirewallDomainList_FromFirewallDomainListId(scope constructs.Construct, id *string, firewallDomainListId *string) IFirewallDomainList {
	_init_.Initialize()

	if err := validateFirewallDomainList_FromFirewallDomainListIdParameters(scope, id, firewallDomainListId); err != nil {
		panic(err)
	}
	var returns IFirewallDomainList

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomainList",
		"fromFirewallDomainListId",
		[]interface{}{scope, id, firewallDomainListId},
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
func FirewallDomainList_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFirewallDomainList_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomainList",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func FirewallDomainList_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateFirewallDomainList_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomainList",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func FirewallDomainList_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateFirewallDomainList_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-route53resolver-alpha.FirewallDomainList",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallDomainList) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := f.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (f *jsiiProxy_FirewallDomainList) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallDomainList) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := f.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallDomainList) GetResourceNameAttribute(nameAttr *string) *string {
	if err := f.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallDomainList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

