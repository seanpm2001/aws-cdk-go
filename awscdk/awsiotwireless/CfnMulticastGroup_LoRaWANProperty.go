package awsiotwireless


// The LoRaWAN information that is to be used with the multicast group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loRaWANProperty := &LoRaWANProperty{
//   	DlClass: jsii.String("dlClass"),
//   	RfRegion: jsii.String("rfRegion"),
//
//   	// the properties below are optional
//   	NumberOfDevicesInGroup: jsii.Number(123),
//   	NumberOfDevicesRequested: jsii.Number(123),
//   }
//
type CfnMulticastGroup_LoRaWANProperty struct {
	// DlClass for LoRaWAN.
	//
	// Valid values are ClassB and ClassC.
	DlClass *string `field:"required" json:"dlClass" yaml:"dlClass"`
	// The frequency band (RFRegion) value.
	RfRegion *string `field:"required" json:"rfRegion" yaml:"rfRegion"`
	// Number of devices that are associated to the multicast group.
	NumberOfDevicesInGroup *float64 `field:"optional" json:"numberOfDevicesInGroup" yaml:"numberOfDevicesInGroup"`
	// Number of devices that are requested to be associated with the multicast group.
	NumberOfDevicesRequested *float64 `field:"optional" json:"numberOfDevicesRequested" yaml:"numberOfDevicesRequested"`
}

