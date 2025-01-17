package awscdklambdapythonalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for bundling.
//
// Example:
//   entry := "/path/to/function"
//   image := awscdk.DockerImage_FromBuild(entry)
//
//   python.NewPythonFunction(this, jsii.String("function"), &PythonFunctionProps{
//   	Entry: jsii.String(Entry),
//   	Runtime: awscdk.Runtime_PYTHON_3_8(),
//   	Bundling: &BundlingOptions{
//   		BuildArgs: map[string]*string{
//   			"PIP_INDEX_URL": jsii.String("https://your.index.url/simple/"),
//   			"PIP_EXTRA_INDEX_URL": jsii.String("https://your.extra-index.url/simple/"),
//   		},
//   	},
//   })
//
// Experimental.
type BundlingOptions struct {
	// The command to run in the container.
	// Default: - run the command defined in the image.
	//
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The entrypoint to run in the container.
	// Default: - run the entrypoint defined in the image.
	//
	// Experimental.
	Entrypoint *[]*string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// The environment variables to pass to the container.
	// Default: - no environment variables.
	//
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Docker [Networking options](https://docs.docker.com/engine/reference/commandline/run/#connect-a-container-to-a-network---network).
	// Default: - no networking options.
	//
	// Experimental.
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Set platform if server is multi-platform capable. _Requires Docker Engine API v1.38+_.
	//
	// Example value: `linux/amd64`.
	// Default: - no platform specified.
	//
	// Experimental.
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// [Security configuration](https://docs.docker.com/engine/reference/run/#security-configuration) when running the docker container.
	// Default: - no security options.
	//
	// Experimental.
	SecurityOpt *string `field:"optional" json:"securityOpt" yaml:"securityOpt"`
	// The user to use when running the container.
	// Default: - root or image default.
	//
	// Experimental.
	User *string `field:"optional" json:"user" yaml:"user"`
	// Docker volumes to mount.
	// Default: - no volumes are mounted.
	//
	// Experimental.
	Volumes *[]*awscdk.DockerVolume `field:"optional" json:"volumes" yaml:"volumes"`
	// Where to mount the specified volumes from.
	// See: https://docs.docker.com/engine/reference/commandline/run/#mount-volumes-from-container---volumes-from
	//
	// Default: - no containers are specified to mount volumes from.
	//
	// Experimental.
	VolumesFrom *[]*string `field:"optional" json:"volumesFrom" yaml:"volumesFrom"`
	// Working directory inside the container.
	// Default: - image default.
	//
	// Experimental.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
	// List of file patterns to exclude when copying assets from source for bundling.
	// Default: - Empty list.
	//
	// Experimental.
	AssetExcludes *[]*string `field:"optional" json:"assetExcludes" yaml:"assetExcludes"`
	// Specify a custom hash for this asset.
	//
	// If `assetHashType` is set it must
	// be set to `AssetHashType.CUSTOM`. For consistency, this custom hash will
	// be SHA256 hashed and encoded as hex. The resulting hash will be the asset
	// hash.
	//
	// NOTE: the hash is used in order to identify a specific revision of the asset, and
	// used for optimizing and caching deployment activities related to this asset such as
	// packaging, uploading to Amazon S3, etc. If you chose to customize the hash, you will
	// need to make sure it is updated every time the asset changes, or otherwise it is
	// possible that some deployments will not be invalidated.
	// Default: - Based on `assetHashType`.
	//
	// Experimental.
	AssetHash *string `field:"optional" json:"assetHash" yaml:"assetHash"`
	// Determines how asset hash is calculated. Assets will get rebuild and uploaded only if their hash has changed.
	//
	// If asset hash is set to `SOURCE` (default), then only changes to the source
	// directory will cause the asset to rebuild. This means, for example, that in
	// order to pick up a new dependency version, a change must be made to the
	// source tree. Ideally, this can be implemented by including a dependency
	// lockfile in your source tree or using fixed dependencies.
	//
	// If the asset hash is set to `OUTPUT`, the hash is calculated after
	// bundling. This means that any change in the output will cause the asset to
	// be invalidated and uploaded. Bear in mind that `pip` adds timestamps to
	// dependencies it installs, which implies that in this mode Python bundles
	// will _always_ get rebuild and uploaded. Normally this is an anti-pattern
	// since build.
	// Default: AssetHashType.SOURCE By default, hash is calculated based on the
	// contents of the source directory. This means that only updates to the
	// source will cause the asset to rebuild.
	//
	// Experimental.
	AssetHashType awscdk.AssetHashType `field:"optional" json:"assetHashType" yaml:"assetHashType"`
	// Optional build arguments to pass to the default container.
	//
	// This can be used to customize
	// the index URLs used for installing dependencies.
	// This is not used if a custom image is provided.
	// Default: - No build arguments.
	//
	// Experimental.
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Which option to use to copy the source files to the docker container and output files back.
	// Default: - BundlingFileAccess.BIND_MOUNT
	//
	// Experimental.
	BundlingFileAccess awscdk.BundlingFileAccess `field:"optional" json:"bundlingFileAccess" yaml:"bundlingFileAccess"`
	// Command hooks.
	// Default: - do not run additional commands.
	//
	// Experimental.
	CommandHooks ICommandHooks `field:"optional" json:"commandHooks" yaml:"commandHooks"`
	// Docker image to use for bundling.
	//
	// If no options are provided, the default bundling image
	// will be used. Dependencies will be installed using the default packaging commands
	// and copied over from into the Lambda asset.
	// Default: - Default bundling image.
	//
	// Experimental.
	Image awscdk.DockerImage `field:"optional" json:"image" yaml:"image"`
	// Output path suffix: the suffix for the directory into which the bundled output is written.
	// Default: - 'python' for a layer, empty string otherwise.
	//
	// Experimental.
	OutputPathSuffix *string `field:"optional" json:"outputPathSuffix" yaml:"outputPathSuffix"`
	// Whether to export Poetry dependencies with hashes.
	//
	// Note that this can cause builds to fail if not all dependencies
	// export with a hash.
	// See: https://github.com/aws/aws-cdk/issues/19232
	//
	// Default: Hashes are NOT included in the exported `requirements.txt` file
	//
	// Experimental.
	PoetryIncludeHashes *bool `field:"optional" json:"poetryIncludeHashes" yaml:"poetryIncludeHashes"`
}

