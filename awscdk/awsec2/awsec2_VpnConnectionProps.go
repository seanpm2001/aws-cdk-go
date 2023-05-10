package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secretValue secretValue
//   var vpc vpc
//
//   vpnConnectionProps := &VpnConnectionProps{
//   	Ip: jsii.String("ip"),
//   	Vpc: vpc,
//
//   	// the properties below are optional
//   	Asn: jsii.Number(123),
//   	StaticRoutes: []*string{
//   		jsii.String("staticRoutes"),
//   	},
//   	TunnelOptions: []vpnTunnelOption{
//   		&vpnTunnelOption{
//   			PreSharedKey: jsii.String("preSharedKey"),
//   			PreSharedKeySecret: secretValue,
//   			TunnelInsideCidr: jsii.String("tunnelInsideCidr"),
//   		},
//   	},
//   }
//
// Experimental.
type VpnConnectionProps struct {
	// The ip address of the customer gateway.
	// Experimental.
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// The ASN of the customer gateway.
	// Experimental.
	Asn *float64 `field:"optional" json:"asn" yaml:"asn"`
	// The static routes to be routed from the VPN gateway to the customer gateway.
	// Experimental.
	StaticRoutes *[]*string `field:"optional" json:"staticRoutes" yaml:"staticRoutes"`
	// The tunnel options for the VPN connection.
	//
	// At most two elements (one per tunnel).
	// Duplicates not allowed.
	// Experimental.
	TunnelOptions *[]*VpnTunnelOption `field:"optional" json:"tunnelOptions" yaml:"tunnelOptions"`
	// The VPC to connect to.
	// Experimental.
	Vpc IVpc `field:"required" json:"vpc" yaml:"vpc"`
}

