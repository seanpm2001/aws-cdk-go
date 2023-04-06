package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityContextProperty := &SecurityContextProperty{
//   	Privileged: jsii.Boolean(false),
//   	ReadOnlyRootFilesystem: jsii.Boolean(false),
//   	RunAsGroup: jsii.Number(123),
//   	RunAsNonRoot: jsii.Boolean(false),
//   	RunAsUser: jsii.Number(123),
//   }
//
type CfnJobDefinition_SecurityContextProperty struct {
	// `CfnJobDefinition.SecurityContextProperty.Privileged`.
	Privileged interface{} `field:"optional" json:"privileged" yaml:"privileged"`
	// `CfnJobDefinition.SecurityContextProperty.ReadOnlyRootFilesystem`.
	ReadOnlyRootFilesystem interface{} `field:"optional" json:"readOnlyRootFilesystem" yaml:"readOnlyRootFilesystem"`
	// `CfnJobDefinition.SecurityContextProperty.RunAsGroup`.
	RunAsGroup *float64 `field:"optional" json:"runAsGroup" yaml:"runAsGroup"`
	// `CfnJobDefinition.SecurityContextProperty.RunAsNonRoot`.
	RunAsNonRoot interface{} `field:"optional" json:"runAsNonRoot" yaml:"runAsNonRoot"`
	// `CfnJobDefinition.SecurityContextProperty.RunAsUser`.
	RunAsUser *float64 `field:"optional" json:"runAsUser" yaml:"runAsUser"`
}

