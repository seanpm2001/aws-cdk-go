package awsmediatailor

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediatailor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::MediaTailor::PlaybackConfiguration`.
//
// Adds a new playback configuration to AWS Elemental MediaTailor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationAliases interface{}
//
//   cfnPlaybackConfiguration := awscdk.Aws_mediatailor.NewCfnPlaybackConfiguration(this, jsii.String("MyCfnPlaybackConfiguration"), &cfnPlaybackConfigurationProps{
//   	adDecisionServerUrl: jsii.String("adDecisionServerUrl"),
//   	name: jsii.String("name"),
//   	videoContentSourceUrl: jsii.String("videoContentSourceUrl"),
//
//   	// the properties below are optional
//   	availSuppression: &availSuppressionProperty{
//   		mode: jsii.String("mode"),
//   		value: jsii.String("value"),
//   	},
//   	bumper: &bumperProperty{
//   		endUrl: jsii.String("endUrl"),
//   		startUrl: jsii.String("startUrl"),
//   	},
//   	cdnConfiguration: &cdnConfigurationProperty{
//   		adSegmentUrlPrefix: jsii.String("adSegmentUrlPrefix"),
//   		contentSegmentUrlPrefix: jsii.String("contentSegmentUrlPrefix"),
//   	},
//   	configurationAliases: map[string]interface{}{
//   		"configurationAliasesKey": configurationAliases,
//   	},
//   	dashConfiguration: &dashConfigurationProperty{
//   		manifestEndpointPrefix: jsii.String("manifestEndpointPrefix"),
//   		mpdLocation: jsii.String("mpdLocation"),
//   		originManifestType: jsii.String("originManifestType"),
//   	},
//   	hlsConfiguration: &hlsConfigurationProperty{
//   		manifestEndpointPrefix: jsii.String("manifestEndpointPrefix"),
//   	},
//   	livePreRollConfiguration: &livePreRollConfigurationProperty{
//   		adDecisionServerUrl: jsii.String("adDecisionServerUrl"),
//   		maxDurationSeconds: jsii.Number(123),
//   	},
//   	manifestProcessingRules: &manifestProcessingRulesProperty{
//   		adMarkerPassthrough: &adMarkerPassthroughProperty{
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   	personalizationThresholdSeconds: jsii.Number(123),
//   	slateAdUrl: jsii.String("slateAdUrl"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	transcodeProfileName: jsii.String("transcodeProfileName"),
//   })
//
type CfnPlaybackConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The URL for the ad decision server (ADS).
	//
	// This includes the specification of static parameters and placeholders for dynamic parameters. MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing you can provide a static VAST URL. The maximum length is 25,000 characters.
	AdDecisionServerUrl() *string
	SetAdDecisionServerUrl(val *string)
	AttrDashConfigurationManifestEndpointPrefix() *string
	AttrHlsConfigurationManifestEndpointPrefix() *string
	AttrPlaybackConfigurationArn() *string
	AttrPlaybackEndpointPrefix() *string
	AttrSessionInitializationEndpointPrefix() *string
	// The configuration for avail suppression, also known as ad suppression.
	//
	// For more information about ad suppression, see [Ad Suppression](https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html) .
	AvailSuppression() interface{}
	SetAvailSuppression(val interface{})
	// The configuration for bumpers.
	//
	// Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see [Bumpers](https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html) .
	Bumper() interface{}
	SetBumper(val interface{})
	// The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
	CdnConfiguration() interface{}
	SetCdnConfiguration(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The player parameters and aliases used as dynamic variables during session initialization.
	//
	// For more information, see [Domain Variables](https://docs.aws.amazon.com/mediatailor/latest/ug/variables-domain.html) .
	ConfigurationAliases() interface{}
	SetConfigurationAliases(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The configuration for DASH content.
	DashConfiguration() interface{}
	SetDashConfiguration(val interface{})
	// `AWS::MediaTailor::PlaybackConfiguration.HlsConfiguration`.
	HlsConfiguration() interface{}
	SetHlsConfiguration(val interface{})
	// The configuration for pre-roll ad insertion.
	LivePreRollConfiguration() interface{}
	SetLivePreRollConfiguration(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The configuration for manifest processing rules.
	//
	// Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
	ManifestProcessingRules() interface{}
	SetManifestProcessingRules(val interface{})
	// The identifier for the playback configuration.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Defines the maximum duration of underfilled ad time (in seconds) allowed in an ad break.
	//
	// If the duration of underfilled ad time exceeds the personalization threshold, then the personalization of the ad break is abandoned and the underlying content is shown. This feature applies to *ad replacement* in live and VOD streams, rather than ad insertion, because it relies on an underlying content stream. For more information about ad break behavior, including ad replacement and insertion, see [Ad Behavior in MediaTailor](https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html) .
	PersonalizationThresholdSeconds() *float64
	SetPersonalizationThresholdSeconds(val *float64)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The URL for a high-quality video asset to transcode and use to fill in time that's not used by ads.
	//
	// MediaTailor shows the slate to fill in gaps in media content. Configuring the slate is optional for non-VPAID configurations. For VPAID, the slate is required because MediaTailor provides it in the slots that are designated for dynamic ad content. The slate must be a high-quality asset that contains both audio and video.
	SlateAdUrl() *string
	SetSlateAdUrl(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags to assign to the playback configuration.
	Tags() awscdk.TagManager
	// The name that is used to associate this playback configuration with a custom transcode profile.
	//
	// This overrides the dynamic transcoding defaults of MediaTailor. Use this only if you have already set up custom profiles with the help of AWS Support.
	TranscodeProfileName() *string
	SetTranscodeProfileName(val *string)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// The URL prefix for the parent manifest for the stream, minus the asset ID.
	//
	// The maximum length is 512 characters.
	VideoContentSourceUrl() *string
	SetVideoContentSourceUrl(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnPlaybackConfiguration
type jsiiProxy_CfnPlaybackConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPlaybackConfiguration) AdDecisionServerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adDecisionServerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) AttrDashConfigurationManifestEndpointPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDashConfigurationManifestEndpointPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) AttrHlsConfigurationManifestEndpointPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrHlsConfigurationManifestEndpointPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) AttrPlaybackConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPlaybackConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) AttrPlaybackEndpointPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPlaybackEndpointPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) AttrSessionInitializationEndpointPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSessionInitializationEndpointPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) AvailSuppression() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"availSuppression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) Bumper() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bumper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) CdnConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdnConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) ConfigurationAliases() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurationAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) DashConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dashConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) HlsConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hlsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) LivePreRollConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"livePreRollConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) ManifestProcessingRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manifestProcessingRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) PersonalizationThresholdSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"personalizationThresholdSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) SlateAdUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slateAdUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) TranscodeProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transcodeProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfiguration) VideoContentSourceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"videoContentSourceUrl",
		&returns,
	)
	return returns
}


// Create a new `AWS::MediaTailor::PlaybackConfiguration`.
func NewCfnPlaybackConfiguration(scope constructs.Construct, id *string, props *CfnPlaybackConfigurationProps) CfnPlaybackConfiguration {
	_init_.Initialize()

	if err := validateNewCfnPlaybackConfigurationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPlaybackConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_mediatailor.CfnPlaybackConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MediaTailor::PlaybackConfiguration`.
