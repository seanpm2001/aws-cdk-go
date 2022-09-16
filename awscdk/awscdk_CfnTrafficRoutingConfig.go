// An experiment to bundle the entire CDK into a single module
package awscdk


// Traffic routing configuration settings.
//
// The type of the {@link CfnCodeDeployBlueGreenHookProps.trafficRoutingConfig} property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrafficRoutingConfig := &cfnTrafficRoutingConfig{
//   	type: monocdk.cfnTrafficRoutingType_ALL_AT_ONCE,
//
//   	// the properties below are optional
//   	timeBasedCanary: &cfnTrafficRoutingTimeBasedCanary{
//   		bakeTimeMins: jsii.Number(123),
//   		stepPercentage: jsii.Number(123),
//   	},
//   	timeBasedLinear: &cfnTrafficRoutingTimeBasedLinear{
//   		bakeTimeMins: jsii.Number(123),
//   		stepPercentage: jsii.Number(123),
//   	},
//   }
//
// Experimental.
type CfnTrafficRoutingConfig struct {
	// The type of traffic shifting used by the blue-green deployment configuration.
	// Experimental.
	Type CfnTrafficRoutingType `field:"required" json:"type" yaml:"type"`
	// The configuration for traffic routing when {@link type} is {@link CfnTrafficRoutingType.TIME_BASED_CANARY}.
	// Experimental.
	TimeBasedCanary *CfnTrafficRoutingTimeBasedCanary `field:"optional" json:"timeBasedCanary" yaml:"timeBasedCanary"`
	// The configuration for traffic routing when {@link type} is {@link CfnTrafficRoutingType.TIME_BASED_LINEAR}.
	// Experimental.
	TimeBasedLinear *CfnTrafficRoutingTimeBasedLinear `field:"optional" json:"timeBasedLinear" yaml:"timeBasedLinear"`
}
