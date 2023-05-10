package awseventstargets


// Rule target input for an AwsApi target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   awsApiInput := &AwsApiInput{
//   	Action: jsii.String("action"),
//   	Service: jsii.String("service"),
//
//   	// the properties below are optional
//   	ApiVersion: jsii.String("apiVersion"),
//   	CatchErrorPattern: jsii.String("catchErrorPattern"),
//   	Parameters: parameters,
//   }
//
// Experimental.
type AwsApiInput struct {
	// The service action to call.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	// Experimental.
	Action *string `field:"required" json:"action" yaml:"action"`
	// The service to call.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	// Experimental.
	Service *string `field:"required" json:"service" yaml:"service"`
	// API version to use for the service.
	// See: https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/locking-api-versions.html
	//
	// Experimental.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// The regex pattern to use to catch API errors.
	//
	// The `code` property of the
	// `Error` object will be tested against this pattern. If there is a match an
	// error will not be thrown.
	// Experimental.
	CatchErrorPattern *string `field:"optional" json:"catchErrorPattern" yaml:"catchErrorPattern"`
	// The parameters for the service action.
	// See: https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html
	//
	// Experimental.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

