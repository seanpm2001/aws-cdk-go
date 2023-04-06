package awsapplicationinsights


// The `AWS::ApplicationInsights::Application SubComponentTypeConfiguration` property type specifies the sub-component configurations for a component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subComponentTypeConfigurationProperty := &SubComponentTypeConfigurationProperty{
//   	SubComponentConfigurationDetails: &SubComponentConfigurationDetailsProperty{
//   		AlarmMetrics: []interface{}{
//   			&AlarmMetricProperty{
//   				AlarmMetricName: jsii.String("alarmMetricName"),
//   			},
//   		},
//   		Logs: []interface{}{
//   			&LogProperty{
//   				LogType: jsii.String("logType"),
//
//   				// the properties below are optional
//   				Encoding: jsii.String("encoding"),
//   				LogGroupName: jsii.String("logGroupName"),
//   				LogPath: jsii.String("logPath"),
//   				PatternSet: jsii.String("patternSet"),
//   			},
//   		},
//   		WindowsEvents: []interface{}{
//   			&WindowsEventProperty{
//   				EventLevels: []*string{
//   					jsii.String("eventLevels"),
//   				},
//   				EventName: jsii.String("eventName"),
//   				LogGroupName: jsii.String("logGroupName"),
//
//   				// the properties below are optional
//   				PatternSet: jsii.String("patternSet"),
//   			},
//   		},
//   	},
//   	SubComponentType: jsii.String("subComponentType"),
//   }
//
type CfnApplication_SubComponentTypeConfigurationProperty struct {
	// The configuration settings of the sub-components.
	SubComponentConfigurationDetails interface{} `field:"required" json:"subComponentConfigurationDetails" yaml:"subComponentConfigurationDetails"`
	// The sub-component type.
	SubComponentType *string `field:"required" json:"subComponentType" yaml:"subComponentType"`
}
