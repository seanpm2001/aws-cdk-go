package awss3objectlambda


// The alias of an Object Lambda Access Point.
//
// For more information, see [How to use a bucket-style alias for your S3 bucket Object Lambda Access Point](https://docs.aws.amazon.com/AmazonS3/latest/userguide/olap-use.html#ol-access-points-alias) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aliasProperty := &AliasProperty{
//   	Status: jsii.String("status"),
//   	Value: jsii.String("value"),
//   }
//
type CfnAccessPoint_AliasProperty struct {
	// The status of the Object Lambda Access Point alias.
	//
	// If the status is `PROVISIONING` , the Object Lambda Access Point is provisioning the alias and the alias is not ready for use yet. If the status is `READY` , the Object Lambda Access Point alias is successfully provisioned and ready for use.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The alias value of the Object Lambda Access Point.
	Value *string `field:"optional" json:"value" yaml:"value"`
}
