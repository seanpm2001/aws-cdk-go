package awscdkapigatewayv2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2/internal"
)

// Represents an integration to an API Route.
// Experimental.
type IIntegration interface {
	awscdk.IResource
	// Id of the integration.
	// Experimental.
	IntegrationId() *string
}

// The jsii proxy for IIntegration
type jsiiProxy_IIntegration struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IIntegration) IntegrationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationId",
		&returns,
	)
	return returns
}

