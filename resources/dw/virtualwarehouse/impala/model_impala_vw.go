// Copyright 2024 Cloudera. All Rights Reserved.
//
// This file is licensed under the Apache License Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0.
//
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS
// OF ANY KIND, either express or implied. Refer to the License for the specific
// permissions and limitations governing your use of the file.

package impala

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const timeZone = time.RFC850

type TagRequest struct {
	Key   string `tfsdk:"key"`   // The tag's name
	Value string `tfsdk:"value"` // The associated value of the tag
}

// TODO Prateek Use or Remove once you are done fixing Impala API
type ServiceConfigReqModel struct {
	ApplicationConfigs map[string]ApplicationConfigReqModel `tfsdk:"application_configs"`
	CommonConfigs      *ApplicationConfigReqModel           `tfsdk:"common_configs"`
	EnableSSO          basetypes.BoolValue                  `tfsdk:"enable_sso"`
	LdapGroups         []string                             `tfsdk:"ldap_groups"`
}

// TODO Prateek Use or Remove once you are done fixing Impala API
type ApplicationConfigReqModel struct {
	ConfigBlocks []*ConfigBlockReqModel `tfsdk:"config_blocks"`
}

// TODO Prateek Use or Remove once you are done fixing Impala API
type ConfigBlockReqModel struct {
	Content *ConfigContentReqModel `tfsdk:"content"`
	ID      *string                `tfsdk:"id"`
}

// TODO Prateek Use or Remove once you are done fixing Impala API
type ConfigContentReqModel struct {
	JSON      string            `tfsdk:"json"`
	KeyValues map[string]string `tfsdk:"key_values"`
	Text      string            `tfsdk:"text"`
}

// TODO Prateek Use or Remove once you are done fixing Impala API
type ImpalaExecutorGroupSetsModel struct {
	Custom1 *ImpalaExecutorGroupSetModel `tfsdk:"custom1"`
	Custom2 *ImpalaExecutorGroupSetModel `tfsdk:"custom2"`
	Custom3 *ImpalaExecutorGroupSetModel `tfsdk:"custom3"`
	Large   *ImpalaExecutorGroupSetModel `tfsdk:"large"`
	Small   *ImpalaExecutorGroupSetModel `tfsdk:"small"`
}

// TODO Prateek Use or Remove once you are done fixing Impala API
type ImpalaExecutorGroupSetModel struct {
	AutoSuspendTimeoutSeconds int32 `tfsdk:"auto_suspend_timeout_seconds"`
	DisableAutoSuspend        bool  `tfsdk:"disable_auto_suspend"`
	ExecGroupSize             int32 `tfsdk:"exec_group_size"`
	MaxExecutorGroups         int32 `tfsdk:"max_executor_groups"`
	MinExecutorGroups         int32 `tfsdk:"min_executor_groups"`
	TriggerScaleDownDelay     int32 `tfsdk:"trigger_scale_down_delay"`
	TriggerScaleUpDelay       int32 `tfsdk:"trigger_scale_up_delay"`
}

type resourceModel struct {
	ID                     types.String `tfsdk:"id"`
	ClusterID              types.String `tfsdk:"cluster_id"`
	DatabaseCatalogID      types.String `tfsdk:"database_catalog_id"`
	Name                   types.String `tfsdk:"name"`
	LastUpdated            types.String `tfsdk:"last_updated"`
	Status                 types.String `tfsdk:"status"`
	ImageVersion           types.String `tfsdk:"image_version"`
	InstanceType           types.String `tfsdk:"instance_type"`
	TShirtSize             types.String `tfsdk:"tshirt_size"`
	NodeCount              types.Int32  `tfsdk:"node_count"`
	AvailabilityZone       types.String `tfsdk:"availability_zone"`
	EnableUnifiedAnalytics types.Bool   `tfsdk:"enable_unified_analytics"`
	ImpalaOptions          types.Object `tfsdk:"impala_options"`
	ImpalaHASettings       types.Object `tfsdk:"impala_ha_settings"`
	Autoscaling            types.Object `tfsdk:"autoscaling"`
	//Config                 types.Object `tfsdk:"config"`
	QueryIsolationOptions  types.Object `tfsdk:"query_isolation_options"`
	Tags                   types.List   `tfsdk:"tags"`
	ResourcePool           types.String `tfsdk:"resource_pool"`
	HiveAuthenticationMode types.String `tfsdk:"hive_authentication_mode"`
	PlatformJwtAuth        types.Bool   `tfsdk:"platform_jwt_auth"`
	ImpalaQueryLog         types.Bool   `tfsdk:"impala_query_log"`
	EbsLLAPSpillGB         types.Int64  `tfsdk:"ebs_llap_spill_gb"`
	// TODO Prateek This does not look like should be in Impala, so delete
	// or enable after talking to Impala team
	// HiveServerHaMode       types.String                `tfsdk:"hive_server_ha_mode"`

	PollingOptions *utils.PollingOptions `tfsdk:"polling_options"`
}

