package awsquicksight


// The top movers and bottom movers computation setup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topBottomMoversComputationProperty := &TopBottomMoversComputationProperty{
//   	Category: &DimensionFieldProperty{
//   		CategoricalDimensionField: &CategoricalDimensionFieldProperty{
//   			Column: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			FieldId: jsii.String("fieldId"),
//
//   			// the properties below are optional
//   			FormatConfiguration: &StringFormatConfigurationProperty{
//   				NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   					NullString: jsii.String("nullString"),
//   				},
//   				NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   					CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   						Symbol: jsii.String("symbol"),
//   					},
//   					NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   					PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   				},
//   			},
//   			HierarchyId: jsii.String("hierarchyId"),
//   		},
//   		DateDimensionField: &DateDimensionFieldProperty{
//   			Column: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			FieldId: jsii.String("fieldId"),
//
//   			// the properties below are optional
//   			DateGranularity: jsii.String("dateGranularity"),
//   			FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   				DateTimeFormat: jsii.String("dateTimeFormat"),
//   				NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   					NullString: jsii.String("nullString"),
//   				},
//   				NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   					CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   						Symbol: jsii.String("symbol"),
//   					},
//   					NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   					PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   				},
//   			},
//   			HierarchyId: jsii.String("hierarchyId"),
//   		},
//   		NumericalDimensionField: &NumericalDimensionFieldProperty{
//   			Column: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			FieldId: jsii.String("fieldId"),
//
//   			// the properties below are optional
//   			FormatConfiguration: &NumberFormatConfigurationProperty{
//   				FormatConfiguration: &NumericFormatConfigurationProperty{
//   					CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   						Symbol: jsii.String("symbol"),
//   					},
//   					NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   					PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   				},
//   			},
//   			HierarchyId: jsii.String("hierarchyId"),
//   		},
//   	},
//   	ComputationId: jsii.String("computationId"),
//   	Time: &DimensionFieldProperty{
//   		CategoricalDimensionField: &CategoricalDimensionFieldProperty{
//   			Column: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			FieldId: jsii.String("fieldId"),
//
//   			// the properties below are optional
//   			FormatConfiguration: &StringFormatConfigurationProperty{
//   				NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   					NullString: jsii.String("nullString"),
//   				},
//   				NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   					CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   						Symbol: jsii.String("symbol"),
//   					},
//   					NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   					PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   				},
//   			},
//   			HierarchyId: jsii.String("hierarchyId"),
//   		},
//   		DateDimensionField: &DateDimensionFieldProperty{
//   			Column: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			FieldId: jsii.String("fieldId"),
//
//   			// the properties below are optional
//   			DateGranularity: jsii.String("dateGranularity"),
//   			FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   				DateTimeFormat: jsii.String("dateTimeFormat"),
//   				NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   					NullString: jsii.String("nullString"),
//   				},
//   				NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   					CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   						Symbol: jsii.String("symbol"),
//   					},
//   					NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   					PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   				},
//   			},
//   			HierarchyId: jsii.String("hierarchyId"),
//   		},
//   		NumericalDimensionField: &NumericalDimensionFieldProperty{
//   			Column: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			FieldId: jsii.String("fieldId"),
//
//   			// the properties below are optional
//   			FormatConfiguration: &NumberFormatConfigurationProperty{
//   				FormatConfiguration: &NumericFormatConfigurationProperty{
//   					CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   						Symbol: jsii.String("symbol"),
//   					},
//   					NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   					PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   				},
//   			},
//   			HierarchyId: jsii.String("hierarchyId"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	MoverSize: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	SortOrder: jsii.String("sortOrder"),
//   	Value: &MeasureFieldProperty{
//   		CalculatedMeasureField: &CalculatedMeasureFieldProperty{
//   			Expression: jsii.String("expression"),
//   			FieldId: jsii.String("fieldId"),
//   		},
//   		CategoricalMeasureField: &CategoricalMeasureFieldProperty{
//   			Column: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			FieldId: jsii.String("fieldId"),
//
//   			// the properties below are optional
//   			AggregationFunction: jsii.String("aggregationFunction"),
//   			FormatConfiguration: &StringFormatConfigurationProperty{
//   				NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   					NullString: jsii.String("nullString"),
//   				},
//   				NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   					CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   						Symbol: jsii.String("symbol"),
//   					},
//   					NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   					PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   				},
//   			},
//   		},
//   		DateMeasureField: &DateMeasureFieldProperty{
//   			Column: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			FieldId: jsii.String("fieldId"),
//
//   			// the properties below are optional
//   			AggregationFunction: jsii.String("aggregationFunction"),
//   			FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   				DateTimeFormat: jsii.String("dateTimeFormat"),
//   				NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   					NullString: jsii.String("nullString"),
//   				},
//   				NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   					CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   						Symbol: jsii.String("symbol"),
//   					},
//   					NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   					PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   				},
//   			},
//   		},
//   		NumericalMeasureField: &NumericalMeasureFieldProperty{
//   			Column: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			FieldId: jsii.String("fieldId"),
//
//   			// the properties below are optional
//   			AggregationFunction: &NumericalAggregationFunctionProperty{
//   				PercentileAggregation: &PercentileAggregationProperty{
//   					PercentileValue: jsii.Number(123),
//   				},
//   				SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   			},
//   			FormatConfiguration: &NumberFormatConfigurationProperty{
//   				FormatConfiguration: &NumericFormatConfigurationProperty{
//   					CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   						Symbol: jsii.String("symbol"),
//   					},
//   					NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						NumberScale: jsii.String("numberScale"),
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   					PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   						DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   							DecimalPlaces: jsii.Number(123),
//   						},
//   						NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   							DisplayMode: jsii.String("displayMode"),
//   						},
//   						NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   							NullString: jsii.String("nullString"),
//   						},
//   						Prefix: jsii.String("prefix"),
//   						SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   							DecimalSeparator: jsii.String("decimalSeparator"),
//   							ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   								Symbol: jsii.String("symbol"),
//   								Visibility: jsii.String("visibility"),
//   							},
//   						},
//   						Suffix: jsii.String("suffix"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnTemplate_TopBottomMoversComputationProperty struct {
	// The category field that is used in a computation.
	Category interface{} `field:"required" json:"category" yaml:"category"`
	// The ID for a computation.
	ComputationId *string `field:"required" json:"computationId" yaml:"computationId"`
	// The time field that is used in a computation.
	Time interface{} `field:"required" json:"time" yaml:"time"`
	// The computation type. Choose from the following options:.
	//
	// - TOP: Top movers computation.
	// - BOTTOM: Bottom movers computation.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The mover size setup of the top and bottom movers computation.
	MoverSize *float64 `field:"optional" json:"moverSize" yaml:"moverSize"`
	// The name of a computation.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The sort order setup of the top and bottom movers computation.
	SortOrder *string `field:"optional" json:"sortOrder" yaml:"sortOrder"`
	// The value field that is used in a computation.
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

