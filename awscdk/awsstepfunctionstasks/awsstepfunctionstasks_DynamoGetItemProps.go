package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Properties for DynamoGetItem Task.
//
// Example:
//   var myTable table
//
//   tasks.NewDynamoGetItem(this, jsii.String("Get Item"), &dynamoGetItemProps{
//   	key: map[string]dynamoAttributeValue{
//   		"messageId": tasks.*dynamoAttributeValue.fromString(jsii.String("message-007")),
//   	},
//   	table: myTable,
//   })
//
// Experimental.
type DynamoGetItemProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// Primary key of the item to retrieve.
	//
	// For the primary key, you must provide all of the attributes.
	// For example, with a simple primary key, you only need to provide a value for the partition key.
	// For a composite primary key, you must provide values for both the partition key and the sort key.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html#DDB-GetItem-request-Key
	//
	// Experimental.
	Key *map[string]DynamoAttributeValue `field:"required" json:"key" yaml:"key"`
	// The name of the table containing the requested item.
	// Experimental.
	Table awsdynamodb.ITable `field:"required" json:"table" yaml:"table"`
	// Determines the read consistency model: If set to true, then the operation uses strongly consistent reads;
	//
	// otherwise, the operation uses eventually consistent reads.
	// Experimental.
	ConsistentRead *bool `field:"optional" json:"consistentRead" yaml:"consistentRead"`
	// One or more substitution tokens for attribute names in an expression.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html#DDB-GetItem-request-ExpressionAttributeNames
	//
	// Experimental.
	ExpressionAttributeNames *map[string]*string `field:"optional" json:"expressionAttributeNames" yaml:"expressionAttributeNames"`
	// An array of DynamoProjectionExpression that identifies one or more attributes to retrieve from the table.
	//
	// These attributes can include scalars, sets, or elements of a JSON document.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html#DDB-GetItem-request-ProjectionExpression
	//
	// Experimental.
	ProjectionExpression *[]DynamoProjectionExpression `field:"optional" json:"projectionExpression" yaml:"projectionExpression"`
	// Determines the level of detail about provisioned throughput consumption that is returned in the response.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html#DDB-GetItem-request-ReturnConsumedCapacity
	//
	// Experimental.
	ReturnConsumedCapacity DynamoConsumedCapacity `field:"optional" json:"returnConsumedCapacity" yaml:"returnConsumedCapacity"`
}
