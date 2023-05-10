package awsconfig


// The maximum frequency at which the AWS Config rule runs evaluations.
//
// Example:
//   // https://docs.aws.amazon.com/config/latest/developerguide/access-keys-rotated.html
//   // https://docs.aws.amazon.com/config/latest/developerguide/access-keys-rotated.html
//   config.NewManagedRule(this, jsii.String("AccessKeysRotated"), &ManagedRuleProps{
//   	Identifier: config.ManagedRuleIdentifiers_ACCESS_KEYS_ROTATED(),
//   	InputParameters: map[string]interface{}{
//   		"maxAccessKeyAge": jsii.Number(60),
//   	},
//
//   	// default is 24 hours
//   	MaximumExecutionFrequency: config.MaximumExecutionFrequency_TWELVE_HOURS,
//   })
//
// Experimental.
type MaximumExecutionFrequency string

const (
	// 1 hour.
	// Experimental.
	MaximumExecutionFrequency_ONE_HOUR MaximumExecutionFrequency = "ONE_HOUR"
	// 3 hours.
	// Experimental.
	MaximumExecutionFrequency_THREE_HOURS MaximumExecutionFrequency = "THREE_HOURS"
	// 6 hours.
	// Experimental.
	MaximumExecutionFrequency_SIX_HOURS MaximumExecutionFrequency = "SIX_HOURS"
	// 12 hours.
	// Experimental.
	MaximumExecutionFrequency_TWELVE_HOURS MaximumExecutionFrequency = "TWELVE_HOURS"
	// 24 hours.
	// Experimental.
	MaximumExecutionFrequency_TWENTY_FOUR_HOURS MaximumExecutionFrequency = "TWENTY_FOUR_HOURS"
)