func (p *resourceModel) setFromDescribeVwResponse(resp *models.DescribeVwResponse) {
	p.ID = types.StringValue(resp.Vw.ID)
	p.DatabaseCatalogID = types.StringValue(resp.Vw.DbcID)
	p.Name = types.StringValue(resp.Vw.Name)
	p.Status = types.StringValue(resp.Vw.Status)
	p.ImageVersion = types.StringValue(resp.Vw.CdhVersion)
	p.LastUpdated = types.StringValue(time.Now().Format(timeZone))

	setStringIfNotEmpty := func(target *types.String, source string) {
		if source != "" {
			*target = types.StringValue(source)
		}
	}

	// Helper function for setting optional int values
	setInt32IfPositive := func(target *types.Int32, source int32) {
		if source > 0 {
			*target = types.Int32Value(source)
		}
	}

	// Helper function for setting optional int64 values
	setInt64IfPositive := func(target *types.Int64, source int32) {
		if source > 0 {
			*target = types.Int64Value(int64(source))
		}
	}

	// Helper function for setting optional pointer-based string values
	setStringIfNotNil := func(target *types.String, source *string) {
		if source != nil {
			*target = types.StringValue(*source)
		}
	}

	/*if resp.Vw.HiveServerHaMode != nil {
		p.InstanceType = types.StringValue(resp.Vw.InstanceType)
	}*/

	setInt32IfPositive(&p.NodeCount, resp.Vw.NodeCount)
	setStringIfNotEmpty(&p.InstanceType, resp.Vw.InstanceType)
	setStringIfNotEmpty(&p.AvailabilityZone, resp.Vw.AvailabilityZone)
	setStringIfNotEmpty(&p.ResourcePool, resp.Vw.ResourcePool)
	setStringIfNotNil(&p.HiveAuthenticationMode, resp.Vw.HiveAuthenticationMode)
	setInt64IfPositive(&p.EbsLLAPSpillGB, resp.Vw.EbsLLAPSpillGB)

	p.EnableUnifiedAnalytics = types.BoolValue(resp.Vw.EnableUnifiedAnalytics)
	p.ImpalaQueryLog = types.BoolValue(resp.Vw.ImpalaQueryLog)

	if resp.Vw.ImpalaOptions != nil {
		p.ImpalaOptions = convertFromAPIImpalaOptions(resp.Vw.ImpalaOptions)
	}
	if resp.Vw.ImpalaHaSettingsOptions != nil {
		p.ImpalaHASettings = convertFromAPIImpalaHASettings(resp.Vw.ImpalaHaSettingsOptions)
	}
	if resp.Vw.AutoscalingOptions != nil {
		p.Autoscaling = convertFromAPIAutoscaling(resp.Vw.AutoscalingOptions)
	}
	if resp.Vw.QueryIsolationOptions != nil {
		p.QueryIsolationOptions = convertFromAPIQueryIsolationOptions(resp.Vw.QueryIsolationOptions)
	}

	if len(resp.Vw.Tags) != 0 {
		p.Tags = convertFromAPITagRequests(resp.Vw.Tags)
	}

	/*if resp.Vw.HiveServerHaMode != nil {
		p.HiveServerHaMode = types.StringValue(*(resp.Vw.HiveServerHaMode))
	}*/

}

func (p *resourceModel) GetPollingOptions() *utils.PollingOptions {
	return p.PollingOptions
}

func convertToAPIImpalaHASettings(model types.Object) *models.ImpalaHASettingsCreateRequest {
	if model.IsUnknown() {
		return nil
	}
	attributes := model.Attributes()

	req := &models.ImpalaHASettingsCreateRequest{}

	if val, ok := attributes["enable_catalog_high_availability"]; ok {
		if value, err := ExtractBoolFromAttribute(context.Background(), val.(basetypes.BoolValue)); err == nil {
			req.EnableCatalogHighAvailability = value
		}
	}

	if val, ok := attributes["enable_shutdown_of_coordinator"]; ok {
		if value, err := ExtractBoolFromAttribute(context.Background(), val.(basetypes.BoolValue)); err == nil {
			req.EnableShutdownOfCoordinator = value
		}
	}

	if val, ok := attributes["enable_statestore_high_availability"]; ok {
		if value, err := ExtractBoolFromAttribute(context.Background(), val.(basetypes.BoolValue)); err == nil {
			req.EnableStatestoreHighAvailability = value
		}
	}

	if val, ok := attributes["high_availability_mode"]; ok {
		if value, err := ExtractStringFromAttribute(context.Background(), val.(basetypes.StringValue)); err == nil {
			req.HighAvailabilityMode = models.ImpalaHighAvailabilityMode(value)
		}
	}

	if val, ok := attributes["num_of_active_coordinators"]; ok {
		if value, err := ExtractInt32FromAttribute(context.Background(), val.(basetypes.Int32Value)); err == nil {
			req.NumOfActiveCoordinators = value
		}
	}

	if val, ok := attributes["shutdown_of_coordinator_delay_secs"]; ok {
		if value, err := ExtractInt32FromAttribute(context.Background(), val.(basetypes.Int32Value)); err == nil {
			req.ShutdownOfCoordinatorDelaySeconds = value
		}
	}

	return req
}

