package awsses

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// An event destination.
//
// Example:
//   var myConfigurationSet configurationSet
//   var myTopic topic
//
//
//   myConfigurationSet.AddEventDestination(jsii.String("ToSns"), &ConfigurationSetEventDestinationOptions{
//   	Destination: ses.EventDestination_SnsTopic(myTopic),
//   })
//
type EventDestination interface {
	// A list of CloudWatch dimensions upon which to categorize your emails.
	Dimensions() *[]*CloudWatchDimension
	// A SNS topic to use as event destination.
	Topic() awssns.ITopic
}

// The jsii proxy struct for EventDestination
type jsiiProxy_EventDestination struct {
	_ byte // padding
}

func (j *jsiiProxy_EventDestination) Dimensions() *[]*CloudWatchDimension {
	var returns *[]*CloudWatchDimension
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventDestination) Topic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}


func NewEventDestination_Override(e EventDestination) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.EventDestination",
		nil, // no parameters
		e,
	)
}

// Use CloudWatch dimensions as event destination.
func EventDestination_CloudWatchDimensions(dimensions *[]*CloudWatchDimension) EventDestination {
	_init_.Initialize()

	if err := validateEventDestination_CloudWatchDimensionsParameters(dimensions); err != nil {
		panic(err)
	}
	var returns EventDestination

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.EventDestination",
		"cloudWatchDimensions",
		[]interface{}{dimensions},
		&returns,
	)

	return returns
}

// Use a SNS topic as event destination.
func EventDestination_SnsTopic(topic awssns.ITopic) EventDestination {
	_init_.Initialize()

	if err := validateEventDestination_SnsTopicParameters(topic); err != nil {
		panic(err)
	}
	var returns EventDestination

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.EventDestination",
		"snsTopic",
		[]interface{}{topic},
		&returns,
	)

	return returns
}
