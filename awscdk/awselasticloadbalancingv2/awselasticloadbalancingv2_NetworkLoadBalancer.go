package awselasticloadbalancingv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Define a new network load balancer.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("VPC"))
//   lb := elbv2.NewNetworkLoadBalancer(this, jsii.String("lb"), &networkLoadBalancerProps{
//   	vpc: vpc,
//   })
//   listener := lb.addListener(jsii.String("listener"), &baseNetworkListenerProps{
//   	port: jsii.Number(80),
//   })
//   listener.addTargets(jsii.String("target"), &addNetworkTargetsProps{
//   	port: jsii.Number(80),
//   })
//
//   httpEndpoint := apigwv2.NewHttpApi(this, jsii.String("HttpProxyPrivateApi"), &httpApiProps{
//   	defaultIntegration: awscdkapigatewayv2integrationsalpha.NewHttpNlbIntegration(jsii.String("DefaultIntegration"), listener),
//   })
//
type NetworkLoadBalancer interface {
	BaseLoadBalancer
	INetworkLoadBalancer
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The ARN of this load balancer.
	//
	// Example value: `arn:aws:elasticloadbalancing:us-west-2:123456789012:loadbalancer/app/my-internal-load-balancer/50dc6c495c0c9188`.
	LoadBalancerArn() *string
	// The canonical hosted zone ID of this load balancer.
	//
	// Example value: `Z2P70J7EXAMPLE`.
	LoadBalancerCanonicalHostedZoneId() *string
	// The DNS name of this load balancer.
	//
	// Example value: `my-load-balancer-424835706.us-west-2.elb.amazonaws.com`
	LoadBalancerDnsName() *string
	// The full name of this load balancer.
	//
	// Example value: `app/my-load-balancer/50dc6c495c0c9188`.
	LoadBalancerFullName() *string
	// The name of this load balancer.
	//
	// Example value: `my-load-balancer`.
	LoadBalancerName() *string
	LoadBalancerSecurityGroups() *[]*string
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
	// The VPC this load balancer has been created in.
	//
	// This property is always defined (not `null` or `undefined`) for sub-classes of `BaseLoadBalancer`.
	Vpc() awsec2.IVpc
	// Add a listener to this load balancer.
	//
	// Returns: The newly created listener.
	AddListener(id *string, props *BaseNetworkListenerProps) NetworkListener
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
	// Enable access logging for this load balancer.
	//
	// A region must be specified on the stack containing the load balancer; you cannot enable logging on
	// environment-agnostic stacks. See https://docs.aws.amazon.com/cdk/latest/guide/environments.html
	LogAccessLogs(bucket awss3.IBucket, prefix *string)
	// Return the given named metric for this Network Load Balancer.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of concurrent TCP flows (or connections) from clients to targets.
	//
	// This metric includes connections in the SYN_SENT and ESTABLISHED states.
	// TCP connections are not terminated at the load balancer, so a client
	// opening a TCP connection to a target counts as a single flow.
	MetricActiveFlowCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of load balancer capacity units (LCU) used by your load balancer.
	MetricConsumedLCUs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of new TCP flows (or connections) established from clients to targets in the time period.
	MetricNewFlowCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of bytes processed by the load balancer, including TCP/IP headers.
	MetricProcessedBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of reset (RST) packets sent from a client to a target.
	//
	// These resets are generated by the client and forwarded by the load balancer.
	MetricTcpClientResetCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of reset (RST) packets generated by the load balancer.
	MetricTcpElbResetCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of reset (RST) packets sent from a target to a client.
	//
	// These resets are generated by the target and forwarded by the load balancer.
	MetricTcpTargetResetCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Remove an attribute from the load balancer.
	RemoveAttribute(key *string)
	ResourcePolicyPrincipal() awsiam.IPrincipal
	// Set a non-standard attribute on the load balancer.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/application-load-balancers.html#load-balancer-attributes
	//
	SetAttribute(key *string, value *string)
	// Returns a string representation of this construct.
	ToString() *string
	ValidateLoadBalancer() *[]*string
}

// The jsii proxy struct for NetworkLoadBalancer
type jsiiProxy_NetworkLoadBalancer struct {
	jsiiProxy_BaseLoadBalancer
	jsiiProxy_INetworkLoadBalancer
}

