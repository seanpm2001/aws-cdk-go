package awswafv2


// A rule statement that inspects for malicious SQL code.
//
// Attackers insert malicious SQL code into web requests to do things like modify your database or extract data from it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var uriPath interface{}
//
//   sqliMatchStatementProperty := &sqliMatchStatementProperty{
//   	fieldToMatch: &fieldToMatchProperty{
//   		allQueryArguments: allQueryArguments,
//   		body: &bodyProperty{
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		cookies: &cookiesProperty{
//   			matchPattern: &cookieMatchPatternProperty{
//   				all: all,
//   				excludedCookies: []*string{
//   					jsii.String("excludedCookies"),
//   				},
//   				includedCookies: []*string{
//   					jsii.String("includedCookies"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		headers: &headersProperty{
//   			matchPattern: &headerMatchPatternProperty{
//   				all: all,
//   				excludedHeaders: []*string{
//   					jsii.String("excludedHeaders"),
//   				},
//   				includedHeaders: []*string{
//   					jsii.String("includedHeaders"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		jsonBody: &jsonBodyProperty{
//   			matchPattern: &jsonMatchPatternProperty{
//   				all: all,
//   				includedPaths: []*string{
//   					jsii.String("includedPaths"),
//   				},
//   			},
//   			matchScope: jsii.String("matchScope"),
//
//   			// the properties below are optional
//   			invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   			oversizeHandling: jsii.String("oversizeHandling"),
//   		},
//   		method: method,
//   		queryString: queryString,
//   		singleHeader: singleHeader,
//   		singleQueryArgument: singleQueryArgument,
//   		uriPath: uriPath,
//   	},
//   	textTransformations: []interface{}{
//   		&textTransformationProperty{
//   			priority: jsii.Number(123),
//   			type: jsii.String("type"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	sensitivityLevel: jsii.String("sensitivityLevel"),
//   }
//
type CfnRuleGroup_SqliMatchStatementProperty struct {
	// The part of the web request that you want AWS WAF to inspect.
	FieldToMatch interface{} `field:"required" json:"fieldToMatch" yaml:"fieldToMatch"`
	// Text transformations eliminate some of the unusual formatting that attackers use in web requests in an effort to bypass detection.
	//
	// If you specify one or more transformations in a rule statement, AWS WAF performs all transformations on the content of the request component identified by `FieldToMatch` , starting from the lowest priority setting, before inspecting the content for a match.
	TextTransformations interface{} `field:"required" json:"textTransformations" yaml:"textTransformations"`
	// `CfnRuleGroup.SqliMatchStatementProperty.SensitivityLevel`.
	SensitivityLevel *string `field:"optional" json:"sensitivityLevel" yaml:"sensitivityLevel"`
}
