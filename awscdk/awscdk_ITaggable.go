// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface to implement tags.
// Experimental.
type ITaggable interface {
	// TagManager to set, remove and format tags.
	// Experimental.
	Tags() TagManager
}

// The jsii proxy for ITaggable
type jsiiProxy_ITaggable struct {
	_ byte // padding
}

func (j *jsiiProxy_ITaggable) Tags() TagManager {
	var returns TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

