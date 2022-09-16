package awsautoscaling


// Input for Signals.renderCreationPolicy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   renderSignalsOptions := &renderSignalsOptions{
//   	desiredCapacity: jsii.Number(123),
//   	minCapacity: jsii.Number(123),
//   }
//
// Experimental.
type RenderSignalsOptions struct {
	// The desiredCapacity of the ASG.
	// Experimental.
	DesiredCapacity *float64 `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// The minSize of the ASG.
	// Experimental.
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
}
