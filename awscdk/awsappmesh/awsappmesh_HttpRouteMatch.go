package awsappmesh


// The criterion for determining a request match for this Route.
//
// Example:
//   var router virtualRouter
//   var node virtualNode
//
//
//   router.addRoute(jsii.String("route-http"), &routeBaseProps{
//   	routeSpec: appmesh.routeSpec.http(&httpRouteSpecOptions{
//   		weightedTargets: []weightedTarget{
//   			&weightedTarget{
//   				virtualNode: node,
//   				weight: jsii.Number(50),
//   			},
//   			&weightedTarget{
//   				virtualNode: node,
//   				weight: jsii.Number(50),
//   			},
//   		},
//   		match: &httpRouteMatch{
//   			path: appmesh.httpRoutePathMatch.startsWith(jsii.String("/path-to-app")),
//   		},
//   	}),
//   })
//
// Experimental.
type HttpRouteMatch struct {
	// Specifies the client request headers to match on.
	//
	// All specified headers
	// must match for the route to match.
	// Experimental.
	Headers *[]HeaderMatch `field:"optional" json:"headers" yaml:"headers"`
	// The HTTP client request method to match on.
	// Experimental.
	Method HttpRouteMethod `field:"optional" json:"method" yaml:"method"`
	// Specifies how is the request matched based on the path part of its URL.
	// Experimental.
	Path HttpRoutePathMatch `field:"optional" json:"path" yaml:"path"`
	// The client request protocol to match on.
	//
	// Applicable only for HTTP2 routes.
	// Experimental.
	Protocol HttpRouteProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The query parameters to match on.
	//
	// All specified query parameters must match for the route to match.
	// Experimental.
	QueryParameters *[]QueryParameterMatch `field:"optional" json:"queryParameters" yaml:"queryParameters"`
}
