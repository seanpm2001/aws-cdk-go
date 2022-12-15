package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
)

// The properties for the QueueProcessingFargateService service.
//
// Example:
//   var cluster cluster
//
//   cluster.enableFargateCapacityProviders()
//
//   queueProcessingFargateService := ecsPatterns.NewQueueProcessingFargateService(this, jsii.String("Service"), &queueProcessingFargateServiceProps{
//   	cluster: cluster,
//   	memoryLimitMiB: jsii.Number(512),
//   	image: ecs.containerImage.fromRegistry(jsii.String("test")),
//   	capacityProviderStrategies: []capacityProviderStrategy{
//   		&capacityProviderStrategy{
//   			capacityProvider: jsii.String("FARGATE_SPOT"),
//   			weight: jsii.Number(2),
//   		},
//   		&capacityProviderStrategy{
//   			capacityProvider: jsii.String("FARGATE"),
//   			weight: jsii.Number(1),
//   		},
//   	},
//   })
//
// Experimental.
type QueueProcessingFargateServiceProps struct {
	// The image used to start a container.
	// Experimental.
	Image awsecs.ContainerImage `field:"required" json:"image" yaml:"image"`
	// A list of Capacity Provider strategies used to place a service.
	// Experimental.
	CapacityProviderStrategies *[]*awsecs.CapacityProviderStrategy `field:"optional" json:"capacityProviderStrategies" yaml:"capacityProviderStrategies"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	// Experimental.
	CircuitBreaker *awsecs.DeploymentCircuitBreaker `field:"optional" json:"circuitBreaker" yaml:"circuitBreaker"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	// Experimental.
	Cluster awsecs.ICluster `field:"optional" json:"cluster" yaml:"cluster"`
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	// Experimental.
	DeploymentController *awsecs.DeploymentController `field:"optional" json:"deploymentController" yaml:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	// Deprecated: - Use `minScalingCapacity` or a literal object instead.
	DesiredTaskCount *float64 `field:"optional" json:"desiredTaskCount" yaml:"desiredTaskCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// Experimental.
	EnableECSManagedTags *bool `field:"optional" json:"enableECSManagedTags" yaml:"enableECSManagedTags"`
	// Flag to indicate whether to enable logging.
	// Experimental.
	EnableLogging *bool `field:"optional" json:"enableLogging" yaml:"enableLogging"`
	// The environment variables to pass to the container.
	//
	// The variable `QUEUE_NAME` with value `queue.queueName` will
	// always be passed.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The name of a family that the task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	// Experimental.
	Family *string `field:"optional" json:"family" yaml:"family"`
	// The log driver to use.
	// Experimental.
	LogDriver awsecs.LogDriver `field:"optional" json:"logDriver" yaml:"logDriver"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	// Experimental.
	MaxHealthyPercent *float64 `field:"optional" json:"maxHealthyPercent" yaml:"maxHealthyPercent"`
	// The maximum number of times that a message can be received by consumers.
	//
	// When this value is exceeded for a message the message will be automatically sent to the Dead Letter Queue.
	// Experimental.
	MaxReceiveCount *float64 `field:"optional" json:"maxReceiveCount" yaml:"maxReceiveCount"`
	// Maximum capacity to scale to.
	// Experimental.
	MaxScalingCapacity *float64 `field:"optional" json:"maxScalingCapacity" yaml:"maxScalingCapacity"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	// Experimental.
	MinHealthyPercent *float64 `field:"optional" json:"minHealthyPercent" yaml:"minHealthyPercent"`
	// Minimum capacity to scale to.
	// Experimental.
	MinScalingCapacity *float64 `field:"optional" json:"minScalingCapacity" yaml:"minScalingCapacity"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	// Experimental.
	PropagateTags awsecs.PropagatedTagSource `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// A queue for which to process items from.
	//
	// If specified and this is a FIFO queue, the queue name must end in the string '.fifo'. See
	// [CreateQueue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueue.html)
	// Experimental.
	Queue awssqs.IQueue `field:"optional" json:"queue" yaml:"queue"`
	// The number of seconds that Dead Letter Queue retains a message.
	// Experimental.
	RetentionPeriod awscdk.Duration `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// The intervals for scaling based on the SQS queue's ApproximateNumberOfMessagesVisible metric.
	//
	// Maps a range of metric values to a particular scaling behavior. See
	// [Simple and Step Scaling Policies for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html)
	// Experimental.
	ScalingSteps *[]*awsapplicationautoscaling.ScalingInterval `field:"optional" json:"scalingSteps" yaml:"scalingSteps"`
	// The secret to expose to the container as an environment variable.
	// Experimental.
	Secrets *map[string]awsecs.Secret `field:"optional" json:"secrets" yaml:"secrets"`
	// The name of the service.
	// Experimental.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Timeout of processing a single message.
	//
	// After dequeuing, the processor has this much time to handle the message and delete it from the queue
	// before it becomes visible again for dequeueing by another processor. Values must be between 0 and (12 hours).
	// Experimental.
	VisibilityTimeout awscdk.Duration `field:"optional" json:"visibilityTimeout" yaml:"visibilityTimeout"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Specifies whether the task's elastic network interface receives a public IP address.
	//
	// If true, each task will receive a public IP address.
	// Experimental.
	AssignPublicIp *bool `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Optional name for the container added.
	// Experimental.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The number of cpu units used by the task.
	//
	// Valid values, which determines your range of valid values for the memory parameter:
	//
	// 256 (.25 vCPU) - Available memory values: 0.5GB, 1GB, 2GB
	//
	// 512 (.5 vCPU) - Available memory values: 1GB, 2GB, 3GB, 4GB
	//
	// 1024 (1 vCPU) - Available memory values: 2GB, 3GB, 4GB, 5GB, 6GB, 7GB, 8GB
	//
	// 2048 (2 vCPU) - Available memory values: Between 4GB and 16GB in 1GB increments
	//
	// 4096 (4 vCPU) - Available memory values: Between 8GB and 30GB in 1GB increments
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Experimental.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// The health check command and associated configuration parameters for the container.
	// Experimental.
	HealthCheck *awsecs.HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The amount (in MiB) of memory used by the task.
	//
	// This field is required and you must use one of the following values, which determines your range of valid values
	// for the cpu parameter:
	//
	// 0.5GB, 1GB, 2GB - Available cpu values: 256 (.25 vCPU)
	//
	// 1GB, 2GB, 3GB, 4GB - Available cpu values: 512 (.5 vCPU)
	//
	// 2GB, 3GB, 4GB, 5GB, 6GB, 7GB, 8GB - Available cpu values: 1024 (1 vCPU)
	//
	// Between 4GB and 16GB in 1GB increments - Available cpu values: 2048 (2 vCPU)
	//
	// Between 8GB and 30GB in 1GB increments - Available cpu values: 4096 (4 vCPU)
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	// Experimental.
	MemoryLimitMiB *float64 `field:"optional" json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// The platform version on which to run your service.
	//
	// If one is not specified, the LATEST platform version is used by default. For more information, see
	// [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	// Experimental.
	PlatformVersion awsecs.FargatePlatformVersion `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// The security groups to associate with the service.
	//
	// If you do not specify a security group, a new security group is created.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The subnets to associate with the service.
	// Experimental.
	TaskSubnets *awsec2.SubnetSelection `field:"optional" json:"taskSubnets" yaml:"taskSubnets"`
}

