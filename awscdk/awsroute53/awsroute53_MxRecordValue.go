package awsroute53


// Properties for a MX record value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mxRecordValue := &mxRecordValue{
//   	hostName: jsii.String("hostName"),
//   	priority: jsii.Number(123),
//   }
//
// Experimental.
type MxRecordValue struct {
	// The mail server host name.
	// Experimental.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// The priority.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
}
