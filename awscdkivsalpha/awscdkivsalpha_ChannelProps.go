// The CDK Construct Library for AWS::IVS
package awscdkivsalpha


// Properties for creating a new Channel.
//
// Example:
//   myChannel := ivs.NewChannel(this, jsii.String("Channel"), &channelProps{
//   	authorized: jsii.Boolean(true),
//   })
//
// Experimental.
type ChannelProps struct {
	// Whether the channel is authorized.
	//
	// If you wish to make an authorized channel, you will need to ensure that
	// a PlaybackKeyPair has been uploaded to your account as this is used to
	// validate the signed JWT that is required for authorization.
	// Experimental.
	Authorized *bool `field:"optional" json:"authorized" yaml:"authorized"`
	// Channel latency mode.
	// Experimental.
	LatencyMode LatencyMode `field:"optional" json:"latencyMode" yaml:"latencyMode"`
	// Channel name.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The channel type, which determines the allowable resolution and bitrate.
	//
	// If you exceed the allowable resolution or bitrate, the stream will disconnect immediately.
	// Experimental.
	Type ChannelType `field:"optional" json:"type" yaml:"type"`
}

