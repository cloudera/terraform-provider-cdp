// Copyright 2023 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package datalake

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/cdp"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client"
	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/client/operations"
	datalakemodels "github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

var (
	_ resource.ResourceWithConfigure   = &awsDatalakeResource{}
	_ resource.ResourceWithImportState = &awsDatalakeResource{}
)

type awsDatalakeResource struct {
	client *cdp.Client
}

func (r *awsDatalakeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func NewAwsDatalakeResource() resource.Resource {
	return &awsDatalakeResource{}
}

func (r *awsDatalakeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_datalake_aws_datalake"
}

func (r *awsDatalakeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = awsDatalakeResourceSchema
}

func (r *awsDatalakeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.client = utils.GetCdpClientForResource(req, resp)
}

func toAwsDatalakeRequest(ctx context.Context, model *awsDatalakeResourceModel) *datalakemodels.CreateAWSDatalakeRequest {
	req := &datalakemodels.CreateAWSDatalakeRequest{}
	req.CloudProviderConfiguration = &datalakemodels.AWSConfigurationRequest{
		InstanceProfile:       model.InstanceProfile.ValueStringPointer(),
		StorageBucketLocation: model.StorageLocationBase.ValueStringPointer(),
	}
	req.DatalakeName = model.DatalakeName.ValueStringPointer()
	req.EnableRangerRaz = model.EnableRangerRaz.ValueBool()
	req.EnvironmentName = model.EnvironmentName.ValueStringPointer()
	if model.Image != nil {
		req.Image = &datalakemodels.ImageRequest{
			CatalogName: model.Image.CatalogName.ValueStringPointer(),
			ID:          model.Image.ID.ValueString(),
			Os:          model.Image.Os.ValueString(),
		}
	}
	req.JavaVersion = int32(model.JavaVersion.ValueInt64())
	req.MultiAz = model.MultiAz.ValueBool()
	req.Recipes = make([]*datalakemodels.InstanceGroupRecipeRequest, len(model.Recipes))
	for i, v := range model.Recipes {
		req.Recipes[i] = &datalakemodels.InstanceGroupRecipeRequest{
			InstanceGroupName: v.InstanceGroupName.ValueStringPointer(),
			RecipeNames:       utils.FromSetValueToStringList(v.RecipeNames),
		}
	}
	req.Runtime = model.Runtime.ValueString()
	req.Scale = datalakemodels.DatalakeScaleType(model.Scale.ValueString())
	if !model.Tags.IsNull() {
		req.Tags = make([]*datalakemodels.DatalakeResourceTagRequest, len(model.Tags.Elements()))
		i := 0
		for k, v := range model.Tags.Elements() {
			key := k
			val, diag := v.(basetypes.StringValuable).ToStringValue(ctx)
			if !diag.HasError() {
				req.Tags[i] = &datalakemodels.DatalakeResourceTagRequest{
					Key:   &key,
					Value: val.ValueStringPointer(),
				}
			}
			i++
		}
	}
	return req
}

func (r *awsDatalakeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state awsDatalakeResourceModel
	diags := req.Plan.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Got error while trying to set plan")
		return
	}

	client := r.client.Datalake

	params := operations.NewCreateAWSDatalakeParamsWithContext(ctx)
	params.WithInput(toAwsDatalakeRequest(ctx, &state))
	responseOk, err := client.Operations.CreateAWSDatalake(params)
	if err != nil {
		utils.AddDatalakeDiagnosticsError(err, &resp.Diagnostics, "create AWS Datalake")
		return
	}

	datalakeResp := responseOk.Payload
	toAwsDatalakeResourceModel(datalakeResp, &state)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !(state.PollingOptions != nil && state.PollingOptions.Async.ValueBool()) {
		stateSaver := func(dlDtl *datalakemodels.DatalakeDetails) {
			datalakeDetailsToAwsDatalakeResourceModel(ctx, dlDtl, &state, state.PollingOptions, &resp.Diagnostics)
			diags = resp.State.Set(ctx, state)
			resp.Diagnostics.Append(diags...)
		}
		if err := waitForDatalakeToBeRunning(ctx, state.DatalakeName.ValueString(), time.Hour, callFailureThreshold, r.client.Datalake, state.PollingOptions, stateSaver); err != nil {
			utils.AddDatalakeDiagnosticsError(err, &resp.Diagnostics, "create AWS Datalake")
			return
		}
	}

	descParams := operations.NewDescribeDatalakeParamsWithContext(ctx)
	descParams.WithInput(&datalakemodels.DescribeDatalakeRequest{DatalakeName: state.DatalakeName.ValueStringPointer()})
	descResponseOk, err := client.Operations.DescribeDatalake(descParams)
	if err != nil {
		utils.AddDatalakeDiagnosticsError(err, &resp.Diagnostics, "create AWS Datalake")
		return
	}

	descDlResp := descResponseOk.Payload
	datalakeDetailsToAwsDatalakeResourceModel(ctx, descDlResp.Datalake, &state, state.PollingOptions, &resp.Diagnostics)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func waitForDatalakeToBeRunning(ctx context.Context, datalakeName string, fallbackPollingTimeout time.Duration, callFailureThresholdDefault int, client *client.Datalake, options *utils.PollingOptions,
	stateSaverCb func(*datalakemodels.DatalakeDetails)) error {
	timeout, err := utils.CalculateTimeoutOrDefault(ctx, options, fallbackPollingTimeout)
	if err != nil {
		return err
	}
	callFailureThreshold, failureThresholdError := utils.CalculateCallFailureThresholdOrDefault(ctx, options, callFailureThresholdDefault)
	if failureThresholdError != nil {
		return failureThresholdError
	}
	callFailedCount := 0
	stateConf := &retry.StateChangeConf{
		Pending: []string{"REQUESTED", "WAIT_FOR_ENVIRONMENT", "ENVIRONMENT_CREATED", "STACK_CREATION_IN_PROGRESS",
			"STACK_CREATION_FINISHED", "EXTERNAL_DATABASE_CREATION_IN_PROGRESS", "EXTERNAL_DATABASE_CREATED",
		},
		Target:       []string{"RUNNING"},
		Delay:        5 * time.Second,
		Timeout:      *timeout,
		PollInterval: 10 * time.Second,
		Refresh: func() (interface{}, string, error) {
			log.Printf("About to describe datalake")
			params := operations.NewDescribeDatalakeParamsWithContext(ctx)
			params.WithInput(&datalakemodels.DescribeDatalakeRequest{DatalakeName: &datalakeName})
			resp, err := client.Operations.DescribeDatalake(params)
			if err != nil {
				callFailedCount++
				if callFailedCount <= callFailureThreshold {
					tflog.Warn(ctx, fmt.Sprintf("Error describing datalake with call failure due to [%s] but threshold limit is not reached yet (%d out of %d).", err.Error(), callFailedCount, callFailureThreshold))
					return nil, "", nil
				}
				tflog.Error(ctx, fmt.Sprintf("Error describing datalake (due to: %s) and call failure threshold limit exceeded.", err))
				return nil, "", err
			}
			callFailedCount = 0
			stateSaverCb(resp.Payload.Datalake)
			log.Printf("Described datalake: %s", resp.GetPayload().Datalake.Status)
			return checkResponseStatusForError(resp)
		},
	}
	_, err = stateConf.WaitForStateContext(ctx)
	return err
}

func checkResponseStatusForError(resp *operations.DescribeDatalakeOK) (interface{}, string, error) {
	if utils.ContainsAsSubstring([]string{"FAILED", "ERROR"}, resp.GetPayload().Datalake.Status) {
		return nil, "", fmt.Errorf("unexpected Data Lake status: %s. Reason: %s", resp.GetPayload().Datalake.Status, resp.GetPayload().Datalake.StatusReason)
	}
	return resp, resp.GetPayload().Datalake.Status, nil
}

func toAwsDatalakeResourceModel(resp *datalakemodels.CreateAWSDatalakeResponse, model *awsDatalakeResourceModel) {
	model.ID = types.StringPointerValue(resp.Datalake.DatalakeName)
	model.CertificateExpirationState = types.StringValue(resp.Datalake.CertificateExpirationState)
	model.CreationDate = types.StringValue(resp.Datalake.CreationDate.String())
	model.Crn = types.StringPointerValue(resp.Datalake.Crn)
	model.DatalakeName = types.StringPointerValue(resp.Datalake.DatalakeName)
	model.EnableRangerRaz = types.BoolValue(resp.Datalake.EnableRangerRaz)
	model.EnvironmentCrn = types.StringValue(resp.Datalake.EnvironmentCrn)
	model.MultiAz = types.BoolValue(resp.Datalake.MultiAz)
	model.Status = types.StringValue(resp.Datalake.Status)
	model.StatusReason = types.StringValue(resp.Datalake.StatusReason)
}

func (r *awsDatalakeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state awsDatalakeResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client.Datalake

	params := operations.NewDescribeDatalakeParamsWithContext(ctx)
	params.WithInput(&datalakemodels.DescribeDatalakeRequest{DatalakeName: state.DatalakeName.ValueStringPointer()})
	responseOk, err := client.Operations.DescribeDatalake(params)
	if err != nil {
		if dlErr, ok := err.(*operations.DescribeDatalakeDefault); ok {
			if cdp.IsDatalakeError(dlErr.GetPayload(), "NOT_FOUND", "") {
				resp.Diagnostics.AddWarning("Resource not found on provider", "Data lake not found, removing from state.")
				tflog.Warn(ctx, "Data lake not found, removing from state", map[string]interface{}{
					"id": state.ID.ValueString(),
				})
				resp.State.RemoveResource(ctx)
				return
			}
		}
		utils.AddDatalakeDiagnosticsError(err, &resp.Diagnostics, "read AWS Datalake")
		return
	}

	datalakeResp := responseOk.Payload
	datalakeDetailsToAwsDatalakeResourceModel(ctx, datalakeResp.Datalake, &state, state.PollingOptions, &resp.Diagnostics)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func datalakeDetailsToAwsDatalakeResourceModel(ctx context.Context, resp *datalakemodels.DatalakeDetails, model *awsDatalakeResourceModel, pollingOptions *utils.PollingOptions, diags *diag.Diagnostics) {
	model.ID = types.StringPointerValue(resp.Crn)
	if resp.AwsConfiguration != nil {
		model.InstanceProfile = types.StringValue(resp.AwsConfiguration.InstanceProfile)
	}
	model.CreationDate = types.StringValue(resp.CreationDate.String())
	model.Crn = types.StringPointerValue(resp.Crn)
	model.DatalakeName = types.StringPointerValue(resp.DatalakeName)
	model.EnableRangerRaz = types.BoolValue(resp.EnableRangerRaz)
	model.PollingOptions = pollingOptions
	model.EnvironmentCrn = types.StringValue(resp.EnvironmentCrn)
	instanceGroups := make([]*instanceGroup, len(resp.InstanceGroups))
	for i, v := range resp.InstanceGroups {
		instanceGroups[i] = &instanceGroup{
			Name: types.StringPointerValue(v.Name),
		}

		instances := make([]*instance, 0, len(v.Instances))
		for _, ins := range v.Instances {
			if ins == nil || ins.ID == nil || len(*ins.ID) == 0 {
				continue
			}
			instances = append(instances, &instance{
				DiscoveryFQDN:   types.StringValue(ins.DiscoveryFQDN),
				ID:              types.StringPointerValue(ins.ID),
				InstanceGroup:   types.StringValue(ins.InstanceGroup),
				InstanceStatus:  types.StringValue(string(ins.InstanceStatus)),
				InstanceTypeVal: types.StringValue(string(ins.InstanceTypeVal)),
				PrivateIP:       types.StringValue(ins.PrivateIP),
				PublicIP:        types.StringValue(ins.PublicIP),
				SSHPort:         types.Int64Value(int64(ins.SSHPort)),
				State:           types.StringPointerValue(ins.State),
				StatusReason:    types.StringValue(ins.StatusReason),
			})
		}
		var instDiags diag.Diagnostics
		instanceGroups[i].Instances, instDiags = types.SetValueFrom(ctx, types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"discovery_fqdn":    types.StringType,
				"id":                types.StringType,
				"instance_group":    types.StringType,
				"instance_status":   types.StringType,
				"instance_type_val": types.StringType,
				"private_ip":        types.StringType,
				"public_ip":         types.StringType,
				"ssh_port":          types.Int64Type,
				"state":             types.StringType,
				"status_reason":     types.StringType,
			},
		}, instances)
		diags.Append(instDiags...)
	}
	model.Scale = types.StringValue(string(resp.Shape))
	model.Status = types.StringValue(resp.Status)
	model.StatusReason = types.StringValue(resp.StatusReason)
	if model.CertificateExpirationState.IsUnknown() {
		model.CertificateExpirationState = types.StringNull()
	}
}

