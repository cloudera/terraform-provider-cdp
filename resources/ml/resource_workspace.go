// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package ml

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/client/operations"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var _ resource.Resource = (*workspaceResource)(nil)

func NewWorkspaceResource() resource.Resource {
	return &workspaceResource{}
}

type workspaceResource struct {
	client *cdp.Client
}

func (r *workspaceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ml_workspace"
}

func (r *workspaceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = workspaceSchema
}

func (r *workspaceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *workspaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data workspaceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Ml

	params := operations.NewCreateWorkspaceParamsWithContext(ctx)
	createReq, convDiag := modelToCreateWorkspaceRequest(ctx, &data)

	if convDiag.HasError() {
		resp.Diagnostics.Append(*convDiag...)
		utils.AddMlDiagnosticsError(errors.New("conversion error"), &resp.Diagnostics, "create Workspace")
		return
	}

	params.WithInput(createReq)
	_, err := client.Operations.CreateWorkspace(params)
	if err != nil {
		utils.AddMlDiagnosticsError(err, &resp.Diagnostics, "create Workspace")
		return
	}

	data.Id = data.WorkspaceName

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func modelToCreateWorkspaceRequest(ctx context.Context, data *workspaceResourceModel) (*models.CreateWorkspaceRequest, *diag.Diagnostics) {
	var diags diag.Diagnostics

	var exDbCfgReq *models.ExistingDatabaseConfig
	if !data.ExistingDatabaseConfig.IsNull() {
		var exDbCfg ExistingDatabaseConfig
		objDiag := data.ExistingDatabaseConfig.As(ctx, &exDbCfg, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
		if objDiag.HasError() {
			for _, v := range objDiag.Errors() {
				tflog.Debug(ctx, "convert ExistingDatabaseConfig error: "+v.Detail())
			}
			diags.Append(objDiag...)
			return nil, &diags
		}
		exDbCfgReq = &models.ExistingDatabaseConfig{
			ExistingDatabaseHost:     exDbCfg.ExistingDatabaseHost.ValueString(),
			ExistingDatabaseName:     exDbCfg.ExistingDatabaseName.ValueString(),
			ExistingDatabasePassword: exDbCfg.ExistingDatabasePassword.ValueString(),
			ExistingDatabasePort:     exDbCfg.ExistingDatabasePort.ValueString(),
			ExistingDatabaseUser:     exDbCfg.ExistingDatabaseUser.ValueString(),
		}
	}

	obTypes := make([]models.OutboundTypes, 0, len(data.OutboundTypes.Elements()))
	if !data.OutboundTypes.IsNull() {
		for _, v := range data.OutboundTypes.Elements() {
			obTypes = append(obTypes, models.OutboundTypes(v.(types.String).ValueString()))
		}
	}

	var provK8sReq *models.ProvisionK8sRequest
	if !data.ProvisionK8sRequest.IsNull() {
		var provK8s ProvisionK8sRequest
		provDiag := data.ProvisionK8sRequest.As(ctx, &provK8s, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
		if provDiag.HasError() {
			for _, v := range provDiag.Errors() {
				tflog.Debug(ctx, "convert ProvisionK8sRequest error: "+v.Detail())
			}
			diags.Append(provDiag...)
			return nil, &diags
		}

		igReqs := make([]*models.InstanceGroup, 0, len(provK8s.InstanceGroups.Elements()))
		if !provK8s.InstanceGroups.IsNull() {
			igs := make([]InstanceGroup, 0, len(provK8s.InstanceGroups.Elements()))
			igsDiag := provK8s.InstanceGroups.ElementsAs(ctx, &igs, false)
			if igsDiag.HasError() {
				for _, v := range igsDiag.Errors() {
					tflog.Debug(ctx, "convert InstanceGroup slice error: "+v.Detail())
				}
				diags.Append(igsDiag...)
				return nil, &diags
			}

			for _, v := range igs {
				var asReq *models.Autoscaling

				if !v.Autoscaling.IsNull() {
					var as Autoscaling
					asDiag := v.Autoscaling.As(ctx, &as, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
					if asDiag.HasError() {
						for _, v := range asDiag.Errors() {
							tflog.Debug(ctx, "convert Autoscaling error: "+v.Detail())
						}
						diags.Append(asDiag...)
						return nil, &diags
					}

					maxInst := int32(as.MaxInstances.ValueInt64())
					minInst := int32(as.MinInstances.ValueInt64())
					asReq = &models.Autoscaling{
						Enabled:      as.Enabled.ValueBool(),
						MaxInstances: &maxInst,
						MinInstances: &minInst,
					}
				}

				ir := make([]string, 0, len(v.IngressRules.Elements()))
				irDiag := v.IngressRules.ElementsAs(ctx, &ir, false)
				if irDiag.HasError() {
					for _, v := range irDiag.Errors() {
						tflog.Debug(ctx, "convert IngressRules error: "+v.Detail())
					}
					diags.Append(irDiag...)
					return nil, &diags
				}

				var rvReq *models.RootVolume
				if !v.RootVolume.IsNull() {
					var rv RootVolume
					rvDiag := v.RootVolume.As(ctx, &rv, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
					if rvDiag.HasError() {
						for _, v := range rvDiag.Errors() {
							tflog.Debug(ctx, "convert RootVolume error: "+v.Detail())
						}
						diags.Append(rvDiag...)
						return nil, &diags
					}
					rvReq = &models.RootVolume{
						Size: rv.Size.ValueInt64Pointer(),
					}
				}

				igReq := &models.InstanceGroup{
					Autoscaling:   asReq,
					IngressRules:  ir,
					InstanceCount: int32(v.InstanceCount.ValueInt64()),
					InstanceTier:  v.InstanceTier.ValueString(),
					InstanceType:  v.InstanceType.ValueStringPointer(),
					Name:          v.Name.ValueString(),
					RootVolume:    rvReq,
				}
				igReqs = append(igReqs, igReq)
			}
		}

		var onReq *models.OverlayNetwork
		if !provK8s.Network.IsNull() {
			var on OverlayNetwork
			onDiag := provK8s.Network.As(ctx, &on, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
			if onDiag.HasError() {
				for _, v := range onDiag.Errors() {
					tflog.Debug(ctx, "convert OverlayNetwork error: "+v.Detail())
				}
				diags.Append(onDiag...)
				return nil, &diags
			}

			var tReq *models.Topology
			if !on.Topology.IsNull() {
				var t Topology
				tDiag := on.Topology.As(ctx, &t, basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true, UnhandledUnknownAsEmpty: true})
				if tDiag.HasError() {
					for _, v := range tDiag.Errors() {
						tflog.Debug(ctx, "convert Topology error: "+v.Detail())
					}
					diags.Append(tDiag...)
					return nil, &diags
				}

				ss := make([]string, 0, len(t.Subnets.Elements()))
				ssDiag := t.Subnets.ElementsAs(ctx, &ss, false)
				if ssDiag.HasError() {
					for _, v := range ssDiag.Errors() {
						tflog.Debug(ctx, "convert Topology error: "+v.Detail())
					}
					diags.Append(ssDiag...)
					return nil, &diags
				}

				tReq = &models.Topology{
					Subnets: ss,
				}
			}

			onReq = &models.OverlayNetwork{
				Plugin:   on.Plugin.ValueString(),
				Topology: tReq,
			}
		}

		provTagReq := make([]*models.ProvisionTag, 0, len(provK8s.Tags.Elements()))
		if !provK8s.Tags.IsNull() {
			tags := make([]ProvisionTag, 0, len(provK8s.Tags.Elements()))
			tagsDiag := provK8s.Tags.ElementsAs(ctx, &tags, false)
			if tagsDiag.HasError() {
				for _, v := range tagsDiag.Errors() {
					tflog.Debug(ctx, "convert Topology error: "+v.Detail())
				}
				diags.Append(tagsDiag...)
				return nil, &diags
			}

			for _, v := range tags {
				provTagReq = append(provTagReq, &models.ProvisionTag{
					Key:   v.Key.ValueStringPointer(),
					Value: v.Value.ValueStringPointer(),
				})
			}
		}

		provK8sReq = &models.ProvisionK8sRequest{
			EnvironmentName: provK8s.EnvironmentName.ValueStringPointer(),
			InstanceGroups:  igReqs,
			Network:         onReq,
			Tags:            provTagReq,
		}
	}

	return &models.CreateWorkspaceRequest{
		AuthorizedIPRanges:          utils.FromSetValueToStringList(data.AuthorizedIPRanges),
		CdswMigrationMode:           data.CdswMigrationMode.ValueString(),
		DisableTLS:                  data.DisableTLS.ValueBool(),
		EnableGovernance:            data.EnableGovernance.ValueBool(),
		EnableModelMetrics:          data.EnableModelMetrics.ValueBool(),
		EnableMonitoring:            data.EnableMonitoring.ValueBool(),
		EnvironmentName:             data.EnvironmentName.ValueStringPointer(),
		ExistingDatabaseConfig:      exDbCfgReq,
		ExistingNFS:                 data.ExistingNFS.ValueString(),
		LoadBalancerIPWhitelists:    utils.FromSetValueToStringList(data.LoadBalancerIPWhitelists),
		MlVersion:                   data.MlVersion.ValueString(),
		NfsVersion:                  data.NfsVersion.ValueString(),
		OutboundTypes:               obTypes,
		PrivateCluster:              data.PrivateCluster.ValueBool(),
		ProvisionK8sRequest:         provK8sReq,
		SkipValidation:              data.SkipValidation.ValueBool(),
		StaticSubdomain:             data.StaticSubdomain.ValueString(),
		SubnetsForLoadBalancers:     utils.FromSetValueToStringList(data.SubnetsForLoadBalancers),
		UsePublicLoadBalancer:       data.UsePublicLoadBalancer.ValueBool(),
		WhitelistAuthorizedIPRanges: data.WhitelistAuthorizedIPRanges.ValueBool(),
		WorkspaceName:               data.WorkspaceName.ValueStringPointer(),
	}, &diags
}

func (r *workspaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data workspaceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Ml

	params := operations.NewDescribeWorkspaceParamsWithContext(ctx)
	params.WithInput(&models.DescribeWorkspaceRequest{
		EnvironmentName: data.EnvironmentName.ValueString(),
		WorkspaceName:   data.WorkspaceName.ValueString(),
	})

	_, err := client.Operations.DescribeWorkspace(params)
	if err != nil {
		utils.AddMlDiagnosticsError(err, &resp.Diagnostics, "read Workspace")
		if d, ok := err.(*operations.DescribeWorkspaceDefault); ok && d.GetPayload() != nil && d.GetPayload().Code == "NOT_FOUND" {
			resp.Diagnostics.AddWarning("Resource not found on provider", "Workspace not found, removing from state.")
			tflog.Warn(ctx, "Workspace not found, removing from state", map[string]interface{}{
				"id": data.Id,
			})
			resp.State.RemoveResource(ctx)
		}
		return
	}
}

func (r *workspaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Warn(ctx, "Update operation is not supported yet.")
}

func (r *workspaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data workspaceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	client := r.client.Ml

	force := false

	params := operations.NewDeleteWorkspaceParamsWithContext(ctx)
	params.WithInput(&models.DeleteWorkspaceRequest{
		EnvironmentName: data.EnvironmentName.ValueString(),
		WorkspaceName:   data.WorkspaceName.ValueString(),
		Force:           &force,
	})

	_, err := client.Operations.DeleteWorkspace(params)
	if err != nil {
		utils.AddMlDiagnosticsError(err, &resp.Diagnostics, "delete Workspace")
		return
	}
}
