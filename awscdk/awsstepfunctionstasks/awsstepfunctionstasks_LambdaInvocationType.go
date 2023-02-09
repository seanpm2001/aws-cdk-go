package awsstepfunctionstasks


// Invocation type of a Lambda.
//
// Example:
//   var fn function
//
//
//   submitJob := tasks.NewLambdaInvoke(this, jsii.String("Invoke Handler"), &lambdaInvokeProps{
//   	lambdaFunction: fn,
//   	payload: sfn.taskInput.fromJsonPathAt(jsii.String("$.input")),
//   	invocationType: tasks.lambdaInvocationType_EVENT,
//   })
//
// Experimental.
type LambdaInvocationType string

const (
	// Invoke the function synchronously.
	//
	// Keep the connection open until the function returns a response or times out.
	// The API response includes the function response and additional data.
	// Experimental.
	LambdaInvocationType_REQUEST_RESPONSE LambdaInvocationType = "REQUEST_RESPONSE"
	// Invoke the function asynchronously.
	//
	// Send events that fail multiple times to the function's dead-letter queue (if it's configured).
	// The API response only includes a status code.
	// Experimental.
	LambdaInvocationType_EVENT LambdaInvocationType = "EVENT"
	// Validate parameter values and verify that the user or role has permission to invoke the function.
	// Experimental.
	LambdaInvocationType_DRY_RUN LambdaInvocationType = "DRY_RUN"
)

