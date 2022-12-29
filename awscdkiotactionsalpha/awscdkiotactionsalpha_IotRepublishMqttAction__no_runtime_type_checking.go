//go:build no_runtime_type_checking

// Receipt rule actions for AWS IoT
package awscdkiotactionsalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateNewIotRepublishMqttActionParameters(topic *string, props *IotRepublishMqttActionProps) error {
	return nil
}
