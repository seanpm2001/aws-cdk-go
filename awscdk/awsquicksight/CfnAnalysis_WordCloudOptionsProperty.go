package awsquicksight


// The word cloud options for a word cloud visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   wordCloudOptionsProperty := &WordCloudOptionsProperty{
//   	CloudLayout: jsii.String("cloudLayout"),
//   	MaximumStringLength: jsii.Number(123),
//   	WordCasing: jsii.String("wordCasing"),
//   	WordOrientation: jsii.String("wordOrientation"),
//   	WordPadding: jsii.String("wordPadding"),
//   	WordScaling: jsii.String("wordScaling"),
//   }
//
type CfnAnalysis_WordCloudOptionsProperty struct {
	// The cloud layout options (fluid, normal) of a word cloud.
	CloudLayout *string `field:"optional" json:"cloudLayout" yaml:"cloudLayout"`
	// The length limit of each word from 1-100.
	MaximumStringLength *float64 `field:"optional" json:"maximumStringLength" yaml:"maximumStringLength"`
	// The word casing options (lower_case, existing_case) for the words in a word cloud.
	WordCasing *string `field:"optional" json:"wordCasing" yaml:"wordCasing"`
	// The word orientation options (horizontal, horizontal_and_vertical) for the words in a word cloud.
	WordOrientation *string `field:"optional" json:"wordOrientation" yaml:"wordOrientation"`
	// The word padding options (none, small, medium, large) for the words in a word cloud.
	WordPadding *string `field:"optional" json:"wordPadding" yaml:"wordPadding"`
	// The word scaling options (emphasize, normal) for the words in a word cloud.
	WordScaling *string `field:"optional" json:"wordScaling" yaml:"wordScaling"`
}

