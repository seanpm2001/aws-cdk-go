package awsmediapackage


// Use `encryptionContractConfiguration` to configure one or more content encryption keys for your endpoints that use SPEKE Version 2.0. The encryption contract defines the content keys used to encrypt the audio and video tracks in your stream. To configure the encryption contract, specify which audio and video encryption presets to use. For more information about these presets, see [SPEKE Version 2.0 Presets](https://docs.aws.amazon.com/mediapackage/latest/ug/drm-content-speke-v2-presets.html) .
//
// Note the following considerations when using `encryptionContractConfiguration` :
//
// - You can use `encryptionContractConfiguration` for DASH endpoints that use SPEKE Version 2.0. SPEKE Version 2.0 relies on the CPIX Version 2.3 specification.
// - You cannot combine an `UNENCRYPTED` preset with `UNENCRYPTED` or `SHARED` presets across `presetSpeke20Audio` and `presetSpeke20Video` .
// - When you use a `SHARED` preset, you must use it for both `presetSpeke20Audio` and `presetSpeke20Video` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionContractConfigurationProperty := &EncryptionContractConfigurationProperty{
//   	PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   	PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-encryptioncontractconfiguration.html
//
type CfnOriginEndpoint_EncryptionContractConfigurationProperty struct {
	// A collection of audio encryption presets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-encryptioncontractconfiguration.html#cfn-mediapackage-originendpoint-encryptioncontractconfiguration-presetspeke20audio
	//
	PresetSpeke20Audio *string `field:"required" json:"presetSpeke20Audio" yaml:"presetSpeke20Audio"`
	// A collection of video encryption presets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-encryptioncontractconfiguration.html#cfn-mediapackage-originendpoint-encryptioncontractconfiguration-presetspeke20video
	//
	PresetSpeke20Video *string `field:"required" json:"presetSpeke20Video" yaml:"presetSpeke20Video"`
}

