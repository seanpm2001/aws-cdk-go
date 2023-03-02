package awssns

import (
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
)

// Properties for creating a new subscription.
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
type SubscriptionProps struct {
	// The subscription endpoint.
	//
	// The meaning of this value depends on the value for 'protocol'.
	// Experimental.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// What type of subscription to add.
	// Experimental.
	Protocol SubscriptionProtocol `field:"required" json:"protocol" yaml:"protocol"`
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The filter policy.
	// Experimental.
	FilterPolicy *map[string]SubscriptionFilter `field:"optional" json:"filterPolicy" yaml:"filterPolicy"`
	// true if raw message delivery is enabled for the subscription.
	//
	// Raw messages are free of JSON formatting and can be
	// sent to HTTP/S and Amazon SQS endpoints. For more information, see GetSubscriptionAttributes in the Amazon Simple
	// Notification Service API Reference.
	// Experimental.
	RawMessageDelivery *bool `field:"optional" json:"rawMessageDelivery" yaml:"rawMessageDelivery"`
	// The region where the topic resides, in the case of cross-region subscriptions.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Arn of role allowing access to firehose delivery stream.
	//
	// Required for a firehose subscription protocol.
	// Experimental.
	SubscriptionRoleArn *string `field:"optional" json:"subscriptionRoleArn" yaml:"subscriptionRoleArn"`
	// The topic to subscribe to.
	// Experimental.
	Topic ITopic `field:"required" json:"topic" yaml:"topic"`
}

