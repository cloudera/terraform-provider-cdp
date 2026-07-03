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
	"reflect"
	"strings"
	"testing"
	"unicode"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

func TestAwsSchemaRequiredFieldsMatchCDPRequest(t *testing.T) {
	// Fields that are intentionally required in the schema but optional in the CDP API.
	schemaStricterExceptions := map[string]bool{
		"vpc_id":     true, // Required in TF for usability; CDP API technically allows omission
		"subnet_ids": true, // Required in TF for usability; CDP API marks as slice without pointer
	}

	assertRequiredFieldsMatch(t,
		reflect.TypeFor[environmentsmodels.CreateAWSEnvironmentRequest](),
		AwsEnvironmentSchema.Attributes,
		"",
		schemaStricterExceptions,
	)
}

func TestAzureSchemaRequiredFieldsMatchCDPRequest(t *testing.T) {
	// Slice fields that are required per the API spec (validate.Required in Validate())
	// but cannot be detected as required via reflection alone ([]T has no pointer indirection).
	schemaStricterExceptions := map[string]bool{
		"existing_network_params.subnet_ids": true,
	}

	assertRequiredFieldsMatch(t,
		reflect.TypeFor[environmentsmodels.CreateAzureEnvironmentRequest](),
		AzureEnvironmentSchema.Attributes,
		"",
		schemaStricterExceptions,
	)
}

func TestGcpSchemaRequiredFieldsMatchCDPRequest(t *testing.T) {
	schemaStricterExceptions := map[string]bool{
		// shared_project_id is Required in TF schema for usability but optional in the CDP API.
		"existing_network_params.shared_project_id": true,
		// Slice field that is required per the API spec (validate.Required in Validate())
		// but cannot be detected as required via reflection alone.
		"existing_network_params.subnet_names": true,
	}

	assertRequiredFieldsMatch(t,
		reflect.TypeFor[environmentsmodels.CreateGCPEnvironmentRequest](),
		GcpEnvironmentSchema.Attributes,
		"",
		schemaStricterExceptions,
	)
}

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
			// CDP field has no corresponding schema field — this is fine for fields the
			// provider doesn't expose yet, or that map differently.
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

		// Recurse into nested structs if the schema attribute has nested attributes.
		nestedAttrs := getNestedAttributes(attr)
		if nestedAttrs != nil {
			nestedType := resolveStructType(field.Type)
			if nestedType != nil {
				assertRequiredFieldsMatch(t, nestedType, nestedAttrs, fullPath, exceptions)
			}
		}
	}
}

// isCDPFieldRequired determines if a CDP struct field is required based on the
// go-swagger code generation pattern: pointer type without omitempty = required.
func isCDPFieldRequired(field reflect.StructField, jsonOpts string) bool {
	hasOmitEmpty := strings.Contains(jsonOpts, "omitempty")
	if hasOmitEmpty {
		return false
	}
	kind := field.Type.Kind()
	return kind == reflect.Pointer
}

// isComputedOnly returns true if the attribute is purely computed (not user-settable).
func isComputedOnly(attr schema.Attribute) bool {
	return attr.IsComputed() && !attr.IsRequired() && !attr.IsOptional()
}

// getNestedAttributes extracts the nested attribute map from a SingleNestedAttribute.
func getNestedAttributes(attr schema.Attribute) map[string]schema.Attribute {
	if nested, ok := attr.(schema.SingleNestedAttribute); ok {
		return nested.Attributes
	}
	return nil
}

// resolveStructType dereferences pointers and returns the underlying struct type, or nil.
func resolveStructType(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t.Kind() == reflect.Struct {
		return t
	}
	return nil
}

// parseJSONTag splits a json struct tag into the field name and remaining options.
func parseJSONTag(tag string) (name string, opts string) {
	parts := strings.SplitN(tag, ",", 2)
	name = parts[0]
	if len(parts) > 1 {
		opts = parts[1]
	}
	return
}

// Tests below ensure that every required CDP struct field has a corresponding schema field.
// If the SDK is regenerated and a new required field appears, the test will fail until
// the field is either added to the schema or explicitly listed as not yet exposed.

func TestAwsRequiredCDPFieldsHaveSchemaCounterpart(t *testing.T) {
	notYetExposed := map[string]bool{}

	assertRequiredCDPFieldsHaveSchema(t,
		reflect.TypeFor[environmentsmodels.CreateAWSEnvironmentRequest](),
		AwsEnvironmentSchema.Attributes,
		"",
		notYetExposed,
	)
}

