package awscdksyntheticsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Options for specifying the s3 location that stores the data of each canary run.
//
// The artifacts bucket location **cannot**
// be updated once the canary is created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import synthetics_alpha "github.com/aws/aws-cdk-go/awscdksyntheticsalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   artifactsBucketLocation := &ArtifactsBucketLocation{
//   	Bucket: bucket,
//
//   	// the properties below are optional
//   	Prefix: jsii.String("prefix"),
//   }
//
// Experimental.
type ArtifactsBucketLocation struct {
	// The s3 location that stores the data of each run.
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The S3 bucket prefix.
	//
	// Specify this if you want a more specific path within the artifacts bucket.
	// Default: - no prefix.
	//
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

