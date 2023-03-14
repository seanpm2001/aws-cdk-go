package awsglue


// Properties for defining a `CfnCrawler`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnCrawlerProps := &CfnCrawlerProps{
//   	Role: jsii.String("role"),
//   	Targets: &TargetsProperty{
//   		CatalogTargets: []interface{}{
//   			&CatalogTargetProperty{
//   				DatabaseName: jsii.String("databaseName"),
//   				Tables: []*string{
//   					jsii.String("tables"),
//   				},
//   			},
//   		},
//   		DynamoDbTargets: []interface{}{
//   			&DynamoDBTargetProperty{
//   				Path: jsii.String("path"),
//   			},
//   		},
//   		JdbcTargets: []interface{}{
//   			&JdbcTargetProperty{
//   				ConnectionName: jsii.String("connectionName"),
//   				Exclusions: []*string{
//   					jsii.String("exclusions"),
//   				},
//   				Path: jsii.String("path"),
//   			},
//   		},
//   		MongoDbTargets: []interface{}{
//   			&MongoDBTargetProperty{
//   				ConnectionName: jsii.String("connectionName"),
//   				Path: jsii.String("path"),
//   			},
//   		},
//   		S3Targets: []interface{}{
//   			&S3TargetProperty{
//   				ConnectionName: jsii.String("connectionName"),
//   				DlqEventQueueArn: jsii.String("dlqEventQueueArn"),
//   				EventQueueArn: jsii.String("eventQueueArn"),
//   				Exclusions: []*string{
//   					jsii.String("exclusions"),
//   				},
//   				Path: jsii.String("path"),
//   				SampleSize: jsii.Number(123),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Classifiers: []*string{
//   		jsii.String("classifiers"),
//   	},
//   	Configuration: jsii.String("configuration"),
//   	CrawlerSecurityConfiguration: jsii.String("crawlerSecurityConfiguration"),
//   	DatabaseName: jsii.String("databaseName"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	RecrawlPolicy: &RecrawlPolicyProperty{
//   		RecrawlBehavior: jsii.String("recrawlBehavior"),
//   	},
//   	Schedule: &ScheduleProperty{
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   	SchemaChangePolicy: &SchemaChangePolicyProperty{
//   		DeleteBehavior: jsii.String("deleteBehavior"),
//   		UpdateBehavior: jsii.String("updateBehavior"),
//   	},
//   	TablePrefix: jsii.String("tablePrefix"),
//   	Tags: tags,
//   }
//
type CfnCrawlerProps struct {
	// The Amazon Resource Name (ARN) of an IAM role that's used to access customer resources, such as Amazon Simple Storage Service (Amazon S3) data.
	Role *string `field:"required" json:"role" yaml:"role"`
	// A collection of targets to crawl.
	Targets interface{} `field:"required" json:"targets" yaml:"targets"`
	// A list of UTF-8 strings that specify the names of custom classifiers that are associated with the crawler.
	Classifiers *[]*string `field:"optional" json:"classifiers" yaml:"classifiers"`
	// Crawler configuration information.
	//
	// This versioned JSON string allows users to specify aspects of a crawler's behavior. For more information, see [Configuring a Crawler](https://docs.aws.amazon.com/glue/latest/dg/crawler-configuration.html) .
	Configuration *string `field:"optional" json:"configuration" yaml:"configuration"`
	// The name of the `SecurityConfiguration` structure to be used by this crawler.
	CrawlerSecurityConfiguration *string `field:"optional" json:"crawlerSecurityConfiguration" yaml:"crawlerSecurityConfiguration"`
	// The name of the database in which the crawler's output is stored.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// A description of the crawler.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the crawler.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A policy that specifies whether to crawl the entire dataset again, or to crawl only folders that were added since the last crawler run.
	RecrawlPolicy interface{} `field:"optional" json:"recrawlPolicy" yaml:"recrawlPolicy"`
	// For scheduled crawlers, the schedule when the crawler runs.
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// The policy that specifies update and delete behaviors for the crawler.
	//
	// The policy tells the crawler what to do in the event that it detects a change in a table that already exists in the customer's database at the time of the crawl. The `SchemaChangePolicy` does not affect whether or how new tables and partitions are added. New tables and partitions are always created regardless of the `SchemaChangePolicy` on a crawler.
	//
	// The SchemaChangePolicy consists of two components, `UpdateBehavior` and `DeleteBehavior` .
	SchemaChangePolicy interface{} `field:"optional" json:"schemaChangePolicy" yaml:"schemaChangePolicy"`
	// The prefix added to the names of tables that are created.
	TablePrefix *string `field:"optional" json:"tablePrefix" yaml:"tablePrefix"`
	// The tags to use with this crawler.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

