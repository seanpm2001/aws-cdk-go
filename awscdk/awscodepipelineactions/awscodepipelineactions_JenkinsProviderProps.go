package awscodepipelineactions


// Example:
//   jenkinsProvider := codepipeline_actions.NewJenkinsProvider(this, jsii.String("JenkinsProvider"), &jenkinsProviderProps{
//   	providerName: jsii.String("MyJenkinsProvider"),
//   	serverUrl: jsii.String("http://my-jenkins.com:8080"),
//   	version: jsii.String("2"),
//   })
//
// Experimental.
type JenkinsProviderProps struct {
	// The name of the Jenkins provider that you set in the AWS CodePipeline plugin configuration of your Jenkins project.
	//
	// Example:
	//   "MyJenkinsProvider"
	//
	// Experimental.
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// The base URL of your Jenkins server.
	//
	// Example:
	//   "http://myjenkins.com:8080"
	//
	// Experimental.
	ServerUrl *string `field:"required" json:"serverUrl" yaml:"serverUrl"`
	// Whether to immediately register a Jenkins Provider for the build category.
	//
	// The Provider will always be registered if you create a {@link JenkinsAction}.
	// Experimental.
	ForBuild *bool `field:"optional" json:"forBuild" yaml:"forBuild"`
	// Whether to immediately register a Jenkins Provider for the test category.
	//
	// The Provider will always be registered if you create a {@link JenkinsTestAction}.
	// Experimental.
	ForTest *bool `field:"optional" json:"forTest" yaml:"forTest"`
	// The version of your provider.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

