# AWS IoT Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

AWS IoT Core lets you connect billions of IoT devices and route trillions of
messages to AWS services without managing infrastructure.

## `TopicRule`

Create a topic rule that give your devices the ability to interact with AWS services.
You can create a topic rule with an action that invoke the Lambda action as following:

```go
func := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
	runtime: lambda.runtime_NODEJS_14_X(),
	handler: jsii.String("index.handler"),
	code: lambda.code.fromInline(jsii.String("\n    exports.handler = (event) => {\n      console.log(\"It is test for lambda action of AWS IoT Rule.\", event);\n    };")),
})

iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
	topicRuleName: jsii.String("MyTopicRule"),
	 // optional
	description: jsii.String("invokes the lambda function"),
	 // optional
	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp FROM 'device/+/data'")),
	actions: []iAction{
		actions.NewLambdaFunctionAction(func),
	},
})
```

Or, you can add an action after constructing the `TopicRule` instance as following:

```go
var func function


topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp FROM 'device/+/data'")),
})
topicRule.addAction(actions.NewLambdaFunctionAction(func))
```

You can also supply `errorAction` as following,
and the IoT Rule will trigger it if a rule's action is unable to perform:

```go
import logs "github.com/aws/aws-cdk-go/awscdk"


logGroup := logs.NewLogGroup(this, jsii.String("MyLogGroup"))

iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp FROM 'device/+/data'")),
	errorAction: actions.NewCloudWatchLogsAction(logGroup),
})
```

If you wanna make the topic rule disable, add property `enabled: false` as following:

```go
iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp FROM 'device/+/data'")),
	enabled: jsii.Boolean(false),
})
```

See also [@aws-cdk/aws-iot-actions](https://docs.aws.amazon.com/cdk/api/latest/docs/aws-iot-actions-readme.html) for other actions.