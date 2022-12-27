package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// The collection of versions of the ADOT Lambda Layer for Java SDK.
//
// Example:
//   var fn function
//
//   layerArn := lambda.adotLambdaLayerJavaSdkVersion_V1_19_0().layerArn(fn.stack, fn.architecture)
//
type AdotLambdaLayerJavaSdkVersion interface {
	LayerVersion() *string
	// The ARN of the Lambda layer.
	LayerArn(scope constructs.IConstruct, architecture Architecture) *string
}

// The jsii proxy struct for AdotLambdaLayerJavaSdkVersion
type jsiiProxy_AdotLambdaLayerJavaSdkVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_AdotLambdaLayerJavaSdkVersion) LayerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerVersion",
		&returns,
	)
	return returns
}


func AdotLambdaLayerJavaSdkVersion_LATEST() AdotLambdaLayerJavaSdkVersion {
	_init_.Initialize()
	var returns AdotLambdaLayerJavaSdkVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AdotLambdaLayerJavaSdkVersion",
		"LATEST",
		&returns,
	)
	return returns
}

func AdotLambdaLayerJavaSdkVersion_V1_19_0() AdotLambdaLayerJavaSdkVersion {
	_init_.Initialize()
	var returns AdotLambdaLayerJavaSdkVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AdotLambdaLayerJavaSdkVersion",
		"V1_19_0",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AdotLambdaLayerJavaSdkVersion) LayerArn(scope constructs.IConstruct, architecture Architecture) *string {
	if err := a.validateLayerArnParameters(scope, architecture); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"layerArn",
		[]interface{}{scope, architecture},
		&returns,
	)

	return returns
}

