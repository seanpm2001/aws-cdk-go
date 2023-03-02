package awsrolesanywhere

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCRL`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCRLProps := &cfnCRLProps{
//   	crlData: jsii.String("crlData"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	trustAnchorArn: jsii.String("trustAnchorArn"),
//   }
//
type CfnCRLProps struct {
	// x509 v3 Certificate Revocation List to revoke auth for corresponding certificates presented in CreateSession operations.
	CrlData *string `field:"required" json:"crlData" yaml:"crlData"`
	// The customer specified name of the resource.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The enabled status of the resource.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A list of Tags.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ARN of the TrustAnchor the certificate revocation list (CRL) will provide revocation for.
	TrustAnchorArn *string `field:"optional" json:"trustAnchorArn" yaml:"trustAnchorArn"`
}

