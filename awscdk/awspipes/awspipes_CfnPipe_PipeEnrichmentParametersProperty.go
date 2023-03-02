package awspipes


// The parameters required to set up enrichment on your pipe.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeEnrichmentParametersProperty := &pipeEnrichmentParametersProperty{
//   	httpParameters: &pipeEnrichmentHttpParametersProperty{
//   		headerParameters: map[string]*string{
//   			"headerParametersKey": jsii.String("headerParameters"),
//   		},
//   		pathParameterValues: []*string{
//   			jsii.String("pathParameterValues"),
//   		},
//   		queryStringParameters: map[string]*string{
//   			"queryStringParametersKey": jsii.String("queryStringParameters"),
//   		},
//   	},
//   	inputTemplate: jsii.String("inputTemplate"),
//   }
//
type CfnPipe_PipeEnrichmentParametersProperty struct {
	// Contains the HTTP parameters to use when the target is a API Gateway REST endpoint or EventBridge ApiDestination.
	//
	// If you specify an API Gateway REST API or EventBridge ApiDestination as a target, you can use this parameter to specify headers, path parameters, and query string keys/values as part of your target invoking request. If you're using ApiDestinations, the corresponding Connection can also have these values configured. In case of any conflicting keys, values from the Connection take precedence.
	HttpParameters interface{} `field:"optional" json:"httpParameters" yaml:"httpParameters"`
	// Valid JSON text passed to the enrichment.
	//
	// In this case, nothing from the event itself is passed to the enrichment. For more information, see [The JavaScript Object Notation (JSON) Data Interchange Format](https://docs.aws.amazon.com/http://www.rfc-editor.org/rfc/rfc7159.txt) .
	InputTemplate *string `field:"optional" json:"inputTemplate" yaml:"inputTemplate"`
}

