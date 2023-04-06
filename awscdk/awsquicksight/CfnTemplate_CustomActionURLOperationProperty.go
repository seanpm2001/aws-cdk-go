package awsquicksight


// The URL operation that opens a link to another webpage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customActionURLOperationProperty := &CustomActionURLOperationProperty{
//   	UrlTarget: jsii.String("urlTarget"),
//   	UrlTemplate: jsii.String("urlTemplate"),
//   }
//
type CfnTemplate_CustomActionURLOperationProperty struct {
	// The target of the `CustomActionURLOperation` .
	//
	// Valid values are defined as follows:
	//
	// - `NEW_TAB` : Opens the target URL in a new browser tab.
	// - `NEW_WINDOW` : Opens the target URL in a new browser window.
	// - `SAME_TAB` : Opens the target URL in the same browser tab.
	UrlTarget *string `field:"required" json:"urlTarget" yaml:"urlTarget"`
	// THe URL link of the `CustomActionURLOperation` .
	UrlTemplate *string `field:"required" json:"urlTemplate" yaml:"urlTemplate"`
}

