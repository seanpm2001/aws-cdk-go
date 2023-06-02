package awseventstargets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Properties to define an ECS Event Task.
//
// Example:
//   import ecs "github.com/aws/aws-cdk-go/awscdk"
//   var cluster iCluster
//   var taskDefinition taskDefinition
//
//
//   rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
//   	Schedule: events.Schedule_Rate(cdk.Duration_Hours(jsii.Number(1))),
//   })
//
//   rule.AddTarget(
//   targets.NewEcsTask(&EcsTaskProps{
//   	Cluster: cluster,
//   	TaskDefinition: taskDefinition,
//   	PropagateTags: ecs.PropagatedTagSource_TASK_DEFINITION,
//   	Tags: []tag{
//   		&tag{
//   			Key: jsii.String("my-tag"),
//   			Value: jsii.String("my-tag-value"),
//   		},
//   	},
//   }))
//
type EcsTaskProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// Cluster where service will be deployed.
	Cluster awsecs.ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// Task Definition of the task that should be started.
	TaskDefinition awsecs.ITaskDefinition `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
	// Container setting overrides.
	//
	// Key is the name of the container to override, value is the
	// values you want to override.
	ContainerOverrides *[]*ContainerOverride `field:"optional" json:"containerOverrides" yaml:"containerOverrides"`
	// The platform version on which to run your task.
	//
	// Unless you have specific compatibility requirements, you don't need to specify this.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html
	//
	PlatformVersion awsecs.FargatePlatformVersion `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// Specifies whether to propagate the tags from the task definition to the task.
	//
	// If no value is specified, the tags are not propagated.
	PropagateTags awsecs.PropagatedTagSource `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// Existing IAM role to run the ECS task.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Existing security groups to use for the task's ENIs.
	//
	// (Only applicable in case the TaskDefinition is configured for AwsVpc networking).
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// In what subnets to place the task's ENIs.
	//
	// (Only applicable in case the TaskDefinition is configured for AwsVpc networking).
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// The metadata that you apply to the task to help you categorize and organize them.
	//
	// Each tag consists of a key and an optional value, both of which you define.
	Tags *[]*Tag `field:"optional" json:"tags" yaml:"tags"`
	// How many tasks should be started when this event is triggered.
	TaskCount *float64 `field:"optional" json:"taskCount" yaml:"taskCount"`
}

