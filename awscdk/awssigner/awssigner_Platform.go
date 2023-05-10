package awssigner

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Platforms that are allowed with signing config.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   signingProfile := signer.NewSigningProfile(this, jsii.String("SigningProfile"), &SigningProfileProps{
//   	Platform: signer.Platform_AWS_LAMBDA_SHA384_ECDSA(),
//   })
//
//   codeSigningConfig := lambda.NewCodeSigningConfig(this, jsii.String("CodeSigningConfig"), &CodeSigningConfigProps{
//   	SigningProfiles: []iSigningProfile{
//   		signingProfile,
//   	},
//   })
//
//   lambda.NewFunction(this, jsii.String("Function"), &FunctionProps{
//   	CodeSigningConfig: CodeSigningConfig,
//   	Runtime: lambda.Runtime_NODEJS_16_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   })
//
// See: https://docs.aws.amazon.com/signer/latest/developerguide/gs-platform.html
//
// Experimental.
type Platform interface {
	// The id of signing platform.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-signer-signingprofile.html#cfn-signer-signingprofile-platformid
	//
	// Experimental.
	PlatformId() *string
}

// The jsii proxy struct for Platform
type jsiiProxy_Platform struct {
	_ byte // padding
}

func (j *jsiiProxy_Platform) PlatformId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformId",
		&returns,
	)
	return returns
}


func Platform_AMAZON_FREE_RTOS_DEFAULT() Platform {
	_init_.Initialize()
	var returns Platform
	_jsii_.StaticGet(
		"monocdk.aws_signer.Platform",
		"AMAZON_FREE_RTOS_DEFAULT",
		&returns,
	)
	return returns
}

func Platform_AMAZON_FREE_RTOS_TI_CC3220SF() Platform {
	_init_.Initialize()
	var returns Platform
	_jsii_.StaticGet(
		"monocdk.aws_signer.Platform",
		"AMAZON_FREE_RTOS_TI_CC3220SF",
		&returns,
	)
	return returns
}

func Platform_AWS_IOT_DEVICE_MANAGEMENT_SHA256_ECDSA() Platform {
	_init_.Initialize()
	var returns Platform
	_jsii_.StaticGet(
		"monocdk.aws_signer.Platform",
		"AWS_IOT_DEVICE_MANAGEMENT_SHA256_ECDSA",
		&returns,
	)
	return returns
}

func Platform_AWS_LAMBDA_SHA384_ECDSA() Platform {
	_init_.Initialize()
	var returns Platform
	_jsii_.StaticGet(
		"monocdk.aws_signer.Platform",
		"AWS_LAMBDA_SHA384_ECDSA",
		&returns,
	)
	return returns
}

