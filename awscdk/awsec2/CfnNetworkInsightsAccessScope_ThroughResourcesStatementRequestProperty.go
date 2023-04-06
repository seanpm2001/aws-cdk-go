package awsec2


// Describes a through resource statement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   throughResourcesStatementRequestProperty := &ThroughResourcesStatementRequestProperty{
//   	ResourceStatement: &ResourceStatementRequestProperty{
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		ResourceTypes: []*string{
//   			jsii.String("resourceTypes"),
//   		},
//   	},
//   }
//
type CfnNetworkInsightsAccessScope_ThroughResourcesStatementRequestProperty struct {
	// The resource statement.
	ResourceStatement interface{} `field:"optional" json:"resourceStatement" yaml:"resourceStatement"`
}

