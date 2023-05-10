package awssigner


// A reference to a Signing Profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   signingProfileAttributes := &SigningProfileAttributes{
//   	SigningProfileName: jsii.String("signingProfileName"),
//   	SigningProfileVersion: jsii.String("signingProfileVersion"),
//   }
//
// Experimental.
type SigningProfileAttributes struct {
	// The name of signing profile.
	// Experimental.
	SigningProfileName *string `field:"required" json:"signingProfileName" yaml:"signingProfileName"`
	// The version of signing profile.
	// Experimental.
	SigningProfileVersion *string `field:"required" json:"signingProfileVersion" yaml:"signingProfileVersion"`
}

