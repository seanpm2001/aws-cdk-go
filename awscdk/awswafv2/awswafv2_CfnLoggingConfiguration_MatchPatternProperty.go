package awswafv2


// The patterns to look for in the JSON body.
//
// AWS WAF inspects the results of these pattern matches against the rule inspection criteria.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   matchPatternProperty := &matchPatternProperty{
//   	all: all,
//   	includedPaths: []*string{
//   		jsii.String("includedPaths"),
//   	},
//   }
//
type CfnLoggingConfiguration_MatchPatternProperty struct {
	// Match all of the elements.
	//
	// You must specify either this setting or the `IncludedPaths` setting, but not both.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Match only the specified include paths.
	//
	// Provide the include paths using JSON Pointer syntax. For example, `"IncludedPaths": ["/dogs/0/name", "/dogs/1/name"]` . For information about this syntax, see the Internet Engineering Task Force (IETF) documentation [JavaScript Object Notation (JSON) Pointer](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc6901) .
	//
	// You must specify either this setting or the `All` setting, but not both.
	//
	// > Don't use this option to include all paths. Instead, use the `All` setting.
	IncludedPaths *[]*string `field:"optional" json:"includedPaths" yaml:"includedPaths"`
}

