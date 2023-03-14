package awsiot


// The configuration that determines how many retries are allowed for each failure type for a job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jobExecutionsRetryConfigProperty := &JobExecutionsRetryConfigProperty{
//   	RetryCriteriaList: []interface{}{
//   		&RetryCriteriaProperty{
//   			FailureType: jsii.String("failureType"),
//   			NumberOfRetries: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnJobTemplate_JobExecutionsRetryConfigProperty struct {
	// `CfnJobTemplate.JobExecutionsRetryConfigProperty.RetryCriteriaList`.
	RetryCriteriaList interface{} `field:"optional" json:"retryCriteriaList" yaml:"retryCriteriaList"`
}

