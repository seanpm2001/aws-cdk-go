package awsecs


// Details on an Elastic Inference accelerator.
//
// For more information, see [Working with Amazon Elastic Inference on Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-inference.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceAcceleratorProperty := &InferenceAcceleratorProperty{
//   	DeviceName: jsii.String("deviceName"),
//   	DeviceType: jsii.String("deviceType"),
//   }
//
type CfnTaskDefinition_InferenceAcceleratorProperty struct {
	// The Elastic Inference accelerator device name.
	//
	// The `deviceName` must also be referenced in a container definition as a [ResourceRequirement](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ResourceRequirement.html) .
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// The Elastic Inference accelerator type to use.
	DeviceType *string `field:"optional" json:"deviceType" yaml:"deviceType"`
}

