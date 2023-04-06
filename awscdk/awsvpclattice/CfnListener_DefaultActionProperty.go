package awsvpclattice


// The action for the default rule.
//
// Each listener has a default rule. Each rule consists of a priority, one or more actions, and one or more conditions. The default rule is the rule that's used if no other rules match. Each rule must include exactly one of the following types of actions: `forward` or `fixed-response` , and it must be the last action to be performed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultActionProperty := &DefaultActionProperty{
//   	Forward: &ForwardProperty{
//   		TargetGroups: []interface{}{
//   			&WeightedTargetGroupProperty{
//   				TargetGroupIdentifier: jsii.String("targetGroupIdentifier"),
//
//   				// the properties below are optional
//   				Weight: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
type CfnListener_DefaultActionProperty struct {
	// Describes a forward action.
	//
	// You can use forward actions to route requests to one or more target groups.
	Forward interface{} `field:"required" json:"forward" yaml:"forward"`
}

