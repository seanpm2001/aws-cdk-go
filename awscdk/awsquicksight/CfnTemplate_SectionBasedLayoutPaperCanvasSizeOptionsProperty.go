package awsquicksight


// The options for a paper canvas of a section-based layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sectionBasedLayoutPaperCanvasSizeOptionsProperty := &SectionBasedLayoutPaperCanvasSizeOptionsProperty{
//   	PaperMargin: &SpacingProperty{
//   		Bottom: jsii.String("bottom"),
//   		Left: jsii.String("left"),
//   		Right: jsii.String("right"),
//   		Top: jsii.String("top"),
//   	},
//   	PaperOrientation: jsii.String("paperOrientation"),
//   	PaperSize: jsii.String("paperSize"),
//   }
//
type CfnTemplate_SectionBasedLayoutPaperCanvasSizeOptionsProperty struct {
	// Defines the spacing between the canvas content and the top, bottom, left, and right edges.
	PaperMargin interface{} `field:"optional" json:"paperMargin" yaml:"paperMargin"`
	// The paper orientation that is used to define canvas dimensions. Choose one of the following options:.
	//
	// - PORTRAIT
	// - LANDSCAPE.
	PaperOrientation *string `field:"optional" json:"paperOrientation" yaml:"paperOrientation"`
	// The paper size that is used to define canvas dimensions.
	PaperSize *string `field:"optional" json:"paperSize" yaml:"paperSize"`
}
