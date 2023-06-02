package awscdksagemakeralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Construction properties for a SageMaker EndpointConfig.
//
// Example:
//   import sagemaker "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//
//   var modelA model
//   var modelB model
//
//
//   endpointConfig := sagemaker.NewEndpointConfig(this, jsii.String("EndpointConfig"), &EndpointConfigProps{
//   	InstanceProductionVariants: []instanceProductionVariantProps{
//   		&instanceProductionVariantProps{
//   			Model: modelA,
//   			VariantName: jsii.String("modelA"),
//   			InitialVariantWeight: jsii.Number(2),
//   		},
//   		&instanceProductionVariantProps{
//   			Model: modelB,
//   			VariantName: jsii.String("variantB"),
//   			InitialVariantWeight: jsii.Number(1),
//   		},
//   	},
//   })
//
// Experimental.
type EndpointConfigProps struct {
	// Optional KMS encryption key associated with this stream.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Name of the endpoint configuration.
	// Experimental.
	EndpointConfigName *string `field:"optional" json:"endpointConfigName" yaml:"endpointConfigName"`
	// A list of instance production variants.
	//
	// You can always add more variants later by calling
	// `EndpointConfig#addInstanceProductionVariant`.
	// Experimental.
	InstanceProductionVariants *[]*InstanceProductionVariantProps `field:"optional" json:"instanceProductionVariants" yaml:"instanceProductionVariants"`
}

