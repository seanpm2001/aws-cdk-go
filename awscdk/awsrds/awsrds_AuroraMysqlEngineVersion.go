package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The versions for the Aurora MySQL cluster engine (those returned by {@link DatabaseClusterEngine.auroraMysql}).
//
// Example:
//   var vpc vpc
//
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
//   	engine: rds.databaseClusterEngine.auroraMysql(&auroraMysqlClusterEngineProps{
//   		version: rds.auroraMysqlEngineVersion_VER_2_08_1(),
//   	}),
//   	credentials: rds.credentials.fromGeneratedSecret(jsii.String("clusteradmin")),
//   	 // Optional - will default to 'admin' username and generated password
//   	instanceProps: &instanceProps{
//   		// optional , defaults to t3.medium
//   		instanceType: ec2.instanceType.of(ec2.instanceClass_BURSTABLE2, ec2.instanceSize_SMALL),
//   		vpcSubnets: &subnetSelection{
//   			subnetType: ec2.subnetType_PRIVATE_WITH_NAT,
//   		},
//   		vpc: vpc,
//   	},
//   })
//
// Experimental.
type AuroraMysqlEngineVersion interface {
	// The full version string, for example, "5.7.mysql_aurora.1.78.3.6".
	// Experimental.
	AuroraMysqlFullVersion() *string
	// The major version of the engine.
	//
	// Currently, it's either "5.7", or "8.0".
	// Experimental.
	AuroraMysqlMajorVersion() *string
}

// The jsii proxy struct for AuroraMysqlEngineVersion
type jsiiProxy_AuroraMysqlEngineVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_AuroraMysqlEngineVersion) AuroraMysqlFullVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auroraMysqlFullVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuroraMysqlEngineVersion) AuroraMysqlMajorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auroraMysqlMajorVersion",
		&returns,
	)
	return returns
}


// Create a new AuroraMysqlEngineVersion with an arbitrary version.
// Experimental.
func AuroraMysqlEngineVersion_Of(auroraMysqlFullVersion *string, auroraMysqlMajorVersion *string) AuroraMysqlEngineVersion {
	_init_.Initialize()

	if err := validateAuroraMysqlEngineVersion_OfParameters(auroraMysqlFullVersion); err != nil {
		panic(err)
	}
	var returns AuroraMysqlEngineVersion

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"of",
		[]interface{}{auroraMysqlFullVersion, auroraMysqlMajorVersion},
		&returns,
	)

	return returns
}

func AuroraMysqlEngineVersion_VER_2_03_2() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_03_2",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_03_3() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_03_3",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_03_4() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_03_4",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_04_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_04_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_04_1() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_04_1",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_04_2() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_04_2",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_04_3() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_04_3",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_04_4() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_04_4",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_04_5() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_04_5",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_04_6() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_04_6",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_04_7() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_04_7",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_04_8() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_04_8",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_05_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_05_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_06_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_06_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_07_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_07_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_07_1() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_07_1",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_07_2() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_07_2",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_08_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_08_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_08_1() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_08_1",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_08_2() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_08_2",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_09_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_09_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_09_1() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_09_1",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_09_2() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_09_2",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_09_3() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_09_3",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_10_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_10_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_10_1() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_10_1",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_2_10_2() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_2_10_2",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_3_01_0() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_3_01_0",
		&returns,
	)
	return returns
}

func AuroraMysqlEngineVersion_VER_5_7_12() AuroraMysqlEngineVersion {
	_init_.Initialize()
	var returns AuroraMysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.AuroraMysqlEngineVersion",
		"VER_5_7_12",
		&returns,
	)
	return returns
}
