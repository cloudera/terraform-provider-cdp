// Copyright 2023 Cloudera. All Rights Reserved.
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
	"errors"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ resource.Resource = &proxyConfigurationResource{}
)

type proxyConfigurationResource struct {
	client *cdp.Client
}

func NewProxyConfigurationResource() resource.Resource {
	return &proxyConfigurationResource{}
}

func (p *proxyConfigurationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environments_proxy_configuration"
}

func (p *proxyConfigurationResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ProxyConfigurationSchema
}

func (p *proxyConfigurationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	p.client = utils.GetCdpClientForResource(req, resp)
}

func joinHostSet(hosts basetypes.SetValue) string {
	if hosts.IsNull() || hosts.IsUnknown() {
		return ""
	}
	str := make([]string, len(hosts.Elements()))
	for i, elem := range hosts.Elements() {
		str[i] = elem.(types.String).ValueString()
	}
	return strings.Join(str, ",")
}

func (p *proxyConfigurationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data proxyConfigurationResourceModel
	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Got Error while trying to get plan")
		return
	}

	client := p.client.Environments

	params := operations.NewCreateProxyConfigParamsWithContext(ctx)
	params.WithInput(&models.CreateProxyConfigRequest{
		ProxyConfigName: data.Name.ValueStringPointer(),
		Description:     data.Description.ValueString(),
		Protocol:        data.Protocol.ValueStringPointer(),
		Host:            data.Host.ValueStringPointer(),
		Port:            func(i int32) *int32 { return &i }(int32(data.Port.ValueInt64())),
		NoProxyHosts:    joinHostSet(data.NoProxyHosts),
		User:            data.User.ValueString(),
		Password:        data.Password.ValueString(),
	})

	responseOk, err := client.Operations.CreateProxyConfig(params)
	if err != nil {
		msg := err.Error()
		if d, ok := err.(utils.EnvironmentErrorPayload); ok && d.GetPayload() != nil {
			msg = d.GetPayload().Message
		}

		resp.Diagnostics.AddError(
			"Create Proxy Configuration",
			"Failed to create proxy configuration, unexpected error: "+msg,
		)
		return
	}

	proxyConfig := responseOk.Payload.ProxyConfig
	data.ID = types.StringValue(*proxyConfig.Crn)

	diags = resp.State.Set(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (p *proxyConfigurationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state proxyConfigurationResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Got Error while trying to get state")
		return
	}

	client := p.client.Environments

	params := operations.NewDeleteProxyConfigParamsWithContext(ctx)
	params.WithInput(&models.DeleteProxyConfigRequest{ProxyConfigName: state.Name.ValueStringPointer()})

	_, err := client.Operations.DeleteProxyConfig(params)
	if err != nil {
		if envErr, ok := err.(*operations.DeleteProxyConfigDefault); ok {
			if cdp.IsEnvironmentsError(envErr.GetPayload(), "NOT_FOUND", "") {
				return
			} else {
				resp.Diagnostics.AddError(
					"Delete Proxy Configuration",
					"Failed to delete proxy configuration: "+envErr.Payload.Message,
				)
				return
			}
		}
		msg := err.Error()
		if d, ok := err.(utils.EnvironmentErrorPayload); ok && d.GetPayload() != nil {
			msg = d.GetPayload().Message
		}

		resp.Diagnostics.AddError(
			"Delete Proxy Configuration",
			"Failed to delete proxy configuration, unexpected error: "+msg,
		)
	}
}

func splitHostsToSet(hostsStr string, diags *diag.Diagnostics, ctx context.Context) basetypes.SetValue {
	hosts := strings.Split(hostsStr, ",")

	result, diag := types.SetValueFrom(ctx, types.StringType, hosts)
	diags.Append(diag...)
	return result
}

func FindProxyConfigurationByName(name string, ctx context.Context, client *client.Environments, diags *diag.Diagnostics) (*models.ProxyConfig, error) {
	params := operations.NewListProxyConfigsParamsWithContext(ctx)
	params.WithInput(&models.ListProxyConfigsRequest{ProxyConfigName: name})

	responseOk, err := client.Operations.ListProxyConfigs(params)
	if err != nil {
		if envErr, ok := err.(*operations.DeleteProxyConfigDefault); ok {
			if cdp.IsEnvironmentsError(envErr.GetPayload(), "NOT_FOUND", "") {
				return nil, errors.New("not found")
			} else {
				diags.AddError(
					"Read Proxy Configuration",
					"Failed to read proxy configuration: "+envErr.Payload.Message,
				)
				return nil, errors.New("unknown error")
			}
		}
		msg := err.Error()
		if d, ok := err.(utils.EnvironmentErrorPayload); ok && d.GetPayload() != nil {
			msg = d.GetPayload().Message
		}

		diags.AddError(
			"Read Proxy Configuration",
			"Failed to read proxy configuration, unexpected error: "+msg,
		)

		return nil, errors.New("unknown error")
	}

	if len(responseOk.Payload.ProxyConfigs) > 0 {
		return responseOk.Payload.ProxyConfigs[0], nil
	} else {
		return nil, errors.New("not found")
	}
}

func (p *proxyConfigurationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state proxyConfigurationResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Got Error while trying to get state")
		return
	}

	client := p.client.Environments

	proxyConfig, err := FindProxyConfigurationByName(state.Name.ValueString(), ctx, client, &resp.Diagnostics)
	if err != nil {
		if err.Error() == "not found" {
			removeProxyConfigFromState(ctx, &resp.Diagnostics, &resp.State, state)
		}
		return
	}

	state.Name = types.StringValue(*proxyConfig.ProxyConfigName)
	state.Description = types.StringValue(proxyConfig.Description)
	state.Protocol = types.StringValue(*proxyConfig.Protocol)
	state.Host = types.StringValue(*proxyConfig.Host)
	state.Port = types.Int64Value(int64(*proxyConfig.Port))
	state.NoProxyHosts = splitHostsToSet(proxyConfig.NoProxyHosts, &diags, ctx)
	state.User = types.StringValue(proxyConfig.User)
	state.Password = types.StringValue(proxyConfig.Password)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func removeProxyConfigFromState(ctx context.Context, diag *diag.Diagnostics, state *tfsdk.State, model proxyConfigurationResourceModel) {
	diag.AddWarning("Resource not found on provider", "Proxy configuration not found, removing resource from state.")
	tflog.Warn(ctx, "Proxy configuration not found, removing resource from state", map[string]interface{}{
		"id": model.ID.ValueString(),
	})
	state.RemoveResource(ctx)
}

func (p *proxyConfigurationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddWarning("Update Proxy Configuration Not Supported", "Update proxy configuration is not supported. Plan changes were not applied.")
}
