package awsec2


// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var vpc vpc
//
//
//   template := ec2.NewLaunchTemplate(this, jsii.String("LaunchTemplate"), &LaunchTemplateProps{
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux(&AmazonLinuxImageProps{
//   		Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX_2,
//   	}),
//   	SecurityGroup: ec2.NewSecurityGroup(this, jsii.String("LaunchTemplateSG"), &SecurityGroupProps{
//   		Vpc: vpc,
//   	}),
//   })
//
type SecurityGroupProps struct {
	// The VPC in which to create the security group.
	Vpc IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Whether to allow all outbound ipv6 traffic by default.
	//
	// If this is set to true, there will only be a single egress rule which allows all
	// outbound ipv6 traffic. If this is set to false, no outbound traffic will be allowed by
	// default and all egress ipv6 traffic must be explicitly authorized.
	//
	// To allow all ipv4 traffic use allowAllOutbound.
	AllowAllIpv6Outbound *bool `field:"optional" json:"allowAllIpv6Outbound" yaml:"allowAllIpv6Outbound"`
	// Whether to allow all outbound traffic by default.
	//
	// If this is set to true, there will only be a single egress rule which allows all
	// outbound traffic. If this is set to false, no outbound traffic will be allowed by
	// default and all egress traffic must be explicitly authorized.
	//
	// To allow all ipv6 traffic use allowAllIpv6Outbound.
	AllowAllOutbound *bool `field:"optional" json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// A description of the security group.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether to disable inline ingress and egress rule optimization.
	//
	// If this is set to true, ingress and egress rules will not be declared under the
	// SecurityGroup in cloudformation, but will be separate elements.
	//
	// Inlining rules is an optimization for producing smaller stack templates. Sometimes
	// this is not desirable, for example when security group access is managed via tags.
	//
	// The default value can be overriden globally by setting the context variable
	// '@aws-cdk/aws-ec2.securityGroupDisableInlineRules'.
	DisableInlineRules *bool `field:"optional" json:"disableInlineRules" yaml:"disableInlineRules"`
	// The name of the security group.
	//
	// For valid values, see the GroupName
	// parameter of the CreateSecurityGroup action in the Amazon EC2 API
	// Reference.
	//
	// It is not recommended to use an explicit group name.
	SecurityGroupName *string `field:"optional" json:"securityGroupName" yaml:"securityGroupName"`
}
