// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package de

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/de/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/de/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var _ resource.Resource = (*serviceResource)(nil)

func NewServiceResource() resource.Resource {
	return &serviceResource{}
}

type serviceResource struct {
	client *cdp.Client
}

func (r *serviceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_de_service"
}

func (r *serviceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = serviceSchema
}

func (r *serviceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *serviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data serviceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.De

	params := operations.NewEnableServiceParamsWithContext(ctx)

	createReq, convDiag := modelToEnableServiceRequest(ctx, &data)

	if convDiag.HasError() {
		resp.Diagnostics.Append(*convDiag...)
		utils.AddDeDiagnosticsError(errors.New("conversion error"), &resp.Diagnostics, "create Service")
		return
	}

	params.WithInput(createReq)

	responseOk, err := client.Operations.EnableService(params)
	if err != nil {
		utils.AddDeDiagnosticsError(err, &resp.Diagnostics, "create Service")
		return
	}

	data.Id = types.StringValue(*responseOk.Payload.Service.ClusterID)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func modelToEnableServiceRequest(ctx context.Context, data *serviceResourceModel) (*models.EnableServiceRequest, *diag.Diagnostics) {
	var diags diag.Diagnostics
	var ret models.EnableServiceRequest

	overrides := make([]*models.ChartValueOverridesRequest, 0, len(data.ChartValueOverrides.Elements()))
	if !data.ChartValueOverrides.IsNull() || len(data.ChartValueOverrides.Elements()) > 0 {
		cvos := make([]*chartValueOverridesRequest, 0, len(data.ChartValueOverrides.Elements()))
		data.ChartValueOverrides.ElementsAs(ctx, &cvos, false)
		for i, cvo := range cvos {
			overrides[i] = &models.ChartValueOverridesRequest{
				ChartName: cvo.ChartName.ValueString(),
				Overrides: cvo.Overrides.ValueString(),
			}
		}
	}

	var conf *models.CustomAzureFilesConfigs
	if !data.CustomAzureFilesConfigs.IsNull() {
		var cfg customAzureFilesConfigs
		objDiag := data.CustomAzureFilesConfigs.As(ctx, &cfg, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
		if objDiag.HasError() {
			for _, v := range objDiag.Errors() {
				tflog.Debug(ctx, "convert CustomAzureFilesConfigs error: "+v.Detail())
			}
			diags.Append(objDiag...)
			return nil, &diags
		}
		conf = &models.CustomAzureFilesConfigs{
			AzureFilesFQDN:     cfg.AzureFilesFQDN.ValueString(),
			ResourceGroup:      cfg.ResourceGroup.ValueStringPointer(),
			StorageAccountName: cfg.StorageAccountName.ValueStringPointer(),
		}
	}

	ret = models.EnableServiceRequest{
		ChartValueOverrides:     overrides,
		CPURequests:             data.CPURequests.ValueString(),
		CustomAzureFilesConfigs: conf,
		DeployPreviousVersion:   data.DeployPreviousVersion.ValueBoolPointer(),
		EnablePrivateNetwork:    data.EnablePrivateNetwork.ValueBool(),
		EnablePublicEndpoint:    data.EnablePublicEndpoint.ValueBoolPointer(),
		EnableWorkloadAnalytics: data.EnableWorkloadAnalytics.ValueBoolPointer(),
		Env:                     data.Env.ValueStringPointer(),
		GpuRequests:             data.GpuRequests.ValueString(),
		InitialInstances:        int32(data.InitialInstances.ValueInt64()),
		InitialSpotInstances:    int32(data.InitialSpotInstances.ValueInt64()),
		InstanceType:            data.InstanceType.ValueStringPointer(),
		LoadbalancerAllowlist:   utils.FromSetValueToStringList(data.LoadbalancerAllowlist),
		MaximumInstances:        utils.Int64To32Pointer(data.MaximumInstances),
		MaximumSpotInstances:    int32(data.MaximumSpotInstances.ValueInt64()),
		MemoryRequests:          data.MemoryRequests.ValueString(),
		MinimumInstances:        utils.Int64To32Pointer(data.MinimumInstances),
		MinimumSpotInstances:    int32(data.MinimumSpotInstances.ValueInt64()),
		Name:                    data.Name.ValueStringPointer(),
		NetworkOutboundType:     data.NetworkOutboundType.ValueString(),
		NfsStorageClass:         data.NfsStorageClass.ValueString(),
		ResourcePool:            data.ResourcePool.ValueString(),
		RootVolumeSize:          int32(data.RootVolumeSize.ValueInt64()),
		SkipValidation:          data.SkipValidation.ValueBoolPointer(),
		Subnets:                 utils.FromSetValueToStringList(data.Subnets),
		Tags:                    utils.FromMapValueToStringMap(data.Tags),
		UseSsd:                  data.UseSsd.ValueBool(),
		WhitelistIps:            utils.FromSetValueToStringList(data.WhitelistIps),
	}

	return &ret, &diags
}

func (r *serviceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data serviceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.De

	params := operations.NewDescribeServiceParamsWithContext(ctx)
	params.WithInput(&models.DescribeServiceRequest{
		ClusterID: data.Id.ValueStringPointer(),
	})

	_, err := client.Operations.DescribeService(params)
	if err != nil {
		utils.AddDeDiagnosticsError(err, &resp.Diagnostics, "read Service")
		if d, ok := err.(*operations.DescribeServiceDefault); ok && d.GetPayload() != nil && d.GetPayload().Code == "NOT_FOUND" {
			resp.Diagnostics.AddWarning("Resource not found on provider", "Service not found, removing from state.")
			tflog.Warn(ctx, "Service not found, removing from state", map[string]interface{}{
				"id": data.Id,
			})
			resp.State.RemoveResource(ctx)
		}
		return
	}
}

func (r *serviceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Warn(ctx, "Update operation is not supported yet.")
}

func (r *serviceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data serviceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	client := r.client.De

	force := false

	params := operations.NewDisableServiceParamsWithContext(ctx)
	params.WithInput(&models.DisableServiceRequest{
		ClusterID: data.Id.ValueStringPointer(),
		Force:     &force,
	})

	_, err := client.Operations.DisableService(params)
	if err != nil {
		utils.AddDeDiagnosticsError(err, &resp.Diagnostics, "delete Service")
		return
	}
}
