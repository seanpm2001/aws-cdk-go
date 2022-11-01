package awselasticloadbalancingv2actions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2actions/internal"
)

// A Listener Action to authenticate with Cognito.
//
// Example:
//   import cognito "github.com/aws/aws-cdk-go/awscdk"
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//   import elbv2 "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/constructs-go/constructs"
//   import actions "github.com/aws/aws-cdk-go/awscdk"
//
//   cognitoStack struct {
//   stack
//   }
//
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   })
//
//   userPool := cognito.NewUserPool(this, jsii.String("UserPool"))
//   userPoolClient := cognito.NewUserPoolClient(this, jsii.String("Client"), &userPoolClientProps{
//   	userPool: userPool,
//
//   	// Required minimal configuration for use with an ELB
//   	generateSecret: jsii.Boolean(true),
//   	authFlows: &authFlow{
//   		userPassword: jsii.Boolean(true),
//   	},
//   	oAuth: &oAuthSettings{
//   		flows: &oAuthFlows{
//   			authorizationCodeGrant: jsii.Boolean(true),
//   		},
//   		scopes: []oAuthScope{
//   			cognito.*oAuthScope_EMAIL(),
//   		},
//   		callbackUrls: []*string{
//   			fmt.Sprintf("https://%v/oauth2/idpresponse", lb.loadBalancerDnsName),
//   		},
//   	},
//   })
//   cfnClient := userPoolClient.node.defaultChild.(cfnUserPoolClient)
//   cfnClient.addPropertyOverride(jsii.String("RefreshTokenValidity"), jsii.Number(1))
//   cfnClient.addPropertyOverride(jsii.String("SupportedIdentityProviders"), []interface{}{
//   	jsii.String("COGNITO"),
//   })
//
//   userPoolDomain := cognito.NewUserPoolDomain(this, jsii.String("Domain"), &userPoolDomainProps{
//   	userPool: userPool,
//   	cognitoDomain: &cognitoDomainOptions{
//   		domainPrefix: jsii.String("test-cdk-prefix"),
//   	},
//   })
//
//   lb.addListener(jsii.String("Listener"), &baseApplicationListenerProps{
//   	port: jsii.Number(443),
//   	certificates: []iListenerCertificate{
//   		certificate,
//   	},
//   	defaultAction: actions.NewAuthenticateCognitoAction(&authenticateCognitoActionProps{
//   		userPool: userPool,
//   		userPoolClient: userPoolClient,
//   		userPoolDomain: userPoolDomain,
//   		next: elbv2.listenerAction.fixedResponse(jsii.Number(200), &fixedResponseOptions{
//   			contentType: jsii.String("text/plain"),
//   			messageBody: jsii.String("Authenticated"),
//   		}),
//   	}),
//   })
//
//   awscdk.NewCfnOutput(this, jsii.String("DNS"), &cfnOutputProps{
//   	value: lb.loadBalancerDnsName,
//   })
//
//   app := awscdk.NewApp()
//   NewCognitoStack(app, jsii.String("integ-cognito"))
//   app.synth()
//
// Experimental.
type AuthenticateCognitoAction interface {
	awselasticloadbalancingv2.ListenerAction
	// Experimental.
	Next() awselasticloadbalancingv2.ListenerAction
	// Called when the action is being used in a listener.
	// Experimental.
	Bind(scope awscdk.Construct, listener awselasticloadbalancingv2.IApplicationListener, associatingConstruct awscdk.IConstruct)
	// Render the actions in this chain.
	// Experimental.
	RenderActions() *[]*awselasticloadbalancingv2.CfnListener_ActionProperty
	// Renumber the "order" fields in the actions array.
	//
	// We don't number for 0 or 1 elements, but otherwise number them 1...#actions
	// so ELB knows about the right order.
	//
	// Do this in `ListenerAction` instead of in `Listener` so that we give
	// users the opportunity to override by subclassing and overriding `renderActions`.
	// Experimental.
	Renumber(actions *[]*awselasticloadbalancingv2.CfnListener_ActionProperty) *[]*awselasticloadbalancingv2.CfnListener_ActionProperty
}

// The jsii proxy struct for AuthenticateCognitoAction
type jsiiProxy_AuthenticateCognitoAction struct {
	internal.Type__awselasticloadbalancingv2ListenerAction
}

func (j *jsiiProxy_AuthenticateCognitoAction) Next() awselasticloadbalancingv2.ListenerAction {
	var returns awselasticloadbalancingv2.ListenerAction
	_jsii_.Get(
		j,
		"next",
		&returns,
	)
	return returns
}


// Authenticate using an identity provide (IdP) that is compliant with OpenID Connect (OIDC).
// Experimental.
func NewAuthenticateCognitoAction(options *AuthenticateCognitoActionProps) AuthenticateCognitoAction {
	_init_.Initialize()

	if err := validateNewAuthenticateCognitoActionParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AuthenticateCognitoAction{}

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2_actions.AuthenticateCognitoAction",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Authenticate using an identity provide (IdP) that is compliant with OpenID Connect (OIDC).
// Experimental.
func NewAuthenticateCognitoAction_Override(a AuthenticateCognitoAction, options *AuthenticateCognitoActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2_actions.AuthenticateCognitoAction",
		[]interface{}{options},
		a,
	)
}

