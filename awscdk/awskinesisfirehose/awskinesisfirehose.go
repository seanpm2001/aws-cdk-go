package awskinesisfirehose

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream",
		reflect.TypeOf((*CfnDeliveryStream)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "amazonOpenSearchServerlessDestinationConfiguration", GoGetter: "AmazonOpenSearchServerlessDestinationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "amazonopensearchserviceDestinationConfiguration", GoGetter: "AmazonopensearchserviceDestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamEncryptionConfigurationInput", GoGetter: "DeliveryStreamEncryptionConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamName", GoGetter: "DeliveryStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamType", GoGetter: "DeliveryStreamType"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchDestinationConfiguration", GoGetter: "ElasticsearchDestinationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "extendedS3DestinationConfiguration", GoGetter: "ExtendedS3DestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "httpEndpointDestinationConfiguration", GoGetter: "HttpEndpointDestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "kinesisStreamSourceConfiguration", GoGetter: "KinesisStreamSourceConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftDestinationConfiguration", GoGetter: "RedshiftDestinationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "s3DestinationConfiguration", GoGetter: "S3DestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "splunkDestinationConfiguration", GoGetter: "SplunkDestinationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeliveryStream{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.AmazonOpenSearchServerlessBufferingHintsProperty",
		reflect.TypeOf((*CfnDeliveryStream_AmazonOpenSearchServerlessBufferingHintsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.AmazonOpenSearchServerlessDestinationConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_AmazonOpenSearchServerlessDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.AmazonOpenSearchServerlessRetryOptionsProperty",
		reflect.TypeOf((*CfnDeliveryStream_AmazonOpenSearchServerlessRetryOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.AmazonopensearchserviceBufferingHintsProperty",
		reflect.TypeOf((*CfnDeliveryStream_AmazonopensearchserviceBufferingHintsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.AmazonopensearchserviceDestinationConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_AmazonopensearchserviceDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.AmazonopensearchserviceRetryOptionsProperty",
		reflect.TypeOf((*CfnDeliveryStream_AmazonopensearchserviceRetryOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.BufferingHintsProperty",
		reflect.TypeOf((*CfnDeliveryStream_BufferingHintsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.CloudWatchLoggingOptionsProperty",
		reflect.TypeOf((*CfnDeliveryStream_CloudWatchLoggingOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.CopyCommandProperty",
		reflect.TypeOf((*CfnDeliveryStream_CopyCommandProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.DataFormatConversionConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_DataFormatConversionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.DeliveryStreamEncryptionConfigurationInputProperty",
		reflect.TypeOf((*CfnDeliveryStream_DeliveryStreamEncryptionConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.DeserializerProperty",
		reflect.TypeOf((*CfnDeliveryStream_DeserializerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.DocumentIdOptionsProperty",
		reflect.TypeOf((*CfnDeliveryStream_DocumentIdOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.DynamicPartitioningConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_DynamicPartitioningConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.ElasticsearchBufferingHintsProperty",
		reflect.TypeOf((*CfnDeliveryStream_ElasticsearchBufferingHintsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.ElasticsearchDestinationConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_ElasticsearchDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.ElasticsearchRetryOptionsProperty",
		reflect.TypeOf((*CfnDeliveryStream_ElasticsearchRetryOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.ExtendedS3DestinationConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_ExtendedS3DestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.HiveJsonSerDeProperty",
		reflect.TypeOf((*CfnDeliveryStream_HiveJsonSerDeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.HttpEndpointCommonAttributeProperty",
		reflect.TypeOf((*CfnDeliveryStream_HttpEndpointCommonAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.HttpEndpointConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_HttpEndpointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.HttpEndpointDestinationConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_HttpEndpointDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.HttpEndpointRequestConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_HttpEndpointRequestConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.InputFormatConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_InputFormatConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.KMSEncryptionConfigProperty",
		reflect.TypeOf((*CfnDeliveryStream_KMSEncryptionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.KinesisStreamSourceConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_KinesisStreamSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.OpenXJsonSerDeProperty",
		reflect.TypeOf((*CfnDeliveryStream_OpenXJsonSerDeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.OrcSerDeProperty",
		reflect.TypeOf((*CfnDeliveryStream_OrcSerDeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.OutputFormatConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_OutputFormatConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.ParquetSerDeProperty",
		reflect.TypeOf((*CfnDeliveryStream_ParquetSerDeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.ProcessingConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_ProcessingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.ProcessorParameterProperty",
		reflect.TypeOf((*CfnDeliveryStream_ProcessorParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.ProcessorProperty",
		reflect.TypeOf((*CfnDeliveryStream_ProcessorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.RedshiftDestinationConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_RedshiftDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.RedshiftRetryOptionsProperty",
		reflect.TypeOf((*CfnDeliveryStream_RedshiftRetryOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.RetryOptionsProperty",
		reflect.TypeOf((*CfnDeliveryStream_RetryOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.S3DestinationConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_S3DestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.SchemaConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_SchemaConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.SerializerProperty",
		reflect.TypeOf((*CfnDeliveryStream_SerializerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.SplunkDestinationConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_SplunkDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.SplunkRetryOptionsProperty",
		reflect.TypeOf((*CfnDeliveryStream_SplunkRetryOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream.VpcConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_VpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStreamProps",
		reflect.TypeOf((*CfnDeliveryStreamProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.DataProcessorBindOptions",
		reflect.TypeOf((*DataProcessorBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.DataProcessorConfig",
		reflect.TypeOf((*DataProcessorConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.DataProcessorIdentifier",
		reflect.TypeOf((*DataProcessorIdentifier)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.DataProcessorProps",
		reflect.TypeOf((*DataProcessorProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_kinesisfirehose.DeliveryStream",
		reflect.TypeOf((*DeliveryStream)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamArn", GoGetter: "DeliveryStreamArn"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamName", GoGetter: "DeliveryStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantPutRecords", GoMethod: "GrantPutRecords"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3Bytes", GoMethod: "MetricBackupToS3Bytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3DataFreshness", GoMethod: "MetricBackupToS3DataFreshness"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3Records", GoMethod: "MetricBackupToS3Records"},
			_jsii_.MemberMethod{JsiiMethod: "metricIncomingBytes", GoMethod: "MetricIncomingBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricIncomingRecords", GoMethod: "MetricIncomingRecords"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_DeliveryStream{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDeliveryStream)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.DeliveryStreamAttributes",
		reflect.TypeOf((*DeliveryStreamAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.DeliveryStreamProps",
		reflect.TypeOf((*DeliveryStreamProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.DestinationBindOptions",
		reflect.TypeOf((*DestinationBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_kinesisfirehose.DestinationConfig",
		reflect.TypeOf((*DestinationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"monocdk.aws_kinesisfirehose.IDataProcessor",
		reflect.TypeOf((*IDataProcessor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			return &jsiiProxy_IDataProcessor{}
		},
	)
	_jsii_.RegisterInterface(
		"monocdk.aws_kinesisfirehose.IDeliveryStream",
		reflect.TypeOf((*IDeliveryStream)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamArn", GoGetter: "DeliveryStreamArn"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamName", GoGetter: "DeliveryStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantPutRecords", GoMethod: "GrantPutRecords"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3Bytes", GoMethod: "MetricBackupToS3Bytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3DataFreshness", GoMethod: "MetricBackupToS3DataFreshness"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3Records", GoMethod: "MetricBackupToS3Records"},
			_jsii_.MemberMethod{JsiiMethod: "metricIncomingBytes", GoMethod: "MetricIncomingBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricIncomingRecords", GoMethod: "MetricIncomingRecords"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IDeliveryStream{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"monocdk.aws_kinesisfirehose.IDestination",
		reflect.TypeOf((*IDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IDestination{}
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_kinesisfirehose.LambdaFunctionProcessor",
		reflect.TypeOf((*LambdaFunctionProcessor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaFunctionProcessor{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDataProcessor)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_kinesisfirehose.StreamEncryption",
		reflect.TypeOf((*StreamEncryption)(nil)).Elem(),
		map[string]interface{}{
			"UNENCRYPTED": StreamEncryption_UNENCRYPTED,
			"CUSTOMER_MANAGED": StreamEncryption_CUSTOMER_MANAGED,
			"AWS_OWNED": StreamEncryption_AWS_OWNED,
		},
	)
}
