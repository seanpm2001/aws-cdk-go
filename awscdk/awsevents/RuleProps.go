package awsevents

import (
	"github.com/aws/constructs-go/constructs/v10"
)

// Properties for defining an EventBridge Rule.
//
// Example:
//   connection := events.NewConnection(this, jsii.String("Connection"), &ConnectionProps{
//   	Authorization: events.Authorization_ApiKey(jsii.String("x-api-key"), awscdk.SecretValue_SecretsManager(jsii.String("ApiSecretName"))),
//   	Description: jsii.String("Connection with API Key x-api-key"),
//   })
//
//   destination := events.NewApiDestination(this, jsii.String("Destination"), &ApiDestinationProps{
//   	Connection: Connection,
//   	Endpoint: jsii.String("https://example.com"),
//   	Description: jsii.String("Calling example.com with API key x-api-key"),
//   })
//
//   rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
//   	Schedule: events.Schedule_Rate(cdk.Duration_Minutes(jsii.Number(1))),
//   	Targets: []iRuleTarget{
//   		targets.NewApiDestination(destination),
//   	},
//   })
//
type RuleProps struct {
	// The scope to use if the source of the rule and its target are in different Stacks (but in the same account & region).
	//
	// This helps dealing with cycles that often arise in these situations.
	CrossStackScope constructs.Construct `field:"optional" json:"crossStackScope" yaml:"crossStackScope"`
	// A description of the rule's purpose.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Additional restrictions for the event to route to the specified target.
	//
	// The method that generates the rule probably imposes some type of event
	// filtering. The filtering implied by what you pass here is added
	// on top of that filtering.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html
	//
	EventPattern *EventPattern `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// A name for the rule.
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// Indicates whether the rule is enabled.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The event bus to associate with this rule.
	EventBus IEventBus `field:"optional" json:"eventBus" yaml:"eventBus"`
	// The schedule or rate (frequency) that determines when EventBridge runs the rule.
	//
	// You must specify this property, the `eventPattern` property, or both.
	//
	// For more information, see Schedule Expression Syntax for
	// Rules in the Amazon EventBridge User Guide.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/scheduled-events.html
	//
	Schedule Schedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Targets to invoke when this rule matches an event.
	//
	// Input will be the full matched event. If you wish to specify custom
	// target input, use `addTarget(target[, inputOptions])`.
	Targets *[]IRuleTarget `field:"optional" json:"targets" yaml:"targets"`
}

