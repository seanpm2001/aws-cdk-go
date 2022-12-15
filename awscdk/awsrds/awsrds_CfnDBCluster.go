package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::RDS::DBCluster`.
//
// The `AWS::RDS::DBCluster` resource creates an Amazon Aurora DB cluster or Multi-AZ DB cluster.
//
// For more information about creating an Aurora DB cluster, see [Creating an Amazon Aurora DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.CreateInstance.html) in the *Amazon Aurora User Guide* .
//
// For more information about creating a Multi-AZ DB cluster, see [Creating a Multi-AZ DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/create-multi-az-db-cluster.html) in the *Amazon RDS User Guide* .
//
// > You can only create this resource in AWS Regions where Amazon Aurora or Multi-AZ DB clusters are supported.
//
// *Updating DB clusters*
//
// When properties labeled " *Update requires:* [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) " are updated, AWS CloudFormation first creates a replacement DB cluster, then changes references from other dependent resources to point to the replacement DB cluster, and finally deletes the old DB cluster.
//
// > We highly recommend that you take a snapshot of the database before updating the stack. If you don't, you lose the data when AWS CloudFormation replaces your DB cluster. To preserve your data, perform the following procedure:
// >
// > - Deactivate any applications that are using the DB cluster so that there's no activity on the DB instance.
// > - Create a snapshot of the DB cluster. For more information, see [Creating a DB Cluster Snapshot](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_CreateSnapshotCluster.html) .
// > - If you want to restore your DB cluster using a DB cluster snapshot, modify the updated template with your DB cluster changes and add the `SnapshotIdentifier` property with the ID of the DB cluster snapshot that you want to use.
// >
// > After you restore a DB cluster with a `SnapshotIdentifier` property, you must specify the same `SnapshotIdentifier` property for any future updates to the DB cluster. When you specify this property for an update, the DB cluster is not restored from the DB cluster snapshot again, and the data in the database is not changed. However, if you don't specify the `SnapshotIdentifier` property, an empty DB cluster is created, and the original DB cluster is deleted. If you specify a property that is different from the previous snapshot restore property, a new DB cluster is restored from the specified `SnapshotIdentifier` property, and the original DB cluster is deleted.
// > - Update the stack.
//
// Currently, when you are updating the stack for an Aurora Serverless DB cluster, you can't include changes to any other properties when you specify one of the following properties: `PreferredBackupWindow` , `PreferredMaintenanceWindow` , and `Port` . This limitation doesn't apply to provisioned DB clusters.
//
// For more information about updating other properties of this resource, see `[ModifyDBCluster](https://docs.aws.amazon.com//AmazonRDS/latest/APIReference/API_ModifyDBCluster.html)` . For more information about updating stacks, see [AWS CloudFormation Stacks Updates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks.html) .
//
// *Deleting DB clusters*
//
// The default `DeletionPolicy` for `AWS::RDS::DBCluster` resources is `Snapshot` . For more information about how AWS CloudFormation deletes resources, see [DeletionPolicy Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBCluster := awscdk.Aws_rds.NewCfnDBCluster(this, jsii.String("MyCfnDBCluster"), &cfnDBClusterProps{
//   	allocatedStorage: jsii.Number(123),
//   	associatedRoles: []interface{}{
//   		&dBClusterRoleProperty{
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			featureName: jsii.String("featureName"),
//   		},
//   	},
//   	autoMinorVersionUpgrade: jsii.Boolean(false),
//   	availabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	backtrackWindow: jsii.Number(123),
//   	backupRetentionPeriod: jsii.Number(123),
//   	copyTagsToSnapshot: jsii.Boolean(false),
//   	databaseName: jsii.String("databaseName"),
//   	dbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	dbClusterInstanceClass: jsii.String("dbClusterInstanceClass"),
//   	dbClusterParameterGroupName: jsii.String("dbClusterParameterGroupName"),
//   	dbInstanceParameterGroupName: jsii.String("dbInstanceParameterGroupName"),
//   	dbSubnetGroupName: jsii.String("dbSubnetGroupName"),
//   	deletionProtection: jsii.Boolean(false),
//   	domain: jsii.String("domain"),
//   	domainIamRoleName: jsii.String("domainIamRoleName"),
//   	enableCloudwatchLogsExports: []*string{
//   		jsii.String("enableCloudwatchLogsExports"),
//   	},
//   	enableHttpEndpoint: jsii.Boolean(false),
//   	enableIamDatabaseAuthentication: jsii.Boolean(false),
//   	engine: jsii.String("engine"),
//   	engineMode: jsii.String("engineMode"),
//   	engineVersion: jsii.String("engineVersion"),
//   	globalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//   	iops: jsii.Number(123),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	masterUsername: jsii.String("masterUsername"),
//   	masterUserPassword: jsii.String("masterUserPassword"),
//   	monitoringInterval: jsii.Number(123),
//   	monitoringRoleArn: jsii.String("monitoringRoleArn"),
//   	networkType: jsii.String("networkType"),
//   	performanceInsightsEnabled: jsii.Boolean(false),
//   	performanceInsightsKmsKeyId: jsii.String("performanceInsightsKmsKeyId"),
//   	performanceInsightsRetentionPeriod: jsii.Number(123),
//   	port: jsii.Number(123),
//   	preferredBackupWindow: jsii.String("preferredBackupWindow"),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	publiclyAccessible: jsii.Boolean(false),
//   	replicationSourceIdentifier: jsii.String("replicationSourceIdentifier"),
//   	restoreType: jsii.String("restoreType"),
//   	scalingConfiguration: &scalingConfigurationProperty{
//   		autoPause: jsii.Boolean(false),
//   		maxCapacity: jsii.Number(123),
//   		minCapacity: jsii.Number(123),
//   		secondsUntilAutoPause: jsii.Number(123),
//   		timeoutAction: jsii.String("timeoutAction"),
//   	},
//   	serverlessV2ScalingConfiguration: &serverlessV2ScalingConfigurationProperty{
//   		maxCapacity: jsii.Number(123),
//   		minCapacity: jsii.Number(123),
//   	},
//   	snapshotIdentifier: jsii.String("snapshotIdentifier"),
//   	sourceDbClusterIdentifier: jsii.String("sourceDbClusterIdentifier"),
//   	sourceRegion: jsii.String("sourceRegion"),
//   	storageEncrypted: jsii.Boolean(false),
//   	storageType: jsii.String("storageType"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	useLatestRestorableTime: jsii.Boolean(false),
//   	vpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   })
//
type CfnDBCluster interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The amount of storage in gibibytes (GiB) to allocate to each DB instance in the Multi-AZ DB cluster.
	//
	// This setting is required to create a Multi-AZ DB cluster.
	//
	// Valid for: Multi-AZ DB clusters only.
	AllocatedStorage() *float64
	SetAllocatedStorage(val *float64)
	// Provides a list of the AWS Identity and Access Management (IAM) roles that are associated with the DB cluster.
	//
	// IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other Amazon Web Services on your behalf.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	AssociatedRoles() interface{}
	SetAssociatedRoles(val interface{})
	// The Amazon Resource Name (ARN) for the DB cluster.
	AttrDbClusterArn() *string
	// The AWS Region -unique, immutable identifier for the DB cluster.
	//
	// This identifier is found in AWS CloudTrail log entries whenever the KMS key for the DB cluster is accessed.
	AttrDbClusterResourceId() *string
	// The connection endpoint for the DB cluster.
	//
	// For example: `mystack-mydbcluster-123456789012.us-east-2.rds.amazonaws.com`
	AttrEndpointAddress() *string
	// The port number that will accept connections on this DB cluster.
	//
	// For example: `3306`.
	AttrEndpointPort() *string
	// The reader endpoint for the DB cluster.
	//
	// For example: `mystack-mydbcluster-ro-123456789012.us-east-2.rds.amazonaws.com`
	AttrReadEndpointAddress() *string
	// A value that indicates whether minor engine upgrades are applied automatically to the DB cluster during the maintenance window.
	//
	// By default, minor engine upgrades are applied automatically.
	//
	// Valid for: Multi-AZ DB clusters only.
	AutoMinorVersionUpgrade() interface{}
	SetAutoMinorVersionUpgrade(val interface{})
	// A list of Availability Zones (AZs) where instances in the DB cluster can be created.
	//
	// For information on AWS Regions and Availability Zones, see [Choosing the Regions and Availability Zones](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.RegionsAndAvailabilityZones.html) in the *Amazon Aurora User Guide* .
	//
	// Valid for: Aurora DB clusters only.
	AvailabilityZones() *[]*string
	SetAvailabilityZones(val *[]*string)
	// The target backtrack window, in seconds. To disable backtracking, set this value to 0.
	//
	// > Currently, Backtrack is only supported for Aurora MySQL DB clusters.
	//
	// Default: 0
	//
	// Constraints:
	//
	// - If specified, this value must be set to a number from 0 to 259,200 (72 hours).
	//
	// Valid for: Aurora MySQL DB clusters only.
	BacktrackWindow() *float64
	SetBacktrackWindow(val *float64)
	// The number of days for which automated backups are retained.
	//
	// Default: 1
	//
	// Constraints:
	//
	// - Must be a value from 1 to 35
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	BackupRetentionPeriod() *float64
	SetBackupRetentionPeriod(val *float64)
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// A value that indicates whether to copy all tags from the DB cluster to snapshots of the DB cluster.
	//
	// The default is not to copy them.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	CopyTagsToSnapshot() interface{}
	SetCopyTagsToSnapshot(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The name of your database.
	//
	// If you don't provide a name, then Amazon RDS won't create a database in this DB cluster. For naming constraints, see [Naming Constraints](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_Limits.html#RDS_Limits.Constraints) in the *Amazon Aurora User Guide* .
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	DatabaseName() *string
	SetDatabaseName(val *string)
	// The DB cluster identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	// - Must contain from 1 to 63 letters, numbers, or hyphens.
	// - First character must be a letter.
	// - Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Example: `my-cluster1`
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	DbClusterIdentifier() *string
	SetDbClusterIdentifier(val *string)
	// The compute and memory capacity of each DB instance in the Multi-AZ DB cluster, for example db.m6gd.xlarge. Not all DB instance classes are available in all AWS Regions , or for all database engines.
	//
	// For the full list of DB instance classes and availability for your engine, see [DB instance class](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) in the *Amazon RDS User Guide* .
	//
	// This setting is required to create a Multi-AZ DB cluster.
	//
	// Valid for: Multi-AZ DB clusters only.
	DbClusterInstanceClass() *string
	SetDbClusterInstanceClass(val *string)
	// The name of the DB cluster parameter group to associate with this DB cluster.
	//
	// > If you apply a parameter group to an existing DB cluster, then its DB instances might need to reboot. This can result in an outage while the DB instances are rebooting.
	// >
	// > If you apply a change to parameter group associated with a stopped DB cluster, then the update stack waits until the DB cluster is started.
	//
	// To list all of the available DB cluster parameter group names, use the following command:
	//
	// `aws rds describe-db-cluster-parameter-groups --query "DBClusterParameterGroups[].DBClusterParameterGroupName" --output text`
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	DbClusterParameterGroupName() *string
	SetDbClusterParameterGroupName(val *string)
	// The name of the DB parameter group to apply to all instances of the DB cluster.
	//
	// > When you apply a parameter group using the `DBInstanceParameterGroupName` parameter, the DB cluster isn't rebooted automatically. Also, parameter changes are applied immediately rather than during the next maintenance window.
	//
	// Default: The existing name setting
	//
	// Constraints:
	//
	// - The DB parameter group must be in the same DB parameter group family as this DB cluster.
	DbInstanceParameterGroupName() *string
	SetDbInstanceParameterGroupName(val *string)
	// A DB subnet group that you want to associate with this DB cluster.
	//
	// If you are restoring a DB cluster to a point in time with `RestoreType` set to `copy-on-write` , and don't specify a DB subnet group name, then the DB cluster is restored with a default DB subnet group.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	DbSubnetGroupName() *string
	SetDbSubnetGroupName(val *string)
	// A value that indicates whether the DB cluster has deletion protection enabled.
	//
	// The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	// Indicates the directory ID of the Active Directory to create the DB cluster.
	//
	// For Amazon Aurora DB clusters, Amazon RDS can use Kerberos authentication to authenticate users that connect to the DB cluster.
	//
	// For more information, see [Kerberos authentication](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/kerberos-authentication.html) in the *Amazon Aurora User Guide* .
	//
	// Valid for: Aurora DB clusters only.
	Domain() *string
	SetDomain(val *string)
	// Specifies the name of the IAM role to use when making API calls to the Directory Service.
	//
	// Valid for: Aurora DB clusters only.
	DomainIamRoleName() *string
	SetDomainIamRoleName(val *string)
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	//
	// The values in the list depend on the DB engine being used. For more information, see [Publishing Database Logs to Amazon CloudWatch Logs](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch) in the *Amazon Aurora User Guide* .
	//
	// *Aurora MySQL*
	//
	// Valid values: `audit` , `error` , `general` , `slowquery`
	//
	// *Aurora PostgreSQL*
	//
	// Valid values: `postgresql`
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	EnableCloudwatchLogsExports() *[]*string
	SetEnableCloudwatchLogsExports(val *[]*string)
	// A value that indicates whether to enable the HTTP endpoint for an Aurora Serverless DB cluster.
	//
	// By default, the HTTP endpoint is disabled.
	//
	// When enabled, the HTTP endpoint provides a connectionless web service API for running SQL queries on the Aurora Serverless DB cluster. You can also query your database from inside the RDS console with the query editor.
	//
	// For more information, see [Using the Data API for Aurora Serverless](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html) in the *Amazon Aurora User Guide* .
	//
	// Valid for: Aurora DB clusters only.
	EnableHttpEndpoint() interface{}
	SetEnableHttpEndpoint(val interface{})
	// A value that indicates whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	//
	// By default, mapping is disabled.
	//
	// For more information, see [IAM Database Authentication](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html) in the *Amazon Aurora User Guide.*
	//
	// Valid for: Aurora DB clusters only.
	EnableIamDatabaseAuthentication() interface{}
	SetEnableIamDatabaseAuthentication(val interface{})
	// The name of the database engine to be used for this DB cluster.
	//
	// Valid Values:
	//
	// - `aurora` (for MySQL 5.6-compatible Aurora)
	// - `aurora-mysql` (for MySQL 5.7-compatible Aurora)
	// - `aurora-postgresql`
	// - `mysql`
	// - `postgres`
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	Engine() *string
	SetEngine(val *string)
	// The DB engine mode of the DB cluster, either `provisioned` , `serverless` , `parallelquery` , `global` , or `multimaster` .
	//
	// The `parallelquery` engine mode isn't required for Aurora MySQL version 1.23 and higher 1.x versions, and version 2.09 and higher 2.x versions.
	//
	// The `global` engine mode isn't required for Aurora MySQL version 1.22 and higher 1.x versions, and `global` engine mode isn't required for any 2.x versions.
	//
	// The `multimaster` engine mode only applies for DB clusters created with Aurora MySQL version 5.6.10a.
	//
	// For Aurora PostgreSQL, the `global` engine mode isn't required, and both the `parallelquery` and the `multimaster` engine modes currently aren't supported.
	//
	// Limitations and requirements apply to some DB engine modes. For more information, see the following sections in the *Amazon Aurora User Guide* :
	//
	// - [Limitations of Aurora Serverless](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html#aurora-serverless.limitations)
	// - [Limitations of Parallel Query](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-mysql-parallel-query.html#aurora-mysql-parallel-query-limitations)
	// - [Limitations of Aurora Global Databases](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html#aurora-global-database.limitations)
	// - [Limitations of Multi-Master Clusters](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-multi-master.html#aurora-multi-master-limitations)
	//
	// Valid for: Aurora DB clusters only.
	EngineMode() *string
	SetEngineMode(val *string)
	// The version number of the database engine to use.
	//
	// To list all of the available engine versions for `aurora` (for MySQL 5.6-compatible Aurora), use the following command:
	//
	// `aws rds describe-db-engine-versions --engine aurora --query "DBEngineVersions[].EngineVersion"`
	//
	// To list all of the available engine versions for `aurora-mysql` (for MySQL 5.7-compatible Aurora), use the following command:
	//
	// `aws rds describe-db-engine-versions --engine aurora-mysql --query "DBEngineVersions[].EngineVersion"`
	//
	// To list all of the available engine versions for `aurora-postgresql` , use the following command:
	//
	// `aws rds describe-db-engine-versions --engine aurora-postgresql --query "DBEngineVersions[].EngineVersion"`
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	EngineVersion() *string
	SetEngineVersion(val *string)
	// If you are configuring an Aurora global database cluster and want your Aurora DB cluster to be a secondary member in the global database cluster, specify the global cluster ID of the global database cluster.
	//
	// To define the primary database cluster of the global cluster, use the [AWS::RDS::GlobalCluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-globalcluster.html) resource.
	//
	// If you aren't configuring a global database cluster, don't specify this property.
	//
	// > To remove the DB cluster from a global database cluster, specify an empty value for the `GlobalClusterIdentifier` property.
	//
	// For information about Aurora global databases, see [Working with Amazon Aurora Global Databases](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html) in the *Amazon Aurora User Guide* .
	//
	// Valid for: Aurora DB clusters only.
	GlobalClusterIdentifier() *string
	SetGlobalClusterIdentifier(val *string)
	// The amount of Provisioned IOPS (input/output operations per second) to be initially allocated for each DB instance in the Multi-AZ DB cluster.
	//
	// For information about valid IOPS values, see [Amazon RDS Provisioned IOPS storage](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS) in the *Amazon RDS User Guide* .
	//
	// This setting is required to create a Multi-AZ DB cluster.
	//
	// Constraints: Must be a multiple between .5 and 50 of the storage amount for the DB cluster.
	//
	// Valid for: Multi-AZ DB clusters only.
	Iops() *float64
	SetIops(val *float64)
	// The Amazon Resource Name (ARN) of the AWS KMS key that is used to encrypt the database instances in the DB cluster, such as `arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef` .
	//
	// If you enable the `StorageEncrypted` property but don't specify this property, the default KMS key is used. If you specify this property, you must set the `StorageEncrypted` property to `true` .
	//
	// If you specify the `SnapshotIdentifier` property, the `StorageEncrypted` property value is inherited from the snapshot, and if the DB cluster is encrypted, the specified `KmsKeyId` property is used.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The name of the master user for the DB cluster.
	//
	// > If you specify the `SourceDBClusterIdentifier` , `SnapshotIdentifier` , or `GlobalClusterIdentifier` property, don't specify this property. The value is inherited from the source DB cluster, the snapshot, or the primary DB cluster for the global database cluster, respectively.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	MasterUsername() *string
	SetMasterUsername(val *string)
	// The master password for the DB instance.
	//
	// > If you specify the `SourceDBClusterIdentifier` , `SnapshotIdentifier` , or `GlobalClusterIdentifier` property, don't specify this property. The value is inherited from the source DB cluster, the snapshot, or the primary DB cluster for the global database cluster, respectively.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	MasterUserPassword() *string
	SetMasterUserPassword(val *string)
	// The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB cluster.
	//
	// To turn off collecting Enhanced Monitoring metrics, specify 0. The default is 0.
	//
	// If `MonitoringRoleArn` is specified, also set `MonitoringInterval` to a value other than 0.
	//
	// Valid Values: `0, 1, 5, 10, 15, 30, 60`
	//
	// Valid for: Multi-AZ DB clusters only.
	MonitoringInterval() *float64
	SetMonitoringInterval(val *float64)
	// The Amazon Resource Name (ARN) for the IAM role that permits RDS to send Enhanced Monitoring metrics to Amazon CloudWatch Logs.
	//
	// An example is `arn:aws:iam:123456789012:role/emaccess` . For information on creating a monitoring role, see [Setting up and enabling Enhanced Monitoring](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling) in the *Amazon RDS User Guide* .
	//
	// If `MonitoringInterval` is set to a value other than 0, supply a `MonitoringRoleArn` value.
	//
	// Valid for: Multi-AZ DB clusters only.
	MonitoringRoleArn() *string
	SetMonitoringRoleArn(val *string)
	// The network type of the DB cluster.
	//
	// Valid values:
	//
	// - `IPV4`
	// - `DUAL`
	//
	// The network type is determined by the `DBSubnetGroup` specified for the DB cluster. A `DBSubnetGroup` can support only the IPv4 protocol or the IPv4 and IPv6 protocols ( `DUAL` ).
	//
	// For more information, see [Working with a DB instance in a VPC](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html) in the *Amazon Aurora User Guide.*
	//
	// Valid for: Aurora DB clusters only.
	NetworkType() *string
	SetNetworkType(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// A value that indicates whether to turn on Performance Insights for the DB cluster.
	//
	// For more information, see [Using Amazon Performance Insights](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html) in the *Amazon RDS User Guide* .
	//
	// Valid for: Multi-AZ DB clusters only.
	PerformanceInsightsEnabled() interface{}
	SetPerformanceInsightsEnabled(val interface{})
	// The AWS KMS key identifier for encryption of Performance Insights data.
	//
	// The AWS KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.
	//
	// If you don't specify a value for `PerformanceInsightsKMSKeyId` , then Amazon RDS uses your default KMS key. There is a default KMS key for your AWS account . Your AWS account has a different default KMS key for each AWS Region .
	//
	// Valid for: Multi-AZ DB clusters only.
	PerformanceInsightsKmsKeyId() *string
	SetPerformanceInsightsKmsKeyId(val *string)
	// The number of days to retain Performance Insights data. The default is 7 days. The following values are valid:.
	//
	// - 7
	// - *month* * 31, where *month* is a number of months from 1-23
	// - 731
	//
	// For example, the following values are valid:
	//
	// - 93 (3 months * 31)
	// - 341 (11 months * 31)
	// - 589 (19 months * 31)
	// - 731
	//
	// If you specify a retention period such as 94, which isn't a valid value, RDS issues an error.
	//
	// Valid for: Multi-AZ DB clusters only.
	PerformanceInsightsRetentionPeriod() *float64
	SetPerformanceInsightsRetentionPeriod(val *float64)
	// The port number on which the DB instances in the DB cluster accept connections.
	//
	// Default:
	//
	// - When `EngineMode` is `provisioned` , `3306` (for both Aurora MySQL and Aurora PostgreSQL)
	// - When `EngineMode` is `serverless` :
	//
	// - `3306` when `Engine` is `aurora` or `aurora-mysql`
	// - `5432` when `Engine` is `aurora-postgresql`
	//
	// > The `No interruption` on update behavior only applies to DB clusters. If you are updating a DB instance, see [Port](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-port) for the AWS::RDS::DBInstance resource.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	Port() *float64
	SetPort(val *float64)
	// The daily time range during which automated backups are created.
	//
	// For more information, see [Backup Window](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.Backups.BackupWindow) in the *Amazon Aurora User Guide.*
	//
	// Constraints:
	//
	// - Must be in the format `hh24:mi-hh24:mi` .
	// - Must be in Universal Coordinated Time (UTC).
	// - Must not conflict with the preferred maintenance window.
	// - Must be at least 30 minutes.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	PreferredBackupWindow() *string
	SetPreferredBackupWindow(val *string)
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	//
	// The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region, occurring on a random day of the week. To see the time blocks available, see [Adjusting the Preferred DB Cluster Maintenance Window](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora) in the *Amazon Aurora User Guide.*
	//
	// Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.
	//
	// Constraints: Minimum 30-minute window.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	// A value that indicates whether the DB cluster is publicly accessible.
	//
	// When the DB cluster is publicly accessible, its Domain Name System (DNS) endpoint resolves to the private IP address from within the DB cluster's virtual private cloud (VPC). It resolves to the public IP address from outside of the DB cluster's VPC. Access to the DB cluster is ultimately controlled by the security group it uses. That public access isn't permitted if the security group assigned to the DB cluster doesn't permit it.
	//
	// When the DB cluster isn't publicly accessible, it is an internal DB cluster with a DNS name that resolves to a private IP address.
	//
	// Default: The default behavior varies depending on whether `DBSubnetGroupName` is specified.
	//
	// If `DBSubnetGroupName` isn't specified, and `PubliclyAccessible` isn't specified, the following applies:
	//
	// - If the default VPC in the target Region doesn’t have an internet gateway attached to it, the DB cluster is private.
	// - If the default VPC in the target Region has an internet gateway attached to it, the DB cluster is public.
	//
	// If `DBSubnetGroupName` is specified, and `PubliclyAccessible` isn't specified, the following applies:
	//
	// - If the subnets are part of a VPC that doesn’t have an internet gateway attached to it, the DB cluster is private.
	// - If the subnets are part of a VPC that has an internet gateway attached to it, the DB cluster is public.
	//
	// Valid for: Multi-AZ DB clusters only.
	PubliclyAccessible() interface{}
	SetPubliclyAccessible(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The Amazon Resource Name (ARN) of the source DB instance or DB cluster if this DB cluster is created as a read replica.
	//
	// Valid for: Aurora DB clusters only.
	ReplicationSourceIdentifier() *string
	SetReplicationSourceIdentifier(val *string)
	// The type of restore to be performed. You can specify one of the following values:.
	//
	// - `full-copy` - The new DB cluster is restored as a full copy of the source DB cluster.
	// - `copy-on-write` - The new DB cluster is restored as a clone of the source DB cluster.
	//
	// Constraints: You can't specify `copy-on-write` if the engine version of the source DB cluster is earlier than 1.11.
	//
	// If you don't specify a `RestoreType` value, then the new DB cluster is restored as a full copy of the source DB cluster.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	RestoreType() *string
	SetRestoreType(val *string)
	// The `ScalingConfiguration` property type specifies the scaling configuration of an Aurora Serverless DB cluster.
	//
	// This property is only supported for Aurora Serverless v1. For Aurora Serverless v2, use `ServerlessV2ScalingConfiguration` property.
	//
	// Valid for: Aurora DB clusters only.
	ScalingConfiguration() interface{}
	SetScalingConfiguration(val interface{})
	// The `ServerlessV2ScalingConfiguration` property type specifies the scaling configuration of an Aurora Serverless V2 DB cluster.
	//
	// This property is only supported for Aurora Serverless v2. For Aurora Serverless v1, use `ScalingConfiguration` property.
	//
	// Valid for: Aurora DB clusters only.
	ServerlessV2ScalingConfiguration() interface{}
	SetServerlessV2ScalingConfiguration(val interface{})
	// The identifier for the DB snapshot or DB cluster snapshot to restore from.
	//
	// You can use either the name or the Amazon Resource Name (ARN) to specify a DB cluster snapshot. However, you can use only the ARN to specify a DB snapshot.
	//
	// After you restore a DB cluster with a `SnapshotIdentifier` property, you must specify the same `SnapshotIdentifier` property for any future updates to the DB cluster. When you specify this property for an update, the DB cluster is not restored from the snapshot again, and the data in the database is not changed. However, if you don't specify the `SnapshotIdentifier` property, an empty DB cluster is created, and the original DB cluster is deleted. If you specify a property that is different from the previous snapshot restore property, a new DB cluster is restored from the specified `SnapshotIdentifier` property, and the original DB cluster is deleted.
	//
	// If you specify the `SnapshotIdentifier` property to restore a DB cluster (as opposed to specifying it for DB cluster updates), then don't specify the following properties:
	//
	// - `GlobalClusterIdentifier`
	// - `MasterUsername`
	// - `MasterUserPassword`
	// - `ReplicationSourceIdentifier`
	// - `RestoreType`
	// - `SourceDBClusterIdentifier`
	// - `SourceRegion`
	// - `StorageEncrypted` (for an encrypted snapshot)
	// - `UseLatestRestorableTime`
	//
	// Constraints:
	//
	// - Must match the identifier of an existing Snapshot.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	SnapshotIdentifier() *string
	SetSnapshotIdentifier(val *string)
	// When restoring a DB cluster to a point in time, the identifier of the source DB cluster from which to restore.
	//
	// Constraints:
	//
	// - Must match the identifier of an existing DBCluster.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	SourceDbClusterIdentifier() *string
	SetSourceDbClusterIdentifier(val *string)
	// The AWS Region which contains the source DB cluster when replicating a DB cluster. For example, `us-east-1` .
	//
	// Valid for: Aurora DB clusters only.
	SourceRegion() *string
	SetSourceRegion(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Indicates whether the DB cluster is encrypted.
	//
	// If you specify the `KmsKeyId` property, then you must enable encryption.
	//
	// If you specify the `SourceDBClusterIdentifier` property, don't specify this property. The value is inherited from the source DB cluster, and if the DB cluster is encrypted, the specified `KmsKeyId` property is used.
	//
	// If you specify the `SnapshotIdentifier` and the specified snapshot is encrypted, don't specify this property. The value is inherited from the snapshot, and the specified `KmsKeyId` property is used.
	//
	// If you specify the `SnapshotIdentifier` and the specified snapshot isn't encrypted, you can use this property to specify that the restored DB cluster is encrypted. Specify the `KmsKeyId` property for the KMS key to use for encryption. If you don't want the restored DB cluster to be encrypted, then don't set this property or set it to `false` .
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	StorageEncrypted() interface{}
	SetStorageEncrypted(val interface{})
	// Specifies the storage type to be associated with the DB cluster.
	//
	// This setting is required to create a Multi-AZ DB cluster.
	//
	// Valid values: `io1`
	//
	// When specified, a value for the `Iops` parameter is required.
	//
	// Default: `io1`
	//
	// Valid for: Multi-AZ DB clusters only.
	StorageType() *string
	SetStorageType(val *string)
	// An optional array of key-value pairs to apply to this DB cluster.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// A value that indicates whether to restore the DB cluster to the latest restorable backup time.
	//
	// By default, the DB cluster is not restored to the latest restorable backup time.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	UseLatestRestorableTime() interface{}
	SetUseLatestRestorableTime(val interface{})
	// A list of EC2 VPC security groups to associate with this DB cluster.
	//
	// If you plan to update the resource, don't specify VPC security groups in a shared VPC.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters.
	VpcSecurityGroupIds() *[]*string
	SetVpcSecurityGroupIds(val *[]*string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDBCluster
type jsiiProxy_CfnDBCluster struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDBCluster) AllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) AssociatedRoles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) AttrDbClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDbClusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) AttrDbClusterResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDbClusterResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) AttrEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) AttrEndpointPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) AttrReadEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrReadEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) AutoMinorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) BacktrackWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backtrackWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) BackupRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) CopyTagsToSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) DbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) DbClusterInstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterInstanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) DbClusterParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) DbInstanceParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) DbSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) DomainIamRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIamRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) EnableCloudwatchLogsExports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enableCloudwatchLogsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) EnableHttpEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttpEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) EnableIamDatabaseAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIamDatabaseAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) EngineMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) GlobalClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) MasterUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) MasterUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) MonitoringInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitoringInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) MonitoringRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) PerformanceInsightsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"performanceInsightsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) PerformanceInsightsKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceInsightsKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) PerformanceInsightsRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceInsightsRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) PreferredBackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) ReplicationSourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) RestoreType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) ScalingConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) ServerlessV2ScalingConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverlessV2ScalingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) SnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) SourceDbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) SourceRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) StorageEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) UseLatestRestorableTime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLatestRestorableTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}


