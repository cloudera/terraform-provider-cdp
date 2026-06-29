// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datahub

import (
	"context"
	"reflect"
	"strings"
	"testing"
	"unicode"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	datahubmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
)

func TestAwsDatahubSchemaRequiredFieldsMatchCDPRequest(t *testing.T) {
	// The schema marks "name" and "environment" as Required for UX, but the CDP struct
	// has no pointer-required fields (all use omitempty). These are intentional strictness.
	schemaStricterExceptions := map[string]bool{
		"name":        true, // Required in TF; CDP uses clusterName with omitempty
		"environment": true, // Required in TF; CDP uses environment with omitempty
		"image.id":    true, // Required within the image block in TF; CDP's ImageRequest.ID uses omitempty
	}

	assertRequiredFieldsMatch(t,
		reflect.TypeFor[datahubmodels.CreateAWSClusterRequest](),
		createFilledAwsDatahubTestObject(),
		"",
		schemaStricterExceptions,
	)
}

func TestAzureDatahubSchemaRequiredFieldsMatchCDPRequest(t *testing.T) {
	schemaStricterExceptions := map[string]bool{
		"name":        true, // Required in TF; CDP uses clusterName with omitempty
		"environment": true, // Required in TF; CDP uses environmentName with omitempty
		"image.id":    true, // Required within the image block in TF; CDP's ImageRequest.ID uses omitempty
	}

	assertRequiredFieldsMatch(t,
		reflect.TypeFor[datahubmodels.CreateAzureClusterRequest](),
		createFilledAzureDatahubTestObject(),
		"",
		schemaStricterExceptions,
	)
}

func TestGcpDatahubSchemaRequiredFieldsMatchCDPRequest(t *testing.T) {
	schemaStricterExceptions := map[string]bool{
		"environment_name": true, // Required in TF; CDP uses environmentName with omitempty
		"image.id":         true, // Required within the image block in TF; CDP's ImageRequest.ID uses omitempty
	}

	assertRequiredFieldsMatch(t,
		reflect.TypeFor[datahubmodels.CreateGCPClusterRequest](),
		createFilledGcpDatahubTestObject(),
		"",
		schemaStricterExceptions,
	)
}

// Tests that every required CDP struct field has a corresponding schema field.
// Note: datahub CDP structs currently have no required fields (all use omitempty),
// but these tests will catch any future SDK regeneration that introduces required fields.

func TestAwsDatahubRequiredCDPFieldsHaveSchemaCounterpart(t *testing.T) {
	notYetExposed := map[string]bool{}

	assertRequiredCDPFieldsHaveSchema(t,
		reflect.TypeFor[datahubmodels.CreateAWSClusterRequest](),
		createFilledAwsDatahubTestObject(),
		"",
		notYetExposed,
	)
}

func TestAzureDatahubRequiredCDPFieldsHaveSchemaCounterpart(t *testing.T) {
	notYetExposed := map[string]bool{}

	assertRequiredCDPFieldsHaveSchema(t,
		reflect.TypeFor[datahubmodels.CreateAzureClusterRequest](),
		createFilledAzureDatahubTestObject(),
		"",
		notYetExposed,
	)
}

func TestGcpDatahubRequiredCDPFieldsHaveSchemaCounterpart(t *testing.T) {
	notYetExposed := map[string]bool{}

	assertRequiredCDPFieldsHaveSchema(t,
		reflect.TypeFor[datahubmodels.CreateGCPClusterRequest](),
		createFilledGcpDatahubTestObject(),
		"",
		notYetExposed,
	)
}

// Tests that every user-facing schema field is accounted for.

func TestAwsDatahubSchemaFieldsCoverage(t *testing.T) {
	tfOnlyFields := map[string]bool{
		"id":              true,
		"crn":             true,
		"status":          true,
		"polling_options": true,
		"destroy_options": true,
		"name":            true, // Maps to CDP's clusterName (different naming)
		"environment":     true, // Maps to CDP's environment (same name but Required handling differs)
		"instance_group":  true, // Maps to CDP's instanceGroups (singular vs plural, different structure)
	}

	assertAllSchemaFieldsMapped(t,
		reflect.TypeFor[datahubmodels.CreateAWSClusterRequest](),
		createFilledAwsDatahubTestObject(),
		tfOnlyFields,
	)
}

func TestAzureDatahubSchemaFieldsCoverage(t *testing.T) {
	tfOnlyFields := map[string]bool{
		"id":                 true,
		"crn":                true,
		"status":             true,
		"polling_options":    true,
		"destroy_options":    true,
		"name":               true, // Maps to CDP's clusterName (different naming)
		"environment":        true, // Maps to CDP's environmentName (different naming)
		"instance_group":     true, // Maps to CDP's instanceGroups (singular vs plural, different structure)
		"cluster_template":   true, // Maps to CDP's clusterTemplateName (alias without _name suffix)
		"cluster_definition": true, // Maps to CDP's clusterDefinitionName (alias without _name suffix)
	}

	assertAllSchemaFieldsMapped(t,
		reflect.TypeFor[datahubmodels.CreateAzureClusterRequest](),
		createFilledAzureDatahubTestObject(),
		tfOnlyFields,
	)
}

func TestGcpDatahubSchemaFieldsCoverage(t *testing.T) {
	tfOnlyFields := map[string]bool{
		"id":              true,
		"crn":             true,
		"status":          true,
		"polling_options": true,
		"destroy_options": true,
		"instance_group":  true, // Maps to CDP's instanceGroups (singular vs plural, different structure)
		"name":            true, // Maps to CDP's clusterName (different naming)
	}

	assertAllSchemaFieldsMapped(t,
		reflect.TypeFor[datahubmodels.CreateGCPClusterRequest](),
		createFilledGcpDatahubTestObject(),
		tfOnlyFields,
	)
}

// --- Schema accessors ---

func createFilledAwsDatahubTestObject() map[string]schema.Attribute {
	res := &awsDatahubResource{}
	schemaResponse := &resource.SchemaResponse{}
	res.Schema(context.TODO(), resource.SchemaRequest{}, schemaResponse)
	return schemaResponse.Schema.Attributes
}

func createFilledAzureDatahubTestObject() map[string]schema.Attribute {
	res := &azureDatahubResource{}
	schemaResponse := &resource.SchemaResponse{}
	res.Schema(context.TODO(), resource.SchemaRequest{}, schemaResponse)
	return schemaResponse.Schema.Attributes
}

func createFilledGcpDatahubTestObject() map[string]schema.Attribute {
	res := &gcpDatahubResource{}
	schemaResponse := &resource.SchemaResponse{}
	res.Schema(context.TODO(), resource.SchemaRequest{}, schemaResponse)
	return schemaResponse.Schema.Attributes
}

// --- Helper functions ---

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
