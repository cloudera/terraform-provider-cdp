// Copyright 2025 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package service

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/df/client/operations"
	dfmodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/df/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ resource.Resource                = (*dfFlowDefinitionResource)(nil)
	_ resource.ResourceWithConfigure   = (*dfFlowDefinitionResource)(nil)
	_ resource.ResourceWithImportState = (*dfFlowDefinitionResource)(nil)
)

type dfFlowDefinitionResource struct {
	client *cdp.Client
}

type flowDefinitionModel struct {
	ID              types.String `tfsdk:"id"`
	Name            types.String `tfsdk:"name"`
	File            types.String `tfsdk:"file"`
	FileSha         types.String `tfsdk:"file_sha"`
	Description     types.String `tfsdk:"description"`
	Comments        types.String `tfsdk:"comments"`
	CollectionCrn   types.String `tfsdk:"collection_crn"`
	Crn             types.String `tfsdk:"crn"`
	FlowVersionCrn  types.String `tfsdk:"flow_version_crn"`
	VersionCount    types.Int32  `tfsdk:"version_count"`
}

func NewDfFlowDefinitionResource() resource.Resource {
	return &dfFlowDefinitionResource{}
}

func (r *dfFlowDefinitionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_df_flow_definition"
}

func (r *dfFlowDefinitionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func (r *dfFlowDefinitionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Imports a NiFi flow definition into the DataFlow catalog.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Required: true, MarkdownDescription: "The name of the flow definition.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"file": schema.StringAttribute{
				Required: true, Sensitive: true,
				MarkdownDescription: "The NiFi flow definition JSON content to import. Updating this uploads a new version. Content is hidden from plan output; use `file_sha` to detect changes.",
			},
			"file_sha": schema.StringAttribute{
				Computed: true, MarkdownDescription: "SHA256 hash of the file content. Changes when the flow definition file is modified.",
			},
			"description": schema.StringAttribute{
				Optional: true, MarkdownDescription: "The description of the flow definition.",
			},
			"comments": schema.StringAttribute{
				Optional: true, MarkdownDescription: "Comments for the flow definition version.",
			},
			"collection_crn": schema.StringAttribute{
				Optional: true, MarkdownDescription: "The CRN of the collection to assign the flow definition to. If unspecified, the flow will not be assigned to a collection.",
			},
			"crn": schema.StringAttribute{
				Computed: true, MarkdownDescription: "The CRN of the flow definition.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"flow_version_crn": schema.StringAttribute{
				Computed: true, MarkdownDescription: "The CRN of the latest flow definition version. Use this to reference the flow in a cdp_df_deployment.",
			},
			"version_count": schema.Int32Attribute{
				Computed: true, MarkdownDescription: "The number of versions.",
			},
		},
	}
}

