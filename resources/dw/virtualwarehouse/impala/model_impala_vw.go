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
	"time"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const timeZone = time.RFC850

/*type TShirtSize string

const (
	TShirtSizeXSmall TShirtSize = "xsmall"
	TShirtSizeSmall  TShirtSize = "small"
	TShirtSizeMedium TShirtSize = "medium"
	TShirtSizeLarge  TShirtSize = "large"
)*/

/*type ImpalaOptionsCreateRequest struct {
	SpillToS3URI      string `tfsdk:"spill_to_s3_uri"`     // S3 URI in "s3://bucket/path" format for spilling to S3.
	ScratchSpaceLimit int32  `tfsdk:"scratch_space_limit"` // Scratch space limit in GiB.
}*/

type ImpalaHASettingsModel struct {
	EnableCatalogHighAvailability    types.Bool   `tfsdk:"enable_catalog_high_availability"`
	EnableShutdownOfCoordinator      types.Bool   `tfsdk:"enable_shutdown_of_coordinator"`
	EnableStatestoreHighAvailability types.Bool   `tfsdk:"enable_statestore_high_availability"`
	HighAvailabilityMode             types.String `tfsdk:"high_availability_mode"`
	NumOfActiveCoordinators          types.Int32  `tfsdk:"num_of_active_coordinators"`
	ShutdownOfCoordinatorDelaySecs   types.Int32  `tfsdk:"shutdown_of_coordinator_delay_secs"`
}

type HASettingsCreateRequest struct {
	HighAvailabilityMode             string `tfsdk:"high_availability_mode"`              // High Availability mode: DISABLED, ACTIVE_PASSIVE, or ACTIVE_ACTIVE.
	EnableShutdownOfCoordinator      bool   `tfsdk:"enable_shutdown_of_coordinator"`      // Enables the shutdown of the coordinator.
	ShutdownOfCoordinatorDelaySecs   int32  `tfsdk:"shutdown_of_coordinator_delay_secs"`  // Delay in seconds before shutting down the coordinator.
	NumOfActiveCoordinators          int32  `tfsdk:"num_of_active_coordinators"`          // Number of active coordinators.
	EnableCatalogHighAvailability    bool   `tfsdk:"enable_catalog_high_availability"`    // Enables high availability for Impala catalog.
	EnableStatestoreHighAvailability bool   `tfsdk:"enable_statestore_high_availability"` // Enables high availability for Impala Statestore.
}

type AutoscalingOptionsUpdateRequest struct {
	MinClusters               int32 `tfsdk:"min_clusters"`                 // Minimum number of available compute groups. Default: 0.
	MaxClusters               int32 `tfsdk:"max_clusters"`                 // Maximum number of available compute groups. Default: 0.
	DisableAutoSuspend        bool  `tfsdk:"disable_auto_suspend"`         // Disable auto-suspend for the Virtual Warehouse.
	AutoSuspendTimeoutSeconds int32 `tfsdk:"auto_suspend_timeout_seconds"` // Threshold for auto-suspend in seconds.
	// HiveScaleWaitTimeSeconds  int32 `tfsdk:"hive_scale_wait_time_seconds"` // Wait time before a scaling event happens.
	// HiveDesiredFreeCapacity              int32                                 `tfsdk:"hive_desired_free_capacity"`                                           // Desired free capacity for Hive.
	ImpalaScaleUpDelaySeconds            int32                                 `tfsdk:"impala_scale_up_delay_seconds"`                                        // Scale-up threshold in seconds for Impala.
	ImpalaScaleDownDelaySeconds          int32                                 `tfsdk:"impala_scale_down_delay_seconds"`                                      // Scale-down threshold in seconds for Impala.
	ImpalaShutdownOfCoordinatorDelaySecs int32                                 `tfsdk:"impala_shutdown_of_coordinator_delay_seconds" tfsdk_deprecated:"true"` // DEPRECATED: Delay in seconds before shutting down Impala coordinator.
	ImpalaNumOfActiveCoordinators        int32                                 `tfsdk:"impala_num_of_active_coordinators" tfsdk_deprecated:"true"`            // DEPRECATED: Number of active Impala coordinators.
	ImpalaExecutorGroupSets              *ImpalaExecutorGroupSetsUpdateRequest `tfsdk:"impala_executor_group_sets"`                                           // Reconfigure executor group sets for workload-aware autoscaling.
}

