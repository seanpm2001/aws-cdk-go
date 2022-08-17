// The CDK Construct Library for AWS::MSK
package awscdkmskalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmskalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsacmpca"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdkmskalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Configuration details related to broker logs.
//
// Example:
//   var vpc vpc
//   var bucket iBucket
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &clusterProps{
//   	clusterName: jsii.String("myCluster"),
//   	kafkaVersion: msk.kafkaVersion_V2_8_1(),
//   	vpc: vpc,
//   	logging: &brokerLogging{
//   		s3: &s3LoggingConfiguration{
//   			bucket: bucket,
//   		},
//   	},
//   })
//
// Experimental.
type BrokerLogging struct {
	// The CloudWatch Logs group that is the destination for broker logs.
	// Experimental.
	CloudwatchLogGroup awslogs.ILogGroup `field:"optional" json:"cloudwatchLogGroup" yaml:"cloudwatchLogGroup"`
	// The Kinesis Data Firehose delivery stream that is the destination for broker logs.
	// Experimental.
	FirehoseDeliveryStreamName *string `field:"optional" json:"firehoseDeliveryStreamName" yaml:"firehoseDeliveryStreamName"`
	// Details of the Amazon S3 destination for broker logs.
	// Experimental.
	S3 *S3LoggingConfiguration `field:"optional" json:"s3" yaml:"s3"`
}

// Configuration properties for client authentication.
//
// Example:
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &clusterProps{
//   	clusterName: jsii.String("myCluster"),
//   	kafkaVersion: msk.kafkaVersion_V2_8_1(),
//   	vpc: vpc,
//   	encryptionInTransit: &encryptionInTransitConfig{
//   		clientBroker: msk.clientBrokerEncryption_TLS,
//   	},
//   	clientAuthentication: msk.clientAuthentication.sasl(&saslAuthProps{
//   		scram: jsii.Boolean(true),
//   	}),
//   })
//
// Experimental.
type ClientAuthentication interface {
	// - properties for SASL authentication.
	// Experimental.
	SaslProps() *SaslAuthProps
	// - properties for TLS authentication.
	// Experimental.
	TlsProps() *TlsAuthProps
}

// The jsii proxy struct for ClientAuthentication
type jsiiProxy_ClientAuthentication struct {
	_ byte // padding
}

func (j *jsiiProxy_ClientAuthentication) SaslProps() *SaslAuthProps {
	var returns *SaslAuthProps
	_jsii_.Get(
		j,
		"saslProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClientAuthentication) TlsProps() *TlsAuthProps {
	var returns *TlsAuthProps
	_jsii_.Get(
		j,
		"tlsProps",
		&returns,
	)
	return returns
}


// SASL authentication.
// Experimental.
func ClientAuthentication_Sasl(props *SaslAuthProps) ClientAuthentication {
	_init_.Initialize()

	var returns ClientAuthentication

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-msk-alpha.ClientAuthentication",
		"sasl",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// TLS authentication.
// Experimental.
func ClientAuthentication_Tls(props *TlsAuthProps) ClientAuthentication {
	_init_.Initialize()

	var returns ClientAuthentication

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-msk-alpha.ClientAuthentication",
		"tls",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Indicates the encryption setting for data in transit between clients and brokers.
//
// Example:
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &clusterProps{
//   	clusterName: jsii.String("myCluster"),
//   	kafkaVersion: msk.kafkaVersion_V2_8_1(),
//   	vpc: vpc,
//   	encryptionInTransit: &encryptionInTransitConfig{
//   		clientBroker: msk.clientBrokerEncryption_TLS,
//   	},
//   	clientAuthentication: msk.clientAuthentication.sasl(&saslAuthProps{
//   		scram: jsii.Boolean(true),
//   	}),
//   })
//
// Experimental.
type ClientBrokerEncryption string

const (
	// TLS means that client-broker communication is enabled with TLS only.
	// Experimental.
	ClientBrokerEncryption_TLS ClientBrokerEncryption = "TLS"
	// TLS_PLAINTEXT means that client-broker communication is enabled for both TLS-encrypted, as well as plaintext data.
	// Experimental.
	ClientBrokerEncryption_TLS_PLAINTEXT ClientBrokerEncryption = "TLS_PLAINTEXT"
	// PLAINTEXT means that client-broker communication is enabled in plaintext only.
	// Experimental.
	ClientBrokerEncryption_PLAINTEXT ClientBrokerEncryption = "PLAINTEXT"
)

// Create a MSK Cluster.
//
// Example:
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &clusterProps{
//   	clusterName: jsii.String("myCluster"),
//   	kafkaVersion: msk.kafkaVersion_V2_8_1(),
//   	vpc: vpc,
//   	encryptionInTransit: &encryptionInTransitConfig{
//   		clientBroker: msk.clientBrokerEncryption_TLS,
//   	},
//   	clientAuthentication: msk.clientAuthentication.sasl(&saslAuthProps{
//   		scram: jsii.Boolean(true),
//   	}),
//   })
//
// Experimental.
type Cluster interface {
	awscdk.Resource
	ICluster
	// Get the list of brokers that a client application can use to bootstrap.
	//
	// Uses a Custom Resource to make an API call to `getBootstrapBrokers` using the Javascript SDK.
	//
	// Returns: - A string containing one or more hostname:port pairs.
	// Experimental.
	BootstrapBrokers() *string
	// Get the list of brokers that a SASL/SCRAM authenticated client application can use to bootstrap.
	//
	// Uses a Custom Resource to make an API call to `getBootstrapBrokers` using the Javascript SDK.
	//
	// Returns: - A string containing one or more dns name (or IP) and SASL SCRAM port pairs.
	// Experimental.
	BootstrapBrokersSaslScram() *string
	// Get the list of brokers that a TLS authenticated client application can use to bootstrap.
	//
	// Uses a Custom Resource to make an API call to `getBootstrapBrokers` using the Javascript SDK.
	//
	// Returns: - A string containing one or more DNS names (or IP) and TLS port pairs.
	// Experimental.
	BootstrapBrokersTls() *string
	// The ARN of cluster.
	// Experimental.
	ClusterArn() *string
	// The physical name of the cluster.
	// Experimental.
	ClusterName() *string
	// Manages connections for the cluster.
	// Experimental.
	Connections() awsec2.Connections
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
	// Key used to encrypt SASL/SCRAM users.
	// Experimental.
	SaslScramAuthenticationKey() awskms.IKey
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Get the ZooKeeper Connection string.
	//
	// Uses a Custom Resource to make an API call to `describeCluster` using the Javascript SDK.
	//
	// Returns: - The connection string to use to connect to the Apache ZooKeeper cluster.
	// Experimental.
	ZookeeperConnectionString() *string
	// Get the ZooKeeper Connection string for a TLS enabled cluster.
	//
	// Uses a Custom Resource to make an API call to `describeCluster` using the Javascript SDK.
	//
	// Returns: - The connection string to use to connect to zookeeper cluster on TLS port.
	// Experimental.
	ZookeeperConnectionStringTls() *string
	// A list of usersnames to register with the cluster.
	//
	// The password will automatically be generated using Secrets
	// Manager and the { username, password } JSON object stored in Secrets Manager as `AmazonMSK_username`.
	//
	// Must be using the SASL/SCRAM authentication mechanism.
	// Experimental.
	AddUser(usernames ...*string)
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

// The jsii proxy struct for Cluster
type jsiiProxy_Cluster struct {
	internal.Type__awscdkResource
	jsiiProxy_ICluster
}

func (j *jsiiProxy_Cluster) BootstrapBrokers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) BootstrapBrokersSaslScram() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersSaslScram",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) BootstrapBrokersTls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersTls",
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

func (j *jsiiProxy_Cluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
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

func (j *jsiiProxy_Cluster) SaslScramAuthenticationKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"saslScramAuthenticationKey",
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

func (j *jsiiProxy_Cluster) ZookeeperConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zookeeperConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cluster) ZookeeperConnectionStringTls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zookeeperConnectionStringTls",
		&returns,
	)
	return returns
}


// Experimental.
func NewCluster(scope constructs.Construct, id *string, props *ClusterProps) Cluster {
	_init_.Initialize()

	j := jsiiProxy_Cluster{}

	_jsii_.Create(
		"@aws-cdk/aws-msk-alpha.Cluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCluster_Override(c Cluster, scope constructs.Construct, id *string, props *ClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-msk-alpha.Cluster",
		[]interface{}{scope, id, props},
		c,
	)
}

// Reference an existing cluster, defined outside of the CDK code, by name.
// Experimental.
func Cluster_FromClusterArn(scope constructs.Construct, id *string, clusterArn *string) ICluster {
	_init_.Initialize()

	var returns ICluster

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-msk-alpha.Cluster",
		"fromClusterArn",
		[]interface{}{scope, id, clusterArn},
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
func Cluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-msk-alpha.Cluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func Cluster_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-msk-alpha.Cluster",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Cluster_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-msk-alpha.Cluster",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cluster) AddUser(usernames ...*string) {
	args := []interface{}{}
	for _, a := range usernames {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addUser",
		args,
	)
}

func (c *jsiiProxy_Cluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
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
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
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

// The Amazon MSK configuration to use for the cluster.
//
// Note: There is currently no Cloudformation Resource to create a Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import msk_alpha "github.com/aws/aws-cdk-go/awscdkmskalpha"
//
//   clusterConfigurationInfo := &clusterConfigurationInfo{
//   	arn: jsii.String("arn"),
//   	revision: jsii.Number(123),
//   }
//
// Experimental.
type ClusterConfigurationInfo struct {
	// The Amazon Resource Name (ARN) of the MSK configuration to use.
	//
	// For example, arn:aws:kafka:us-east-1:123456789012:configuration/example-configuration-name/abcdabcd-1234-abcd-1234-abcd123e8e8e-1.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The revision of the Amazon MSK configuration to use.
	// Experimental.
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
}

// The level of monitoring for the MSK cluster.
// See: https://docs.aws.amazon.com/msk/latest/developerguide/monitoring.html#metrics-details
//
// Experimental.
type ClusterMonitoringLevel string

const (
	// Default metrics are the essential metrics to monitor.
	// Experimental.
	ClusterMonitoringLevel_DEFAULT ClusterMonitoringLevel = "DEFAULT"
	// Per Broker metrics give you metrics at the broker level.
	// Experimental.
	ClusterMonitoringLevel_PER_BROKER ClusterMonitoringLevel = "PER_BROKER"
	// Per Topic Per Broker metrics help you understand volume at the topic level.
	// Experimental.
	ClusterMonitoringLevel_PER_TOPIC_PER_BROKER ClusterMonitoringLevel = "PER_TOPIC_PER_BROKER"
	// Per Topic Per Partition metrics help you understand consumer group lag at the topic partition level.
	// Experimental.
	ClusterMonitoringLevel_PER_TOPIC_PER_PARTITION ClusterMonitoringLevel = "PER_TOPIC_PER_PARTITION"
)

// Properties for a MSK Cluster.
//
// Example:
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &clusterProps{
//   	clusterName: jsii.String("myCluster"),
//   	kafkaVersion: msk.kafkaVersion_V2_8_1(),
//   	vpc: vpc,
//   	encryptionInTransit: &encryptionInTransitConfig{
//   		clientBroker: msk.clientBrokerEncryption_TLS,
//   	},
//   	clientAuthentication: msk.clientAuthentication.sasl(&saslAuthProps{
//   		scram: jsii.Boolean(true),
//   	}),
//   })
//
// Experimental.
type ClusterProps struct {
	// The physical name of the cluster.
	// Experimental.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The version of Apache Kafka.
	// Experimental.
	KafkaVersion KafkaVersion `field:"required" json:"kafkaVersion" yaml:"kafkaVersion"`
	// Defines the virtual networking environment for this cluster.
	//
	// Must have at least 2 subnets in two different AZs.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Configuration properties for client authentication.
	//
	// MSK supports using private TLS certificates or SASL/SCRAM to authenticate the identity of clients.
	// Experimental.
	ClientAuthentication ClientAuthentication `field:"optional" json:"clientAuthentication" yaml:"clientAuthentication"`
	// The Amazon MSK configuration to use for the cluster.
	// Experimental.
	ConfigurationInfo *ClusterConfigurationInfo `field:"optional" json:"configurationInfo" yaml:"configurationInfo"`
	// Information about storage volumes attached to MSK broker nodes.
	// Experimental.
	EbsStorageInfo *EbsStorageInfo `field:"optional" json:"ebsStorageInfo" yaml:"ebsStorageInfo"`
	// Config details for encryption in transit.
	// Experimental.
	EncryptionInTransit *EncryptionInTransitConfig `field:"optional" json:"encryptionInTransit" yaml:"encryptionInTransit"`
	// The EC2 instance type that you want Amazon MSK to use when it creates your brokers.
	// See: https://docs.aws.amazon.com/msk/latest/developerguide/msk-create-cluster.html#broker-instance-types
	//
	// Experimental.
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Configure your MSK cluster to send broker logs to different destination types.
	// Experimental.
	Logging *BrokerLogging `field:"optional" json:"logging" yaml:"logging"`
	// Cluster monitoring configuration.
	// Experimental.
	Monitoring *MonitoringConfiguration `field:"optional" json:"monitoring" yaml:"monitoring"`
	// Number of Apache Kafka brokers deployed in each Availability Zone.
	// Experimental.
	NumberOfBrokerNodes *float64 `field:"optional" json:"numberOfBrokerNodes" yaml:"numberOfBrokerNodes"`
	// What to do when this resource is deleted from a stack.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The AWS security groups to associate with the elastic network interfaces in order to specify who can connect to and communicate with the Amazon MSK cluster.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Where to place the nodes within the VPC.
	//
	// Amazon MSK distributes the broker nodes evenly across the subnets that you specify.
	// The subnets that you specify must be in distinct Availability Zones.
	// Client subnets can't be in Availability Zone us-east-1e.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

// EBS volume information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import msk_alpha "github.com/aws/aws-cdk-go/awscdkmskalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var key key
//
//   ebsStorageInfo := &ebsStorageInfo{
//   	encryptionKey: key,
//   	volumeSize: jsii.Number(123),
//   }
//
// Experimental.
type EbsStorageInfo struct {
	// The AWS KMS key for encrypting data at rest.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The size in GiB of the EBS volume for the data drive on each broker node.
	// Experimental.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
}

// The settings for encrypting data in transit.
//
// Example:
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &clusterProps{
//   	clusterName: jsii.String("myCluster"),
//   	kafkaVersion: msk.kafkaVersion_V2_8_1(),
//   	vpc: vpc,
//   	encryptionInTransit: &encryptionInTransitConfig{
//   		clientBroker: msk.clientBrokerEncryption_TLS,
//   	},
//   	clientAuthentication: msk.clientAuthentication.sasl(&saslAuthProps{
//   		scram: jsii.Boolean(true),
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/msk/latest/developerguide/msk-encryption.html#msk-encryption-in-transit
//
// Experimental.
type EncryptionInTransitConfig struct {
	// Indicates the encryption setting for data in transit between clients and brokers.
	// Experimental.
	ClientBroker ClientBrokerEncryption `field:"optional" json:"clientBroker" yaml:"clientBroker"`
	// Indicates that data communication among the broker nodes of the cluster is encrypted.
	// Experimental.
	EnableInCluster *bool `field:"optional" json:"enableInCluster" yaml:"enableInCluster"`
}

// Represents a MSK Cluster.
// Experimental.
type ICluster interface {
	awsec2.IConnectable
	awscdk.IResource
	// The ARN of cluster.
	// Experimental.
	ClusterArn() *string
	// The physical name of the cluster.
	// Experimental.
	ClusterName() *string
}

// The jsii proxy for ICluster
type jsiiProxy_ICluster struct {
	internal.Type__awsec2IConnectable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ICluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_ICluster) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// Kafka cluster version.
//
// Example:
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &clusterProps{
//   	clusterName: jsii.String("myCluster"),
//   	kafkaVersion: msk.kafkaVersion_V2_8_1(),
//   	vpc: vpc,
//   	encryptionInTransit: &encryptionInTransitConfig{
//   		clientBroker: msk.clientBrokerEncryption_TLS,
//   	},
//   	clientAuthentication: msk.clientAuthentication.sasl(&saslAuthProps{
//   		scram: jsii.Boolean(true),
//   	}),
//   })
//
// Experimental.
type KafkaVersion interface {
	// cluster version number.
	// Experimental.
	Version() *string
}

// The jsii proxy struct for KafkaVersion
type jsiiProxy_KafkaVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_KafkaVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Custom cluster version.
// Experimental.
func KafkaVersion_Of(version *string) KafkaVersion {
	_init_.Initialize()

	var returns KafkaVersion

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"of",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func KafkaVersion_V1_1_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V1_1_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_2_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_2_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_3_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_3_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_4_1_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_4_1_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_5_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_5_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_6_0() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_6_0",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_6_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_6_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_6_2() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_6_2",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_6_3() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_6_3",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_7_0() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_7_0",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_7_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_7_1",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_7_2() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_7_2",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_8_0() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_8_0",
		&returns,
	)
	return returns
}

func KafkaVersion_V2_8_1() KafkaVersion {
	_init_.Initialize()
	var returns KafkaVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-msk-alpha.KafkaVersion",
		"V2_8_1",
		&returns,
	)
	return returns
}

// Monitoring Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import msk_alpha "github.com/aws/aws-cdk-go/awscdkmskalpha"
//
//   monitoringConfiguration := &monitoringConfiguration{
//   	clusterMonitoringLevel: msk_alpha.clusterMonitoringLevel_DEFAULT,
//   	enablePrometheusJmxExporter: jsii.Boolean(false),
//   	enablePrometheusNodeExporter: jsii.Boolean(false),
//   }
//
// Experimental.
type MonitoringConfiguration struct {
	// Specifies the level of monitoring for the MSK cluster.
	// Experimental.
	ClusterMonitoringLevel ClusterMonitoringLevel `field:"optional" json:"clusterMonitoringLevel" yaml:"clusterMonitoringLevel"`
	// Indicates whether you want to enable or disable the JMX Exporter.
	// Experimental.
	EnablePrometheusJmxExporter *bool `field:"optional" json:"enablePrometheusJmxExporter" yaml:"enablePrometheusJmxExporter"`
	// Indicates whether you want to enable or disable the Prometheus Node Exporter.
	//
	// You can use the Prometheus Node Exporter to get CPU and disk metrics for the broker nodes.
	// Experimental.
	EnablePrometheusNodeExporter *bool `field:"optional" json:"enablePrometheusNodeExporter" yaml:"enablePrometheusNodeExporter"`
}

// Details of the Amazon S3 destination for broker logs.
//
// Example:
//   var vpc vpc
//   var bucket iBucket
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &clusterProps{
//   	clusterName: jsii.String("myCluster"),
//   	kafkaVersion: msk.kafkaVersion_V2_8_1(),
//   	vpc: vpc,
//   	logging: &brokerLogging{
//   		s3: &s3LoggingConfiguration{
//   			bucket: bucket,
//   		},
//   	},
//   })
//
// Experimental.
type S3LoggingConfiguration struct {
	// The S3 bucket that is the destination for broker logs.
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The S3 prefix that is the destination for broker logs.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

// SASL authentication properties.
//
// Example:
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &clusterProps{
//   	clusterName: jsii.String("myCluster"),
//   	kafkaVersion: msk.kafkaVersion_V2_8_1(),
//   	vpc: vpc,
//   	encryptionInTransit: &encryptionInTransitConfig{
//   		clientBroker: msk.clientBrokerEncryption_TLS,
//   	},
//   	clientAuthentication: msk.clientAuthentication.sasl(&saslAuthProps{
//   		scram: jsii.Boolean(true),
//   	}),
//   })
//
// Experimental.
type SaslAuthProps struct {
	// Enable IAM access control.
	// Experimental.
	Iam *bool `field:"optional" json:"iam" yaml:"iam"`
	// KMS Key to encrypt SASL/SCRAM secrets.
	//
	// You must use a customer master key (CMK) when creating users in secrets manager.
	// You cannot use a Secret with Amazon MSK that uses the default Secrets Manager encryption key.
	// Experimental.
	Key awskms.IKey `field:"optional" json:"key" yaml:"key"`
	// Enable SASL/SCRAM authentication.
	// Experimental.
	Scram *bool `field:"optional" json:"scram" yaml:"scram"`
}

// TLS authentication properties.
//
// Example:
//   import acmpca "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("Cluster"), &clusterProps{
//   	clusterName: jsii.String("myCluster"),
//   	kafkaVersion: msk.kafkaVersion_V2_8_1(),
//   	vpc: vpc,
//   	encryptionInTransit: &encryptionInTransitConfig{
//   		clientBroker: msk.clientBrokerEncryption_TLS,
//   	},
//   	clientAuthentication: msk.clientAuthentication.tls(&tlsAuthProps{
//   		certificateAuthorities: []iCertificateAuthority{
//   			acmpca.certificateAuthority.fromCertificateAuthorityArn(this, jsii.String("CertificateAuthority"), jsii.String("arn:aws:acm-pca:us-west-2:1234567890:certificate-authority/11111111-1111-1111-1111-111111111111")),
//   		},
//   	}),
//   })
//
// Experimental.
type TlsAuthProps struct {
	// List of ACM Certificate Authorities to enable TLS authentication.
	// Experimental.
	CertificateAuthorities *[]awsacmpca.ICertificateAuthority `field:"optional" json:"certificateAuthorities" yaml:"certificateAuthorities"`
}
