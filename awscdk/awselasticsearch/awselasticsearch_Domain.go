package awselasticsearch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awselasticsearch/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/constructs-go/constructs/v3"
)

// Provides an Elasticsearch domain.
//
// Example:
//   domain := es.NewDomain(this, jsii.String("Domain"), &domainProps{
//   	version: es.elasticsearchVersion_V7_4(),
//   	ebs: &ebsOptions{
//   		volumeSize: jsii.Number(100),
//   		volumeType: ec2.ebsDeviceVolumeType_GENERAL_PURPOSE_SSD,
//   	},
//   	nodeToNodeEncryption: jsii.Boolean(true),
//   	encryptionAtRest: &encryptionAtRestOptions{
//   		enabled: jsii.Boolean(true),
//   	},
//   })
//
// Deprecated: use opensearchservice module instead.
type Domain interface {
	awscdk.Resource
	awsec2.IConnectable
	IDomain
	// Log group that application logs are logged to.
	// Deprecated: use opensearchservice module instead.
	AppLogGroup() awslogs.ILogGroup
	// Log group that audit logs are logged to.
	// Deprecated: use opensearchservice module instead.
	AuditLogGroup() awslogs.ILogGroup
	// Manages network connections to the domain.
	//
	// This will throw an error in case the domain
	// is not placed inside a VPC.
	// Deprecated: use opensearchservice module instead.
	Connections() awsec2.Connections
	// Arn of the Elasticsearch domain.
	// Deprecated: use opensearchservice module instead.
	DomainArn() *string
	// Endpoint of the Elasticsearch domain.
	// Deprecated: use opensearchservice module instead.
	DomainEndpoint() *string
	// Domain name of the Elasticsearch domain.
	// Deprecated: use opensearchservice module instead.
	DomainName() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Deprecated: use opensearchservice module instead.
	Env() *awscdk.ResourceEnvironment
	// Master user password if fine grained access control is configured.
	// Deprecated: use opensearchservice module instead.
	MasterUserPassword() awscdk.SecretValue
	// The construct tree node associated with this construct.
	// Deprecated: use opensearchservice module instead.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Deprecated: use opensearchservice module instead.
	PhysicalName() *string
	// Log group that slow indices are logged to.
	// Deprecated: use opensearchservice module instead.
	SlowIndexLogGroup() awslogs.ILogGroup
	// Log group that slow searches are logged to.
	// Deprecated: use opensearchservice module instead.
	SlowSearchLogGroup() awslogs.ILogGroup
	// The stack in which this resource is defined.
	// Deprecated: use opensearchservice module instead.
	Stack() awscdk.Stack
	// Add policy statements to the domain access policy.
	// Deprecated: use opensearchservice module instead.
	AddAccessPolicies(accessPolicyStatements ...awsiam.PolicyStatement)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Deprecated: use opensearchservice module instead.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Deprecated: use opensearchservice module instead.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Deprecated: use opensearchservice module instead.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Deprecated: use opensearchservice module instead.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grant read permissions for an index in this domain to an IAM principal (Role/Group/User).
	// Deprecated: use opensearchservice module instead.
	GrantIndexRead(index *string, identity awsiam.IGrantable) awsiam.Grant
	// Grant read/write permissions for an index in this domain to an IAM principal (Role/Group/User).
	// Deprecated: use opensearchservice module instead.
	GrantIndexReadWrite(index *string, identity awsiam.IGrantable) awsiam.Grant
	// Grant write permissions for an index in this domain to an IAM principal (Role/Group/User).
	// Deprecated: use opensearchservice module instead.
	GrantIndexWrite(index *string, identity awsiam.IGrantable) awsiam.Grant
	// Grant read permissions for a specific path in this domain to an IAM principal (Role/Group/User).
	// Deprecated: use opensearchservice module instead.
	GrantPathRead(path *string, identity awsiam.IGrantable) awsiam.Grant
	// Grant read/write permissions for a specific path in this domain to an IAM principal (Role/Group/User).
	// Deprecated: use opensearchservice module instead.
	GrantPathReadWrite(path *string, identity awsiam.IGrantable) awsiam.Grant
	// Grant write permissions for a specific path in this domain to an IAM principal (Role/Group/User).
	// Deprecated: use opensearchservice module instead.
	GrantPathWrite(path *string, identity awsiam.IGrantable) awsiam.Grant
	// Grant read permissions for this domain and its contents to an IAM principal (Role/Group/User).
	// Deprecated: use opensearchservice module instead.
	GrantRead(identity awsiam.IGrantable) awsiam.Grant
	// Grant read/write permissions for this domain and its contents to an IAM principal (Role/Group/User).
	// Deprecated: use opensearchservice module instead.
	GrantReadWrite(identity awsiam.IGrantable) awsiam.Grant
	// Grant write permissions for this domain and its contents to an IAM principal (Role/Group/User).
	// Deprecated: use opensearchservice module instead.
	GrantWrite(identity awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this Domain.
	// Deprecated: use opensearchservice module instead.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for automated snapshot failures.
	// Deprecated: use opensearchservice module instead.
	MetricAutomatedSnapshotFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the cluster blocking index writes.
	// Deprecated: use opensearchservice module instead.
	MetricClusterIndexWritesBlocked(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the time the cluster status is red.
	// Deprecated: use opensearchservice module instead.
	MetricClusterStatusRed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the time the cluster status is yellow.
	// Deprecated: use opensearchservice module instead.
	MetricClusterStatusYellow(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for CPU utilization.
	// Deprecated: use opensearchservice module instead.
	MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the storage space of nodes in the cluster.
	// Deprecated: use opensearchservice module instead.
	MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for indexing latency.
	// Deprecated: use opensearchservice module instead.
	MetricIndexingLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for JVM memory pressure.
	// Deprecated: use opensearchservice module instead.
	MetricJVMMemoryPressure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for KMS key errors.
	// Deprecated: use opensearchservice module instead.
	MetricKMSKeyError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for KMS key being inaccessible.
	// Deprecated: use opensearchservice module instead.
	MetricKMSKeyInaccessible(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for master CPU utilization.
	// Deprecated: use opensearchservice module instead.
	MetricMasterCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for master JVM memory pressure.
	// Deprecated: use opensearchservice module instead.
	MetricMasterJVMMemoryPressure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of nodes.
	// Deprecated: use opensearchservice module instead.
	MetricNodes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for number of searchable documents.
	// Deprecated: use opensearchservice module instead.
	MetricSearchableDocuments(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for search latency.
	// Deprecated: use opensearchservice module instead.
	MetricSearchLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: use opensearchservice module instead.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: use opensearchservice module instead.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Deprecated: use opensearchservice module instead.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: use opensearchservice module instead.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: use opensearchservice module instead.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Deprecated: use opensearchservice module instead.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Deprecated: use opensearchservice module instead.
	Validate() *[]*string
}

// The jsii proxy struct for Domain
type jsiiProxy_Domain struct {
	internal.Type__awscdkResource
	internal.Type__awsec2IConnectable
	jsiiProxy_IDomain
}

func (j *jsiiProxy_Domain) AppLogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"appLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) AuditLogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"auditLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) DomainArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) DomainEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) MasterUserPassword() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"masterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) SlowIndexLogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"slowIndexLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) SlowSearchLogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"slowSearchLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Domain) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Deprecated: use opensearchservice module instead.
func NewDomain(scope constructs.Construct, id *string, props *DomainProps) Domain {
	_init_.Initialize()

	if err := validateNewDomainParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Domain{}

	_jsii_.Create(
		"monocdk.aws_elasticsearch.Domain",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Deprecated: use opensearchservice module instead.
func NewDomain_Override(d Domain, scope constructs.Construct, id *string, props *DomainProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_elasticsearch.Domain",
		[]interface{}{scope, id, props},
		d,
	)
}

// Creates a Domain construct that represents an external domain.
// Deprecated: use opensearchservice module instead.
func Domain_FromDomainAttributes(scope constructs.Construct, id *string, attrs *DomainAttributes) IDomain {
	_init_.Initialize()

	if err := validateDomain_FromDomainAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IDomain

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticsearch.Domain",
		"fromDomainAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Creates a Domain construct that represents an external domain via domain endpoint.
// Deprecated: use opensearchservice module instead.
func Domain_FromDomainEndpoint(scope constructs.Construct, id *string, domainEndpoint *string) IDomain {
	_init_.Initialize()

	if err := validateDomain_FromDomainEndpointParameters(scope, id, domainEndpoint); err != nil {
		panic(err)
	}
	var returns IDomain

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticsearch.Domain",
		"fromDomainEndpoint",
		[]interface{}{scope, id, domainEndpoint},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Deprecated: use opensearchservice module instead.
func Domain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDomain_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticsearch.Domain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Deprecated: use opensearchservice module instead.
func Domain_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	if err := validateDomain_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticsearch.Domain",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) AddAccessPolicies(accessPolicyStatements ...awsiam.PolicyStatement) {
	args := []interface{}{}
	for _, a := range accessPolicyStatements {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addAccessPolicies",
		args,
	)
}

func (d *jsiiProxy_Domain) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := d.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (d *jsiiProxy_Domain) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := d.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) GetResourceNameAttribute(nameAttr *string) *string {
	if err := d.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) GrantIndexRead(index *string, identity awsiam.IGrantable) awsiam.Grant {
	if err := d.validateGrantIndexReadParameters(index, identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantIndexRead",
		[]interface{}{index, identity},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) GrantIndexReadWrite(index *string, identity awsiam.IGrantable) awsiam.Grant {
	if err := d.validateGrantIndexReadWriteParameters(index, identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantIndexReadWrite",
		[]interface{}{index, identity},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) GrantIndexWrite(index *string, identity awsiam.IGrantable) awsiam.Grant {
	if err := d.validateGrantIndexWriteParameters(index, identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantIndexWrite",
		[]interface{}{index, identity},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) GrantPathRead(path *string, identity awsiam.IGrantable) awsiam.Grant {
	if err := d.validateGrantPathReadParameters(path, identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantPathRead",
		[]interface{}{path, identity},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) GrantPathReadWrite(path *string, identity awsiam.IGrantable) awsiam.Grant {
	if err := d.validateGrantPathReadWriteParameters(path, identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantPathReadWrite",
		[]interface{}{path, identity},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) GrantPathWrite(path *string, identity awsiam.IGrantable) awsiam.Grant {
	if err := d.validateGrantPathWriteParameters(path, identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantPathWrite",
		[]interface{}{path, identity},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) GrantRead(identity awsiam.IGrantable) awsiam.Grant {
	if err := d.validateGrantReadParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantRead",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) GrantReadWrite(identity awsiam.IGrantable) awsiam.Grant {
	if err := d.validateGrantReadWriteParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantReadWrite",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) GrantWrite(identity awsiam.IGrantable) awsiam.Grant {
	if err := d.validateGrantWriteParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantWrite",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) MetricAutomatedSnapshotFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricAutomatedSnapshotFailureParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricAutomatedSnapshotFailure",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) MetricClusterIndexWritesBlocked(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricClusterIndexWritesBlockedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricClusterIndexWritesBlocked",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) MetricClusterStatusRed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricClusterStatusRedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricClusterStatusRed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) MetricClusterStatusYellow(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricClusterStatusYellowParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricClusterStatusYellow",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricCPUUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricCPUUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricFreeStorageSpaceParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFreeStorageSpace",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) MetricIndexingLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricIndexingLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricIndexingLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) MetricJVMMemoryPressure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricJVMMemoryPressureParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricJVMMemoryPressure",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) MetricKMSKeyError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricKMSKeyErrorParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricKMSKeyError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) MetricKMSKeyInaccessible(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricKMSKeyInaccessibleParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricKMSKeyInaccessible",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) MetricMasterCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricMasterCPUUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricMasterCPUUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) MetricMasterJVMMemoryPressure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricMasterJVMMemoryPressureParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricMasterJVMMemoryPressure",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) MetricNodes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricNodesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricNodes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) MetricSearchableDocuments(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricSearchableDocumentsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricSearchableDocuments",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) MetricSearchLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricSearchLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricSearchLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Domain) OnSynthesize(session constructs.ISynthesisSession) {
	if err := d.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (d *jsiiProxy_Domain) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) Prepare() {
	_jsii_.InvokeVoid(
		d,
		"prepare",
		nil, // no parameters
	)
}

func (d *jsiiProxy_Domain) Synthesize(session awscdk.ISynthesisSession) {
	if err := d.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		[]interface{}{session},
	)
}

func (d *jsiiProxy_Domain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_Domain) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

