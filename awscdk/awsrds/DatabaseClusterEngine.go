package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A database cluster engine.
//
// Provides mapping to the serverless application
// used for secret rotation.
//
// Example:
//   var vpc vpc
//
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
//   		Version: rds.AuroraMysqlEngineVersion_VER_2_08_1(),
//   	}),
//   	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer"), &ProvisionedClusterInstanceProps{
//   		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_R6G, ec2.InstanceSize_XLARGE4),
//   	}),
//   	ServerlessV2MinCapacity: jsii.Number(6.5),
//   	ServerlessV2MaxCapacity: jsii.Number(64),
//   	Readers: []iClusterInstance{
//   		rds.ClusterInstance_ServerlessV2(jsii.String("reader1"), &ServerlessV2ClusterInstanceProps{
//   			ScaleWithWriter: jsii.Boolean(true),
//   		}),
//   		rds.ClusterInstance_*ServerlessV2(jsii.String("reader2")),
//   	},
//   	Vpc: Vpc,
//   })
//
type DatabaseClusterEngine interface {
}

// The jsii proxy struct for DatabaseClusterEngine
type jsiiProxy_DatabaseClusterEngine struct {
	_ byte // padding
}

func NewDatabaseClusterEngine() DatabaseClusterEngine {
	_init_.Initialize()

	j := jsiiProxy_DatabaseClusterEngine{}

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseClusterEngine",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewDatabaseClusterEngine_Override(d DatabaseClusterEngine) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.DatabaseClusterEngine",
		nil, // no parameters
		d,
	)
}

// Creates a new plain Aurora database cluster engine.
func DatabaseClusterEngine_Aurora(props *AuroraClusterEngineProps) IClusterEngine {
	_init_.Initialize()

	if err := validateDatabaseClusterEngine_AuroraParameters(props); err != nil {
		panic(err)
	}
	var returns IClusterEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseClusterEngine",
		"aurora",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new Aurora MySQL database cluster engine.
func DatabaseClusterEngine_AuroraMysql(props *AuroraMysqlClusterEngineProps) IClusterEngine {
	_init_.Initialize()

	if err := validateDatabaseClusterEngine_AuroraMysqlParameters(props); err != nil {
		panic(err)
	}
	var returns IClusterEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseClusterEngine",
		"auroraMysql",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Creates a new Aurora PostgreSQL database cluster engine.
func DatabaseClusterEngine_AuroraPostgres(props *AuroraPostgresClusterEngineProps) IClusterEngine {
	_init_.Initialize()

	if err := validateDatabaseClusterEngine_AuroraPostgresParameters(props); err != nil {
		panic(err)
	}
	var returns IClusterEngine

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseClusterEngine",
		"auroraPostgres",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func DatabaseClusterEngine_AURORA() IClusterEngine {
	_init_.Initialize()
	var returns IClusterEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseClusterEngine",
		"AURORA",
		&returns,
	)
	return returns
}

func DatabaseClusterEngine_AURORA_MYSQL() IClusterEngine {
	_init_.Initialize()
	var returns IClusterEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseClusterEngine",
		"AURORA_MYSQL",
		&returns,
	)
	return returns
}

func DatabaseClusterEngine_AURORA_POSTGRESQL() IClusterEngine {
	_init_.Initialize()
	var returns IClusterEngine
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.DatabaseClusterEngine",
		"AURORA_POSTGRESQL",
		&returns,
	)
	return returns
}

