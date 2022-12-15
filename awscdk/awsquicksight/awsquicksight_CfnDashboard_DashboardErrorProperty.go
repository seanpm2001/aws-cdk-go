package awsquicksight


// Dashboard error.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashboardErrorProperty := &dashboardErrorProperty{
//   	message: jsii.String("message"),
//   	type: jsii.String("type"),
//   }
//
type CfnDashboard_DashboardErrorProperty struct {
	// Message.
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Type.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

