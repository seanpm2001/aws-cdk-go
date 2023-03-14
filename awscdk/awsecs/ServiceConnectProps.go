package awsecs


// Interface for Service Connect configuration.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var container containerDefinition
//
//
//   container.AddPortMappings(&PortMapping{
//   	Name: jsii.String("api"),
//   	ContainerPort: jsii.Number(8080),
//   })
//
//   taskDefinition.AddContainer(container)
//
//   cluster.AddDefaultCloudMapNamespace(&CloudMapNamespaceOptions{
//   	Name: jsii.String("local"),
//   })
//
//   service := ecs.NewFargateService(this, jsii.String("Service"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	ServiceConnectConfiguration: &ServiceConnectProps{
//   		Services: []serviceConnectService{
//   			&serviceConnectService{
//   				PortMappingName: jsii.String("api"),
//   				DnsName: jsii.String("http-api"),
//   				Port: jsii.Number(80),
//   			},
//   		},
//   	},
//   })
//
type ServiceConnectProps struct {
	// The log driver configuration to use for the Service Connect agent logs.
	LogDriver LogDriver `field:"optional" json:"logDriver" yaml:"logDriver"`
	// The cloudmap namespace to register this service into.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The list of Services, including a port mapping, terse client alias, and optional intermediate DNS name.
	//
	// This property may be left blank if the current ECS service does not need to advertise any ports via Service Connect.
	Services *[]*ServiceConnectService `field:"optional" json:"services" yaml:"services"`
}

