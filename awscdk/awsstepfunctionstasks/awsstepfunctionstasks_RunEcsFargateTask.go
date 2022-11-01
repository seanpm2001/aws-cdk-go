package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Start a service on an ECS cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var containerDefinition containerDefinition
//   var securityGroup securityGroup
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var taskDefinition taskDefinition
//
//   runEcsFargateTask := awscdk.Aws_stepfunctions_tasks.NewRunEcsFargateTask(&runEcsFargateTaskProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//
//   	// the properties below are optional
//   	assignPublicIp: jsii.Boolean(false),
//   	containerOverrides: []containerOverride{
//   		&containerOverride{
//   			containerDefinition: containerDefinition,
//
//   			// the properties below are optional
//   			command: []*string{
//   				jsii.String("command"),
//   			},
//   			cpu: jsii.Number(123),
//   			environment: []taskEnvironmentVariable{
//   				&taskEnvironmentVariable{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			memoryLimit: jsii.Number(123),
//   			memoryReservation: jsii.Number(123),
//   		},
//   	},
//   	integrationPattern: awscdk.Aws_stepfunctions.serviceIntegrationPattern_FIRE_AND_FORGET,
//   	platformVersion: awscdk.Aws_ecs.fargatePlatformVersion_LATEST,
//   	securityGroup: securityGroup,
//   	subnets: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnetName: jsii.String("subnetName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: awscdk.Aws_ec2.subnetType_ISOLATED,
//   	},
//   })
//
// Deprecated: replaced by `EcsRunTask`.
type RunEcsFargateTask interface {
	EcsRunTaskBase
	// Manage allowed network traffic for this service.
	// Deprecated: replaced by `EcsRunTask`.
	Connections() awsec2.Connections
	// Called when the task object is used in a workflow.
	// Deprecated: replaced by `EcsRunTask`.
	Bind(task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig
	// Deprecated: replaced by `EcsRunTask`.
	ConfigureAwsVpcNetworking(vpc awsec2.IVpc, assignPublicIp *bool, subnetSelection *awsec2.SubnetSelection, securityGroup awsec2.ISecurityGroup)
}

// The jsii proxy struct for RunEcsFargateTask
type jsiiProxy_RunEcsFargateTask struct {
	jsiiProxy_EcsRunTaskBase
}

func (j *jsiiProxy_RunEcsFargateTask) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}


// Deprecated: replaced by `EcsRunTask`.
func NewRunEcsFargateTask(props *RunEcsFargateTaskProps) RunEcsFargateTask {
	_init_.Initialize()

	if err := validateNewRunEcsFargateTaskParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RunEcsFargateTask{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.RunEcsFargateTask",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Deprecated: replaced by `EcsRunTask`.
func NewRunEcsFargateTask_Override(r RunEcsFargateTask, props *RunEcsFargateTaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.RunEcsFargateTask",
		[]interface{}{props},
		r,
	)
}

func (r *jsiiProxy_RunEcsFargateTask) Bind(task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig {
	if err := r.validateBindParameters(task); err != nil {
		panic(err)
	}
	var returns *awsstepfunctions.StepFunctionsTaskConfig

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{task},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RunEcsFargateTask) ConfigureAwsVpcNetworking(vpc awsec2.IVpc, assignPublicIp *bool, subnetSelection *awsec2.SubnetSelection, securityGroup awsec2.ISecurityGroup) {
	if err := r.validateConfigureAwsVpcNetworkingParameters(vpc, subnetSelection); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"configureAwsVpcNetworking",
		[]interface{}{vpc, assignPublicIp, subnetSelection, securityGroup},
	)
}

