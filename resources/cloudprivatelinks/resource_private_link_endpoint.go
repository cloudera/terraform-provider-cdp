// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package cloudprivatelinks

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	cdp "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/cloudprivatelinks/client/operations"
	cloudprivatelinkmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/cloudprivatelinks/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var _ resource.Resource = &privateLinkEndpointResource{}

func NewPrivateLinkEndpointResource() resource.Resource {
	return &privateLinkEndpointResource{}
}

type privateLinkEndpointResource struct {
	client *cdp.Client
}

func (r *privateLinkEndpointResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloudprivatelinks_private_link_endpoint"
}

func (r *privateLinkEndpointResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *privateLinkEndpointResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var model cloudPrivateLinkResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}

	csp := model.CloudServiceProvider.ValueString()
	var createReq *cloudprivatelinkmodels.CreatePrivateLinkEndpointRequest

	switch csp {
	case "AWS":
		createReq = fromModelToAwsRequest(model, ctx)
	case "AZURE":
		createReq = fromModelToAzureRequest(model, ctx)
	default:
		resp.Diagnostics.AddError("Unsupported CSP", fmt.Sprintf("cloud_service_provider %q is not supported", csp))
		return
	}

	params := operations.NewCreatePrivateLinkEndpointParamsWithContext(ctx).WithInput(createReq)
	result, err := r.client.Cloudprivatelinks.Operations.CreatePrivateLinkEndpoint(params)
	if err != nil {
		resp.Diagnostics.AddError("Create failed", fmt.Sprintf("Error creating Private Link endpoint: %s", err))
		return
	}

	trackingID := result.GetPayload().TrackingID
	model.TrackingID = types.StringValue(trackingID)

	tflog.Debug(ctx, "Waiting for Private Link endpoint", map[string]interface{}{"trackingId": trackingID})
	opsClient, ok := r.client.Cloudprivatelinks.Operations.(*operations.Client)
	if !ok {
		resp.Diagnostics.AddError("Client error", "Failed to cast operations client")
		return
	}
	_, err = waitForEndpointReady(ctx, opsClient, trackingID)
	if err != nil {
		resp.Diagnostics.AddError("Endpoint not ready", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *privateLinkEndpointResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var model cloudPrivateLinkResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read is a no-op if the API has no get-by-ID endpoint; state is set on Create
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *privateLinkEndpointResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "Private Link endpoints cannot be updated in place")
}

func (r *privateLinkEndpointResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var model cloudPrivateLinkResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}

	csp := model.CloudServiceProvider.ValueString()
	var deleteReq *cloudprivatelinkmodels.DeletePrivateLinkEndpointRequest

	switch csp {
	case "AWS":
		awsReq := fromModelToAwsRequest(model, ctx)
		cspVal := cloudprivatelinkmodels.CloudServiceProvider(csp)
		deleteReq = &cloudprivatelinkmodels.DeletePrivateLinkEndpointRequest{
			CloudServiceProvider: &cspVal,
			AwsAccountInfo: &cloudprivatelinkmodels.AWSAccountInfo{
				CredentialCrn:           awsReq.AwsAccountDetails.CredentialCrn,
				CrossAccountRoleDetails: awsReq.AwsAccountDetails.CrossAccountRoleDetails,
				Region:                  awsReq.AwsAccountDetails.Region,
				VpcID:                   awsReq.AwsAccountDetails.VpcID,
			},
		}
	case "AZURE":
		azureReq := fromModelToAzureRequest(model, ctx)
		cspVal := cloudprivatelinkmodels.CloudServiceProvider(csp)
		deleteReq = &cloudprivatelinkmodels.DeletePrivateLinkEndpointRequest{
			CloudServiceProvider: &cspVal,
			AzureAccountInfo: &cloudprivatelinkmodels.AzureAccountInfo{
				AzureClientSecretCredential: azureReq.AzureAccountDetails.AzureClientSecretCredential,
				CredentialCrn:               azureReq.AzureAccountDetails.CredentialCrn,
				Location:                    azureReq.AzureAccountDetails.Location,
				VNetID:                      azureReq.AzureAccountDetails.VNetID,
			},
		}
	default:
		resp.Diagnostics.AddError("Unsupported CSP", fmt.Sprintf("cloud_service_provider %q is not supported", csp))
		return
	}

	params := operations.NewDeletePrivateLinkEndpointParamsWithContext(ctx).WithInput(deleteReq)
	_, err := r.client.Cloudprivatelinks.Operations.DeletePrivateLinkEndpoint(params)
	if err != nil {
		resp.Diagnostics.AddError("Delete failed", fmt.Sprintf("Error deleting Private Link endpoint: %s", err))
	}
}
