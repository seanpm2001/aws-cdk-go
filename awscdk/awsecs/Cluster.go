package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery"
	"github.com/aws/constructs-go/constructs/v10"
)

// A regional grouping of one or more container instances on which you can run tasks and services.
//
// Example:
//   vpc := ec2.Vpc_FromLookup(this, jsii.String("Vpc"), &VpcLookupOptions{
//   	IsDefault: jsii.Boolean(true),
//   })
//
//   cluster := ecs.NewCluster(this, jsii.String("FargateCluster"), &ClusterProps{
//   	Vpc: Vpc,
//   })
//
//   taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TD"), &TaskDefinitionProps{
//   	MemoryMiB: jsii.String("512"),
//   	Cpu: jsii.String("256"),
//   	Compatibility: ecs.Compatibility_FARGATE,
//   })
//
//   containerDefinition := taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("foo/bar")),
//   	MemoryLimitMiB: jsii.Number(256),
//   })
//
//   runTask := tasks.NewEcsRunTask(this, jsii.String("RunFargate"), &EcsRunTaskProps{
//   	IntegrationPattern: sfn.IntegrationPattern_RUN_JOB,
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	AssignPublicIp: jsii.Boolean(true),
//   	ContainerOverrides: []containerOverride{
//   		&containerOverride{
//   			ContainerDefinition: *ContainerDefinition,
//   			Environment: []taskEnvironmentVariable{
//   				&taskEnvironmentVariable{
//   					Name: jsii.String("SOME_KEY"),
//   					Value: sfn.JsonPath_StringAt(jsii.String("$.SomeKey")),
//   				},
//   			},
//   		},
//   	},
//   	LaunchTarget: tasks.NewEcsFargateLaunchTarget(),
//   })
//
type Cluster interface {
	awscdk.Resource
	ICluster
	// Getter for autoscaling group added to cluster.
	AutoscalingGroup() awsautoscaling.IAutoScalingGroup
	// Getter for _capacityProviderNames added to cluster.
	CapacityProviderNames() *[]*string
	// The Amazon Resource Name (ARN) that identifies the cluster.
	ClusterArn() *string
	// The name of the cluster.
	ClusterName() *string
	// Manage the allowed network connections for the cluster with Security Groups.
	Connections() awsec2.Connections
	// Getter for _defaultCapacityProviderStrategy.
	//
	// This is necessary to correctly create Capacity Provider Associations.
	DefaultCapacityProviderStrategy() *[]*CapacityProviderStrategy
	// Getter for namespace added to cluster.
	DefaultCloudMapNamespace() awsservicediscovery.INamespace
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// Getter for execute command configuration associated with the cluster.
	ExecuteCommandConfiguration() *ExecuteCommandConfiguration
	// Whether the cluster has EC2 capacity associated with it.
	HasEc2Capacity() *bool
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
	// The VPC associated with the cluster.
	Vpc() awsec2.IVpc
	// This method adds an Auto Scaling Group Capacity Provider to a cluster.
	AddAsgCapacityProvider(provider AsgCapacityProvider, options *AddAutoScalingGroupCapacityOptions)
	// It is highly recommended to use `Cluster.addAsgCapacityProvider` instead of this method.
	//
	// This method adds compute capacity to a cluster by creating an AutoScalingGroup with the specified options.
	//
	// Returns the AutoScalingGroup so you can add autoscaling settings to it.
	AddCapacity(id *string, options *AddCapacityOptions) awsautoscaling.AutoScalingGroup
	// Add default capacity provider strategy for this cluster.
	AddDefaultCapacityProviderStrategy(defaultCapacityProviderStrategy *[]*CapacityProviderStrategy)
	// Add an AWS Cloud Map DNS namespace for this cluster.
	//
	// NOTE: HttpNamespaces are supported only for use cases involving Service Connect. For use cases involving both Service-
	// Discovery and Service Connect, customers should manage the HttpNamespace outside of the Cluster.addDefaultCloudMapNamespace method.
	AddDefaultCloudMapNamespace(options *CloudMapNamespaceOptions) awsservicediscovery.INamespace
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
	// Enable the Fargate capacity providers for this cluster.
	EnableFargateCapacityProviders()
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
	// This method returns the specifed CloudWatch metric for this cluster.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// This method returns the CloudWatch metric for this clusters CPU reservation.
	MetricCpuReservation(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// This method returns the CloudWatch metric for this clusters CPU utilization.
	MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// This method returns the CloudWatch metric for this clusters memory reservation.
	MetricMemoryReservation(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// This method returns the CloudWatch metric for this clusters memory utilization.
	MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Cluster
type jsiiProxy_Cluster struct {
	internal.Type__awscdkResource
	jsiiProxy_ICluster
}

func (j *jsiiProxy_Cluster) AutoscalingGroup() awsautoscaling.IAutoScalingGroup {
	var returns awsautoscaling.IAutoScalingGroup
	_jsii_.Get(
		j,
		"autoscalingGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) CapacityProviderNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"capacityProviderNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DefaultCapacityProviderStrategy() *[]*CapacityProviderStrategy {
	var returns *[]*CapacityProviderStrategy
	_jsii_.Get(
		j,
		"defaultCapacityProviderStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) DefaultCloudMapNamespace() awsservicediscovery.INamespace {
	var returns awsservicediscovery.INamespace
	_jsii_.Get(
		j,
		"defaultCloudMapNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ExecuteCommandConfiguration() *ExecuteCommandConfiguration {
	var returns *ExecuteCommandConfiguration
	_jsii_.Get(
		j,
		"executeCommandConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) HasEc2Capacity() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"hasEc2Capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Constructs a new instance of the Cluster class.
func NewCluster(scope constructs.Construct, id *string, props *ClusterProps) Cluster {
	_init_.Initialize()

	if err := validateNewClusterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Cluster{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.Cluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the Cluster class.
func NewCluster_Override(c Cluster, scope constructs.Construct, id *string, props *ClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.Cluster",
		[]interface{}{scope, id, props},
		c,
	)
}

// Import an existing cluster to the stack from the cluster ARN.
//
// This does not provide access to the vpc, hasEc2Capacity, or connections -
// use the `fromClusterAttributes` method to access those properties.
func Cluster_FromClusterArn(scope constructs.Construct, id *string, clusterArn *string) ICluster {
	_init_.Initialize()

	if err := validateCluster_FromClusterArnParameters(scope, id, clusterArn); err != nil {
		panic(err)
	}
	var returns ICluster

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Cluster",
		"fromClusterArn",
		[]interface{}{scope, id, clusterArn},
		&returns,
	)

	return returns
}

// Import an existing cluster to the stack from its attributes.
func Cluster_FromClusterAttributes(scope constructs.Construct, id *string, attrs *ClusterAttributes) ICluster {
	_init_.Initialize()

	if err := validateCluster_FromClusterAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns ICluster

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Cluster",
		"fromClusterAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Cluster.
func Cluster_IsCluster(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCluster_IsClusterParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Cluster",
		"isCluster",
		[]interface{}{x},
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
func Cluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Cluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Cluster_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCluster_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Cluster",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Cluster_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCluster_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Cluster",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) AddAsgCapacityProvider(provider AsgCapacityProvider, options *AddAutoScalingGroupCapacityOptions) {
	if err := c.validateAddAsgCapacityProviderParameters(provider, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addAsgCapacityProvider",
		[]interface{}{provider, options},
	)
}

func (c *jsiiProxy_Cluster) AddCapacity(id *string, options *AddCapacityOptions) awsautoscaling.AutoScalingGroup {
	if err := c.validateAddCapacityParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsautoscaling.AutoScalingGroup

	_jsii_.Invoke(
		c,
		"addCapacity",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) AddDefaultCapacityProviderStrategy(defaultCapacityProviderStrategy *[]*CapacityProviderStrategy) {
	if err := c.validateAddDefaultCapacityProviderStrategyParameters(defaultCapacityProviderStrategy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDefaultCapacityProviderStrategy",
		[]interface{}{defaultCapacityProviderStrategy},
	)
}

func (c *jsiiProxy_Cluster) AddDefaultCloudMapNamespace(options *CloudMapNamespaceOptions) awsservicediscovery.INamespace {
	if err := c.validateAddDefaultCloudMapNamespaceParameters(options); err != nil {
		panic(err)
	}
	var returns awsservicediscovery.INamespace

	_jsii_.Invoke(
		c,
		"addDefaultCloudMapNamespace",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := c.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_Cluster) EnableFargateCapacityProviders() {
	_jsii_.InvokeVoid(
		c,
		"enableFargateCapacityProviders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cluster) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := c.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) GetResourceNameAttribute(nameAttr *string) *string {
	if err := c.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := c.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) MetricCpuReservation(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := c.validateMetricCpuReservationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricCpuReservation",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := c.validateMetricCpuUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricCpuUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) MetricMemoryReservation(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := c.validateMetricMemoryReservationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricMemoryReservation",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := c.validateMetricMemoryUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricMemoryUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

