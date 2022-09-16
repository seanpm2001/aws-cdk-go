package awsec2


// Direction of traffic to allow all by default.
//
// Example:
//   var instanceType instanceType
//
//
//   provider := ec2.natProvider.instance(&natInstanceProps{
//   	instanceType: instanceType,
//   	defaultAllowedTraffic: ec2.natTrafficDirection_OUTBOUND_ONLY,
//   })
//   ec2.NewVpc(this, jsii.String("TheVPC"), &vpcProps{
//   	natGatewayProvider: provider,
//   })
//   provider.connections.allowFrom(ec2.peer.ipv4(jsii.String("1.2.3.4/8")), ec2.port.tcp(jsii.Number(80)))
//
// Experimental.
type NatTrafficDirection string

const (
	// Allow all outbound traffic and disallow all inbound traffic.
	// Experimental.
	NatTrafficDirection_OUTBOUND_ONLY NatTrafficDirection = "OUTBOUND_ONLY"
	// Allow all outbound and inbound traffic.
	// Experimental.
	NatTrafficDirection_INBOUND_AND_OUTBOUND NatTrafficDirection = "INBOUND_AND_OUTBOUND"
	// Disallow all outbound and inbound traffic.
	// Experimental.
	NatTrafficDirection_NONE NatTrafficDirection = "NONE"
)