func (j *jsiiProxy_NetworkLoadBalancer) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancer) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancer) LoadBalancerCanonicalHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerCanonicalHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancer) LoadBalancerDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancer) LoadBalancerFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancer) LoadBalancerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancer) LoadBalancerSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancer) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancer) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewNetworkLoadBalancer(scope constructs.Construct, id *string, props *NetworkLoadBalancerProps) NetworkLoadBalancer {
	_init_.Initialize()

	if err := validateNewNetworkLoadBalancerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkLoadBalancer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkLoadBalancer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNetworkLoadBalancer_Override(n NetworkLoadBalancer, scope constructs.Construct, id *string, props *NetworkLoadBalancerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkLoadBalancer",
		[]interface{}{scope, id, props},
		n,
	)
}

// Looks up the network load balancer.
func NetworkLoadBalancer_FromLookup(scope constructs.Construct, id *string, options *NetworkLoadBalancerLookupOptions) INetworkLoadBalancer {
	_init_.Initialize()

	if err := validateNetworkLoadBalancer_FromLookupParameters(scope, id, options); err != nil {
		panic(err)
	}
	var returns INetworkLoadBalancer

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkLoadBalancer",
		"fromLookup",
		[]interface{}{scope, id, options},
		&returns,
	)

	return returns
}

func NetworkLoadBalancer_FromNetworkLoadBalancerAttributes(scope constructs.Construct, id *string, attrs *NetworkLoadBalancerAttributes) INetworkLoadBalancer {
	_init_.Initialize()

	if err := validateNetworkLoadBalancer_FromNetworkLoadBalancerAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns INetworkLoadBalancer

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkLoadBalancer",
		"fromNetworkLoadBalancerAttributes",
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
func NetworkLoadBalancer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkLoadBalancer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkLoadBalancer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func NetworkLoadBalancer_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateNetworkLoadBalancer_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkLoadBalancer",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func NetworkLoadBalancer_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateNetworkLoadBalancer_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.NetworkLoadBalancer",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancer) AddListener(id *string, props *BaseNetworkListenerProps) NetworkListener {
	if err := n.validateAddListenerParameters(id, props); err != nil {
		panic(err)
	}
	var returns NetworkListener

	_jsii_.Invoke(
		n,
		"addListener",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := n.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (n *jsiiProxy_NetworkLoadBalancer) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancer) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := n.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancer) GetResourceNameAttribute(nameAttr *string) *string {
	if err := n.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancer) LogAccessLogs(bucket awss3.IBucket, prefix *string) {
	if err := n.validateLogAccessLogsParameters(bucket); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"logAccessLogs",
		[]interface{}{bucket, prefix},
	)
}

func (n *jsiiProxy_NetworkLoadBalancer) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := n.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancer) MetricActiveFlowCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := n.validateMetricActiveFlowCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
		"metricActiveFlowCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancer) MetricConsumedLCUs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := n.validateMetricConsumedLCUsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
		"metricConsumedLCUs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancer) MetricNewFlowCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := n.validateMetricNewFlowCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
		"metricNewFlowCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancer) MetricProcessedBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := n.validateMetricProcessedBytesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
		"metricProcessedBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancer) MetricTcpClientResetCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := n.validateMetricTcpClientResetCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
		"metricTcpClientResetCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancer) MetricTcpElbResetCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := n.validateMetricTcpElbResetCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
		"metricTcpElbResetCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancer) MetricTcpTargetResetCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := n.validateMetricTcpTargetResetCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
		"metricTcpTargetResetCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancer) RemoveAttribute(key *string) {
	if err := n.validateRemoveAttributeParameters(key); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"removeAttribute",
		[]interface{}{key},
	)
}

func (n *jsiiProxy_NetworkLoadBalancer) ResourcePolicyPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal

	_jsii_.Invoke(
		n,
		"resourcePolicyPrincipal",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancer) SetAttribute(key *string, value *string) {
	if err := n.validateSetAttributeParameters(key); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"setAttribute",
		[]interface{}{key, value},
	)
}

func (n *jsiiProxy_NetworkLoadBalancer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancer) ValidateLoadBalancer() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"validateLoadBalancer",
		nil, // no parameters
		&returns,
	)

	return returns
}