// Create a new `AWS::RDS::DBCluster`.
func NewCfnDBCluster(scope awscdk.Construct, id *string, props *CfnDBClusterProps) CfnDBCluster {
	_init_.Initialize()

	if err := validateNewCfnDBClusterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDBCluster{}

	_jsii_.Create(
		"monocdk.aws_rds.CfnDBCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::RDS::DBCluster`.
func NewCfnDBCluster_Override(c CfnDBCluster, scope awscdk.Construct, id *string, props *CfnDBClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_rds.CfnDBCluster",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetAllocatedStorage(val *float64) {
	_jsii_.Set(
		j,
		"allocatedStorage",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetAssociatedRoles(val interface{}) {
	if err := j.validateSetAssociatedRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatedRoles",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetAutoMinorVersionUpgrade(val interface{}) {
	if err := j.validateSetAutoMinorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetAvailabilityZones(val *[]*string) {
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetBacktrackWindow(val *float64) {
	_jsii_.Set(
		j,
		"backtrackWindow",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetBackupRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"backupRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetCopyTagsToSnapshot(val interface{}) {
	if err := j.validateSetCopyTagsToSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToSnapshot",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetDbClusterIdentifier(val *string) {
	_jsii_.Set(
		j,
		"dbClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetDbClusterInstanceClass(val *string) {
	_jsii_.Set(
		j,
		"dbClusterInstanceClass",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetDbClusterParameterGroupName(val *string) {
	_jsii_.Set(
		j,
		"dbClusterParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetDbInstanceParameterGroupName(val *string) {
	_jsii_.Set(
		j,
		"dbInstanceParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetDbSubnetGroupName(val *string) {
	_jsii_.Set(
		j,
		"dbSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetDomain(val *string) {
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetDomainIamRoleName(val *string) {
	_jsii_.Set(
		j,
		"domainIamRoleName",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetEnableCloudwatchLogsExports(val *[]*string) {
	_jsii_.Set(
		j,
		"enableCloudwatchLogsExports",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetEnableHttpEndpoint(val interface{}) {
	if err := j.validateSetEnableHttpEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHttpEndpoint",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetEnableIamDatabaseAuthentication(val interface{}) {
	if err := j.validateSetEnableIamDatabaseAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableIamDatabaseAuthentication",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetEngine(val *string) {
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetEngineMode(val *string) {
	_jsii_.Set(
		j,
		"engineMode",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetEngineVersion(val *string) {
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetGlobalClusterIdentifier(val *string) {
	_jsii_.Set(
		j,
		"globalClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetIops(val *float64) {
	_jsii_.Set(
		j,
		"iops",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetMasterUsername(val *string) {
	_jsii_.Set(
		j,
		"masterUsername",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetMasterUserPassword(val *string) {
	_jsii_.Set(
		j,
		"masterUserPassword",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetMonitoringInterval(val *float64) {
	_jsii_.Set(
		j,
		"monitoringInterval",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetMonitoringRoleArn(val *string) {
	_jsii_.Set(
		j,
		"monitoringRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetNetworkType(val *string) {
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetPerformanceInsightsEnabled(val interface{}) {
	if err := j.validateSetPerformanceInsightsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetPerformanceInsightsKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"performanceInsightsKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetPerformanceInsightsRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"performanceInsightsRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetPreferredBackupWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredBackupWindow",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetPreferredMaintenanceWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetPubliclyAccessible(val interface{}) {
	if err := j.validateSetPubliclyAccessibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetReplicationSourceIdentifier(val *string) {
	_jsii_.Set(
		j,
		"replicationSourceIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetRestoreType(val *string) {
	_jsii_.Set(
		j,
		"restoreType",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetScalingConfiguration(val interface{}) {
	if err := j.validateSetScalingConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scalingConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetServerlessV2ScalingConfiguration(val interface{}) {
	if err := j.validateSetServerlessV2ScalingConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverlessV2ScalingConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetSnapshotIdentifier(val *string) {
	_jsii_.Set(
		j,
		"snapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetSourceDbClusterIdentifier(val *string) {
	_jsii_.Set(
		j,
		"sourceDbClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetSourceRegion(val *string) {
	_jsii_.Set(
		j,
		"sourceRegion",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetStorageEncrypted(val interface{}) {
	if err := j.validateSetStorageEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageEncrypted",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetStorageType(val *string) {
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetUseLatestRestorableTime(val interface{}) {
	if err := j.validateSetUseLatestRestorableTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLatestRestorableTime",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetVpcSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnDBCluster_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDBCluster_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.CfnDBCluster",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDBCluster_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnDBCluster_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.CfnDBCluster",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDBCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDBCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.CfnDBCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBCluster_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_rds.CfnDBCluster",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDBCluster) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDBCluster) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDBCluster) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDBCluster) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDBCluster) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDBCluster) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDBCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDBCluster) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBCluster) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBCluster) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDBCluster) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDBCluster) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDBCluster) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBCluster) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDBCluster) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDBCluster) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBCluster) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBCluster) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDBCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBCluster) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBCluster) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

