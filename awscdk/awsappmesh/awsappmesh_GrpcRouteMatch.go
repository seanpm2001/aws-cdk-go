package awsappmesh


// The criterion for determining a request match for this Route.
//
// At least one match type must be selected.
//
// Example:
//   var router virtualRouter
//   var node virtualNode
//
//
//   router.addRoute(jsii.String("route-http"), &routeBaseProps{
//   	routeSpec: appmesh.routeSpec.grpc(&grpcRouteSpecOptions{
//   		weightedTargets: []weightedTarget{
//   			&weightedTarget{
//   				virtualNode: node,
//   			},
//   		},
//   		match: &grpcRouteMatch{
//   			serviceName: jsii.String("my-service.default.svc.cluster.local"),
//   		},
//   		timeout: &grpcTimeout{
//   			idle: cdk.duration.seconds(jsii.Number(2)),
//   			perRequest: cdk.*duration.seconds(jsii.Number(1)),
//   		},
//   	}),
//   })
//
// Experimental.
type GrpcRouteMatch struct {
	// Create metadata based gRPC route match.
	//
	// All specified metadata must match for the route to match.
	// Experimental.
	Metadata *[]HeaderMatch `field:"optional" json:"metadata" yaml:"metadata"`
	// The method name to match from the request.
	//
	// If the method name is specified, service name must be also provided.
	// Experimental.
	MethodName *string `field:"optional" json:"methodName" yaml:"methodName"`
	// Create service name based gRPC route match.
	// Experimental.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}
