package awssagemaker


// A list of algorithms that were used to create a model package.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceAlgorithmSpecificationProperty := &SourceAlgorithmSpecificationProperty{
//   	SourceAlgorithms: []interface{}{
//   		&SourceAlgorithmProperty{
//   			AlgorithmName: jsii.String("algorithmName"),
//
//   			// the properties below are optional
//   			ModelDataUrl: jsii.String("modelDataUrl"),
//   		},
//   	},
//   }
//
type CfnModelPackage_SourceAlgorithmSpecificationProperty struct {
	// A list of the algorithms that were used to create a model package.
	SourceAlgorithms interface{} `field:"required" json:"sourceAlgorithms" yaml:"sourceAlgorithms"`
}
