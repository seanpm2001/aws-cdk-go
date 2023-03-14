package awsec2


// Amazon Linux image properties.
//
// Example:
//   sg := ec2.SecurityGroup_FromSecurityGroupId(this, jsii.String("FsxSecurityGroup"), jsii.String("{SECURITY-GROUP-ID}"))
//   fs := fsx.LustreFileSystem_FromLustreFileSystemAttributes(this, jsii.String("FsxLustreFileSystem"), &FileSystemAttributes{
//   	DnsName: jsii.String("{FILE-SYSTEM-DNS-NAME}"),
//   	FileSystemId: jsii.String("{FILE-SYSTEM-ID}"),
//   	SecurityGroup: sg,
//   })
//
//   vpc := ec2.Vpc_FromVpcAttributes(this, jsii.String("Vpc"), &VpcAttributes{
//   	AvailabilityZones: []*string{
//   		jsii.String("us-west-2a"),
//   		jsii.String("us-west-2b"),
//   	},
//   	PublicSubnetIds: []*string{
//   		jsii.String("{US-WEST-2A-SUBNET-ID}"),
//   		jsii.String("{US-WEST-2B-SUBNET-ID}"),
//   	},
//   	VpcId: jsii.String("{VPC-ID}"),
//   })
//
//   inst := ec2.NewInstance(this, jsii.String("inst"), &InstanceProps{
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_T2, ec2.InstanceSize_LARGE),
//   	MachineImage: ec2.NewAmazonLinuxImage(&AmazonLinuxImageProps{
//   		Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX_2,
//   	}),
//   	Vpc: Vpc,
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PUBLIC,
//   	},
//   })
//
//   fs.Connections.AllowDefaultPortFrom(inst)
//
type AmazonLinuxImageProps struct {
	// Whether the AMI ID is cached to be stable between deployments.
	//
	// By default, the newest image is used on each deployment. This will cause
	// instances to be replaced whenever a new version is released, and may cause
	// downtime if there aren't enough running instances in the AutoScalingGroup
	// to reschedule the tasks on.
	//
	// If set to true, the AMI ID will be cached in `cdk.context.json` and the
	// same value will be used on future runs. Your instances will not be replaced
	// but your AMI version will grow old over time. To refresh the AMI lookup,
	// you will have to evict the value from the cache using the `cdk context`
	// command. See https://docs.aws.amazon.com/cdk/latest/guide/context.html for
	// more information.
	//
	// Can not be set to `true` in environment-agnostic stacks.
	CachedInContext *bool `field:"optional" json:"cachedInContext" yaml:"cachedInContext"`
	// CPU Type.
	CpuType AmazonLinuxCpuType `field:"optional" json:"cpuType" yaml:"cpuType"`
	// What edition of Amazon Linux to use.
	Edition AmazonLinuxEdition `field:"optional" json:"edition" yaml:"edition"`
	// What generation of Amazon Linux to use.
	Generation AmazonLinuxGeneration `field:"optional" json:"generation" yaml:"generation"`
	// What kernel version of Amazon Linux to use.
	Kernel AmazonLinuxKernel `field:"optional" json:"kernel" yaml:"kernel"`
	// What storage backed image to use.
	Storage AmazonLinuxStorage `field:"optional" json:"storage" yaml:"storage"`
	// Initial user data.
	UserData UserData `field:"optional" json:"userData" yaml:"userData"`
	// Virtualization type.
	Virtualization AmazonLinuxVirt `field:"optional" json:"virtualization" yaml:"virtualization"`
}