func convertFromAPIImpalaHASettings(apiModel *models.ImpalaHASettingsOptionsResponse) types.Object {
	ctx := context.Background()

	if apiModel == nil {
		tflog.Debug(ctx, "apiModel is nil, returning ObjectNull")
		return types.ObjectNull(map[string]attr.Type{
			"enable_catalog_high_availability":    types.BoolType,
			"enable_shutdown_of_coordinator":      types.BoolType,
			"enable_statestore_high_availability": types.BoolType,
			"high_availability_mode":              types.StringType,
			"num_of_active_coordinators":          types.Int32Type,
			"shutdown_of_coordinator_delay_secs":  types.Int32Type,
		})
	}

	attributeTypes := map[string]attr.Type{
		"enable_catalog_high_availability":    types.BoolType,
		"enable_shutdown_of_coordinator":      types.BoolType,
		"enable_statestore_high_availability": types.BoolType,
		"high_availability_mode":              types.StringType,
		"num_of_active_coordinators":          types.Int32Type,
		"shutdown_of_coordinator_delay_secs":  types.Int32Type,
	}

	attributeValues := map[string]attr.Value{
		"enable_catalog_high_availability":    types.BoolNull(),
		"enable_shutdown_of_coordinator":      types.BoolNull(),
		"enable_statestore_high_availability": types.BoolNull(),
		"high_availability_mode":              types.StringNull(),
		"num_of_active_coordinators":          types.Int32Null(),
		"shutdown_of_coordinator_delay_secs":  types.Int32Null(),
	}

	attributeValues["enable_catalog_high_availability"] = types.BoolValue(apiModel.EnableCatalogHighAvailability)
	attributeValues["enable_shutdown_of_coordinator"] = types.BoolValue(apiModel.EnableShutdownOfCoordinator)
	attributeValues["enable_statestore_high_availability"] = types.BoolValue(apiModel.EnableStatestoreHighAvailability)
	if apiModel.HighAvailabilityMode != "" {
		attributeValues["high_availability_mode"] = types.StringValue(string(apiModel.HighAvailabilityMode))
	}
	if apiModel.NumOfActiveCoordinators != 0 {
		attributeValues["num_of_active_coordinators"] = types.Int32Value(apiModel.NumOfActiveCoordinators)
	}
	if apiModel.ShutdownOfCoordinatorDelaySeconds != 0 {
		attributeValues["shutdown_of_coordinator_delay_secs"] = types.Int32Value(apiModel.ShutdownOfCoordinatorDelaySeconds)
	}
	ret, err := types.ObjectValue(attributeTypes, attributeValues)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Error creating ObjectValue: %v", err))
	}
	return ret
}

func convertToAPIImpalaOptions(model types.Object) *models.ImpalaOptionsCreateRequest {
	if model.IsUnknown() {
		return nil
	}
	attributes := model.Attributes()

	scratchSpaceLimit, hasScratchSpace := attributes["scratch_space_limit"]
	spillToS3URI, hasSpillToS3URI := attributes["spill_to_s3_uri"]

	req := &models.ImpalaOptionsCreateRequest{}

	if hasScratchSpace && !scratchSpaceLimit.IsUnknown() {
		intValue, err := ExtractInt32FromAttribute(context.Background(), scratchSpaceLimit)
		if err != nil {
			fmt.Printf("Error extracting int32 for ScratchSpaceLimit: %v\n", err)
		} else {
			req.ScratchSpaceLimit = intValue
			fmt.Printf("Assigned value to ScratchSpaceLimit: %d\n", req.ScratchSpaceLimit)
		}
		// If this is set, don't set spillToS3URI
		return req
	}

	if hasSpillToS3URI && !spillToS3URI.IsUnknown() {
		valueString, err := ExtractStringFromAttribute(context.Background(), spillToS3URI)
		if err != nil {
			fmt.Printf("Error extracting string for SpillToS3URI: %v\n", err)
		} else {
			req.SpillToS3URI = valueString
			fmt.Printf("Assigned value to SpillToS3URI: %s\n", req.SpillToS3URI)
		}
	}

	return req
}

func convertFromAPIImpalaOptions(apiModel *models.ImpalaOptionsResponse) types.Object {

	if apiModel == nil {
		return types.ObjectNull(map[string]attr.Type{
			"scratch_space_limit": types.Int32Type,
			"spill_to_s3_uri":     types.StringType,
		})
	}

	attributeTypes := map[string]attr.Type{
		"scratch_space_limit": types.Int32Type,
		"spill_to_s3_uri":     types.StringType,
	}

	attributeValues := map[string]attr.Value{
		"scratch_space_limit": types.Int32Null(),
		"spill_to_s3_uri":     types.StringNull(),
	}

	if apiModel.ScratchSpaceLimit != 0 {
		attributeValues["scratch_space_limit"] = types.Int32Value(apiModel.ScratchSpaceLimit)
	} else {
		attributeValues["scratch_space_limit"] = types.Int32Null() // Null if missing
	}

	if apiModel.SpillToS3URI != "" {
		attributeValues["spill_to_s3_uri"] = types.StringValue(apiModel.SpillToS3URI)
	} else {
		attributeValues["spill_to_s3_uri"] = types.StringNull() // Null if empty
	}

	// Return as types.Object
	ret, _ := types.ObjectValue(attributeTypes, attributeValues)
	return ret
}

func convertToAPIQueryIsolationOptions(model types.Object) *models.QueryIsolationOptionsRequest {
	if model.IsUnknown() {
		return nil
	}
	attributes := model.Attributes()

	req := &models.QueryIsolationOptionsRequest{}

	if val, ok := attributes["max_nodes_per_query"]; ok {
		if value, err := ExtractInt32FromAttribute(context.Background(), val.(basetypes.Int32Value)); err == nil {
			req.MaxNodesPerQuery = value
		}
	}

	if val, ok := attributes["max_queries"]; ok {
		if value, err := ExtractInt32FromAttribute(context.Background(), val.(basetypes.Int32Value)); err == nil {
			req.MaxQueries = value
		}
	}

	return req
}

