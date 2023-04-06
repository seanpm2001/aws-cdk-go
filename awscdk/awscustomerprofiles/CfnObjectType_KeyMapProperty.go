package awscustomerprofiles


// A unique key map that can be used to map data to the profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyMapProperty := &KeyMapProperty{
//   	Name: jsii.String("name"),
//   	ObjectTypeKeyList: []interface{}{
//   		&ObjectTypeKeyProperty{
//   			FieldNames: []*string{
//   				jsii.String("fieldNames"),
//   			},
//   			StandardIdentifiers: []*string{
//   				jsii.String("standardIdentifiers"),
//   			},
//   		},
//   	},
//   }
//
type CfnObjectType_KeyMapProperty struct {
	// Name of the key.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of ObjectTypeKey.
	ObjectTypeKeyList interface{} `field:"optional" json:"objectTypeKeyList" yaml:"objectTypeKeyList"`
}
