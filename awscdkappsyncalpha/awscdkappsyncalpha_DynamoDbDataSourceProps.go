// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for an AppSync DynamoDB datasource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appsync_alpha "github.com/aws/aws-cdk-go/awscdkappsyncalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var graphqlApi graphqlApi
//   var role role
//   var table table
//
//   dynamoDbDataSourceProps := &dynamoDbDataSourceProps{
//   	api: graphqlApi,
//   	table: table,
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	readOnlyAccess: jsii.Boolean(false),
//   	serviceRole: role,
//   	useCallerCredentials: jsii.Boolean(false),
//   }
//
// Experimental.
type DynamoDbDataSourceProps struct {
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
	// The DynamoDB table backing this data source.
	// Experimental.
	Table awsdynamodb.ITable `field:"required" json:"table" yaml:"table"`
	// Specify whether this DS is read only or has read and write permissions to the DynamoDB table.
	// Experimental.
	ReadOnlyAccess *bool `field:"optional" json:"readOnlyAccess" yaml:"readOnlyAccess"`
	// use credentials of caller to access DynamoDB.
	// Experimental.
	UseCallerCredentials *bool `field:"optional" json:"useCallerCredentials" yaml:"useCallerCredentials"`
}