func (r *awsDatalakeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *awsDatalakeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state awsDatalakeResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	dlName := state.DatalakeName.ValueString()
	if len(dlName) == 0 {
		dlName = state.ID.ValueString()
	}
	client := r.client.Datalake
	params := operations.NewDeleteDatalakeParamsWithContext(ctx)
	params.WithInput(&datalakemodels.DeleteDatalakeRequest{
		DatalakeName: &dlName,
		Force:        false,
	})
	_, err := client.Operations.DeleteDatalake(params)
	if err != nil {
		if dlErr, ok := err.(*operations.DescribeDatalakeDefault); ok {
			if cdp.IsDatalakeError(dlErr.GetPayload(), "NOT_FOUND", "") {
				tflog.Info(ctx, "Data lake already deleted", map[string]interface{}{
					"id": state.ID.ValueString(),
				})
				return
			}
		}
		utils.AddDatalakeDiagnosticsError(err, &resp.Diagnostics, "delete AWS Datalake")
		return
	}

	if err := waitForDatalakeToBeDeleted(ctx, state.DatalakeName.ValueString(), time.Hour, r.client.Datalake, state.PollingOptions); err != nil {
		utils.AddDatalakeDiagnosticsError(err, &resp.Diagnostics, "delete AWS Datalake")
		return
	}

}