func convertFromAPIQueryIsolationOptions(apiModel *models.QueryIsolationOptionsResponse) types.Object {
	if apiModel == nil {
		return types.ObjectNull(map[string]attr.Type{
			"max_nodes_per_query": types.Int32Type,
			"max_queries":         types.Int32Type,
		})
	}

	attributeTypes := map[string]attr.Type{
		"max_nodes_per_query": types.Int32Type,
		"max_queries":         types.Int32Type,
	}

	attributeValues := map[string]attr.Value{
		"max_nodes_per_query": types.Int32Null(),
		"max_queries":         types.Int32Null(),
	}

	// Handle MaxNodesPerQuery
	if apiModel.MaxNodesPerQuery != 0 {
		attributeValues["max_nodes_per_query"] = types.Int32Value(apiModel.MaxNodesPerQuery)
	} else {
		attributeValues["max_nodes_per_query"] = types.Int32Null()
	}

	// Handle MaxQueries
	if apiModel.MaxQueries != 0 {
		attributeValues["max_queries"] = types.Int32Value(apiModel.MaxQueries)
	} else {
		attributeValues["max_queries"] = types.Int32Null()
	}

	// Return as types.Object
	ret, _ := types.ObjectValue(attributeTypes, attributeValues)
	return ret
}

func ConvertToServiceConfigReqModel(model *ServiceConfigReqModel) *models.ServiceConfigReq {
	converted := &models.ServiceConfigReq{
		ApplicationConfigs: make(map[string]models.ApplicationConfigReq),
		EnableSSO:          model.EnableSSO.ValueBool(),
		LdapGroups:         model.LdapGroups,
	}

	for key, value := range model.ApplicationConfigs {
		converted.ApplicationConfigs[key] = ConvertToApplicationConfigReqModel(value)
	}

	if model.CommonConfigs != nil {
		commonConfig := ConvertToApplicationConfigReqModel(*model.CommonConfigs)
		converted.CommonConfigs = &commonConfig
	}

	return converted
}

func ConvertToApplicationConfigReqModel(req ApplicationConfigReqModel) models.ApplicationConfigReq {
	converted := models.ApplicationConfigReq{
		ConfigBlocks: make([]*models.ConfigBlockReq, len(req.ConfigBlocks)),
	}
	for i, block := range req.ConfigBlocks {
		converted.ConfigBlocks[i] = ConvertToConfigBlockReqModel(*block)
	}
	return converted
}

func ConvertToConfigBlockReqModel(req ConfigBlockReqModel) *models.ConfigBlockReq {
	return &models.ConfigBlockReq{
		Content: ConvertToConfigContentReqModel(*req.Content),
		ID:      req.ID,
	}
}

func ConvertToConfigContentReqModel(req ConfigContentReqModel) *models.ConfigContentReq {
	return &models.ConfigContentReq{
		JSON:      req.JSON,
		KeyValues: req.KeyValues,
		Text:      req.Text,
	}
}

func convertToAPIAutoscaling(model types.Object) *models.AutoscalingOptionsCreateRequest {
	if model.IsUnknown() {
		return nil
	}
	attributes := model.Attributes()

	// TODO Remove debugging
	for key, val := range attributes {
		tflog.Debug(context.Background(), fmt.Sprintf("Key: %s, IsUnknown: %t, Value: %+v", key, val.IsUnknown(), val))
		fmt.Printf("Key: %s, IsUnknown: %t, Value: %+v, type %v,\n", key, val.IsUnknown(), val, val.Type(context.Background()))
	}

	req := &models.AutoscalingOptionsCreateRequest{}

	if val, ok := attributes["auto_suspend_timeout_seconds"]; ok {
		// Debug statement to be deleted
		fmt.Printf("Found key 'auto_suspend_timeout_seconds' in attributes: %+v\n", val)
		intValue, err := ExtractInt32FromAttribute(context.Background(), val)
		if err != nil {
			fmt.Printf("Error extracting int32: %v\n", err)
		} else {
			req.AutoSuspendTimeoutSeconds = intValue
			fmt.Printf("Assigned value to AutoSuspendTimeoutSeconds: %d\n", req.AutoSuspendTimeoutSeconds)
		}
	}

	if val, ok := attributes["disable_auto_suspend"]; ok {
		boolValue, err := ExtractBoolFromAttribute(context.Background(), val)
		if err != nil {
			fmt.Printf("Error extracting bool: %v\n", err)
		} else {
			req.DisableAutoSuspend = boolValue
			fmt.Printf("Assigned value to DisableAutoSuspend: %t\n", req.DisableAutoSuspend)
		}
	}

	/*if val, ok := attributes["impala_num_of_active_coordinators"]; ok && !val.IsUnknown() {
		tfv, err := val.ToTerraformValue(context.Background())
		if err == nil {
			var valueInt32 int32
			if err := tfv.As(&valueInt32); err == nil {
				req.ImpalaNumOfActiveCoordinators = valueInt32
			}
		}
	}*/

	if val, ok := attributes["impala_scale_down_delay_seconds"]; ok {
		intValue, err := ExtractInt32FromAttribute(context.Background(), val)
		if err != nil {
			fmt.Printf("Error extracting int32: %v\n", err)
		} else {
			req.ImpalaScaleDownDelaySeconds = intValue
			fmt.Printf("Assigned value to ImpalaScaleDownDelaySeconds: %d\n", req.ImpalaScaleDownDelaySeconds)
		}
	}

	if val, ok := attributes["impala_scale_up_delay_seconds"]; ok {
		intValue, err := ExtractInt32FromAttribute(context.Background(), val)
		if err != nil {
			fmt.Printf("Error extracting int32 for impala_scale_up_delay_seconds: %v\n", err)
		} else {
			req.ImpalaScaleUpDelaySeconds = intValue
			fmt.Printf("Assigned value to ImpalaScaleUpDelaySeconds: %d\n", req.ImpalaScaleUpDelaySeconds)
		}
	}

	/*if val, ok := attributes["impala_shutdown_of_coordinator_delay_seconds"]; ok && !val.IsUnknown() {
		tfv, err := val.ToTerraformValue(context.Background())
		if err == nil {
			var valueInt32 int32
			if err := tfv.As(&valueInt32); err == nil {
				req.ImpalaShutdownOfCoordinatorDelaySeconds = valueInt32
			}
		}
	}*/

	if val, ok := attributes["max_clusters"]; ok {
		intValue, err := ExtractInt32FromAttribute(context.Background(), val)
		if err != nil {
			fmt.Printf("Error extracting int32 for max_clusters: %v\n", err)
		} else {
			req.MaxClusters = &intValue
			fmt.Printf("Assigned value to MaxClusters: %d\n", *req.MaxClusters)
		}
	}

	if val, ok := attributes["min_clusters"]; ok {
		intValue, err := ExtractInt32FromAttribute(context.Background(), val)
		if err != nil {
			fmt.Printf("Error extracting int32 for min_clusters: %v\n", err)
		} else {
			req.MinClusters = &intValue
			fmt.Printf("Assigned value to MinClusters: %d\n", *req.MinClusters)
		}
	}

	/*if val, ok := attributes["impala_executor_group_sets"]; ok && !val.IsUnknown() {
		tfv, err := val.ToTerraformValue(context.Background())
		if err == nil {
			var impalaExecutorGroupSets ImpalaExecutorGroupSetsModel
			if err := tfv.As(&impalaExecutorGroupSets); err == nil {
				req.ImpalaExecutorGroupSets = convertToAPIExecutorGroupSets(&impalaExecutorGroupSets)
			}
		}
	}*/

	return req
}

