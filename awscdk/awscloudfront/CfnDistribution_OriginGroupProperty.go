package awscloudfront


// An origin group includes two origins (a primary origin and a second origin to failover to) and a failover criteria that you specify.
//
// You create an origin group to support origin failover in CloudFront. When you create or update a distribution, you can specifiy the origin group instead of a single origin, and CloudFront will failover from the primary origin to the second origin under the failover conditions that you've chosen.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originGroupProperty := &OriginGroupProperty{
//   	FailoverCriteria: &OriginGroupFailoverCriteriaProperty{
//   		StatusCodes: &StatusCodesProperty{
//   			Items: []interface{}{
//   				jsii.Number(123),
//   			},
//   			Quantity: jsii.Number(123),
//   		},
//   	},
//   	Id: jsii.String("id"),
//   	Members: &OriginGroupMembersProperty{
//   		Items: []interface{}{
//   			&OriginGroupMemberProperty{
//   				OriginId: jsii.String("originId"),
//   			},
//   		},
//   		Quantity: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroup.html
//
type CfnDistribution_OriginGroupProperty struct {
	// A complex type that contains information about the failover criteria for an origin group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroup.html#cfn-cloudfront-distribution-origingroup-failovercriteria
	//
	FailoverCriteria interface{} `field:"required" json:"failoverCriteria" yaml:"failoverCriteria"`
	// The origin group's ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroup.html#cfn-cloudfront-distribution-origingroup-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// A complex type that contains information about the origins in an origin group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroup.html#cfn-cloudfront-distribution-origingroup-members
	//
	Members interface{} `field:"required" json:"members" yaml:"members"`
}

