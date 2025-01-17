package awscdkappconfigalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Defines the extensible base implementation for extension association resources.
// Experimental.
type IExtensible interface {
	// Adds an extension association to the derived resource.
	// Experimental.
	AddExtension(extension IExtension)
	// Adds an extension defined by the action point and event destination and also creates an extension association to the derived resource.
	// Experimental.
	On(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_BAKING extension with the provided event destination and also creates an extension association to the derived resource.
	// Experimental.
	OnDeploymentBaking(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_COMPLETE extension with the provided event destination and also creates an extension association to the derived resource.
	// Experimental.
	OnDeploymentComplete(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_ROLLED_BACK extension with the provided event destination and also creates an extension association to the derived resource.
	// Experimental.
	OnDeploymentRolledBack(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_START extension with the provided event destination and also creates an extension association to the derived resource.
	// Experimental.
	OnDeploymentStart(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_STEP extension with the provided event destination and also creates an extension association to the derived resource.
	// Experimental.
	OnDeploymentStep(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds a PRE_CREATE_HOSTED_CONFIGURATION_VERSION extension with the provided event destination and also creates an extension association to the derived resource.
	// Experimental.
	PreCreateHostedConfigurationVersion(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds a PRE_START_DEPLOYMENT extension with the provided event destination and also creates an extension association to the derived resource.
	// Experimental.
	PreStartDeployment(eventDestination IEventDestination, options *ExtensionOptions)
}

// The jsii proxy for IExtensible
type jsiiProxy_IExtensible struct {
	_ byte // padding
}

func (i *jsiiProxy_IExtensible) AddExtension(extension IExtension) {
	if err := i.validateAddExtensionParameters(extension); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addExtension",
		[]interface{}{extension},
	)
}

func (i *jsiiProxy_IExtensible) On(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validateOnParameters(actionPoint, eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"on",
		[]interface{}{actionPoint, eventDestination, options},
	)
}

func (i *jsiiProxy_IExtensible) OnDeploymentBaking(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validateOnDeploymentBakingParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onDeploymentBaking",
		[]interface{}{eventDestination, options},
	)
}

func (i *jsiiProxy_IExtensible) OnDeploymentComplete(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validateOnDeploymentCompleteParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onDeploymentComplete",
		[]interface{}{eventDestination, options},
	)
}

func (i *jsiiProxy_IExtensible) OnDeploymentRolledBack(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validateOnDeploymentRolledBackParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onDeploymentRolledBack",
		[]interface{}{eventDestination, options},
	)
}

func (i *jsiiProxy_IExtensible) OnDeploymentStart(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validateOnDeploymentStartParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onDeploymentStart",
		[]interface{}{eventDestination, options},
	)
}

func (i *jsiiProxy_IExtensible) OnDeploymentStep(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validateOnDeploymentStepParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onDeploymentStep",
		[]interface{}{eventDestination, options},
	)
}

func (i *jsiiProxy_IExtensible) PreCreateHostedConfigurationVersion(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validatePreCreateHostedConfigurationVersionParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"preCreateHostedConfigurationVersion",
		[]interface{}{eventDestination, options},
	)
}

func (i *jsiiProxy_IExtensible) PreStartDeployment(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := i.validatePreStartDeploymentParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"preStartDeployment",
		[]interface{}{eventDestination, options},
	)
}

