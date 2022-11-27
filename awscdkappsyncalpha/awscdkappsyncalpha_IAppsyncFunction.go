// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkappsyncalpha/v2/internal"
)

// Interface for AppSync Functions.
// Experimental.
type IAppsyncFunction interface {
	awscdk.IResource
	// the ARN of the AppSync function.
	// Experimental.
	FunctionArn() *string
	// the name of this AppSync Function.
	// Experimental.
	FunctionId() *string
}

// The jsii proxy for IAppsyncFunction
type jsiiProxy_IAppsyncFunction struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IAppsyncFunction) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppsyncFunction) FunctionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionId",
		&returns,
	)
	return returns
}
