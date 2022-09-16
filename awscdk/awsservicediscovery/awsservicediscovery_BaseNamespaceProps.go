package awsservicediscovery


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baseNamespaceProps := &baseNamespaceProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
// Experimental.
type BaseNamespaceProps struct {
	// A name for the Namespace.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the Namespace.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}
