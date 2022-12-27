package awslakeformation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   databaseResourceProperty := &databaseResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	name: jsii.String("name"),
//   }
//
type CfnTagAssociation_DatabaseResourceProperty struct {
	// `CfnTagAssociation.DatabaseResourceProperty.CatalogId`.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// `CfnTagAssociation.DatabaseResourceProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
}

