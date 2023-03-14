package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// A interface VPC endpoint.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var vpc vpc
//
//
//   ec2.NewInterfaceVpcEndpoint(this, jsii.String("VPC Endpoint"), &InterfaceVpcEndpointProps{
//   	Vpc: Vpc,
//   	Service: ec2.NewInterfaceVpcEndpointService(jsii.String("com.amazonaws.vpce.us-east-1.vpce-svc-uuddlrlrbastrtsvc"), jsii.Number(443)),
//   	// Choose which availability zones to place the VPC endpoint in, based on
//   	// available AZs
//   	Subnets: &SubnetSelection{
//   		AvailabilityZones: []*string{
//   			jsii.String("us-east-1a"),
//   			jsii.String("us-east-1c"),
//   		},
//   	},
//   })
//
type InterfaceVpcEndpoint interface {
	VpcEndpoint
	IInterfaceVpcEndpoint
	// Access to network connections.
	Connections() Connections
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
	PolicyDocument() awsiam.PolicyDocument
	SetPolicyDocument(val awsiam.PolicyDocument)
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The date and time the interface VPC endpoint was created.
	VpcEndpointCreationTimestamp() *string
	// The DNS entries for the interface VPC endpoint.
	//
	// Each entry is a combination of the hosted zone ID and the DNS name.
	// The entries are ordered as follows: regional public DNS, zonal public DNS, private DNS, and wildcard DNS.
	// This order is not enforced for AWS Marketplace services.
	//
	// The following is an example. In the first entry, the hosted zone ID is Z1HUB23UULQXV
	// and the DNS name is vpce-01abc23456de78f9g-12abccd3.ec2.us-east-1.vpce.amazonaws.com.
	//
	// ["Z1HUB23UULQXV:vpce-01abc23456de78f9g-12abccd3.ec2.us-east-1.vpce.amazonaws.com",
	// "Z1HUB23UULQXV:vpce-01abc23456de78f9g-12abccd3-us-east-1a.ec2.us-east-1.vpce.amazonaws.com",
	// "Z1C12344VYDITB0:ec2.us-east-1.amazonaws.com"]
	//
	// If you update the PrivateDnsEnabled or SubnetIds properties, the DNS entries in the list will change.
	VpcEndpointDnsEntries() *[]*string
	// The interface VPC endpoint identifier.
	VpcEndpointId() *string
	// One or more network interfaces for the interface VPC endpoint.
	VpcEndpointNetworkInterfaceIds() *[]*string
	// Adds a statement to the policy document of the VPC endpoint. The statement must have a Principal.
	//
	// Not all interface VPC endpoints support policy. For more information
	// see https://docs.aws.amazon.com/vpc/latest/userguide/vpce-interface.html
	AddToPolicy(statement awsiam.PolicyStatement)
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
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for InterfaceVpcEndpoint
type jsiiProxy_InterfaceVpcEndpoint struct {
	jsiiProxy_VpcEndpoint
	jsiiProxy_IInterfaceVpcEndpoint
}

func (j *jsiiProxy_InterfaceVpcEndpoint) Connections() Connections {
	var returns Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceVpcEndpoint) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceVpcEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceVpcEndpoint) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceVpcEndpoint) PolicyDocument() awsiam.PolicyDocument {
	var returns awsiam.PolicyDocument
	_jsii_.Get(
		j,
		"policyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceVpcEndpoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceVpcEndpoint) VpcEndpointCreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointCreationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceVpcEndpoint) VpcEndpointDnsEntries() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcEndpointDnsEntries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceVpcEndpoint) VpcEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceVpcEndpoint) VpcEndpointNetworkInterfaceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcEndpointNetworkInterfaceIds",
		&returns,
	)
	return returns
}


func NewInterfaceVpcEndpoint(scope constructs.Construct, id *string, props *InterfaceVpcEndpointProps) InterfaceVpcEndpoint {
	_init_.Initialize()

	if err := validateNewInterfaceVpcEndpointParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_InterfaceVpcEndpoint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewInterfaceVpcEndpoint_Override(i InterfaceVpcEndpoint, scope constructs.Construct, id *string, props *InterfaceVpcEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpoint",
		[]interface{}{scope, id, props},
		i,
	)
}

func (j *jsiiProxy_InterfaceVpcEndpoint)SetPolicyDocument(val awsiam.PolicyDocument) {
	_jsii_.Set(
		j,
		"policyDocument",
		val,
	)
}

// Imports an existing interface VPC endpoint.
func InterfaceVpcEndpoint_FromInterfaceVpcEndpointAttributes(scope constructs.Construct, id *string, attrs *InterfaceVpcEndpointAttributes) IInterfaceVpcEndpoint {
	_init_.Initialize()

	if err := validateInterfaceVpcEndpoint_FromInterfaceVpcEndpointAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IInterfaceVpcEndpoint

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpoint",
		"fromInterfaceVpcEndpointAttributes",
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
func InterfaceVpcEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateInterfaceVpcEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func InterfaceVpcEndpoint_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateInterfaceVpcEndpoint_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpoint",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func InterfaceVpcEndpoint_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateInterfaceVpcEndpoint_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InterfaceVpcEndpoint",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InterfaceVpcEndpoint) AddToPolicy(statement awsiam.PolicyStatement) {
	if err := i.validateAddToPolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addToPolicy",
		[]interface{}{statement},
	)
}

func (i *jsiiProxy_InterfaceVpcEndpoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_InterfaceVpcEndpoint) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InterfaceVpcEndpoint) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := i.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InterfaceVpcEndpoint) GetResourceNameAttribute(nameAttr *string) *string {
	if err := i.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InterfaceVpcEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

