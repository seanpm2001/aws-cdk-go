// An experiment to bundle the entire CDK into a single module
package awscdk


// Options for creating a lazy string token.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   lazyStringValueOptions := &lazyStringValueOptions{
//   	displayHint: jsii.String("displayHint"),
//   }
//
// Experimental.
type LazyStringValueOptions struct {
	// Use the given name as a display hint.
	// Experimental.
	DisplayHint *string `field:"optional" json:"displayHint" yaml:"displayHint"`
}
