package awscdkschedulertargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2"
)

// Experimental.
type ScheduleTargetBase interface {
	// Experimental.
	TargetArn() *string
	// Experimental.
	AddTargetActionToRole(schedule awscdkscheduleralpha.ISchedule, role awsiam.IRole)
	// Experimental.
	Bind(schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig
	// Experimental.
	BindBaseTargetConfig(_schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig
}

// The jsii proxy struct for ScheduleTargetBase
type jsiiProxy_ScheduleTargetBase struct {
	_ byte // padding
}

func (j *jsiiProxy_ScheduleTargetBase) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewScheduleTargetBase_Override(s ScheduleTargetBase, baseProps *ScheduleTargetBaseProps, targetArn *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.ScheduleTargetBase",
		[]interface{}{baseProps, targetArn},
		s,
	)
}

func (s *jsiiProxy_ScheduleTargetBase) AddTargetActionToRole(schedule awscdkscheduleralpha.ISchedule, role awsiam.IRole) {
	if err := s.validateAddTargetActionToRoleParameters(schedule, role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addTargetActionToRole",
		[]interface{}{schedule, role},
	)
}

func (s *jsiiProxy_ScheduleTargetBase) Bind(schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
	if err := s.validateBindParameters(schedule); err != nil {
		panic(err)
	}
	var returns *awscdkscheduleralpha.ScheduleTargetConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduleTargetBase) BindBaseTargetConfig(_schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
	if err := s.validateBindBaseTargetConfigParameters(_schedule); err != nil {
		panic(err)
	}
	var returns *awscdkscheduleralpha.ScheduleTargetConfig

	_jsii_.Invoke(
		s,
		"bindBaseTargetConfig",
		[]interface{}{_schedule},
		&returns,
	)

	return returns
}
