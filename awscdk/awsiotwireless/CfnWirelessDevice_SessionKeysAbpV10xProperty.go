package awsiotwireless


// LoRaWAN object for create APIs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sessionKeysAbpV10xProperty := &SessionKeysAbpV10xProperty{
//   	AppSKey: jsii.String("appSKey"),
//   	NwkSKey: jsii.String("nwkSKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-sessionkeysabpv10x.html
//
type CfnWirelessDevice_SessionKeysAbpV10xProperty struct {
	// The AppSKey is a secret key, which you should handle in a similar way as you would an application password.
	//
	// You can protect the AppSKey value by storing it in the AWS Secrets Manager and use the [secretsmanager](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html#dynamic-references-secretsmanager) to reference this value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-sessionkeysabpv10x.html#cfn-iotwireless-wirelessdevice-sessionkeysabpv10x-appskey
	//
	AppSKey *string `field:"required" json:"appSKey" yaml:"appSKey"`
	// The NwkSKey is a secret key, which you should handle in a similar way as you would an application password.
	//
	// You can protect the NwkSKey value by storing it in the AWS Secrets Manager and use the [secretsmanager](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html#dynamic-references-secretsmanager) to reference this value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdevice-sessionkeysabpv10x.html#cfn-iotwireless-wirelessdevice-sessionkeysabpv10x-nwkskey
	//
	NwkSKey *string `field:"required" json:"nwkSKey" yaml:"nwkSKey"`
}

