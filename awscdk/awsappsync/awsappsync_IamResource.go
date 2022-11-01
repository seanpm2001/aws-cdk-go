package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A class used to generate resource arns for AppSync.
//
// Example:
//   var api graphqlApi
//   role := iam.NewRole(this, jsii.String("Role"), &roleProps{
//   	assumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
//   })
//
//   api.grant(role, appsync.iamResource.custom(jsii.String("types/Mutation/fields/updateExample")), jsii.String("appsync:GraphQL"))
//
// Experimental.
type IamResource interface {
	// Return the Resource ARN.
	// Experimental.
	ResourceArns(api GraphqlApi) *[]*string
}

// The jsii proxy struct for IamResource
type jsiiProxy_IamResource struct {
	_ byte // padding
}

// Generate the resource names that accepts all types: `*`.
// Experimental.
func IamResource_All() IamResource {
	_init_.Initialize()

	var returns IamResource

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.IamResource",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Generate the resource names given custom arns.
// Experimental.
func IamResource_Custom(arns ...*string) IamResource {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range arns {
		args = append(args, a)
	}

	var returns IamResource

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.IamResource",
		"custom",
		args,
		&returns,
	)

	return returns
}

// Generate the resource names given a type and fields.
// Experimental.
func IamResource_OfType(type_ *string, fields ...*string) IamResource {
	_init_.Initialize()

	if err := validateIamResource_OfTypeParameters(type_); err != nil {
		panic(err)
	}
	args := []interface{}{type_}
	for _, a := range fields {
		args = append(args, a)
	}

	var returns IamResource

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.IamResource",
		"ofType",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamResource) ResourceArns(api GraphqlApi) *[]*string {
	if err := i.validateResourceArnsParameters(api); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"resourceArns",
		[]interface{}{api},
		&returns,
	)

	return returns
}