func convertToAPIExecutorGroupSets(model *ImpalaExecutorGroupSetsModel) *models.ImpalaExecutorGroupSetsCreateRequest {
	if model == nil {
		return nil
	}

	return &models.ImpalaExecutorGroupSetsCreateRequest{
		Custom1: convertToAPIExecutorGroupSet(model.Custom1),
		Custom2: convertToAPIExecutorGroupSet(model.Custom2),
		Custom3: convertToAPIExecutorGroupSet(model.Custom3),
		Large:   convertToAPIExecutorGroupSet(model.Large),
		Small:   convertToAPIExecutorGroupSet(model.Small),
	}
}

func convertToAPIExecutorGroupSet(model *ImpalaExecutorGroupSetModel) *models.ImpalaExecutorGroupSetCreateRequest {
	if model == nil {
		return nil
	}

	return &models.ImpalaExecutorGroupSetCreateRequest{
		AutoSuspendTimeoutSeconds: model.AutoSuspendTimeoutSeconds,
		DisableAutoSuspend:        model.DisableAutoSuspend,
		ExecGroupSize:             model.ExecGroupSize,
		MaxExecutorGroups:         model.MaxExecutorGroups,
		MinExecutorGroups:         model.MinExecutorGroups,
		TriggerScaleDownDelay:     model.TriggerScaleDownDelay,
		TriggerScaleUpDelay:       model.TriggerScaleUpDelay,
	}
}

