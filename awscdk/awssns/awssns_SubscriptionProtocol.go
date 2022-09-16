package awssns


// The type of subscription, controlling the type of the endpoint parameter.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   var stream deliveryStream
//
//
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//
//   sns.NewSubscription(this, jsii.String("Subscription"), &subscriptionProps{
//   	topic: topic,
//   	endpoint: stream.deliveryStreamArn,
//   	protocol: sns.subscriptionProtocol_FIREHOSE,
//   	subscriptionRoleArn: jsii.String("SAMPLE_ARN"),
//   })
//
// Experimental.
type SubscriptionProtocol string

const (
	// JSON-encoded message is POSTED to an HTTP url.
	// Experimental.
	SubscriptionProtocol_HTTP SubscriptionProtocol = "HTTP"
	// JSON-encoded message is POSTed to an HTTPS url.
	// Experimental.
	SubscriptionProtocol_HTTPS SubscriptionProtocol = "HTTPS"
	// Notifications are sent via email.
	// Experimental.
	SubscriptionProtocol_EMAIL SubscriptionProtocol = "EMAIL"
	// Notifications are JSON-encoded and sent via mail.
	// Experimental.
	SubscriptionProtocol_EMAIL_JSON SubscriptionProtocol = "EMAIL_JSON"
	// Notification is delivered by SMS.
	// Experimental.
	SubscriptionProtocol_SMS SubscriptionProtocol = "SMS"
	// Notifications are enqueued into an SQS queue.
	// Experimental.
	SubscriptionProtocol_SQS SubscriptionProtocol = "SQS"
	// JSON-encoded notifications are sent to a mobile app endpoint.
	// Experimental.
	SubscriptionProtocol_APPLICATION SubscriptionProtocol = "APPLICATION"
	// Notifications trigger a Lambda function.
	// Experimental.
	SubscriptionProtocol_LAMBDA SubscriptionProtocol = "LAMBDA"
	// Notifications put records into a firehose delivery stream.
	// Experimental.
	SubscriptionProtocol_FIREHOSE SubscriptionProtocol = "FIREHOSE"
)
