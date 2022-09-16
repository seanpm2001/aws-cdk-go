package awsrds


// Properties for Oracle Enterprise Edition instance engines.
//
// Used in {@link DatabaseInstanceEngine.oracleEe}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var oracleEngineVersion oracleEngineVersion
//
//   oracleEeInstanceEngineProps := &oracleEeInstanceEngineProps{
//   	version: oracleEngineVersion,
//   }
//
// Experimental.
type OracleEeInstanceEngineProps struct {
	// The exact version of the engine to use.
	// Experimental.
	Version OracleEngineVersion `field:"required" json:"version" yaml:"version"`
}
