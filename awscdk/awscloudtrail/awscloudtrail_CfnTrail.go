package awscloudtrail

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudtrail/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::CloudTrail::Trail`.
//
// Creates a trail that specifies the settings for delivery of log data to an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrail := awscdk.Aws_cloudtrail.NewCfnTrail(this, jsii.String("MyCfnTrail"), &cfnTrailProps{
//   	isLogging: jsii.Boolean(false),
//   	s3BucketName: jsii.String("s3BucketName"),
//
//   	// the properties below are optional
//   	cloudWatchLogsLogGroupArn: jsii.String("cloudWatchLogsLogGroupArn"),
//   	cloudWatchLogsRoleArn: jsii.String("cloudWatchLogsRoleArn"),
//   	enableLogFileValidation: jsii.Boolean(false),
//   	eventSelectors: []interface{}{
//   		&eventSelectorProperty{
//   			dataResources: []interface{}{
//   				&dataResourceProperty{
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   			excludeManagementEventSources: []*string{
//   				jsii.String("excludeManagementEventSources"),
//   			},
//   			includeManagementEvents: jsii.Boolean(false),
//   			readWriteType: jsii.String("readWriteType"),
//   		},
//   	},
//   	includeGlobalServiceEvents: jsii.Boolean(false),
//   	insightSelectors: []interface{}{
//   		&insightSelectorProperty{
//   			insightType: jsii.String("insightType"),
//   		},
//   	},
//   	isMultiRegionTrail: jsii.Boolean(false),
//   	isOrganizationTrail: jsii.Boolean(false),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   	snsTopicName: jsii.String("snsTopicName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	trailName: jsii.String("trailName"),
//   })
//
type CfnTrail interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `Ref` returns the ARN of the CloudTrail trail, such as `arn:aws:cloudtrail:us-east-2:123456789012:trail/myCloudTrail` .
	AttrArn() *string
	// `Ref` returns the ARN of the Amazon SNS topic that's associated with the CloudTrail trail, such as `arn:aws:sns:us-east-2:123456789012:mySNSTopic` .
	AttrSnsTopicArn() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Specifies a log group name using an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs are delivered.
	//
	// You must use a log group that exists in your account.
	//
	// Not required unless you specify `CloudWatchLogsRoleArn` .
	CloudWatchLogsLogGroupArn() *string
	SetCloudWatchLogsLogGroupArn(val *string)
	// Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's log group.
	//
	// You must use a role that exists in your account.
	CloudWatchLogsRoleArn() *string
	SetCloudWatchLogsRoleArn(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// Specifies whether log file validation is enabled. The default is false.
	//
	// > When you disable log file integrity validation, the chain of digest files is broken after one hour. CloudTrail does not create digest files for log files that were delivered during a period in which log file integrity validation was disabled. For example, if you enable log file integrity validation at noon on January 1, disable it at noon on January 2, and re-enable it at noon on January 10, digest files will not be created for the log files delivered from noon on January 2 to noon on January 10. The same applies whenever you stop CloudTrail logging or delete a trail.
	EnableLogFileValidation() interface{}
	SetEnableLogFileValidation(val interface{})
	// Use event selectors to further specify the management and data event settings for your trail.
	//
	// By default, trails created without specific event selectors will be configured to log all read and write management events, and no data events. When an event occurs in your account, CloudTrail evaluates the event selector for all trails. For each trail, if the event matches any event selector, the trail processes and logs the event. If the event doesn't match any event selector, the trail doesn't log the event.
	//
	// You can configure up to five event selectors for a trail.
	EventSelectors() interface{}
	SetEventSelectors(val interface{})
	// Specifies whether the trail is publishing events from global services such as IAM to the log files.
	IncludeGlobalServiceEvents() interface{}
	SetIncludeGlobalServiceEvents(val interface{})
	// A JSON string that contains the insight types you want to log on a trail.
	//
	// `ApiCallRateInsight` and `ApiErrorRateInsight` are valid Insight types.
	//
	// The `ApiCallRateInsight` Insights type analyzes write-only management API calls that are aggregated per minute against a baseline API call volume.
	//
	// The `ApiErrorRateInsight` Insights type analyzes management API calls that result in error codes. The error is shown if the API call is unsuccessful.
	InsightSelectors() interface{}
	SetInsightSelectors(val interface{})
	// Whether the CloudTrail trail is currently logging AWS API calls.
	IsLogging() interface{}
	SetIsLogging(val interface{})
	// Specifies whether the trail applies only to the current region or to all regions.
	//
	// The default is false. If the trail exists only in the current region and this value is set to true, shadow trails (replications of the trail) will be created in the other regions. If the trail exists in all regions and this value is set to false, the trail will remain in the region where it was created, and its shadow trails in other regions will be deleted. As a best practice, consider using trails that log events in all regions.
	IsMultiRegionTrail() interface{}
	SetIsMultiRegionTrail(val interface{})
	// Specifies whether the trail is applied to all accounts in an organization in AWS Organizations , or only for the current AWS account .
	//
	// The default is false, and cannot be true unless the call is made on behalf of an AWS account that is the management account or delegated administrator account for an organization in AWS Organizations . If the trail is not an organization trail and this is set to `true` , the trail will be created in all AWS accounts that belong to the organization. If the trail is an organization trail and this is set to `false` , the trail will remain in the current AWS account but be deleted from all member accounts in the organization.
	IsOrganizationTrail() interface{}
	SetIsOrganizationTrail(val interface{})
	// Specifies the AWS KMS key ID to use to encrypt the logs delivered by CloudTrail.
	//
	// The value can be an alias name prefixed by "alias/", a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
	//
	// CloudTrail also supports AWS KMS multi-Region keys. For more information about multi-Region keys, see [Using multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) in the *AWS Key Management Service Developer Guide* .
	//
	// Examples:
	//
	// - alias/MyAliasName
	// - arn:aws:kms:us-east-2:123456789012:alias/MyAliasName
	// - arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012
	// - 12345678-1234-1234-1234-123456789012.
	KmsKeyId() *string
	SetKmsKeyId(val *string)
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
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// Specifies the name of the Amazon S3 bucket designated for publishing log files.
	//
	// See [Amazon S3 Bucket Naming Requirements](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/create_trail_naming_policy.html) .
	S3BucketName() *string
	SetS3BucketName(val *string)
	// Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery.
	//
	// For more information, see [Finding Your CloudTrail Log Files](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-find-log-files.html) . The maximum length is 200 characters.
	S3KeyPrefix() *string
	SetS3KeyPrefix(val *string)
	// Specifies the name of the Amazon SNS topic defined for notification of log file delivery.
	//
	// The maximum length is 256 characters.
	SnsTopicName() *string
	SetSnsTopicName(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A custom set of tags (key-value pairs) for this trail.
	Tags() awscdk.TagManager
	// Specifies the name of the trail. The name must meet the following requirements:.
	//
	// - Contain only ASCII letters (a-z, A-Z), numbers (0-9), periods (.), underscores (_), or dashes (-)
	// - Start with a letter or number, and end with a letter or number
	// - Be between 3 and 128 characters
	// - Have no adjacent periods, underscores or dashes. Names like `my-_namespace` and `my--namespace` are not valid.
	// - Not be in IP address format (for example, 192.168.5.4)
	TrailName() *string
	SetTrailName(val *string)
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

// The jsii proxy struct for CfnTrail
type jsiiProxy_CfnTrail struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTrail) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) AttrSnsTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSnsTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) CloudWatchLogsLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogsLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) CloudWatchLogsRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogsRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) EnableLogFileValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLogFileValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) EventSelectors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventSelectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) IncludeGlobalServiceEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeGlobalServiceEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) InsightSelectors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insightSelectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) IsLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) IsMultiRegionTrail() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMultiRegionTrail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) IsOrganizationTrail() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isOrganizationTrail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) S3KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) SnsTopicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsTopicName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) TrailName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trailName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTrail) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CloudTrail::Trail`.
func NewCfnTrail(scope awscdk.Construct, id *string, props *CfnTrailProps) CfnTrail {
	_init_.Initialize()

	if err := validateNewCfnTrailParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTrail{}

	_jsii_.Create(
		"monocdk.aws_cloudtrail.CfnTrail",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CloudTrail::Trail`.
func NewCfnTrail_Override(c CfnTrail, scope awscdk.Construct, id *string, props *CfnTrailProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudtrail.CfnTrail",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTrail)SetCloudWatchLogsLogGroupArn(val *string) {
	_jsii_.Set(
		j,
		"cloudWatchLogsLogGroupArn",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetCloudWatchLogsRoleArn(val *string) {
	_jsii_.Set(
		j,
		"cloudWatchLogsRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetEnableLogFileValidation(val interface{}) {
	if err := j.validateSetEnableLogFileValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLogFileValidation",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetEventSelectors(val interface{}) {
	if err := j.validateSetEventSelectorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventSelectors",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetIncludeGlobalServiceEvents(val interface{}) {
	if err := j.validateSetIncludeGlobalServiceEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeGlobalServiceEvents",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetInsightSelectors(val interface{}) {
	if err := j.validateSetInsightSelectorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insightSelectors",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetIsLogging(val interface{}) {
	if err := j.validateSetIsLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isLogging",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetIsMultiRegionTrail(val interface{}) {
	if err := j.validateSetIsMultiRegionTrailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isMultiRegionTrail",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetIsOrganizationTrail(val interface{}) {
	if err := j.validateSetIsOrganizationTrailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isOrganizationTrail",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetS3BucketName(val *string) {
	if err := j.validateSetS3BucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BucketName",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetS3KeyPrefix(val *string) {
	_jsii_.Set(
		j,
		"s3KeyPrefix",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetSnsTopicName(val *string) {
	_jsii_.Set(
		j,
		"snsTopicName",
		val,
	)
}

func (j *jsiiProxy_CfnTrail)SetTrailName(val *string) {
	_jsii_.Set(
		j,
		"trailName",
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
func CfnTrail_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTrail_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudtrail.CfnTrail",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnTrail_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnTrail_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudtrail.CfnTrail",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnTrail_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTrail_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudtrail.CfnTrail",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTrail_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_cloudtrail.CfnTrail",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTrail) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTrail) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTrail) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTrail) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTrail) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTrail) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTrail) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTrail) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnTrail) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnTrail) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTrail) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTrail) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnTrail) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTrail) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTrail) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTrail) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnTrail) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTrail) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnTrail) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTrail) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTrail) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

