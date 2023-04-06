// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Create a new API Gateway HTTP API endpoint.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//   var booksDefaultFn function
//
//   booksIntegration := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(jsii.String("BooksIntegration"), booksDefaultFn)
//
//   httpApi := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/books"),
//   	Methods: []httpMethod{
//   		apigwv2.*httpMethod_GET,
//   	},
//   	Integration: booksIntegration,
//   })
//
// Experimental.
type HttpApi interface {
	awscdk.Resource
	IApi
	IHttpApi
	// Get the default endpoint for this API.
	// Experimental.
	ApiEndpoint() *string
	// The identifier of this API Gateway API.
	// Experimental.
	ApiId() *string
	// The default stage of this API.
	// Experimental.
	DefaultStage() IHttpStage
	// Specifies whether clients can invoke this HTTP API by using the default execute-api endpoint.
	// Experimental.
	DisableExecuteApiEndpoint() *bool
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The identifier of this API Gateway HTTP API.
	// Experimental.
	HttpApiId() *string
	// A human friendly name for this HTTP API.
	//
	// Note that this is different from `httpApiId`.
	// Experimental.
	HttpApiName() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Get the URL to the default stage of this API.
	//
	// Returns `undefined` if `createDefaultStage` is unset.
	// Experimental.
	Url() *string
	// Add multiple routes that uses the same configuration.
	//
	// The routes all go to the same path, but for different
	// methods.
	// Experimental.
	AddRoutes(options *AddRoutesOptions) *[]HttpRoute
	// Add a new stage.
	// Experimental.
	AddStage(id *string, options *HttpStageOptions) HttpStage
	// Add a new VpcLink.
	// Experimental.
	AddVpcLink(options *VpcLinkProps) VpcLink
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Return the given named metric for this Api Gateway.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of client-side errors captured in a given period.
	// Experimental.
	MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the total number API requests in a given period.
	// Experimental.
	MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the amount of data processed in bytes.
	// Experimental.
	MetricDataProcessed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the time between when API Gateway relays a request to the backend and when it receives a response from the backend.
	// Experimental.
	MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time between when API Gateway receives a request from a client and when it returns a response to the client.
	//
	// The latency includes the integration latency and other API Gateway overhead.
	// Experimental.
	MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of server-side errors captured in a given period.
	// Experimental.
	MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HttpApi
type jsiiProxy_HttpApi struct {
	internal.Type__awscdkResource
	jsiiProxy_IApi
	jsiiProxy_IHttpApi
}

func (j *jsiiProxy_HttpApi) ApiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApi) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApi) DefaultStage() IHttpStage {
	var returns IHttpStage
	_jsii_.Get(
		j,
		"defaultStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApi) DisableExecuteApiEndpoint() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"disableExecuteApiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApi) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApi) HttpApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApi) HttpApiName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpApiName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApi) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApi) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApi) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


// Experimental.
func NewHttpApi(scope constructs.Construct, id *string, props *HttpApiProps) HttpApi {
	_init_.Initialize()

	if err := validateNewHttpApiParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpApi{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpApi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpApi_Override(h HttpApi, scope constructs.Construct, id *string, props *HttpApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpApi",
		[]interface{}{scope, id, props},
		h,
	)
}

// Import an existing HTTP API into this CDK app.
// Experimental.
func HttpApi_FromHttpApiAttributes(scope constructs.Construct, id *string, attrs *HttpApiAttributes) IHttpApi {
	_init_.Initialize()

	if err := validateHttpApi_FromHttpApiAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IHttpApi

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpApi",
		"fromHttpApiAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func HttpApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHttpApi_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func HttpApi_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateHttpApi_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpApi",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func HttpApi_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateHttpApi_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-apigatewayv2-alpha.HttpApi",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpApi) AddRoutes(options *AddRoutesOptions) *[]HttpRoute {
	if err := h.validateAddRoutesParameters(options); err != nil {
		panic(err)
	}
	var returns *[]HttpRoute

	_jsii_.Invoke(
		h,
		"addRoutes",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpApi) AddStage(id *string, options *HttpStageOptions) HttpStage {
	if err := h.validateAddStageParameters(id, options); err != nil {
		panic(err)
	}
	var returns HttpStage

	_jsii_.Invoke(
		h,
		"addStage",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpApi) AddVpcLink(options *VpcLinkProps) VpcLink {
	if err := h.validateAddVpcLinkParameters(options); err != nil {
		panic(err)
	}
	var returns VpcLink

	_jsii_.Invoke(
		h,
		"addVpcLink",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := h.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (h *jsiiProxy_HttpApi) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpApi) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := h.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpApi) GetResourceNameAttribute(nameAttr *string) *string {
	if err := h.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpApi) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := h.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpApi) MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := h.validateMetricClientErrorParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricClientError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpApi) MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := h.validateMetricCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpApi) MetricDataProcessed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := h.validateMetricDataProcessedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricDataProcessed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpApi) MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := h.validateMetricIntegrationLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricIntegrationLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpApi) MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := h.validateMetricLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpApi) MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := h.validateMetricServerErrorParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricServerError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

