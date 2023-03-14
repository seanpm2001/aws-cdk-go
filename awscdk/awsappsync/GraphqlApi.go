package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticsearch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchservice"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
)

// An AppSync GraphQL API.
//
// Example:
//   api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
//   	Name: jsii.String("api"),
//   	Schema: appsync.SchemaFile_FromAsset(path.join(__dirname, jsii.String("schema.graphql"))),
//   })
//
//   httpDs := api.AddHttpDataSource(jsii.String("ds"), jsii.String("https://states.amazonaws.com"), &HttpDataSourceOptions{
//   	Name: jsii.String("httpDsWithStepF"),
//   	Description: jsii.String("from appsync to StepFunctions Workflow"),
//   	AuthorizationConfig: &AwsIamConfig{
//   		SigningRegion: jsii.String("us-east-1"),
//   		SigningServiceName: jsii.String("states"),
//   	},
//   })
//
//   httpDs.CreateResolver(jsii.String("MutationCallStepFunctionResolver"), &BaseResolverProps{
//   	TypeName: jsii.String("Mutation"),
//   	FieldName: jsii.String("callStepFunction"),
//   	RequestMappingTemplate: appsync.MappingTemplate_FromFile(jsii.String("request.vtl")),
//   	ResponseMappingTemplate: appsync.MappingTemplate_*FromFile(jsii.String("response.vtl")),
//   })
//
type GraphqlApi interface {
	GraphqlApiBase
	// an unique AWS AppSync GraphQL API identifier i.e. 'lxz775lwdrgcndgz3nurvac7oa'.
	ApiId() *string
	// the configured API key, if present.
	ApiKey() *string
	// The AppSyncDomainName of the associated custom domain.
	AppSyncDomainName() *string
	// the ARN of the API.
	Arn() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// the URL of the endpoint created by AppSync.
	GraphqlUrl() *string
	// the CloudWatch Log Group for this API.
	LogGroup() awslogs.ILogGroup
	// The Authorization Types for this GraphQL Api.
	Modes() *[]AuthorizationType
	// the name of the API.
	Name() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// the schema attached to this api.
	Schema() ISchema
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// add a new DynamoDB data source to this API.
	AddDynamoDbDataSource(id *string, table awsdynamodb.ITable, options *DataSourceOptions) DynamoDbDataSource
	// add a new elasticsearch data source to this API.
	// Deprecated: - use `addOpenSearchDataSource`.
	AddElasticsearchDataSource(id *string, domain awselasticsearch.IDomain, options *DataSourceOptions) ElasticsearchDataSource
	// add a new http data source to this API.
	AddHttpDataSource(id *string, endpoint *string, options *HttpDataSourceOptions) HttpDataSource
	// add a new Lambda data source to this API.
	AddLambdaDataSource(id *string, lambdaFunction awslambda.IFunction, options *DataSourceOptions) LambdaDataSource
	// add a new dummy data source to this API.
	//
	// Useful for pipeline resolvers
	// and for backend changes that don't require a data source.
	AddNoneDataSource(id *string, options *DataSourceOptions) NoneDataSource
	// add a new OpenSearch data source to this API.
	AddOpenSearchDataSource(id *string, domain awsopensearchservice.IDomain, options *DataSourceOptions) OpenSearchDataSource
	// add a new Rds data source to this API.
	AddRdsDataSource(id *string, serverlessCluster awsrds.IServerlessCluster, secretStore awssecretsmanager.ISecret, databaseName *string, options *DataSourceOptions) RdsDataSource
	// Add schema dependency to a given construct.
	AddSchemaDependency(construct awscdk.CfnResource) *bool
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// creates a new resolver for this datasource and API using the given properties.
	CreateResolver(id *string, props *ExtendedResolverProps) Resolver
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Adds an IAM policy statement associated with this GraphQLApi to an IAM principal's policy.
	Grant(grantee awsiam.IGrantable, resources IamResource, actions ...*string) awsiam.Grant
	// Adds an IAM policy statement for Mutation access to this GraphQLApi to an IAM principal's policy.
	GrantMutation(grantee awsiam.IGrantable, fields ...*string) awsiam.Grant
	// Adds an IAM policy statement for Query access to this GraphQLApi to an IAM principal's policy.
	GrantQuery(grantee awsiam.IGrantable, fields ...*string) awsiam.Grant
	// Adds an IAM policy statement for Subscription access to this GraphQLApi to an IAM principal's policy.
	GrantSubscription(grantee awsiam.IGrantable, fields ...*string) awsiam.Grant
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for GraphqlApi
type jsiiProxy_GraphqlApi struct {
	jsiiProxy_GraphqlApiBase
}

func (j *jsiiProxy_GraphqlApi) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApi) ApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApi) AppSyncDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appSyncDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApi) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApi) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApi) GraphqlUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphqlUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApi) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApi) Modes() *[]AuthorizationType {
	var returns *[]AuthorizationType
	_jsii_.Get(
		j,
		"modes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApi) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApi) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApi) Schema() ISchema {
	var returns ISchema
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApi) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewGraphqlApi(scope constructs.Construct, id *string, props *GraphqlApiProps) GraphqlApi {
	_init_.Initialize()

	if err := validateNewGraphqlApiParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GraphqlApi{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.GraphqlApi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewGraphqlApi_Override(g GraphqlApi, scope constructs.Construct, id *string, props *GraphqlApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.GraphqlApi",
		[]interface{}{scope, id, props},
		g,
	)
}

// Import a GraphQL API through this function.
func GraphqlApi_FromGraphqlApiAttributes(scope constructs.Construct, id *string, attrs *GraphqlApiAttributes) IGraphqlApi {
	_init_.Initialize()

	if err := validateGraphqlApi_FromGraphqlApiAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IGraphqlApi

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.GraphqlApi",
		"fromGraphqlApiAttributes",
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
func GraphqlApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGraphqlApi_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.GraphqlApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func GraphqlApi_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateGraphqlApi_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.GraphqlApi",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func GraphqlApi_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateGraphqlApi_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.GraphqlApi",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphqlApi) AddDynamoDbDataSource(id *string, table awsdynamodb.ITable, options *DataSourceOptions) DynamoDbDataSource {
	if err := g.validateAddDynamoDbDataSourceParameters(id, table, options); err != nil {
		panic(err)
	}
	var returns DynamoDbDataSource

	_jsii_.Invoke(
		g,
		"addDynamoDbDataSource",
		[]interface{}{id, table, options},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphqlApi) AddElasticsearchDataSource(id *string, domain awselasticsearch.IDomain, options *DataSourceOptions) ElasticsearchDataSource {
	if err := g.validateAddElasticsearchDataSourceParameters(id, domain, options); err != nil {
		panic(err)
	}
	var returns ElasticsearchDataSource

	_jsii_.Invoke(
		g,
		"addElasticsearchDataSource",
		[]interface{}{id, domain, options},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphqlApi) AddHttpDataSource(id *string, endpoint *string, options *HttpDataSourceOptions) HttpDataSource {
	if err := g.validateAddHttpDataSourceParameters(id, endpoint, options); err != nil {
		panic(err)
	}
	var returns HttpDataSource

	_jsii_.Invoke(
		g,
		"addHttpDataSource",
		[]interface{}{id, endpoint, options},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphqlApi) AddLambdaDataSource(id *string, lambdaFunction awslambda.IFunction, options *DataSourceOptions) LambdaDataSource {
	if err := g.validateAddLambdaDataSourceParameters(id, lambdaFunction, options); err != nil {
		panic(err)
	}
	var returns LambdaDataSource

	_jsii_.Invoke(
		g,
		"addLambdaDataSource",
		[]interface{}{id, lambdaFunction, options},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphqlApi) AddNoneDataSource(id *string, options *DataSourceOptions) NoneDataSource {
	if err := g.validateAddNoneDataSourceParameters(id, options); err != nil {
		panic(err)
	}
	var returns NoneDataSource

	_jsii_.Invoke(
		g,
		"addNoneDataSource",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphqlApi) AddOpenSearchDataSource(id *string, domain awsopensearchservice.IDomain, options *DataSourceOptions) OpenSearchDataSource {
	if err := g.validateAddOpenSearchDataSourceParameters(id, domain, options); err != nil {
		panic(err)
	}
	var returns OpenSearchDataSource

	_jsii_.Invoke(
		g,
		"addOpenSearchDataSource",
		[]interface{}{id, domain, options},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphqlApi) AddRdsDataSource(id *string, serverlessCluster awsrds.IServerlessCluster, secretStore awssecretsmanager.ISecret, databaseName *string, options *DataSourceOptions) RdsDataSource {
	if err := g.validateAddRdsDataSourceParameters(id, serverlessCluster, secretStore, options); err != nil {
		panic(err)
	}
	var returns RdsDataSource

	_jsii_.Invoke(
		g,
		"addRdsDataSource",
		[]interface{}{id, serverlessCluster, secretStore, databaseName, options},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphqlApi) AddSchemaDependency(construct awscdk.CfnResource) *bool {
	if err := g.validateAddSchemaDependencyParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		g,
		"addSchemaDependency",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphqlApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := g.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (g *jsiiProxy_GraphqlApi) CreateResolver(id *string, props *ExtendedResolverProps) Resolver {
	if err := g.validateCreateResolverParameters(id, props); err != nil {
		panic(err)
	}
	var returns Resolver

	_jsii_.Invoke(
		g,
		"createResolver",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphqlApi) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphqlApi) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := g.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphqlApi) GetResourceNameAttribute(nameAttr *string) *string {
	if err := g.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphqlApi) Grant(grantee awsiam.IGrantable, resources IamResource, actions ...*string) awsiam.Grant {
	if err := g.validateGrantParameters(grantee, resources); err != nil {
		panic(err)
	}
	args := []interface{}{grantee, resources}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		g,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphqlApi) GrantMutation(grantee awsiam.IGrantable, fields ...*string) awsiam.Grant {
	if err := g.validateGrantMutationParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range fields {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		g,
		"grantMutation",
		args,
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphqlApi) GrantQuery(grantee awsiam.IGrantable, fields ...*string) awsiam.Grant {
	if err := g.validateGrantQueryParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range fields {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		g,
		"grantQuery",
		args,
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphqlApi) GrantSubscription(grantee awsiam.IGrantable, fields ...*string) awsiam.Grant {
	if err := g.validateGrantSubscriptionParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range fields {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		g,
		"grantSubscription",
		args,
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphqlApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

