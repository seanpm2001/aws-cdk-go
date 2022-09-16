package awsecs


// The properties to import from the service using the Fargate launch type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//
//   fargateServiceAttributes := &fargateServiceAttributes{
//   	cluster: cluster,
//
//   	// the properties below are optional
//   	serviceArn: jsii.String("serviceArn"),
//   	serviceName: jsii.String("serviceName"),
//   }
//
// Experimental.
type FargateServiceAttributes struct {
	// The cluster that hosts the service.
	// Experimental.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The service ARN.
	// Experimental.
	ServiceArn *string `field:"optional" json:"serviceArn" yaml:"serviceArn"`
	// The name of the service.
	// Experimental.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}
