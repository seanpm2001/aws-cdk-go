package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnEIP`.
//
// Example:
//   var instance instance
//
//   var myZone hostedZone
//
//
//   elasticIp := ec2.NewCfnEIP(this, jsii.String("EIP"), &cfnEIPProps{
//   	domain: jsii.String("vpc"),
//   	instanceId: instance.instanceId,
//   })
//   route53.NewARecord(this, jsii.String("ARecord"), &aRecordProps{
//   	zone: myZone,
//   	target: route53.recordTarget.fromIpAddresses(elasticIp.ref),
//   })
//
type CfnEIPProps struct {
	// Indicates whether the Elastic IP address is for use with instances in a VPC or instance in EC2-Classic.
	//
	// Default: If the Region supports EC2-Classic, the default is `standard` . Otherwise, the default is `vpc` .
	//
	// Use when allocating an address for use with a VPC if the Region supports EC2-Classic.
	//
	// If you define an Elastic IP address and associate it with a VPC that is defined in the same template, you must declare a dependency on the VPC-gateway attachment by using the [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) on this resource.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The ID of the instance.
	//
	// > Updates to the `InstanceId` property may require *some interruptions* . Updates on an EIP reassociates the address on its associated resource.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// A unique set of Availability Zones, Local Zones, or Wavelength Zones from which AWS advertises IP addresses.
	//
	// Use this parameter to limit the IP address to this location. IP addresses cannot move between network border groups.
	//
	// Use [DescribeAvailabilityZones](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAvailabilityZones.html) to view the network border groups.
	//
	// You cannot use a network border group with EC2 Classic. If you attempt this operation on EC2 Classic, you receive an `InvalidParameterCombination` error.
	NetworkBorderGroup *string `field:"optional" json:"networkBorderGroup" yaml:"networkBorderGroup"`
	// The ID of an address pool that you own.
	//
	// Use this parameter to let Amazon EC2 select an address from the address pool.
	//
	// > Updates to the `PublicIpv4Pool` property may require *some interruptions* . Updates on an EIP reassociates the address on its associated resource.
	PublicIpv4Pool *string `field:"optional" json:"publicIpv4Pool" yaml:"publicIpv4Pool"`
	// Any tags assigned to the Elastic IP address.
	//
	// > Updates to the `Tags` property may require *some interruptions* . Updates on an EIP reassociates the address on its associated resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The Elastic IP address you are accepting for transfer.
	TransferAddress *string `field:"optional" json:"transferAddress" yaml:"transferAddress"`
}

