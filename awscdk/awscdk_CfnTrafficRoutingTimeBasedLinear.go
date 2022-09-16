// An experiment to bundle the entire CDK into a single module
package awscdk


// The traffic routing configuration if {@link CfnTrafficRoutingConfig.type} is {@link CfnTrafficRoutingType.TIME_BASED_LINEAR}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrafficRoutingTimeBasedLinear := &cfnTrafficRoutingTimeBasedLinear{
//   	bakeTimeMins: jsii.Number(123),
//   	stepPercentage: jsii.Number(123),
//   }
//
// Experimental.
type CfnTrafficRoutingTimeBasedLinear struct {
	// The number of minutes between the first and second traffic shifts of a time-based linear deployment.
	// Experimental.
	BakeTimeMins *float64 `field:"optional" json:"bakeTimeMins" yaml:"bakeTimeMins"`
	// The percentage of traffic that is shifted at the start of each increment of a time-based linear deployment.
	//
	// The step percentage must be 14% or greater.
	// Experimental.
	StepPercentage *float64 `field:"optional" json:"stepPercentage" yaml:"stepPercentage"`
}
