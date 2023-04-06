package awsquicksight


// The parameters for OpenSearch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amazonElasticsearchParametersProperty := &AmazonElasticsearchParametersProperty{
//   	Domain: jsii.String("domain"),
//   }
//
type CfnDataSource_AmazonElasticsearchParametersProperty struct {
	// The OpenSearch domain.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
}