func (r *dfFlowDefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan flowDefinitionModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	jsonParams := map[string]string{
		"file": "flow.json",
		"name": plan.Name.ValueString(),
	}
	if !plan.Description.IsNull() && plan.Description.ValueString() != "" {
		jsonParams["description"] = plan.Description.ValueString()
	}
	if !plan.Comments.IsNull() && plan.Comments.ValueString() != "" {
		jsonParams["comments"] = plan.Comments.ValueString()
	}
	if !plan.CollectionCrn.IsNull() && plan.CollectionCrn.ValueString() != "" {
		jsonParams["collectionCrn"] = plan.CollectionCrn.ValueString()
	}

	headers := map[string]string{}
	if v := plan.Name.ValueString(); v != "" {
		headers["Flow-Definition-Name"] = url.PathEscape(v)
	}
	if !plan.Description.IsNull() && plan.Description.ValueString() != "" {
		headers["Flow-Definition-Description"] = url.PathEscape(plan.Description.ValueString())
	}
	if !plan.Comments.IsNull() && plan.Comments.ValueString() != "" {
		headers["Flow-Definition-Comments"] = url.PathEscape(plan.Comments.ValueString())
	}
	if !plan.CollectionCrn.IsNull() && plan.CollectionCrn.ValueString() != "" {
		headers["Flow-Definition-Collection-Identifier"] = url.PathEscape(plan.CollectionCrn.ValueString())
	}

	respBody, err := r.dfUpload(ctx, "/api/v1/df/importFlowDefinition", jsonParams, headers, plan.File.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error importing flow definition", err.Error())
		return
	}

	var result struct {
		Crn          string `json:"crn"`
		Name         string `json:"name"`
		VersionCount int32  `json:"versionCount"`
		Versions     []struct {
			Crn     string `json:"crn"`
			Version int32  `json:"version"`
		} `json:"versions"`
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		resp.Diagnostics.AddError("Error parsing response", fmt.Sprintf("%s (body: %s)", err, string(respBody)))
		return
	}

	plan.Crn = types.StringValue(result.Crn)
	plan.ID = types.StringValue(result.Crn)
	plan.Name = types.StringValue(result.Name)
	plan.VersionCount = types.Int32Value(result.VersionCount)
	plan.FlowVersionCrn = types.StringValue(latestVersionCrn(result.Versions))
	computeFlowFileSha(&plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// computeFlowFileSha sets the SHA256 hash of the file content.
func computeFlowFileSha(m *flowDefinitionModel) {
	if m.File.IsNull() || m.File.ValueString() == "" {
		m.FileSha = types.StringValue("")
		return
	}
	hash := sha256.Sum256([]byte(m.File.ValueString()))
	m.FileSha = types.StringValue(fmt.Sprintf("%x", hash))
}

func (r *dfFlowDefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state flowDefinitionModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewDescribeFlowParamsWithContext(ctx).WithInput(&dfmodels.DescribeFlowRequest{
		FlowCrn: state.Crn.ValueStringPointer(),
	})
	result, err := r.client.Df.Operations.DescribeFlow(params)
	if err != nil {
		if strings.Contains(err.Error(), "NOT_FOUND") {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error reading flow definition", err.Error())
		return
	}

	flow := result.GetPayload().FlowDetail
	if flow == nil {
		resp.State.RemoveResource(ctx)
		return
	}
	state.Crn = types.StringPointerValue(flow.Crn)
	state.ID = types.StringPointerValue(flow.Crn)
	state.Name = types.StringPointerValue(flow.Name)
	state.VersionCount = types.Int32PointerValue(flow.VersionCount)
	state.FlowVersionCrn = types.StringValue(latestFlowVersionCrn(flow.Versions))

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *dfFlowDefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan flowDefinitionModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state flowDefinitionModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	jsonParams := map[string]string{
		"file":    "flow.json",
		"flowCrn": state.Crn.ValueString(),
	}
	if !plan.Comments.IsNull() && plan.Comments.ValueString() != "" {
		jsonParams["comments"] = plan.Comments.ValueString()
	}

	headers := map[string]string{}
	if !plan.Comments.IsNull() && plan.Comments.ValueString() != "" {
		headers["Flow-Definition-Comments"] = url.PathEscape(plan.Comments.ValueString())
	}

	_, err := r.dfUpload(ctx, "/api/v1/df/importFlowDefinitionVersion", jsonParams, headers, plan.File.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error uploading new flow definition version", err.Error())
		return
	}

	descParams := operations.NewDescribeFlowParamsWithContext(ctx).WithInput(&dfmodels.DescribeFlowRequest{
		FlowCrn: state.Crn.ValueStringPointer(),
	})
	descResult, errDesc := r.client.Df.Operations.DescribeFlow(descParams)
	if errDesc != nil {
		resp.Diagnostics.AddError("Error reading flow definition after version upload", errDesc.Error())
		return
	}

	flow := descResult.GetPayload().FlowDetail
	plan.Crn = state.Crn
	plan.ID = state.ID
	plan.Name = types.StringPointerValue(flow.Name)
	plan.VersionCount = types.Int32PointerValue(flow.VersionCount)
	plan.FlowVersionCrn = types.StringValue(latestFlowVersionCrn(flow.Versions))
	computeFlowFileSha(&plan)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *dfFlowDefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state flowDefinitionModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := operations.NewDeleteFlowParamsWithContext(ctx).WithInput(&dfmodels.DeleteFlowRequest{
		FlowCrn: state.Crn.ValueStringPointer(),
	})
	if _, err := r.client.Df.Operations.DeleteFlow(params); err != nil {
		if !strings.Contains(err.Error(), "NOT_FOUND") {
			resp.Diagnostics.AddError("Error deleting flow definition", err.Error())
		}
	}
}

func (r *dfFlowDefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// latestVersionCrn returns the CRN of the highest-versioned entry from the import response.
func latestVersionCrn(versions []struct {
	Crn     string `json:"crn"`
	Version int32  `json:"version"`
}) string {
	if len(versions) == 0 {
		return ""
	}
	best := versions[0]
	for _, v := range versions[1:] {
		if v.Version > best.Version {
			best = v
		}
	}
	return best.Crn
}

// latestFlowVersionCrn returns the CRN of the highest-versioned FlowVersion from DescribeFlow.
func latestFlowVersionCrn(versions []*dfmodels.FlowVersion) string {
	if len(versions) == 0 {
		return ""
	}
	best := versions[0]
	for _, v := range versions[1:] {
		if v.Version > best.Version {
			best = v
		}
	}
	if best.Crn == nil {
		return ""
	}
	return *best.Crn
}

// dfUpload implements the two-step CDP flow upload (matching CDP CLI behavior):
// Step 1: POST JSON {"file":"flow.json","name":"..."} to CDP API → get 308 redirect URL
// Step 2: POST raw flow bytes to the redirect URL with metadata in headers
func (r *dfFlowDefinitionResource) dfUpload(ctx context.Context, apiPath string, jsonParams map[string]string, dfHeaders map[string]string, flowContent string) ([]byte, error) {
	credentials, err := r.client.GetCredentials()
	if err != nil {
		return nil, err
	}

	endpoint, err := r.client.GetCdpApiEndpoint()
	if err != nil {
		return nil, err
	}

	logFile, _ := os.OpenFile("/tmp/dfupload.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	defer logFile.Close()
	dbg := log.New(logFile, "", log.LstdFlags)

	// Step 1: Send JSON body to get the 308 redirect URL
	apiURL := strings.TrimRight(endpoint, "/") + apiPath
	jsonBody, _ := json.Marshal(jsonParams)

	dbg.Printf("Step1: URL=%s json_body=%s", apiURL, string(jsonBody))

	req1, err := http.NewRequestWithContext(ctx, "POST", apiURL, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}
	req1.Header.Set("Content-Type", "application/json")
	req1.ContentLength = int64(len(jsonBody))

	date := cdp.FormatDate()
	auth, err := cdp.AuthHeader(ctx, r.client.GetLogger(), credentials.AccessKeyId, credentials.PrivateKey, "POST", apiPath, date)
	if err != nil {
		return nil, fmt.Errorf("failed to sign request: %w", err)
	}
	req1.Header.Set("x-altus-auth", auth)
	req1.Header.Set("x-altus-date", date)

	httpClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp1, err := httpClient.Do(req1)
	if err != nil {
		return nil, err
	}
	defer resp1.Body.Close()

	dbg.Printf("Step1 response: status=%d location=%s", resp1.StatusCode, resp1.Header.Get("Location"))

	if resp1.StatusCode != http.StatusPermanentRedirect {
		respBody, _ := io.ReadAll(resp1.Body)
		if resp1.StatusCode >= 400 {
			return nil, fmt.Errorf("API returned status %d: %s", resp1.StatusCode, string(respBody))
		}
		return respBody, nil
	}

	location := resp1.Header.Get("Location")
	if location == "" {
		return nil, fmt.Errorf("308 redirect with no Location header")
	}

	// Step 2: POST raw flow content to the redirect URL with metadata headers
	redirectURL, err := url.Parse(location)
	if err != nil {
		return nil, err
	}

	flowBytes := []byte(flowContent)
	req2, err := http.NewRequestWithContext(ctx, "POST", location, bytes.NewReader(flowBytes))
	if err != nil {
		return nil, err
	}
	req2.Header.Set("Content-Type", "application/json")
	req2.ContentLength = int64(len(flowBytes))
	for k, v := range dfHeaders {
		if v != "" {
			req2.Header.Set(k, v)
		}
	}

	date = cdp.FormatDate()
	redirectPath := redirectURL.Path
	if redirectURL.RawQuery != "" {
		redirectPath = redirectPath + "?" + redirectURL.RawQuery
	}
	auth, err = cdp.AuthHeader(ctx, r.client.GetLogger(), credentials.AccessKeyId, credentials.PrivateKey, "POST", redirectPath, date)
	if err != nil {
		return nil, fmt.Errorf("failed to sign redirect request: %w", err)
	}
	req2.Header.Set("x-altus-auth", auth)
	req2.Header.Set("x-altus-date", date)

	dbg.Printf("Step2: URL=%s body_len=%d headers=%v", location, len(flowBytes), dfHeaders)

	resp2, err := httpClient.Do(req2)
	if err != nil {
		return nil, err
	}
	defer resp2.Body.Close()

	respBody, err := io.ReadAll(resp2.Body)
	if err != nil {
		return nil, err
	}

	dbg.Printf("Step2 response: status=%d body=%s", resp2.StatusCode, string(respBody[:min(500, len(respBody))]))

	if resp2.StatusCode >= 400 {
		return nil, fmt.Errorf("API returned status %d: %s", resp2.StatusCode, string(respBody))
	}

	return respBody, nil
}
