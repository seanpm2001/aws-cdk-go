package awsquicksight


// The source controls that are used in a `CascadingControlConfiguration` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cascadingControlSourceProperty := &CascadingControlSourceProperty{
//   	ColumnToMatch: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	SourceSheetControlId: jsii.String("sourceSheetControlId"),
//   }
//
type CfnDashboard_CascadingControlSourceProperty struct {
	// The column identifier that determines which column to look up for the source sheet control.
	ColumnToMatch interface{} `field:"optional" json:"columnToMatch" yaml:"columnToMatch"`
	// The source sheet control ID of a `CascadingControlSource` .
	SourceSheetControlId *string `field:"optional" json:"sourceSheetControlId" yaml:"sourceSheetControlId"`
}

