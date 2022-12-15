package awswafv2


// A single label name condition for a `Condition` in a logging filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   labelNameConditionProperty := &labelNameConditionProperty{
//   	labelName: jsii.String("labelName"),
//   }
//
type CfnLoggingConfiguration_LabelNameConditionProperty struct {
	// The label name that a log record must contain in order to meet the condition.
	//
	// This must be a fully qualified label name. Fully qualified labels have a prefix, optional namespaces, and label name. The prefix identifies the rule group or web ACL context of the rule that added the label.
	LabelName *string `field:"required" json:"labelName" yaml:"labelName"`
}

