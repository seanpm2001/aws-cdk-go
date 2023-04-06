package awscloudfront


// Properties for creating a Origin Request Policy.
//
// Example:
//   // Creating a custom origin request policy for a Distribution -- all parameters optional
//   var bucketOrigin s3Origin
//
//   myOriginRequestPolicy := cloudfront.NewOriginRequestPolicy(this, jsii.String("OriginRequestPolicy"), &OriginRequestPolicyProps{
//   	OriginRequestPolicyName: jsii.String("MyPolicy"),
//   	Comment: jsii.String("A default policy"),
//   	CookieBehavior: cloudfront.OriginRequestCookieBehavior_None(),
//   	HeaderBehavior: cloudfront.OriginRequestHeaderBehavior_All(jsii.String("CloudFront-Is-Android-Viewer")),
//   	QueryStringBehavior: cloudfront.OriginRequestQueryStringBehavior_AllowList(jsii.String("username")),
//   })
//
//   cloudfront.NewDistribution(this, jsii.String("myDistCustomPolicy"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: bucketOrigin,
//   		OriginRequestPolicy: myOriginRequestPolicy,
//   	},
//   })
//
type OriginRequestPolicyProps struct {
	// A comment to describe the origin request policy.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The cookies from viewer requests to include in origin requests.
	CookieBehavior OriginRequestCookieBehavior `field:"optional" json:"cookieBehavior" yaml:"cookieBehavior"`
	// The HTTP headers to include in origin requests.
	//
	// These can include headers from viewer requests and additional headers added by CloudFront.
	HeaderBehavior OriginRequestHeaderBehavior `field:"optional" json:"headerBehavior" yaml:"headerBehavior"`
	// A unique name to identify the origin request policy.
	//
	// The name must only include '-', '_', or alphanumeric characters.
	OriginRequestPolicyName *string `field:"optional" json:"originRequestPolicyName" yaml:"originRequestPolicyName"`
	// The URL query strings from viewer requests to include in origin requests.
	QueryStringBehavior OriginRequestQueryStringBehavior `field:"optional" json:"queryStringBehavior" yaml:"queryStringBehavior"`
}
