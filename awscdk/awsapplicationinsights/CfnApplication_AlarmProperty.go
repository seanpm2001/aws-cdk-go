package awsapplicationinsights


// The `AWS::ApplicationInsights::Application Alarm` property type defines a CloudWatch alarm to be monitored for the component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alarmProperty := &AlarmProperty{
//   	AlarmName: jsii.String("alarmName"),
//
//   	// the properties below are optional
//   	Severity: jsii.String("severity"),
//   }
//
type CfnApplication_AlarmProperty struct {
	// The name of the CloudWatch alarm to be monitored for the component.
	AlarmName *string `field:"required" json:"alarmName" yaml:"alarmName"`
	// Indicates the degree of outage when the alarm goes off.
	Severity *string `field:"optional" json:"severity" yaml:"severity"`
}
