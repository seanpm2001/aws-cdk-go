package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oidcMemberDefinitionProperty := &oidcMemberDefinitionProperty{
//   	oidcGroups: []*string{
//   		jsii.String("oidcGroups"),
//   	},
//   }
//
type CfnWorkteam_OidcMemberDefinitionProperty struct {
	// `CfnWorkteam.OidcMemberDefinitionProperty.OidcGroups`.
	OidcGroups *[]*string `field:"required" json:"oidcGroups" yaml:"oidcGroups"`
}
