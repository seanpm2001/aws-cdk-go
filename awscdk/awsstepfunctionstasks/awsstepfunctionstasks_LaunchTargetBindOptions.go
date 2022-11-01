package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
)

// Options for binding a launch target to an ECS run job task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var taskDefinition taskDefinition
//
//   launchTargetBindOptions := &launchTargetBindOptions{
//   	taskDefinition: taskDefinition,
//
//   	// the properties below are optional
//   	cluster: cluster,
//   }
//
// Experimental.
type LaunchTargetBindOptions struct {
	// Task definition to run Docker containers in Amazon ECS.
	// Experimental.
	TaskDefinition awsecs.ITaskDefinition `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
	// A regional grouping of one or more container instances on which you can run tasks and services.
	// Experimental.
	Cluster awsecs.ICluster `field:"optional" json:"cluster" yaml:"cluster"`
}

