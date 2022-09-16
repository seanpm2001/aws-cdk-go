package awsappmesh


// Represents TLS properties for listener.
//
// Example:
//   // A Virtual Node with listener TLS from an ACM provided certificate
//   var cert certificate
//   var mesh mesh
//
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &virtualNodeProps{
//   	mesh: mesh,
//   	serviceDiscovery: appmesh.serviceDiscovery.dns(jsii.String("node")),
//   	listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener.grpc(&grpcVirtualNodeListenerOptions{
//   			port: jsii.Number(80),
//   			tls: &listenerTlsOptions{
//   				mode: appmesh.tlsMode_STRICT,
//   				certificate: appmesh.tlsCertificate.acm(cert),
//   			},
//   		}),
//   	},
//   })
//
//   // A Virtual Gateway with listener TLS from a customer provided file certificate
//   gateway := appmesh.NewVirtualGateway(this, jsii.String("gateway"), &virtualGatewayProps{
//   	mesh: mesh,
//   	listeners: []virtualGatewayListener{
//   		appmesh.*virtualGatewayListener.grpc(&grpcGatewayListenerOptions{
//   			port: jsii.Number(8080),
//   			tls: &listenerTlsOptions{
//   				mode: appmesh.*tlsMode_STRICT,
//   				certificate: appmesh.*tlsCertificate.file(jsii.String("path/to/certChain"), jsii.String("path/to/privateKey")),
//   			},
//   		}),
//   	},
//   	virtualGatewayName: jsii.String("gateway"),
//   })
//
//   // A Virtual Gateway with listener TLS from a SDS provided certificate
//   gateway2 := appmesh.NewVirtualGateway(this, jsii.String("gateway2"), &virtualGatewayProps{
//   	mesh: mesh,
//   	listeners: []*virtualGatewayListener{
//   		appmesh.*virtualGatewayListener.http2(&http2GatewayListenerOptions{
//   			port: jsii.Number(8080),
//   			tls: &listenerTlsOptions{
//   				mode: appmesh.*tlsMode_STRICT,
//   				certificate: appmesh.*tlsCertificate.sds(jsii.String("secrete_certificate")),
//   			},
//   		}),
//   	},
//   	virtualGatewayName: jsii.String("gateway2"),
//   })
//
// Experimental.
type ListenerTlsOptions struct {
	// Represents TLS certificate.
	// Experimental.
	Certificate TlsCertificate `field:"required" json:"certificate" yaml:"certificate"`
	// The TLS mode.
	// Experimental.
	Mode TlsMode `field:"required" json:"mode" yaml:"mode"`
	// Represents a listener's TLS validation context.
	//
	// The client certificate will only be validated if the client provides it, enabling mutual TLS.
	// Experimental.
	MutualTlsValidation *MutualTlsValidation `field:"optional" json:"mutualTlsValidation" yaml:"mutualTlsValidation"`
}
