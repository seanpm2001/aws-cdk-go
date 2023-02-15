package awsomics


// A genome reference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceItemProperty := &referenceItemProperty{
//   	referenceArn: jsii.String("referenceArn"),
//   }
//
type CfnAnnotationStore_ReferenceItemProperty struct {
	// The reference's ARN.
	ReferenceArn *string `field:"required" json:"referenceArn" yaml:"referenceArn"`
}