type ImpalaExecutorGroupSetsUpdateRequest struct {
	Small   ImpalaExecutorGroupSetUpdateRequest `json:"small"`
	Custom1 ImpalaExecutorGroupSetUpdateRequest `json:"custom1"`
	Custom2 ImpalaExecutorGroupSetUpdateRequest `json:"custom2"`
	Custom3 ImpalaExecutorGroupSetUpdateRequest `json:"custom3"`
	Large   ImpalaExecutorGroupSetUpdateRequest `json:"large"`
}

// ImpalaExecutorGroupSetsCreateRequest represents the configuration of executor group sets for workload-aware autoscaling.
type ImpalaExecutorGroupSetsCreateRequest struct {
	Small   ImpalaExecutorGroupSetCreateRequest  `tfsdk:"small"`   // Configure small executor group set for workload-aware autoscaling. Required.
	Custom1 *ImpalaExecutorGroupSetCreateRequest `tfsdk:"custom1"` // Configure first optional custom executor group set.
	Custom2 *ImpalaExecutorGroupSetCreateRequest `tfsdk:"custom2"` // Configure second optional custom executor group set.
	Custom3 *ImpalaExecutorGroupSetCreateRequest `tfsdk:"custom3"` // Configure third optional custom executor group set.
	Large   ImpalaExecutorGroupSetCreateRequest  `tfsdk:"large"`   // Configure large executor group set for workload-aware autoscaling. Required.
}

type ImpalaExecutorGroupSetCreateRequest struct {
	ExecGroupSize             int   `tfsdk:"exec_group_size"`                        // Set number of executors per executor group. Required.
	MinExecutorGroups         int   `tfsdk:"min_executor_groups"`                    // Set minimum number of executor groups allowed. Required.
	MaxExecutorGroups         int   `tfsdk:"max_executor_groups"`                    // Set maximum number of executor groups allowed. Required.
	AutoSuspendTimeoutSeconds *int  `tfsdk:"auto_suspend_timeout_seconds,omitempty"` // Set auto suspend threshold. Optional.
	DisableAutoSuspend        *bool `tfsdk:"disable_auto_suspend,omitempty"`         // Turn off auto suspend. Optional.
	TriggerScaleUpDelay       *int  `tfsdk:"trigger_scale_up_delay,omitempty"`       // Set scale-up threshold in seconds. Optional.
	TriggerScaleDownDelay     *int  `tfsdk:"trigger_scale_down_delay,omitempty"`     // Set scale-down threshold in seconds. Optional.
}

type ImpalaExecutorGroupSetUpdateRequest struct {
	ExecGroupSize             *int  `tfsdk:"exec_group_size,omitempty"`              // Set number of executors per executor group. Optional.
	MinExecutorGroups         *int  `tfsdk:"min_executor_groups,omitempty"`          // Set minimum number of executor groups allowed. Optional.
	MaxExecutorGroups         *int  `tfsdk:"max_executor_groups,omitempty"`          // Set maximum number of executor groups allowed. Optional.
	AutoSuspendTimeoutSeconds *int  `tfsdk:"auto_suspend_timeout_seconds,omitempty"` // Set auto suspend threshold. Optional.
	DisableAutoSuspend        *bool `tfsdk:"disable_auto_suspend,omitempty"`         // Turn off auto suspend. Optional.
	TriggerScaleUpDelay       *int  `tfsdk:"trigger_scale_up_delay,omitempty"`       // Set scale-up threshold in seconds. Optional.
	TriggerScaleDownDelay     *int  `tfsdk:"trigger_scale_down_delay,omitempty"`     // Set scale-down threshold in seconds. Optional.
	DeleteGroupSet            *bool `tfsdk:"delete_group_set,omitempty"`             // Delete the executor group set. Optional.
}

type ServiceConfigReq struct {
	CommonConfigs      ApplicationConfigReq            `json:"common_configs"`
	ApplicationConfigs map[string]ApplicationConfigReq `json:"application_configs"`
	LdapGroups         []string                        `json:"ldap_groups,omitempty"`
	EnableSSO          bool                            `json:"enable_sso"`
}

type ApplicationConfigReq struct {
	ConfigBlocks []ConfigBlockReq `json:"config_blocks"`
}

type ConfigBlockReq struct {
	// Define fields based on the actual ConfigBlock schema
	Name    string `json:"name"`
	Value   string `json:"value"`
	Type    string `json:"type"`
	Enabled bool   `json:"enabled"`
}