// Authenticate using an identity provider (IdP) that is compliant with OpenID Connect (OIDC).
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/listener-authenticate-users.html#oidc-requirements
//
// Experimental.
func AuthenticateCognitoAction_AuthenticateOidc(options *awselasticloadbalancingv2.AuthenticateOidcOptions) awselasticloadbalancingv2.ListenerAction {
	_init_.Initialize()

	if err := validateAuthenticateCognitoAction_AuthenticateOidcParameters(options); err != nil {
		panic(err)
	}
	var returns awselasticloadbalancingv2.ListenerAction

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancingv2_actions.AuthenticateCognitoAction",
		"authenticateOidc",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Return a fixed response.
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#fixed-response-actions
//
// Experimental.
func AuthenticateCognitoAction_FixedResponse(statusCode *float64, options *awselasticloadbalancingv2.FixedResponseOptions) awselasticloadbalancingv2.ListenerAction {
	_init_.Initialize()

	if err := validateAuthenticateCognitoAction_FixedResponseParameters(statusCode, options); err != nil {
		panic(err)
	}
	var returns awselasticloadbalancingv2.ListenerAction

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancingv2_actions.AuthenticateCognitoAction",
		"fixedResponse",
		[]interface{}{statusCode, options},
		&returns,
	)

	return returns
}

// Forward to one or more Target Groups.
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#forward-actions
//
// Experimental.
func AuthenticateCognitoAction_Forward(targetGroups *[]awselasticloadbalancingv2.IApplicationTargetGroup, options *awselasticloadbalancingv2.ForwardOptions) awselasticloadbalancingv2.ListenerAction {
	_init_.Initialize()

	if err := validateAuthenticateCognitoAction_ForwardParameters(targetGroups, options); err != nil {
		panic(err)
	}
	var returns awselasticloadbalancingv2.ListenerAction

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancingv2_actions.AuthenticateCognitoAction",
		"forward",
		[]interface{}{targetGroups, options},
		&returns,
	)

	return returns
}

// Redirect to a different URI.
//
// A URI consists of the following components:
// protocol://hostname:port/path?query. You must modify at least one of the
// following components to avoid a redirect loop: protocol, hostname, port, or
// path. Any components that you do not modify retain their original values.
//
// You can reuse URI components using the following reserved keywords:
//
// - `#{protocol}`
// - `#{host}`
// - `#{port}`
// - `#{path}` (the leading "/" is removed)
// - `#{query}`
//
// For example, you can change the path to "/new/#{path}", the hostname to
// "example.#{host}", or the query to "#{query}&value=xyz".
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#redirect-actions
//
// Experimental.
func AuthenticateCognitoAction_Redirect(options *awselasticloadbalancingv2.RedirectOptions) awselasticloadbalancingv2.ListenerAction {
	_init_.Initialize()

	if err := validateAuthenticateCognitoAction_RedirectParameters(options); err != nil {
		panic(err)
	}
	var returns awselasticloadbalancingv2.ListenerAction

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancingv2_actions.AuthenticateCognitoAction",
		"redirect",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Forward to one or more Target Groups which are weighted differently.
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html#forward-actions
//
// Experimental.
func AuthenticateCognitoAction_WeightedForward(targetGroups *[]*awselasticloadbalancingv2.WeightedTargetGroup, options *awselasticloadbalancingv2.ForwardOptions) awselasticloadbalancingv2.ListenerAction {
	_init_.Initialize()

	if err := validateAuthenticateCognitoAction_WeightedForwardParameters(targetGroups, options); err != nil {
		panic(err)
	}
	var returns awselasticloadbalancingv2.ListenerAction

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancingv2_actions.AuthenticateCognitoAction",
		"weightedForward",
		[]interface{}{targetGroups, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthenticateCognitoAction) Bind(scope awscdk.Construct, listener awselasticloadbalancingv2.IApplicationListener, associatingConstruct awscdk.IConstruct) {
	if err := a.validateBindParameters(scope, listener); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"bind",
		[]interface{}{scope, listener, associatingConstruct},
	)
}

func (a *jsiiProxy_AuthenticateCognitoAction) RenderActions() *[]*awselasticloadbalancingv2.CfnListener_ActionProperty {
	var returns *[]*awselasticloadbalancingv2.CfnListener_ActionProperty

	_jsii_.Invoke(
		a,
		"renderActions",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuthenticateCognitoAction) Renumber(actions *[]*awselasticloadbalancingv2.CfnListener_ActionProperty) *[]*awselasticloadbalancingv2.CfnListener_ActionProperty {
	if err := a.validateRenumberParameters(actions); err != nil {
		panic(err)
	}
	var returns *[]*awselasticloadbalancingv2.CfnListener_ActionProperty

	_jsii_.Invoke(
		a,
		"renumber",
		[]interface{}{actions},
		&returns,
	)

	return returns
}

