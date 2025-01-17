package awscdkapigatewayv2alpha


// Supported HTTP methods.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//   var booksDefaultFn function
//
//
//   getBooksIntegration := awscdkapigatewayv2integrationsalpha.NewHttpUrlIntegration(jsii.String("GetBooksIntegration"), jsii.String("https://get-books-proxy.example.com"))
//   booksDefaultIntegration := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(jsii.String("BooksIntegration"), booksDefaultFn)
//
//   httpApi := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/books"),
//   	Methods: []httpMethod{
//   		apigwv2.*httpMethod_GET,
//   	},
//   	Integration: getBooksIntegration,
//   })
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/books"),
//   	Methods: []*httpMethod{
//   		apigwv2.*httpMethod_ANY,
//   	},
//   	Integration: booksDefaultIntegration,
//   })
//
// Experimental.
type HttpMethod string

const (
	// HTTP ANY.
	// Experimental.
	HttpMethod_ANY HttpMethod = "ANY"
	// HTTP DELETE.
	// Experimental.
	HttpMethod_DELETE HttpMethod = "DELETE"
	// HTTP GET.
	// Experimental.
	HttpMethod_GET HttpMethod = "GET"
	// HTTP HEAD.
	// Experimental.
	HttpMethod_HEAD HttpMethod = "HEAD"
	// HTTP OPTIONS.
	// Experimental.
	HttpMethod_OPTIONS HttpMethod = "OPTIONS"
	// HTTP PATCH.
	// Experimental.
	HttpMethod_PATCH HttpMethod = "PATCH"
	// HTTP POST.
	// Experimental.
	HttpMethod_POST HttpMethod = "POST"
	// HTTP PUT.
	// Experimental.
	HttpMethod_PUT HttpMethod = "PUT"
)

