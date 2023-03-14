package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessLogSettingProperty := &AccessLogSettingProperty{
//   	DestinationArn: jsii.String("destinationArn"),
//   	Format: jsii.String("format"),
//   }
//
type CfnHttpApi_AccessLogSettingProperty struct {
	// `CfnHttpApi.AccessLogSettingProperty.DestinationArn`.
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// `CfnHttpApi.AccessLogSettingProperty.Format`.
	Format *string `field:"optional" json:"format" yaml:"format"`
}