func TestAzureRequiredCDPFieldsHaveSchemaCounterpart(t *testing.T) {
	notYetExposed := map[string]bool{}

	assertRequiredCDPFieldsHaveSchema(t,
		reflect.TypeFor[environmentsmodels.CreateAzureEnvironmentRequest](),
		AzureEnvironmentSchema.Attributes,
		"",
		notYetExposed,
	)
}

func TestGcpRequiredCDPFieldsHaveSchemaCounterpart(t *testing.T) {
	notYetExposed := map[string]bool{}

	assertRequiredCDPFieldsHaveSchema(t,
		reflect.TypeFor[environmentsmodels.CreateGCPEnvironmentRequest](),
		GcpEnvironmentSchema.Attributes,
		"",
		notYetExposed,
	)
}

// assertRequiredCDPFieldsHaveSchema checks that every required field on the CDP struct
// has a corresponding field in the Terraform schema. This catches new required fields
// introduced by SDK regeneration that haven't been added to the schema yet.
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

		// Recurse into nested required structs.
		nestedAttrs := getNestedAttributes(attr)
		if nestedAttrs != nil {
			nestedType := resolveStructType(field.Type)
			if nestedType != nil {
				assertRequiredCDPFieldsHaveSchema(t, nestedType, nestedAttrs, fullPath, notYetExposed)
			}
		}
	}
}

// Tests below ensure that every user-facing schema field is accounted for in the
// alignment check. If someone adds a new field to the schema, they must either:
// - Ensure it has a matching CDP struct field (auto-checked by the alignment test), OR
// - Add it to the tfOnlyFields set (explicitly acknowledging it has no CDP counterpart).
// This prevents new fields from silently bypassing the required-field alignment check.

func TestAwsSchemaFieldsCoverage(t *testing.T) {
	tfOnlyFields := map[string]bool{
		"id":              true,
		"crn":             true,
		"status":          true,
		"status_reason":   true,
		"tunnel_type":     true,
		"polling_options": true,
		"delete_options":  true,
		"freeipa":         true,
		"compute_cluster": true, // Maps to computeClusterConfiguration + enableComputeCluster (composite)
	}

	assertAllSchemaFieldsMapped(t,
		reflect.TypeFor[environmentsmodels.CreateAWSEnvironmentRequest](),
		AwsEnvironmentSchema.Attributes,
		tfOnlyFields,
	)
}

func TestAzureSchemaFieldsCoverage(t *testing.T) {
	tfOnlyFields := map[string]bool{
		"id":              true,
		"crn":             true,
		"status":          true,
		"status_reason":   true,
		"tunnel_type":     true,
		"polling_options": true,
		"delete_options":  true,
		"freeipa":         true,
		"compute_cluster": true, // Maps to computeClusterConfiguration + enableComputeCluster (composite)
	}

	assertAllSchemaFieldsMapped(t,
		reflect.TypeFor[environmentsmodels.CreateAzureEnvironmentRequest](),
		AzureEnvironmentSchema.Attributes,
		tfOnlyFields,
	)
}

func TestGcpSchemaFieldsCoverage(t *testing.T) {
	tfOnlyFields := map[string]bool{
		"id":              true,
		"crn":             true,
		"status":          true,
		"status_reason":   true,
		"tunnel_type":     true,
		"polling_options": true,
		"delete_options":  true,
		"freeipa":         true,
	}

	assertAllSchemaFieldsMapped(t,
		reflect.TypeFor[environmentsmodels.CreateGCPEnvironmentRequest](),
		GcpEnvironmentSchema.Attributes,
		tfOnlyFields,
	)
}

// assertAllSchemaFieldsMapped checks that every non-computed-only schema field either
// maps to a field on the CDP struct (by convention-based name conversion) or is
// explicitly listed in tfOnlyFields.
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

// collectCDPSchemaNames returns the set of snake_case names derived from a CDP struct's json tags.
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