type TagRequest struct {
	Key   string `tfsdk:"key"`   // The tag's name
	Value string `tfsdk:"value"` // The associated value of the tag
}

type QueryIsolationOptionsModel struct {
	MaxNodesPerQuery types.Int32 `tfsdk:"max_nodes_per_query"`
	MaxQueries       types.Int32 `tfsdk:"max_queries"`
}

type ServiceConfigReqModel struct {
	ApplicationConfigs map[string]ApplicationConfigReqModel `tfsdk:"application_configs"`
	CommonConfigs      *ApplicationConfigReqModel           `tfsdk:"common_configs"`
	EnableSSO          basetypes.BoolValue                  `tfsdk:"enable_sso"`
	LdapGroups         []string                             `tfsdk:"ldap_groups"`
}

type ApplicationConfigReqModel struct {
	ConfigBlocks []*ConfigBlockReqModel `tfsdk:"config_blocks"`
}

type ConfigBlockReqModel struct {
	Content *ConfigContentReqModel `tfsdk:"content"`
	ID      *string                `tfsdk:"id"`
}

type ConfigContentReqModel struct {
	JSON      string            `tfsdk:"json"`
	KeyValues map[string]string `tfsdk:"key_values"`
	Text      string            `tfsdk:"text"`
}

// AutoscalingModel represents the Terraform model for autoscaling options.
type AutoscalingModel struct {
	AutoSuspendTimeoutSeconds types.Int32 `tfsdk:"auto_suspend_timeout_seconds"`
	DisableAutoSuspend        types.Bool  `tfsdk:"disable_auto_suspend"`
	// HiveDesiredFreeCapacity                 types.Int32                   `tfsdk:"hive_desired_free_capacity"`
	// HiveScaleWaitTimeSeconds                types.Int32                   `tfsdk:"hive_scale_wait_time_seconds"`
	ImpalaNumOfActiveCoordinators           types.Int32                   `tfsdk:"impala_num_of_active_coordinators"`
	ImpalaScaleDownDelaySeconds             types.Int32                   `tfsdk:"impala_scale_down_delay_seconds"`
	ImpalaScaleUpDelaySeconds               types.Int32                   `tfsdk:"impala_scale_up_delay_seconds"`
	ImpalaShutdownOfCoordinatorDelaySeconds types.Int32                   `tfsdk:"impala_shutdown_of_coordinator_delay_seconds"`
	MaxClusters                             types.Int32                   `tfsdk:"max_clusters"`
	MinClusters                             types.Int32                   `tfsdk:"min_clusters"`
	ImpalaExecutorGroupSets                 *ImpalaExecutorGroupSetsModel `tfsdk:"impala_executor_group_sets"`
}

// ImpalaExecutorGroupSetsModel represents the Terraform model for executor group sets.
type ImpalaExecutorGroupSetsModel struct {
	Custom1 *ImpalaExecutorGroupSetModel `tfsdk:"custom1"`
	Custom2 *ImpalaExecutorGroupSetModel `tfsdk:"custom2"`
	Custom3 *ImpalaExecutorGroupSetModel `tfsdk:"custom3"`
	Large   *ImpalaExecutorGroupSetModel `tfsdk:"large"`
	Small   *ImpalaExecutorGroupSetModel `tfsdk:"small"`
}

