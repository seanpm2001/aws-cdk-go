package awssagemaker


// The content of the model card.
//
// It follows the [model card json schema](https://docs.aws.amazon.com/sagemaker/latest/dg/model-cards.html#model-cards-json-schema) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var value interface{}
//
//   contentProperty := &contentProperty{
//   	additionalInformation: &additionalInformationProperty{
//   		caveatsAndRecommendations: jsii.String("caveatsAndRecommendations"),
//   		customDetails: map[string]*string{
//   			"customDetailsKey": jsii.String("customDetails"),
//   		},
//   		ethicalConsiderations: jsii.String("ethicalConsiderations"),
//   	},
//   	businessDetails: &businessDetailsProperty{
//   		businessProblem: jsii.String("businessProblem"),
//   		businessStakeholders: jsii.String("businessStakeholders"),
//   		lineOfBusiness: jsii.String("lineOfBusiness"),
//   	},
//   	evaluationDetails: []interface{}{
//   		&evaluationDetailProperty{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			datasets: []*string{
//   				jsii.String("datasets"),
//   			},
//   			evaluationJobArn: jsii.String("evaluationJobArn"),
//   			evaluationObservation: jsii.String("evaluationObservation"),
//   			metadata: map[string]*string{
//   				"metadataKey": jsii.String("metadata"),
//   			},
//   			metricGroups: []interface{}{
//   				&metricGroupProperty{
//   					metricData: []interface{}{
//   						&metricDataItemsProperty{
//   							name: jsii.String("name"),
//   							type: jsii.String("type"),
//   							value: value,
//
//   							// the properties below are optional
//   							notes: jsii.String("notes"),
//   							xAxisName: []*string{
//   								jsii.String("xAxisName"),
//   							},
//   							yAxisName: []*string{
//   								jsii.String("yAxisName"),
//   							},
//   						},
//   					},
//   					name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   	intendedUses: &intendedUsesProperty{
//   		explanationsForRiskRating: jsii.String("explanationsForRiskRating"),
//   		factorsAffectingModelEfficiency: jsii.String("factorsAffectingModelEfficiency"),
//   		intendedUses: jsii.String("intendedUses"),
//   		purposeOfModel: jsii.String("purposeOfModel"),
//   		riskRating: jsii.String("riskRating"),
//   	},
//   	modelOverview: &modelOverviewProperty{
//   		algorithmType: jsii.String("algorithmType"),
//   		inferenceEnvironment: &inferenceEnvironmentProperty{
//   			containerImage: []*string{
//   				jsii.String("containerImage"),
//   			},
//   		},
//   		modelArtifact: []*string{
//   			jsii.String("modelArtifact"),
//   		},
//   		modelCreator: jsii.String("modelCreator"),
//   		modelDescription: jsii.String("modelDescription"),
//   		modelId: jsii.String("modelId"),
//   		modelName: jsii.String("modelName"),
//   		modelOwner: jsii.String("modelOwner"),
//   		modelVersion: jsii.Number(123),
//   		problemType: jsii.String("problemType"),
//   	},
//   	trainingDetails: &trainingDetailsProperty{
//   		objectiveFunction: &objectiveFunctionProperty{
//   			function: &functionProperty{
//   				condition: jsii.String("condition"),
//   				facet: jsii.String("facet"),
//   				function: jsii.String("function"),
//   			},
//   			notes: jsii.String("notes"),
//   		},
//   		trainingJobDetails: &trainingJobDetailsProperty{
//   			hyperParameters: []interface{}{
//   				&trainingHyperParameterProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			trainingArn: jsii.String("trainingArn"),
//   			trainingDatasets: []*string{
//   				jsii.String("trainingDatasets"),
//   			},
//   			trainingEnvironment: &trainingEnvironmentProperty{
//   				containerImage: []*string{
//   					jsii.String("containerImage"),
//   				},
//   			},
//   			trainingMetrics: []interface{}{
//   				&trainingMetricProperty{
//   					name: jsii.String("name"),
//   					value: jsii.Number(123),
//
//   					// the properties below are optional
//   					notes: jsii.String("notes"),
//   				},
//   			},
//   			userProvidedHyperParameters: []interface{}{
//   				&trainingHyperParameterProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			userProvidedTrainingMetrics: []interface{}{
//   				&trainingMetricProperty{
//   					name: jsii.String("name"),
//   					value: jsii.Number(123),
//
//   					// the properties below are optional
//   					notes: jsii.String("notes"),
//   				},
//   			},
//   		},
//   		trainingObservations: jsii.String("trainingObservations"),
//   	},
//   }
//
type CfnModelCard_ContentProperty struct {
	// Additional information about the model.
	AdditionalInformation interface{} `field:"optional" json:"additionalInformation" yaml:"additionalInformation"`
	// Information about how the model supports business goals.
	BusinessDetails interface{} `field:"optional" json:"businessDetails" yaml:"businessDetails"`
	// An overview about the model's evaluation.
	EvaluationDetails interface{} `field:"optional" json:"evaluationDetails" yaml:"evaluationDetails"`
	// The intended usage of the model.
	IntendedUses interface{} `field:"optional" json:"intendedUses" yaml:"intendedUses"`
	// An overview about the model.
	ModelOverview interface{} `field:"optional" json:"modelOverview" yaml:"modelOverview"`
	// An overview about model training.
	TrainingDetails interface{} `field:"optional" json:"trainingDetails" yaml:"trainingDetails"`
}
