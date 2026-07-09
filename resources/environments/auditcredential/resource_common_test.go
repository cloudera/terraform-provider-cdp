// Copyright 2026 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package auditcredential

import (
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

func createRawAwsAuditCredentialResource(id string) tftypes.Value {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"id":              tftypes.String,
				"credential_name": tftypes.String,
				"description":     tftypes.String,
				"crn":             tftypes.String,
				"role_arn":        tftypes.String,
			},
		},
		map[string]tftypes.Value{
			"id":              tftypes.NewValue(tftypes.String, id),
			"credential_name": tftypes.NewValue(tftypes.String, "test-audit-cred"),
			"description":     tftypes.NewValue(tftypes.String, nil),
			"crn":             tftypes.NewValue(tftypes.String, "test-crn"),
			"role_arn":        tftypes.NewValue(tftypes.String, "arn:aws:iam::123456789012:role/audit-role"),
		},
	)
}

func createRawAzureAuditCredentialResource(id string) tftypes.Value {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"id":              tftypes.String,
				"credential_name": tftypes.String,
				"description":     tftypes.String,
				"crn":             tftypes.String,
				"subscription_id": tftypes.String,
				"tenant_id":       tftypes.String,
				"app_based": tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"application_id": tftypes.String,
						"secret_key":     tftypes.String,
					},
				},
			},
		},
		map[string]tftypes.Value{
			"id":              tftypes.NewValue(tftypes.String, id),
			"credential_name": tftypes.NewValue(tftypes.String, "test-audit-cred"),
			"description":     tftypes.NewValue(tftypes.String, nil),
			"crn":             tftypes.NewValue(tftypes.String, "test-crn"),
			"subscription_id": tftypes.NewValue(tftypes.String, "sub-id-123"),
			"tenant_id":       tftypes.NewValue(tftypes.String, "tenant-id-456"),
			"app_based": tftypes.NewValue(
				tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"application_id": tftypes.String,
						"secret_key":     tftypes.String,
					},
				},
				map[string]tftypes.Value{
					"application_id": tftypes.NewValue(tftypes.String, "app-id-789"),
					"secret_key":     tftypes.NewValue(tftypes.String, "secret-key-abc"),
				},
			),
		},
	)
}

func createRawAwsGovCloudAuditCredentialResource(id string) tftypes.Value {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"id":              tftypes.String,
				"credential_name": tftypes.String,
				"description":     tftypes.String,
				"crn":             tftypes.String,
				"role_arn":        tftypes.String,
			},
		},
		map[string]tftypes.Value{
			"id":              tftypes.NewValue(tftypes.String, id),
			"credential_name": tftypes.NewValue(tftypes.String, "test-audit-cred"),
			"description":     tftypes.NewValue(tftypes.String, nil),
			"crn":             tftypes.NewValue(tftypes.String, "test-crn"),
			"role_arn":        tftypes.NewValue(tftypes.String, "arn:aws-us-gov:iam::123456789012:role/audit-role"),
		},
	)
}

func createRawGcpAuditCredentialResource(id string) tftypes.Value {
	return tftypes.NewValue(
		tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"id":              tftypes.String,
				"credential_name": tftypes.String,
				"description":     tftypes.String,
				"crn":             tftypes.String,
				"credential_key":  tftypes.String,
			},
		},
		map[string]tftypes.Value{
			"id":              tftypes.NewValue(tftypes.String, id),
			"credential_name": tftypes.NewValue(tftypes.String, "test-audit-cred"),
			"description":     tftypes.NewValue(tftypes.String, nil),
			"crn":             tftypes.NewValue(tftypes.String, "test-crn"),
			"credential_key":  tftypes.NewValue(tftypes.String, "eyJ0ZXN0IjoidmFsdWUifQ=="),
		},
	)
}

func testAuditCredential() *models.Credential {
	return &models.Credential{
		CredentialName: new("test-audit-cred"),
		Crn:            new("crn:cdp:environments:us-west-1:1234:credential:test-audit-cred"),
		CloudPlatform:  new("AWS"),
		AwsCredentialProperties: &models.AwsCredentialProperties{
			RoleArn: "arn:aws:iam::123456789012:role/audit-role",
		},
	}
}

func testAzureAuditCredential() *models.Credential {
	return &models.Credential{
		CredentialName: new("test-audit-cred"),
		Crn:            new("crn:cdp:environments:us-west-1:1234:credential:test-audit-cred"),
		CloudPlatform:  new("AZURE"),
		AzureCredentialProperties: &models.AzureCredentialProperties{
			SubscriptionID: "sub-id-123",
			TenantID:       "tenant-id-456",
			AppID:          "app-id-789",
		},
	}
}

func testGcpAuditCredential() *models.Credential {
	return &models.Credential{
		CredentialName: new("test-audit-cred"),
		Crn:            new("crn:cdp:environments:us-west-1:1234:credential:test-audit-cred"),
		CloudPlatform:  new("GCP"),
	}
}
