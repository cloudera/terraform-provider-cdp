package environments

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	rsschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func TestAwsEnvironmentSchemaMatchesEnvironmentConfigSchema(t *testing.T) {
	assertSchemaPathsMatch(
		t,
		"AWS",
		AwsEnvironmentSchema.Attributes,
		awsSchema.Attributes,
	)
}

func TestAzureEnvironmentSchemaMatchesEnvironmentConfigSchema(t *testing.T) {
	assertSchemaPathsMatch(
		t,
		"Azure",
		AzureEnvironmentSchema.Attributes,
		azureSchema.Attributes,
	)
}

func TestGcpEnvironmentSchemaMatchesEnvironmentConfigSchema(t *testing.T) {
	assertSchemaPathsMatch(
		t,
		"GCP",
		GcpEnvironmentSchema.Attributes,
		gcpSchema.Attributes,
	)
}

func assertSchemaPathsMatch(
	t *testing.T,
	name string,
	resourceAttrs map[string]rsschema.Attribute,
	dataSourceAttrs map[string]dsschema.Attribute,
) {
	t.Helper()

	resourcePaths := collectResourceSchemaPaths(resourceAttrs, "")
	dataSourcePaths := collectDataSourceSchemaPaths(dataSourceAttrs, "")

	onlyInResource, onlyInDataSource := diffPathSets(resourcePaths, dataSourcePaths)

	if len(onlyInResource) == 0 && len(onlyInDataSource) == 0 {
		return
	}

	var b strings.Builder
	_, _ = fmt.Fprintf(&b, "%s schema mismatch\n", name)

	if len(onlyInResource) > 0 {
		b.WriteString("\nFields only in resource schema:\n")
		for _, p := range onlyInResource {
			_, _ = fmt.Fprintf(&b, "  - %s\n", p)
		}
	}

	if len(onlyInDataSource) > 0 {
		b.WriteString("\nFields only in datasource schema:\n")
		for _, p := range onlyInDataSource {
			_, _ = fmt.Fprintf(&b, "  - %s\n", p)
		}
	}

	t.Fatal(b.String())
}

func collectResourceSchemaPaths(attrs map[string]rsschema.Attribute, prefix string) map[string]struct{} {
	result := make(map[string]struct{})

	for name, attr := range attrs {
		path := joinPath(prefix, name)
		result[path] = struct{}{}

		switch a := attr.(type) {
		case rsschema.SingleNestedAttribute:
			for k, v := range collectResourceSchemaPaths(a.Attributes, path) {
				result[k] = v
			}
		case rsschema.ListNestedAttribute:
			for k, v := range collectResourceSchemaPaths(a.NestedObject.Attributes, path) {
				result[k] = v
			}
		case rsschema.SetNestedAttribute:
			for k, v := range collectResourceSchemaPaths(a.NestedObject.Attributes, path) {
				result[k] = v
			}
		case rsschema.MapNestedAttribute:
			for k, v := range collectResourceSchemaPaths(a.NestedObject.Attributes, path) {
				result[k] = v
			}
		}
	}

	return result
}

func collectDataSourceSchemaPaths(attrs map[string]dsschema.Attribute, prefix string) map[string]struct{} {
	result := make(map[string]struct{})

	for name, attr := range attrs {
		path := joinPath(prefix, name)
		result[path] = struct{}{}

		switch a := attr.(type) {
		case dsschema.SingleNestedAttribute:
			for k, v := range collectDataSourceSchemaPaths(a.Attributes, path) {
				result[k] = v
			}
		case dsschema.ListNestedAttribute:
			for k, v := range collectDataSourceSchemaPaths(a.NestedObject.Attributes, path) {
				result[k] = v
			}
		case dsschema.SetNestedAttribute:
			for k, v := range collectDataSourceSchemaPaths(a.NestedObject.Attributes, path) {
				result[k] = v
			}
		case dsschema.MapNestedAttribute:
			for k, v := range collectDataSourceSchemaPaths(a.NestedObject.Attributes, path) {
				result[k] = v
			}
		}
	}

	return result
}

func joinPath(prefix, name string) string {
	if prefix == "" {
		return name
	}
	return prefix + "." + name
}

func diffPathSets(left, right map[string]struct{}) ([]string, []string) {
	var onlyLeft []string
	var onlyRight []string

	for k := range left {
		if _, ok := right[k]; !ok {
			onlyLeft = append(onlyLeft, k)
		}
	}

	for k := range right {
		if _, ok := left[k]; !ok {
			onlyRight = append(onlyRight, k)
		}
	}

	slices.Sort(onlyLeft)
	slices.Sort(onlyRight)

	return onlyLeft, onlyRight
}
