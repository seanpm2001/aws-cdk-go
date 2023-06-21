package awsconnect


// Information about a section from an evaluation form.
//
// A section can contain sections and/or questions. Evaluation forms can only contain sections and subsections (two level nesting).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var evaluationFormItemProperty_ evaluationFormItemProperty
//
//   evaluationFormSectionProperty := &EvaluationFormSectionProperty{
//   	RefId: jsii.String("refId"),
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
//   	Instructions: jsii.String("instructions"),
//   	Items: []interface{}{
//   		&evaluationFormItemProperty{
//   			Question: &EvaluationFormQuestionProperty{
//   				QuestionType: jsii.String("questionType"),
//   				RefId: jsii.String("refId"),
//   				Title: jsii.String("title"),
//
//   				// the properties below are optional
//   				Instructions: jsii.String("instructions"),
//   				NotApplicableEnabled: jsii.Boolean(false),
//   				QuestionTypeProperties: &EvaluationFormQuestionTypePropertiesProperty{
//   					Numeric: &EvaluationFormNumericQuestionPropertiesProperty{
//   						MaxValue: jsii.Number(123),
//   						MinValue: jsii.Number(123),
//
//   						// the properties below are optional
//   						Automation: &EvaluationFormNumericQuestionAutomationProperty{
//   							PropertyValue: &NumericQuestionPropertyValueAutomationProperty{
//   								Label: jsii.String("label"),
//   							},
//   						},
//   						Options: []interface{}{
//   							&EvaluationFormNumericQuestionOptionProperty{
//   								MaxValue: jsii.Number(123),
//   								MinValue: jsii.Number(123),
//
//   								// the properties below are optional
//   								AutomaticFail: jsii.Boolean(false),
//   								Score: jsii.Number(123),
//   							},
//   						},
//   					},
//   					SingleSelect: &EvaluationFormSingleSelectQuestionPropertiesProperty{
//   						Options: []interface{}{
//   							&EvaluationFormSingleSelectQuestionOptionProperty{
//   								RefId: jsii.String("refId"),
//   								Text: jsii.String("text"),
//
//   								// the properties below are optional
//   								AutomaticFail: jsii.Boolean(false),
//   								Score: jsii.Number(123),
//   							},
//   						},
//
//   						// the properties below are optional
//   						Automation: &EvaluationFormSingleSelectQuestionAutomationProperty{
//   							Options: []interface{}{
//   								&EvaluationFormSingleSelectQuestionAutomationOptionProperty{
//   									RuleCategory: &SingleSelectQuestionRuleCategoryAutomationProperty{
//   										Category: jsii.String("category"),
//   										Condition: jsii.String("condition"),
//   										OptionRefId: jsii.String("optionRefId"),
//   									},
//   								},
//   							},
//
//   							// the properties below are optional
//   							DefaultOptionRefId: jsii.String("defaultOptionRefId"),
//   						},
//   						DisplayAs: jsii.String("displayAs"),
//   					},
//   				},
//   				Weight: jsii.Number(123),
//   			},
//   			Section: &EvaluationFormSectionProperty{
//   				RefId: jsii.String("refId"),
//   				Title: jsii.String("title"),
//
//   				// the properties below are optional
//   				Instructions: jsii.String("instructions"),
//   				Items: []interface{}{
//   					evaluationFormItemProperty_,
//   				},
//   				Weight: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Weight: jsii.Number(123),
//   }
//
type CfnEvaluationForm_EvaluationFormSectionProperty struct {
	// The identifier of the section. An identifier must be unique within the evaluation form.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 40.
	RefId *string `field:"required" json:"refId" yaml:"refId"`
	// The title of the section.
	//
	// *Length Constraints* : Minimum length of 1. Maximum length of 128.
	Title *string `field:"required" json:"title" yaml:"title"`
	// The instructions of the section.
	Instructions *string `field:"optional" json:"instructions" yaml:"instructions"`
	// The items of the section.
	//
	// *Minimum* : 1.
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	// The scoring weight of the section.
	//
	// *Minimum* : 0
	//
	// *Maximum* : 100.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}
