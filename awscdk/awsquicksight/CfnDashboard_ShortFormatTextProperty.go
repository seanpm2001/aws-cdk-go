package awsquicksight


// The text format for the title.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   shortFormatTextProperty := &ShortFormatTextProperty{
//   	PlainText: jsii.String("plainText"),
//   	RichText: jsii.String("richText"),
//   }
//
type CfnDashboard_ShortFormatTextProperty struct {
	// Plain text format.
	PlainText *string `field:"optional" json:"plainText" yaml:"plainText"`
	// Rich text.
	//
	// Examples of rich text include bold, underline, and italics.
	RichText *string `field:"optional" json:"richText" yaml:"richText"`
}

