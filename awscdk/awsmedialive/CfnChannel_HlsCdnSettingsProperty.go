package awsmedialive


// The settings for the CDN of an HLS output.
//
// The parent of this entity is HlsGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsCdnSettingsProperty := &HlsCdnSettingsProperty{
//   	HlsAkamaiSettings: &HlsAkamaiSettingsProperty{
//   		ConnectionRetryInterval: jsii.Number(123),
//   		FilecacheDuration: jsii.Number(123),
//   		HttpTransferMode: jsii.String("httpTransferMode"),
//   		NumRetries: jsii.Number(123),
//   		RestartDelay: jsii.Number(123),
//   		Salt: jsii.String("salt"),
//   		Token: jsii.String("token"),
//   	},
//   	HlsBasicPutSettings: &HlsBasicPutSettingsProperty{
//   		ConnectionRetryInterval: jsii.Number(123),
//   		FilecacheDuration: jsii.Number(123),
//   		NumRetries: jsii.Number(123),
//   		RestartDelay: jsii.Number(123),
//   	},
//   	HlsMediaStoreSettings: &HlsMediaStoreSettingsProperty{
//   		ConnectionRetryInterval: jsii.Number(123),
//   		FilecacheDuration: jsii.Number(123),
//   		MediaStoreStorageClass: jsii.String("mediaStoreStorageClass"),
//   		NumRetries: jsii.Number(123),
//   		RestartDelay: jsii.Number(123),
//   	},
//   	HlsS3Settings: &HlsS3SettingsProperty{
//   		CannedAcl: jsii.String("cannedAcl"),
//   	},
//   	HlsWebdavSettings: &HlsWebdavSettingsProperty{
//   		ConnectionRetryInterval: jsii.Number(123),
//   		FilecacheDuration: jsii.Number(123),
//   		HttpTransferMode: jsii.String("httpTransferMode"),
//   		NumRetries: jsii.Number(123),
//   		RestartDelay: jsii.Number(123),
//   	},
//   }
//
type CfnChannel_HlsCdnSettingsProperty struct {
	// Sets up Akamai as the downstream system for the HLS output group.
	HlsAkamaiSettings interface{} `field:"optional" json:"hlsAkamaiSettings" yaml:"hlsAkamaiSettings"`
	// The settings for Basic Put for the HLS output.
	HlsBasicPutSettings interface{} `field:"optional" json:"hlsBasicPutSettings" yaml:"hlsBasicPutSettings"`
	// Sets up MediaStore as the destination for the HLS output.
	HlsMediaStoreSettings interface{} `field:"optional" json:"hlsMediaStoreSettings" yaml:"hlsMediaStoreSettings"`
	// Sets up Amazon S3 as the destination for this HLS output.
	HlsS3Settings interface{} `field:"optional" json:"hlsS3Settings" yaml:"hlsS3Settings"`
	// The settings for Web VTT captions in the HLS output group.
	//
	// The parent of this entity is HlsGroupSettings.
	HlsWebdavSettings interface{} `field:"optional" json:"hlsWebdavSettings" yaml:"hlsWebdavSettings"`
}