func convertFromAPIAutoscaling(apiModel *models.AutoscalingOptionsResponse) types.Object {
	if apiModel == nil {
		return types.ObjectNull(map[string]attr.Type{
			"auto_suspend_timeout_seconds": types.Int32Type,
			"disable_auto_suspend":         types.BoolType,
			//"impala_num_of_active_coordinators":            types.Int32Type,
			"impala_scale_down_delay_seconds": types.Int32Type,
			"impala_scale_up_delay_seconds":   types.Int32Type,
			//"impala_shutdown_of_coordinator_delay_seconds": types.Int32Type,
			"max_clusters": types.Int32Type,
			"min_clusters": types.Int32Type,
			/*"impala_executor_group_sets": types.ListType{
				ElemType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"custom1": types.ObjectType{},
						"custom2": types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"auto_suspend_timeout_seconds": types.Int32Type,
								"disable_auto_suspend":         types.BoolType,
								"exec_group_size":              types.Int32Type,
								"max_executor_groups":          types.Int32Type,
								"min_executor_groups":          types.Int32Type,
								"trigger_scale_down_delay":     types.Int32Type,
								"trigger_scale_up_delay":       types.Int32Type,
							},
						},
						"custom3": types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"auto_suspend_timeout_seconds": types.Int32Type,
								"disable_auto_suspend":         types.BoolType,
								"exec_group_size":              types.Int32Type,
								"max_executor_groups":          types.Int32Type,
								"min_executor_groups":          types.Int32Type,
								"trigger_scale_down_delay":     types.Int32Type,
								"trigger_scale_up_delay":       types.Int32Type,
							},
						},
						"large": types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"auto_suspend_timeout_seconds": types.Int32Type,
								"disable_auto_suspend":         types.BoolType,
								"exec_group_size":              types.Int32Type,
								"max_executor_groups":          types.Int32Type,
								"min_executor_groups":          types.Int32Type,
								"trigger_scale_down_delay":     types.Int32Type,
								"trigger_scale_up_delay":       types.Int32Type,
							},
						},
						"small": types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"auto_suspend_timeout_seconds": types.Int32Type,
								"disable_auto_suspend":         types.BoolType,
								"exec_group_size":              types.Int32Type,
								"max_executor_groups":          types.Int32Type,
								"min_executor_groups":          types.Int32Type,
								"trigger_scale_down_delay":     types.Int32Type,
								"trigger_scale_up_delay":       types.Int32Type,
							},
						},
					},
				},
			},*/
		})
	}

	attributeTypes := map[string]attr.Type{
		"auto_suspend_timeout_seconds": types.Int32Type,
		"disable_auto_suspend":         types.BoolType,
		//"impala_num_of_active_coordinators":            types.Int32Type,
		"impala_scale_down_delay_seconds": types.Int32Type,
		"impala_scale_up_delay_seconds":   types.Int32Type,
		//"impala_shutdown_of_coordinator_delay_seconds": types.Int32Type,
		"max_clusters": types.Int32Type,
		"min_clusters": types.Int32Type,
		/*"impala_executor_group_sets": types.ListType{
			ElemType: types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"custom1": types.ObjectType{},
					"custom2": types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"auto_suspend_timeout_seconds": types.Int32Type,
							"disable_auto_suspend":         types.BoolType,
							"exec_group_size":              types.Int32Type,
							"max_executor_groups":          types.Int32Type,
							"min_executor_groups":          types.Int32Type,
							"trigger_scale_down_delay":     types.Int32Type,
							"trigger_scale_up_delay":       types.Int32Type,
						},
					},
					"custom3": types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"auto_suspend_timeout_seconds": types.Int32Type,
							"disable_auto_suspend":         types.BoolType,
							"exec_group_size":              types.Int32Type,
							"max_executor_groups":          types.Int32Type,
							"min_executor_groups":          types.Int32Type,
							"trigger_scale_down_delay":     types.Int32Type,
							"trigger_scale_up_delay":       types.Int32Type,
						},
					},
					"large": types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"auto_suspend_timeout_seconds": types.Int32Type,
							"disable_auto_suspend":         types.BoolType,
							"exec_group_size":              types.Int32Type,
							"max_executor_groups":          types.Int32Type,
							"min_executor_groups":          types.Int32Type,
							"trigger_scale_down_delay":     types.Int32Type,
							"trigger_scale_up_delay":       types.Int32Type,
						},
					},
					"small": types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"auto_suspend_timeout_seconds": types.Int32Type,
							"disable_auto_suspend":         types.BoolType,
							"exec_group_size":              types.Int32Type,
							"max_executor_groups":          types.Int32Type,
							"min_executor_groups":          types.Int32Type,
							"trigger_scale_down_delay":     types.Int32Type,
							"trigger_scale_up_delay":       types.Int32Type,
						},
					},
				},
			},
		},*/
	}

	attributeValues := map[string]attr.Value{
		"auto_suspend_timeout_seconds": types.Int32Null(),
		"disable_auto_suspend":         types.BoolNull(),
		//"impala_num_of_active_coordinators":            types.Int32Null(),
		"impala_scale_down_delay_seconds": types.Int32Null(),
		"impala_scale_up_delay_seconds":   types.Int32Null(),
		//"impala_shutdown_of_coordinator_delay_seconds": types.Int32Null(),
		"max_clusters": types.Int32Null(),
		"min_clusters": types.Int32Null(),
		/*"impala_executor_group_sets": types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"small": types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"exec_group_size":              types.Int64Type,
						"min_executor_groups":          types.Int64Type,
						"max_executor_groups":          types.Int64Type,
						"auto_suspend_timeout_seconds": types.NumberType,
						"disable_auto_suspend":         types.BoolType,
						"trigger_scale_up_delay":       types.Int64Type,
						"trigger_scale_down_delay":     types.Int64Type,
					},
				},
				"custom1": types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"exec_group_size":              types.Int64Type,
						"min_executor_groups":          types.Int64Type,
						"max_executor_groups":          types.Int64Type,
						"auto_suspend_timeout_seconds": types.NumberType,
						"disable_auto_suspend":         types.BoolType,
						"trigger_scale_up_delay":       types.Int64Type,
						"trigger_scale_down_delay":     types.Int64Type,
					},
				},
				"custom2": types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"exec_group_size":              types.Int64Type,
						"min_executor_groups":          types.Int64Type,
						"max_executor_groups":          types.Int64Type,
						"auto_suspend_timeout_seconds": types.NumberType,
						"disable_auto_suspend":         types.BoolType,
						"trigger_scale_up_delay":       types.Int64Type,
						"trigger_scale_down_delay":     types.Int64Type,
					},
				},
				"custom3": types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"exec_group_size":              types.Int64Type,
						"min_executor_groups":          types.Int64Type,
						"max_executor_groups":          types.Int64Type,
						"auto_suspend_timeout_seconds": types.NumberType,
						"disable_auto_suspend":         types.BoolType,
						"trigger_scale_up_delay":       types.Int64Type,
						"trigger_scale_down_delay":     types.Int64Type,
					},
				},
				"large": types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"exec_group_size":              types.Int64Type,
						"min_executor_groups":          types.Int64Type,
						"max_executor_groups":          types.Int64Type,
						"auto_suspend_timeout_seconds": types.NumberType,
						"disable_auto_suspend":         types.BoolType,
						"trigger_scale_up_delay":       types.Int64Type,
						"trigger_scale_down_delay":     types.Int64Type,
					},
				},
			},
		}),*/
	}

	// Populate non-null values
	if apiModel.AutoSuspendTimeoutSeconds != 0 {
		attributeValues["auto_suspend_timeout_seconds"] = types.Int32Value(apiModel.AutoSuspendTimeoutSeconds)
	}
	attributeValues["disable_auto_suspend"] = types.BoolValue(apiModel.DisableAutoSuspend)
	/*if apiModel.ImpalaNumOfActiveCoordinators != 0 {
		attributeValues["impala_num_of_active_coordinators"] = types.Int32Value(apiModel.ImpalaNumOfActiveCoordinators)
	}*/
	if apiModel.ImpalaScaleDownDelaySeconds != 0 {
		attributeValues["impala_scale_down_delay_seconds"] = types.Int32Value(apiModel.ImpalaScaleDownDelaySeconds)
	}
	if apiModel.ImpalaScaleUpDelaySeconds != 0 {
		attributeValues["impala_scale_up_delay_seconds"] = types.Int32Value(apiModel.ImpalaScaleUpDelaySeconds)
	}
	/*if apiModel.ImpalaShutdownOfCoordinatorDelaySeconds != 0 {
		attributeValues["impala_shutdown_of_coordinator_delay_seconds"] = types.Int32Value(apiModel.ImpalaShutdownOfCoordinatorDelaySeconds)
	}*/
	if apiModel.MaxClusters != 0 {
		attributeValues["max_clusters"] = types.Int32Value(apiModel.MaxClusters)
	}
	if apiModel.MinClusters != 0 {
		attributeValues["min_clusters"] = types.Int32Value(apiModel.MinClusters)
	}

	// Convert ImpalaExecutorGroupSets if available
	/*if apiModel.ImpalaExecutorGroupSets != nil {
		attributeValues["impala_executor_group_sets"] = convertFromAPIImpalaExecutorGroupSets(apiModel.ImpalaExecutorGroupSets)
	}*/

	// Return as types.Object
	ret, _ := types.ObjectValue(attributeTypes, attributeValues)
	return ret
}

