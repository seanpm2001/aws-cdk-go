package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticsearch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchservice"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
)

// Base Class for GraphQL API.
type GraphqlApiBase interface {
	awscdk.Resource
	IGraphqlApi
	// an unique AWS AppSync GraphQL API identifier i.e. 'lxz775lwdrgcndgz3nurvac7oa'.
	ApiId() *string
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
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
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
	// Add schema dependency if not imported.
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
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for GraphqlApiBase
type jsiiProxy_GraphqlApiBase struct {
	internal.Type__awscdkResource
	jsiiProxy_IGraphqlApi
}

func (j *jsiiProxy_GraphqlApiBase) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApiBase) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApiBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApiBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApiBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApiBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewGraphqlApiBase_Override(g GraphqlApiBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.GraphqlApiBase",
		[]interface{}{scope, id, props},
		g,
	)
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
func GraphqlApiBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGraphqlApiBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.GraphqlApiBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func GraphqlApiBase_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateGraphqlApiBase_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.GraphqlApiBase",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func GraphqlApiBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateGraphqlApiBase_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.GraphqlApiBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphqlApiBase) AddDynamoDbDataSource(id *string, table awsdynamodb.ITable, options *DataSourceOptions) DynamoDbDataSource {
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

func (g *jsiiProxy_GraphqlApiBase) AddElasticsearchDataSource(id *string, domain awselasticsearch.IDomain, options *DataSourceOptions) ElasticsearchDataSource {
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

func (g *jsiiProxy_GraphqlApiBase) AddHttpDataSource(id *string, endpoint *string, options *HttpDataSourceOptions) HttpDataSource {
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

func (g *jsiiProxy_GraphqlApiBase) AddLambdaDataSource(id *string, lambdaFunction awslambda.IFunction, options *DataSourceOptions) LambdaDataSource {
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

func (g *jsiiProxy_GraphqlApiBase) AddNoneDataSource(id *string, options *DataSourceOptions) NoneDataSource {
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

func (g *jsiiProxy_GraphqlApiBase) AddOpenSearchDataSource(id *string, domain awsopensearchservice.IDomain, options *DataSourceOptions) OpenSearchDataSource {
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

func (g *jsiiProxy_GraphqlApiBase) AddRdsDataSource(id *string, serverlessCluster awsrds.IServerlessCluster, secretStore awssecretsmanager.ISecret, databaseName *string, options *DataSourceOptions) RdsDataSource {
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

func (g *jsiiProxy_GraphqlApiBase) AddSchemaDependency(construct awscdk.CfnResource) *bool {
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

func (g *jsiiProxy_GraphqlApiBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := g.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (g *jsiiProxy_GraphqlApiBase) CreateResolver(id *string, props *ExtendedResolverProps) Resolver {
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

func (g *jsiiProxy_GraphqlApiBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphqlApiBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (g *jsiiProxy_GraphqlApiBase) GetResourceNameAttribute(nameAttr *string) *string {
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

func (g *jsiiProxy_GraphqlApiBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

