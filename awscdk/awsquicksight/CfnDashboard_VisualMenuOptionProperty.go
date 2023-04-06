package awsquicksight


// The menu options for a visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   visualMenuOptionProperty := &VisualMenuOptionProperty{
//   	AvailabilityStatus: jsii.String("availabilityStatus"),
//   }
//
type CfnDashboard_VisualMenuOptionProperty struct {
	// The availaiblity status of a visual's menu options.
	AvailabilityStatus *string `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
}

