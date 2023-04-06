package awsec2


// Options for creating a SystemD configuration file.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var vpc vpc
//   var instanceType instanceType
//
//
//   ec2.NewInstance(this, jsii.String("Instance"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux(&AmazonLinuxImageProps{
//   		// Amazon Linux 2 uses SystemD
//   		Generation: ec2.AmazonLinuxGeneration,
//   		AMAZON_LINUX_2: AMAZON_LINUX_2,
//   	}),
//
//   	Init: ec2.CloudFormationInit_FromElements([]interface{}{
//   		ec2.InitService_SystemdConfigFile(jsii.String("simpleserver"), &SystemdConfigFileOptions{
//   			Command: jsii.String("/usr/bin/python3 -m http.server 8080"),
//   			Cwd: jsii.String("/var/www/html"),
//   		}),
//   		ec2.InitService_Enable(jsii.String("simpleserver"), &InitServiceOptions{
//   			ServiceManager: ec2.ServiceManager_SYSTEMD,
//   		}),
//   		ec2.InitFile_FromString(jsii.String("/var/www/html/index.html"), jsii.String("Hello! It's working!")),
//   	}),
//   })
//
type SystemdConfigFileOptions struct {
	// The command to run to start this service.
	Command *string `field:"required" json:"command" yaml:"command"`
	// Start the service after the networking part of the OS comes up.
	AfterNetwork *bool `field:"optional" json:"afterNetwork" yaml:"afterNetwork"`
	// The working directory for the command.
	Cwd *string `field:"optional" json:"cwd" yaml:"cwd"`
	// A description of this service.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The group to execute the process under.
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Keep the process running all the time.
	//
	// Restarts the process when it exits for any reason other
	// than the machine shutting down.
	KeepRunning *bool `field:"optional" json:"keepRunning" yaml:"keepRunning"`
	// The user to execute the process under.
	User *string `field:"optional" json:"user" yaml:"user"`
}
