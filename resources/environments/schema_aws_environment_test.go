// Copyright 2025 Cloudera. All Rights Reserved.
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

var awsSchemaElements = []SchemaTestCaseStructure{
	{
		name:             "id field must exist and be valid",
		field:            "id",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "crn should exist and be valid",
		field:            "crn",
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
		name:             "authentication should exist and be valid",
		field:            "authentication",
		computed:         false,
		shouldBeRequired: true,
	},
	{
		name:             "create_private_subnets should exist and be valid",
		field:            "create_private_subnets",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "create_service_endpoints should exist and be valid",
		field:            "create_service_endpoints",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "s3_guard_table_name should exist and be valid",
		field:            "s3_guard_table_name",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "credential_name should exist and be valid",
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
		name:             "enable_tunnel should exist and be valid",
		field:            "enable_tunnel",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "encryption_key_arn should exist and be valid",
		field:            "encryption_key_arn",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "endpoint_access_gateway_scheme should exist and be valid",
		field:            "endpoint_access_gateway_scheme",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "endpoint_access_gateway_subnet_ids should exist and be valid",
		field:            "endpoint_access_gateway_subnet_ids",
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
		name:             "cascading_delete should exist and be valid",
		field:            "cascading_delete",
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
		name:             "network_cidr should exist and be valid",
		field:            "network_cidr",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "proxy_config_name should exist and be valid",
		field:            "proxy_config_name",
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
		name:             "subnet_ids should exist and be valid",
		field:            "subnet_ids",
		computed:         false,
		shouldBeRequired: true,
	},
	{
		name:             "tags should exist and be valid",
		field:            "tags",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "tunnel_type should exist and be valid",
		field:            "tunnel_type",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "workload_analytics should exist and be valid",
		field:            "workload_analytics",
		computed:         true,
		shouldBeRequired: false,
	},
	{
		name:             "vpc_id should exist and be valid",
		field:            "vpc_id",
		computed:         false,
		shouldBeRequired: true,
	},
}

func TestAwsSchemaContainsElements(t *testing.T) {
	for _, test := range awsSchemaElements {
		performResourceSchemaValidation(t, test, AwsEnvironmentSchema.Attributes[test.field])
	}
}