func waitForDatalakeToBeDeleted(ctx context.Context, datalakeName string, fallbackPollingTimeout time.Duration, datalake *client.Datalake, options *utils.PollingOptions) error {
	timeout, err := utils.CalculateTimeoutOrDefault(ctx, options, fallbackPollingTimeout)
	if err != nil {
		return err
	}
	failureThreshold, failureThresholdErr := utils.CalculateCallFailureThresholdOrDefault(ctx, options, callFailureThreshold)
	if failureThresholdErr != nil {
		return failureThresholdErr
	}
	callFailedCount := 0
	stateConf := &retry.StateChangeConf{
		Pending: []string{"DELETE_REQUESTED", "STACK_DELETION_IN_PROGRESS", "STACK_DELETED", "EXTERNAL_DATABASE_DELETION_IN_PROGRESS", "DELETED"},
		Target:  []string{},
		Timeout: *timeout,
		Refresh: func() (interface{}, string, error) {
			params := operations.NewDescribeDatalakeParamsWithContext(ctx)
			params.WithInput(&datalakemodels.DescribeDatalakeRequest{DatalakeName: &datalakeName})
			resp, err := datalake.Operations.DescribeDatalake(params)
			if err != nil {
				if dlErr, ok := err.(*operations.DescribeDatalakeDefault); ok {
					if cdp.IsDatalakeError(dlErr.GetPayload(), "NOT_FOUND", "") {
						return nil, "", nil
					}
				}
				callFailedCount++
				if callFailedCount <= failureThreshold {
					tflog.Warn(ctx, fmt.Sprintf("Error describing datalake with call failure due to [%s] but threshold limit is not reached yet (%d out of %d).", err.Error(), callFailedCount, callFailureThreshold))
					return nil, "", nil
				}
				return nil, "", err
			}
			if resp.GetPayload().Datalake == nil {
				return nil, "", nil
			}
			return checkResponseStatusForError(resp)
		},
	}
	_, err = stateConf.WaitForStateContext(ctx)

	return err
}
