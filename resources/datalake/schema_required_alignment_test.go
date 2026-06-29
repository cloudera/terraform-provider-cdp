// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datalake

import (
	"reflect"
	"strings"
	"testing"
	"unicode"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	datalakemodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

func TestAwsDatalakeSchemaRequiredFieldsMatchCDPRequest(t *testing.T) {
	// AWS datalake flattens CloudProviderConfiguration fields (instance_profile, storage_location_base)
	// to top-level schema attributes, so they won't match the nested CDP struct path.
	schemaStricterExceptions := map[string]bool{
		"instance_profile":      true, // Maps to CloudProviderConfiguration.InstanceProfile (composite mapping)
		"storage_location_base": true, // Maps to CloudProviderConfiguration.StorageBucketLocation (composite mapping)
		"environment_name":      true, // Deprecated; schema uses "environment" as the primary Required field instead
	}

	assertRequiredFieldsMatch(t,
		reflect.TypeFor[datalakemodels.CreateAWSDatalakeRequest](),
		createFilledAwsTestObject(),
		"",
		schemaStricterExceptions,
	)
}

func TestAzureDatalakeSchemaRequiredFieldsMatchCDPRequest(t *testing.T) {
	// Azure datalake flattens CloudProviderConfiguration fields to top-level schema attributes.
	schemaStricterExceptions := map[string]bool{
		"managed_identity":      true, // Maps to CloudProviderConfiguration.ManagedIdentity (composite mapping)
		"storage_location_base": true, // Maps to CloudProviderConfiguration.StorageLocation (composite mapping)
		"environment_name":      true, // Deprecated; schema uses "environment" as the primary Required field instead
	}

	assertRequiredFieldsMatch(t,
		reflect.TypeFor[datalakemodels.CreateAzureDatalakeRequest](),
		createFilledAzureTestObject(),
		"",
		schemaStricterExceptions,
	)
}

func TestGcpDatalakeSchemaRequiredFieldsMatchCDPRequest(t *testing.T) {
	schemaStricterExceptions := map[string]bool{
		"environment_name": true, // Deprecated; schema uses "environment" as the primary Required field instead
	}

	assertRequiredFieldsMatch(t,
		reflect.TypeFor[datalakemodels.CreateGCPDatalakeRequest](),
		createFilledTestObject(),
		"",
		schemaStricterExceptions,
	)
}

// Tests that every required CDP struct field has a corresponding schema field.

func TestAwsDatalakeRequiredCDPFieldsHaveSchemaCounterpart(t *testing.T) {
	// cloud_provider_configuration is required in CDP but AWS schema flattens its children
	// to top-level attributes (instance_profile, storage_location_base).
	notYetExposed := map[string]bool{
		"cloud_provider_configuration": true,
	}

	assertRequiredCDPFieldsHaveSchema(t,
		reflect.TypeFor[datalakemodels.CreateAWSDatalakeRequest](),
		createFilledAwsTestObject(),
		"",
		notYetExposed,
	)
}

func TestAzureDatalakeRequiredCDPFieldsHaveSchemaCounterpart(t *testing.T) {
	// cloud_provider_configuration is required in CDP but Azure schema flattens its children
	// to top-level attributes (managed_identity, storage_location_base).
	notYetExposed := map[string]bool{
		"cloud_provider_configuration": true,
	}

	assertRequiredCDPFieldsHaveSchema(t,
		reflect.TypeFor[datalakemodels.CreateAzureDatalakeRequest](),
		createFilledAzureTestObject(),
		"",
		notYetExposed,
	)
}

func TestGcpDatalakeRequiredCDPFieldsHaveSchemaCounterpart(t *testing.T) {
	notYetExposed := map[string]bool{}

	assertRequiredCDPFieldsHaveSchema(t,
		reflect.TypeFor[datalakemodels.CreateGCPDatalakeRequest](),
		createFilledTestObject(),
		"",
		notYetExposed,
	)
}

// Tests that every user-facing schema field is accounted for.

func TestAwsDatalakeSchemaFieldsCoverage(t *testing.T) {
	tfOnlyFields := map[string]bool{
		"id":                           true,
		"crn":                          true,
		"status":                       true,
		"status_reason":                true,
		"creation_date":                true,
		"environment_crn":              true,
		"polling_options":              true,
		"delete_options":               true,
		"certificate_expiration_state": true,
		"instance_profile":             true, // Flat field mapping to CloudProviderConfiguration.InstanceProfile
		"storage_location_base":        true, // Flat field mapping to CloudProviderConfiguration.StorageBucketLocation
		"environment_name":             true, // Deprecated alias for environment
		"environment":                  true, // Maps to CDP's environmentName
		"custom_instance_groups":       true, // Maps to CDP's customInstanceGroups but with different naming
		"recipes":                      true, // Maps to CDP's recipes but with different structure
	}

	assertAllSchemaFieldsMapped(t,
		reflect.TypeFor[datalakemodels.CreateAWSDatalakeRequest](),
		createFilledAwsTestObject(),
		tfOnlyFields,
	)
}

func TestAzureDatalakeSchemaFieldsCoverage(t *testing.T) {
	tfOnlyFields := map[string]bool{
		"id":                           true,
		"crn":                          true,
		"status":                       true,
		"status_reason":                true,
		"creation_date":                true,
		"environment_crn":              true,
		"polling_options":              true,
		"delete_options":               true,
		"certificate_expiration_state": true,
		"managed_identity":             true, // Flat field mapping to CloudProviderConfiguration.ManagedIdentity
		"storage_location_base":        true, // Flat field mapping to CloudProviderConfiguration.StorageLocation
		"environment_name":             true, // Deprecated alias for environment
		"environment":                  true, // Maps to CDP's environmentName
		"custom_instance_groups":       true, // Maps to CDP's customInstanceGroups but with different naming
		"recipes":                      true, // Maps to CDP's recipes but with different structure
	}

	assertAllSchemaFieldsMapped(t,
		reflect.TypeFor[datalakemodels.CreateAzureDatalakeRequest](),
		createFilledAzureTestObject(),
		tfOnlyFields,
	)
}

func TestGcpDatalakeSchemaFieldsCoverage(t *testing.T) {
	tfOnlyFields := map[string]bool{
		"id":                     true,
		"crn":                    true,
		"status":                 true,
		"status_reason":          true,
		"creation_date":          true,
		"environment_crn":        true,
		"polling_options":        true,
		"delete_options":         true,
		"environment_name":       true, // Deprecated alias for environment
		"environment":            true, // Maps to CDP's environmentName
		"custom_instance_groups": true, // Maps to CDP's customInstanceGroups but with different naming
		"recipes":                true, // Maps to CDP's recipes but with different structure
	}

	assertAllSchemaFieldsMapped(t,
		reflect.TypeFor[datalakemodels.CreateGCPDatalakeRequest](),
		createFilledTestObject(),
		tfOnlyFields,
	)
}

// --- Helper functions ---
// Note: createFilledAwsTestObject(), createFilledAzureTestObject(), and createFilledTestObject() (GCP)
// are defined in their respective schema_*_datalake_test.go files and reused here.

func assertRequiredFieldsMatch(t *testing.T, cdpType reflect.Type, schemaAttrs map[string]schema.Attribute, path string, exceptions map[string]bool) {
	t.Helper()

	for field := range cdpType.Fields() {
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		jsonName, opts := parseJSONTag(jsonTag)
		schemaName := camelToSnake(jsonName)
		fullPath := schemaName
		if path != "" {
			fullPath = path + "." + schemaName
		}

		attr, exists := schemaAttrs[schemaName]
		if !exists {
			continue
		}

		if isComputedOnly(attr) {
			continue
		}

		cdpRequired := isCDPFieldRequired(field, opts)
		schemaRequired := attr.IsRequired()

		if cdpRequired && !schemaRequired && !exceptions[fullPath] {
			t.Errorf("%s: field is required in CDP request struct but not required in Terraform schema. "+
				"Fix: mark the field as Required in the schema, or add %q to schemaStricterExceptions if this is intentional", fullPath, fullPath)
		}

		if schemaRequired && !cdpRequired && !exceptions[fullPath] {
			t.Errorf("%s: field is required in Terraform schema but not required in CDP request struct. "+
				"Fix: change the field to Optional in the schema, or add %q to schemaStricterExceptions if this is intentional", fullPath, fullPath)
		}

		nestedAttrs := getNestedAttributes(attr)
		if nestedAttrs != nil {
			nestedType := resolveStructType(field.Type)
			if nestedType != nil {
				assertRequiredFieldsMatch(t, nestedType, nestedAttrs, fullPath, exceptions)
			}
		}
	}
}

func assertRequiredCDPFieldsHaveSchema(t *testing.T, cdpType reflect.Type, schemaAttrs map[string]schema.Attribute, path string, notYetExposed map[string]bool) {
	t.Helper()

	for field := range cdpType.Fields() {
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		jsonName, opts := parseJSONTag(jsonTag)
		if !isCDPFieldRequired(field, opts) {
			continue
		}

		schemaName := camelToSnake(jsonName)
		fullPath := schemaName
		if path != "" {
			fullPath = path + "." + schemaName
		}

		if notYetExposed[fullPath] {
			continue
		}

		attr, exists := schemaAttrs[schemaName]
		if !exists {
			t.Errorf("%s: required field in CDP request struct has no corresponding Terraform schema field. "+
				"Fix: add %q as a Required attribute in the schema, or add %q to notYetExposed if intentionally omitted", fullPath, schemaName, fullPath)
			continue
		}

		nestedAttrs := getNestedAttributes(attr)
		if nestedAttrs != nil {
			nestedType := resolveStructType(field.Type)
			if nestedType != nil {
				assertRequiredCDPFieldsHaveSchema(t, nestedType, nestedAttrs, fullPath, notYetExposed)
			}
		}
	}
}

func assertAllSchemaFieldsMapped(t *testing.T, cdpType reflect.Type, schemaAttrs map[string]schema.Attribute, tfOnlyFields map[string]bool) {
	t.Helper()

	cdpFieldNames := collectCDPSchemaNames(cdpType)

	for schemaName, attr := range schemaAttrs {
		if isComputedOnly(attr) {
			continue
		}
		if tfOnlyFields[schemaName] {
			continue
		}
		if !cdpFieldNames[schemaName] {
			t.Errorf("schema field %q has no matching CDP struct field and is not listed in tfOnlyFields. "+
				"Fix: if this field maps to a CDP request field, ensure the naming convention matches (camelCase json tag -> snake_case schema name). "+
				"Otherwise, add %q to tfOnlyFields in this test to acknowledge it has no CDP counterpart", schemaName, schemaName)
		}
	}
}

func collectCDPSchemaNames(cdpType reflect.Type) map[string]bool {
	names := make(map[string]bool)
	for field := range cdpType.Fields() {
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}
		jsonName, _ := parseJSONTag(jsonTag)
		names[camelToSnake(jsonName)] = true
	}
	return names
}

