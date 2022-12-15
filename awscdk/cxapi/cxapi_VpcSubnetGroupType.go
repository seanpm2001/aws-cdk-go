package cxapi


// The type of subnet group.
//
// Same as SubnetType in the @aws-cdk/aws-ec2 package,
// but we can't use that because of cyclical dependencies.
// Experimental.
type VpcSubnetGroupType string

const (
	// Public subnet group type.
	// Experimental.
	VpcSubnetGroupType_PUBLIC VpcSubnetGroupType = "PUBLIC"
	// Private subnet group type.
	// Experimental.
	VpcSubnetGroupType_PRIVATE VpcSubnetGroupType = "PRIVATE"
	// Isolated subnet group type.
	// Experimental.
	VpcSubnetGroupType_ISOLATED VpcSubnetGroupType = "ISOLATED"
)