// ConvertFromImpalaExecutorGroupSetsModel converts from Impala executor group sets.
/*func convertFromAPIImpalaExecutorGroupSets(apiModel *models.ImpalaExecutorGroupSetsResponse) types.Object {
	if apiModel == nil {
		return types.ObjectNull(map[string]attr.Type{})
	}

	groupTypes := map[string]attr.Type{
		"custom1": types.ObjectType{},
		"custom2": types.ObjectType{},
		"custom3": types.ObjectType{},
		"large":   types.ObjectType{},
		"small":   types.ObjectType{},
	}

	groupValues := map[string]attr.Value{
		"custom1": convertFromAPIImpalaExecutorGroupSet(apiModel.Custom1),
		"custom2": convertFromAPIImpalaExecutorGroupSet(apiModel.Custom2),
		"custom3": convertFromAPIImpalaExecutorGroupSet(apiModel.Custom3),
		"large":   convertFromAPIImpalaExecutorGroupSet(apiModel.Large),
		"small":   convertFromAPIImpalaExecutorGroupSet(apiModel.Small),
	}

	ret, _ := types.ObjectValue(groupTypes, groupValues)
	return ret
}*/

func convertFromAPIImpalaExecutorGroupSets(apiModel *models.ImpalaExecutorGroupSetsResponse) types.List {
	if apiModel == nil {
		return types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"group_type": types.StringType,
				"settings": types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"auto_suspend_timeout_seconds": types.Int32Type,
						"disable_auto_suspend":         types.BoolType,
						"exec_group_size":              types.Int32Type,
						"max_executor_groups":          types.Int32Type,
						"min_executor_groups":          types.Int32Type,
						"trigger_scale_down_delay":     types.Int32Type,
						"trigger_scale_up_delay":       types.Int32Type,
					},
				},
			},
		})
	}

	groupList := []attr.Value{}
	groupTypes := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"group_type": types.StringType,
			"settings": types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"auto_suspend_timeout_seconds": types.Int32Type,
					"disable_auto_suspend":         types.BoolType,
					"exec_group_size":              types.Int32Type,
					"max_executor_groups":          types.Int32Type,
					"min_executor_groups":          types.Int32Type,
					"trigger_scale_down_delay":     types.Int32Type,
					"trigger_scale_up_delay":       types.Int32Type,
				},
			},
		},
	}

	// Iterate over group types
	groupMappings := map[string]*models.ImpalaExecutorGroupSetResponse{
		"custom1": apiModel.Custom1,
		"custom2": apiModel.Custom2,
		"custom3": apiModel.Custom3,
		"large":   apiModel.Large,
		"small":   apiModel.Small,
	}

	for groupType, groupData := range groupMappings {
		if groupData != nil {
			groupValue, _ := types.ObjectValue(groupTypes.AttrTypes, map[string]attr.Value{
				"group_type": types.StringValue(groupType),
				"settings":   convertFromAPIImpalaExecutorGroupSet(groupData),
			})
			groupList = append(groupList, groupValue)
		}
	}

	ret, _ := types.ListValue(groupTypes, groupList)
	return ret
}

