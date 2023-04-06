// The CDK Construct Library for AWS::Amplify
package awscdkamplifyalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkamplifyalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkamplifyalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An Amplify Console application.
//
// Example:
//   amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &AppProps{
//   	SourceCodeProvider: amplify.NewGitHubSourceCodeProvider(&GitHubSourceCodeProviderProps{
//   		Owner: jsii.String("<user>"),
//   		Repository: jsii.String("<repo>"),
//   		OauthToken: awscdk.SecretValue_SecretsManager(jsii.String("my-github-token")),
//   	}),
//   	AutoBranchCreation: &AutoBranchCreation{
//   		 // Automatically connect branches that match a pattern set
//   		Patterns: []*string{
//   			jsii.String("feature/*"),
//   			jsii.String("test/*"),
//   		},
//   	},
//   	AutoBranchDeletion: jsii.Boolean(true),
//   })
//
// Experimental.
type App interface {
	awscdk.Resource
	IApp
	awsiam.IGrantable
	// The application id.
	// Experimental.
	AppId() *string
	// The name of the application.
	// Experimental.
	AppName() *string
	// The ARN of the application.
	// Experimental.
	Arn() *string
	// The default domain of the application.
	// Experimental.
	DefaultDomain() *string
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
	// The principal to grant permissions to.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
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
	// Adds an environment variable to the auto created branch.
	//
	// All environment variables that you add are encrypted to prevent rogue
	// access so you can use them to store secret information.
	// Experimental.
	AddAutoBranchEnvironment(name *string, value *string) App
	// Adds a branch to this application.
	// Experimental.
	AddBranch(id *string, options *BranchOptions) Branch
	// Adds a custom rewrite/redirect rule to this application.
	// Experimental.
	AddCustomRule(rule CustomRule) App
	// Adds a domain to this application.
	// Experimental.
	AddDomain(id *string, options *DomainOptions) Domain
	// Adds an environment variable to this application.
	//
	// All environment variables that you add are encrypted to prevent rogue
	// access so you can use them to store secret information.
	// Experimental.
	AddEnvironment(name *string, value *string) App
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
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for App
type jsiiProxy_App struct {
	internal.Type__awscdkResource
	jsiiProxy_IApp
	internal.Type__awsiamIGrantable
}

func (j *jsiiProxy_App) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) AppName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) DefaultDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_App) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewApp(scope constructs.Construct, id *string, props *AppProps) App {
	_init_.Initialize()

	if err := validateNewAppParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_App{}

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.App",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewApp_Override(a App, scope constructs.Construct, id *string, props *AppProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-amplify-alpha.App",
		[]interface{}{scope, id, props},
		a,
	)
}

// Import an existing application.
// Experimental.
func App_FromAppId(scope constructs.Construct, id *string, appId *string) IApp {
	_init_.Initialize()

	if err := validateApp_FromAppIdParameters(scope, id, appId); err != nil {
		panic(err)
	}
	var returns IApp

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-amplify-alpha.App",
		"fromAppId",
		[]interface{}{scope, id, appId},
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
func App_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-amplify-alpha.App",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func App_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateApp_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-amplify-alpha.App",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func App_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateApp_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-amplify-alpha.App",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) AddAutoBranchEnvironment(name *string, value *string) App {
	if err := a.validateAddAutoBranchEnvironmentParameters(name, value); err != nil {
		panic(err)
	}
	var returns App

	_jsii_.Invoke(
		a,
		"addAutoBranchEnvironment",
		[]interface{}{name, value},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) AddBranch(id *string, options *BranchOptions) Branch {
	if err := a.validateAddBranchParameters(id, options); err != nil {
		panic(err)
	}
	var returns Branch

	_jsii_.Invoke(
		a,
		"addBranch",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) AddCustomRule(rule CustomRule) App {
	if err := a.validateAddCustomRuleParameters(rule); err != nil {
		panic(err)
	}
	var returns App

	_jsii_.Invoke(
		a,
		"addCustomRule",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) AddDomain(id *string, options *DomainOptions) Domain {
	if err := a.validateAddDomainParameters(id, options); err != nil {
		panic(err)
	}
	var returns Domain

	_jsii_.Invoke(
		a,
		"addDomain",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) AddEnvironment(name *string, value *string) App {
	if err := a.validateAddEnvironmentParameters(name, value); err != nil {
		panic(err)
	}
	var returns App

	_jsii_.Invoke(
		a,
		"addEnvironment",
		[]interface{}{name, value},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := a.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_App) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := a.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) GetResourceNameAttribute(nameAttr *string) *string {
	if err := a.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

