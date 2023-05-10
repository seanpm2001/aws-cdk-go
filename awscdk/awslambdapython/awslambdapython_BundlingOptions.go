package awslambdapython

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Options for bundling.
//
// Example:
//   entry := "/path/to/function"
//   image := awscdk.DockerImage_FromBuild(entry)
//
//   lambda.NewPythonFunction(this, jsii.String("function"), &PythonFunctionProps{
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
	// Experimental.
	AssetHashType awscdk.AssetHashType `field:"optional" json:"assetHashType" yaml:"assetHashType"`
	// Optional build arguments to pass to the default container.
	//
	// This can be used to customize
	// the index URLs used for installing dependencies.
	// This is not used if a custom image is provided.
	// Experimental.
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Environment variables defined when bundling runs.
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Docker image to use for bundling.
	//
	// If no options are provided, the default bundling image
	// will be used. Dependencies will be installed using the default packaging commands
	// and copied over from into the Lambda asset.
	// Experimental.
	Image awscdk.DockerImage `field:"optional" json:"image" yaml:"image"`
	// Output path suffix: the suffix for the directory into which the bundled output is written.
	// Experimental.
	OutputPathSuffix *string `field:"optional" json:"outputPathSuffix" yaml:"outputPathSuffix"`
}

