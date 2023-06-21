package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Common options for creating cluster instances (both serverless and provisioned).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var clusterInstanceType clusterInstanceType
//   var key key
//   var parameterGroup parameterGroup
//
//   clusterInstanceProps := &ClusterInstanceProps{
//   	InstanceType: clusterInstanceType,
//
//   	// the properties below are optional
//   	AllowMajorVersionUpgrade: jsii.Boolean(false),
//   	AutoMinorVersionUpgrade: jsii.Boolean(false),
//   	EnablePerformanceInsights: jsii.Boolean(false),
//   	InstanceIdentifier: jsii.String("instanceIdentifier"),
//   	IsFromLegacyInstanceProps: jsii.Boolean(false),
//   	ParameterGroup: parameterGroup,
//   	Parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	PerformanceInsightEncryptionKey: key,
//   	PerformanceInsightRetention: awscdk.Aws_rds.PerformanceInsightRetention_DEFAULT,
//   	PromotionTier: jsii.Number(123),
//   	PubliclyAccessible: jsii.Boolean(false),
//   }
//
type ClusterInstanceProps struct {
	// Whether to allow upgrade of major version for the DB instance.
	AllowMajorVersionUpgrade *bool `field:"optional" json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// Whether to enable automatic upgrade of minor version for the DB instance.
	AutoMinorVersionUpgrade *bool `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// Whether to enable Performance Insights for the DB instance.
	EnablePerformanceInsights *bool `field:"optional" json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// The identifier for the database instance.
	InstanceIdentifier *string `field:"optional" json:"instanceIdentifier" yaml:"instanceIdentifier"`
	// The DB parameter group to associate with the instance.
	//
	// This is only needed if you need to configure different parameter
	// groups for each individual instance, otherwise you should not
	// provide this and just use the cluster parameter group.
	ParameterGroup IParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// The parameters in the DBParameterGroup to create automatically.
	//
	// You can only specify parameterGroup or parameters but not both.
	// You need to use a versioned engine to auto-generate a DBParameterGroup.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The AWS KMS key for encryption of Performance Insights data.
	PerformanceInsightEncryptionKey awskms.IKey `field:"optional" json:"performanceInsightEncryptionKey" yaml:"performanceInsightEncryptionKey"`
	// The amount of time, in days, to retain Performance Insights data.
	PerformanceInsightRetention PerformanceInsightRetention `field:"optional" json:"performanceInsightRetention" yaml:"performanceInsightRetention"`
	// Indicates whether the DB instance is an internet-facing instance.
	PubliclyAccessible *bool `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The type of cluster instance to create.
	//
	// Can be either
	// provisioned or serverless v2.
	InstanceType ClusterInstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// Only used for migrating existing clusters from using `instanceProps` to `writer` and `readers`.
	IsFromLegacyInstanceProps *bool `field:"optional" json:"isFromLegacyInstanceProps" yaml:"isFromLegacyInstanceProps"`
	// The promotion tier of the cluster instance.
	//
	// This matters more for serverlessV2 instances. If a serverless
	// instance is in tier 0-1 then it will scale with the writer.
	//
	// For provisioned instances this just determines the failover priority.
	// If multiple instances have the same priority then one will be picked at random.
	PromotionTier *float64 `field:"optional" json:"promotionTier" yaml:"promotionTier"`
}