// ImpalaExecutorGroupSetModel represents an individual executor group set.
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
	ID                     types.String                `tfsdk:"id"`
	ClusterID              types.String                `tfsdk:"cluster_id"`
	DatabaseCatalogID      types.String                `tfsdk:"database_catalog_id"`
	Name                   types.String                `tfsdk:"name"`
	LastUpdated            types.String                `tfsdk:"last_updated"`
	Status                 types.String                `tfsdk:"status"`
	ImageVersion           types.String                `tfsdk:"image_version"`
	InstanceType           types.String                `tfsdk:"instance_type"`
	TShirtSize             types.String                `tfsdk:"tshirt_size"`
	NodeCount              types.Int32                 `tfsdk:"node_count"`
	AvailabilityZone       types.String                `tfsdk:"availability_zone"`
	EnableUnifiedAnalytics types.Bool                  `tfsdk:"enable_unified_analytics"`
	ImpalaOptions          types.Object                `tfsdk:"impala_options"`
	ImpalaHASettings       *ImpalaHASettingsModel      `tfsdk:"impala_ha_settings"`
	Autoscaling            *AutoscalingModel           `tfsdk:"autoscaling"`
	Config                 *ServiceConfigReqModel      `tfsdk:"config"`
	QueryIsolationOptions  *QueryIsolationOptionsModel `tfsdk:"query_isolation_options"`
	Tags                   *[]TagRequest               `tfsdk:"tags"`
	ResourcePool           types.String                `tfsdk:"resource_pool"`
	HiveAuthenticationMode types.String                `tfsdk:"hive_authentication_mode"`
	PlatformJwtAuth        types.Bool                  `tfsdk:"platform_jwt_auth"`
	ImpalaQueryLog         types.Bool                  `tfsdk:"impala_query_log"`
	EbsLLAPSpillGB         types.Int64                 `tfsdk:"ebs_llap_spill_gb"`
	// HiveServerHaMode       types.String                `tfsdk:"hive_server_ha_mode"`

	PollingOptions *utils.PollingOptions `tfsdk:"polling_options"`
}

type QueryIsolationOptionsRequest struct {
	MaxQueries       int32 `json:"maxQueries,omitempty"`       // Default: 0, disables query isolation when 0
	MaxNodesPerQuery int32 `json:"maxNodesPerQuery,omitempty"` // Default: 0, disables query isolation when 0
}

type ImpalaOptionsModel struct {
	ScratchSpaceLimit basetypes.Int32Value  `tfsdk:"scratch_space_limit"`
	SpillToS3URI      basetypes.StringValue `tfsdk:"spill_to_s3_uri"`
}

type AutoscalingOptionsCreateRequest struct {
	MinClusters               int32 `json:"minClusters,omitempty"`
	MaxClusters               int32 `json:"maxClusters,omitempty"`
	DisableAutoSuspend        bool  `json:"disableAutoSuspend,omitempty"`
	AutoSuspendTimeoutSeconds int32 `json:"autoSuspendTimeoutSeconds,omitempty"`
	EnableUnifiedAnalytics    *bool `json:"enableUnifiedAnalytics,omitempty"` // Deprecated, nullable if not used
	// HiveScaleWaitTimeSeconds  int32 `json:"hiveScaleWaitTimeSeconds,omitempty"`
	// HiveDesiredFreeCapacity                 int32                                `json:"hiveDesiredFreeCapacity,omitempty"`
	ImpalaHighAvailabilityMode              string                               `json:"impalaHighAvailabilityMode,omitempty"`
	ImpalaScaleUpDelaySeconds               int32                                `json:"impalaScaleUpDelaySeconds,omitempty"`
	ImpalaScaleDownDelaySeconds             int32                                `json:"impalaScaleDownDelaySeconds,omitempty"`
	ImpalaEnableShutdownOfCoordinator       bool                                 `json:"impalaEnableShutdownOfCoordinator,omitempty"`
	ImpalaShutdownOfCoordinatorDelaySeconds int32                                `json:"impalaShutdownOfCoordinatorDelaySeconds,omitempty"`
	ImpalaNumOfActiveCoordinators           int32                                `json:"impalaNumOfActiveCoordinators,omitempty"`
	ImpalaEnableCatalogHighAvailability     types.Bool                           `json:"impalaEnableCatalogHighAvailability,omitempty"`
	ImpalaExecutorGroupSets                 ImpalaExecutorGroupSetsCreateRequest `json:"impalaExecutorGroupSets"`
}

