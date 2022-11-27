// The CDK Construct Library for AWS::ServiceCatalogAppRegistry
package awscdkservicecatalogappregistryalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkservicecatalogappregistryalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Contains static factory methods with which you can build the input needed for application associator to work.
//
// Example:
//   app := awscdk.NewApp()
//   associatedApp := appreg.NewApplicationAssociator(app, jsii.String("AssociatedApplication"), &applicationAssociatorProps{
//   	applications: []targetApplication{
//   		appreg.*targetApplication.createApplicationStack(&createTargetApplicationOptions{
//   			applicationName: jsii.String("MyAssociatedApplication"),
//   			// 'Application containing stacks deployed via CDK.' is the default
//   			applicationDescription: jsii.String("Associated Application description"),
//   			stackName: jsii.String("MyAssociatedApplicationStack"),
//   			// AWS Account and Region that are implied by the current CLI configuration is the default
//   			env: &environment{
//   				account: jsii.String("123456789012"),
//   				region: jsii.String("us-east-1"),
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type TargetApplication interface {
	// Called when the ApplicationAssociator is initialized.
	// Experimental.
	Bind(scope constructs.Construct) *BindTargetApplicationResult
}

// The jsii proxy struct for TargetApplication
type jsiiProxy_TargetApplication struct {
	_ byte // padding
}

// Experimental.
func NewTargetApplication_Override(t TargetApplication) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.TargetApplication",
		nil, // no parameters
		t,
	)
}

// Factory method to build the input using the provided application name and stack props.
// Experimental.
func TargetApplication_CreateApplicationStack(options *CreateTargetApplicationOptions) TargetApplication {
	_init_.Initialize()

	if err := validateTargetApplication_CreateApplicationStackParameters(options); err != nil {
		panic(err)
	}
	var returns TargetApplication

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.TargetApplication",
		"createApplicationStack",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Factory method to build the input using the provided application ARN.
// Experimental.
func TargetApplication_ExistingApplicationFromArn(options *ExistingTargetApplicationOptions) TargetApplication {
	_init_.Initialize()

	if err := validateTargetApplication_ExistingApplicationFromArnParameters(options); err != nil {
		panic(err)
	}
	var returns TargetApplication

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.TargetApplication",
		"existingApplicationFromArn",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TargetApplication) Bind(scope constructs.Construct) *BindTargetApplicationResult {
	if err := t.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *BindTargetApplicationResult

	_jsii_.Invoke(
		t,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}
