// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
)

// properties used for creating the DomainName.
//
// Example:
//   import acm "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//   var handler function
//
//
//   certArn := "arn:aws:acm:us-east-1:111111111111:certificate"
//   domainName := "example.com"
//
//   dn := apigwv2.NewDomainName(this, jsii.String("DN"), &domainNameProps{
//   	domainName: jsii.String(domainName),
//   	certificate: acm.certificate.fromCertificateArn(this, jsii.String("cert"), certArn),
//   })
//   api := apigwv2.NewHttpApi(this, jsii.String("HttpProxyProdApi"), &httpApiProps{
//   	defaultIntegration: awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(jsii.String("DefaultIntegration"), handler),
//   	// https://${dn.domainName}/foo goes to prodApi $default stage
//   	defaultDomainMapping: &domainMappingOptions{
//   		domainName: dn,
//   		mappingKey: jsii.String("foo"),
//   	},
//   })
//
// Experimental.
type DomainNameProps struct {
	// The ACM certificate for this domain name.
	//
	// Certificate can be both ACM issued or imported.
	// Experimental.
	Certificate awscertificatemanager.ICertificate `field:"required" json:"certificate" yaml:"certificate"`
	// The user-friendly name of the certificate that will be used by the endpoint for this domain name.
	// Experimental.
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// The type of endpoint for this DomainName.
	// Experimental.
	EndpointType EndpointType `field:"optional" json:"endpointType" yaml:"endpointType"`
	// A public certificate issued by ACM to validate that you own a custom domain.
	//
	// This parameter is required
	// only when you configure mutual TLS authentication and you specify an ACM imported or private CA certificate
	// for `certificate`. The ownership certificate validates that you have permissions to use the domain name.
	// Experimental.
	OwnershipCertificate awscertificatemanager.ICertificate `field:"optional" json:"ownershipCertificate" yaml:"ownershipCertificate"`
	// The Transport Layer Security (TLS) version + cipher suite for this domain name.
	// Experimental.
	SecurityPolicy SecurityPolicy `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
	// The custom domain name.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The mutual TLS authentication configuration for a custom domain name.
	// Experimental.
	Mtls *MTLSConfig `field:"optional" json:"mtls" yaml:"mtls"`
}
