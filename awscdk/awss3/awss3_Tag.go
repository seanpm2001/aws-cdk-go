package awss3


// Tag.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tag := &tag{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
// Experimental.
type Tag struct {
	// key to e tagged.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// additional value.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}
