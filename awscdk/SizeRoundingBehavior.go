// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Rounding behaviour when converting between units of `Size`.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   awscdk.Size_Mebibytes(jsii.Number(2)).ToKibibytes() // yields 2048
//   awscdk.Size_Kibibytes(jsii.Number(2050)).ToMebibytes(&SizeConversionOptions{
//   	Rounding: awscdk.SizeRoundingBehavior_FLOOR,
//   })
//
type SizeRoundingBehavior string

const (
	// Fail the conversion if the result is not an integer.
	SizeRoundingBehavior_FAIL SizeRoundingBehavior = "FAIL"
	// If the result is not an integer, round it to the closest integer less than the result.
	SizeRoundingBehavior_FLOOR SizeRoundingBehavior = "FLOOR"
	// Don't round.
	//
	// Return even if the result is a fraction.
	SizeRoundingBehavior_NONE SizeRoundingBehavior = "NONE"
)

