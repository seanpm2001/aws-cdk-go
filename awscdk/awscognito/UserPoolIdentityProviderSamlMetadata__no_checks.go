//go:build no_runtime_type_checking

package awscognito

// Building without runtime type checking enabled, so all the below just return nil

func validateUserPoolIdentityProviderSamlMetadata_FileParameters(fileContent *string) error {
	return nil
}

func validateUserPoolIdentityProviderSamlMetadata_UrlParameters(url *string) error {
	return nil
}

