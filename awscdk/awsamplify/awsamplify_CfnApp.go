package awsamplify

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsamplify/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::Amplify::App`.
//
// The AWS::Amplify::App resource creates Apps in the Amplify Console. An App is a collection of branches.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApp := awscdk.Aws_amplify.NewCfnApp(this, jsii.String("MyCfnApp"), &cfnAppProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	accessToken: jsii.String("accessToken"),
//   	autoBranchCreationConfig: &autoBranchCreationConfigProperty{
//   		autoBranchCreationPatterns: []*string{
//   			jsii.String("autoBranchCreationPatterns"),
//   		},
//   		basicAuthConfig: &basicAuthConfigProperty{
//   			enableBasicAuth: jsii.Boolean(false),
//   			password: jsii.String("password"),
//   			username: jsii.String("username"),
//   		},
//   		buildSpec: jsii.String("buildSpec"),
//   		enableAutoBranchCreation: jsii.Boolean(false),
//   		enableAutoBuild: jsii.Boolean(false),
//   		enablePerformanceMode: jsii.Boolean(false),
//   		enablePullRequestPreview: jsii.Boolean(false),
//   		environmentVariables: []interface{}{
//   			&environmentVariableProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		pullRequestEnvironmentName: jsii.String("pullRequestEnvironmentName"),
//   		stage: jsii.String("stage"),
//   	},
//   	basicAuthConfig: &basicAuthConfigProperty{
//   		enableBasicAuth: jsii.Boolean(false),
//   		password: jsii.String("password"),
//   		username: jsii.String("username"),
//   	},
//   	buildSpec: jsii.String("buildSpec"),
//   	customHeaders: jsii.String("customHeaders"),
//   	customRules: []interface{}{
//   		&customRuleProperty{
//   			source: jsii.String("source"),
//   			target: jsii.String("target"),
//
//   			// the properties below are optional
//   			condition: jsii.String("condition"),
//   			status: jsii.String("status"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	enableBranchAutoDeletion: jsii.Boolean(false),
//   	environmentVariables: []interface{}{
//   		&environmentVariableProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	iamServiceRole: jsii.String("iamServiceRole"),
//   	oauthToken: jsii.String("oauthToken"),
//   	repository: jsii.String("repository"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnApp interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The personal access token for a GitHub repository for an Amplify app.
	//
	// The personal access token is used to authorize access to a GitHub repository using the Amplify GitHub App. The token is not stored.
	//
	// Use `AccessToken` for GitHub repositories only. To authorize access to a repository provider such as Bitbucket or CodeCommit, use `OauthToken` .
	//
	// You must specify either `AccessToken` or `OauthToken` when you create a new app.
	//
	// Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see [Migrating an existing OAuth app to the Amplify GitHub App](https://docs.aws.amazon.com/amplify/latest/userguide/setting-up-GitHub-access.html#migrating-to-github-app-auth) in the *Amplify User Guide* .
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 255.
	AccessToken() *string
	SetAccessToken(val *string)
	// Unique Id for the Amplify App.
	AttrAppId() *string
	// Name for the Amplify App.
	AttrAppName() *string
	// ARN for the Amplify App.
	AttrArn() *string
	// Default domain for the Amplify App.
	AttrDefaultDomain() *string
	// Sets the configuration for your automatic branch creation.
	AutoBranchCreationConfig() interface{}
	SetAutoBranchCreationConfig(val interface{})
	// The credentials for basic authorization for an Amplify app.
	//
	// You must base64-encode the authorization credentials and provide them in the format `user:password` .
	BasicAuthConfig() interface{}
	SetBasicAuthConfig(val interface{})
	// The build specification (build spec) for an Amplify app.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 25000.
	//
	// *Pattern:* (?s).+
	BuildSpec() *string
	SetBuildSpec(val *string)
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The custom HTTP headers for an Amplify app.
	//
	// *Length Constraints:* Minimum length of 0. Maximum length of 25000.
	//
	// *Pattern:* (?s).*
	CustomHeaders() *string
	SetCustomHeaders(val *string)
	// The custom rewrite and redirect rules for an Amplify app.
	CustomRules() interface{}
	SetCustomRules(val interface{})
	// The description for an Amplify app.
	//
	// *Length Constraints:* Maximum length of 1000.
	//
	// *Pattern:* (?s).*
	Description() *string
	SetDescription(val *string)
	// Automatically disconnect a branch in the Amplify Console when you delete a branch from your Git repository.
	EnableBranchAutoDeletion() interface{}
	SetEnableBranchAutoDeletion(val interface{})
	// The environment variables map for an Amplify app.
	EnvironmentVariables() interface{}
	SetEnvironmentVariables(val interface{})
	// The AWS Identity and Access Management (IAM) service role for the Amazon Resource Name (ARN) of the Amplify app.
	//
	// *Length Constraints:* Minimum length of 0. Maximum length of 1000.
	//
	// *Pattern:* (?s).*
	IamServiceRole() *string
	SetIamServiceRole(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The name for an Amplify app.
	//
	// *Length Constraints:* Minimum length of 1. Maximum length of 255.
	//
	// *Pattern:* (?s).+
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The OAuth token for a third-party source control system for an Amplify app.
	//
	// The OAuth token is used to create a webhook and a read-only deploy key using SSH cloning. The OAuth token is not stored.
	//
	// Use `OauthToken` for repository providers other than GitHub, such as Bitbucket or CodeCommit. To authorize access to GitHub as your repository provider, use `AccessToken` .
	//
	// You must specify either `OauthToken` or `AccessToken` when you create a new app.
	//
	// Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see [Migrating an existing OAuth app to the Amplify GitHub App](https://docs.aws.amazon.com/amplify/latest/userguide/setting-up-GitHub-access.html#migrating-to-github-app-auth) in the *Amplify User Guide* .
	//
	// *Length Constraints:* Maximum length of 1000.
	//
	// *Pattern:* (?s).*
	OauthToken() *string
	SetOauthToken(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The repository for an Amplify app.
	//
	// *Pattern:* (?s).*
	Repository() *string
	SetRepository(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The tag for an Amplify app.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnApp
type jsiiProxy_CfnApp struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApp) AccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) AttrAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) AttrAppName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAppName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) AttrDefaultDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDefaultDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) AutoBranchCreationConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoBranchCreationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) BasicAuthConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"basicAuthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) BuildSpec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CustomHeaders() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) CustomRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) EnableBranchAutoDeletion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBranchAutoDeletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) EnvironmentVariables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) IamServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamServiceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) OauthToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Repository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApp) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Amplify::App`.
func NewCfnApp(scope awscdk.Construct, id *string, props *CfnAppProps) CfnApp {
	_init_.Initialize()

	if err := validateNewCfnAppParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApp{}

	_jsii_.Create(
		"monocdk.aws_amplify.CfnApp",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Amplify::App`.
func NewCfnApp_Override(c CfnApp, scope awscdk.Construct, id *string, props *CfnAppProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_amplify.CfnApp",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApp)SetAccessToken(val *string) {
	_jsii_.Set(
		j,
		"accessToken",
		val,
	)
}

func (j *jsiiProxy_CfnApp)SetAutoBranchCreationConfig(val interface{}) {
	if err := j.validateSetAutoBranchCreationConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoBranchCreationConfig",
		val,
	)
}

func (j *jsiiProxy_CfnApp)SetBasicAuthConfig(val interface{}) {
	if err := j.validateSetBasicAuthConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"basicAuthConfig",
		val,
	)
}

func (j *jsiiProxy_CfnApp)SetBuildSpec(val *string) {
	_jsii_.Set(
		j,
		"buildSpec",
		val,
	)
}

func (j *jsiiProxy_CfnApp)SetCustomHeaders(val *string) {
	_jsii_.Set(
		j,
		"customHeaders",
		val,
	)
}

func (j *jsiiProxy_CfnApp)SetCustomRules(val interface{}) {
	if err := j.validateSetCustomRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customRules",
		val,
	)
}

func (j *jsiiProxy_CfnApp)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnApp)SetEnableBranchAutoDeletion(val interface{}) {
	if err := j.validateSetEnableBranchAutoDeletionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableBranchAutoDeletion",
		val,
	)
}

func (j *jsiiProxy_CfnApp)SetEnvironmentVariables(val interface{}) {
	if err := j.validateSetEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_CfnApp)SetIamServiceRole(val *string) {
	_jsii_.Set(
		j,
		"iamServiceRole",
		val,
	)
}

func (j *jsiiProxy_CfnApp)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnApp)SetOauthToken(val *string) {
	_jsii_.Set(
		j,
		"oauthToken",
		val,
	)
}

func (j *jsiiProxy_CfnApp)SetRepository(val *string) {
	_jsii_.Set(
		j,
		"repository",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnApp_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApp_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplify.CfnApp",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnApp_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnApp_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplify.CfnApp",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplify.CfnApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApp_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_amplify.CfnApp",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApp) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApp) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApp) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApp) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApp) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApp) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApp) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnApp) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApp) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApp) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnApp) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnApp) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApp) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApp) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}
