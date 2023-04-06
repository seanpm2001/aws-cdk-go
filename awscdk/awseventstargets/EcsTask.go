package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets/internal"
)

// Start a task on an ECS cluster.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var taskDefinition taskDefinition
//   var role role
//
//
//   ecsTaskTarget := awscdk.NewEcsTask(&EcsTaskProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	Role: Role,
//   })
//
//   awscdk.NewRule(this, jsii.String("ScheduleRule"), &RuleProps{
//   	Schedule: awscdk.Schedule_Cron(&CronOptions{
//   		Minute: jsii.String("0"),
//   		Hour: jsii.String("4"),
//   	}),
//   	Targets: []iRuleTarget{
//   		ecsTaskTarget,
//   	},
//   })
//
type EcsTask interface {
	awsevents.IRuleTarget
	// The security groups associated with the task.
	//
	// Only applicable with awsvpc network mode.
	SecurityGroups() *[]awsec2.ISecurityGroup
	// Allows using tasks as target of EventBridge events.
	Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for EcsTask
type jsiiProxy_EcsTask struct {
	internal.Type__awseventsIRuleTarget
}

func (j *jsiiProxy_EcsTask) SecurityGroups() *[]awsec2.ISecurityGroup {
	var returns *[]awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}


func NewEcsTask(props *EcsTaskProps) EcsTask {
	_init_.Initialize()

	if err := validateNewEcsTaskParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsTask{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.EcsTask",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewEcsTask_Override(e EcsTask, props *EcsTaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.EcsTask",
		[]interface{}{props},
		e,
	)
}

func (e *jsiiProxy_EcsTask) Bind(_rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	if err := e.validateBindParameters(_rule); err != nil {
		panic(err)
	}
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_rule, _id},
		&returns,
	)

	return returns
}