func convertFromAPIImpalaExecutorGroupSet(apiModel *models.ImpalaExecutorGroupSetResponse) types.Object {
	if apiModel == nil {
		return types.ObjectNull(map[string]attr.Type{
			"auto_suspend_timeout_seconds": types.Int32Type,
			"disable_auto_suspend":         types.BoolType,
			"exec_group_size":              types.Int32Type,
			"max_executor_groups":          types.Int32Type,
			"min_executor_groups":          types.Int32Type,
			"trigger_scale_down_delay":     types.Int32Type,
			"trigger_scale_up_delay":       types.Int32Type,
		})
	}

	attributeTypes := map[string]attr.Type{
		"auto_suspend_timeout_seconds": types.Int32Type,
		"disable_auto_suspend":         types.BoolType,
		"exec_group_size":              types.Int32Type,
		"max_executor_groups":          types.Int32Type,
		"min_executor_groups":          types.Int32Type,
		"trigger_scale_down_delay":     types.Int32Type,
		"trigger_scale_up_delay":       types.Int32Type,
	}

	attributeValues := map[string]attr.Value{
		"auto_suspend_timeout_seconds": types.Int32Value(apiModel.AutoSuspendTimeoutSeconds),
		"disable_auto_suspend":         types.BoolValue(apiModel.DisableAutoSuspend),
		"exec_group_size":              types.Int32Value(apiModel.ExecGroupSize),
		"max_executor_groups":          types.Int32Value(apiModel.MaxExecutorGroups),
		"min_executor_groups":          types.Int32Value(apiModel.MinExecutorGroups),
		"trigger_scale_down_delay":     types.Int32Value(apiModel.TriggerScaleDownDelay),
		"trigger_scale_up_delay":       types.Int32Value(apiModel.TriggerScaleUpDelay),
	}

	ret, _ := types.ObjectValue(attributeTypes, attributeValues)
	return ret
}

func convertToAPITagRequests(model types.List) []*models.TagRequest {
	if model.IsUnknown() || model.IsNull() {
		return nil
	}

	var tagRequests []*models.TagRequest

	// Extract values from the list
	values, _ := model.ToTerraformValue(context.Background())
	var tags []TagRequest
	if err := values.As(&tags); err != nil {
		return nil
	}

	// Convert each TagRequest into API format
	for _, tag := range tags {
		tagRequests = append(tagRequests, &models.TagRequest{
			Key:   &tag.Key,
			Value: &tag.Value,
		})
	}

	return tagRequests
}

func convertFromAPITagRequests(apiTags []*models.TagResponse) types.List {
	attributeType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":   types.StringType,
			"value": types.StringType,
		},
	}

	if apiTags == nil || len(apiTags) == 0 {
		return types.ListNull(attributeType)
	}

	var tagValues []attr.Value

	for _, apiTag := range apiTags {
		if apiTag == nil {
			continue
		}
		tagValues = append(tagValues, types.ObjectValueMust(
			map[string]attr.Type{
				"key":   types.StringType,
				"value": types.StringType,
			},
			map[string]attr.Value{
				"key":   types.StringValue(*apiTag.Key),
				"value": types.StringValue(*apiTag.Value),
			},
		))
	}

	return types.ListValueMust(attributeType, tagValues)
}

func ExtractInt32FromAttribute(ctx context.Context, val attr.Value) (int32, error) {
	if val.IsUnknown() {
		return 0, fmt.Errorf("value is unknown")
	}

	tfv, err := val.ToTerraformValue(ctx)
	if err != nil {
		return 0, fmt.Errorf("error converting to Terraform value: %v", err)
	}

	convertedValue, err := basetypes.Int32Type{}.ValueFromTerraform(ctx, tfv)
	if err != nil {
		return 0, fmt.Errorf("error converting Terraform value to Int32Type: %v", err)
	}

	stringValue := convertedValue.String()
	intValue, err := strconv.ParseInt(stringValue, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("error parsing string to int32: %v", err)
	}

	return int32(intValue), nil
}

func ExtractBoolFromAttribute(ctx context.Context, val attr.Value) (bool, error) {
	if val.IsUnknown() {
		return false, fmt.Errorf("value is unknown")
	}

	tfv, err := val.ToTerraformValue(ctx)
	if err != nil {
		return false, fmt.Errorf("error converting to Terraform value: %v", err)
	}

	convertedValue, err := basetypes.BoolType{}.ValueFromTerraform(ctx, tfv)
	if err != nil {
		return false, fmt.Errorf("error converting Terraform value to BoolType: %v", err)
	}

	boolString := convertedValue.String()
	boolValue, err := strconv.ParseBool(boolString)
	if err != nil {
		return false, fmt.Errorf("error parsing string to bool: %v", err)
	}

	return boolValue, nil
}

func ExtractStringFromAttribute(ctx context.Context, attr attr.Value) (string, error) {
	if attr.IsUnknown() {
		return "", fmt.Errorf("attribute value is unknown")
	}

	tfv, err := attr.ToTerraformValue(ctx)
	if err != nil {
		return "", fmt.Errorf("error converting to Terraform value: %w", err)
	}

	var valueString string
	if err := tfv.As(&valueString); err != nil {
		return "", fmt.Errorf("error converting Terraform value to string: %w", err)
	}

	return valueString, nil
}
