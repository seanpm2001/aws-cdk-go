//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsiotactions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiot"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
)

func (s *jsiiProxy_SqsQueueAction) validateBindParameters(rule awsiot.ITopicRule) error {
	if rule == nil {
		return fmt.Errorf("parameter rule is required, but nil was provided")
	}

	return nil
}

func validateNewSqsQueueActionParameters(queue awssqs.IQueue, props *SqsQueueActionProps) error {
	if queue == nil {
		return fmt.Errorf("parameter queue is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

