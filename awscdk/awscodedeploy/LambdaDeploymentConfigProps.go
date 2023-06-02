package awscodedeploy


// Construction properties of `LambdaDeploymentConfig`.
//
// Example:
//   var application lambdaApplication
//   var alias alias
//   config := codedeploy.NewLambdaDeploymentConfig(this, jsii.String("CustomConfig"), &LambdaDeploymentConfigProps{
//   	TrafficRouting: codedeploy.NewTimeBasedCanaryTrafficRouting(&TimeBasedCanaryTrafficRoutingProps{
//   		Interval: awscdk.Duration_Minutes(jsii.Number(15)),
//   		Percentage: jsii.Number(5),
//   	}),
//   })
//   deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &LambdaDeploymentGroupProps{
//   	Application: Application,
//   	Alias: Alias,
//   	DeploymentConfig: config,
//   })
//
type LambdaDeploymentConfigProps struct {
	// The physical, human-readable name of the Deployment Configuration.
	DeploymentConfigName *string `field:"optional" json:"deploymentConfigName" yaml:"deploymentConfigName"`
	// The configuration that specifies how traffic is shifted from the 'blue' target group to the 'green' target group during a deployment.
	TrafficRouting TrafficRouting `field:"optional" json:"trafficRouting" yaml:"trafficRouting"`
}

