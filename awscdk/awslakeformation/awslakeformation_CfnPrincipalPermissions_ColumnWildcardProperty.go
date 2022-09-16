package awslakeformation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnWildcardProperty := &columnWildcardProperty{
//   	excludedColumnNames: []*string{
//   		jsii.String("excludedColumnNames"),
//   	},
//   }
//
type CfnPrincipalPermissions_ColumnWildcardProperty struct {
	// `CfnPrincipalPermissions.ColumnWildcardProperty.ExcludedColumnNames`.
	ExcludedColumnNames *[]*string `field:"optional" json:"excludedColumnNames" yaml:"excludedColumnNames"`
}
