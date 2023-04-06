package awssagemaker


// A collection of settings that apply to an `RSessionGateway` app.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rSessionAppSettingsProperty := &RSessionAppSettingsProperty{
//   	CustomImages: []interface{}{
//   		&CustomImageProperty{
//   			AppImageConfigName: jsii.String("appImageConfigName"),
//   			ImageName: jsii.String("imageName"),
//
//   			// the properties below are optional
//   			ImageVersionNumber: jsii.Number(123),
//   		},
//   	},
//   	DefaultResourceSpec: &ResourceSpecProperty{
//   		InstanceType: jsii.String("instanceType"),
//   		LifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   		SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   		SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   	},
//   }
//
type CfnDomain_RSessionAppSettingsProperty struct {
	// A list of custom SageMaker images that are configured to run as a RSession app.
	CustomImages interface{} `field:"optional" json:"customImages" yaml:"customImages"`
	// Specifies the ARNs of a SageMaker image and SageMaker image version, and the instance type that the version runs on.
	DefaultResourceSpec interface{} `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
}
