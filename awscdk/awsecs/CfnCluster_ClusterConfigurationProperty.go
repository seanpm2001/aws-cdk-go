package awsecs


// The execute command configuration for the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterConfigurationProperty := &ClusterConfigurationProperty{
//   	ExecuteCommandConfiguration: &ExecuteCommandConfigurationProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		LogConfiguration: &ExecuteCommandLogConfigurationProperty{
//   			CloudWatchEncryptionEnabled: jsii.Boolean(false),
//   			CloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   			S3BucketName: jsii.String("s3BucketName"),
//   			S3EncryptionEnabled: jsii.Boolean(false),
//   			S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   		},
//   		Logging: jsii.String("logging"),
//   	},
//   }
//
type CfnCluster_ClusterConfigurationProperty struct {
	// The details of the execute command configuration.
	ExecuteCommandConfiguration interface{} `field:"optional" json:"executeCommandConfiguration" yaml:"executeCommandConfiguration"`
}