func (p *resourceModel) setFromDescribeVwResponse(resp *models.DescribeVwResponse) {
	p.ID = types.StringValue(resp.Vw.ID)
	p.DatabaseCatalogID = types.StringValue(resp.Vw.DbcID)
	p.Name = types.StringValue(resp.Vw.Name)
	p.Status = types.StringValue(resp.Vw.Status)
	p.ImageVersion = types.StringValue(resp.Vw.CdhVersion)
	p.LastUpdated = types.StringValue(time.Now().Format(timeZone))

	/*if resp.Vw.HiveServerHaMode != nil {
		p.InstanceType = types.StringValue(resp.Vw.InstanceType)
	}*/

	if resp.Vw.NodeCount != 0 {
		p.NodeCount = types.Int32Value(resp.Vw.NodeCount)
	}

	if resp.Vw.InstanceType != "" {
		p.InstanceType = types.StringValue(resp.Vw.InstanceType)
	}

	if resp.Vw.AvailabilityZone != "" {
		p.AvailabilityZone = types.StringValue(resp.Vw.AvailabilityZone)
	}

	p.EnableUnifiedAnalytics = types.BoolValue(resp.Vw.EnableUnifiedAnalytics)

	if resp.Vw.ImpalaOptions != nil {
		p.ImpalaOptions = convertFromAPIImpalaOptions(resp.Vw.ImpalaOptions)
	}
	if resp.Vw.ImpalaHaSettingsOptions != nil {
		p.ImpalaHASettings = convertFromAPIModel(resp.Vw.ImpalaHaSettingsOptions)
	}

	if resp.Vw.AutoscalingOptions != nil {
		p.Autoscaling = ConvertFromAutoscalingModel(resp.Vw.AutoscalingOptions)
	}

	if resp.Vw.QueryIsolationOptions != nil {
		p.QueryIsolationOptions = convertFromAPIQueryIsolationOptions(resp.Vw.QueryIsolationOptions)
	}

	if len(resp.Vw.Tags) != 0 {
		if p.Tags == nil {
			p.Tags = &[]TagRequest{}
		}

		for i := 0; i < len(resp.Vw.Tags); i++ {
			newTag := TagRequest{
				Key:   *(resp.Vw.Tags[i].Key),
				Value: *(resp.Vw.Tags[i].Value),
			}
			*p.Tags = append(*p.Tags, newTag)
		}
	}

	if resp.Vw.ResourcePool != "" {
		p.ResourcePool = types.StringValue(resp.Vw.ResourcePool)
	}

	if resp.Vw.HiveAuthenticationMode != nil {
		p.HiveAuthenticationMode = types.StringValue(*(resp.Vw.HiveAuthenticationMode))
	}

	p.ImpalaQueryLog = types.BoolValue(resp.Vw.ImpalaQueryLog)

	if resp.Vw.EbsLLAPSpillGB != 0 {
		p.EbsLLAPSpillGB = types.Int64Value(int64(resp.Vw.EbsLLAPSpillGB))
	}

	/*if resp.Vw.HiveServerHaMode != nil {
		p.HiveServerHaMode = types.StringValue(*(resp.Vw.HiveServerHaMode))
	}*/

}

func (p *resourceModel) GetPollingOptions() *utils.PollingOptions {
	return p.PollingOptions
}

func convertToAPIModel(model *ImpalaHASettingsModel) *models.ImpalaHASettingsCreateRequest {
	if model == nil {
		return nil
	}

	return &models.ImpalaHASettingsCreateRequest{
		EnableCatalogHighAvailability:     model.EnableCatalogHighAvailability.ValueBool(),
		EnableShutdownOfCoordinator:       model.EnableShutdownOfCoordinator.ValueBool(),
		EnableStatestoreHighAvailability:  model.EnableStatestoreHighAvailability.ValueBool(),
		HighAvailabilityMode:              models.ImpalaHighAvailabilityMode(model.HighAvailabilityMode.ValueString()),
		NumOfActiveCoordinators:           model.NumOfActiveCoordinators.ValueInt32(),
		ShutdownOfCoordinatorDelaySeconds: model.ShutdownOfCoordinatorDelaySecs.ValueInt32(),
	}
}

