package awscodepipelineactions


// The type of the Jenkins Action that determines its CodePipeline Category - Build, or Test.
//
// Note that a Jenkins provider, even if it has the same name,
// must be separately registered for each type.
//
// Example:
//   var jenkinsProvider jenkinsProvider
//
//   buildAction := codepipeline_actions.NewJenkinsAction(&jenkinsActionProps{
//   	actionName: jsii.String("JenkinsBuild"),
//   	jenkinsProvider: jenkinsProvider,
//   	projectName: jsii.String("MyProject"),
//   	type: codepipeline_actions.jenkinsActionType_BUILD,
//   })
//
// Experimental.
type JenkinsActionType string

const (
	// The Action will have the Build Category.
	// Experimental.
	JenkinsActionType_BUILD JenkinsActionType = "BUILD"
	// The Action will have the Test Category.
	// Experimental.
	JenkinsActionType_TEST JenkinsActionType = "TEST"
)
