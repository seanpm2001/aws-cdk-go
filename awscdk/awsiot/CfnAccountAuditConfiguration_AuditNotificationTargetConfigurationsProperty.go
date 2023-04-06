package awsiot


// The configuration of the audit notification target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   auditNotificationTargetConfigurationsProperty := &AuditNotificationTargetConfigurationsProperty{
//   	Sns: &AuditNotificationTargetProperty{
//   		Enabled: jsii.Boolean(false),
//   		RoleArn: jsii.String("roleArn"),
//   		TargetArn: jsii.String("targetArn"),
//   	},
//   }
//
type CfnAccountAuditConfiguration_AuditNotificationTargetConfigurationsProperty struct {
	// The `Sns` notification target.
	Sns interface{} `field:"optional" json:"sns" yaml:"sns"`
}
