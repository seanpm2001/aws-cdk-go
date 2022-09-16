package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Properties for EmrModifyInstanceGroupByName.
//
// Example:
//   tasks.NewEmrModifyInstanceGroupByName(this, jsii.String("Task"), &emrModifyInstanceGroupByNameProps{
//   	clusterId: jsii.String("ClusterId"),
//   	instanceGroupName: sfn.jsonPath.stringAt(jsii.String("$.InstanceGroupName")),
//   	instanceGroup: &instanceGroupModifyConfigProperty{
//   		instanceCount: jsii.Number(1),
//   	},
//   })
//
// Experimental.
type EmrModifyInstanceGroupByNameProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The ClusterId to update.
	// Experimental.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The JSON that you want to provide to your ModifyInstanceGroup call as input.
	//
	// This uses the same syntax as the ModifyInstanceGroups API.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_ModifyInstanceGroups.html
	//
	// Experimental.
	InstanceGroup *EmrModifyInstanceGroupByName_InstanceGroupModifyConfigProperty `field:"required" json:"instanceGroup" yaml:"instanceGroup"`
	// The InstanceGroupName to update.
	// Experimental.
	InstanceGroupName *string `field:"required" json:"instanceGroupName" yaml:"instanceGroupName"`
}
