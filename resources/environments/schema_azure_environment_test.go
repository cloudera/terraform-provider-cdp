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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"testing"
)

var schemaElements = []SchemaTestCaseStructure{
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
		computed:         false,
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
		computed:         true,
		shouldBeRequired: false,
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
		name:             "new_network_params should exist and be valid",
		field:            "new_network_params",
		computed:         true,
		shouldBeRequired: false,
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
}

func TestSchemaContainsElements(t *testing.T) {
	for _, test := range schemaElements {
		performValidation(t, test, AzureEnvironmentSchema.Attributes[test.field])
	}
}

func performValidation(t *testing.T, test SchemaTestCaseStructure, attr schema.Attribute) {
	t.Run(test.name, func(t *testing.T) {
		if attr == nil {
			t.Errorf("The following field does not exists, however it should: " + test.field)
			t.FailNow()
		}
		if attr.IsRequired() != test.shouldBeRequired {
			t.Errorf("The '%s' filed's >required< property should be: %t", test.field, test.shouldBeRequired)
		}
		if attr.IsComputed() != test.computed {
			t.Errorf("The '%s' filed's >computed< property should be: %t", test.field, test.computed)
		}
	})
}
