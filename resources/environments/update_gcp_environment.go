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

	"github.com/hashicorp/terraform-plugin-framework/resource"

	environmentsclient "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
)

func updateGcpEnvironment(ctx context.Context, plan *gcpEnvironmentResourceModel, state *gcpEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return executeUpdateOperations(ctx, plan, state, client, resp,
		updateGcpCredentialIfChanged,
		updateGcpSshKeyIfChanged,
		updateGcpProxyConfigurationIfChanged,
		updateGcpCatalogIfChanged,
		updateGcpEndpointAccessGatewayIfChanged,
		updateGcpCustomDockerRegistryIfChanged,
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

func updateGcpCustomDockerRegistryIfChanged(ctx context.Context, plan *gcpEnvironmentResourceModel, state *gcpEnvironmentResourceModel, client *environmentsclient.Environments, resp *resource.UpdateResponse) *resource.UpdateResponse {
	return updateCustomDockerRegistryIfChanged(ctx, client, state.CustomDockerRegistry, plan.CustomDockerRegistry, plan.EnvironmentName.ValueStringPointer(), resp)
}
