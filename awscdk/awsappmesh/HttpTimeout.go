package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Represents timeouts for HTTP protocols.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var mesh mesh
//   var service service
//
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &VirtualNodeProps{
//   	Mesh: Mesh,
//   	ServiceDiscovery: appmesh.ServiceDiscovery_CloudMap(service),
//   	Listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener_Http(&HttpVirtualNodeListenerOptions{
//   			Port: jsii.Number(8080),
//   			HealthCheck: appmesh.HealthCheck_Http(&HttpHealthCheckOptions{
//   				HealthyThreshold: jsii.Number(3),
//   				Interval: cdk.Duration_Seconds(jsii.Number(5)),
//   				Path: jsii.String("/ping"),
//   				Timeout: cdk.Duration_*Seconds(jsii.Number(2)),
//   				UnhealthyThreshold: jsii.Number(2),
//   			}),
//   			Timeout: &HttpTimeout{
//   				Idle: cdk.Duration_*Seconds(jsii.Number(5)),
//   			},
//   		}),
//   	},
//   	BackendDefaults: &BackendDefaults{
//   		TlsClientPolicy: &TlsClientPolicy{
//   			Validation: &TlsValidation{
//   				Trust: appmesh.TlsValidationTrust_File(jsii.String("/keys/local_cert_chain.pem")),
//   			},
//   		},
//   	},
//   	AccessLog: appmesh.AccessLog_FromFilePath(jsii.String("/dev/stdout")),
//   })
//
//   cdk.Tags_Of(node).Add(jsii.String("Environment"), jsii.String("Dev"))
//
type HttpTimeout struct {
	// Represents an idle timeout.
	//
	// The amount of time that a connection may be idle.
	Idle awscdk.Duration `field:"optional" json:"idle" yaml:"idle"`
	// Represents per request timeout.
	PerRequest awscdk.Duration `field:"optional" json:"perRequest" yaml:"perRequest"`
}

