package awslakeformation


// Properties for defining a `CfnTag`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTagProps := &CfnTagProps{
//   	TagKey: jsii.String("tagKey"),
//   	TagValues: []*string{
//   		jsii.String("tagValues"),
//   	},
//
//   	// the properties below are optional
//   	CatalogId: jsii.String("catalogId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-tag.html
//
type CfnTagProps struct {
	// UTF-8 string, not less than 1 or more than 255 bytes long, matching the [single-line string pattern](https://docs.aws.amazon.com/lake-formation/latest/dg/aws-lake-formation-api-aws-lake-formation-api-common.html) .
	//
	// The key-name for the LF-tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-tag.html#cfn-lakeformation-tag-tagkey
	//
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// An array of UTF-8 strings, not less than 1 or more than 50 strings.
	//
	// A list of possible values of the corresponding `TagKey` of an LF-tag key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-tag.html#cfn-lakeformation-tag-tagvalues
	//
	TagValues *[]*string `field:"required" json:"tagValues" yaml:"tagValues"`
	// Catalog id string, not less than 1 or more than 255 bytes long, matching the [single-line string pattern](https://docs.aws.amazon.com/lake-formation/latest/dg/aws-lake-formation-api-aws-lake-formation-api-common.html) .
	//
	// The identifier for the Data Catalog . By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your AWS Lake Formation environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-tag.html#cfn-lakeformation-tag-catalogid
	//
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
}

