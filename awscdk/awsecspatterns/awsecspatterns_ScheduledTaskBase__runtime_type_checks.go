//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsecspatterns

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/awseventstargets"
	"github.com/aws/constructs-go/constructs/v3"
)

func (s *jsiiProxy_ScheduledTaskBase) validateAddTaskAsTargetParameters(ecsTaskTarget awseventstargets.EcsTask) error {
	if ecsTaskTarget == nil {
		return fmt.Errorf("parameter ecsTaskTarget is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ScheduledTaskBase) validateAddTaskDefinitionToEventTargetParameters(taskDefinition awsecs.TaskDefinition) error {
	if taskDefinition == nil {
		return fmt.Errorf("parameter taskDefinition is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ScheduledTaskBase) validateCreateAWSLogDriverParameters(prefix *string) error {
	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ScheduledTaskBase) validateGetDefaultClusterParameters(scope awscdk.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ScheduledTaskBase) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_ScheduledTaskBase) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func validateScheduledTaskBase_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewScheduledTaskBaseParameters(scope constructs.Construct, id *string, props *ScheduledTaskBaseProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}
