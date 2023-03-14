package awsconnectcampaigns


// Contains predictive dialer configuration for an outbound campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   predictiveDialerConfigProperty := &PredictiveDialerConfigProperty{
//   	BandwidthAllocation: jsii.Number(123),
//   }
//
type CfnCampaign_PredictiveDialerConfigProperty struct {
	// Bandwidth allocation for the predictive dialer.
	BandwidthAllocation *float64 `field:"required" json:"bandwidthAllocation" yaml:"bandwidthAllocation"`
}

