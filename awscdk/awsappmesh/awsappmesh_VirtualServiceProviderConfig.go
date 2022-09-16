package awsappmesh


// Properties for a VirtualService provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var mesh mesh
//
//   virtualServiceProviderConfig := &virtualServiceProviderConfig{
//   	mesh: mesh,
//
//   	// the properties below are optional
//   	virtualNodeProvider: &virtualNodeServiceProviderProperty{
//   		virtualNodeName: jsii.String("virtualNodeName"),
//   	},
//   	virtualRouterProvider: &virtualRouterServiceProviderProperty{
//   		virtualRouterName: jsii.String("virtualRouterName"),
//   	},
//   }
//
// Experimental.
type VirtualServiceProviderConfig struct {
	// Mesh the Provider is using.
	// Experimental.
	Mesh IMesh `field:"required" json:"mesh" yaml:"mesh"`
	// Virtual Node based provider.
	// Experimental.
	VirtualNodeProvider *CfnVirtualService_VirtualNodeServiceProviderProperty `field:"optional" json:"virtualNodeProvider" yaml:"virtualNodeProvider"`
	// Virtual Router based provider.
	// Experimental.
	VirtualRouterProvider *CfnVirtualService_VirtualRouterServiceProviderProperty `field:"optional" json:"virtualRouterProvider" yaml:"virtualRouterProvider"`
}
