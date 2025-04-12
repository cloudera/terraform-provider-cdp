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
)

var azureSchemaElements = []SchemaTestCaseStructure{
	{
		name:             "id field must exist and be valid",
		field:            "id",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "polling_options should exist and be valid",
		field:            "polling_options",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "crn should exist and be valid",
		field:            "crn",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "create_private_endpoints should exist and be valid",
		field:            "create_private_endpoints",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "credential_name must exist and be valid",
		field:            "credential_name",
		computed:         false,
		shouldBeRequired: true,
	},
	{
		name:             "description should exist and be valid",
		field:            "description",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "enable_outbound_load_balancer should exist and be valid",
		field:            "enable_outbound_load_balancer",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "enable_tunnel should exist and be valid",
		field:            "enable_tunnel",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "endpoint_access_gateway_scheme should exist and be valid",
		field:            "endpoint_access_gateway_scheme",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "endpoint_access_gateway_subnet_ids should exist and be valid",
		field:            "endpoint_access_gateway_subnet_ids",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "encryption_key_resource_group_name should exist and be valid",
		field:            "encryption_key_resource_group_name",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "encryption_key_url should exist and be valid",
		field:            "encryption_key_url",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "encryption_at_host should exist and be valid",
		field:            "encryption_at_host",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "environment_name should exist and be valid",
		field:            "environment_name",
		computed:         false,
		shouldBeRequired: true,
	},
	{
		name:             "existing_network_params should exist and be valid",
		field:            "existing_network_params",
		computed:         false,
		shouldBeRequired: true,
	},
	{
		name:             "freeipa should exist and be valid",
		field:            "freeipa",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "log_storage should exist and be valid",
		field:            "log_storage",
		computed:         false,
		shouldBeRequired: true,
	},
	{
		name:             "proxy_config_name should exist and be valid",
		field:            "proxy_config_name",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "public_key should exist and be valid",
		field:            "public_key",
		computed:         false,
		shouldBeRequired: true,
	},
	{
		name:             "region should exist and be valid",
		field:            "region",
		computed:         false,
		shouldBeRequired: true,
	},
	{
		name:             "report_deployment_logs should exist and be valid",
		field:            "report_deployment_logs",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "resource_group_name should exist and be valid",
		field:            "resource_group_name",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "security_access should exist and be valid",
		field:            "security_access",
		computed:         false,
		shouldBeRequired: true,
	},
	{
		name:             "status should exist and be valid",
		field:            "status",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "status_reason should exist and be valid",
		field:            "status_reason",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "encryption_user_managed_identity should exist and be valid",
		field:            "encryption_user_managed_identity",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "tags should exist and be valid",
		field:            "tags",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "use_public_ip should exist and be valid",
		field:            "use_public_ip",
		computed:         false,
		shouldBeRequired: true,
	},
	{
		name:             "workload_analytics should exist and be valid",
		field:            "workload_analytics",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "compute_cluster should exist and be valid",
		field:            "compute_cluster",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "delete_options should exist and be valid",
		field:            "delete_options",
		computed:         false,
		shouldBeRequired: false,
	},
	{
		name:             "cascading_delete should exist and be valid",
		field:            "cascading_delete",
		computed:         true,
		shouldBeRequired: false,
	},
}

func TestAzureEnvironmentSchemaSchemaAttributeNumber(t *testing.T) {
	expected := 32
	if len(AzureEnvironmentSchema.Attributes) != expected {
		t.Errorf("The number of fields in the AzureEnvironment schema should be: %d but it is: %d", expected, len(AzureEnvironmentSchema.Attributes))
		t.FailNow()
	}
}

func TestAzureSchemaContainsElements(t *testing.T) {
	for _, test := range azureSchemaElements {
		performResourceSchemaValidation(t, test, AzureEnvironmentSchema.Attributes[test.field])
	}
}
