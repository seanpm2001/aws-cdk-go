package awswisdom


// Configuration information for Amazon AppIntegrations to automatically ingest content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appIntegrationsConfigurationProperty := &AppIntegrationsConfigurationProperty{
//   	AppIntegrationArn: jsii.String("appIntegrationArn"),
//
//   	// the properties below are optional
//   	ObjectFields: []*string{
//   		jsii.String("objectFields"),
//   	},
//   }
//
type CfnKnowledgeBase_AppIntegrationsConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the AppIntegrations DataIntegration to use for ingesting content.
	//
	// - For [Salesforce](https://docs.aws.amazon.com/https://developer.salesforce.com/docs/atlas.en-us.knowledge_dev.meta/knowledge_dev/sforce_api_objects_knowledge__kav.htm) , your AppIntegrations DataIntegration must have an ObjectConfiguration if objectFields is not provided, including at least `Id` , `ArticleNumber` , `VersionNumber` , `Title` , `PublishStatus` , and `IsDeleted` as source fields.
	// - For [ServiceNow](https://docs.aws.amazon.com/https://developer.servicenow.com/dev.do#!/reference/api/rome/rest/knowledge-management-api) , your AppIntegrations DataIntegration must have an ObjectConfiguration if objectFields is not provided, including at least `number` , `short_description` , `sys_mod_count` , `workflow_state` , and `active` as source fields.
	// - For [Zendesk](https://docs.aws.amazon.com/https://developer.zendesk.com/api-reference/help_center/help-center-api/articles/) , your AppIntegrations DataIntegration must have an ObjectConfiguration if `objectFields` is not provided, including at least `id` , `title` , `updated_at` , and `draft` as source fields.
	// - For [SharePoint](https://docs.aws.amazon.com/https://learn.microsoft.com/en-us/sharepoint/dev/sp-add-ins/sharepoint-net-server-csom-jsom-and-rest-api-index) , your AppIntegrations DataIntegration must have a FileConfiguration, including only file extensions that are among `docx` , `pdf` , `html` , `htm` , and `txt` .
	AppIntegrationArn *string `field:"required" json:"appIntegrationArn" yaml:"appIntegrationArn"`
	// The fields from the source that are made available to your agents in Wisdom.
	//
	// Optional if ObjectConfiguration is included in the provided DataIntegration.
	//
	// - For [Salesforce](https://docs.aws.amazon.com/https://developer.salesforce.com/docs/atlas.en-us.knowledge_dev.meta/knowledge_dev/sforce_api_objects_knowledge__kav.htm) , you must include at least `Id` , `ArticleNumber` , `VersionNumber` , `Title` , `PublishStatus` , and `IsDeleted` .
	// - For [ServiceNow](https://docs.aws.amazon.com/https://developer.servicenow.com/dev.do#!/reference/api/rome/rest/knowledge-management-api) , you must include at least `number` , `short_description` , `sys_mod_count` , `workflow_state` , and `active` .
	// - For [Zendesk](https://docs.aws.amazon.com/https://developer.zendesk.com/api-reference/help_center/help-center-api/articles/) , you must include at least `id` , `title` , `updated_at` , and `draft` .
	//
	// Make sure to include additional fields. These fields are indexed and used to source recommendations.
	ObjectFields *[]*string `field:"optional" json:"objectFields" yaml:"objectFields"`
}