func isCDPFieldRequired(field reflect.StructField, jsonOpts string) bool {
	hasOmitEmpty := strings.Contains(jsonOpts, "omitempty")
	if hasOmitEmpty {
		return false
	}
	return field.Type.Kind() == reflect.Pointer
}

func isComputedOnly(attr schema.Attribute) bool {
	return attr.IsComputed() && !attr.IsRequired() && !attr.IsOptional()
}

func getNestedAttributes(attr schema.Attribute) map[string]schema.Attribute {
	if nested, ok := attr.(schema.SingleNestedAttribute); ok {
		return nested.Attributes
	}
	return nil
}

func resolveStructType(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t.Kind() == reflect.Struct {
		return t
	}
	return nil
}

func parseJSONTag(tag string) (name string, opts string) {
	parts := strings.SplitN(tag, ",", 2)
	name = parts[0]
	if len(parts) > 1 {
		opts = parts[1]
	}
	return
}

func camelToSnake(s string) string {
	var result strings.Builder
	runes := []rune(s)

	for i, r := range runes {
		if unicode.IsUpper(r) {
			if i > 0 {
				prevUpper := unicode.IsUpper(runes[i-1])
				nextLower := i+1 < len(runes) && unicode.IsLower(runes[i+1])

				if !prevUpper || nextLower {
					result.WriteRune('_')
				}
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}
