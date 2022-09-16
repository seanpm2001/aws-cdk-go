package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   functionSAMPTProperty := &functionSAMPTProperty{
//   	functionName: jsii.String("functionName"),
//   }
//
type CfnStateMachine_FunctionSAMPTProperty struct {
	// `CfnStateMachine.FunctionSAMPTProperty.FunctionName`.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
}
