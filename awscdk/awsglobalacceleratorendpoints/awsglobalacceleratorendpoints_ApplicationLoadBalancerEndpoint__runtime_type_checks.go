//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsglobalacceleratorendpoints

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
)

func validateNewApplicationLoadBalancerEndpointParameters(loadBalancer awselasticloadbalancingv2.IApplicationLoadBalancer, options *ApplicationLoadBalancerEndpointOptions) error {
	if loadBalancer == nil {
		return fmt.Errorf("parameter loadBalancer is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}
