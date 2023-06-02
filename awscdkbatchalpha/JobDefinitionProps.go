package awscdkbatchalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props common to all JobDefinitions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import batch_alpha "github.com/aws/aws-cdk-go/awscdkbatchalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//   var retryStrategy retryStrategy
//
//   jobDefinitionProps := &JobDefinitionProps{
//   	JobDefinitionName: jsii.String("jobDefinitionName"),
//   	Parameters: map[string]interface{}{
//   		"parametersKey": parameters,
//   	},
//   	RetryAttempts: jsii.Number(123),
//   	RetryStrategies: []*retryStrategy{
//   		retryStrategy,
//   	},
//   	SchedulingPriority: jsii.Number(123),
//   	Timeout: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
// Experimental.
type JobDefinitionProps struct {
	// The name of this job definition.
	// Experimental.
	JobDefinitionName *string `field:"optional" json:"jobDefinitionName" yaml:"jobDefinitionName"`
	// The default parameters passed to the container These parameters can be referenced in the `command` that you give to the container.
	// See: https://docs.aws.amazon.com/batch/latest/userguide/job_definition_parameters.html#parameters
	//
	// Experimental.
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The number of times to retry a job.
	//
	// The job is retried on failure the same number of attempts as the value.
	// Experimental.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// Defines the retry behavior for this job.
	// Experimental.
	RetryStrategies *[]RetryStrategy `field:"optional" json:"retryStrategies" yaml:"retryStrategies"`
	// The priority of this Job.
	//
	// Only used in Fairshare Scheduling
	// to decide which job to run first when there are multiple jobs
	// with the same share identifier.
	// Experimental.
	SchedulingPriority *float64 `field:"optional" json:"schedulingPriority" yaml:"schedulingPriority"`
	// The timeout time for jobs that are submitted with this job definition.
	//
	// After the amount of time you specify passes,
	// Batch terminates your jobs if they aren't finished.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

