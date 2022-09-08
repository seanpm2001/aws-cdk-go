package awss3notifications

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3notifications/internal"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
)

// Use an SQS queue as a bucket notification destination.
//
// Example:
//   var myQueue queue
//
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   bucket.addEventNotification(s3.eventType_OBJECT_REMOVED,
//   s3n.NewSqsDestination(myQueue), &notificationKeyFilter{
//   	prefix: jsii.String("foo/"),
//   	suffix: jsii.String(".jpg"),
//   })
//
// Experimental.
type SqsDestination interface {
	awss3.IBucketNotificationDestination
	// Allows using SQS queues as destinations for bucket notifications.
	//
	// Use `bucket.onEvent(event, queue)` to subscribe.
	// Experimental.
	Bind(_scope awscdk.Construct, bucket awss3.IBucket) *awss3.BucketNotificationDestinationConfig
}

// The jsii proxy struct for SqsDestination
type jsiiProxy_SqsDestination struct {
	internal.Type__awss3IBucketNotificationDestination
}

// Experimental.
func NewSqsDestination(queue awssqs.IQueue) SqsDestination {
	_init_.Initialize()

	if err := validateNewSqsDestinationParameters(queue); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqsDestination{}

	_jsii_.Create(
		"monocdk.aws_s3_notifications.SqsDestination",
		[]interface{}{queue},
		&j,
	)

	return &j
}

// Experimental.
func NewSqsDestination_Override(s SqsDestination, queue awssqs.IQueue) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_s3_notifications.SqsDestination",
		[]interface{}{queue},
		s,
	)
}

func (s *jsiiProxy_SqsDestination) Bind(_scope awscdk.Construct, bucket awss3.IBucket) *awss3.BucketNotificationDestinationConfig {
	if err := s.validateBindParameters(_scope, bucket); err != nil {
		panic(err)
	}
	var returns *awss3.BucketNotificationDestinationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, bucket},
		&returns,
	)

	return returns
}

