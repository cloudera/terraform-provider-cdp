// Copyright 2024 Cloudera. All Rights Reserved.
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
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/cloudera/terraform-provider-cdp/utils/test"
)

var gcpSchemaElements = []test.ResourceSchemaTestCaseStructure{
	{
		Name:             "id field must exist and be valid",
		Field:            "id",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "polling_options should exist and be valid",
		Field:            "polling_options",
		Computed:         false,
		ShouldBeRequired: false,
		AttributeType:    schema.SingleNestedAttribute{},
	},
	{
		Name:             "crn should exist and be valid",
		Field:            "crn",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "credential_name must exist and be valid",
		Field:            "credential_name",
		Computed:         false,
		ShouldBeRequired: true,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "description should exist and be valid",
		Field:            "description",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "enable_tunnel should exist and be valid",
		Field:            "enable_tunnel",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.BoolAttribute{},
	},
	{
		Name:             "endpoint_access_gateway_scheme should exist and be valid",
		Field:            "endpoint_access_gateway_scheme",
		Computed:         false,
		ShouldBeRequired: false,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "environment_name should exist and be valid",
		Field:            "environment_name",
		Computed:         false,
		ShouldBeRequired: true,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "existing_network_params should exist and be valid",
		Field:            "existing_network_params",
		Computed:         false,
		ShouldBeRequired: true,
		AttributeType:    schema.SingleNestedAttribute{},
	},
	{
		Name:             "security_access should exist and be valid",
		Field:            "security_access",
		Computed:         false,
		ShouldBeRequired: false,
		AttributeType:    schema.SingleNestedAttribute{},
	},
	{
		Name:             "freeipa should exist and be valid",
		Field:            "freeipa",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.SingleNestedAttribute{},
	},
	{
		Name:             "log_storage should exist and be valid",
		Field:            "log_storage",
		Computed:         false,
		ShouldBeRequired: false,
		AttributeType:    schema.SingleNestedAttribute{},
	},
	{
		Name:             "region should exist and be valid",
		Field:            "region",
		Computed:         false,
		ShouldBeRequired: true,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "public_key should exist and be valid",
		Field:            "public_key",
		Computed:         false,
		ShouldBeRequired: true,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "report_deployment_logs should exist and be valid",
		Field:            "report_deployment_logs",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.BoolAttribute{},
	},
	{
		Name:             "status should exist and be valid",
		Field:            "status",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "status_reason should exist and be valid",
		Field:            "status_reason",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "encryption_key should exist and be valid",
		Field:            "encryption_key",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "availability_zones should exist and be valid",
		Field:            "availability_zones",
		Computed:         false,
		ShouldBeRequired: false,
		AttributeType:    schema.SetAttribute{},
	},
	{
		Name:             "tags should exist and be valid",
		Field:            "tags",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.MapAttribute{},
	},
	{
		Name:             "proxy_config_name should exist and be valid",
		Field:            "proxy_config_name",
		Computed:         false,
		ShouldBeRequired: false,
		AttributeType:    schema.StringAttribute{},
	},
	{
		Name:             "use_public_ip should exist and be valid",
		Field:            "use_public_ip",
		Computed:         false,
		ShouldBeRequired: true,
		AttributeType:    schema.BoolAttribute{},
	},
	{
		Name:             "workload_analytics should exist and be valid",
		Field:            "workload_analytics",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.BoolAttribute{},
	},
	{
		Name:             "delete_options should exist and be valid",
		Field:            "delete_options",
		Computed:         false,
		ShouldBeRequired: false,
		AttributeType:    schema.SingleNestedAttribute{},
	},
	{
		Name:             "cascading_delete should exist and be valid",
		Field:            "cascading_delete",
		Computed:         true,
		ShouldBeRequired: false,
		AttributeType:    schema.BoolAttribute{},
	},
}

func TestGcpEnvironmentSchemaSchemaAttributeNumber(t *testing.T) {
	expected := 25
	if len(GcpEnvironmentSchema.Attributes) != expected {
		t.Errorf("The number of fields in the AzureEnvironment schema should be: %d but it is: %d", expected, len(GcpEnvironmentSchema.Attributes))
		t.FailNow()
	}
}

func TestGcpSchemaContainsElements(t *testing.T) {
	test.PerformResourceSchemaValidation(t, GcpEnvironmentSchema.Attributes, gcpSchemaElements)
}
