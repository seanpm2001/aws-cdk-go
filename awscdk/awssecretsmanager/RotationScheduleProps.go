package awssecretsmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Construction properties for a RotationSchedule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var function_ function
//   var hostedRotation hostedRotation
//   var secret secret
//
//   rotationScheduleProps := &RotationScheduleProps{
//   	Secret: secret,
//
//   	// the properties below are optional
//   	AutomaticallyAfter: cdk.Duration_Minutes(jsii.Number(30)),
//   	HostedRotation: hostedRotation,
//   	RotationLambda: function_,
//   }
//
type RotationScheduleProps struct {
	// Specifies the number of days after the previous rotation before Secrets Manager triggers the next automatic rotation.
	//
	// A value of zero will disable automatic rotation - `Duration.days(0)`.
	AutomaticallyAfter awscdk.Duration `field:"optional" json:"automaticallyAfter" yaml:"automaticallyAfter"`
	// Hosted rotation.
	HostedRotation HostedRotation `field:"optional" json:"hostedRotation" yaml:"hostedRotation"`
	// A Lambda function that can rotate the secret.
	RotationLambda awslambda.IFunction `field:"optional" json:"rotationLambda" yaml:"rotationLambda"`
	// The secret to rotate.
	//
	// If hosted rotation is used, this must be a JSON string with the following format:
	//
	// ```
	// {
	//   "engine": <required: database engine>,
	//   "host": <required: instance host name>,
	//   "username": <required: username>,
	//   "password": <required: password>,
	//   "dbname": <optional: database name>,
	//   "port": <optional: if not specified, default port will be used>,
	//   "masterarn": <required for multi user rotation: the arn of the master secret which will be used to create users/change passwords>
	// }
	// ```
	//
	// This is typically the case for a secret referenced from an `AWS::SecretsManager::SecretTargetAttachment`
	// or an `ISecret` returned by the `attach()` method of `Secret`.
	Secret ISecret `field:"required" json:"secret" yaml:"secret"`
}
