//go:build !no_runtime_type_checking

package awsappsync

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/awselasticsearch"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awsopensearchservice"
	"github.com/aws/aws-cdk-go/awscdk/awsrds"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v3"
)

func (g *jsiiProxy_GraphqlApi) validateAddDynamoDbDataSourceParameters(id *string, table awsdynamodb.ITable, options *DataSourceOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if table == nil {
		return fmt.Errorf("parameter table is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_GraphqlApi) validateAddElasticsearchDataSourceParameters(id *string, domain awselasticsearch.IDomain, options *DataSourceOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if domain == nil {
		return fmt.Errorf("parameter domain is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_GraphqlApi) validateAddHttpDataSourceParameters(id *string, endpoint *string, options *HttpDataSourceOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if endpoint == nil {
		return fmt.Errorf("parameter endpoint is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_GraphqlApi) validateAddLambdaDataSourceParameters(id *string, lambdaFunction awslambda.IFunction, options *DataSourceOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if lambdaFunction == nil {
		return fmt.Errorf("parameter lambdaFunction is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_GraphqlApi) validateAddMutationParameters(fieldName *string, field ResolvableField) error {
	if fieldName == nil {
		return fmt.Errorf("parameter fieldName is required, but nil was provided")
	}

	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GraphqlApi) validateAddNoneDataSourceParameters(id *string, options *DataSourceOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_GraphqlApi) validateAddOpenSearchDataSourceParameters(id *string, domain awsopensearchservice.IDomain, options *DataSourceOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if domain == nil {
		return fmt.Errorf("parameter domain is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_GraphqlApi) validateAddQueryParameters(fieldName *string, field ResolvableField) error {
	if fieldName == nil {
		return fmt.Errorf("parameter fieldName is required, but nil was provided")
	}

	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GraphqlApi) validateAddRdsDataSourceParameters(id *string, serverlessCluster awsrds.IServerlessCluster, secretStore awssecretsmanager.ISecret, options *DataSourceOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if serverlessCluster == nil {
		return fmt.Errorf("parameter serverlessCluster is required, but nil was provided")
	}

	if secretStore == nil {
		return fmt.Errorf("parameter secretStore is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_GraphqlApi) validateAddSchemaDependencyParameters(construct awscdk.CfnResource) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GraphqlApi) validateAddSubscriptionParameters(fieldName *string, field ResolvableField) error {
	if fieldName == nil {
		return fmt.Errorf("parameter fieldName is required, but nil was provided")
	}

	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GraphqlApi) validateAddToSchemaParameters(addition *string) error {
	if addition == nil {
		return fmt.Errorf("parameter addition is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GraphqlApi) validateAddTypeParameters(type_ IIntermediateType) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GraphqlApi) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GraphqlApi) validateCreateResolverParameters(props *ExtendedResolverProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_GraphqlApi) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	if arnAttr == nil {
		return fmt.Errorf("parameter arnAttr is required, but nil was provided")
	}

	if arnComponents == nil {
		return fmt.Errorf("parameter arnComponents is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(arnComponents, func() string { return "parameter arnComponents" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_GraphqlApi) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GraphqlApi) validateGrantParameters(grantee awsiam.IGrantable, resources IamResource) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	if resources == nil {
		return fmt.Errorf("parameter resources is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GraphqlApi) validateGrantMutationParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GraphqlApi) validateGrantQueryParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GraphqlApi) validateGrantSubscriptionParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GraphqlApi) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GraphqlApi) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func validateGraphqlApi_FromGraphqlApiAttributesParameters(scope constructs.Construct, id *string, attrs *GraphqlApiAttributes) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if attrs == nil {
		return fmt.Errorf("parameter attrs is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(attrs, func() string { return "parameter attrs" }); err != nil {
		return err
	}

	return nil
}

func validateGraphqlApi_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateGraphqlApi_IsResourceParameters(construct awscdk.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewGraphqlApiParameters(scope constructs.Construct, id *string, props *GraphqlApiProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

