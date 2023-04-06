package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventSourceProperty := &EventSourceProperty{
//   	Properties: &CloudWatchEventEventProperty{
//   		Method: jsii.String("method"),
//   		Path: jsii.String("path"),
//
//   		// the properties below are optional
//   		RestApiId: jsii.String("restApiId"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
type CfnStateMachine_EventSourceProperty struct {
	// `CfnStateMachine.EventSourceProperty.Properties`.
	Properties interface{} `field:"required" json:"properties" yaml:"properties"`
	// `CfnStateMachine.EventSourceProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
}
