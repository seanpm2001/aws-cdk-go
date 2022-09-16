package awsssm


// Properties needed to create a String SSM parameter.
//
// Example:
//   ssm.NewStringParameter(this, jsii.String("Parameter"), &stringParameterProps{
//   	allowedPattern: jsii.String(".*"),
//   	description: jsii.String("The value Foo"),
//   	parameterName: jsii.String("FooParameter"),
//   	stringValue: jsii.String("Foo"),
//   	tier: ssm.parameterTier_ADVANCED,
//   })
//
// Experimental.
type StringParameterProps struct {
	// A regular expression used to validate the parameter value.
	//
	// For example, for String types with values restricted to
	// numbers, you can specify the following: ``^\d+$``.
	// Experimental.
	AllowedPattern *string `field:"optional" json:"allowedPattern" yaml:"allowedPattern"`
	// Information about the parameter that you want to add to the system.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the parameter.
	// Experimental.
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
	// Experimental.
	SimpleName *bool `field:"optional" json:"simpleName" yaml:"simpleName"`
	// The tier of the string parameter.
	// Experimental.
	Tier ParameterTier `field:"optional" json:"tier" yaml:"tier"`
	// The value of the parameter.
	//
	// It may not reference another parameter and ``{{}}`` cannot be used in the value.
	// Experimental.
	StringValue *string `field:"required" json:"stringValue" yaml:"stringValue"`
	// The data type of the parameter, such as `text` or `aws:ec2:image`.
	// Experimental.
	DataType ParameterDataType `field:"optional" json:"dataType" yaml:"dataType"`
	// The type of the string parameter.
	// Experimental.
	Type ParameterType `field:"optional" json:"type" yaml:"type"`
}
