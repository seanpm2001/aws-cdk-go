package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
)

// Attributes for EKS clusters.
//
// Example:
//   var cluster cluster
//   var asg autoScalingGroup
//
//   importedCluster := eks.cluster.fromClusterAttributes(this, jsii.String("ImportedCluster"), &clusterAttributes{
//   	clusterName: cluster.clusterName,
//   	clusterSecurityGroupId: cluster.clusterSecurityGroupId,
//   })
//
//   importedCluster.connectAutoScalingGroupCapacity(asg, &autoScalingGroupOptions{
//   })
//
// Experimental.
type ClusterAttributes struct {
	// The physical name of the Cluster.
	// Experimental.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The certificate-authority-data for your cluster.
	// Experimental.
	ClusterCertificateAuthorityData *string `field:"optional" json:"clusterCertificateAuthorityData" yaml:"clusterCertificateAuthorityData"`
	// Amazon Resource Name (ARN) or alias of the customer master key (CMK).
	// Experimental.
	ClusterEncryptionConfigKeyArn *string `field:"optional" json:"clusterEncryptionConfigKeyArn" yaml:"clusterEncryptionConfigKeyArn"`
	// The API Server endpoint URL.
	// Experimental.
	ClusterEndpoint *string `field:"optional" json:"clusterEndpoint" yaml:"clusterEndpoint"`
	// A security group id to associate with the Cluster Handler's Lambdas.
	//
	// The Cluster Handler's Lambdas are responsible for calling AWS's EKS API.
	// Experimental.
	ClusterHandlerSecurityGroupId *string `field:"optional" json:"clusterHandlerSecurityGroupId" yaml:"clusterHandlerSecurityGroupId"`
	// The cluster security group that was created by Amazon EKS for the cluster.
	// Experimental.
	ClusterSecurityGroupId *string `field:"optional" json:"clusterSecurityGroupId" yaml:"clusterSecurityGroupId"`
	// Environment variables to use when running `kubectl` against this cluster.
	// Experimental.
	KubectlEnvironment *map[string]*string `field:"optional" json:"kubectlEnvironment" yaml:"kubectlEnvironment"`
	// An IAM role that can perform kubectl operations against this cluster.
	//
	// The role should be mapped to the `system:masters` Kubernetes RBAC role.
	//
	// This role is directly passed to the lambda handler that sends Kube Ctl commands
	// to the cluster.
	// Experimental.
	KubectlLambdaRole awsiam.IRole `field:"optional" json:"kubectlLambdaRole" yaml:"kubectlLambdaRole"`
	// An AWS Lambda Layer which includes `kubectl`, Helm and the AWS CLI.
	//
	// This layer
	// is used by the kubectl handler to apply manifests and install helm charts.
	//
	// The handler expects the layer to include the following executables:
	//
	//     helm/helm
	//     kubectl/kubectl
	// awscli/aws.
	// Experimental.
	KubectlLayer awslambda.ILayerVersion `field:"optional" json:"kubectlLayer" yaml:"kubectlLayer"`
	// Amount of memory to allocate to the provider's lambda function.
	// Experimental.
	KubectlMemory awscdk.Size `field:"optional" json:"kubectlMemory" yaml:"kubectlMemory"`
	// Subnets to host the `kubectl` compute resources.
	//
	// If not specified, the k8s
	// endpoint is expected to be accessible publicly.
	// Experimental.
	KubectlPrivateSubnetIds *[]*string `field:"optional" json:"kubectlPrivateSubnetIds" yaml:"kubectlPrivateSubnetIds"`
	// KubectlProvider for issuing kubectl commands.
	// Experimental.
	KubectlProvider IKubectlProvider `field:"optional" json:"kubectlProvider" yaml:"kubectlProvider"`
	// An IAM role with cluster administrator and "system:masters" permissions.
	// Experimental.
	KubectlRoleArn *string `field:"optional" json:"kubectlRoleArn" yaml:"kubectlRoleArn"`
	// A security group to use for `kubectl` execution.
	//
	// If not specified, the k8s
	// endpoint is expected to be accessible publicly.
	// Experimental.
	KubectlSecurityGroupId *string `field:"optional" json:"kubectlSecurityGroupId" yaml:"kubectlSecurityGroupId"`
	// An AWS Lambda Layer which includes the NPM dependency `proxy-agent`.
	//
	// This layer
	// is used by the onEvent handler to route AWS SDK requests through a proxy.
	//
	// The handler expects the layer to include the following node_modules:
	//
	// proxy-agent.
	// Experimental.
	OnEventLayer awslambda.ILayerVersion `field:"optional" json:"onEventLayer" yaml:"onEventLayer"`
	// An Open ID Connect provider for this cluster that can be used to configure service accounts.
	//
	// You can either import an existing provider using `iam.OpenIdConnectProvider.fromProviderArn`,
	// or create a new provider using `new eks.OpenIdConnectProvider`
	// Experimental.
	OpenIdConnectProvider awsiam.IOpenIdConnectProvider `field:"optional" json:"openIdConnectProvider" yaml:"openIdConnectProvider"`
	// Indicates whether Kubernetes resources added through `addManifest()` can be automatically pruned.
	//
	// When this is enabled (default), prune labels will be
	// allocated and injected to each resource. These labels will then be used
	// when issuing the `kubectl apply` operation with the `--prune` switch.
	// Experimental.
	Prune *bool `field:"optional" json:"prune" yaml:"prune"`
	// Additional security groups associated with this cluster.
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The VPC in which this Cluster was created.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

