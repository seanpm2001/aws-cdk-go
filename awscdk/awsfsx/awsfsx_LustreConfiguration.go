package awsfsx


// The configuration for the Amazon FSx for Lustre file system.
//
// Example:
//   var vpc vpc
//
//
//   fileSystem := fsx.NewLustreFileSystem(this, jsii.String("FsxLustreFileSystem"), &lustreFileSystemProps{
//   	lustreConfiguration: &lustreConfiguration{
//   		deploymentType: fsx.lustreDeploymentType_SCRATCH_2,
//   	},
//   	storageCapacityGiB: jsii.Number(1200),
//   	vpc: vpc,
//   	vpcSubnet: vpc.privateSubnets[jsii.Number(0)],
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-lustreconfiguration.html
//
// Experimental.
type LustreConfiguration struct {
	// The type of backing file system deployment used by FSx.
	// Experimental.
	DeploymentType LustreDeploymentType `field:"required" json:"deploymentType" yaml:"deploymentType"`
	// The path in Amazon S3 where the root of your Amazon FSx file system is exported.
	//
	// The path must use the same
	// Amazon S3 bucket as specified in ImportPath. If you only specify a bucket name, such as s3://import-bucket, you
	// get a 1:1 mapping of file system objects to S3 bucket objects. This mapping means that the input data in S3 is
	// overwritten on export. If you provide a custom prefix in the export path, such as
	// s3://import-bucket/[custom-optional-prefix], Amazon FSx exports the contents of your file system to that export
	// prefix in the Amazon S3 bucket.
	// Experimental.
	ExportPath *string `field:"optional" json:"exportPath" yaml:"exportPath"`
	// For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk.
	//
	// Allowed values are between 1 and 512,000.
	// Experimental.
	ImportedFileChunkSizeMiB *float64 `field:"optional" json:"importedFileChunkSizeMiB" yaml:"importedFileChunkSizeMiB"`
	// The path to the Amazon S3 bucket (including the optional prefix) that you're using as the data repository for your Amazon FSx for Lustre file system.
	//
	// Must be of the format "s3://{bucketName}/optional-prefix" and cannot
	// exceed 900 characters.
	// Experimental.
	ImportPath *string `field:"optional" json:"importPath" yaml:"importPath"`
	// Required for the PERSISTENT_1 deployment type, describes the amount of read and write throughput for each 1 tebibyte of storage, in MB/s/TiB.
	//
	// Valid values are 50, 100, 200.
	// Experimental.
	PerUnitStorageThroughput *float64 `field:"optional" json:"perUnitStorageThroughput" yaml:"perUnitStorageThroughput"`
	// The preferred day and time to perform weekly maintenance.
	//
	// The first digit is the day of the week, starting at 1
	// for Monday, then the following are hours and minutes in the UTC time zone, 24 hour clock. For example: '2:20:30'
	// is Tuesdays at 20:30.
	// Experimental.
	WeeklyMaintenanceStartTime LustreMaintenanceTime `field:"optional" json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

