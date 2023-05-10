package awsmsk


// The settings for encrypting data in transit.
//
// Example:
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &ClusterProps{
//   	ClusterName: jsii.String("myCluster"),
//   	KafkaVersion: msk.KafkaVersion_V2_8_1(),
//   	Vpc: Vpc,
//   	EncryptionInTransit: &EncryptionInTransitConfig{
//   		ClientBroker: msk.ClientBrokerEncryption_TLS,
//   	},
//   	ClientAuthentication: msk.ClientAuthentication_Sasl(&SaslAuthProps{
//   		Scram: jsii.Boolean(true),
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/msk/latest/developerguide/msk-encryption.html#msk-encryption-in-transit
//
// Experimental.
type EncryptionInTransitConfig struct {
	// Indicates the encryption setting for data in transit between clients and brokers.
	// Experimental.
	ClientBroker ClientBrokerEncryption `field:"optional" json:"clientBroker" yaml:"clientBroker"`
	// Indicates that data communication among the broker nodes of the cluster is encrypted.
	// Experimental.
	EnableInCluster *bool `field:"optional" json:"enableInCluster" yaml:"enableInCluster"`
}

