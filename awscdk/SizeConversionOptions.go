// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Options for how to convert time to a different unit.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   awscdk.Size_Mebibytes(jsii.Number(2)).ToKibibytes() // yields 2048
//   awscdk.Size_Kibibytes(jsii.Number(2050)).ToMebibytes(&SizeConversionOptions{
//   	Rounding: awscdk.SizeRoundingBehavior_FLOOR,
//   })
//
type SizeConversionOptions struct {
	// How conversions should behave when it encounters a non-integer result.
	Rounding SizeRoundingBehavior `field:"optional" json:"rounding" yaml:"rounding"`
}

