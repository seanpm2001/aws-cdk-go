package triggers

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

// Props for `Trigger`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var construct construct
//   var function_ function
//
//   triggerProps := &TriggerProps{
//   	Handler: function_,
//
//   	// the properties below are optional
//   	ExecuteAfter: []*construct{
//   		construct,
//   	},
//   	ExecuteBefore: []*construct{
//   		construct,
//   	},
//   	ExecuteOnHandlerChange: jsii.Boolean(false),
//   	InvocationType: awscdk.Triggers.InvocationType_EVENT,
//   	Timeout: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
type TriggerProps struct {
	// Adds trigger dependencies. Execute this trigger only after these construct scopes have been provisioned.
	//
	// You can also use `trigger.executeAfter()` to add additional dependencies.
	ExecuteAfter *[]constructs.Construct `field:"optional" json:"executeAfter" yaml:"executeAfter"`
	// Adds this trigger as a dependency on other constructs.
	//
	// This means that this
	// trigger will get executed *before* the given construct(s).
	//
	// You can also use `trigger.executeBefore()` to add additional dependants.
	ExecuteBefore *[]constructs.Construct `field:"optional" json:"executeBefore" yaml:"executeBefore"`
	// Re-executes the trigger every time the handler changes.
	//
	// This implies that the trigger is associated with the `currentVersion` of
	// the handler, which gets recreated every time the handler or its
	// configuration is updated.
	ExecuteOnHandlerChange *bool `field:"optional" json:"executeOnHandlerChange" yaml:"executeOnHandlerChange"`
	// The AWS Lambda function of the handler to execute.
	Handler awslambda.Function `field:"required" json:"handler" yaml:"handler"`
	// The invocation type to invoke the Lambda function with.
	InvocationType InvocationType `field:"optional" json:"invocationType" yaml:"invocationType"`
	// The timeout of the invocation call of the Lambda function to be triggered.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

