package awskinesisanalytics


// The information required to specify a Maven reference.
//
// You can use Maven references to specify dependency JAR files.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mavenReferenceProperty := &MavenReferenceProperty{
//   	ArtifactId: jsii.String("artifactId"),
//   	GroupId: jsii.String("groupId"),
//   	Version: jsii.String("version"),
//   }
//
type CfnApplicationV2_MavenReferenceProperty struct {
	// The artifact ID of the Maven reference.
	ArtifactId *string `field:"required" json:"artifactId" yaml:"artifactId"`
	// The group ID of the Maven reference.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// The version of the Maven reference.
	Version *string `field:"required" json:"version" yaml:"version"`
}
