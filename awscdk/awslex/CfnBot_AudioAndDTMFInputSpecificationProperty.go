package awslex


// Specifies the audio and DTMF input specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioAndDTMFInputSpecificationProperty := &AudioAndDTMFInputSpecificationProperty{
//   	StartTimeoutMs: jsii.Number(123),
//
//   	// the properties below are optional
//   	AudioSpecification: &AudioSpecificationProperty{
//   		EndTimeoutMs: jsii.Number(123),
//   		MaxLengthMs: jsii.Number(123),
//   	},
//   	DtmfSpecification: &DTMFSpecificationProperty{
//   		DeletionCharacter: jsii.String("deletionCharacter"),
//   		EndCharacter: jsii.String("endCharacter"),
//   		EndTimeoutMs: jsii.Number(123),
//   		MaxLength: jsii.Number(123),
//   	},
//   }
//
type CfnBot_AudioAndDTMFInputSpecificationProperty struct {
	// Time for which a bot waits before assuming that the customer isn't going to speak or press a key.
	//
	// This timeout is shared between Audio and DTMF inputs.
	StartTimeoutMs *float64 `field:"required" json:"startTimeoutMs" yaml:"startTimeoutMs"`
	// Specifies the settings on audio input.
	AudioSpecification interface{} `field:"optional" json:"audioSpecification" yaml:"audioSpecification"`
	// Specifies the settings on DTMF input.
	DtmfSpecification interface{} `field:"optional" json:"dtmfSpecification" yaml:"dtmfSpecification"`
}

