// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Representation of validation results.
//
// Models a tree of validation errors so that we have as much information as possible
// about the failure that occurred.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var validationResults validationResults
//
//   validationResult := monocdk.NewValidationResult(jsii.String("errorMessage"), validationResults)
//
// Experimental.
type ValidationResult interface {
	// Experimental.
	ErrorMessage() *string
	// Experimental.
	IsSuccess() *bool
	// Experimental.
	Results() ValidationResults
	// Turn a failed validation into an exception.
	// Experimental.
	AssertSuccess()
	// Return a string rendering of the tree of validation failures.
	// Experimental.
	ErrorTree() *string
	// Wrap this result with an error message, if it concerns an error.
	// Experimental.
	Prefix(message *string) ValidationResult
}

// The jsii proxy struct for ValidationResult
type jsiiProxy_ValidationResult struct {
	_ byte // padding
}

func (j *jsiiProxy_ValidationResult) ErrorMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidationResult) IsSuccess() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ValidationResult) Results() ValidationResults {
	var returns ValidationResults
	_jsii_.Get(
		j,
		"results",
		&returns,
	)
	return returns
}


// Experimental.
func NewValidationResult(errorMessage *string, results ValidationResults) ValidationResult {
	_init_.Initialize()

	j := jsiiProxy_ValidationResult{}

	_jsii_.Create(
		"monocdk.ValidationResult",
		[]interface{}{errorMessage, results},
		&j,
	)

	return &j
}

// Experimental.
func NewValidationResult_Override(v ValidationResult, errorMessage *string, results ValidationResults) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.ValidationResult",
		[]interface{}{errorMessage, results},
		v,
	)
}

func (v *jsiiProxy_ValidationResult) AssertSuccess() {
	_jsii_.InvokeVoid(
		v,
		"assertSuccess",
		nil, // no parameters
	)
}

func (v *jsiiProxy_ValidationResult) ErrorTree() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"errorTree",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ValidationResult) Prefix(message *string) ValidationResult {
	if err := v.validatePrefixParameters(message); err != nil {
		panic(err)
	}
	var returns ValidationResult

	_jsii_.Invoke(
		v,
		"prefix",
		[]interface{}{message},
		&returns,
	)

	return returns
}

