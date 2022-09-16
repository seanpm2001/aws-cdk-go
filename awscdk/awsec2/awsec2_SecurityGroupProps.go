package awsec2


// Example:
//   var vpc vpc
//
//
//   securityGroup1 := ec2.NewSecurityGroup(this, jsii.String("SecurityGroup1"), &securityGroupProps{
//   	vpc: vpc,
//   })
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   	securityGroup: securityGroup1,
//   })
//
//   securityGroup2 := ec2.NewSecurityGroup(this, jsii.String("SecurityGroup2"), &securityGroupProps{
//   	vpc: vpc,
//   })
//   lb.addSecurityGroup(securityGroup2)
//
// Experimental.
type SecurityGroupProps struct {
	// The VPC in which to create the security group.
	// Experimental.
	Vpc IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Whether to allow all outbound traffic by default.
	//
	// If this is set to true, there will only be a single egress rule which allows all
	// outbound traffic. If this is set to false, no outbound traffic will be allowed by
	// default and all egress traffic must be explicitly authorized.
	// Experimental.
	AllowAllOutbound *bool `field:"optional" json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// A description of the security group.
	// Experimental.
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
	// Experimental.
	DisableInlineRules *bool `field:"optional" json:"disableInlineRules" yaml:"disableInlineRules"`
	// The name of the security group.
	//
	// For valid values, see the GroupName
	// parameter of the CreateSecurityGroup action in the Amazon EC2 API
	// Reference.
	//
	// It is not recommended to use an explicit group name.
	// Experimental.
	SecurityGroupName *string `field:"optional" json:"securityGroupName" yaml:"securityGroupName"`
}
