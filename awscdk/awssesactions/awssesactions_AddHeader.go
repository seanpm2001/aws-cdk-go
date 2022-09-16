package awssesactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsses"
	"github.com/aws/aws-cdk-go/awscdk/awssesactions/internal"
)

// Adds a header to the received email.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//   import actions "github.com/aws/aws-cdk-go/awscdk"
//
//
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//
//   ses.NewReceiptRuleSet(this, jsii.String("RuleSet"), &receiptRuleSetProps{
//   	rules: []receiptRuleOptions{
//   		&receiptRuleOptions{
//   			recipients: []*string{
//   				jsii.String("hello@aws.com"),
//   			},
//   			actions: []iReceiptRuleAction{
//   				actions.NewAddHeader(&addHeaderProps{
//   					name: jsii.String("X-Special-Header"),
//   					value: jsii.String("aws"),
//   				}),
//   				actions.NewS3(&s3Props{
//   					bucket: bucket,
//   					objectKeyPrefix: jsii.String("emails/"),
//   					topic: topic,
//   				}),
//   			},
//   		},
//   		&receiptRuleOptions{
//   			recipients: []*string{
//   				jsii.String("aws.com"),
//   			},
//   			actions: []*iReceiptRuleAction{
//   				actions.NewSns(&snsProps{
//   					topic: topic,
//   				}),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type AddHeader interface {
	awsses.IReceiptRuleAction
	// Returns the receipt rule action specification.
	// Experimental.
	Bind(_rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig
}

// The jsii proxy struct for AddHeader
type jsiiProxy_AddHeader struct {
	internal.Type__awssesIReceiptRuleAction
}

// Experimental.
func NewAddHeader(props *AddHeaderProps) AddHeader {
	_init_.Initialize()

	if err := validateNewAddHeaderParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AddHeader{}

	_jsii_.Create(
		"monocdk.aws_ses_actions.AddHeader",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewAddHeader_Override(a AddHeader, props *AddHeaderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ses_actions.AddHeader",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_AddHeader) Bind(_rule awsses.IReceiptRule) *awsses.ReceiptRuleActionConfig {
	if err := a.validateBindParameters(_rule); err != nil {
		panic(err)
	}
	var returns *awsses.ReceiptRuleActionConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_rule},
		&returns,
	)

	return returns
}
