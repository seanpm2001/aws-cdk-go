//go:build !no_runtime_type_checking

package awselasticloadbalancingv2

import (
	"fmt"
)

func (i *jsiiProxy_IpTarget) validateAttachToApplicationTargetGroupParameters(targetGroup IApplicationTargetGroup) error {
	if targetGroup == nil {
		return fmt.Errorf("parameter targetGroup is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IpTarget) validateAttachToNetworkTargetGroupParameters(targetGroup INetworkTargetGroup) error {
	if targetGroup == nil {
		return fmt.Errorf("parameter targetGroup is required, but nil was provided")
	}

	return nil
}

func validateNewIpTargetParameters(ipAddress *string) error {
	if ipAddress == nil {
		return fmt.Errorf("parameter ipAddress is required, but nil was provided")
	}

	return nil
}
