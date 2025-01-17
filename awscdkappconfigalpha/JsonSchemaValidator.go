package awscdkappconfigalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkappconfigalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Defines a JSON Schema validator.
//
// Example:
//   var application application
//   var fn function
//
//
//   appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfiguration"), &HostedConfigurationProps{
//   	Application: Application,
//   	Content: appconfig.ConfigurationContent_FromInlineText(jsii.String("This is my configuration content.")),
//   	Validators: []iValidator{
//   		appconfig.JsonSchemaValidator_FromFile(jsii.String("schema.json")),
//   		appconfig.LambdaValidator_FromFunction(fn),
//   	},
//   })
//
// Experimental.
type JsonSchemaValidator interface {
	IValidator
	// The content of the validator.
	// Experimental.
	Content() *string
	// The type of validator.
	// Experimental.
	Type() ValidatorType
}

// The jsii proxy struct for JsonSchemaValidator
type jsiiProxy_JsonSchemaValidator struct {
	jsiiProxy_IValidator
}

func (j *jsiiProxy_JsonSchemaValidator) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JsonSchemaValidator) Type() ValidatorType {
	var returns ValidatorType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Experimental.
func NewJsonSchemaValidator_Override(j JsonSchemaValidator) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-appconfig-alpha.JsonSchemaValidator",
		nil, // no parameters
		j,
	)
}

// Defines a JSON Schema validator from a file.
// Experimental.
func JsonSchemaValidator_FromFile(path *string) JsonSchemaValidator {
	_init_.Initialize()

	if err := validateJsonSchemaValidator_FromFileParameters(path); err != nil {
		panic(err)
	}
	var returns JsonSchemaValidator

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.JsonSchemaValidator",
		"fromFile",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Defines a JSON Schema validator from inline code.
// Experimental.
func JsonSchemaValidator_FromInline(code *string) JsonSchemaValidator {
	_init_.Initialize()

	if err := validateJsonSchemaValidator_FromInlineParameters(code); err != nil {
		panic(err)
	}
	var returns JsonSchemaValidator

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-appconfig-alpha.JsonSchemaValidator",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

