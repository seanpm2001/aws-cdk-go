package cxapi


// Construction properties for CloudAssemblyBuilder.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cloudAssemblyBuilder cloudAssemblyBuilder
//
//   cloudAssemblyBuilderProps := &CloudAssemblyBuilderProps{
//   	AssetOutdir: jsii.String("assetOutdir"),
//   	ParentBuilder: cloudAssemblyBuilder,
//   }
//
// Experimental.
type CloudAssemblyBuilderProps struct {
	// Use the given asset output directory.
	// Experimental.
	AssetOutdir *string `field:"optional" json:"assetOutdir" yaml:"assetOutdir"`
	// If this builder is for a nested assembly, the parent assembly builder.
	// Experimental.
	ParentBuilder CloudAssemblyBuilder `field:"optional" json:"parentBuilder" yaml:"parentBuilder"`
}

