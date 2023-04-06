package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   legacyCustomOriginProperty := &LegacyCustomOriginProperty{
//   	DnsName: jsii.String("dnsName"),
//   	OriginProtocolPolicy: jsii.String("originProtocolPolicy"),
//   	OriginSslProtocols: []*string{
//   		jsii.String("originSslProtocols"),
//   	},
//
//   	// the properties below are optional
//   	HttpPort: jsii.Number(123),
//   	HttpsPort: jsii.Number(123),
//   }
//
type CfnDistribution_LegacyCustomOriginProperty struct {
	// `CfnDistribution.LegacyCustomOriginProperty.DNSName`.
	DnsName *string `field:"required" json:"dnsName" yaml:"dnsName"`
	// `CfnDistribution.LegacyCustomOriginProperty.OriginProtocolPolicy`.
	OriginProtocolPolicy *string `field:"required" json:"originProtocolPolicy" yaml:"originProtocolPolicy"`
	// `CfnDistribution.LegacyCustomOriginProperty.OriginSSLProtocols`.
	OriginSslProtocols *[]*string `field:"required" json:"originSslProtocols" yaml:"originSslProtocols"`
	// `CfnDistribution.LegacyCustomOriginProperty.HTTPPort`.
	HttpPort *float64 `field:"optional" json:"httpPort" yaml:"httpPort"`
	// `CfnDistribution.LegacyCustomOriginProperty.HTTPSPort`.
	HttpsPort *float64 `field:"optional" json:"httpsPort" yaml:"httpsPort"`
}
