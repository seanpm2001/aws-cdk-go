package awscodestarnotifications


// The level of detail to include in the notifications for this resource.
// Experimental.
type DetailType string

const (
	// BASIC will include only the contents of the event as it would appear in AWS CloudWatch.
	// Experimental.
	DetailType_BASIC DetailType = "BASIC"
	// FULL will include any supplemental information provided by AWS CodeStar Notifications and/or the service for the resource for which the notification is created.
	// Experimental.
	DetailType_FULL DetailType = "FULL"
)

