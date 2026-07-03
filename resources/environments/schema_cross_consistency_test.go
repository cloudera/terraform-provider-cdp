// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package environments

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getEnvironmentConfigProviderSchemas(t *testing.T) (aws, azure, gcp dsschema.SingleNestedAttribute) {
	t.Helper()
	ds := &environmentConfigDataSource{}
	resp := &datasource.SchemaResponse{}
	ds.Schema(context.TODO(), datasource.SchemaRequest{}, resp)

	awsAttr, ok := resp.Schema.Attributes["aws"].(dsschema.SingleNestedAttribute)
	require.True(t, ok, "aws attribute must be a SingleNestedAttribute")
	azureAttr, ok := resp.Schema.Attributes["azure"].(dsschema.SingleNestedAttribute)
	require.True(t, ok, "azure attribute must be a SingleNestedAttribute")
	gcpAttr, ok := resp.Schema.Attributes["gcp"].(dsschema.SingleNestedAttribute)
	require.True(t, ok, "gcp attribute must be a SingleNestedAttribute")

	return awsAttr, azureAttr, gcpAttr
}

func extractNestedFieldNames(t *testing.T, parent dsschema.SingleNestedAttribute, fieldName string) []string {
	t.Helper()
	attr, ok := parent.Attributes[fieldName]
	require.True(t, ok, "field %q must exist in parent schema", fieldName)
	nested, ok := attr.(dsschema.SingleNestedAttribute)
	require.True(t, ok, "field %q must be a SingleNestedAttribute", fieldName)

	names := make([]string, 0, len(nested.Attributes))
	for name := range nested.Attributes {
		names = append(names, name)
	}
	return names
}

// TestFreeIpaResourceFieldsExistInDatasourceConfig verifies that every field in the
// resource FreeIpaSchema also exists in the datasource freeIpaConfigSchema. This catches
// drift when someone adds a field to the resource schema but forgets the datasource config.
func TestFreeIpaResourceFieldsExistInDatasourceConfig(t *testing.T) {
	resourceFields := FreeIpaSchema.Attributes
	datasourceFields := freeIpaConfigSchema.Attributes

	for fieldName := range resourceFields {
		assert.Contains(t, datasourceFields, fieldName,
			"FreeIPA resource field %q is missing from the datasource config schema (freeIpaConfigSchema)", fieldName)
	}
}

// TestFreeIpaDatasourceConfigFieldsExistInResource verifies that every field in the
// datasource freeIpaConfigSchema also exists in the resource FreeIpaSchema.
func TestFreeIpaDatasourceConfigFieldsExistInResource(t *testing.T) {
	resourceFields := FreeIpaSchema.Attributes
	datasourceFields := freeIpaConfigSchema.Attributes

	for fieldName := range datasourceFields {
		assert.Contains(t, resourceFields, fieldName,
			"FreeIPA datasource config field %q is missing from the resource schema (FreeIpaSchema)", fieldName)
	}
}

// TestSecurityAccessResourceFieldsExistInAwsConfig verifies that every field in the
// resource securityAccess schema also exists in the AWS datasource config security_access.
func TestSecurityAccessResourceFieldsExistInAwsConfig(t *testing.T) {
	aws, _, _ := getEnvironmentConfigProviderSchemas(t)
	dsFields := extractNestedFieldNames(t, aws, "security_access")

	for fieldName := range securityAccess.Attributes {
		assert.Contains(t, dsFields, fieldName,
			"securityAccess resource field %q is missing from AWS datasource config security_access", fieldName)
	}
}

// TestSecurityAccessResourceFieldsExistInAzureConfig verifies that every field in the
// resource securityAccess schema also exists in the Azure datasource config security_access.
func TestSecurityAccessResourceFieldsExistInAzureConfig(t *testing.T) {
	_, azure, _ := getEnvironmentConfigProviderSchemas(t)
	dsFields := extractNestedFieldNames(t, azure, "security_access")

	for fieldName := range securityAccess.Attributes {
		assert.Contains(t, dsFields, fieldName,
			"securityAccess resource field %q is missing from Azure datasource config security_access", fieldName)
	}
}

// TestPollingOptionsConsistentAcrossProviders verifies that the polling_options nested
// schema has the same number of fields across aws, azure, and gcp in the environment config.
func TestPollingOptionsConsistentAcrossProviders(t *testing.T) {
	aws, azure, gcp := getEnvironmentConfigProviderSchemas(t)

	awsFields := extractNestedFieldNames(t, aws, "polling_options")
	azureFields := extractNestedFieldNames(t, azure, "polling_options")
	gcpFields := extractNestedFieldNames(t, gcp, "polling_options")

	assert.Equal(t, len(awsFields), len(azureFields),
		"AWS and Azure polling_options should have the same number of fields")
	assert.Equal(t, len(awsFields), len(gcpFields),
		"AWS and GCP polling_options should have the same number of fields")

	for _, name := range awsFields {
		assert.Contains(t, azureFields, name,
			"polling_options field %q exists in AWS but not in Azure", name)
		assert.Contains(t, gcpFields, name,
			"polling_options field %q exists in AWS but not in GCP", name)
	}
}

// TestDeleteOptionsConsistentAcrossProviders verifies that the delete_options nested
// schema has the same number of fields across aws, azure, and gcp in the environment config.
func TestDeleteOptionsConsistentAcrossProviders(t *testing.T) {
	aws, azure, gcp := getEnvironmentConfigProviderSchemas(t)

	awsFields := extractNestedFieldNames(t, aws, "delete_options")
	azureFields := extractNestedFieldNames(t, azure, "delete_options")
	gcpFields := extractNestedFieldNames(t, gcp, "delete_options")

	assert.Equal(t, len(awsFields), len(azureFields),
		"AWS and Azure delete_options should have the same number of fields")
	assert.Equal(t, len(awsFields), len(gcpFields),
		"AWS and GCP delete_options should have the same number of fields")

	for _, name := range awsFields {
		assert.Contains(t, azureFields, name,
			"delete_options field %q exists in AWS but not in Azure", name)
		assert.Contains(t, gcpFields, name,
			"delete_options field %q exists in AWS but not in GCP", name)
	}
}

// TestFreeIpaConsistentAcrossProviders verifies that the freeipa nested schema has the
// same fields across aws, azure, and gcp in the environment config datasource.
func TestFreeIpaConsistentAcrossProviders(t *testing.T) {
	aws, azure, gcp := getEnvironmentConfigProviderSchemas(t)

	awsFields := extractNestedFieldNames(t, aws, "freeipa")
	azureFields := extractNestedFieldNames(t, azure, "freeipa")
	gcpFields := extractNestedFieldNames(t, gcp, "freeipa")

	assert.Equal(t, len(awsFields), len(azureFields),
		"AWS and Azure freeipa should have the same number of fields")
	assert.Equal(t, len(awsFields), len(gcpFields),
		"AWS and GCP freeipa should have the same number of fields")

	for _, name := range awsFields {
		assert.Contains(t, azureFields, name,
			"freeipa field %q exists in AWS but not in Azure", name)
		assert.Contains(t, gcpFields, name,
			"freeipa field %q exists in AWS but not in GCP", name)
	}
}
