package awsopensearchservice

import (
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// Whether the domain should encrypt data at rest, and if so, the AWS Key Management Service (KMS) key to use.
//
// Can only be used to create a new domain,
// not update an existing one. Requires Elasticsearch version 5.1 or later or OpenSearch version 1.0 or later.
//
// Example:
//   domain := opensearch.NewDomain(this, jsii.String("Domain"), &domainProps{
//   	version: opensearch.engineVersion_OPENSEARCH_1_0(),
//   	enforceHttps: jsii.Boolean(true),
//   	nodeToNodeEncryption: jsii.Boolean(true),
//   	encryptionAtRest: &encryptionAtRestOptions{
//   		enabled: jsii.Boolean(true),
//   	},
//   	fineGrainedAccessControl: &advancedSecurityOptions{
//   		masterUserName: jsii.String("master-user"),
//   	},
//   	logging: &loggingOptions{
//   		auditLogEnabled: jsii.Boolean(true),
//   		slowSearchLogEnabled: jsii.Boolean(true),
//   		appLogEnabled: jsii.Boolean(true),
//   		slowIndexLogEnabled: jsii.Boolean(true),
//   	},
//   })
//
// Experimental.
type EncryptionAtRestOptions struct {
	// Specify true to enable encryption at rest.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Supply if using KMS key for encryption at rest.
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

