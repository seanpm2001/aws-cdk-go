package awsses


// Properties for defining a `CfnVdmAttributes`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVdmAttributesProps := &CfnVdmAttributesProps{
//   	DashboardAttributes: &DashboardAttributesProperty{
//   		EngagementMetrics: jsii.String("engagementMetrics"),
//   	},
//   	GuardianAttributes: &GuardianAttributesProperty{
//   		OptimizedSharedDelivery: jsii.String("optimizedSharedDelivery"),
//   	},
//   }
//
type CfnVdmAttributesProps struct {
	// Specifies additional settings for your VDM configuration as applicable to the Dashboard.
	DashboardAttributes interface{} `field:"optional" json:"dashboardAttributes" yaml:"dashboardAttributes"`
	// Specifies additional settings for your VDM configuration as applicable to the Guardian.
	GuardianAttributes interface{} `field:"optional" json:"guardianAttributes" yaml:"guardianAttributes"`
}

