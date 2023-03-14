package awscloudfront


// An object that determines whether any HTTP headers (and if so, which headers) are included in requests that CloudFront sends to the origin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   headersConfigProperty := &HeadersConfigProperty{
//   	HeaderBehavior: jsii.String("headerBehavior"),
//
//   	// the properties below are optional
//   	Headers: []*string{
//   		jsii.String("headers"),
//   	},
//   }
//
type CfnOriginRequestPolicy_HeadersConfigProperty struct {
	// Determines whether any HTTP headers are included in requests that CloudFront sends to the origin. Valid values are:.
	//
	// - `none` – No HTTP headers in viewer requests are included in requests that CloudFront sends to the origin. Even when this field is set to `none` , any headers that are listed in a `CachePolicy` *are* included in origin requests.
	// - `whitelist` – Only the HTTP headers that are listed in the `Headers` type are included in requests that CloudFront sends to the origin.
	// - `allViewer` – All HTTP headers in viewer requests are included in requests that CloudFront sends to the origin.
	// - `allViewerAndWhitelistCloudFront` – All HTTP headers in viewer requests and the additional CloudFront headers that are listed in the `Headers` type are included in requests that CloudFront sends to the origin. The additional headers are added by CloudFront.
	// - `allExcept` – All HTTP headers in viewer requests are included in requests that CloudFront sends to the origin, **except** for those listed in the `Headers` type, which are not included.
	HeaderBehavior *string `field:"required" json:"headerBehavior" yaml:"headerBehavior"`
	// Contains a list of HTTP header names.
	Headers *[]*string `field:"optional" json:"headers" yaml:"headers"`
}

