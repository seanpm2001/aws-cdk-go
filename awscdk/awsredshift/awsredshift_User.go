package awsredshift

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsredshift/internal"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v3"
)

// A user in a Redshift cluster.
//
// Example:
//   user := awscdk.NewUser(this, jsii.String("User"), &userProps{
//   	cluster: cluster,
//   	databaseName: jsii.String("databaseName"),
//   })
//   cluster.addRotationMultiUser(jsii.String("MultiUserRotation"), &rotationMultiUserOptions{
//   	secret: user.secret,
//   })
//
// Experimental.
type User interface {
	awscdk.Construct
	IUser
	// The cluster where the table is located.
	// Experimental.
	Cluster() ICluster
	// The name of the database where the table is located.
	// Experimental.
	DatabaseName() *string
	// Experimental.
	DatabaseProps() *DatabaseOptions
	// Experimental.
	SetDatabaseProps(val *DatabaseOptions)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The password of the user.
	// Experimental.
	Password() awscdk.SecretValue
	// The Secrets Manager secret of the user.
	// Experimental.
	Secret() awssecretsmanager.ISecret
	// The name of the user.
	// Experimental.
	Username() *string
	// Grant this user privilege to access a table.
	// Experimental.
	AddTablePrivileges(table ITable, actions ...TableAction)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be destroyed (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	//
	// This resource is destroyed by default.
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
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
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
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
}

// The jsii proxy struct for User
type jsiiProxy_User struct {
	internal.Type__awscdkConstruct
	jsiiProxy_IUser
}

func (j *jsiiProxy_User) Cluster() ICluster {
	var returns ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) DatabaseProps() *DatabaseOptions {
	var returns *DatabaseOptions
	_jsii_.Get(
		j,
		"databaseProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Password() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Secret() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_User) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}


// Experimental.
func NewUser(scope constructs.Construct, id *string, props *UserProps) User {
	_init_.Initialize()

	if err := validateNewUserParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_User{}

	_jsii_.Create(
		"monocdk.aws_redshift.User",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewUser_Override(u User, scope constructs.Construct, id *string, props *UserProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_redshift.User",
		[]interface{}{scope, id, props},
		u,
	)
}

func (j *jsiiProxy_User)SetDatabaseProps(val *DatabaseOptions) {
	if err := j.validateSetDatabasePropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseProps",
		val,
	)
}

// Specify a Redshift user using credentials that already exist.
// Experimental.
func User_FromUserAttributes(scope constructs.Construct, id *string, attrs *UserAttributes) IUser {
	_init_.Initialize()

	if err := validateUser_FromUserAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IUser

	_jsii_.StaticInvoke(
		"monocdk.aws_redshift.User",
		"fromUserAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func User_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUser_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_redshift.User",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) AddTablePrivileges(table ITable, actions ...TableAction) {
	if err := u.validateAddTablePrivilegesParameters(table); err != nil {
		panic(err)
	}
	args := []interface{}{table}
	for _, a := range actions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		u,
		"addTablePrivileges",
		args,
	)
}

func (u *jsiiProxy_User) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := u.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (u *jsiiProxy_User) OnPrepare() {
	_jsii_.InvokeVoid(
		u,
		"onPrepare",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) OnSynthesize(session constructs.ISynthesisSession) {
	if err := u.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (u *jsiiProxy_User) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) Prepare() {
	_jsii_.InvokeVoid(
		u,
		"prepare",
		nil, // no parameters
	)
}

func (u *jsiiProxy_User) Synthesize(session awscdk.ISynthesisSession) {
	if err := u.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"synthesize",
		[]interface{}{session},
	)
}

func (u *jsiiProxy_User) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_User) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

