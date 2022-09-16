package awss3


// A bucket associated with a specific Region when creating Multi-Region Access Points.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   regionProperty := &regionProperty{
//   	bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	accountId: jsii.String("accountId"),
//   }
//
type CfnMultiRegionAccessPoint_RegionProperty struct {
	// The name of the associated bucket for the Region.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// `CfnMultiRegionAccessPoint.RegionProperty.AccountId`.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
}
