package awsssm


// Properties needed to create a StringList SSM Parameter.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   // Create a new SSM Parameter holding a String
//   param := ssm.NewStringParameter(stack, jsii.String("StringParameter"), &StringParameterProps{
//   	// description: 'Some user-friendly description',
//   	// name: 'ParameterName',
//   	StringValue: jsii.String("Initial parameter value"),
//   })
//
//   // Grant read access to some Role
//   param.grantRead(role)
//
//   // Create a new SSM Parameter holding a StringList
//   listParameter := ssm.NewStringListParameter(stack, jsii.String("StringListParameter"), &StringListParameterProps{
//   	// description: 'Some user-friendly description',
//   	// name: 'ParameterName',
//   	StringListValue: []*string{
//   		jsii.String("Initial parameter value A"),
//   		jsii.String("Initial parameter value B"),
//   	},
//   })
//
type StringListParameterProps struct {
	// A regular expression used to validate the parameter value.
	//
	// For example, for String types with values restricted to
	// numbers, you can specify the following: ``^\d+$``.
	AllowedPattern *string `field:"optional" json:"allowedPattern" yaml:"allowedPattern"`
	// Information about the parameter that you want to add to the system.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the parameter.
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// Indicates of the parameter name is a simple name (i.e. does not include "/" separators).
	//
	// This is only required only if `parameterName` is a token, which means we
	// are unable to detect if the name is simple or "path-like" for the purpose
	// of rendering SSM parameter ARNs.
	//
	// If `parameterName` is not specified, `simpleName` must be `true` (or
	// undefined) since the name generated by AWS CloudFormation is always a
	// simple name.
	SimpleName *bool `field:"optional" json:"simpleName" yaml:"simpleName"`
	// The tier of the string parameter.
	Tier ParameterTier `field:"optional" json:"tier" yaml:"tier"`
	// The values of the parameter.
	//
	// It may not reference another parameter and ``{{}}`` cannot be used in the value.
	StringListValue *[]*string `field:"required" json:"stringListValue" yaml:"stringListValue"`
}