func NewCfnPlaybackConfiguration_Override(c CfnPlaybackConfiguration, scope constructs.Construct, id *string, props *CfnPlaybackConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_mediatailor.CfnPlaybackConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration)SetAdDecisionServerUrl(val *string) {
	if err := j.validateSetAdDecisionServerUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adDecisionServerUrl",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration)SetAvailSuppression(val interface{}) {
	if err := j.validateSetAvailSuppressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availSuppression",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration)SetBumper(val interface{}) {
	if err := j.validateSetBumperParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bumper",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration)SetCdnConfiguration(val interface{}) {
	if err := j.validateSetCdnConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdnConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration)SetConfigurationAliases(val interface{}) {
	if err := j.validateSetConfigurationAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationAliases",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration)SetDashConfiguration(val interface{}) {
	if err := j.validateSetDashConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dashConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration)SetHlsConfiguration(val interface{}) {
	if err := j.validateSetHlsConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hlsConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration)SetLivePreRollConfiguration(val interface{}) {
	if err := j.validateSetLivePreRollConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"livePreRollConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration)SetManifestProcessingRules(val interface{}) {
	if err := j.validateSetManifestProcessingRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manifestProcessingRules",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration)SetPersonalizationThresholdSeconds(val *float64) {
	_jsii_.Set(
		j,
		"personalizationThresholdSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration)SetSlateAdUrl(val *string) {
	_jsii_.Set(
		j,
		"slateAdUrl",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration)SetTranscodeProfileName(val *string) {
	_jsii_.Set(
		j,
		"transcodeProfileName",
		val,
	)
}

func (j *jsiiProxy_CfnPlaybackConfiguration)SetVideoContentSourceUrl(val *string) {
	if err := j.validateSetVideoContentSourceUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"videoContentSourceUrl",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPlaybackConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPlaybackConfiguration_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediatailor.CfnPlaybackConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPlaybackConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnPlaybackConfiguration_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediatailor.CfnPlaybackConfiguration",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnPlaybackConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPlaybackConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mediatailor.CfnPlaybackConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPlaybackConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_mediatailor.CfnPlaybackConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPlaybackConfiguration) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPlaybackConfiguration) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnPlaybackConfiguration) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPlaybackConfiguration) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPlaybackConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnPlaybackConfiguration) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnPlaybackConfiguration) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPlaybackConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPlaybackConfiguration) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

