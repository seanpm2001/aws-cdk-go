package awswafv2


// A rule statement used to run the rules that are defined in a managed rule group.
//
// To use this, provide the vendor name and the name of the rule group in this statement.
//
// You cannot nest a `ManagedRuleGroupStatement` , for example for use inside a `NotStatement` or `OrStatement` . It can only be referenced as a top-level statement within a rule.
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
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   managedRuleGroupStatementProperty := &managedRuleGroupStatementProperty{
//   	name: jsii.String("name"),
//   	vendorName: jsii.String("vendorName"),
//
//   	// the properties below are optional
//   	excludedRules: []interface{}{
//   		&excludedRuleProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   	managedRuleGroupConfigs: []interface{}{
//   		&managedRuleGroupConfigProperty{
//   			loginPath: jsii.String("loginPath"),
//   			passwordField: &fieldIdentifierProperty{
//   				identifier: jsii.String("identifier"),
//   			},
//   			payloadType: jsii.String("payloadType"),
//   			usernameField: &fieldIdentifierProperty{
//   				identifier: jsii.String("identifier"),
//   			},
//   		},
//   	},
//   	scopeDownStatement: &statementProperty{
//   		andStatement: &andStatementProperty{
//   			statements: []interface{}{
//   				statementProperty_,
//   			},
//   		},
//   		byteMatchStatement: &byteMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: &singleHeaderProperty{
//   					name: jsii.String("name"),
//   				},
//   				singleQueryArgument: &singleQueryArgumentProperty{
//   					name: jsii.String("name"),
//   				},
//   				uriPath: uriPath,
//   			},
//   			positionalConstraint: jsii.String("positionalConstraint"),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			searchString: jsii.String("searchString"),
//   			searchStringBase64: jsii.String("searchStringBase64"),
//   		},
//   		geoMatchStatement: &geoMatchStatementProperty{
//   			countryCodes: []*string{
//   				jsii.String("countryCodes"),
//   			},
//   			forwardedIpConfig: &forwardedIPConfigurationProperty{
//   				fallbackBehavior: jsii.String("fallbackBehavior"),
//   				headerName: jsii.String("headerName"),
//   			},
//   		},
//   		ipSetReferenceStatement: map[string]interface{}{
//   			"arn": jsii.String("arn"),
//
//   			// the properties below are optional
//   			"ipSetForwardedIpConfig": map[string]*string{
//   				"fallbackBehavior": jsii.String("fallbackBehavior"),
//   				"headerName": jsii.String("headerName"),
//   				"position": jsii.String("position"),
//   			},
//   		},
//   		labelMatchStatement: &labelMatchStatementProperty{
//   			key: jsii.String("key"),
//   			scope: jsii.String("scope"),
//   		},
//   		managedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   			name: jsii.String("name"),
//   			vendorName: jsii.String("vendorName"),
//
//   			// the properties below are optional
//   			excludedRules: []interface{}{
//   				&excludedRuleProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   			managedRuleGroupConfigs: []interface{}{
//   				&managedRuleGroupConfigProperty{
//   					loginPath: jsii.String("loginPath"),
//   					passwordField: &fieldIdentifierProperty{
//   						identifier: jsii.String("identifier"),
//   					},
//   					payloadType: jsii.String("payloadType"),
//   					usernameField: &fieldIdentifierProperty{
//   						identifier: jsii.String("identifier"),
//   					},
//   				},
//   			},
//   			scopeDownStatement: statementProperty_,
//   			version: jsii.String("version"),
//   		},
//   		notStatement: &notStatementProperty{
//   			statement: statementProperty_,
//   		},
//   		orStatement: &orStatementProperty{
//   			statements: []interface{}{
//   				statementProperty_,
//   			},
//   		},
//   		rateBasedStatement: &rateBasedStatementProperty{
//   			aggregateKeyType: jsii.String("aggregateKeyType"),
//   			limit: jsii.Number(123),
//
//   			// the properties below are optional
//   			forwardedIpConfig: &forwardedIPConfigurationProperty{
//   				fallbackBehavior: jsii.String("fallbackBehavior"),
//   				headerName: jsii.String("headerName"),
//   			},
//   			scopeDownStatement: statementProperty_,
//   		},
//   		regexMatchStatement: &regexMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: &singleHeaderProperty{
//   					name: jsii.String("name"),
//   				},
//   				singleQueryArgument: &singleQueryArgumentProperty{
//   					name: jsii.String("name"),
//   				},
//   				uriPath: uriPath,
//   			},
//   			regexString: jsii.String("regexString"),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   			arn: jsii.String("arn"),
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: &singleHeaderProperty{
//   					name: jsii.String("name"),
//   				},
//   				singleQueryArgument: &singleQueryArgumentProperty{
//   					name: jsii.String("name"),
//   				},
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		ruleGroupReferenceStatement: &ruleGroupReferenceStatementProperty{
//   			arn: jsii.String("arn"),
//
//   			// the properties below are optional
//   			excludedRules: []interface{}{
//   				&excludedRuleProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   		},
//   		sizeConstraintStatement: &sizeConstraintStatementProperty{
//   			comparisonOperator: jsii.String("comparisonOperator"),
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: &singleHeaderProperty{
//   					name: jsii.String("name"),
//   				},
//   				singleQueryArgument: &singleQueryArgumentProperty{
//   					name: jsii.String("name"),
//   				},
//   				uriPath: uriPath,
//   			},
//   			size: jsii.Number(123),
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		sqliMatchStatement: &sqliMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: &singleHeaderProperty{
//   					name: jsii.String("name"),
//   				},
//   				singleQueryArgument: &singleQueryArgumentProperty{
//   					name: jsii.String("name"),
//   				},
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			sensitivityLevel: jsii.String("sensitivityLevel"),
//   		},
//   		xssMatchStatement: &xssMatchStatementProperty{
//   			fieldToMatch: &fieldToMatchProperty{
//   				allQueryArguments: allQueryArguments,
//   				body: &bodyProperty{
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				cookies: &cookiesProperty{
//   					matchPattern: &cookieMatchPatternProperty{
//   						all: all,
//   						excludedCookies: []*string{
//   							jsii.String("excludedCookies"),
//   						},
//   						includedCookies: []*string{
//   							jsii.String("includedCookies"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				headers: &headersProperty{
//   					matchPattern: &headerMatchPatternProperty{
//   						all: all,
//   						excludedHeaders: []*string{
//   							jsii.String("excludedHeaders"),
//   						},
//   						includedHeaders: []*string{
//   							jsii.String("includedHeaders"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				jsonBody: &jsonBodyProperty{
//   					matchPattern: &jsonMatchPatternProperty{
//   						all: all,
//   						includedPaths: []*string{
//   							jsii.String("includedPaths"),
//   						},
//   					},
//   					matchScope: jsii.String("matchScope"),
//
//   					// the properties below are optional
//   					invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   					oversizeHandling: jsii.String("oversizeHandling"),
//   				},
//   				method: method,
//   				queryString: queryString,
//   				singleHeader: &singleHeaderProperty{
//   					name: jsii.String("name"),
//   				},
//   				singleQueryArgument: &singleQueryArgumentProperty{
//   					name: jsii.String("name"),
//   				},
//   				uriPath: uriPath,
//   			},
//   			textTransformations: []interface{}{
//   				&textTransformationProperty{
//   					priority: jsii.Number(123),
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   	version: jsii.String("version"),
//   }
//
type CfnWebACL_ManagedRuleGroupStatementProperty struct {
	// The name of the managed rule group.
	//
	// You use this, along with the vendor name, to identify the rule group.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the managed rule group vendor.
	//
	// You use this, along with the rule group name, to identify the rule group.
	VendorName *string `field:"required" json:"vendorName" yaml:"vendorName"`
	// Rules in the referenced rule group whose actions are set to `Count` .
	//
	// > Instead of this option, use `RuleActionOverrides` . It accepts any valid action setting, including `Count` .
	ExcludedRules interface{} `field:"optional" json:"excludedRules" yaml:"excludedRules"`
	// Additional information that's used by a managed rule group. Many managed rule groups don't require this.
	//
	// Use the `AWSManagedRulesBotControlRuleSet` configuration object to configure the protection level that you want the Bot Control rule group to use.
	ManagedRuleGroupConfigs interface{} `field:"optional" json:"managedRuleGroupConfigs" yaml:"managedRuleGroupConfigs"`
	// An optional nested statement that narrows the scope of the web requests that are evaluated by the managed rule group.
	//
	// Requests are only evaluated by the rule group if they match the scope-down statement. You can use any nestable `Statement` in the scope-down statement, and you can nest statements at any level, the same as you can for a rule statement.
	ScopeDownStatement interface{} `field:"optional" json:"scopeDownStatement" yaml:"scopeDownStatement"`
	// The version of the managed rule group to use.
	//
	// If you specify this, the version setting is fixed until you change it. If you don't specify this, AWS WAF uses the vendor's default version, and then keeps the version at the vendor's default when the vendor updates the managed rule group settings.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

