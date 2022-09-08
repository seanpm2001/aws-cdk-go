package awsappmesh


// All Properties for Envoy Access logs for mesh endpoints.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessLogConfig := &accessLogConfig{
//   	virtualGatewayAccessLog: &virtualGatewayAccessLogProperty{
//   		file: &virtualGatewayFileAccessLogProperty{
//   			path: jsii.String("path"),
//   		},
//   	},
//   	virtualNodeAccessLog: &accessLogProperty{
//   		file: &fileAccessLogProperty{
//   			path: jsii.String("path"),
//   		},
//   	},
//   }
//
// Experimental.
type AccessLogConfig struct {
	// VirtualGateway CFN configuration for Access Logging.
	// Experimental.
	VirtualGatewayAccessLog *CfnVirtualGateway_VirtualGatewayAccessLogProperty `field:"optional" json:"virtualGatewayAccessLog" yaml:"virtualGatewayAccessLog"`
	// VirtualNode CFN configuration for Access Logging.
	// Experimental.
	VirtualNodeAccessLog *CfnVirtualNode_AccessLogProperty `field:"optional" json:"virtualNodeAccessLog" yaml:"virtualNodeAccessLog"`
}

