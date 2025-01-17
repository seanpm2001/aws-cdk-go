package awsses


// Properties for a dedicated IP pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dedicatedIpPoolProps := &DedicatedIpPoolProps{
//   	DedicatedIpPoolName: jsii.String("dedicatedIpPoolName"),
//   }
//
type DedicatedIpPoolProps struct {
	// A name for the dedicated IP pool.
	// Default: - a CloudFormation generated name.
	//
	DedicatedIpPoolName *string `field:"optional" json:"dedicatedIpPoolName" yaml:"dedicatedIpPoolName"`
}

