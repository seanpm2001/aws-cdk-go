package awsbatch


// An object that represents a job timeout configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeoutProperty := &timeoutProperty{
//   	attemptDurationSeconds: jsii.Number(123),
//   }
//
type CfnJobDefinition_TimeoutProperty struct {
	// The job timeout time (in seconds) that's measured from the job attempt's `startedAt` timestamp.
	//
	// After this time passes, AWS Batch terminates your jobs if they aren't finished. The minimum value for the timeout is 60 seconds.
	AttemptDurationSeconds *float64 `field:"optional" json:"attemptDurationSeconds" yaml:"attemptDurationSeconds"`
}

