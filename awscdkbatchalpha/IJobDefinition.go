// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkbatchalpha/v2/internal"
)

// An interface representing a job definition - either a new one, created with the CDK, *using the `JobDefinition` class, or existing ones, referenced using the `JobDefinition.fromJobDefinitionArn` method.
// Experimental.
type IJobDefinition interface {
	awscdk.IResource
	// The ARN of this batch job definition.
	// Experimental.
	JobDefinitionArn() *string
	// The name of the batch job definition.
	// Experimental.
	JobDefinitionName() *string
}

// The jsii proxy for IJobDefinition
type jsiiProxy_IJobDefinition struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IJobDefinition) JobDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobDefinition) JobDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionName",
		&returns,
	)
	return returns
}

