package awssupportapp


// Properties for defining a `CfnAccountAlias`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccountAliasProps := &CfnAccountAliasProps{
//   	AccountAlias: jsii.String("accountAlias"),
//   }
//
type CfnAccountAliasProps struct {
	// An alias or short name for an AWS account .
	AccountAlias *string `field:"required" json:"accountAlias" yaml:"accountAlias"`
}
