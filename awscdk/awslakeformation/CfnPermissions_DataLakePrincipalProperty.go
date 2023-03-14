package awslakeformation


// The Lake Formation principal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataLakePrincipalProperty := &DataLakePrincipalProperty{
//   	DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   }
//
type CfnPermissions_DataLakePrincipalProperty struct {
	// An identifier for the Lake Formation principal.
	DataLakePrincipalIdentifier *string `field:"optional" json:"dataLakePrincipalIdentifier" yaml:"dataLakePrincipalIdentifier"`
}

