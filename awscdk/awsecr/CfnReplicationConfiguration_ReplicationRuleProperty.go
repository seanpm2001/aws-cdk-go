package awsecr


// An array of objects representing the replication destinations and repository filters for a replication configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationRuleProperty := &ReplicationRuleProperty{
//   	Destinations: []interface{}{
//   		&ReplicationDestinationProperty{
//   			Region: jsii.String("region"),
//   			RegistryId: jsii.String("registryId"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	RepositoryFilters: []interface{}{
//   		&RepositoryFilterProperty{
//   			Filter: jsii.String("filter"),
//   			FilterType: jsii.String("filterType"),
//   		},
//   	},
//   }
//
type CfnReplicationConfiguration_ReplicationRuleProperty struct {
	// An array of objects representing the destination for a replication rule.
	Destinations interface{} `field:"required" json:"destinations" yaml:"destinations"`
	// An array of objects representing the filters for a replication rule.
	//
	// Specifying a repository filter for a replication rule provides a method for controlling which repositories in a private registry are replicated.
	RepositoryFilters interface{} `field:"optional" json:"repositoryFilters" yaml:"repositoryFilters"`
}

