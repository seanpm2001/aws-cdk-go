package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An option group.
//
// Example:
//   // Set open cursors with parameter group
//   parameterGroup := rds.NewParameterGroup(this, jsii.String("ParameterGroup"), &parameterGroupProps{
//   	engine: rds.databaseInstanceEngine.oracleSe2(&oracleSe2InstanceEngineProps{
//   		version: rds.oracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
//   	}),
//   	parameters: map[string]*string{
//   		"open_cursors": jsii.String("2500"),
//   	},
//   })
//
//   optionGroup := rds.NewOptionGroup(this, jsii.String("OptionGroup"), &optionGroupProps{
//   	engine: rds.*databaseInstanceEngine.oracleSe2(&oracleSe2InstanceEngineProps{
//   		version: rds.*oracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
//   	}),
//   	configurations: []optionConfiguration{
//   		&optionConfiguration{
//   			name: jsii.String("LOCATOR"),
//   		},
//   		&optionConfiguration{
//   			name: jsii.String("OEM"),
//   			port: jsii.Number(1158),
//   			vpc: vpc,
//   		},
//   	},
//   })
//
//   // Allow connections to OEM
//   optionGroup.optionConnections.oEM.connections.allowDefaultPortFromAnyIpv4()
//
//   // Database instance with production values
//   instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &databaseInstanceProps{
//   	engine: rds.*databaseInstanceEngine.oracleSe2(&oracleSe2InstanceEngineProps{
//   		version: rds.*oracleEngineVersion_VER_19_0_0_0_2020_04_R1(),
//   	}),
//   	licenseModel: rds.licenseModel_BRING_YOUR_OWN_LICENSE,
//   	instanceType: ec2.instanceType.of(ec2.instanceClass_BURSTABLE3, ec2.instanceSize_MEDIUM),
//   	multiAz: jsii.Boolean(true),
//   	storageType: rds.storageType_IO1,
//   	credentials: rds.credentials.fromUsername(jsii.String("syscdk")),
//   	vpc: vpc,
//   	databaseName: jsii.String("ORCL"),
//   	storageEncrypted: jsii.Boolean(true),
//   	backupRetention: cdk.duration.days(jsii.Number(7)),
//   	monitoringInterval: cdk.*duration.seconds(jsii.Number(60)),
//   	enablePerformanceInsights: jsii.Boolean(true),
//   	cloudwatchLogsExports: []*string{
//   		jsii.String("trace"),
//   		jsii.String("audit"),
//   		jsii.String("alert"),
//   		jsii.String("listener"),
//   	},
//   	cloudwatchLogsRetention: logs.retentionDays_ONE_MONTH,
//   	autoMinorVersionUpgrade: jsii.Boolean(true),
//   	 // required to be true if LOCATOR is used in the option group
//   	optionGroup: optionGroup,
//   	parameterGroup: parameterGroup,
//   	removalPolicy: awscdk.RemovalPolicy_DESTROY,
//   })
//
//   // Allow connections on default port from any IPV4
//   instance.connections.allowDefaultPortFromAnyIpv4()
//
//   // Rotate the master user password every 30 days
//   instance.addRotationSingleUser()
//
//   // Add alarm for high CPU
//   // Add alarm for high CPU
//   cloudwatch.NewAlarm(this, jsii.String("HighCPU"), &alarmProps{
//   	metric: instance.metricCPUUtilization(),
//   	threshold: jsii.Number(90),
//   	evaluationPeriods: jsii.Number(1),
//   })
//
//   // Trigger Lambda function on instance availability events
//   fn := lambda.NewFunction(this, jsii.String("Function"), &functionProps{
//   	code: lambda.code.fromInline(jsii.String("exports.handler = (event) => console.log(event);")),
//   	handler: jsii.String("index.handler"),
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   })
//
//   availabilityRule := instance.onEvent(jsii.String("Availability"), &onEventOptions{
//   	target: targets.NewLambdaFunction(fn),
//   })
//   availabilityRule.addEventPattern(&eventPattern{
//   	detail: map[string]interface{}{
//   		"EventCategories": []interface{}{
//   			jsii.String("availability"),
//   		},
//   	},
//   })
//
type OptionGroup interface {
	awscdk.Resource
	IOptionGroup
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
	// The connections object for the options.
	OptionConnections() *map[string]awsec2.Connections
	// The name of the option group.
	OptionGroupName() *string
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Adds a configuration to this OptionGroup.
	//
	// This method is a no-op for an imported OptionGroup.
	AddConfiguration(configuration *OptionConfiguration) *bool
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

// The jsii proxy struct for OptionGroup
type jsiiProxy_OptionGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_IOptionGroup
}

func (j *jsiiProxy_OptionGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OptionGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OptionGroup) OptionConnections() *map[string]awsec2.Connections {
	var returns *map[string]awsec2.Connections
	_jsii_.Get(
		j,
		"optionConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OptionGroup) OptionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OptionGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OptionGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewOptionGroup(scope constructs.Construct, id *string, props *OptionGroupProps) OptionGroup {
	_init_.Initialize()

	if err := validateNewOptionGroupParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_OptionGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.OptionGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewOptionGroup_Override(o OptionGroup, scope constructs.Construct, id *string, props *OptionGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.OptionGroup",
		[]interface{}{scope, id, props},
		o,
	)
}

// Import an existing option group.
func OptionGroup_FromOptionGroupName(scope constructs.Construct, id *string, optionGroupName *string) IOptionGroup {
	_init_.Initialize()

	if err := validateOptionGroup_FromOptionGroupNameParameters(scope, id, optionGroupName); err != nil {
		panic(err)
	}
	var returns IOptionGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.OptionGroup",
		"fromOptionGroupName",
		[]interface{}{scope, id, optionGroupName},
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
func OptionGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOptionGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.OptionGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func OptionGroup_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateOptionGroup_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.OptionGroup",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func OptionGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateOptionGroup_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.OptionGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OptionGroup) AddConfiguration(configuration *OptionConfiguration) *bool {
	if err := o.validateAddConfigurationParameters(configuration); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		o,
		"addConfiguration",
		[]interface{}{configuration},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OptionGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := o.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (o *jsiiProxy_OptionGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OptionGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := o.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OptionGroup) GetResourceNameAttribute(nameAttr *string) *string {
	if err := o.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OptionGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

