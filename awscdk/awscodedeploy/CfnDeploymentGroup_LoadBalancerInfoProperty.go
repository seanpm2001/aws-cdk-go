package awscodedeploy


// The `LoadBalancerInfo` property type specifies information about the load balancer or target group used for an AWS CodeDeploy deployment group.
//
// For more information, see [Integrating CodeDeploy with Elastic Load Balancing](https://docs.aws.amazon.com/codedeploy/latest/userguide/integrations-aws-elastic-load-balancing.html) in the *AWS CodeDeploy User Guide* .
//
// For AWS CloudFormation to use the properties specified in `LoadBalancerInfo` , the `DeploymentStyle.DeploymentOption` property must be set to `WITH_TRAFFIC_CONTROL` . If `DeploymentStyle.DeploymentOption` is not set to `WITH_TRAFFIC_CONTROL` , AWS CloudFormation ignores any settings specified in `LoadBalancerInfo` .
//
// > AWS CloudFormation supports blue/green deployments on the AWS Lambda compute platform only.
//
// `LoadBalancerInfo` is a property of the [DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerInfoProperty := &LoadBalancerInfoProperty{
//   	ElbInfoList: []interface{}{
//   		&ELBInfoProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	TargetGroupInfoList: []interface{}{
//   		&TargetGroupInfoProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	TargetGroupPairInfoList: []interface{}{
//   		&TargetGroupPairInfoProperty{
//   			ProdTrafficRoute: &TrafficRouteProperty{
//   				ListenerArns: []*string{
//   					jsii.String("listenerArns"),
//   				},
//   			},
//   			TargetGroups: []interface{}{
//   				&TargetGroupInfoProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			TestTrafficRoute: &TrafficRouteProperty{
//   				ListenerArns: []*string{
//   					jsii.String("listenerArns"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-loadbalancerinfo.html
//
type CfnDeploymentGroup_LoadBalancerInfoProperty struct {
	// An array that contains information about the load balancer to use for load balancing in a deployment.
	//
	// In Elastic Load Balancing, load balancers are used with Classic Load Balancers.
	//
	// > Adding more than one load balancer to the array is not supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-loadbalancerinfo.html#cfn-codedeploy-deploymentgroup-loadbalancerinfo-elbinfolist
	//
	ElbInfoList interface{} `field:"optional" json:"elbInfoList" yaml:"elbInfoList"`
	// An array that contains information about the target group to use for load balancing in a deployment.
	//
	// In Elastic Load Balancing , target groups are used with Application Load Balancers .
	//
	// > Adding more than one target group to the array is not supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-loadbalancerinfo.html#cfn-codedeploy-deploymentgroup-loadbalancerinfo-targetgroupinfolist
	//
	TargetGroupInfoList interface{} `field:"optional" json:"targetGroupInfoList" yaml:"targetGroupInfoList"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-loadbalancerinfo.html#cfn-codedeploy-deploymentgroup-loadbalancerinfo-targetgrouppairinfolist
	//
	TargetGroupPairInfoList interface{} `field:"optional" json:"targetGroupPairInfoList" yaml:"targetGroupPairInfoList"`
}

