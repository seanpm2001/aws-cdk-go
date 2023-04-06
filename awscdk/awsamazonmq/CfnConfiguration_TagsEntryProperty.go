package awsamazonmq


// A key-value pair to associate with the configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagsEntryProperty := &TagsEntryProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
type CfnConfiguration_TagsEntryProperty struct {
	// The key in a key-value pair.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value in a key-value pair.
	Value *string `field:"required" json:"value" yaml:"value"`
}
