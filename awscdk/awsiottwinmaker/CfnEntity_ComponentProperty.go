package awsiottwinmaker


// The entity component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataValueProperty_ dataValueProperty
//   var definition interface{}
//   var error interface{}
//   var relationshipValue interface{}
//
//   componentProperty := &ComponentProperty{
//   	ComponentName: jsii.String("componentName"),
//   	ComponentTypeId: jsii.String("componentTypeId"),
//   	DefinedIn: jsii.String("definedIn"),
//   	Description: jsii.String("description"),
//   	Properties: map[string]interface{}{
//   		"propertiesKey": &PropertyProperty{
//   			"definition": definition,
//   			"value": &dataValueProperty{
//   				"booleanValue": jsii.Boolean(false),
//   				"doubleValue": jsii.Number(123),
//   				"expression": jsii.String("expression"),
//   				"integerValue": jsii.Number(123),
//   				"listValue": []interface{}{
//   					dataValueProperty_,
//   				},
//   				"longValue": jsii.Number(123),
//   				"mapValue": map[string]interface{}{
//   					"mapValueKey": dataValueProperty_,
//   				},
//   				"relationshipValue": relationshipValue,
//   				"stringValue": jsii.String("stringValue"),
//   			},
//   		},
//   	},
//   	PropertyGroups: map[string]interface{}{
//   		"propertyGroupsKey": &PropertyGroupProperty{
//   			"groupType": jsii.String("groupType"),
//   			"propertyNames": []*string{
//   				jsii.String("propertyNames"),
//   			},
//   		},
//   	},
//   	Status: &StatusProperty{
//   		Error: error,
//   		State: jsii.String("state"),
//   	},
//   }
//
type CfnEntity_ComponentProperty struct {
	// The name of the component.
	ComponentName *string `field:"optional" json:"componentName" yaml:"componentName"`
	// The ID of the ComponentType.
	ComponentTypeId *string `field:"optional" json:"componentTypeId" yaml:"componentTypeId"`
	// The name of the property definition set in the request.
	DefinedIn *string `field:"optional" json:"definedIn" yaml:"definedIn"`
	// The description of the component.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An object that maps strings to the properties to set in the component type.
	//
	// Each string in the mapping must be unique to this object.
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// An object that maps strings to the property groups in the component type.
	//
	// Each string in the mapping must be unique to this object.
	PropertyGroups interface{} `field:"optional" json:"propertyGroups" yaml:"propertyGroups"`
	// The status of the component.
	Status interface{} `field:"optional" json:"status" yaml:"status"`
}

