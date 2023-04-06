package awsquicksight


// A *logical table* is a unit that joins and that data transformations operate on.
//
// A logical table has a source, which can be either a physical table or result of a join. When a logical table points to a physical table, the logical table acts as a mutable copy of that physical table through transform operations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logicalTableProperty := &LogicalTableProperty{
//   	Alias: jsii.String("alias"),
//   	Source: &LogicalTableSourceProperty{
//   		DataSetArn: jsii.String("dataSetArn"),
//   		JoinInstruction: &JoinInstructionProperty{
//   			LeftOperand: jsii.String("leftOperand"),
//   			OnClause: jsii.String("onClause"),
//   			RightOperand: jsii.String("rightOperand"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			LeftJoinKeyProperties: &JoinKeyPropertiesProperty{
//   				UniqueKey: jsii.Boolean(false),
//   			},
//   			RightJoinKeyProperties: &JoinKeyPropertiesProperty{
//   				UniqueKey: jsii.Boolean(false),
//   			},
//   		},
//   		PhysicalTableId: jsii.String("physicalTableId"),
//   	},
//
//   	// the properties below are optional
//   	DataTransforms: []interface{}{
//   		&TransformOperationProperty{
//   			CastColumnTypeOperation: &CastColumnTypeOperationProperty{
//   				ColumnName: jsii.String("columnName"),
//   				NewColumnType: jsii.String("newColumnType"),
//
//   				// the properties below are optional
//   				Format: jsii.String("format"),
//   			},
//   			CreateColumnsOperation: &CreateColumnsOperationProperty{
//   				Columns: []interface{}{
//   					&CalculatedColumnProperty{
//   						ColumnId: jsii.String("columnId"),
//   						ColumnName: jsii.String("columnName"),
//   						Expression: jsii.String("expression"),
//   					},
//   				},
//   			},
//   			FilterOperation: &FilterOperationProperty{
//   				ConditionExpression: jsii.String("conditionExpression"),
//   			},
//   			ProjectOperation: &ProjectOperationProperty{
//   				ProjectedColumns: []*string{
//   					jsii.String("projectedColumns"),
//   				},
//   			},
//   			RenameColumnOperation: &RenameColumnOperationProperty{
//   				ColumnName: jsii.String("columnName"),
//   				NewColumnName: jsii.String("newColumnName"),
//   			},
//   			TagColumnOperation: &TagColumnOperationProperty{
//   				ColumnName: jsii.String("columnName"),
//   				Tags: []columnTagProperty{
//   					&columnTagProperty{
//   						ColumnDescription: &ColumnDescriptionProperty{
//   							Text: jsii.String("text"),
//   						},
//   						ColumnGeographicRole: jsii.String("columnGeographicRole"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnDataSet_LogicalTableProperty struct {
	// A display name for the logical table.
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// Source of this logical table.
	Source interface{} `field:"required" json:"source" yaml:"source"`
	// Transform operations that act on this logical table.
	//
	// For this structure to be valid, only one of the attributes can be non-null.
	DataTransforms interface{} `field:"optional" json:"dataTransforms" yaml:"dataTransforms"`
}
