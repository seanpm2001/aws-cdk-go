package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awsopensearchservice"
)

// Properties for the OpenSearch Data Source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var domain domain
//   var graphqlApi graphqlApi
//   var role role
//
//   openSearchDataSourceProps := &openSearchDataSourceProps{
//   	api: graphqlApi,
//   	domain: domain,
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	serviceRole: role,
//   }
//
// Experimental.
type OpenSearchDataSourceProps struct {
	// The API to attach this data source to.
	// Experimental.
	Api IGraphqlApi `field:"required" json:"api" yaml:"api"`
	// the description of the data source.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the data source.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The IAM service role to be assumed by AppSync to interact with the data source.
	// Experimental.
	ServiceRole awsiam.IRole `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// The OpenSearch domain containing the endpoint for the data source.
	// Experimental.
	Domain awsopensearchservice.IDomain `field:"required" json:"domain" yaml:"domain"`
}
