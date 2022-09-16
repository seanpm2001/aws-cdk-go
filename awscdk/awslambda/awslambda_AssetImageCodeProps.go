package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/assets"
	"github.com/aws/aws-cdk-go/awscdk/awsecrassets"
)

// Properties to initialize a new AssetImage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkMode networkMode
//   var platform platform
//
//   assetImageCodeProps := &assetImageCodeProps{
//   	buildArgs: map[string]*string{
//   		"buildArgsKey": jsii.String("buildArgs"),
//   	},
//   	cmd: []*string{
//   		jsii.String("cmd"),
//   	},
//   	entrypoint: []*string{
//   		jsii.String("entrypoint"),
//   	},
//   	exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	extraHash: jsii.String("extraHash"),
//   	file: jsii.String("file"),
//   	follow: awscdk.Assets.followMode_NEVER,
//   	followSymlinks: monocdk.symlinkFollowMode_NEVER,
//   	ignoreMode: monocdk.ignoreMode_GLOB,
//   	invalidation: &dockerImageAssetInvalidationOptions{
//   		buildArgs: jsii.Boolean(false),
//   		extraHash: jsii.Boolean(false),
//   		file: jsii.Boolean(false),
//   		networkMode: jsii.Boolean(false),
//   		platform: jsii.Boolean(false),
//   		repositoryName: jsii.Boolean(false),
//   		target: jsii.Boolean(false),
//   	},
//   	networkMode: networkMode,
//   	platform: platform,
//   	repositoryName: jsii.String("repositoryName"),
//   	target: jsii.String("target"),
//   	workingDirectory: jsii.String("workingDirectory"),
//   }
//
// Experimental.
type AssetImageCodeProps struct {
	// Glob patterns to exclude from the copy.
	// Experimental.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	// Deprecated: use `followSymlinks` instead.
	Follow assets.FollowMode `field:"optional" json:"follow" yaml:"follow"`
	// The ignore behavior to use for exclude patterns.
	// Experimental.
	IgnoreMode awscdk.IgnoreMode `field:"optional" json:"ignoreMode" yaml:"ignoreMode"`
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	// Experimental.
	ExtraHash *string `field:"optional" json:"extraHash" yaml:"extraHash"`
	// A strategy for how to handle symlinks.
	// Experimental.
	FollowSymlinks awscdk.SymlinkFollowMode `field:"optional" json:"followSymlinks" yaml:"followSymlinks"`
	// Build args to pass to the `docker build` command.
	//
	// Since Docker build arguments are resolved before deployment, keys and
	// values cannot refer to unresolved tokens (such as `lambda.functionArn` or
	// `queue.queueUrl`).
	// Experimental.
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Path to the Dockerfile (relative to the directory).
	// Experimental.
	File *string `field:"optional" json:"file" yaml:"file"`
	// Options to control which parameters are used to invalidate the asset hash.
	// Experimental.
	Invalidation *awsecrassets.DockerImageAssetInvalidationOptions `field:"optional" json:"invalidation" yaml:"invalidation"`
	// Networking mode for the RUN commands during build.
	//
	// Support docker API 1.25+.
	// Experimental.
	NetworkMode awsecrassets.NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Platform to build for.
	//
	// _Requires Docker Buildx_.
	// Experimental.
	Platform awsecrassets.Platform `field:"optional" json:"platform" yaml:"platform"`
	// ECR repository name.
	//
	// Specify this property if you need to statically address the image, e.g.
	// from a Kubernetes Pod. Note, this is only the repository name, without the
	// registry and the tag parts.
	// Deprecated: to control the location of docker image assets, please override
	// `Stack.addDockerImageAsset`. this feature will be removed in future
	// releases.
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// Docker target to build to.
	// Experimental.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Specify or override the CMD on the specified Docker image or Dockerfile.
	//
	// This needs to be in the 'exec form', viz., `[ 'executable', 'param1', 'param2' ]`.
	// See: https://docs.docker.com/engine/reference/builder/#cmd
	//
	// Experimental.
	Cmd *[]*string `field:"optional" json:"cmd" yaml:"cmd"`
	// Specify or override the ENTRYPOINT on the specified Docker image or Dockerfile.
	//
	// An ENTRYPOINT allows you to configure a container that will run as an executable.
	// This needs to be in the 'exec form', viz., `[ 'executable', 'param1', 'param2' ]`.
	// See: https://docs.docker.com/engine/reference/builder/#entrypoint
	//
	// Experimental.
	Entrypoint *[]*string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// Specify or override the WORKDIR on the specified Docker image or Dockerfile.
	//
	// A WORKDIR allows you to configure the working directory the container will use.
	// See: https://docs.docker.com/engine/reference/builder/#workdir
	//
	// Experimental.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}
