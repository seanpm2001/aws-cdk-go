package awslex


// Contains information about code hooks that Amazon Lex calls during a conversation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeHookSpecificationProperty := &CodeHookSpecificationProperty{
//   	LambdaCodeHook: &LambdaCodeHookProperty{
//   		CodeHookInterfaceVersion: jsii.String("codeHookInterfaceVersion"),
//   		LambdaArn: jsii.String("lambdaArn"),
//   	},
//   }
//
type CfnBot_CodeHookSpecificationProperty struct {
	// Specifies a Lambda function that verifies requests to a bot or fulfills the user's request to a bot.
	LambdaCodeHook interface{} `field:"required" json:"lambdaCodeHook" yaml:"lambdaCodeHook"`
}
