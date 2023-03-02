package awsec2


// Properties for defining a `CfnIPAMPoolCidr`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIPAMPoolCidrProps := &cfnIPAMPoolCidrProps{
//   	ipamPoolId: jsii.String("ipamPoolId"),
//
//   	// the properties below are optional
//   	cidr: jsii.String("cidr"),
//   	netmaskLength: jsii.Number(123),
//   }
//
type CfnIPAMPoolCidrProps struct {
	// The ID of the IPAM pool.
	IpamPoolId *string `field:"required" json:"ipamPoolId" yaml:"ipamPoolId"`
	// The CIDR provisioned to the IPAM pool.
	//
	// A CIDR is a representation of an IP address and its associated network mask (or netmask) and refers to a range of IP addresses. An IPv4 CIDR example is `10.24.34.0/23` . An IPv6 CIDR example is `2001:DB8::/32` .
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// The netmask length of the CIDR you'd like to provision to a pool.
	//
	// Can be used for provisioning Amazon-provided IPv6 CIDRs to top-level pools and for provisioning CIDRs to pools with source pools. Cannot be used to provision BYOIP CIDRs to top-level pools. "NetmaskLength" or "Cidr" is required.
	NetmaskLength *float64 `field:"optional" json:"netmaskLength" yaml:"netmaskLength"`
}

