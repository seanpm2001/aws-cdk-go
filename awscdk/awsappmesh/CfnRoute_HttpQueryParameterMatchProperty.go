package awsappmesh


// An object representing the query parameter to match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpQueryParameterMatchProperty := &HttpQueryParameterMatchProperty{
//   	Exact: jsii.String("exact"),
//   }
//
type CfnRoute_HttpQueryParameterMatchProperty struct {
	// The exact query parameter to match on.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
}
