package awsapplicationautoscaling


// The service that supports Application AutoScaling.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var code code
//
//
//   handler := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
//   	runtime: lambda.runtime_PYTHON_3_7(),
//   	handler: jsii.String("index.handler"),
//   	code: code,
//
//   	reservedConcurrentExecutions: jsii.Number(2),
//   })
//
//   fnVer := handler.currentVersion
//
//   target := appscaling.NewScalableTarget(this, jsii.String("ScalableTarget"), &scalableTargetProps{
//   	serviceNamespace: appscaling.serviceNamespace_LAMBDA,
//   	maxCapacity: jsii.Number(100),
//   	minCapacity: jsii.Number(10),
//   	resourceId: fmt.Sprintf("function:%v:%v", handler.functionName, fnVer.version),
//   	scalableDimension: jsii.String("lambda:function:ProvisionedConcurrency"),
//   })
//
//   target.scaleToTrackMetric(jsii.String("PceTracking"), &basicTargetTrackingScalingPolicyProps{
//   	targetValue: jsii.Number(0.9),
//   	predefinedMetric: appscaling.predefinedMetric_LAMBDA_PROVISIONED_CONCURRENCY_UTILIZATION,
//   })
//
// Experimental.
type ServiceNamespace string

const (
	// Elastic Container Service.
	// Experimental.
	ServiceNamespace_ECS ServiceNamespace = "ECS"
	// Elastic Map Reduce.
	// Experimental.
	ServiceNamespace_ELASTIC_MAP_REDUCE ServiceNamespace = "ELASTIC_MAP_REDUCE"
	// Elastic Compute Cloud.
	// Experimental.
	ServiceNamespace_EC2 ServiceNamespace = "EC2"
	// App Stream.
	// Experimental.
	ServiceNamespace_APPSTREAM ServiceNamespace = "APPSTREAM"
	// Dynamo DB.
	// Experimental.
	ServiceNamespace_DYNAMODB ServiceNamespace = "DYNAMODB"
	// Relational Database Service.
	// Experimental.
	ServiceNamespace_RDS ServiceNamespace = "RDS"
	// SageMaker.
	// Experimental.
	ServiceNamespace_SAGEMAKER ServiceNamespace = "SAGEMAKER"
	// Custom Resource.
	// Experimental.
	ServiceNamespace_CUSTOM_RESOURCE ServiceNamespace = "CUSTOM_RESOURCE"
	// Lambda.
	// Experimental.
	ServiceNamespace_LAMBDA ServiceNamespace = "LAMBDA"
	// Comprehend.
	// Experimental.
	ServiceNamespace_COMPREHEND ServiceNamespace = "COMPREHEND"
	// Kafka.
	// Experimental.
	ServiceNamespace_KAFKA ServiceNamespace = "KAFKA"
	// ElastiCache.
	// Experimental.
	ServiceNamespace_ELASTICACHE ServiceNamespace = "ELASTICACHE"
)

