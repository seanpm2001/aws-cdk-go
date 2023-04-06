package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The properties for enabling scaling based on CPU utilization.
//
// Example:
//   var target applicationTargetGroup
//   var service baseService
//
//   scaling := service.AutoScaleTaskCount(&EnableScalingProps{
//   	MaxCapacity: jsii.Number(10),
//   })
//   scaling.ScaleOnCpuUtilization(jsii.String("CpuScaling"), &CpuUtilizationScalingProps{
//   	TargetUtilizationPercent: jsii.Number(50),
//   })
//
//   scaling.ScaleOnRequestCount(jsii.String("RequestScaling"), &RequestCountScalingProps{
//   	RequestsPerTarget: jsii.Number(10000),
//   	TargetGroup: target,
//   })
//
type CpuUtilizationScalingProps struct {
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the scalable resource. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// scalable resource.
	DisableScaleIn *bool `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// A name for the scaling policy.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// Period after a scale in activity completes before another scale in activity can start.
	ScaleInCooldown awscdk.Duration `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Period after a scale out activity completes before another scale out activity can start.
	ScaleOutCooldown awscdk.Duration `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
	// The target value for CPU utilization across all tasks in the service.
	TargetUtilizationPercent *float64 `field:"required" json:"targetUtilizationPercent" yaml:"targetUtilizationPercent"`
}
