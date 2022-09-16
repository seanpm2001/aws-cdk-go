package awsappmesh


// Represent the GRPC Node Listener prorperty.
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
type GrpcVirtualNodeListenerOptions struct {
	// Connection pool for http listeners.
	// Experimental.
	ConnectionPool *GrpcConnectionPool `field:"optional" json:"connectionPool" yaml:"connectionPool"`
	// The health check information for the listener.
	// Experimental.
	HealthCheck HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Represents the configuration for enabling outlier detection.
	// Experimental.
	OutlierDetection *OutlierDetection `field:"optional" json:"outlierDetection" yaml:"outlierDetection"`
	// Port to listen for connections on.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Timeout for GRPC protocol.
	// Experimental.
	Timeout *GrpcTimeout `field:"optional" json:"timeout" yaml:"timeout"`
	// Represents the configuration for enabling TLS on a listener.
	// Experimental.
	Tls *ListenerTlsOptions `field:"optional" json:"tls" yaml:"tls"`
}
