package awsdynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// Properties of a DynamoDB Table.
//
// Use {@link TableProps} for all table properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var key key
//
//   tableOptions := &tableOptions{
//   	partitionKey: &attribute{
//   		name: jsii.String("name"),
//   		type: awscdk.Aws_dynamodb.attributeType_BINARY,
//   	},
//
//   	// the properties below are optional
//   	billingMode: awscdk.*Aws_dynamodb.billingMode_PAY_PER_REQUEST,
//   	contributorInsightsEnabled: jsii.Boolean(false),
//   	encryption: awscdk.*Aws_dynamodb.tableEncryption_DEFAULT,
//   	encryptionKey: key,
//   	pointInTimeRecovery: jsii.Boolean(false),
//   	readCapacity: jsii.Number(123),
//   	removalPolicy: monocdk.removalPolicy_DESTROY,
//   	replicationRegions: []*string{
//   		jsii.String("replicationRegions"),
//   	},
//   	replicationTimeout: duration,
//   	serverSideEncryption: jsii.Boolean(false),
//   	sortKey: &attribute{
//   		name: jsii.String("name"),
//   		type: awscdk.*Aws_dynamodb.*attributeType_BINARY,
//   	},
//   	stream: awscdk.*Aws_dynamodb.streamViewType_NEW_IMAGE,
//   	tableClass: awscdk.*Aws_dynamodb.tableClass_STANDARD,
//   	timeToLiveAttribute: jsii.String("timeToLiveAttribute"),
//   	waitForReplicationToFinish: jsii.Boolean(false),
//   	writeCapacity: jsii.Number(123),
//   }
//
// Experimental.
type TableOptions struct {
	// Partition key attribute definition.
	// Experimental.
	PartitionKey *Attribute `field:"required" json:"partitionKey" yaml:"partitionKey"`
	// Sort key attribute definition.
	// Experimental.
	SortKey *Attribute `field:"optional" json:"sortKey" yaml:"sortKey"`
	// Specify how you are charged for read and write throughput and how you manage capacity.
	// Experimental.
	BillingMode BillingMode `field:"optional" json:"billingMode" yaml:"billingMode"`
	// Whether CloudWatch contributor insights is enabled.
	// Experimental.
	ContributorInsightsEnabled *bool `field:"optional" json:"contributorInsightsEnabled" yaml:"contributorInsightsEnabled"`
	// Whether server-side encryption with an AWS managed customer master key is enabled.
	//
	// This property cannot be set if `serverSideEncryption` is set.
	//
	// > **NOTE**: if you set this to `CUSTOMER_MANAGED` and `encryptionKey` is not
	// > specified, the key that the Tablet generates for you will be created with
	// > default permissions. If you are using CDKv2, these permissions will be
	// > sufficient to enable the key for use with DynamoDB tables.  If you are
	// > using CDKv1, make sure the feature flag
	// > `@aws-cdk/aws-kms:defaultKeyPolicies` is set to `true` in your `cdk.json`.
	// Experimental.
	Encryption TableEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// External KMS key to use for table encryption.
	//
	// This property can only be set if `encryption` is set to `TableEncryption.CUSTOMER_MANAGED`.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Whether point-in-time recovery is enabled.
	// Experimental.
	PointInTimeRecovery *bool `field:"optional" json:"pointInTimeRecovery" yaml:"pointInTimeRecovery"`
	// The read capacity for the table.
	//
	// Careful if you add Global Secondary Indexes, as
	// those will share the table's provisioned throughput.
	//
	// Can only be provided if billingMode is Provisioned.
	// Experimental.
	ReadCapacity *float64 `field:"optional" json:"readCapacity" yaml:"readCapacity"`
	// The removal policy to apply to the DynamoDB Table.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Regions where replica tables will be created.
	// Experimental.
	ReplicationRegions *[]*string `field:"optional" json:"replicationRegions" yaml:"replicationRegions"`
	// The timeout for a table replication operation in a single region.
	// Experimental.
	ReplicationTimeout awscdk.Duration `field:"optional" json:"replicationTimeout" yaml:"replicationTimeout"`
	// Whether server-side encryption with an AWS managed customer master key is enabled.
	//
	// This property cannot be set if `encryption` and/or `encryptionKey` is set.
	// Deprecated: This property is deprecated. In order to obtain the same behavior as
	// enabling this, set the `encryption` property to `TableEncryption.AWS_MANAGED` instead.
	ServerSideEncryption *bool `field:"optional" json:"serverSideEncryption" yaml:"serverSideEncryption"`
	// When an item in the table is modified, StreamViewType determines what information is written to the stream for this table.
	// Experimental.
	Stream StreamViewType `field:"optional" json:"stream" yaml:"stream"`
	// Specify the table class.
	// Experimental.
	TableClass TableClass `field:"optional" json:"tableClass" yaml:"tableClass"`
	// The name of TTL attribute.
	// Experimental.
	TimeToLiveAttribute *string `field:"optional" json:"timeToLiveAttribute" yaml:"timeToLiveAttribute"`
	// Indicates whether CloudFormation stack waits for replication to finish.
	//
	// If set to false, the CloudFormation resource will mark the resource as
	// created and replication will be completed asynchronously. This property is
	// ignored if replicationRegions property is not set.
	//
	// DO NOT UNSET this property if adding/removing multiple replicationRegions
	// in one deployment, as CloudFormation only supports one region replication
	// at a time. CDK overcomes this limitation by waiting for replication to
	// finish before starting new replicationRegion.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-replicas
	//
	// Experimental.
	WaitForReplicationToFinish *bool `field:"optional" json:"waitForReplicationToFinish" yaml:"waitForReplicationToFinish"`
	// The write capacity for the table.
	//
	// Careful if you add Global Secondary Indexes, as
	// those will share the table's provisioned throughput.
	//
	// Can only be provided if billingMode is Provisioned.
	// Experimental.
	WriteCapacity *float64 `field:"optional" json:"writeCapacity" yaml:"writeCapacity"`
}
