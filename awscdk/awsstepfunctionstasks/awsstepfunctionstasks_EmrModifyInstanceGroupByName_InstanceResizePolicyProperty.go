package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Custom policy for requesting termination protection or termination of specific instances when shrinking an instance group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   instanceResizePolicyProperty := &instanceResizePolicyProperty{
//   	instancesToProtect: []*string{
//   		jsii.String("instancesToProtect"),
//   	},
//   	instancesToTerminate: []*string{
//   		jsii.String("instancesToTerminate"),
//   	},
//   	instanceTerminationTimeout: duration,
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceResizePolicy.html
//
// Experimental.
type EmrModifyInstanceGroupByName_InstanceResizePolicyProperty struct {
	// Specific list of instances to be protected when shrinking an instance group.
	// Experimental.
	InstancesToProtect *[]*string `field:"optional" json:"instancesToProtect" yaml:"instancesToProtect"`
	// Specific list of instances to be terminated when shrinking an instance group.
	// Experimental.
	InstancesToTerminate *[]*string `field:"optional" json:"instancesToTerminate" yaml:"instancesToTerminate"`
	// Decommissioning timeout override for the specific list of instances to be terminated.
	// Experimental.
	InstanceTerminationTimeout awscdk.Duration `field:"optional" json:"instanceTerminationTimeout" yaml:"instanceTerminationTimeout"`
}
