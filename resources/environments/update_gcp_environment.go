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
	"context"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	environmentsmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

func updateGcpEnvironment(ctx context.Context, plan *gcpEnvironmentResourceModel, state *gcpEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return executeUpdateOperations(ctx, plan, state, client, resp,
		updateGcpEndpointAccessGatewayIfChanged,
		updateGcpCustomDockerRegistryIfChanged,
		updateGcpProxyConfigurationIfChanged,
		updateGcpAvailabilityZonesIfChanged,
		updateGcpTelemetryFeaturesIfChanged,
		updateGcpSecurityAccessIfChanged,
		updateGcpCredentialIfChanged,
		updateGcpCatalogIfChanged,
		updateGcpSshKeyIfChanged,
	)
}

func updateGcpCredentialIfChanged(ctx context.Context, plan *gcpEnvironmentResourceModel, state *gcpEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateCredential(ctx, client, plan.CredentialName, &state.CredentialName, plan.EnvironmentName.ValueStringPointer(), resp)
}

func updateGcpSshKeyIfChanged(ctx context.Context, plan *gcpEnvironmentResourceModel, state *gcpEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateSshKeyIfChanged(ctx, client, plan.PublicKey, &state.PublicKey, plan.EnvironmentName.ValueStringPointer(), resp)
}

func updateGcpProxyConfigurationIfChanged(ctx context.Context, plan *gcpEnvironmentResourceModel, state *gcpEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateProxyConfigurationIfChanged(ctx, client, &state.ProxyConfigName, &plan.ProxyConfigName, plan.EnvironmentName.ValueStringPointer(), resp)
}

func updateGcpCatalogIfChanged(ctx context.Context, plan *gcpEnvironmentResourceModel, state *gcpEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateCatalogIfChanged(ctx, plan.FreeIpa, &state.FreeIpa, plan.EnvironmentName.ValueString(), client, resp)
}

func updateGcpEndpointAccessGatewayIfChanged(ctx context.Context, plan *gcpEnvironmentResourceModel, state *gcpEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateEndpointAccessGatewayIfChanged(ctx, client, plan.EndpointAccessGatewayScheme, plan.EndpointAccessGatewaySubnetIds, &state.EndpointAccessGatewayScheme, &state.EndpointAccessGatewaySubnetIds, plan.EnvironmentName.ValueString(), plan.PollingOptions, resp)
}

func updateGcpSecurityAccessIfChanged(ctx context.Context, plan *gcpEnvironmentResourceModel, state *gcpEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	if plan.SecurityAccess == nil {
		return resp
	}
	return updateSecurityAccessIfChanged(ctx, client,
		plan.SecurityAccess.DefaultSecurityGroupId,
		plan.SecurityAccess.SecurityGroupIdForKnox,
		&state.SecurityAccess.DefaultSecurityGroupId,
		&state.SecurityAccess.SecurityGroupIdForKnox,
		plan.EnvironmentName.ValueStringPointer(), resp)
}

func updateGcpCustomDockerRegistryIfChanged(ctx context.Context, plan *gcpEnvironmentResourceModel, state *gcpEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateCustomDockerRegistryIfChanged(ctx, client, state.CustomDockerRegistry, plan.CustomDockerRegistry, plan.EnvironmentName.ValueStringPointer(), resp)
}

func updateGcpTelemetryFeaturesIfChanged(ctx context.Context, plan *gcpEnvironmentResourceModel, state *gcpEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateTelemetryFeaturesIfChanged(ctx, client, plan.WorkloadAnalytics, &state.WorkloadAnalytics, plan.EnvironmentName.ValueStringPointer(), resp)
}

func updateGcpAvailabilityZonesIfChanged(ctx context.Context, plan *gcpEnvironmentResourceModel, state *gcpEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	if plan.AvailabilityZones == nil || reflect.DeepEqual(plan.AvailabilityZones, state.AvailabilityZones) {
		return resp
	}
	if len(plan.AvailabilityZones) == 0 {
		resp.Diagnostics.AddError("Invalid availability zone setup", "availability_zones must be a non-empty, known value.")
		return resp
	}
	tflog.Info(ctx, fmt.Sprintf("Updating GCP availability zones for environment '%s'", plan.EnvironmentName.ValueString()))
	request := environmentsmodels.UpdateGcpAvailabilityZonesRequest{
		AvailabilityZones: utils.FromTfStringSliceToStringSlice(plan.AvailabilityZones),
		Environment:       plan.EnvironmentName.ValueStringPointer(),
	}
	params := operations.NewUpdateGcpAvailabilityZonesParams()
	params.WithInput(&request)
	_, err := client.Operations.UpdateGcpAvailabilityZonesContext(ctx, params)
	if err != nil {
		utils.AddEnvironmentDiagnosticsError(err, &resp.Diagnostics, "update GCP availability zones")
	} else {
		state.AvailabilityZones = plan.AvailabilityZones
	}
	return resp
}
