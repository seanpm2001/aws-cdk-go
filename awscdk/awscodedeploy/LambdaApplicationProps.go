package awscodedeploy


// Construction properties for `LambdaApplication`.
//
// Example:
//   application := codedeploy.NewLambdaApplication(this, jsii.String("CodeDeployApplication"), &LambdaApplicationProps{
//   	ApplicationName: jsii.String("MyApplication"),
//   })
//
type LambdaApplicationProps struct {
	// The physical, human-readable name of the CodeDeploy Application.
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
}