func TestCamelToSnake(t *testing.T) {
	cases := []struct{ in, want string }{
		{"credentialName", "credential_name"},
		{"vpcId", "vpc_id"},
		{"s3GuardTableName", "s3_guard_table_name"},
		{"environmentName", "environment_name"},
		{"encryptionKeyArn", "encryption_key_arn"},
		{"endpointAccessGatewayScheme", "endpoint_access_gateway_scheme"},
		{"usePublicIp", "use_public_ip"},
		{"networkId", "network_id"},
		{"aksPrivateDnsZoneId", "aks_private_dns_zone_id"},
		{"subnetIds", "subnet_ids"},
		{"encryptionKeyUrl", "encryption_key_url"},
		{"logStorage", "log_storage"},
		{"securityAccess", "security_access"},
	}
	for _, tc := range cases {
		got := camelToSnake(tc.in)
		if got != tc.want {
			t.Errorf("camelToSnake(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

// --- AWS Credential alignment tests ---

func TestAwsCredentialSchemaRequiredFieldsMatchCDPRequest(t *testing.T) {
	schemaStricterExceptions := map[string]bool{}

	assertRequiredFieldsMatch(t,
		reflect.TypeFor[environmentsmodels.CreateAWSCredentialRequest](),
		createFilledAwsCredentialTestObject(),
		"",
		schemaStricterExceptions,
	)
}

func TestAwsCredentialRequiredCDPFieldsHaveSchemaCounterpart(t *testing.T) {
	notYetExposed := map[string]bool{}

	assertRequiredCDPFieldsHaveSchema(t,
		reflect.TypeFor[environmentsmodels.CreateAWSCredentialRequest](),
		createFilledAwsCredentialTestObject(),
		"",
		notYetExposed,
	)
}

func TestAwsCredentialSchemaFieldsCoverage(t *testing.T) {
	tfOnlyFields := map[string]bool{
		"id":                        true,
		"crn":                       true,
		"skip_org_policy_decisions": true, // TF-only computed field
		"verify_permissions":        true, // TF-only computed field
	}

	assertAllSchemaFieldsMapped(t,
		reflect.TypeFor[environmentsmodels.CreateAWSCredentialRequest](),
		createFilledAwsCredentialTestObject(),
		tfOnlyFields,
	)
}

// --- Azure Credential alignment tests ---

func TestAzureCredentialSchemaRequiredFieldsMatchCDPRequest(t *testing.T) {
	schemaStricterExceptions := map[string]bool{
		"subscription_id":          true, // Required in TF for usability; CDP marks as optional (omitempty)
		"tenant_id":                true, // Required in TF for usability; CDP marks as optional (omitempty)
		"app_based.application_id": true, // Required in TF for usability; CDP marks as optional (omitempty)
		"app_based.secret_key":     true, // Required in TF for usability; CDP marks as optional (omitempty)
	}

	assertRequiredFieldsMatch(t,
		reflect.TypeFor[environmentsmodels.CreateAzureCredentialRequest](),
		createFilledAzureCredentialTestObject(),
		"",
		schemaStricterExceptions,
	)
}

func TestAzureCredentialRequiredCDPFieldsHaveSchemaCounterpart(t *testing.T) {
	notYetExposed := map[string]bool{}

	assertRequiredCDPFieldsHaveSchema(t,
		reflect.TypeFor[environmentsmodels.CreateAzureCredentialRequest](),
		createFilledAzureCredentialTestObject(),
		"",
		notYetExposed,
	)
}

func TestAzureCredentialSchemaFieldsCoverage(t *testing.T) {
	tfOnlyFields := map[string]bool{
		"id":  true,
		"crn": true,
	}

	assertAllSchemaFieldsMapped(t,
		reflect.TypeFor[environmentsmodels.CreateAzureCredentialRequest](),
		createFilledAzureCredentialTestObject(),
		tfOnlyFields,
	)
}

// --- GCP Credential alignment tests ---

func TestGcpCredentialSchemaRequiredFieldsMatchCDPRequest(t *testing.T) {
	schemaStricterExceptions := map[string]bool{}

	assertRequiredFieldsMatch(t,
		reflect.TypeFor[environmentsmodels.CreateGCPCredentialRequest](),
		createFilledGcpCredentialTestObject(),
		"",
		schemaStricterExceptions,
	)
}

func TestGcpCredentialRequiredCDPFieldsHaveSchemaCounterpart(t *testing.T) {
	notYetExposed := map[string]bool{}

	assertRequiredCDPFieldsHaveSchema(t,
		reflect.TypeFor[environmentsmodels.CreateGCPCredentialRequest](),
		createFilledGcpCredentialTestObject(),
		"",
		notYetExposed,
	)
}

func TestGcpCredentialSchemaFieldsCoverage(t *testing.T) {
	tfOnlyFields := map[string]bool{
		"id":  true,
		"crn": true,
	}

	assertAllSchemaFieldsMapped(t,
		reflect.TypeFor[environmentsmodels.CreateGCPCredentialRequest](),
		createFilledGcpCredentialTestObject(),
		tfOnlyFields,
	)
}

// --- ID Broker Mappings alignment tests ---

func TestIDBrokerMappingsSchemaRequiredFieldsMatchCDPRequest(t *testing.T) {
	schemaStricterExceptions := map[string]bool{
		"environment_name":  true, // Required in TF; CDP also marks as required but field name matches
		"ranger_audit_role": true, // Required in TF for usability; CDP uses string with omitempty (not detectable as required via reflection)
		"data_access_role":  true, // Required in CDP but Optional+Computed in TF; resource manages it via state
	}

	assertRequiredFieldsMatch(t,
		reflect.TypeFor[environmentsmodels.SetIDBrokerMappingsRequest](),
		IDBrokerMappingSchema.Attributes,
		"",
		schemaStricterExceptions,
	)
}

func TestIDBrokerMappingsRequiredCDPFieldsHaveSchemaCounterpart(t *testing.T) {
	notYetExposed := map[string]bool{}

	assertRequiredCDPFieldsHaveSchema(t,
		reflect.TypeFor[environmentsmodels.SetIDBrokerMappingsRequest](),
		IDBrokerMappingSchema.Attributes,
		"",
		notYetExposed,
	)
}

func TestIDBrokerMappingsSchemaFieldsCoverage(t *testing.T) {
	tfOnlyFields := map[string]bool{
		"id":               true,
		"environment_crn":  true, // TF-only; not in SetIDBrokerMappingsRequest
		"mappings_version": true, // TF-only computed field
	}

	assertAllSchemaFieldsMapped(t,
		reflect.TypeFor[environmentsmodels.SetIDBrokerMappingsRequest](),
		IDBrokerMappingSchema.Attributes,
		tfOnlyFields,
	)
}

// --- Proxy Configuration alignment tests ---

func TestProxyConfigSchemaRequiredFieldsMatchCDPRequest(t *testing.T) {
	schemaStricterExceptions := map[string]bool{}

	assertRequiredFieldsMatch(t,
		reflect.TypeFor[environmentsmodels.CreateProxyConfigRequest](),
		ProxyConfigurationSchema.Attributes,
		"",
		schemaStricterExceptions,
	)
}

func TestProxyConfigRequiredCDPFieldsHaveSchemaCounterpart(t *testing.T) {
	notYetExposed := map[string]bool{
		"proxy_config_name": true, // Schema uses "name" instead (different naming convention)
	}

	assertRequiredCDPFieldsHaveSchema(t,
		reflect.TypeFor[environmentsmodels.CreateProxyConfigRequest](),
		ProxyConfigurationSchema.Attributes,
		"",
		notYetExposed,
	)
}

func TestProxyConfigSchemaFieldsCoverage(t *testing.T) {
	tfOnlyFields := map[string]bool{
		"id":   true,
		"name": true, // Maps to CDP's proxyConfigName (different naming convention)
	}

	assertAllSchemaFieldsMapped(t,
		reflect.TypeFor[environmentsmodels.CreateProxyConfigRequest](),
		ProxyConfigurationSchema.Attributes,
		tfOnlyFields,
	)
}

// --- Azure Image Terms alignment tests ---

func TestAzureImageTermsSchemaRequiredFieldsMatchCDPRequest(t *testing.T) {
	schemaStricterExceptions := map[string]bool{}

	assertRequiredFieldsMatch(t,
		reflect.TypeFor[environmentsmodels.UpdateAzureImageTermsPolicyRequest](),
		AzureImageTermsPolicySchema.Attributes,
		"",
		schemaStricterExceptions,
	)
}

func TestAzureImageTermsRequiredCDPFieldsHaveSchemaCounterpart(t *testing.T) {
	notYetExposed := map[string]bool{}

	assertRequiredCDPFieldsHaveSchema(t,
		reflect.TypeFor[environmentsmodels.UpdateAzureImageTermsPolicyRequest](),
		AzureImageTermsPolicySchema.Attributes,
		"",
		notYetExposed,
	)
}

func TestAzureImageTermsSchemaFieldsCoverage(t *testing.T) {
	tfOnlyFields := map[string]bool{
		"id": true,
	}

	assertAllSchemaFieldsMapped(t,
		reflect.TypeFor[environmentsmodels.UpdateAzureImageTermsPolicyRequest](),
		AzureImageTermsPolicySchema.Attributes,
		tfOnlyFields,
	)
}

// camelToSnake converts a camelCase string to snake_case.
// Handles consecutive uppercase like "ID" -> "id", "URL" -> "url", "DNS" -> "dns".
func camelToSnake(s string) string {
	var result strings.Builder
	runes := []rune(s)

	for i, r := range runes {
		if unicode.IsUpper(r) {
			if i > 0 {
				// Insert underscore before uppercase unless:
				// - previous char is already uppercase AND next char is uppercase or end
				//   (we're in the middle of an acronym like "DNS" or "ID")
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
