package awspipes


// The collection of event patterns used to filter events.
//
// To remove a filter, specify a `FilterCriteria` object with an empty array of `Filter` objects.
//
// For more information, see [Events and Event Patterns](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) in the *Amazon EventBridge User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterCriteriaProperty := &FilterCriteriaProperty{
//   	Filters: []interface{}{
//   		&FilterProperty{
//   			Pattern: jsii.String("pattern"),
//   		},
//   	},
//   }
//
type CfnPipe_FilterCriteriaProperty struct {
	// The event patterns.
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
}
