// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   dockerImageAssetSource := &DockerImageAssetSource{
//   	SourceHash: jsii.String("sourceHash"),
//
//   	// the properties below are optional
//   	DirectoryName: jsii.String("directoryName"),
//   	DockerBuildArgs: map[string]*string{
//   		"dockerBuildArgsKey": jsii.String("dockerBuildArgs"),
//   	},
//   	DockerBuildSecrets: map[string]*string{
//   		"dockerBuildSecretsKey": jsii.String("dockerBuildSecrets"),
//   	},
//   	DockerBuildTarget: jsii.String("dockerBuildTarget"),
//   	DockerCacheFrom: []dockerCacheOption{
//   		&dockerCacheOption{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Params: map[string]*string{
//   				"paramsKey": jsii.String("params"),
//   			},
//   		},
//   	},
//   	DockerCacheTo: &dockerCacheOption{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Params: map[string]*string{
//   			"paramsKey": jsii.String("params"),
//   		},
//   	},
//   	DockerFile: jsii.String("dockerFile"),
//   	DockerOutputs: []*string{
//   		jsii.String("dockerOutputs"),
//   	},
//   	Executable: []*string{
//   		jsii.String("executable"),
//   	},
//   	NetworkMode: jsii.String("networkMode"),
//   	Platform: jsii.String("platform"),
//   }
//
type DockerImageAssetSource struct {
	// The hash of the contents of the docker build context.
	//
	// This hash is used
	// throughout the system to identify this image and avoid duplicate work
	// in case the source did not change.
	//
	// NOTE: this means that if you wish to update your docker image, you
	// must make a modification to the source (e.g. add some metadata to your Dockerfile).
	SourceHash *string `field:"required" json:"sourceHash" yaml:"sourceHash"`
	// The directory where the Dockerfile is stored, must be relative to the cloud assembly root.
	DirectoryName *string `field:"optional" json:"directoryName" yaml:"directoryName"`
	// Build args to pass to the `docker build` command.
	//
	// Since Docker build arguments are resolved before deployment, keys and
	// values cannot refer to unresolved tokens (such as `lambda.functionArn` or
	// `queue.queueUrl`).
	//
	// Only allowed when `directoryName` is specified.
	DockerBuildArgs *map[string]*string `field:"optional" json:"dockerBuildArgs" yaml:"dockerBuildArgs"`
	// Build secrets to pass to the `docker build` command.
	//
	// Since Docker build secrets are resolved before deployment, keys and
	// values cannot refer to unresolved tokens (such as `lambda.functionArn` or
	// `queue.queueUrl`).
	//
	// Only allowed when `directoryName` is specified.
	DockerBuildSecrets *map[string]*string `field:"optional" json:"dockerBuildSecrets" yaml:"dockerBuildSecrets"`
	// Docker target to build to.
	//
	// Only allowed when `directoryName` is specified.
	DockerBuildTarget *string `field:"optional" json:"dockerBuildTarget" yaml:"dockerBuildTarget"`
	// Cache from options to pass to the `docker build` command.
	DockerCacheFrom *[]*DockerCacheOption `field:"optional" json:"dockerCacheFrom" yaml:"dockerCacheFrom"`
	// Cache to options to pass to the `docker build` command.
	DockerCacheTo *DockerCacheOption `field:"optional" json:"dockerCacheTo" yaml:"dockerCacheTo"`
	// Path to the Dockerfile (relative to the directory).
	//
	// Only allowed when `directoryName` is specified.
	DockerFile *string `field:"optional" json:"dockerFile" yaml:"dockerFile"`
	// Outputs to pass to the `docker build` command.
	DockerOutputs *[]*string `field:"optional" json:"dockerOutputs" yaml:"dockerOutputs"`
	// An external command that will produce the packaged asset.
	//
	// The command should produce the name of a local Docker image on `stdout`.
	Executable *[]*string `field:"optional" json:"executable" yaml:"executable"`
	// Networking mode for the RUN commands during build. _Requires Docker Engine API v1.25+_.
	//
	// Specify this property to build images on a specific networking mode.
	NetworkMode *string `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Platform to build for. _Requires Docker Buildx_.
	//
	// Specify this property to build images on a specific platform.
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
}

