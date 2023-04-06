package awswafv2


// The processing guidance for a rule, used by AWS WAF to determine whether a web request matches the rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allQueryArguments interface{}
//   var andStatementProperty_ andStatementProperty
//   var managedRuleGroupStatementProperty_ managedRuleGroupStatementProperty
//   var method interface{}
//   var notStatementProperty_ notStatementProperty
//   var orStatementProperty_ orStatementProperty
//   var queryString interface{}
//   var rateBasedStatementProperty_ rateBasedStatementProperty
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   statementProperty := &statementProperty{
//   	AndStatement: &andStatementProperty{
//   		Statements: []interface{}{
//   			&statementProperty{
//   				AndStatement: andStatementProperty_,
//   				ByteMatchStatement: &ByteMatchStatementProperty{
//   					FieldToMatch: &FieldToMatchProperty{
//   						AllQueryArguments: allQueryArguments,
//   						Body: &BodyProperty{
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Cookies: &CookiesProperty{
//   							MatchPattern: &CookieMatchPatternProperty{
//   								All: all,
//   								ExcludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								IncludedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Headers: &HeadersProperty{
//   							MatchPattern: &HeaderMatchPatternProperty{
//   								All: all,
//   								ExcludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								IncludedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						JsonBody: &JsonBodyProperty{
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
//   						UriPath: uriPath,
//   					},
//   					PositionalConstraint: jsii.String("positionalConstraint"),
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					SearchString: jsii.String("searchString"),
//   					SearchStringBase64: jsii.String("searchStringBase64"),
//   				},
//   				GeoMatchStatement: &GeoMatchStatementProperty{
//   					CountryCodes: []*string{
//   						jsii.String("countryCodes"),
//   					},
//   					ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
//   						HeaderName: jsii.String("headerName"),
//   					},
//   				},
//   				IpSetReferenceStatement: map[string]interface{}{
//   					"arn": jsii.String("arn"),
//
//   					// the properties below are optional
//   					"ipSetForwardedIpConfig": map[string]*string{
//   						"fallbackBehavior": jsii.String("fallbackBehavior"),
//   						"headerName": jsii.String("headerName"),
//   						"position": jsii.String("position"),
//   					},
//   				},
//   				LabelMatchStatement: &LabelMatchStatementProperty{
//   					Key: jsii.String("key"),
//   					Scope: jsii.String("scope"),
//   				},
//   				ManagedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   					Name: jsii.String("name"),
//   					VendorName: jsii.String("vendorName"),
//
//   					// the properties below are optional
//   					ExcludedRules: []interface{}{
//   						&ExcludedRuleProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					ManagedRuleGroupConfigs: []interface{}{
//   						&ManagedRuleGroupConfigProperty{
//   							AwsManagedRulesAtpRuleSet: &AWSManagedRulesATPRuleSetProperty{
//   								LoginPath: jsii.String("loginPath"),
//
//   								// the properties below are optional
//   								RequestInspection: &RequestInspectionProperty{
//   									PasswordField: &FieldIdentifierProperty{
//   										Identifier: jsii.String("identifier"),
//   									},
//   									PayloadType: jsii.String("payloadType"),
//   									UsernameField: &FieldIdentifierProperty{
//   										Identifier: jsii.String("identifier"),
//   									},
//   								},
//   								ResponseInspection: &ResponseInspectionProperty{
//   									BodyContains: &ResponseInspectionBodyContainsProperty{
//   										FailureStrings: []*string{
//   											jsii.String("failureStrings"),
//   										},
//   										SuccessStrings: []*string{
//   											jsii.String("successStrings"),
//   										},
//   									},
//   									Header: &ResponseInspectionHeaderProperty{
//   										FailureValues: []*string{
//   											jsii.String("failureValues"),
//   										},
//   										Name: jsii.String("name"),
//   										SuccessValues: []*string{
//   											jsii.String("successValues"),
//   										},
//   									},
//   									Json: &ResponseInspectionJsonProperty{
//   										FailureValues: []*string{
//   											jsii.String("failureValues"),
//   										},
//   										Identifier: jsii.String("identifier"),
//   										SuccessValues: []*string{
//   											jsii.String("successValues"),
//   										},
//   									},
//   									StatusCode: &ResponseInspectionStatusCodeProperty{
//   										FailureCodes: []interface{}{
//   											jsii.Number(123),
//   										},
//   										SuccessCodes: []interface{}{
//   											jsii.Number(123),
//   										},
//   									},
//   								},
//   							},
//   							AwsManagedRulesBotControlRuleSet: &AWSManagedRulesBotControlRuleSetProperty{
//   								InspectionLevel: jsii.String("inspectionLevel"),
//   							},
//   							LoginPath: jsii.String("loginPath"),
//   							PasswordField: &FieldIdentifierProperty{
//   								Identifier: jsii.String("identifier"),
//   							},
//   							PayloadType: jsii.String("payloadType"),
//   							UsernameField: &FieldIdentifierProperty{
//   								Identifier: jsii.String("identifier"),
//   							},
//   						},
//   					},
//   					RuleActionOverrides: []interface{}{
//   						&RuleActionOverrideProperty{
//   							ActionToUse: &RuleActionProperty{
//   								Allow: &AllowActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Block: &BlockActionProperty{
//   									CustomResponse: &CustomResponseProperty{
//   										ResponseCode: jsii.Number(123),
//
//   										// the properties below are optional
//   										CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   										ResponseHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Captcha: &CaptchaActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Challenge: &ChallengeActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Count: &CountActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					ScopeDownStatement: statementProperty_,
//   					Version: jsii.String("version"),
//   				},
//   				NotStatement: &notStatementProperty{
//   					Statement: statementProperty_,
//   				},
//   				OrStatement: &orStatementProperty{
//   					Statements: []interface{}{
//   						statementProperty_,
//   					},
//   				},
//   				RateBasedStatement: &rateBasedStatementProperty{
//   					AggregateKeyType: jsii.String("aggregateKeyType"),
//   					Limit: jsii.Number(123),
//
//   					// the properties below are optional
//   					ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
//   						HeaderName: jsii.String("headerName"),
//   					},
//   					ScopeDownStatement: statementProperty_,
//   				},
//   				RegexMatchStatement: &RegexMatchStatementProperty{
//   					FieldToMatch: &FieldToMatchProperty{
//   						AllQueryArguments: allQueryArguments,
//   						Body: &BodyProperty{
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Cookies: &CookiesProperty{
//   							MatchPattern: &CookieMatchPatternProperty{
//   								All: all,
//   								ExcludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								IncludedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Headers: &HeadersProperty{
//   							MatchPattern: &HeaderMatchPatternProperty{
//   								All: all,
//   								ExcludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								IncludedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						JsonBody: &JsonBodyProperty{
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
//   						UriPath: uriPath,
//   					},
//   					RegexString: jsii.String("regexString"),
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				RegexPatternSetReferenceStatement: &RegexPatternSetReferenceStatementProperty{
//   					Arn: jsii.String("arn"),
//   					FieldToMatch: &FieldToMatchProperty{
//   						AllQueryArguments: allQueryArguments,
//   						Body: &BodyProperty{
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Cookies: &CookiesProperty{
//   							MatchPattern: &CookieMatchPatternProperty{
//   								All: all,
//   								ExcludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								IncludedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Headers: &HeadersProperty{
//   							MatchPattern: &HeaderMatchPatternProperty{
//   								All: all,
//   								ExcludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								IncludedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						JsonBody: &JsonBodyProperty{
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
//   						UriPath: uriPath,
//   					},
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				RuleGroupReferenceStatement: &RuleGroupReferenceStatementProperty{
//   					Arn: jsii.String("arn"),
//
//   					// the properties below are optional
//   					ExcludedRules: []interface{}{
//   						&ExcludedRuleProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					RuleActionOverrides: []interface{}{
//   						&RuleActionOverrideProperty{
//   							ActionToUse: &RuleActionProperty{
//   								Allow: &AllowActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Block: &BlockActionProperty{
//   									CustomResponse: &CustomResponseProperty{
//   										ResponseCode: jsii.Number(123),
//
//   										// the properties below are optional
//   										CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   										ResponseHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Captcha: &CaptchaActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Challenge: &ChallengeActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Count: &CountActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//   							Name: jsii.String("name"),
//   						},
//   					},
//   				},
//   				SizeConstraintStatement: &SizeConstraintStatementProperty{
//   					ComparisonOperator: jsii.String("comparisonOperator"),
//   					FieldToMatch: &FieldToMatchProperty{
//   						AllQueryArguments: allQueryArguments,
//   						Body: &BodyProperty{
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Cookies: &CookiesProperty{
//   							MatchPattern: &CookieMatchPatternProperty{
//   								All: all,
//   								ExcludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								IncludedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Headers: &HeadersProperty{
//   							MatchPattern: &HeaderMatchPatternProperty{
//   								All: all,
//   								ExcludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								IncludedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						JsonBody: &JsonBodyProperty{
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
//   						UriPath: uriPath,
//   					},
//   					Size: jsii.Number(123),
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				SqliMatchStatement: &SqliMatchStatementProperty{
//   					FieldToMatch: &FieldToMatchProperty{
//   						AllQueryArguments: allQueryArguments,
//   						Body: &BodyProperty{
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Cookies: &CookiesProperty{
//   							MatchPattern: &CookieMatchPatternProperty{
//   								All: all,
//   								ExcludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								IncludedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Headers: &HeadersProperty{
//   							MatchPattern: &HeaderMatchPatternProperty{
//   								All: all,
//   								ExcludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								IncludedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						JsonBody: &JsonBodyProperty{
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
//   						UriPath: uriPath,
//   					},
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					SensitivityLevel: jsii.String("sensitivityLevel"),
//   				},
//   				XssMatchStatement: &XssMatchStatementProperty{
//   					FieldToMatch: &FieldToMatchProperty{
//   						AllQueryArguments: allQueryArguments,
//   						Body: &BodyProperty{
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Cookies: &CookiesProperty{
//   							MatchPattern: &CookieMatchPatternProperty{
//   								All: all,
//   								ExcludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								IncludedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Headers: &HeadersProperty{
//   							MatchPattern: &HeaderMatchPatternProperty{
//   								All: all,
//   								ExcludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								IncludedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						JsonBody: &JsonBodyProperty{
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
//   						UriPath: uriPath,
//   					},
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	ByteMatchStatement: &ByteMatchStatementProperty{
//   		FieldToMatch: &FieldToMatchProperty{
//   			AllQueryArguments: allQueryArguments,
//   			Body: &BodyProperty{
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Cookies: &CookiesProperty{
//   				MatchPattern: &CookieMatchPatternProperty{
//   					All: all,
//   					ExcludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					IncludedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Headers: &HeadersProperty{
//   				MatchPattern: &HeaderMatchPatternProperty{
//   					All: all,
//   					ExcludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					IncludedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			JsonBody: &JsonBodyProperty{
//   				MatchPattern: &JsonMatchPatternProperty{
//   					All: all,
//   					IncludedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Method: method,
//   			QueryString: queryString,
//   			SingleHeader: singleHeader,
//   			SingleQueryArgument: singleQueryArgument,
//   			UriPath: uriPath,
//   		},
//   		PositionalConstraint: jsii.String("positionalConstraint"),
//   		TextTransformations: []interface{}{
//   			&TextTransformationProperty{
//   				Priority: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		SearchString: jsii.String("searchString"),
//   		SearchStringBase64: jsii.String("searchStringBase64"),
//   	},
//   	GeoMatchStatement: &GeoMatchStatementProperty{
//   		CountryCodes: []*string{
//   			jsii.String("countryCodes"),
//   		},
//   		ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   			FallbackBehavior: jsii.String("fallbackBehavior"),
//   			HeaderName: jsii.String("headerName"),
//   		},
//   	},
//   	IpSetReferenceStatement: map[string]interface{}{
//   		"arn": jsii.String("arn"),
//
//   		// the properties below are optional
//   		"ipSetForwardedIpConfig": map[string]*string{
//   			"fallbackBehavior": jsii.String("fallbackBehavior"),
//   			"headerName": jsii.String("headerName"),
//   			"position": jsii.String("position"),
//   		},
//   	},
//   	LabelMatchStatement: &LabelMatchStatementProperty{
//   		Key: jsii.String("key"),
//   		Scope: jsii.String("scope"),
//   	},
//   	ManagedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   		Name: jsii.String("name"),
//   		VendorName: jsii.String("vendorName"),
//
//   		// the properties below are optional
//   		ExcludedRules: []interface{}{
//   			&ExcludedRuleProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		ManagedRuleGroupConfigs: []interface{}{
//   			&ManagedRuleGroupConfigProperty{
//   				AwsManagedRulesAtpRuleSet: &AWSManagedRulesATPRuleSetProperty{
//   					LoginPath: jsii.String("loginPath"),
//
//   					// the properties below are optional
//   					RequestInspection: &RequestInspectionProperty{
//   						PasswordField: &FieldIdentifierProperty{
//   							Identifier: jsii.String("identifier"),
//   						},
//   						PayloadType: jsii.String("payloadType"),
//   						UsernameField: &FieldIdentifierProperty{
//   							Identifier: jsii.String("identifier"),
//   						},
//   					},
//   					ResponseInspection: &ResponseInspectionProperty{
//   						BodyContains: &ResponseInspectionBodyContainsProperty{
//   							FailureStrings: []*string{
//   								jsii.String("failureStrings"),
//   							},
//   							SuccessStrings: []*string{
//   								jsii.String("successStrings"),
//   							},
//   						},
//   						Header: &ResponseInspectionHeaderProperty{
//   							FailureValues: []*string{
//   								jsii.String("failureValues"),
//   							},
//   							Name: jsii.String("name"),
//   							SuccessValues: []*string{
//   								jsii.String("successValues"),
//   							},
//   						},
//   						Json: &ResponseInspectionJsonProperty{
//   							FailureValues: []*string{
//   								jsii.String("failureValues"),
//   							},
//   							Identifier: jsii.String("identifier"),
//   							SuccessValues: []*string{
//   								jsii.String("successValues"),
//   							},
//   						},
//   						StatusCode: &ResponseInspectionStatusCodeProperty{
//   							FailureCodes: []interface{}{
//   								jsii.Number(123),
//   							},
//   							SuccessCodes: []interface{}{
//   								jsii.Number(123),
//   							},
//   						},
//   					},
//   				},
//   				AwsManagedRulesBotControlRuleSet: &AWSManagedRulesBotControlRuleSetProperty{
//   					InspectionLevel: jsii.String("inspectionLevel"),
//   				},
//   				LoginPath: jsii.String("loginPath"),
//   				PasswordField: &FieldIdentifierProperty{
//   					Identifier: jsii.String("identifier"),
//   				},
//   				PayloadType: jsii.String("payloadType"),
//   				UsernameField: &FieldIdentifierProperty{
//   					Identifier: jsii.String("identifier"),
//   				},
//   			},
//   		},
//   		RuleActionOverrides: []interface{}{
//   			&RuleActionOverrideProperty{
//   				ActionToUse: &RuleActionProperty{
//   					Allow: &AllowActionProperty{
//   						CustomRequestHandling: &CustomRequestHandlingProperty{
//   							InsertHeaders: []interface{}{
//   								&CustomHTTPHeaderProperty{
//   									Name: jsii.String("name"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   					Block: &BlockActionProperty{
//   						CustomResponse: &CustomResponseProperty{
//   							ResponseCode: jsii.Number(123),
//
//   							// the properties below are optional
//   							CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   							ResponseHeaders: []interface{}{
//   								&CustomHTTPHeaderProperty{
//   									Name: jsii.String("name"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   					Captcha: &CaptchaActionProperty{
//   						CustomRequestHandling: &CustomRequestHandlingProperty{
//   							InsertHeaders: []interface{}{
//   								&CustomHTTPHeaderProperty{
//   									Name: jsii.String("name"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   					Challenge: &ChallengeActionProperty{
//   						CustomRequestHandling: &CustomRequestHandlingProperty{
//   							InsertHeaders: []interface{}{
//   								&CustomHTTPHeaderProperty{
//   									Name: jsii.String("name"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   					Count: &CountActionProperty{
//   						CustomRequestHandling: &CustomRequestHandlingProperty{
//   							InsertHeaders: []interface{}{
//   								&CustomHTTPHeaderProperty{
//   									Name: jsii.String("name"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		ScopeDownStatement: &statementProperty{
//   			AndStatement: &andStatementProperty{
//   				Statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			ByteMatchStatement: &ByteMatchStatementProperty{
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				PositionalConstraint: jsii.String("positionalConstraint"),
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				SearchString: jsii.String("searchString"),
//   				SearchStringBase64: jsii.String("searchStringBase64"),
//   			},
//   			GeoMatchStatement: &GeoMatchStatementProperty{
//   				CountryCodes: []*string{
//   					jsii.String("countryCodes"),
//   				},
//   				ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   					HeaderName: jsii.String("headerName"),
//   				},
//   			},
//   			IpSetReferenceStatement: map[string]interface{}{
//   				"arn": jsii.String("arn"),
//
//   				// the properties below are optional
//   				"ipSetForwardedIpConfig": map[string]*string{
//   					"fallbackBehavior": jsii.String("fallbackBehavior"),
//   					"headerName": jsii.String("headerName"),
//   					"position": jsii.String("position"),
//   				},
//   			},
//   			LabelMatchStatement: &LabelMatchStatementProperty{
//   				Key: jsii.String("key"),
//   				Scope: jsii.String("scope"),
//   			},
//   			ManagedRuleGroupStatement: managedRuleGroupStatementProperty_,
//   			NotStatement: &notStatementProperty{
//   				Statement: statementProperty_,
//   			},
//   			OrStatement: &orStatementProperty{
//   				Statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			RateBasedStatement: &rateBasedStatementProperty{
//   				AggregateKeyType: jsii.String("aggregateKeyType"),
//   				Limit: jsii.Number(123),
//
//   				// the properties below are optional
//   				ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   					HeaderName: jsii.String("headerName"),
//   				},
//   				ScopeDownStatement: statementProperty_,
//   			},
//   			RegexMatchStatement: &RegexMatchStatementProperty{
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				RegexString: jsii.String("regexString"),
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			RegexPatternSetReferenceStatement: &RegexPatternSetReferenceStatementProperty{
//   				Arn: jsii.String("arn"),
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			RuleGroupReferenceStatement: &RuleGroupReferenceStatementProperty{
//   				Arn: jsii.String("arn"),
//
//   				// the properties below are optional
//   				ExcludedRules: []interface{}{
//   					&ExcludedRuleProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				RuleActionOverrides: []interface{}{
//   					&RuleActionOverrideProperty{
//   						ActionToUse: &RuleActionProperty{
//   							Allow: &AllowActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Block: &BlockActionProperty{
//   								CustomResponse: &CustomResponseProperty{
//   									ResponseCode: jsii.Number(123),
//
//   									// the properties below are optional
//   									CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   									ResponseHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Captcha: &CaptchaActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Challenge: &ChallengeActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Count: &CountActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   			SizeConstraintStatement: &SizeConstraintStatementProperty{
//   				ComparisonOperator: jsii.String("comparisonOperator"),
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				Size: jsii.Number(123),
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			SqliMatchStatement: &SqliMatchStatementProperty{
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				SensitivityLevel: jsii.String("sensitivityLevel"),
//   			},
//   			XssMatchStatement: &XssMatchStatementProperty{
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   		},
//   		Version: jsii.String("version"),
//   	},
//   	NotStatement: &notStatementProperty{
//   		Statement: &statementProperty{
//   			AndStatement: &andStatementProperty{
//   				Statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			ByteMatchStatement: &ByteMatchStatementProperty{
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				PositionalConstraint: jsii.String("positionalConstraint"),
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				SearchString: jsii.String("searchString"),
//   				SearchStringBase64: jsii.String("searchStringBase64"),
//   			},
//   			GeoMatchStatement: &GeoMatchStatementProperty{
//   				CountryCodes: []*string{
//   					jsii.String("countryCodes"),
//   				},
//   				ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   					HeaderName: jsii.String("headerName"),
//   				},
//   			},
//   			IpSetReferenceStatement: map[string]interface{}{
//   				"arn": jsii.String("arn"),
//
//   				// the properties below are optional
//   				"ipSetForwardedIpConfig": map[string]*string{
//   					"fallbackBehavior": jsii.String("fallbackBehavior"),
//   					"headerName": jsii.String("headerName"),
//   					"position": jsii.String("position"),
//   				},
//   			},
//   			LabelMatchStatement: &LabelMatchStatementProperty{
//   				Key: jsii.String("key"),
//   				Scope: jsii.String("scope"),
//   			},
//   			ManagedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   				Name: jsii.String("name"),
//   				VendorName: jsii.String("vendorName"),
//
//   				// the properties below are optional
//   				ExcludedRules: []interface{}{
//   					&ExcludedRuleProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				ManagedRuleGroupConfigs: []interface{}{
//   					&ManagedRuleGroupConfigProperty{
//   						AwsManagedRulesAtpRuleSet: &AWSManagedRulesATPRuleSetProperty{
//   							LoginPath: jsii.String("loginPath"),
//
//   							// the properties below are optional
//   							RequestInspection: &RequestInspectionProperty{
//   								PasswordField: &FieldIdentifierProperty{
//   									Identifier: jsii.String("identifier"),
//   								},
//   								PayloadType: jsii.String("payloadType"),
//   								UsernameField: &FieldIdentifierProperty{
//   									Identifier: jsii.String("identifier"),
//   								},
//   							},
//   							ResponseInspection: &ResponseInspectionProperty{
//   								BodyContains: &ResponseInspectionBodyContainsProperty{
//   									FailureStrings: []*string{
//   										jsii.String("failureStrings"),
//   									},
//   									SuccessStrings: []*string{
//   										jsii.String("successStrings"),
//   									},
//   								},
//   								Header: &ResponseInspectionHeaderProperty{
//   									FailureValues: []*string{
//   										jsii.String("failureValues"),
//   									},
//   									Name: jsii.String("name"),
//   									SuccessValues: []*string{
//   										jsii.String("successValues"),
//   									},
//   								},
//   								Json: &ResponseInspectionJsonProperty{
//   									FailureValues: []*string{
//   										jsii.String("failureValues"),
//   									},
//   									Identifier: jsii.String("identifier"),
//   									SuccessValues: []*string{
//   										jsii.String("successValues"),
//   									},
//   								},
//   								StatusCode: &ResponseInspectionStatusCodeProperty{
//   									FailureCodes: []interface{}{
//   										jsii.Number(123),
//   									},
//   									SuccessCodes: []interface{}{
//   										jsii.Number(123),
//   									},
//   								},
//   							},
//   						},
//   						AwsManagedRulesBotControlRuleSet: &AWSManagedRulesBotControlRuleSetProperty{
//   							InspectionLevel: jsii.String("inspectionLevel"),
//   						},
//   						LoginPath: jsii.String("loginPath"),
//   						PasswordField: &FieldIdentifierProperty{
//   							Identifier: jsii.String("identifier"),
//   						},
//   						PayloadType: jsii.String("payloadType"),
//   						UsernameField: &FieldIdentifierProperty{
//   							Identifier: jsii.String("identifier"),
//   						},
//   					},
//   				},
//   				RuleActionOverrides: []interface{}{
//   					&RuleActionOverrideProperty{
//   						ActionToUse: &RuleActionProperty{
//   							Allow: &AllowActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Block: &BlockActionProperty{
//   								CustomResponse: &CustomResponseProperty{
//   									ResponseCode: jsii.Number(123),
//
//   									// the properties below are optional
//   									CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   									ResponseHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Captcha: &CaptchaActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Challenge: &ChallengeActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Count: &CountActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				ScopeDownStatement: statementProperty_,
//   				Version: jsii.String("version"),
//   			},
//   			NotStatement: notStatementProperty_,
//   			OrStatement: &orStatementProperty{
//   				Statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			RateBasedStatement: &rateBasedStatementProperty{
//   				AggregateKeyType: jsii.String("aggregateKeyType"),
//   				Limit: jsii.Number(123),
//
//   				// the properties below are optional
//   				ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   					HeaderName: jsii.String("headerName"),
//   				},
//   				ScopeDownStatement: statementProperty_,
//   			},
//   			RegexMatchStatement: &RegexMatchStatementProperty{
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				RegexString: jsii.String("regexString"),
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			RegexPatternSetReferenceStatement: &RegexPatternSetReferenceStatementProperty{
//   				Arn: jsii.String("arn"),
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			RuleGroupReferenceStatement: &RuleGroupReferenceStatementProperty{
//   				Arn: jsii.String("arn"),
//
//   				// the properties below are optional
//   				ExcludedRules: []interface{}{
//   					&ExcludedRuleProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				RuleActionOverrides: []interface{}{
//   					&RuleActionOverrideProperty{
//   						ActionToUse: &RuleActionProperty{
//   							Allow: &AllowActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Block: &BlockActionProperty{
//   								CustomResponse: &CustomResponseProperty{
//   									ResponseCode: jsii.Number(123),
//
//   									// the properties below are optional
//   									CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   									ResponseHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Captcha: &CaptchaActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Challenge: &ChallengeActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Count: &CountActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   			SizeConstraintStatement: &SizeConstraintStatementProperty{
//   				ComparisonOperator: jsii.String("comparisonOperator"),
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				Size: jsii.Number(123),
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			SqliMatchStatement: &SqliMatchStatementProperty{
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				SensitivityLevel: jsii.String("sensitivityLevel"),
//   			},
//   			XssMatchStatement: &XssMatchStatementProperty{
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	OrStatement: &orStatementProperty{
//   		Statements: []interface{}{
//   			&statementProperty{
//   				AndStatement: &andStatementProperty{
//   					Statements: []interface{}{
//   						statementProperty_,
//   					},
//   				},
//   				ByteMatchStatement: &ByteMatchStatementProperty{
//   					FieldToMatch: &FieldToMatchProperty{
//   						AllQueryArguments: allQueryArguments,
//   						Body: &BodyProperty{
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Cookies: &CookiesProperty{
//   							MatchPattern: &CookieMatchPatternProperty{
//   								All: all,
//   								ExcludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								IncludedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Headers: &HeadersProperty{
//   							MatchPattern: &HeaderMatchPatternProperty{
//   								All: all,
//   								ExcludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								IncludedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						JsonBody: &JsonBodyProperty{
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
//   						UriPath: uriPath,
//   					},
//   					PositionalConstraint: jsii.String("positionalConstraint"),
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					SearchString: jsii.String("searchString"),
//   					SearchStringBase64: jsii.String("searchStringBase64"),
//   				},
//   				GeoMatchStatement: &GeoMatchStatementProperty{
//   					CountryCodes: []*string{
//   						jsii.String("countryCodes"),
//   					},
//   					ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
//   						HeaderName: jsii.String("headerName"),
//   					},
//   				},
//   				IpSetReferenceStatement: map[string]interface{}{
//   					"arn": jsii.String("arn"),
//
//   					// the properties below are optional
//   					"ipSetForwardedIpConfig": map[string]*string{
//   						"fallbackBehavior": jsii.String("fallbackBehavior"),
//   						"headerName": jsii.String("headerName"),
//   						"position": jsii.String("position"),
//   					},
//   				},
//   				LabelMatchStatement: &LabelMatchStatementProperty{
//   					Key: jsii.String("key"),
//   					Scope: jsii.String("scope"),
//   				},
//   				ManagedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   					Name: jsii.String("name"),
//   					VendorName: jsii.String("vendorName"),
//
//   					// the properties below are optional
//   					ExcludedRules: []interface{}{
//   						&ExcludedRuleProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					ManagedRuleGroupConfigs: []interface{}{
//   						&ManagedRuleGroupConfigProperty{
//   							AwsManagedRulesAtpRuleSet: &AWSManagedRulesATPRuleSetProperty{
//   								LoginPath: jsii.String("loginPath"),
//
//   								// the properties below are optional
//   								RequestInspection: &RequestInspectionProperty{
//   									PasswordField: &FieldIdentifierProperty{
//   										Identifier: jsii.String("identifier"),
//   									},
//   									PayloadType: jsii.String("payloadType"),
//   									UsernameField: &FieldIdentifierProperty{
//   										Identifier: jsii.String("identifier"),
//   									},
//   								},
//   								ResponseInspection: &ResponseInspectionProperty{
//   									BodyContains: &ResponseInspectionBodyContainsProperty{
//   										FailureStrings: []*string{
//   											jsii.String("failureStrings"),
//   										},
//   										SuccessStrings: []*string{
//   											jsii.String("successStrings"),
//   										},
//   									},
//   									Header: &ResponseInspectionHeaderProperty{
//   										FailureValues: []*string{
//   											jsii.String("failureValues"),
//   										},
//   										Name: jsii.String("name"),
//   										SuccessValues: []*string{
//   											jsii.String("successValues"),
//   										},
//   									},
//   									Json: &ResponseInspectionJsonProperty{
//   										FailureValues: []*string{
//   											jsii.String("failureValues"),
//   										},
//   										Identifier: jsii.String("identifier"),
//   										SuccessValues: []*string{
//   											jsii.String("successValues"),
//   										},
//   									},
//   									StatusCode: &ResponseInspectionStatusCodeProperty{
//   										FailureCodes: []interface{}{
//   											jsii.Number(123),
//   										},
//   										SuccessCodes: []interface{}{
//   											jsii.Number(123),
//   										},
//   									},
//   								},
//   							},
//   							AwsManagedRulesBotControlRuleSet: &AWSManagedRulesBotControlRuleSetProperty{
//   								InspectionLevel: jsii.String("inspectionLevel"),
//   							},
//   							LoginPath: jsii.String("loginPath"),
//   							PasswordField: &FieldIdentifierProperty{
//   								Identifier: jsii.String("identifier"),
//   							},
//   							PayloadType: jsii.String("payloadType"),
//   							UsernameField: &FieldIdentifierProperty{
//   								Identifier: jsii.String("identifier"),
//   							},
//   						},
//   					},
//   					RuleActionOverrides: []interface{}{
//   						&RuleActionOverrideProperty{
//   							ActionToUse: &RuleActionProperty{
//   								Allow: &AllowActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Block: &BlockActionProperty{
//   									CustomResponse: &CustomResponseProperty{
//   										ResponseCode: jsii.Number(123),
//
//   										// the properties below are optional
//   										CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   										ResponseHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Captcha: &CaptchaActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Challenge: &ChallengeActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Count: &CountActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					ScopeDownStatement: statementProperty_,
//   					Version: jsii.String("version"),
//   				},
//   				NotStatement: &notStatementProperty{
//   					Statement: statementProperty_,
//   				},
//   				OrStatement: orStatementProperty_,
//   				RateBasedStatement: &rateBasedStatementProperty{
//   					AggregateKeyType: jsii.String("aggregateKeyType"),
//   					Limit: jsii.Number(123),
//
//   					// the properties below are optional
//   					ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   						FallbackBehavior: jsii.String("fallbackBehavior"),
//   						HeaderName: jsii.String("headerName"),
//   					},
//   					ScopeDownStatement: statementProperty_,
//   				},
//   				RegexMatchStatement: &RegexMatchStatementProperty{
//   					FieldToMatch: &FieldToMatchProperty{
//   						AllQueryArguments: allQueryArguments,
//   						Body: &BodyProperty{
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Cookies: &CookiesProperty{
//   							MatchPattern: &CookieMatchPatternProperty{
//   								All: all,
//   								ExcludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								IncludedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Headers: &HeadersProperty{
//   							MatchPattern: &HeaderMatchPatternProperty{
//   								All: all,
//   								ExcludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								IncludedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						JsonBody: &JsonBodyProperty{
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
//   						UriPath: uriPath,
//   					},
//   					RegexString: jsii.String("regexString"),
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				RegexPatternSetReferenceStatement: &RegexPatternSetReferenceStatementProperty{
//   					Arn: jsii.String("arn"),
//   					FieldToMatch: &FieldToMatchProperty{
//   						AllQueryArguments: allQueryArguments,
//   						Body: &BodyProperty{
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Cookies: &CookiesProperty{
//   							MatchPattern: &CookieMatchPatternProperty{
//   								All: all,
//   								ExcludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								IncludedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Headers: &HeadersProperty{
//   							MatchPattern: &HeaderMatchPatternProperty{
//   								All: all,
//   								ExcludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								IncludedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						JsonBody: &JsonBodyProperty{
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
//   						UriPath: uriPath,
//   					},
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				RuleGroupReferenceStatement: &RuleGroupReferenceStatementProperty{
//   					Arn: jsii.String("arn"),
//
//   					// the properties below are optional
//   					ExcludedRules: []interface{}{
//   						&ExcludedRuleProperty{
//   							Name: jsii.String("name"),
//   						},
//   					},
//   					RuleActionOverrides: []interface{}{
//   						&RuleActionOverrideProperty{
//   							ActionToUse: &RuleActionProperty{
//   								Allow: &AllowActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Block: &BlockActionProperty{
//   									CustomResponse: &CustomResponseProperty{
//   										ResponseCode: jsii.Number(123),
//
//   										// the properties below are optional
//   										CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   										ResponseHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Captcha: &CaptchaActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Challenge: &ChallengeActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   								Count: &CountActionProperty{
//   									CustomRequestHandling: &CustomRequestHandlingProperty{
//   										InsertHeaders: []interface{}{
//   											&CustomHTTPHeaderProperty{
//   												Name: jsii.String("name"),
//   												Value: jsii.String("value"),
//   											},
//   										},
//   									},
//   								},
//   							},
//   							Name: jsii.String("name"),
//   						},
//   					},
//   				},
//   				SizeConstraintStatement: &SizeConstraintStatementProperty{
//   					ComparisonOperator: jsii.String("comparisonOperator"),
//   					FieldToMatch: &FieldToMatchProperty{
//   						AllQueryArguments: allQueryArguments,
//   						Body: &BodyProperty{
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Cookies: &CookiesProperty{
//   							MatchPattern: &CookieMatchPatternProperty{
//   								All: all,
//   								ExcludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								IncludedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Headers: &HeadersProperty{
//   							MatchPattern: &HeaderMatchPatternProperty{
//   								All: all,
//   								ExcludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								IncludedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						JsonBody: &JsonBodyProperty{
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
//   						UriPath: uriPath,
//   					},
//   					Size: jsii.Number(123),
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				SqliMatchStatement: &SqliMatchStatementProperty{
//   					FieldToMatch: &FieldToMatchProperty{
//   						AllQueryArguments: allQueryArguments,
//   						Body: &BodyProperty{
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Cookies: &CookiesProperty{
//   							MatchPattern: &CookieMatchPatternProperty{
//   								All: all,
//   								ExcludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								IncludedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Headers: &HeadersProperty{
//   							MatchPattern: &HeaderMatchPatternProperty{
//   								All: all,
//   								ExcludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								IncludedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						JsonBody: &JsonBodyProperty{
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
//   						UriPath: uriPath,
//   					},
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					SensitivityLevel: jsii.String("sensitivityLevel"),
//   				},
//   				XssMatchStatement: &XssMatchStatementProperty{
//   					FieldToMatch: &FieldToMatchProperty{
//   						AllQueryArguments: allQueryArguments,
//   						Body: &BodyProperty{
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Cookies: &CookiesProperty{
//   							MatchPattern: &CookieMatchPatternProperty{
//   								All: all,
//   								ExcludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								IncludedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Headers: &HeadersProperty{
//   							MatchPattern: &HeaderMatchPatternProperty{
//   								All: all,
//   								ExcludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								IncludedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						JsonBody: &JsonBodyProperty{
//   							MatchPattern: &JsonMatchPatternProperty{
//   								All: all,
//   								IncludedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							MatchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							OversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						Method: method,
//   						QueryString: queryString,
//   						SingleHeader: singleHeader,
//   						SingleQueryArgument: singleQueryArgument,
//   						UriPath: uriPath,
//   					},
//   					TextTransformations: []interface{}{
//   						&TextTransformationProperty{
//   							Priority: jsii.Number(123),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	RateBasedStatement: &rateBasedStatementProperty{
//   		AggregateKeyType: jsii.String("aggregateKeyType"),
//   		Limit: jsii.Number(123),
//
//   		// the properties below are optional
//   		ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   			FallbackBehavior: jsii.String("fallbackBehavior"),
//   			HeaderName: jsii.String("headerName"),
//   		},
//   		ScopeDownStatement: &statementProperty{
//   			AndStatement: &andStatementProperty{
//   				Statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			ByteMatchStatement: &ByteMatchStatementProperty{
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				PositionalConstraint: jsii.String("positionalConstraint"),
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				SearchString: jsii.String("searchString"),
//   				SearchStringBase64: jsii.String("searchStringBase64"),
//   			},
//   			GeoMatchStatement: &GeoMatchStatementProperty{
//   				CountryCodes: []*string{
//   					jsii.String("countryCodes"),
//   				},
//   				ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   					FallbackBehavior: jsii.String("fallbackBehavior"),
//   					HeaderName: jsii.String("headerName"),
//   				},
//   			},
//   			IpSetReferenceStatement: map[string]interface{}{
//   				"arn": jsii.String("arn"),
//
//   				// the properties below are optional
//   				"ipSetForwardedIpConfig": map[string]*string{
//   					"fallbackBehavior": jsii.String("fallbackBehavior"),
//   					"headerName": jsii.String("headerName"),
//   					"position": jsii.String("position"),
//   				},
//   			},
//   			LabelMatchStatement: &LabelMatchStatementProperty{
//   				Key: jsii.String("key"),
//   				Scope: jsii.String("scope"),
//   			},
//   			ManagedRuleGroupStatement: &managedRuleGroupStatementProperty{
//   				Name: jsii.String("name"),
//   				VendorName: jsii.String("vendorName"),
//
//   				// the properties below are optional
//   				ExcludedRules: []interface{}{
//   					&ExcludedRuleProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				ManagedRuleGroupConfigs: []interface{}{
//   					&ManagedRuleGroupConfigProperty{
//   						AwsManagedRulesAtpRuleSet: &AWSManagedRulesATPRuleSetProperty{
//   							LoginPath: jsii.String("loginPath"),
//
//   							// the properties below are optional
//   							RequestInspection: &RequestInspectionProperty{
//   								PasswordField: &FieldIdentifierProperty{
//   									Identifier: jsii.String("identifier"),
//   								},
//   								PayloadType: jsii.String("payloadType"),
//   								UsernameField: &FieldIdentifierProperty{
//   									Identifier: jsii.String("identifier"),
//   								},
//   							},
//   							ResponseInspection: &ResponseInspectionProperty{
//   								BodyContains: &ResponseInspectionBodyContainsProperty{
//   									FailureStrings: []*string{
//   										jsii.String("failureStrings"),
//   									},
//   									SuccessStrings: []*string{
//   										jsii.String("successStrings"),
//   									},
//   								},
//   								Header: &ResponseInspectionHeaderProperty{
//   									FailureValues: []*string{
//   										jsii.String("failureValues"),
//   									},
//   									Name: jsii.String("name"),
//   									SuccessValues: []*string{
//   										jsii.String("successValues"),
//   									},
//   								},
//   								Json: &ResponseInspectionJsonProperty{
//   									FailureValues: []*string{
//   										jsii.String("failureValues"),
//   									},
//   									Identifier: jsii.String("identifier"),
//   									SuccessValues: []*string{
//   										jsii.String("successValues"),
//   									},
//   								},
//   								StatusCode: &ResponseInspectionStatusCodeProperty{
//   									FailureCodes: []interface{}{
//   										jsii.Number(123),
//   									},
//   									SuccessCodes: []interface{}{
//   										jsii.Number(123),
//   									},
//   								},
//   							},
//   						},
//   						AwsManagedRulesBotControlRuleSet: &AWSManagedRulesBotControlRuleSetProperty{
//   							InspectionLevel: jsii.String("inspectionLevel"),
//   						},
//   						LoginPath: jsii.String("loginPath"),
//   						PasswordField: &FieldIdentifierProperty{
//   							Identifier: jsii.String("identifier"),
//   						},
//   						PayloadType: jsii.String("payloadType"),
//   						UsernameField: &FieldIdentifierProperty{
//   							Identifier: jsii.String("identifier"),
//   						},
//   					},
//   				},
//   				RuleActionOverrides: []interface{}{
//   					&RuleActionOverrideProperty{
//   						ActionToUse: &RuleActionProperty{
//   							Allow: &AllowActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Block: &BlockActionProperty{
//   								CustomResponse: &CustomResponseProperty{
//   									ResponseCode: jsii.Number(123),
//
//   									// the properties below are optional
//   									CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   									ResponseHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Captcha: &CaptchaActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Challenge: &ChallengeActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Count: &CountActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				ScopeDownStatement: statementProperty_,
//   				Version: jsii.String("version"),
//   			},
//   			NotStatement: &notStatementProperty{
//   				Statement: statementProperty_,
//   			},
//   			OrStatement: &orStatementProperty{
//   				Statements: []interface{}{
//   					statementProperty_,
//   				},
//   			},
//   			RateBasedStatement: rateBasedStatementProperty_,
//   			RegexMatchStatement: &RegexMatchStatementProperty{
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				RegexString: jsii.String("regexString"),
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			RegexPatternSetReferenceStatement: &RegexPatternSetReferenceStatementProperty{
//   				Arn: jsii.String("arn"),
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			RuleGroupReferenceStatement: &RuleGroupReferenceStatementProperty{
//   				Arn: jsii.String("arn"),
//
//   				// the properties below are optional
//   				ExcludedRules: []interface{}{
//   					&ExcludedRuleProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				RuleActionOverrides: []interface{}{
//   					&RuleActionOverrideProperty{
//   						ActionToUse: &RuleActionProperty{
//   							Allow: &AllowActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Block: &BlockActionProperty{
//   								CustomResponse: &CustomResponseProperty{
//   									ResponseCode: jsii.Number(123),
//
//   									// the properties below are optional
//   									CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   									ResponseHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Captcha: &CaptchaActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Challenge: &ChallengeActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   							Count: &CountActionProperty{
//   								CustomRequestHandling: &CustomRequestHandlingProperty{
//   									InsertHeaders: []interface{}{
//   										&CustomHTTPHeaderProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   							},
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   			SizeConstraintStatement: &SizeConstraintStatementProperty{
//   				ComparisonOperator: jsii.String("comparisonOperator"),
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				Size: jsii.Number(123),
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			SqliMatchStatement: &SqliMatchStatementProperty{
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				SensitivityLevel: jsii.String("sensitivityLevel"),
//   			},
//   			XssMatchStatement: &XssMatchStatementProperty{
//   				FieldToMatch: &FieldToMatchProperty{
//   					AllQueryArguments: allQueryArguments,
//   					Body: &BodyProperty{
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Cookies: &CookiesProperty{
//   						MatchPattern: &CookieMatchPatternProperty{
//   							All: all,
//   							ExcludedCookies: []*string{
//   								jsii.String("excludedCookies"),
//   							},
//   							IncludedCookies: []*string{
//   								jsii.String("includedCookies"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Headers: &HeadersProperty{
//   						MatchPattern: &HeaderMatchPatternProperty{
//   							All: all,
//   							ExcludedHeaders: []*string{
//   								jsii.String("excludedHeaders"),
//   							},
//   							IncludedHeaders: []*string{
//   								jsii.String("includedHeaders"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					JsonBody: &JsonBodyProperty{
//   						MatchPattern: &JsonMatchPatternProperty{
//   							All: all,
//   							IncludedPaths: []*string{
//   								jsii.String("includedPaths"),
//   							},
//   						},
//   						MatchScope: jsii.String("matchScope"),
//
//   						// the properties below are optional
//   						InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   						OversizeHandling: jsii.String("oversizeHandling"),
//   					},
//   					Method: method,
//   					QueryString: queryString,
//   					SingleHeader: singleHeader,
//   					SingleQueryArgument: singleQueryArgument,
//   					UriPath: uriPath,
//   				},
//   				TextTransformations: []interface{}{
//   					&TextTransformationProperty{
//   						Priority: jsii.Number(123),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	RegexMatchStatement: &RegexMatchStatementProperty{
//   		FieldToMatch: &FieldToMatchProperty{
//   			AllQueryArguments: allQueryArguments,
//   			Body: &BodyProperty{
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Cookies: &CookiesProperty{
//   				MatchPattern: &CookieMatchPatternProperty{
//   					All: all,
//   					ExcludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					IncludedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Headers: &HeadersProperty{
//   				MatchPattern: &HeaderMatchPatternProperty{
//   					All: all,
//   					ExcludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					IncludedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			JsonBody: &JsonBodyProperty{
//   				MatchPattern: &JsonMatchPatternProperty{
//   					All: all,
//   					IncludedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Method: method,
//   			QueryString: queryString,
//   			SingleHeader: singleHeader,
//   			SingleQueryArgument: singleQueryArgument,
//   			UriPath: uriPath,
//   		},
//   		RegexString: jsii.String("regexString"),
//   		TextTransformations: []interface{}{
//   			&TextTransformationProperty{
//   				Priority: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	RegexPatternSetReferenceStatement: &RegexPatternSetReferenceStatementProperty{
//   		Arn: jsii.String("arn"),
//   		FieldToMatch: &FieldToMatchProperty{
//   			AllQueryArguments: allQueryArguments,
//   			Body: &BodyProperty{
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Cookies: &CookiesProperty{
//   				MatchPattern: &CookieMatchPatternProperty{
//   					All: all,
//   					ExcludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					IncludedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Headers: &HeadersProperty{
//   				MatchPattern: &HeaderMatchPatternProperty{
//   					All: all,
//   					ExcludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					IncludedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			JsonBody: &JsonBodyProperty{
//   				MatchPattern: &JsonMatchPatternProperty{
//   					All: all,
//   					IncludedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Method: method,
//   			QueryString: queryString,
//   			SingleHeader: singleHeader,
//   			SingleQueryArgument: singleQueryArgument,
//   			UriPath: uriPath,
//   		},
//   		TextTransformations: []interface{}{
//   			&TextTransformationProperty{
//   				Priority: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	RuleGroupReferenceStatement: &RuleGroupReferenceStatementProperty{
//   		Arn: jsii.String("arn"),
//
//   		// the properties below are optional
//   		ExcludedRules: []interface{}{
//   			&ExcludedRuleProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		RuleActionOverrides: []interface{}{
//   			&RuleActionOverrideProperty{
//   				ActionToUse: &RuleActionProperty{
//   					Allow: &AllowActionProperty{
//   						CustomRequestHandling: &CustomRequestHandlingProperty{
//   							InsertHeaders: []interface{}{
//   								&CustomHTTPHeaderProperty{
//   									Name: jsii.String("name"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   					Block: &BlockActionProperty{
//   						CustomResponse: &CustomResponseProperty{
//   							ResponseCode: jsii.Number(123),
//
//   							// the properties below are optional
//   							CustomResponseBodyKey: jsii.String("customResponseBodyKey"),
//   							ResponseHeaders: []interface{}{
//   								&CustomHTTPHeaderProperty{
//   									Name: jsii.String("name"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   					Captcha: &CaptchaActionProperty{
//   						CustomRequestHandling: &CustomRequestHandlingProperty{
//   							InsertHeaders: []interface{}{
//   								&CustomHTTPHeaderProperty{
//   									Name: jsii.String("name"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   					Challenge: &ChallengeActionProperty{
//   						CustomRequestHandling: &CustomRequestHandlingProperty{
//   							InsertHeaders: []interface{}{
//   								&CustomHTTPHeaderProperty{
//   									Name: jsii.String("name"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   					Count: &CountActionProperty{
//   						CustomRequestHandling: &CustomRequestHandlingProperty{
//   							InsertHeaders: []interface{}{
//   								&CustomHTTPHeaderProperty{
//   									Name: jsii.String("name"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	SizeConstraintStatement: &SizeConstraintStatementProperty{
//   		ComparisonOperator: jsii.String("comparisonOperator"),
//   		FieldToMatch: &FieldToMatchProperty{
//   			AllQueryArguments: allQueryArguments,
//   			Body: &BodyProperty{
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Cookies: &CookiesProperty{
//   				MatchPattern: &CookieMatchPatternProperty{
//   					All: all,
//   					ExcludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					IncludedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Headers: &HeadersProperty{
//   				MatchPattern: &HeaderMatchPatternProperty{
//   					All: all,
//   					ExcludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					IncludedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			JsonBody: &JsonBodyProperty{
//   				MatchPattern: &JsonMatchPatternProperty{
//   					All: all,
//   					IncludedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Method: method,
//   			QueryString: queryString,
//   			SingleHeader: singleHeader,
//   			SingleQueryArgument: singleQueryArgument,
//   			UriPath: uriPath,
//   		},
//   		Size: jsii.Number(123),
//   		TextTransformations: []interface{}{
//   			&TextTransformationProperty{
//   				Priority: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	SqliMatchStatement: &SqliMatchStatementProperty{
//   		FieldToMatch: &FieldToMatchProperty{
//   			AllQueryArguments: allQueryArguments,
//   			Body: &BodyProperty{
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Cookies: &CookiesProperty{
//   				MatchPattern: &CookieMatchPatternProperty{
//   					All: all,
//   					ExcludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					IncludedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Headers: &HeadersProperty{
//   				MatchPattern: &HeaderMatchPatternProperty{
//   					All: all,
//   					ExcludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					IncludedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			JsonBody: &JsonBodyProperty{
//   				MatchPattern: &JsonMatchPatternProperty{
//   					All: all,
//   					IncludedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Method: method,
//   			QueryString: queryString,
//   			SingleHeader: singleHeader,
//   			SingleQueryArgument: singleQueryArgument,
//   			UriPath: uriPath,
//   		},
//   		TextTransformations: []interface{}{
//   			&TextTransformationProperty{
//   				Priority: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		SensitivityLevel: jsii.String("sensitivityLevel"),
//   	},
//   	XssMatchStatement: &XssMatchStatementProperty{
//   		FieldToMatch: &FieldToMatchProperty{
//   			AllQueryArguments: allQueryArguments,
//   			Body: &BodyProperty{
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Cookies: &CookiesProperty{
//   				MatchPattern: &CookieMatchPatternProperty{
//   					All: all,
//   					ExcludedCookies: []*string{
//   						jsii.String("excludedCookies"),
//   					},
//   					IncludedCookies: []*string{
//   						jsii.String("includedCookies"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Headers: &HeadersProperty{
//   				MatchPattern: &HeaderMatchPatternProperty{
//   					All: all,
//   					ExcludedHeaders: []*string{
//   						jsii.String("excludedHeaders"),
//   					},
//   					IncludedHeaders: []*string{
//   						jsii.String("includedHeaders"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			JsonBody: &JsonBodyProperty{
//   				MatchPattern: &JsonMatchPatternProperty{
//   					All: all,
//   					IncludedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				MatchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   				OversizeHandling: jsii.String("oversizeHandling"),
//   			},
//   			Method: method,
//   			QueryString: queryString,
//   			SingleHeader: singleHeader,
//   			SingleQueryArgument: singleQueryArgument,
//   			UriPath: uriPath,
//   		},
//   		TextTransformations: []interface{}{
//   			&TextTransformationProperty{
//   				Priority: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   }
//
type CfnWebACL_StatementProperty struct {
	// A logical rule statement used to combine other rule statements with AND logic.
	//
	// You provide more than one `Statement` within the `AndStatement` .
	AndStatement interface{} `field:"optional" json:"andStatement" yaml:"andStatement"`
	// A rule statement that defines a string match search for AWS WAF to apply to web requests.
	//
	// The byte match statement provides the bytes to search for, the location in requests that you want AWS WAF to search, and other settings. The bytes to search for are typically a string that corresponds with ASCII characters. In the AWS WAF console and the developer guide, this is called a string match statement.
	ByteMatchStatement interface{} `field:"optional" json:"byteMatchStatement" yaml:"byteMatchStatement"`
	// A rule statement that labels web requests by country and region and that matches against web requests based on country code.
	//
	// A geo match rule labels every request that it inspects regardless of whether it finds a match.
	//
	// - To manage requests only by country, you can use this statement by itself and specify the countries that you want to match against in the `CountryCodes` array.
	// - Otherwise, configure your geo match rule with Count action so that it only labels requests. Then, add one or more label match rules to run after the geo match rule and configure them to match against the geographic labels and handle the requests as needed.
	//
	// AWS WAF labels requests using the alpha-2 country and region codes from the International Organization for Standardization (ISO) 3166 standard. AWS WAF determines the codes using either the IP address in the web request origin or, if you specify it, the address in the geo match `ForwardedIPConfig` .
	//
	// If you use the web request origin, the label formats are `awswaf:clientip:geo:region:<ISO country code>-<ISO region code>` and `awswaf:clientip:geo:country:<ISO country code>` .
	//
	// If you use a forwarded IP address, the label formats are `awswaf:forwardedip:geo:region:<ISO country code>-<ISO region code>` and `awswaf:forwardedip:geo:country:<ISO country code>` .
	//
	// For additional details, see [Geographic match rule statement](https://docs.aws.amazon.com/waf/latest/developerguide/waf-rule-statement-type-geo-match.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	GeoMatchStatement interface{} `field:"optional" json:"geoMatchStatement" yaml:"geoMatchStatement"`
	// A rule statement used to detect web requests coming from particular IP addresses or address ranges.
	//
	// To use this, create an `IPSet` that specifies the addresses you want to detect, then use the ARN of that set in this statement.
	//
	// Each IP set rule statement references an IP set. You create and maintain the set independent of your rules. This allows you to use the single set in multiple rules. When you update the referenced set, AWS WAF automatically updates all rules that reference it.
	IpSetReferenceStatement interface{} `field:"optional" json:"ipSetReferenceStatement" yaml:"ipSetReferenceStatement"`
	// A rule statement to match against labels that have been added to the web request by rules that have already run in the web ACL.
	//
	// The label match statement provides the label or namespace string to search for. The label string can represent a part or all of the fully qualified label name that had been added to the web request. Fully qualified labels have a prefix, optional namespaces, and label name. The prefix identifies the rule group or web ACL context of the rule that added the label. If you do not provide the fully qualified name in your label match string, AWS WAF performs the search for labels that were added in the same context as the label match statement.
	LabelMatchStatement interface{} `field:"optional" json:"labelMatchStatement" yaml:"labelMatchStatement"`
	// A rule statement used to run the rules that are defined in a managed rule group.
	//
	// To use this, provide the vendor name and the name of the rule group in this statement.
	//
	// You cannot nest a `ManagedRuleGroupStatement` , for example for use inside a `NotStatement` or `OrStatement` . It can only be referenced as a top-level statement within a rule.
	ManagedRuleGroupStatement interface{} `field:"optional" json:"managedRuleGroupStatement" yaml:"managedRuleGroupStatement"`
	// A logical rule statement used to negate the results of another rule statement.
	//
	// You provide one `Statement` within the `NotStatement` .
	NotStatement interface{} `field:"optional" json:"notStatement" yaml:"notStatement"`
	// A logical rule statement used to combine other rule statements with OR logic.
	//
	// You provide more than one `Statement` within the `OrStatement` .
	OrStatement interface{} `field:"optional" json:"orStatement" yaml:"orStatement"`
	// A rate-based rule tracks the rate of requests for each originating IP address, and triggers the rule action when the rate exceeds a limit that you specify on the number of requests in any 5-minute time span.
	//
	// You can use this to put a temporary block on requests from an IP address that is sending excessive requests.
	//
	// AWS WAF tracks and manages web requests separately for each instance of a rate-based rule that you use. For example, if you provide the same rate-based rule settings in two web ACLs, each of the two rule statements represents a separate instance of the rate-based rule and gets its own tracking and management by AWS WAF . If you define a rate-based rule inside a rule group, and then use that rule group in multiple places, each use creates a separate instance of the rate-based rule that gets its own tracking and management by AWS WAF .
	//
	// When the rule action triggers, AWS WAF blocks additional requests from the IP address until the request rate falls below the limit.
	//
	// You can optionally nest another statement inside the rate-based statement, to narrow the scope of the rule so that it only counts requests that match the nested statement. For example, based on recent requests that you have seen from an attacker, you might create a rate-based rule with a nested AND rule statement that contains the following nested statements:
	//
	// - An IP match statement with an IP set that specifies the address 192.0.2.44.
	// - A string match statement that searches in the User-Agent header for the string BadBot.
	//
	// In this rate-based rule, you also define a rate limit. For this example, the rate limit is 1,000. Requests that meet the criteria of both of the nested statements are counted. If the count exceeds 1,000 requests per five minutes, the rule action triggers. Requests that do not meet the criteria of both of the nested statements are not counted towards the rate limit and are not affected by this rule.
	//
	// You cannot nest a `RateBasedStatement` inside another statement, for example inside a `NotStatement` or `OrStatement` . You can define a `RateBasedStatement` inside a web ACL and inside a rule group.
	RateBasedStatement interface{} `field:"optional" json:"rateBasedStatement" yaml:"rateBasedStatement"`
	// A rule statement used to search web request components for a match against a single regular expression.
	RegexMatchStatement interface{} `field:"optional" json:"regexMatchStatement" yaml:"regexMatchStatement"`
	// A rule statement used to search web request components for matches with regular expressions.
	//
	// To use this, create a `RegexPatternSet` that specifies the expressions that you want to detect, then use the ARN of that set in this statement. A web request matches the pattern set rule statement if the request component matches any of the patterns in the set.
	//
	// Each regex pattern set rule statement references a regex pattern set. You create and maintain the set independent of your rules. This allows you to use the single set in multiple rules. When you update the referenced set, AWS WAF automatically updates all rules that reference it.
	RegexPatternSetReferenceStatement interface{} `field:"optional" json:"regexPatternSetReferenceStatement" yaml:"regexPatternSetReferenceStatement"`
	// A rule statement used to run the rules that are defined in a `RuleGroup` .
	//
	// To use this, create a rule group with your rules, then provide the ARN of the rule group in this statement.
	//
	// You cannot nest a `RuleGroupReferenceStatement` , for example for use inside a `NotStatement` or `OrStatement` . You can only use a rule group reference statement at the top level inside a web ACL.
	RuleGroupReferenceStatement interface{} `field:"optional" json:"ruleGroupReferenceStatement" yaml:"ruleGroupReferenceStatement"`
	// A rule statement that compares a number of bytes against the size of a request component, using a comparison operator, such as greater than (>) or less than (<).
	//
	// For example, you can use a size constraint statement to look for query strings that are longer than 100 bytes.
	//
	// If you configure AWS WAF to inspect the request body, AWS WAF inspects only the first 8192 bytes (8 KB). If the request body for your web requests never exceeds 8192 bytes, you could use a size constraint statement to block requests that have a request body greater than 8192 bytes.
	//
	// If you choose URI for the value of Part of the request to filter on, the slash (/) in the URI counts as one character. For example, the URI `/logo.jpg` is nine characters long.
	SizeConstraintStatement interface{} `field:"optional" json:"sizeConstraintStatement" yaml:"sizeConstraintStatement"`
	// A rule statement that inspects for malicious SQL code.
	//
	// Attackers insert malicious SQL code into web requests to do things like modify your database or extract data from it.
	SqliMatchStatement interface{} `field:"optional" json:"sqliMatchStatement" yaml:"sqliMatchStatement"`
	// A rule statement that inspects for cross-site scripting (XSS) attacks.
	//
	// In XSS attacks, the attacker uses vulnerabilities in a benign website as a vehicle to inject malicious client-site scripts into other legitimate web browsers.
	XssMatchStatement interface{} `field:"optional" json:"xssMatchStatement" yaml:"xssMatchStatement"`
}
