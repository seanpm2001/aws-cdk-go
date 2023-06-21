package awsssmcontacts


// Information about the days of the week that the on-call rotation coverage includes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   shiftCoverageProperty := &ShiftCoverageProperty{
//   	CoverageTimes: []interface{}{
//   		&CoverageTimeProperty{
//   			EndTime: jsii.String("endTime"),
//   			StartTime: jsii.String("startTime"),
//   		},
//   	},
//   	DayOfWeek: jsii.String("dayOfWeek"),
//   }
//
type CfnRotation_ShiftCoverageProperty struct {
	// The start and end times of the shift.
	CoverageTimes interface{} `field:"required" json:"coverageTimes" yaml:"coverageTimes"`
	// A list of days on which the schedule is active.
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
}
