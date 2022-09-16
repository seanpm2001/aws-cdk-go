package cloudassemblyschema


// Query input for looking up a load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerContextQuery := &loadBalancerContextQuery{
//   	account: jsii.String("account"),
//   	loadBalancerType: awscdk.Cloud_assembly_schema.loadBalancerType_NETWORK,
//   	region: jsii.String("region"),
//
//   	// the properties below are optional
//   	loadBalancerArn: jsii.String("loadBalancerArn"),
//   	loadBalancerTags: []tag{
//   		&tag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	lookupRoleArn: jsii.String("lookupRoleArn"),
//   }
//
// Experimental.
type LoadBalancerContextQuery struct {
	// Filter load balancers by their type.
	// Experimental.
	LoadBalancerType LoadBalancerType `field:"required" json:"loadBalancerType" yaml:"loadBalancerType"`
	// Find by load balancer's ARN.
	// Experimental.
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Match load balancer tags.
	// Experimental.
	LoadBalancerTags *[]*Tag `field:"optional" json:"loadBalancerTags" yaml:"loadBalancerTags"`
	// Query account.
	// Experimental.
	Account *string `field:"required" json:"account" yaml:"account"`
	// Query region.
	// Experimental.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	// Experimental.
	LookupRoleArn *string `field:"optional" json:"lookupRoleArn" yaml:"lookupRoleArn"`
}
