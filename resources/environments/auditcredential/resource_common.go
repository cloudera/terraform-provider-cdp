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
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

func findAuditCredentialByName(ctx context.Context, cdpClient *cdp.Client, credentialName string) (*environmentsmodels.Credential, error) {
	params := operations.NewListAuditCredentialsParams()
	resp, err := cdpClient.Environments.Operations.ListAuditCredentialsContext(ctx, params)
	if err != nil {
		return nil, err
	}
	if resp.GetPayload() == nil {
		return nil, nil
	}
	for _, c := range resp.GetPayload().Credentials {
		if c != nil && c.CredentialName != nil && *c.CredentialName == credentialName {
			return c, nil
		}
	}
	return nil, nil
}

func mapAuditCredentialResponse(c *environmentsmodels.Credential, id, credentialName, crn, description *types.String) {
	*id = types.StringPointerValue(c.CredentialName)
	*credentialName = types.StringPointerValue(c.CredentialName)
	*crn = types.StringPointerValue(c.Crn)
	if c.Description != "" {
		*description = types.StringValue(c.Description)
	} else {
		*description = types.StringNull()
	}
}

func deleteAuditCredential(ctx context.Context, cdpClient *cdp.Client, credentialName string, diagnostics *diag.Diagnostics, errMsg string) {
	params := operations.NewDeleteAuditCredentialParams()
	params.WithInput(&environmentsmodels.DeleteAuditCredentialRequest{
		CredentialName: &credentialName,
	})
	_, err := cdpClient.Environments.Operations.DeleteAuditCredentialContext(ctx, params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, diagnostics, errMsg)
	}
}

func credentialOrError(credential *environmentsmodels.Credential, diagnostics *diag.Diagnostics, operation string) bool {
	if credential == nil {
		diagnostics.AddError(
			"Unexpected API response",
			"CDP returned an empty credential in response to "+operation)
		return false
	}
	return true
}
