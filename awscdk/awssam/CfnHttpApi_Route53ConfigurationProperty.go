package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   route53ConfigurationProperty := &Route53ConfigurationProperty{
//   	DistributedDomainName: jsii.String("distributedDomainName"),
//   	EvaluateTargetHealth: jsii.Boolean(false),
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   	HostedZoneName: jsii.String("hostedZoneName"),
//   	IpV6: jsii.Boolean(false),
//   }
//
type CfnHttpApi_Route53ConfigurationProperty struct {
	// `CfnHttpApi.Route53ConfigurationProperty.DistributedDomainName`.
	DistributedDomainName *string `field:"optional" json:"distributedDomainName" yaml:"distributedDomainName"`
	// `CfnHttpApi.Route53ConfigurationProperty.EvaluateTargetHealth`.
	EvaluateTargetHealth interface{} `field:"optional" json:"evaluateTargetHealth" yaml:"evaluateTargetHealth"`
	// `CfnHttpApi.Route53ConfigurationProperty.HostedZoneId`.
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
	// `CfnHttpApi.Route53ConfigurationProperty.HostedZoneName`.
	HostedZoneName *string `field:"optional" json:"hostedZoneName" yaml:"hostedZoneName"`
	// `CfnHttpApi.Route53ConfigurationProperty.IpV6`.
	IpV6 interface{} `field:"optional" json:"ipV6" yaml:"ipV6"`
}

