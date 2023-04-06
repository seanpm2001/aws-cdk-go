package awselasticache


// A member of a Global datastore.
//
// It contains the Replication Group Id, the Amazon region and the role of the replication group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   globalReplicationGroupMemberProperty := &GlobalReplicationGroupMemberProperty{
//   	ReplicationGroupId: jsii.String("replicationGroupId"),
//   	ReplicationGroupRegion: jsii.String("replicationGroupRegion"),
//   	Role: jsii.String("role"),
//   }
//
type CfnGlobalReplicationGroup_GlobalReplicationGroupMemberProperty struct {
	// The replication group id of the Global datastore member.
	ReplicationGroupId *string `field:"optional" json:"replicationGroupId" yaml:"replicationGroupId"`
	// The Amazon region of the Global datastore member.
	ReplicationGroupRegion *string `field:"optional" json:"replicationGroupRegion" yaml:"replicationGroupRegion"`
	// Indicates the role of the replication group, `PRIMARY` or `SECONDARY` .
	Role *string `field:"optional" json:"role" yaml:"role"`
}