func convertFromAPIModel(apiModel *models.ImpalaHASettingsOptionsResponse) *ImpalaHASettingsModel {
	if apiModel == nil {
		return nil
	}

	return &ImpalaHASettingsModel{
		EnableCatalogHighAvailability:    types.BoolValue(apiModel.EnableCatalogHighAvailability),
		EnableShutdownOfCoordinator:      types.BoolValue(apiModel.EnableShutdownOfCoordinator),
		EnableStatestoreHighAvailability: types.BoolValue(apiModel.EnableStatestoreHighAvailability),
		HighAvailabilityMode:             types.StringValue(string(apiModel.HighAvailabilityMode)),
		NumOfActiveCoordinators:          types.Int32Value(apiModel.NumOfActiveCoordinators),
		ShutdownOfCoordinatorDelaySecs:   types.Int32Value(apiModel.ShutdownOfCoordinatorDelaySeconds),
	}
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
		val, err := scratchSpaceLimit.ToTerraformValue(context.Background())
		if err == nil {
			var valueInt32 int32
			if err := val.As(&valueInt32); err == nil {
				req.ScratchSpaceLimit = valueInt32
			}
		}
	}

	if hasSpillToS3URI && !spillToS3URI.IsUnknown() {
		val, err := spillToS3URI.ToTerraformValue(context.Background())
		if err == nil {
			var valueString string
			if err := val.As(&valueString); err == nil {
				req.SpillToS3URI = valueString
			}
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

	impalaOptions := map[string]attr.Value{}

	// Handle ScratchSpaceLimit
	if apiModel.ScratchSpaceLimit != 0 {
		impalaOptions["scratch_space_limit"] = types.Int32Value(apiModel.ScratchSpaceLimit)
	} else {
		impalaOptions["scratch_space_limit"] = types.Int32Null() // Null if missing
	}

	// Handle SpillToS3URI
	if apiModel.SpillToS3URI != "" {
		impalaOptions["spill_to_s3_uri"] = types.StringValue(apiModel.SpillToS3URI)
	} else {
		impalaOptions["spill_to_s3_uri"] = types.StringNull() // Null if empty
	}

	// Return as types.Object
	ret, _ := types.ObjectValue(attributeTypes, attributeValues)
	return ret
}

func convertToAPIQueryIsolationOptions(model *QueryIsolationOptionsModel) *models.QueryIsolationOptionsRequest {
	if model == nil {
		return nil
	}

	return &models.QueryIsolationOptionsRequest{
		MaxNodesPerQuery: model.MaxNodesPerQuery.ValueInt32(),
		MaxQueries:       model.MaxQueries.ValueInt32(),
	}
}

func convertFromAPIQueryIsolationOptions(apiModel *models.QueryIsolationOptionsResponse) *QueryIsolationOptionsModel {
	if apiModel == nil {
		return nil
	}

	return &QueryIsolationOptionsModel{
		MaxNodesPerQuery: types.Int32Value(apiModel.MaxNodesPerQuery),
		MaxQueries:       types.Int32Value(apiModel.MaxQueries),
	}
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

func ConvertToAutoscalingModel(req *AutoscalingModel) *models.AutoscalingOptionsCreateRequest {
	if req == nil {
		return nil
	}

	return &models.AutoscalingOptionsCreateRequest{
		AutoSuspendTimeoutSeconds: req.AutoSuspendTimeoutSeconds.ValueInt32(),
		DisableAutoSuspend:        req.DisableAutoSuspend.ValueBool(),
		// HiveDesiredFreeCapacity:                 req.HiveDesiredFreeCapacity.ValueInt32(),
		// HiveScaleWaitTimeSeconds:                req.HiveScaleWaitTimeSeconds.ValueInt32(),
		ImpalaNumOfActiveCoordinators:           req.ImpalaNumOfActiveCoordinators.ValueInt32(),
		ImpalaScaleDownDelaySeconds:             req.ImpalaScaleDownDelaySeconds.ValueInt32(),
		ImpalaScaleUpDelaySeconds:               req.ImpalaScaleUpDelaySeconds.ValueInt32(),
		ImpalaShutdownOfCoordinatorDelaySeconds: req.ImpalaShutdownOfCoordinatorDelaySeconds.ValueInt32(),
		MaxClusters:                             convertIntPtr(req.MaxClusters.ValueInt32()),
		MinClusters:                             convertIntPtr(req.MinClusters.ValueInt32()),
		ImpalaExecutorGroupSets:                 ConvertToImpalaExecutorGroupSetsModel(req.ImpalaExecutorGroupSets),
	}
}

// ConvertToImpalaExecutorGroupSetsModel converts Impala executor group sets.
func ConvertToImpalaExecutorGroupSetsModel(req *ImpalaExecutorGroupSetsModel) *models.ImpalaExecutorGroupSetsCreateRequest {
	if req == nil {
		return nil
	}

	return &models.ImpalaExecutorGroupSetsCreateRequest{
		Custom1: ConvertToImpalaExecutorGroupSetModel(req.Custom1),
		Custom2: ConvertToImpalaExecutorGroupSetModel(req.Custom2),
		Custom3: ConvertToImpalaExecutorGroupSetModel(req.Custom3),
		Large:   ConvertToImpalaExecutorGroupSetModel(req.Large),
		Small:   ConvertToImpalaExecutorGroupSetModel(req.Small),
	}
}

// ConvertToImpalaExecutorGroupSetModel converts an executor group set.
func ConvertToImpalaExecutorGroupSetModel(req *ImpalaExecutorGroupSetModel) *models.ImpalaExecutorGroupSetCreateRequest {
	if req == nil {
		return nil
	}

	return &models.ImpalaExecutorGroupSetCreateRequest{
		AutoSuspendTimeoutSeconds: req.AutoSuspendTimeoutSeconds,
		DisableAutoSuspend:        req.DisableAutoSuspend,
		ExecGroupSize:             req.ExecGroupSize,
		MaxExecutorGroups:         req.MaxExecutorGroups,
		MinExecutorGroups:         req.MinExecutorGroups,
		TriggerScaleDownDelay:     req.TriggerScaleDownDelay,
		TriggerScaleUpDelay:       req.TriggerScaleUpDelay,
	}
}

// convertIntPtr safely converts an int32 to a *int32.
func convertIntPtr(ptr int32) *int32 {
	if ptr == 0 {
		return nil
	}
	r := &ptr
	return r
}

func ConvertFromAutoscalingModel(req *models.AutoscalingOptionsResponse) *AutoscalingModel {
	if req == nil {
		return nil
	}

	return &AutoscalingModel{
		AutoSuspendTimeoutSeconds: types.Int32Value(req.AutoSuspendTimeoutSeconds),
		DisableAutoSuspend:        types.BoolValue(req.DisableAutoSuspend),
		// HiveDesiredFreeCapacity:                 types.Int32Value(req.HiveDesiredFreeCapacity),
		// HiveScaleWaitTimeSeconds:                types.Int32Value(req.HiveScaleWaitTimeSeconds),
		ImpalaNumOfActiveCoordinators:           types.Int32Value(req.ImpalaNumOfActiveCoordinators),
		ImpalaScaleDownDelaySeconds:             types.Int32Value(req.ImpalaScaleDownDelaySeconds),
		ImpalaScaleUpDelaySeconds:               types.Int32Value(req.ImpalaScaleUpDelaySeconds),
		ImpalaShutdownOfCoordinatorDelaySeconds: types.Int32Value(req.ImpalaShutdownOfCoordinatorDelaySeconds),
		MaxClusters:                             convertInt32Ptr(&req.MaxClusters),
		MinClusters:                             convertInt32Ptr(&req.MinClusters),
		ImpalaExecutorGroupSets:                 ConvertFromImpalaExecutorGroupSetsModel(req.ImpalaExecutorGroupSets),
	}
}

// ConvertFromImpalaExecutorGroupSetsModel converts from Impala executor group sets.
func ConvertFromImpalaExecutorGroupSetsModel(req *models.ImpalaExecutorGroupSetsResponse) *ImpalaExecutorGroupSetsModel {
	if req == nil {
		return nil
	}

	return &ImpalaExecutorGroupSetsModel{
		Custom1: ConvertFromImpalaExecutorGroupSetModel(req.Custom1),
		Custom2: ConvertFromImpalaExecutorGroupSetModel(req.Custom2),
		Custom3: ConvertFromImpalaExecutorGroupSetModel(req.Custom3),
		Large:   ConvertFromImpalaExecutorGroupSetModel(req.Large),
		Small:   ConvertFromImpalaExecutorGroupSetModel(req.Small),
	}
}

// ConvertFromImpalaExecutorGroupSetModel converts from an executor group set.
func ConvertFromImpalaExecutorGroupSetModel(req *models.ImpalaExecutorGroupSetResponse) *ImpalaExecutorGroupSetModel {
	if req == nil {
		return nil
	}

	return &ImpalaExecutorGroupSetModel{
		AutoSuspendTimeoutSeconds: req.AutoSuspendTimeoutSeconds,
		DisableAutoSuspend:        req.DisableAutoSuspend,
		ExecGroupSize:             req.ExecGroupSize,
		MaxExecutorGroups:         req.MaxExecutorGroups,
		MinExecutorGroups:         req.MinExecutorGroups,
		TriggerScaleDownDelay:     req.TriggerScaleDownDelay,
		TriggerScaleUpDelay:       req.TriggerScaleUpDelay,
	}
}

// convertInt32Ptr safely converts an int32 pointer to a Terraform Int32 type.
func convertInt32Ptr(value *int32) types.Int32 {
	if value == nil {
		return types.Int32Null()
	}
	return types.Int32Value(*value)
}
