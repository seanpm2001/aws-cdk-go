package awsappflow


// The operation to be performed on the provided source fields.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorOperatorProperty := &ConnectorOperatorProperty{
//   	Amplitude: jsii.String("amplitude"),
//   	CustomConnector: jsii.String("customConnector"),
//   	Datadog: jsii.String("datadog"),
//   	Dynatrace: jsii.String("dynatrace"),
//   	GoogleAnalytics: jsii.String("googleAnalytics"),
//   	InforNexus: jsii.String("inforNexus"),
//   	Marketo: jsii.String("marketo"),
//   	Pardot: jsii.String("pardot"),
//   	S3: jsii.String("s3"),
//   	Salesforce: jsii.String("salesforce"),
//   	SapoData: jsii.String("sapoData"),
//   	ServiceNow: jsii.String("serviceNow"),
//   	Singular: jsii.String("singular"),
//   	Slack: jsii.String("slack"),
//   	Trendmicro: jsii.String("trendmicro"),
//   	Veeva: jsii.String("veeva"),
//   	Zendesk: jsii.String("zendesk"),
//   }
//
type CfnFlow_ConnectorOperatorProperty struct {
	// The operation to be performed on the provided Amplitude source fields.
	Amplitude *string `field:"optional" json:"amplitude" yaml:"amplitude"`
	// Operators supported by the custom connector.
	CustomConnector *string `field:"optional" json:"customConnector" yaml:"customConnector"`
	// The operation to be performed on the provided Datadog source fields.
	Datadog *string `field:"optional" json:"datadog" yaml:"datadog"`
	// The operation to be performed on the provided Dynatrace source fields.
	Dynatrace *string `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// The operation to be performed on the provided Google Analytics source fields.
	GoogleAnalytics *string `field:"optional" json:"googleAnalytics" yaml:"googleAnalytics"`
	// The operation to be performed on the provided Infor Nexus source fields.
	InforNexus *string `field:"optional" json:"inforNexus" yaml:"inforNexus"`
	// The operation to be performed on the provided Marketo source fields.
	Marketo *string `field:"optional" json:"marketo" yaml:"marketo"`
	// The operation to be performed on the provided Salesforce Pardot source fields.
	Pardot *string `field:"optional" json:"pardot" yaml:"pardot"`
	// The operation to be performed on the provided Amazon S3 source fields.
	S3 *string `field:"optional" json:"s3" yaml:"s3"`
	// The operation to be performed on the provided Salesforce source fields.
	Salesforce *string `field:"optional" json:"salesforce" yaml:"salesforce"`
	// The operation to be performed on the provided SAPOData source fields.
	SapoData *string `field:"optional" json:"sapoData" yaml:"sapoData"`
	// The operation to be performed on the provided ServiceNow source fields.
	ServiceNow *string `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// The operation to be performed on the provided Singular source fields.
	Singular *string `field:"optional" json:"singular" yaml:"singular"`
	// The operation to be performed on the provided Slack source fields.
	Slack *string `field:"optional" json:"slack" yaml:"slack"`
	// The operation to be performed on the provided Trend Micro source fields.
	Trendmicro *string `field:"optional" json:"trendmicro" yaml:"trendmicro"`
	// The operation to be performed on the provided Veeva source fields.
	Veeva *string `field:"optional" json:"veeva" yaml:"veeva"`
	// The operation to be performed on the provided Zendesk source fields.
	Zendesk *string `field:"optional" json:"zendesk" yaml:"zendesk"`
}
