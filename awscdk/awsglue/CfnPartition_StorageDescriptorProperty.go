package awsglue


// Describes the physical storage of table data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//   var skewedColumnValueLocationMaps interface{}
//
//   storageDescriptorProperty := &StorageDescriptorProperty{
//   	BucketColumns: []*string{
//   		jsii.String("bucketColumns"),
//   	},
//   	Columns: []interface{}{
//   		&ColumnProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Comment: jsii.String("comment"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Compressed: jsii.Boolean(false),
//   	InputFormat: jsii.String("inputFormat"),
//   	Location: jsii.String("location"),
//   	NumberOfBuckets: jsii.Number(123),
//   	OutputFormat: jsii.String("outputFormat"),
//   	Parameters: parameters,
//   	SchemaReference: &SchemaReferenceProperty{
//   		SchemaId: &SchemaIdProperty{
//   			RegistryName: jsii.String("registryName"),
//   			SchemaArn: jsii.String("schemaArn"),
//   			SchemaName: jsii.String("schemaName"),
//   		},
//   		SchemaVersionId: jsii.String("schemaVersionId"),
//   		SchemaVersionNumber: jsii.Number(123),
//   	},
//   	SerdeInfo: &SerdeInfoProperty{
//   		Name: jsii.String("name"),
//   		Parameters: parameters,
//   		SerializationLibrary: jsii.String("serializationLibrary"),
//   	},
//   	SkewedInfo: &SkewedInfoProperty{
//   		SkewedColumnNames: []*string{
//   			jsii.String("skewedColumnNames"),
//   		},
//   		SkewedColumnValueLocationMaps: skewedColumnValueLocationMaps,
//   		SkewedColumnValues: []*string{
//   			jsii.String("skewedColumnValues"),
//   		},
//   	},
//   	SortColumns: []interface{}{
//   		&OrderProperty{
//   			Column: jsii.String("column"),
//
//   			// the properties below are optional
//   			SortOrder: jsii.Number(123),
//   		},
//   	},
//   	StoredAsSubDirectories: jsii.Boolean(false),
//   }
//
type CfnPartition_StorageDescriptorProperty struct {
	// A list of reducer grouping columns, clustering columns, and bucketing columns in the table.
	BucketColumns *[]*string `field:"optional" json:"bucketColumns" yaml:"bucketColumns"`
	// A list of the `Columns` in the table.
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// `True` if the data in the table is compressed, or `False` if not.
	Compressed interface{} `field:"optional" json:"compressed" yaml:"compressed"`
	// The input format: `SequenceFileInputFormat` (binary), or `TextInputFormat` , or a custom format.
	InputFormat *string `field:"optional" json:"inputFormat" yaml:"inputFormat"`
	// The physical location of the table.
	//
	// By default, this takes the form of the warehouse location, followed by the database location in the warehouse, followed by the table name.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The number of buckets.
	//
	// You must specify this property if the partition contains any dimension columns.
	NumberOfBuckets *float64 `field:"optional" json:"numberOfBuckets" yaml:"numberOfBuckets"`
	// The output format: `SequenceFileOutputFormat` (binary), or `IgnoreKeyTextOutputFormat` , or a custom format.
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// The user-supplied properties in key-value form.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// An object that references a schema stored in the AWS Glue Schema Registry.
	SchemaReference interface{} `field:"optional" json:"schemaReference" yaml:"schemaReference"`
	// The serialization/deserialization (SerDe) information.
	SerdeInfo interface{} `field:"optional" json:"serdeInfo" yaml:"serdeInfo"`
	// The information about values that appear frequently in a column (skewed values).
	SkewedInfo interface{} `field:"optional" json:"skewedInfo" yaml:"skewedInfo"`
	// A list specifying the sort order of each bucket in the table.
	SortColumns interface{} `field:"optional" json:"sortColumns" yaml:"sortColumns"`
	// `True` if the table data is stored in subdirectories, or `False` if not.
	StoredAsSubDirectories interface{} `field:"optional" json:"storedAsSubDirectories" yaml:"storedAsSubDirectories"`
}
