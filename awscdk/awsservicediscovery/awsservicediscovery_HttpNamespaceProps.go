package awsservicediscovery


// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("aws-servicediscovery-integ"))
//
//   namespace := servicediscovery.NewHttpNamespace(stack, jsii.String("MyNamespace"), &HttpNamespaceProps{
//   	Name: jsii.String("covfefe"),
//   })
//
//   service1 := namespace.CreateService(jsii.String("NonIpService"), &BaseServiceProps{
//   	Description: jsii.String("service registering non-ip instances"),
//   })
//
//   service1.RegisterNonIpInstance(jsii.String("NonIpInstance"), &NonIpInstanceBaseProps{
//   	CustomAttributes: map[string]*string{
//   		"arn": jsii.String("arn:aws:s3:::mybucket"),
//   	},
//   })
//
//   service2 := namespace.CreateService(jsii.String("IpService"), &BaseServiceProps{
//   	Description: jsii.String("service registering ip instances"),
//   	HealthCheck: &HealthCheckConfig{
//   		Type: servicediscovery.HealthCheckType_HTTP,
//   		ResourcePath: jsii.String("/check"),
//   	},
//   })
//
//   service2.RegisterIpInstance(jsii.String("IpInstance"), &IpInstanceBaseProps{
//   	Ipv4: jsii.String("54.239.25.192"),
//   })
//
//   app.Synth()
//
// Experimental.
type HttpNamespaceProps struct {
	// A name for the Namespace.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the Namespace.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

