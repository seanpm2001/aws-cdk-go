package awskinesisanalytics


// Property key-value pairs passed into an application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   propertyGroupProperty := &PropertyGroupProperty{
//   	PropertyGroupId: jsii.String("propertyGroupId"),
//   	PropertyMap: map[string]*string{
//   		"propertyMapKey": jsii.String("propertyMap"),
//   	},
//   }
//
type CfnApplicationV2_PropertyGroupProperty struct {
	// Describes the key of an application execution property key-value pair.
	PropertyGroupId *string `field:"optional" json:"propertyGroupId" yaml:"propertyGroupId"`
	// Describes the value of an application execution property key-value pair.
	PropertyMap interface{} `field:"optional" json:"propertyMap" yaml:"propertyMap"`
}

