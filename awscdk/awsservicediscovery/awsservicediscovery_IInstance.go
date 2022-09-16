package awsservicediscovery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsservicediscovery/internal"
)

// Experimental.
type IInstance interface {
	awscdk.IResource
	// The id of the instance resource.
	// Experimental.
	InstanceId() *string
	// The Cloudmap service this resource is registered to.
	// Experimental.
	Service() IService
}

// The jsii proxy for IInstance
type jsiiProxy_IInstance struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IInstance) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstance) Service() IService {
	var returns IService
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}
