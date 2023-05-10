package awsservicediscovery


// Experimental.
type NamespaceType string

const (
	// Choose this option if you want your application to use only API calls to discover registered instances.
	// Experimental.
	NamespaceType_HTTP NamespaceType = "HTTP"
	// Choose this option if you want your application to be able to discover instances using either API calls or using DNS queries in a VPC.
	// Experimental.
	NamespaceType_DNS_PRIVATE NamespaceType = "DNS_PRIVATE"
	// Choose this option if you want your application to be able to discover instances using either API calls or using public DNS queries.
	//
	// You aren't required to use both methods.
	// Experimental.
	NamespaceType_DNS_PUBLIC NamespaceType = "DNS_PUBLIC"
)

