package awsamplifyuibuilder


// Properties for defining a `CfnComponent`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var componentChildProperty_ componentChildProperty
//   var componentPropertyProperty_ componentPropertyProperty
//   var overrides interface{}
//   var predicateProperty_ predicateProperty
//
//   cfnComponentProps := &CfnComponentProps{
//   	BindingProperties: map[string]interface{}{
//   		"bindingPropertiesKey": &ComponentBindingPropertiesValueProperty{
//   			"bindingProperties": &ComponentBindingPropertiesValuePropertiesProperty{
//   				"bucket": jsii.String("bucket"),
//   				"defaultValue": jsii.String("defaultValue"),
//   				"field": jsii.String("field"),
//   				"key": jsii.String("key"),
//   				"model": jsii.String("model"),
//   				"predicates": []interface{}{
//   					&predicateProperty{
//   						"and": []interface{}{
//   							predicateProperty_,
//   						},
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operator": jsii.String("operator"),
//   						"or": []interface{}{
//   							predicateProperty_,
//   						},
//   					},
//   				},
//   				"userAttribute": jsii.String("userAttribute"),
//   			},
//   			"defaultValue": jsii.String("defaultValue"),
//   			"type": jsii.String("type"),
//   		},
//   	},
//   	ComponentType: jsii.String("componentType"),
//   	Name: jsii.String("name"),
//   	Overrides: overrides,
//   	Properties: map[string]interface{}{
//   		"propertiesKey": &componentPropertyProperty{
//   			"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   				"property": jsii.String("property"),
//
//   				// the properties below are optional
//   				"field": jsii.String("field"),
//   			},
//   			"bindings": map[string]interface{}{
//   				"bindingsKey": &FormBindingElementProperty{
//   					"element": jsii.String("element"),
//   					"property": jsii.String("property"),
//   				},
//   			},
//   			"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   				"property": jsii.String("property"),
//
//   				// the properties below are optional
//   				"field": jsii.String("field"),
//   			},
//   			"componentName": jsii.String("componentName"),
//   			"concat": []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			"condition": &ComponentConditionPropertyProperty{
//   				"else": componentPropertyProperty_,
//   				"field": jsii.String("field"),
//   				"operand": jsii.String("operand"),
//   				"operandType": jsii.String("operandType"),
//   				"operator": jsii.String("operator"),
//   				"property": jsii.String("property"),
//   				"then": componentPropertyProperty_,
//   			},
//   			"configured": jsii.Boolean(false),
//   			"defaultValue": jsii.String("defaultValue"),
//   			"event": jsii.String("event"),
//   			"importedValue": jsii.String("importedValue"),
//   			"model": jsii.String("model"),
//   			"property": jsii.String("property"),
//   			"type": jsii.String("type"),
//   			"userAttribute": jsii.String("userAttribute"),
//   			"value": jsii.String("value"),
//   		},
//   	},
//   	Variants: []interface{}{
//   		&ComponentVariantProperty{
//   			Overrides: overrides,
//   			VariantValues: map[string]*string{
//   				"variantValuesKey": jsii.String("variantValues"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	AppId: jsii.String("appId"),
//   	Children: []interface{}{
//   		&componentChildProperty{
//   			ComponentType: jsii.String("componentType"),
//   			Name: jsii.String("name"),
//   			Properties: map[string]interface{}{
//   				"propertiesKey": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": map[string]interface{}{
//   						"bindingsKey": &FormBindingElementProperty{
//   							"element": jsii.String("element"),
//   							"property": jsii.String("property"),
//   						},
//   					},
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			Children: []interface{}{
//   				componentChildProperty_,
//   			},
//   			Events: map[string]interface{}{
//   				"eventsKey": &ComponentEventProperty{
//   					"action": jsii.String("action"),
//   					"parameters": &ActionParametersProperty{
//   						"anchor": &componentPropertyProperty{
//   							"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"bindings": map[string]interface{}{
//   								"bindingsKey": &FormBindingElementProperty{
//   									"element": jsii.String("element"),
//   									"property": jsii.String("property"),
//   								},
//   							},
//   							"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"componentName": jsii.String("componentName"),
//   							"concat": []interface{}{
//   								componentPropertyProperty_,
//   							},
//   							"condition": &ComponentConditionPropertyProperty{
//   								"else": componentPropertyProperty_,
//   								"field": jsii.String("field"),
//   								"operand": jsii.String("operand"),
//   								"operandType": jsii.String("operandType"),
//   								"operator": jsii.String("operator"),
//   								"property": jsii.String("property"),
//   								"then": componentPropertyProperty_,
//   							},
//   							"configured": jsii.Boolean(false),
//   							"defaultValue": jsii.String("defaultValue"),
//   							"event": jsii.String("event"),
//   							"importedValue": jsii.String("importedValue"),
//   							"model": jsii.String("model"),
//   							"property": jsii.String("property"),
//   							"type": jsii.String("type"),
//   							"userAttribute": jsii.String("userAttribute"),
//   							"value": jsii.String("value"),
//   						},
//   						"fields": map[string]interface{}{
//   							"fieldsKey": &componentPropertyProperty{
//   								"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   									"property": jsii.String("property"),
//
//   									// the properties below are optional
//   									"field": jsii.String("field"),
//   								},
//   								"bindings": map[string]interface{}{
//   									"bindingsKey": &FormBindingElementProperty{
//   										"element": jsii.String("element"),
//   										"property": jsii.String("property"),
//   									},
//   								},
//   								"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   									"property": jsii.String("property"),
//
//   									// the properties below are optional
//   									"field": jsii.String("field"),
//   								},
//   								"componentName": jsii.String("componentName"),
//   								"concat": []interface{}{
//   									componentPropertyProperty_,
//   								},
//   								"condition": &ComponentConditionPropertyProperty{
//   									"else": componentPropertyProperty_,
//   									"field": jsii.String("field"),
//   									"operand": jsii.String("operand"),
//   									"operandType": jsii.String("operandType"),
//   									"operator": jsii.String("operator"),
//   									"property": jsii.String("property"),
//   									"then": componentPropertyProperty_,
//   								},
//   								"configured": jsii.Boolean(false),
//   								"defaultValue": jsii.String("defaultValue"),
//   								"event": jsii.String("event"),
//   								"importedValue": jsii.String("importedValue"),
//   								"model": jsii.String("model"),
//   								"property": jsii.String("property"),
//   								"type": jsii.String("type"),
//   								"userAttribute": jsii.String("userAttribute"),
//   								"value": jsii.String("value"),
//   							},
//   						},
//   						"global": &componentPropertyProperty{
//   							"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"bindings": map[string]interface{}{
//   								"bindingsKey": &FormBindingElementProperty{
//   									"element": jsii.String("element"),
//   									"property": jsii.String("property"),
//   								},
//   							},
//   							"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"componentName": jsii.String("componentName"),
//   							"concat": []interface{}{
//   								componentPropertyProperty_,
//   							},
//   							"condition": &ComponentConditionPropertyProperty{
//   								"else": componentPropertyProperty_,
//   								"field": jsii.String("field"),
//   								"operand": jsii.String("operand"),
//   								"operandType": jsii.String("operandType"),
//   								"operator": jsii.String("operator"),
//   								"property": jsii.String("property"),
//   								"then": componentPropertyProperty_,
//   							},
//   							"configured": jsii.Boolean(false),
//   							"defaultValue": jsii.String("defaultValue"),
//   							"event": jsii.String("event"),
//   							"importedValue": jsii.String("importedValue"),
//   							"model": jsii.String("model"),
//   							"property": jsii.String("property"),
//   							"type": jsii.String("type"),
//   							"userAttribute": jsii.String("userAttribute"),
//   							"value": jsii.String("value"),
//   						},
//   						"id": &componentPropertyProperty{
//   							"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"bindings": map[string]interface{}{
//   								"bindingsKey": &FormBindingElementProperty{
//   									"element": jsii.String("element"),
//   									"property": jsii.String("property"),
//   								},
//   							},
//   							"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"componentName": jsii.String("componentName"),
//   							"concat": []interface{}{
//   								componentPropertyProperty_,
//   							},
//   							"condition": &ComponentConditionPropertyProperty{
//   								"else": componentPropertyProperty_,
//   								"field": jsii.String("field"),
//   								"operand": jsii.String("operand"),
//   								"operandType": jsii.String("operandType"),
//   								"operator": jsii.String("operator"),
//   								"property": jsii.String("property"),
//   								"then": componentPropertyProperty_,
//   							},
//   							"configured": jsii.Boolean(false),
//   							"defaultValue": jsii.String("defaultValue"),
//   							"event": jsii.String("event"),
//   							"importedValue": jsii.String("importedValue"),
//   							"model": jsii.String("model"),
//   							"property": jsii.String("property"),
//   							"type": jsii.String("type"),
//   							"userAttribute": jsii.String("userAttribute"),
//   							"value": jsii.String("value"),
//   						},
//   						"model": jsii.String("model"),
//   						"state": &MutationActionSetStateParameterProperty{
//   							"componentName": jsii.String("componentName"),
//   							"property": jsii.String("property"),
//   							"set": &componentPropertyProperty{
//   								"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   									"property": jsii.String("property"),
//
//   									// the properties below are optional
//   									"field": jsii.String("field"),
//   								},
//   								"bindings": map[string]interface{}{
//   									"bindingsKey": &FormBindingElementProperty{
//   										"element": jsii.String("element"),
//   										"property": jsii.String("property"),
//   									},
//   								},
//   								"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   									"property": jsii.String("property"),
//
//   									// the properties below are optional
//   									"field": jsii.String("field"),
//   								},
//   								"componentName": jsii.String("componentName"),
//   								"concat": []interface{}{
//   									componentPropertyProperty_,
//   								},
//   								"condition": &ComponentConditionPropertyProperty{
//   									"else": componentPropertyProperty_,
//   									"field": jsii.String("field"),
//   									"operand": jsii.String("operand"),
//   									"operandType": jsii.String("operandType"),
//   									"operator": jsii.String("operator"),
//   									"property": jsii.String("property"),
//   									"then": componentPropertyProperty_,
//   								},
//   								"configured": jsii.Boolean(false),
//   								"defaultValue": jsii.String("defaultValue"),
//   								"event": jsii.String("event"),
//   								"importedValue": jsii.String("importedValue"),
//   								"model": jsii.String("model"),
//   								"property": jsii.String("property"),
//   								"type": jsii.String("type"),
//   								"userAttribute": jsii.String("userAttribute"),
//   								"value": jsii.String("value"),
//   							},
//   						},
//   						"target": &componentPropertyProperty{
//   							"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"bindings": map[string]interface{}{
//   								"bindingsKey": &FormBindingElementProperty{
//   									"element": jsii.String("element"),
//   									"property": jsii.String("property"),
//   								},
//   							},
//   							"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"componentName": jsii.String("componentName"),
//   							"concat": []interface{}{
//   								componentPropertyProperty_,
//   							},
//   							"condition": &ComponentConditionPropertyProperty{
//   								"else": componentPropertyProperty_,
//   								"field": jsii.String("field"),
//   								"operand": jsii.String("operand"),
//   								"operandType": jsii.String("operandType"),
//   								"operator": jsii.String("operator"),
//   								"property": jsii.String("property"),
//   								"then": componentPropertyProperty_,
//   							},
//   							"configured": jsii.Boolean(false),
//   							"defaultValue": jsii.String("defaultValue"),
//   							"event": jsii.String("event"),
//   							"importedValue": jsii.String("importedValue"),
//   							"model": jsii.String("model"),
//   							"property": jsii.String("property"),
//   							"type": jsii.String("type"),
//   							"userAttribute": jsii.String("userAttribute"),
//   							"value": jsii.String("value"),
//   						},
//   						"type": &componentPropertyProperty{
//   							"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"bindings": map[string]interface{}{
//   								"bindingsKey": &FormBindingElementProperty{
//   									"element": jsii.String("element"),
//   									"property": jsii.String("property"),
//   								},
//   							},
//   							"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"componentName": jsii.String("componentName"),
//   							"concat": []interface{}{
//   								componentPropertyProperty_,
//   							},
//   							"condition": &ComponentConditionPropertyProperty{
//   								"else": componentPropertyProperty_,
//   								"field": jsii.String("field"),
//   								"operand": jsii.String("operand"),
//   								"operandType": jsii.String("operandType"),
//   								"operator": jsii.String("operator"),
//   								"property": jsii.String("property"),
//   								"then": componentPropertyProperty_,
//   							},
//   							"configured": jsii.Boolean(false),
//   							"defaultValue": jsii.String("defaultValue"),
//   							"event": jsii.String("event"),
//   							"importedValue": jsii.String("importedValue"),
//   							"model": jsii.String("model"),
//   							"property": jsii.String("property"),
//   							"type": jsii.String("type"),
//   							"userAttribute": jsii.String("userAttribute"),
//   							"value": jsii.String("value"),
//   						},
//   						"url": &componentPropertyProperty{
//   							"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"bindings": map[string]interface{}{
//   								"bindingsKey": &FormBindingElementProperty{
//   									"element": jsii.String("element"),
//   									"property": jsii.String("property"),
//   								},
//   							},
//   							"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"componentName": jsii.String("componentName"),
//   							"concat": []interface{}{
//   								componentPropertyProperty_,
//   							},
//   							"condition": &ComponentConditionPropertyProperty{
//   								"else": componentPropertyProperty_,
//   								"field": jsii.String("field"),
//   								"operand": jsii.String("operand"),
//   								"operandType": jsii.String("operandType"),
//   								"operator": jsii.String("operator"),
//   								"property": jsii.String("property"),
//   								"then": componentPropertyProperty_,
//   							},
//   							"configured": jsii.Boolean(false),
//   							"defaultValue": jsii.String("defaultValue"),
//   							"event": jsii.String("event"),
//   							"importedValue": jsii.String("importedValue"),
//   							"model": jsii.String("model"),
//   							"property": jsii.String("property"),
//   							"type": jsii.String("type"),
//   							"userAttribute": jsii.String("userAttribute"),
//   							"value": jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	CollectionProperties: map[string]interface{}{
//   		"collectionPropertiesKey": &ComponentDataConfigurationProperty{
//   			"model": jsii.String("model"),
//
//   			// the properties below are optional
//   			"identifiers": []*string{
//   				jsii.String("identifiers"),
//   			},
//   			"predicate": &predicateProperty{
//   				"and": []interface{}{
//   					predicateProperty_,
//   				},
//   				"field": jsii.String("field"),
//   				"operand": jsii.String("operand"),
//   				"operator": jsii.String("operator"),
//   				"or": []interface{}{
//   					predicateProperty_,
//   				},
//   			},
//   			"sort": []interface{}{
//   				&SortPropertyProperty{
//   					"direction": jsii.String("direction"),
//   					"field": jsii.String("field"),
//   				},
//   			},
//   		},
//   	},
//   	EnvironmentName: jsii.String("environmentName"),
//   	Events: map[string]interface{}{
//   		"eventsKey": &ComponentEventProperty{
//   			"action": jsii.String("action"),
//   			"parameters": &ActionParametersProperty{
//   				"anchor": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": map[string]interface{}{
//   						"bindingsKey": &FormBindingElementProperty{
//   							"element": jsii.String("element"),
//   							"property": jsii.String("property"),
//   						},
//   					},
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"fields": map[string]interface{}{
//   					"fieldsKey": &componentPropertyProperty{
//   						"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   							"property": jsii.String("property"),
//
//   							// the properties below are optional
//   							"field": jsii.String("field"),
//   						},
//   						"bindings": map[string]interface{}{
//   							"bindingsKey": &FormBindingElementProperty{
//   								"element": jsii.String("element"),
//   								"property": jsii.String("property"),
//   							},
//   						},
//   						"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   							"property": jsii.String("property"),
//
//   							// the properties below are optional
//   							"field": jsii.String("field"),
//   						},
//   						"componentName": jsii.String("componentName"),
//   						"concat": []interface{}{
//   							componentPropertyProperty_,
//   						},
//   						"condition": &ComponentConditionPropertyProperty{
//   							"else": componentPropertyProperty_,
//   							"field": jsii.String("field"),
//   							"operand": jsii.String("operand"),
//   							"operandType": jsii.String("operandType"),
//   							"operator": jsii.String("operator"),
//   							"property": jsii.String("property"),
//   							"then": componentPropertyProperty_,
//   						},
//   						"configured": jsii.Boolean(false),
//   						"defaultValue": jsii.String("defaultValue"),
//   						"event": jsii.String("event"),
//   						"importedValue": jsii.String("importedValue"),
//   						"model": jsii.String("model"),
//   						"property": jsii.String("property"),
//   						"type": jsii.String("type"),
//   						"userAttribute": jsii.String("userAttribute"),
//   						"value": jsii.String("value"),
//   					},
//   				},
//   				"global": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": map[string]interface{}{
//   						"bindingsKey": &FormBindingElementProperty{
//   							"element": jsii.String("element"),
//   							"property": jsii.String("property"),
//   						},
//   					},
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"id": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": map[string]interface{}{
//   						"bindingsKey": &FormBindingElementProperty{
//   							"element": jsii.String("element"),
//   							"property": jsii.String("property"),
//   						},
//   					},
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"model": jsii.String("model"),
//   				"state": &MutationActionSetStateParameterProperty{
//   					"componentName": jsii.String("componentName"),
//   					"property": jsii.String("property"),
//   					"set": &componentPropertyProperty{
//   						"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   							"property": jsii.String("property"),
//
//   							// the properties below are optional
//   							"field": jsii.String("field"),
//   						},
//   						"bindings": map[string]interface{}{
//   							"bindingsKey": &FormBindingElementProperty{
//   								"element": jsii.String("element"),
//   								"property": jsii.String("property"),
//   							},
//   						},
//   						"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   							"property": jsii.String("property"),
//
//   							// the properties below are optional
//   							"field": jsii.String("field"),
//   						},
//   						"componentName": jsii.String("componentName"),
//   						"concat": []interface{}{
//   							componentPropertyProperty_,
//   						},
//   						"condition": &ComponentConditionPropertyProperty{
//   							"else": componentPropertyProperty_,
//   							"field": jsii.String("field"),
//   							"operand": jsii.String("operand"),
//   							"operandType": jsii.String("operandType"),
//   							"operator": jsii.String("operator"),
//   							"property": jsii.String("property"),
//   							"then": componentPropertyProperty_,
//   						},
//   						"configured": jsii.Boolean(false),
//   						"defaultValue": jsii.String("defaultValue"),
//   						"event": jsii.String("event"),
//   						"importedValue": jsii.String("importedValue"),
//   						"model": jsii.String("model"),
//   						"property": jsii.String("property"),
//   						"type": jsii.String("type"),
//   						"userAttribute": jsii.String("userAttribute"),
//   						"value": jsii.String("value"),
//   					},
//   				},
//   				"target": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": map[string]interface{}{
//   						"bindingsKey": &FormBindingElementProperty{
//   							"element": jsii.String("element"),
//   							"property": jsii.String("property"),
//   						},
//   					},
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"type": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": map[string]interface{}{
//   						"bindingsKey": &FormBindingElementProperty{
//   							"element": jsii.String("element"),
//   							"property": jsii.String("property"),
//   						},
//   					},
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"url": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": map[string]interface{}{
//   						"bindingsKey": &FormBindingElementProperty{
//   							"element": jsii.String("element"),
//   							"property": jsii.String("property"),
//   						},
//   					},
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	SchemaVersion: jsii.String("schemaVersion"),
//   	SourceId: jsii.String("sourceId"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnComponentProps struct {
	// The information to connect a component's properties to data at runtime.
	//
	// You can't specify `tags` as a valid property for `bindingProperties` .
	BindingProperties interface{} `field:"required" json:"bindingProperties" yaml:"bindingProperties"`
	// The type of the component.
	//
	// This can be an Amplify custom UI component or another custom component.
	ComponentType *string `field:"required" json:"componentType" yaml:"componentType"`
	// The name of the component.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Describes the component's properties that can be overriden in a customized instance of the component.
	//
	// You can't specify `tags` as a valid property for `overrides` .
	Overrides interface{} `field:"required" json:"overrides" yaml:"overrides"`
	// Describes the component's properties.
	//
	// You can't specify `tags` as a valid property for `properties` .
	Properties interface{} `field:"required" json:"properties" yaml:"properties"`
	// A list of the component's variants.
	//
	// A variant is a unique style configuration of a main component.
	Variants interface{} `field:"required" json:"variants" yaml:"variants"`
	// The unique ID of the Amplify app associated with the component.
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
	// A list of the component's `ComponentChild` instances.
	Children interface{} `field:"optional" json:"children" yaml:"children"`
	// The data binding configuration for the component's properties.
	//
	// Use this for a collection component. You can't specify `tags` as a valid property for `collectionProperties` .
	CollectionProperties interface{} `field:"optional" json:"collectionProperties" yaml:"collectionProperties"`
	// The name of the backend environment that is a part of the Amplify app.
	EnvironmentName *string `field:"optional" json:"environmentName" yaml:"environmentName"`
	// Describes the events that can be raised on the component.
	//
	// Use for the workflow feature in Amplify Studio that allows you to bind events and actions to components.
	Events interface{} `field:"optional" json:"events" yaml:"events"`
	// The schema version of the component when it was imported.
	SchemaVersion *string `field:"optional" json:"schemaVersion" yaml:"schemaVersion"`
	// The unique ID of the component in its original source system, such as Figma.
	SourceId *string `field:"optional" json:"sourceId" yaml:"sourceId"`
	// One or more key-value pairs to use when tagging the component.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

