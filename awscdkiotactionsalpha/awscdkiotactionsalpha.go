// Receipt rule actions for AWS IoT
package awscdkiotactionsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkiotactionsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-cdk-go/awscdkiotactionsalpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdkiotalpha/v2"
	"github.com/aws/aws-cdk-go/awscdkioteventsalpha/v2"
	"github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha/v2"
)

// The action to send data to Amazon CloudWatch Logs.
//
// Example:
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//
//
//   logGroup := logs.NewLogGroup(this, jsii.String("MyLogGroup"))
//
//   iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp FROM 'device/+/data'")),
//   	errorAction: actions.NewCloudWatchLogsAction(logGroup),
//   })
//
// Experimental.
type CloudWatchLogsAction interface {
	awscdkiotalpha.IAction
}

// The jsii proxy struct for CloudWatchLogsAction
type jsiiProxy_CloudWatchLogsAction struct {
	internal.Type__awscdkiotalphaIAction
}

// Experimental.
func NewCloudWatchLogsAction(logGroup awslogs.ILogGroup, props *CloudWatchLogsActionProps) CloudWatchLogsAction {
	_init_.Initialize()

	j := jsiiProxy_CloudWatchLogsAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.CloudWatchLogsAction",
		[]interface{}{logGroup, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudWatchLogsAction_Override(c CloudWatchLogsAction, logGroup awslogs.ILogGroup, props *CloudWatchLogsActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.CloudWatchLogsAction",
		[]interface{}{logGroup, props},
		c,
	)
}

// Configuration properties of an action for CloudWatch Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import iot_actions_alpha "github.com/aws/aws-cdk-go/awscdkiotactionsalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   cloudWatchLogsActionProps := &cloudWatchLogsActionProps{
//   	role: role,
//   }
//
// Experimental.
type CloudWatchLogsActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

// The action to capture an Amazon CloudWatch metric.
//
// Example:
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, namespace, unit, value, timestamp FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewCloudWatchPutMetricAction(&cloudWatchPutMetricActionProps{
//   			metricName: jsii.String("${topic(2)}"),
//   			metricNamespace: jsii.String("${namespace}"),
//   			metricUnit: jsii.String("${unit}"),
//   			metricValue: jsii.String("${value}"),
//   			metricTimestamp: jsii.String("${timestamp}"),
//   		}),
//   	},
//   })
//
// Experimental.
type CloudWatchPutMetricAction interface {
	awscdkiotalpha.IAction
}

// The jsii proxy struct for CloudWatchPutMetricAction
type jsiiProxy_CloudWatchPutMetricAction struct {
	internal.Type__awscdkiotalphaIAction
}

// Experimental.
func NewCloudWatchPutMetricAction(props *CloudWatchPutMetricActionProps) CloudWatchPutMetricAction {
	_init_.Initialize()

	j := jsiiProxy_CloudWatchPutMetricAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.CloudWatchPutMetricAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudWatchPutMetricAction_Override(c CloudWatchPutMetricAction, props *CloudWatchPutMetricActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.CloudWatchPutMetricAction",
		[]interface{}{props},
		c,
	)
}

// Configuration properties of an action for CloudWatch metric.
//
// Example:
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, namespace, unit, value, timestamp FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewCloudWatchPutMetricAction(&cloudWatchPutMetricActionProps{
//   			metricName: jsii.String("${topic(2)}"),
//   			metricNamespace: jsii.String("${namespace}"),
//   			metricUnit: jsii.String("${unit}"),
//   			metricValue: jsii.String("${value}"),
//   			metricTimestamp: jsii.String("${timestamp}"),
//   		}),
//   	},
//   })
//
// Experimental.
type CloudWatchPutMetricActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The CloudWatch metric name.
	//
	// Supports substitution templates.
	// See: https://docs.aws.amazon.com/iot/latest/developerguide/iot-substitution-templates.html
	//
	// Experimental.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The CloudWatch metric namespace name.
	//
	// Supports substitution templates.
	// See: https://docs.aws.amazon.com/iot/latest/developerguide/iot-substitution-templates.html
	//
	// Experimental.
	MetricNamespace *string `field:"required" json:"metricNamespace" yaml:"metricNamespace"`
	// The metric unit supported by CloudWatch.
	//
	// Supports substitution templates.
	// See: https://docs.aws.amazon.com/iot/latest/developerguide/iot-substitution-templates.html
	//
	// Experimental.
	MetricUnit *string `field:"required" json:"metricUnit" yaml:"metricUnit"`
	// A string that contains the CloudWatch metric value.
	//
	// Supports substitution templates.
	// See: https://docs.aws.amazon.com/iot/latest/developerguide/iot-substitution-templates.html
	//
	// Experimental.
	MetricValue *string `field:"required" json:"metricValue" yaml:"metricValue"`
	// A string that contains the timestamp, expressed in seconds in Unix epoch time.
	//
	// Supports substitution templates.
	// See: https://docs.aws.amazon.com/iot/latest/developerguide/iot-substitution-templates.html
	//
	// Experimental.
	MetricTimestamp *string `field:"optional" json:"metricTimestamp" yaml:"metricTimestamp"`
}

// The action to change the state of an Amazon CloudWatch alarm.
//
// Example:
//   import cloudwatch "github.com/aws/aws-cdk-go/awscdk"
//
//
//   metric := cloudwatch.NewMetric(&metricProps{
//   	namespace: jsii.String("MyNamespace"),
//   	metricName: jsii.String("MyMetric"),
//   	dimensions: map[string]interface{}{
//   		"MyDimension": jsii.String("MyDimensionValue"),
//   	},
//   })
//   alarm := cloudwatch.NewAlarm(this, jsii.String("MyAlarm"), &alarmProps{
//   	metric: metric,
//   	threshold: jsii.Number(100),
//   	evaluationPeriods: jsii.Number(3),
//   	datapointsToAlarm: jsii.Number(2),
//   })
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewCloudWatchSetAlarmStateAction(alarm, &cloudWatchSetAlarmStateActionProps{
//   			reason: jsii.String("AWS Iot Rule action is triggered"),
//   			alarmStateToSet: cloudwatch.alarmState_ALARM,
//   		}),
//   	},
//   })
//
// Experimental.
type CloudWatchSetAlarmStateAction interface {
	awscdkiotalpha.IAction
}

// The jsii proxy struct for CloudWatchSetAlarmStateAction
type jsiiProxy_CloudWatchSetAlarmStateAction struct {
	internal.Type__awscdkiotalphaIAction
}

// Experimental.
func NewCloudWatchSetAlarmStateAction(alarm awscloudwatch.IAlarm, props *CloudWatchSetAlarmStateActionProps) CloudWatchSetAlarmStateAction {
	_init_.Initialize()

	j := jsiiProxy_CloudWatchSetAlarmStateAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.CloudWatchSetAlarmStateAction",
		[]interface{}{alarm, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudWatchSetAlarmStateAction_Override(c CloudWatchSetAlarmStateAction, alarm awscloudwatch.IAlarm, props *CloudWatchSetAlarmStateActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.CloudWatchSetAlarmStateAction",
		[]interface{}{alarm, props},
		c,
	)
}

// Configuration properties of an action for CloudWatch alarm.
//
// Example:
//   import cloudwatch "github.com/aws/aws-cdk-go/awscdk"
//
//
//   metric := cloudwatch.NewMetric(&metricProps{
//   	namespace: jsii.String("MyNamespace"),
//   	metricName: jsii.String("MyMetric"),
//   	dimensions: map[string]interface{}{
//   		"MyDimension": jsii.String("MyDimensionValue"),
//   	},
//   })
//   alarm := cloudwatch.NewAlarm(this, jsii.String("MyAlarm"), &alarmProps{
//   	metric: metric,
//   	threshold: jsii.Number(100),
//   	evaluationPeriods: jsii.Number(3),
//   	datapointsToAlarm: jsii.Number(2),
//   })
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewCloudWatchSetAlarmStateAction(alarm, &cloudWatchSetAlarmStateActionProps{
//   			reason: jsii.String("AWS Iot Rule action is triggered"),
//   			alarmStateToSet: cloudwatch.alarmState_ALARM,
//   		}),
//   	},
//   })
//
// Experimental.
type CloudWatchSetAlarmStateActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The value of the alarm state to set.
	// Experimental.
	AlarmStateToSet awscloudwatch.AlarmState `field:"required" json:"alarmStateToSet" yaml:"alarmStateToSet"`
	// The reason for the alarm change.
	// Experimental.
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
}

// Common properties shared by Actions it access to AWS service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import iot_actions_alpha "github.com/aws/aws-cdk-go/awscdkiotactionsalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   commonActionProps := &commonActionProps{
//   	role: role,
//   }
//
// Experimental.
type CommonActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

// The action to put the record from an MQTT message to the DynamoDB table.
//
// Example:
//   import dynamodb "github.com/aws/aws-cdk-go/awscdk"
//
//   var table table
//
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewDynamoDBv2PutItemAction(table),
//   	},
//   })
//
// Experimental.
type DynamoDBv2PutItemAction interface {
	awscdkiotalpha.IAction
}

// The jsii proxy struct for DynamoDBv2PutItemAction
type jsiiProxy_DynamoDBv2PutItemAction struct {
	internal.Type__awscdkiotalphaIAction
}

// Experimental.
func NewDynamoDBv2PutItemAction(table awsdynamodb.ITable, props *DynamoDBv2PutItemActionProps) DynamoDBv2PutItemAction {
	_init_.Initialize()

	j := jsiiProxy_DynamoDBv2PutItemAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.DynamoDBv2PutItemAction",
		[]interface{}{table, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDynamoDBv2PutItemAction_Override(d DynamoDBv2PutItemAction, table awsdynamodb.ITable, props *DynamoDBv2PutItemActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.DynamoDBv2PutItemAction",
		[]interface{}{table, props},
		d,
	)
}

// Configuration properties of an action for the dynamodb table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import iot_actions_alpha "github.com/aws/aws-cdk-go/awscdkiotactionsalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   dynamoDBv2PutItemActionProps := &dynamoDBv2PutItemActionProps{
//   	role: role,
//   }
//
// Experimental.
type DynamoDBv2PutItemActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

// The action to put the record from an MQTT message to the Kinesis Data Firehose stream.
//
// Example:
//   import firehose "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha"
//   import destinations "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha"
//
//
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   stream := firehose.NewDeliveryStream(this, jsii.String("MyStream"), &deliveryStreamProps{
//   	destinations: []iDestination{
//   		destinations.NewS3Bucket(bucket),
//   	},
//   })
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewFirehosePutRecordAction(stream, &firehosePutRecordActionProps{
//   			batchMode: jsii.Boolean(true),
//   			recordSeparator: actions.firehoseRecordSeparator_NEWLINE,
//   		}),
//   	},
//   })
//
// Experimental.
type FirehosePutRecordAction interface {
	awscdkiotalpha.IAction
}

// The jsii proxy struct for FirehosePutRecordAction
type jsiiProxy_FirehosePutRecordAction struct {
	internal.Type__awscdkiotalphaIAction
}

// Experimental.
func NewFirehosePutRecordAction(stream awscdkkinesisfirehosealpha.IDeliveryStream, props *FirehosePutRecordActionProps) FirehosePutRecordAction {
	_init_.Initialize()

	j := jsiiProxy_FirehosePutRecordAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.FirehosePutRecordAction",
		[]interface{}{stream, props},
		&j,
	)

	return &j
}

// Experimental.
func NewFirehosePutRecordAction_Override(f FirehosePutRecordAction, stream awscdkkinesisfirehosealpha.IDeliveryStream, props *FirehosePutRecordActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.FirehosePutRecordAction",
		[]interface{}{stream, props},
		f,
	)
}

// Configuration properties of an action for the Kinesis Data Firehose stream.
//
// Example:
//   import firehose "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha"
//   import destinations "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha"
//
//
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   stream := firehose.NewDeliveryStream(this, jsii.String("MyStream"), &deliveryStreamProps{
//   	destinations: []iDestination{
//   		destinations.NewS3Bucket(bucket),
//   	},
//   })
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewFirehosePutRecordAction(stream, &firehosePutRecordActionProps{
//   			batchMode: jsii.Boolean(true),
//   			recordSeparator: actions.firehoseRecordSeparator_NEWLINE,
//   		}),
//   	},
//   })
//
// Experimental.
type FirehosePutRecordActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Whether to deliver the Kinesis Data Firehose stream as a batch by using `PutRecordBatch`.
	//
	// When batchMode is true and the rule's SQL statement evaluates to an Array, each Array
	// element forms one record in the PutRecordBatch request. The resulting array can't have
	// more than 500 records.
	// Experimental.
	BatchMode *bool `field:"optional" json:"batchMode" yaml:"batchMode"`
	// A character separator that will be used to separate records written to the Kinesis Data Firehose stream.
	// Experimental.
	RecordSeparator FirehoseRecordSeparator `field:"optional" json:"recordSeparator" yaml:"recordSeparator"`
}

// Record Separator to be used to separate records.
//
// Example:
//   import firehose "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha"
//   import destinations "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha"
//
//
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   stream := firehose.NewDeliveryStream(this, jsii.String("MyStream"), &deliveryStreamProps{
//   	destinations: []iDestination{
//   		destinations.NewS3Bucket(bucket),
//   	},
//   })
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewFirehosePutRecordAction(stream, &firehosePutRecordActionProps{
//   			batchMode: jsii.Boolean(true),
//   			recordSeparator: actions.firehoseRecordSeparator_NEWLINE,
//   		}),
//   	},
//   })
//
// Experimental.
type FirehoseRecordSeparator string

const (
	// Separate by a new line.
	// Experimental.
	FirehoseRecordSeparator_NEWLINE FirehoseRecordSeparator = "NEWLINE"
	// Separate by a tab.
	// Experimental.
	FirehoseRecordSeparator_TAB FirehoseRecordSeparator = "TAB"
	// Separate by a windows new line.
	// Experimental.
	FirehoseRecordSeparator_WINDOWS_NEWLINE FirehoseRecordSeparator = "WINDOWS_NEWLINE"
	// Separate by a commma.
	// Experimental.
	FirehoseRecordSeparator_COMMA FirehoseRecordSeparator = "COMMA"
)

// The action to put the message from an MQTT message to the IoT Events input.
//
// Example:
//   import iotevents "github.com/aws/aws-cdk-go/awscdkioteventsalpha"
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//
//   var role iRole
//
//
//   input := iotevents.NewInput(this, jsii.String("MyInput"), &inputProps{
//   	attributeJsonPaths: []*string{
//   		jsii.String("payload.temperature"),
//   		jsii.String("payload.transactionId"),
//   	},
//   })
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewIotEventsPutMessageAction(input, &iotEventsPutMessageActionProps{
//   			batchMode: jsii.Boolean(true),
//   			 // optional property, default is 'false'
//   			messageId: jsii.String("${payload.transactionId}"),
//   			 // optional property, default is a new UUID
//   			role: role,
//   		}),
//   	},
//   })
//
// Experimental.
type IotEventsPutMessageAction interface {
	awscdkiotalpha.IAction
}

// The jsii proxy struct for IotEventsPutMessageAction
type jsiiProxy_IotEventsPutMessageAction struct {
	internal.Type__awscdkiotalphaIAction
}

// Experimental.
func NewIotEventsPutMessageAction(input awscdkioteventsalpha.IInput, props *IotEventsPutMessageActionProps) IotEventsPutMessageAction {
	_init_.Initialize()

	j := jsiiProxy_IotEventsPutMessageAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.IotEventsPutMessageAction",
		[]interface{}{input, props},
		&j,
	)

	return &j
}

// Experimental.
func NewIotEventsPutMessageAction_Override(i IotEventsPutMessageAction, input awscdkioteventsalpha.IInput, props *IotEventsPutMessageActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.IotEventsPutMessageAction",
		[]interface{}{input, props},
		i,
	)
}

// Configuration properties of an action for the IoT Events.
//
// Example:
//   import iotevents "github.com/aws/aws-cdk-go/awscdkioteventsalpha"
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//
//   var role iRole
//
//
//   input := iotevents.NewInput(this, jsii.String("MyInput"), &inputProps{
//   	attributeJsonPaths: []*string{
//   		jsii.String("payload.temperature"),
//   		jsii.String("payload.transactionId"),
//   	},
//   })
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewIotEventsPutMessageAction(input, &iotEventsPutMessageActionProps{
//   			batchMode: jsii.Boolean(true),
//   			 // optional property, default is 'false'
//   			messageId: jsii.String("${payload.transactionId}"),
//   			 // optional property, default is a new UUID
//   			role: role,
//   		}),
//   	},
//   })
//
// Experimental.
type IotEventsPutMessageActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Whether to process the event actions as a batch.
	//
	// When batchMode is true, you can't specify a messageId.
	//
	// When batchMode is true and the rule SQL statement evaluates to an Array,
	// each Array element is treated as a separate message when Events by calling BatchPutMessage.
	// The resulting array can't have more than 10 messages.
	// Experimental.
	BatchMode *bool `field:"optional" json:"batchMode" yaml:"batchMode"`
	// The ID of the message.
	//
	// When batchMode is true, you can't specify a messageId--a new UUID value will be assigned.
	// Assign a value to this property to ensure that only one input (message) with a given messageId will be processed by an AWS IoT Events detector.
	// Experimental.
	MessageId *string `field:"optional" json:"messageId" yaml:"messageId"`
}

// The action to put the record from an MQTT message to republish another MQTT topic.
//
// Example:
//   iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp, temperature FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewIotRepublishMqttAction(jsii.String("${topic()}/republish"), &iotRepublishMqttActionProps{
//   			qualityOfService: actions.mqttQualityOfService_AT_LEAST_ONCE,
//   		}),
//   	},
//   })
//
// Experimental.
type IotRepublishMqttAction interface {
	awscdkiotalpha.IAction
}

// The jsii proxy struct for IotRepublishMqttAction
type jsiiProxy_IotRepublishMqttAction struct {
	internal.Type__awscdkiotalphaIAction
}

// Experimental.
func NewIotRepublishMqttAction(topic *string, props *IotRepublishMqttActionProps) IotRepublishMqttAction {
	_init_.Initialize()

	j := jsiiProxy_IotRepublishMqttAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.IotRepublishMqttAction",
		[]interface{}{topic, props},
		&j,
	)

	return &j
}

// Experimental.
func NewIotRepublishMqttAction_Override(i IotRepublishMqttAction, topic *string, props *IotRepublishMqttActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.IotRepublishMqttAction",
		[]interface{}{topic, props},
		i,
	)
}

// Configuration properties of an action to republish MQTT messages.
//
// Example:
//   iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp, temperature FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewIotRepublishMqttAction(jsii.String("${topic()}/republish"), &iotRepublishMqttActionProps{
//   			qualityOfService: actions.mqttQualityOfService_AT_LEAST_ONCE,
//   		}),
//   	},
//   })
//
// Experimental.
type IotRepublishMqttActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The Quality of Service (QoS) level to use when republishing messages.
	// See: https://docs.aws.amazon.com/iot/latest/developerguide/mqtt.html#mqtt-qos
	//
	// Experimental.
	QualityOfService MqttQualityOfService `field:"optional" json:"qualityOfService" yaml:"qualityOfService"`
}

// The action to put the record from an MQTT message to the Kinesis Data stream.
//
// Example:
//   import kinesis "github.com/aws/aws-cdk-go/awscdk"
//
//
//   stream := kinesis.NewStream(this, jsii.String("MyStream"))
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewKinesisPutRecordAction(stream, &kinesisPutRecordActionProps{
//   			partitionKey: jsii.String("${newuuid()}"),
//   		}),
//   	},
//   })
//
// Experimental.
type KinesisPutRecordAction interface {
	awscdkiotalpha.IAction
}

// The jsii proxy struct for KinesisPutRecordAction
type jsiiProxy_KinesisPutRecordAction struct {
	internal.Type__awscdkiotalphaIAction
}

// Experimental.
func NewKinesisPutRecordAction(stream awskinesis.IStream, props *KinesisPutRecordActionProps) KinesisPutRecordAction {
	_init_.Initialize()

	j := jsiiProxy_KinesisPutRecordAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.KinesisPutRecordAction",
		[]interface{}{stream, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKinesisPutRecordAction_Override(k KinesisPutRecordAction, stream awskinesis.IStream, props *KinesisPutRecordActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.KinesisPutRecordAction",
		[]interface{}{stream, props},
		k,
	)
}

// Configuration properties of an action for the Kinesis Data stream.
//
// Example:
//   import kinesis "github.com/aws/aws-cdk-go/awscdk"
//
//
//   stream := kinesis.NewStream(this, jsii.String("MyStream"))
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewKinesisPutRecordAction(stream, &kinesisPutRecordActionProps{
//   			partitionKey: jsii.String("${newuuid()}"),
//   		}),
//   	},
//   })
//
// Experimental.
type KinesisPutRecordActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The partition key used to determine to which shard the data is written.
	//
	// The partition key is usually composed of an expression (for example, ${topic()} or ${timestamp()}).
	// See: https://docs.aws.amazon.com/kinesis/latest/APIReference/API_PutRecord.html#API_PutRecord_RequestParameters
	//
	// Experimental.
	PartitionKey *string `field:"required" json:"partitionKey" yaml:"partitionKey"`
}

// The action to invoke an AWS Lambda function, passing in an MQTT message.
//
// Example:
//   func := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromInline(jsii.String("\n    exports.handler = (event) => {\n      console.log(\"It is test for lambda action of AWS IoT Rule.\", event);\n    };")),
//   })
//
//   iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp, temperature FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewLambdaFunctionAction(func),
//   	},
//   })
//
// Experimental.
type LambdaFunctionAction interface {
	awscdkiotalpha.IAction
}

// The jsii proxy struct for LambdaFunctionAction
type jsiiProxy_LambdaFunctionAction struct {
	internal.Type__awscdkiotalphaIAction
}

// Experimental.
func NewLambdaFunctionAction(func_ awslambda.IFunction) LambdaFunctionAction {
	_init_.Initialize()

	j := jsiiProxy_LambdaFunctionAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.LambdaFunctionAction",
		[]interface{}{func_},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaFunctionAction_Override(l LambdaFunctionAction, func_ awslambda.IFunction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.LambdaFunctionAction",
		[]interface{}{func_},
		l,
	)
}

// MQTT Quality of Service (QoS) indicates the level of assurance for delivery of an MQTT Message.
//
// Example:
//   iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp, temperature FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewIotRepublishMqttAction(jsii.String("${topic()}/republish"), &iotRepublishMqttActionProps{
//   			qualityOfService: actions.mqttQualityOfService_AT_LEAST_ONCE,
//   		}),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/iot/latest/developerguide/mqtt.html#mqtt-qos
//
// Experimental.
type MqttQualityOfService string

const (
	// QoS level 0.
	//
	// Sent zero or more times.
	// This level should be used for messages that are sent over reliable communication links or that can be missed without a problem.
	// Experimental.
	MqttQualityOfService_ZERO_OR_MORE_TIMES MqttQualityOfService = "ZERO_OR_MORE_TIMES"
	// QoS level 1.
	//
	// Sent at least one time, and then repeatedly until a PUBACK response is received.
	// The message is not considered complete until the sender receives a PUBACK response to indicate successful delivery.
	// Experimental.
	MqttQualityOfService_AT_LEAST_ONCE MqttQualityOfService = "AT_LEAST_ONCE"
)

// The action to write the data from an MQTT message to an Amazon S3 bucket.
//
// Example:
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//
//   iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewS3PutObjectAction(bucket, &s3PutObjectActionProps{
//   			key: jsii.String("${year}/${month}/${day}/${topic(2)}"),
//   		}),
//   	},
//   })
//
// Experimental.
type S3PutObjectAction interface {
	awscdkiotalpha.IAction
}

// The jsii proxy struct for S3PutObjectAction
type jsiiProxy_S3PutObjectAction struct {
	internal.Type__awscdkiotalphaIAction
}

// Experimental.
func NewS3PutObjectAction(bucket awss3.IBucket, props *S3PutObjectActionProps) S3PutObjectAction {
	_init_.Initialize()

	j := jsiiProxy_S3PutObjectAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.S3PutObjectAction",
		[]interface{}{bucket, props},
		&j,
	)

	return &j
}

// Experimental.
func NewS3PutObjectAction_Override(s S3PutObjectAction, bucket awss3.IBucket, props *S3PutObjectActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.S3PutObjectAction",
		[]interface{}{bucket, props},
		s,
	)
}

// Configuration properties of an action for s3.
//
// Example:
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//
//   iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewS3PutObjectAction(bucket, &s3PutObjectActionProps{
//   			key: jsii.String("${year}/${month}/${day}/${topic(2)}"),
//   		}),
//   	},
//   })
//
// Experimental.
type S3PutObjectActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The Amazon S3 canned ACL that controls access to the object identified by the object key.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#canned-acl
	//
	// Experimental.
	AccessControl awss3.BucketAccessControl `field:"optional" json:"accessControl" yaml:"accessControl"`
	// The path to the file where the data is written.
	//
	// Supports substitution templates.
	// See: https://docs.aws.amazon.com/iot/latest/developerguide/iot-substitution-templates.html
	//
	// Experimental.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

// SNS topic action message format options.
//
// Example:
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//
//
//   topic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewSnsTopicAction(topic, &snsTopicActionProps{
//   			messageFormat: actions.snsActionMessageFormat_JSON,
//   		}),
//   	},
//   })
//
// Experimental.
type SnsActionMessageFormat string

const (
	// RAW message format.
	// Experimental.
	SnsActionMessageFormat_RAW SnsActionMessageFormat = "RAW"
	// JSON message format.
	// Experimental.
	SnsActionMessageFormat_JSON SnsActionMessageFormat = "JSON"
)

// The action to write the data from an MQTT message to an Amazon SNS topic.
//
// Example:
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//
//
//   topic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewSnsTopicAction(topic, &snsTopicActionProps{
//   			messageFormat: actions.snsActionMessageFormat_JSON,
//   		}),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/iot/latest/developerguide/sns-rule-action.html
//
// Experimental.
type SnsTopicAction interface {
	awscdkiotalpha.IAction
}

// The jsii proxy struct for SnsTopicAction
type jsiiProxy_SnsTopicAction struct {
	internal.Type__awscdkiotalphaIAction
}

// Experimental.
func NewSnsTopicAction(topic awssns.ITopic, props *SnsTopicActionProps) SnsTopicAction {
	_init_.Initialize()

	j := jsiiProxy_SnsTopicAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.SnsTopicAction",
		[]interface{}{topic, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSnsTopicAction_Override(s SnsTopicAction, topic awssns.ITopic, props *SnsTopicActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.SnsTopicAction",
		[]interface{}{topic, props},
		s,
	)
}

// Configuration options for the SNS topic action.
//
// Example:
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//
//
//   topic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewSnsTopicAction(topic, &snsTopicActionProps{
//   			messageFormat: actions.snsActionMessageFormat_JSON,
//   		}),
//   	},
//   })
//
// Experimental.
type SnsTopicActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The message format of the message to publish.
	//
	// SNS uses this setting to determine if the payload should be parsed and relevant platform-specific bits of the payload should be extracted.
	// See: https://docs.aws.amazon.com/sns/latest/dg/sns-message-and-json-formats.html
	//
	// Experimental.
	MessageFormat SnsActionMessageFormat `field:"optional" json:"messageFormat" yaml:"messageFormat"`
}

// The action to write the data from an MQTT message to an Amazon SQS queue.
//
// Example:
//   import sqs "github.com/aws/aws-cdk-go/awscdk"
//
//
//   queue := sqs.NewQueue(this, jsii.String("MyQueue"))
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewSqsQueueAction(queue, &sqsQueueActionProps{
//   			useBase64: jsii.Boolean(true),
//   		}),
//   	},
//   })
//
// Experimental.
type SqsQueueAction interface {
	awscdkiotalpha.IAction
}

// The jsii proxy struct for SqsQueueAction
type jsiiProxy_SqsQueueAction struct {
	internal.Type__awscdkiotalphaIAction
}

// Experimental.
func NewSqsQueueAction(queue awssqs.IQueue, props *SqsQueueActionProps) SqsQueueAction {
	_init_.Initialize()

	j := jsiiProxy_SqsQueueAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.SqsQueueAction",
		[]interface{}{queue, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSqsQueueAction_Override(s SqsQueueAction, queue awssqs.IQueue, props *SqsQueueActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.SqsQueueAction",
		[]interface{}{queue, props},
		s,
	)
}

// Configuration properties of an action for SQS.
//
// Example:
//   import sqs "github.com/aws/aws-cdk-go/awscdk"
//
//
//   queue := sqs.NewQueue(this, jsii.String("MyQueue"))
//
//   topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewSqsQueueAction(queue, &sqsQueueActionProps{
//   			useBase64: jsii.Boolean(true),
//   		}),
//   	},
//   })
//
// Experimental.
type SqsQueueActionProps struct {
	// The IAM role that allows access to AWS service.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Specifies whether to use Base64 encoding.
	// Experimental.
	UseBase64 *bool `field:"optional" json:"useBase64" yaml:"useBase64"`
}
