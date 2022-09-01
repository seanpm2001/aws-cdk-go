// CDK Integration Testing Constructs
package awscdkintegtestsalpha


// The result of an Assertion wrapping the actual result data in another struct.
//
// Needed to access the whole message via getAtt() on the custom resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   assertionResult := &assertionResult{
//   	data: jsii.String("data"),
//
//   	// the properties below are optional
//   	failed: jsii.Boolean(false),
//   }
//
// Experimental.
type AssertionResult struct {
	// The result of an assertion.
	// Experimental.
	Data *string `field:"required" json:"data" yaml:"data"`
	// Whether or not the assertion failed.
	// Experimental.
	Failed *bool `field:"optional" json:"failed" yaml:"failed"`
}

