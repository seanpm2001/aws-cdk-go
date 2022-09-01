package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Configuration props for EKS Fargate.
//
// Example:
//   cluster := eks.NewFargateCluster(this, jsii.String("MyCluster"), &fargateClusterProps{
//   	version: eks.kubernetesVersion_V1_21(),
//   })
//
type FargateClusterProps struct {
	// The Kubernetes version to run in the cluster.
	Version KubernetesVersion `field:"required" json:"version" yaml:"version"`
	// Name for the cluster.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Determines whether a CloudFormation output with the name of the cluster will be synthesized.
	OutputClusterName *bool `field:"optional" json:"outputClusterName" yaml:"outputClusterName"`
	// Determines whether a CloudFormation output with the `aws eks update-kubeconfig` command will be synthesized.
	//
	// This command will include
	// the cluster name and, if applicable, the ARN of the masters IAM role.
	OutputConfigCommand *bool `field:"optional" json:"outputConfigCommand" yaml:"outputConfigCommand"`
	// Role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Security Group to use for Control Plane ENIs.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// The VPC in which to create the Cluster.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where to place EKS Control Plane ENIs.
	//
	// If you want to create public load balancers, this must include public subnets.
	//
	// For example, to only select private subnets, supply the following:
	//
	// `vpcSubnets: [{ subnetType: ec2.SubnetType.PRIVATE_WITH_NAT }]`
	VpcSubnets *[]*awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// Install the AWS Load Balancer Controller onto the cluster.
	// See: https://kubernetes-sigs.github.io/aws-load-balancer-controller
	//
	AlbController *AlbControllerOptions `field:"optional" json:"albController" yaml:"albController"`
	// Custom environment variables when interacting with the EKS endpoint to manage the cluster lifecycle.
	ClusterHandlerEnvironment *map[string]*string `field:"optional" json:"clusterHandlerEnvironment" yaml:"clusterHandlerEnvironment"`
	// A security group to associate with the Cluster Handler's Lambdas.
	//
	// The Cluster Handler's Lambdas are responsible for calling AWS's EKS API.
	//
	// Requires `placeClusterHandlerInVpc` to be set to true.
	ClusterHandlerSecurityGroup awsec2.ISecurityGroup `field:"optional" json:"clusterHandlerSecurityGroup" yaml:"clusterHandlerSecurityGroup"`
	// The cluster log types which you want to enable.
	ClusterLogging *[]ClusterLoggingTypes `field:"optional" json:"clusterLogging" yaml:"clusterLogging"`
	// Controls the "eks.amazonaws.com/compute-type" annotation in the CoreDNS configuration on your cluster to determine which compute type to use for CoreDNS.
	CoreDnsComputeType CoreDnsComputeType `field:"optional" json:"coreDnsComputeType" yaml:"coreDnsComputeType"`
	// Configure access to the Kubernetes API server endpoint..
	// See: https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html
	//
	EndpointAccess EndpointAccess `field:"optional" json:"endpointAccess" yaml:"endpointAccess"`
	// Environment variables for the kubectl execution.
	//
	// Only relevant for kubectl enabled clusters.
	KubectlEnvironment *map[string]*string `field:"optional" json:"kubectlEnvironment" yaml:"kubectlEnvironment"`
	// An AWS Lambda Layer which includes `kubectl`, Helm and the AWS CLI.
	//
	// By default, the provider will use the layer included in the
	// "aws-lambda-layer-kubectl" SAR application which is available in all
	// commercial regions.
	//
	// To deploy the layer locally, visit
	// https://github.com/aws-samples/aws-lambda-layer-kubectl/blob/master/cdk/README.md
	// for instructions on how to prepare the .zip file and then define it in your
	// app as follows:
	//
	// ```ts
	// const layer = new lambda.LayerVersion(this, 'kubectl-layer', {
	//    code: lambda.Code.fromAsset(`${__dirname}/layer.zip`),
	//    compatibleRuntimes: [lambda.Runtime.PROVIDED],
	// });
	// ```.
	// See: https://github.com/aws-samples/aws-lambda-layer-kubectl
	//
	KubectlLayer awslambda.ILayerVersion `field:"optional" json:"kubectlLayer" yaml:"kubectlLayer"`
	// Amount of memory to allocate to the provider's lambda function.
	KubectlMemory awscdk.Size `field:"optional" json:"kubectlMemory" yaml:"kubectlMemory"`
	// An IAM role that will be added to the `system:masters` Kubernetes RBAC group.
	// See: https://kubernetes.io/docs/reference/access-authn-authz/rbac/#default-roles-and-role-bindings
	//
	MastersRole awsiam.IRole `field:"optional" json:"mastersRole" yaml:"mastersRole"`
	// An AWS Lambda Layer which includes the NPM dependency `proxy-agent`.
	//
	// This layer
	// is used by the onEvent handler to route AWS SDK requests through a proxy.
	//
	// By default, the provider will use the layer included in the
	// "aws-lambda-layer-node-proxy-agent" SAR application which is available in all
	// commercial regions.
	//
	// To deploy the layer locally define it in your app as follows:
	//
	// ```ts
	// const layer = new lambda.LayerVersion(this, 'proxy-agent-layer', {
	//    code: lambda.Code.fromAsset(`${__dirname}/layer.zip`),
	//    compatibleRuntimes: [lambda.Runtime.NODEJS_14_X],
	// });
	// ```.
	OnEventLayer awslambda.ILayerVersion `field:"optional" json:"onEventLayer" yaml:"onEventLayer"`
	// Determines whether a CloudFormation output with the ARN of the "masters" IAM role will be synthesized (if `mastersRole` is specified).
	OutputMastersRoleArn *bool `field:"optional" json:"outputMastersRoleArn" yaml:"outputMastersRoleArn"`
	// If set to true, the cluster handler functions will be placed in the private subnets of the cluster vpc, subject to the `vpcSubnets` selection strategy.
	PlaceClusterHandlerInVpc *bool `field:"optional" json:"placeClusterHandlerInVpc" yaml:"placeClusterHandlerInVpc"`
	// Indicates whether Kubernetes resources added through `addManifest()` can be automatically pruned.
	//
	// When this is enabled (default), prune labels will be
	// allocated and injected to each resource. These labels will then be used
	// when issuing the `kubectl apply` operation with the `--prune` switch.
	Prune *bool `field:"optional" json:"prune" yaml:"prune"`
	// KMS secret for envelope encryption for Kubernetes secrets.
	SecretsEncryptionKey awskms.IKey `field:"optional" json:"secretsEncryptionKey" yaml:"secretsEncryptionKey"`
	// The CIDR block to assign Kubernetes service IP addresses from.
	// See: https://docs.aws.amazon.com/eks/latest/APIReference/API_KubernetesNetworkConfigRequest.html#AmazonEKS-Type-KubernetesNetworkConfigRequest-serviceIpv4Cidr
	//
	ServiceIpv4Cidr *string `field:"optional" json:"serviceIpv4Cidr" yaml:"serviceIpv4Cidr"`
	// Fargate Profile to create along with the cluster.
	DefaultProfile *FargateProfileOptions `field:"optional" json:"defaultProfile" yaml:"defaultProfile"`
}

