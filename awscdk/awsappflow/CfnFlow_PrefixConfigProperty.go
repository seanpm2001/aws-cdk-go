package awsappflow


// Specifies elements that Amazon AppFlow includes in the file and folder names in the flow destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   prefixConfigProperty := &PrefixConfigProperty{
//   	PathPrefixHierarchy: []*string{
//   		jsii.String("pathPrefixHierarchy"),
//   	},
//   	PrefixFormat: jsii.String("prefixFormat"),
//   	PrefixType: jsii.String("prefixType"),
//   }
//
type CfnFlow_PrefixConfigProperty struct {
	// `CfnFlow.PrefixConfigProperty.PathPrefixHierarchy`.
	PathPrefixHierarchy *[]*string `field:"optional" json:"pathPrefixHierarchy" yaml:"pathPrefixHierarchy"`
	// Determines the level of granularity for the date and time that's included in the prefix.
	PrefixFormat *string `field:"optional" json:"prefixFormat" yaml:"prefixFormat"`
	// Determines the format of the prefix, and whether it applies to the file name, file path, or both.
	PrefixType *string `field:"optional" json:"prefixType" yaml:"prefixType"`
}
