package awsappmesh


// Properties specific for a TCP Based Routes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var virtualNode virtualNode
//
//   tcpRouteSpecOptions := &tcpRouteSpecOptions{
//   	weightedTargets: []weightedTarget{
//   		&weightedTarget{
//   			virtualNode: virtualNode,
//
//   			// the properties below are optional
//   			weight: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	priority: jsii.Number(123),
//   	timeout: &tcpTimeout{
//   		idle: duration,
//   	},
//   }
//
// Experimental.
type TcpRouteSpecOptions struct {
	// The priority for the route.
	//
	// When a Virtual Router has multiple routes, route match is performed in the
	// order of specified value, where 0 is the highest priority, and first matched route is selected.
	// Experimental.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// List of targets that traffic is routed to when a request matches the route.
	// Experimental.
	WeightedTargets *[]*WeightedTarget `field:"required" json:"weightedTargets" yaml:"weightedTargets"`
	// An object that represents a tcp timeout.
	// Experimental.
	Timeout *TcpTimeout `field:"optional" json:"timeout" yaml:"timeout"`
}
