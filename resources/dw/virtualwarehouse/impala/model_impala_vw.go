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
	"time"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
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

type HASettingsCreateRequest struct {
	HighAvailabilityMode             string `tfsdk:"high_availability_mode"`              // High Availability mode: DISABLED, ACTIVE_PASSIVE, or ACTIVE_ACTIVE.
	EnableShutdownOfCoordinator      bool   `tfsdk:"enable_shutdown_of_coordinator"`      // Enables the shutdown of the coordinator.
	ShutdownOfCoordinatorDelaySecs   int32  `tfsdk:"shutdown_of_coordinator_delay_secs"`  // Delay in seconds before shutting down the coordinator.
	NumOfActiveCoordinators          int32  `tfsdk:"num_of_active_coordinators"`          // Number of active coordinators.
	EnableCatalogHighAvailability    bool   `tfsdk:"enable_catalog_high_availability"`    // Enables high availability for Impala catalog.
	EnableStatestoreHighAvailability bool   `tfsdk:"enable_statestore_high_availability"` // Enables high availability for Impala Statestore.
}

type AutoscalingOptionsUpdateRequest struct {
	MinClusters                          int32                                 `tfsdk:"min_clusters"`                                                      // Minimum number of available compute groups. Default: 0.
	MaxClusters                          int32                                 `tfsdk:"max_clusters"`                                                      // Maximum number of available compute groups. Default: 0.
	DisableAutoSuspend                   bool                                  `tfsdk:"disable_auto_suspend"`                                              // Disable auto-suspend for the Virtual Warehouse.
	AutoSuspendTimeoutSeconds            int32                                 `tfsdk:"auto_suspend_timeout_seconds"`                                      // Threshold for auto-suspend in seconds.
	HiveScaleWaitTimeSeconds             int32                                 `tfsdk:"hive_scale_wait_time_seconds"`                                      // Wait time before a scaling event happens.
	HiveDesiredFreeCapacity              int32                                 `tfsdk:"hive_desired_free_capacity"`                                        // Desired free capacity for Hive.
	ImpalaScaleUpDelaySeconds            int32                                 `tfsdk:"impala_scale_up_delay_seconds"`                                     // Scale-up threshold in seconds for Impala.
	ImpalaScaleDownDelaySeconds          int32                                 `tfsdk:"impala_scale_down_delay_seconds"`                                   // Scale-down threshold in seconds for Impala.
	ImpalaShutdownOfCoordinatorDelaySecs int32                                 `tfsdk:"impala_shutdown_of_coordinator_delay_secs" tfsdk_deprecated:"true"` // DEPRECATED: Delay in seconds before shutting down Impala coordinator.
	ImpalaNumOfActiveCoordinators        int32                                 `tfsdk:"impala_num_of_active_coordinators" tfsdk_deprecated:"true"`         // DEPRECATED: Number of active Impala coordinators.
	ImpalaExecutorGroupSets              *ImpalaExecutorGroupSetsUpdateRequest `tfsdk:"impala_executor_group_sets"`                                        // Reconfigure executor group sets for workload-aware autoscaling.
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
	Key   string `json:"key"`   // The tag's name
	Value string `json:"value"` // The associated value of the tag
}

type resourceModel struct {
	ID                     types.String                            `tfsdk:"id"`
	ClusterID              types.String                            `tfsdk:"cluster_id"`
	DatabaseCatalogID      types.String                            `tfsdk:"database_catalog_id"`
	Name                   types.String                            `tfsdk:"name"`
	LastUpdated            types.String                            `tfsdk:"last_updated"`
	Status                 types.String                            `tfsdk:"status"`
	ImageVersion           types.String                            `tfsdk:"image_version"`
	InstanceType           types.String                            `tfsdk:"instance_type"`
	TShirtSize             types.String                            `tfsdk:"tshirt_size"`
	NodeCount              types.Int32                             `tfsdk:"node_count"`
	AvailabilityZone       types.String                            `tfsdk:"availability_zone"`
	EnableUnifiedAnalytics types.Bool                              `tfsdk:"enable_unified_analytics"`
	ImpalaOptions          *models.ImpalaOptionsCreateRequest      `tfsdk:"impala_options"`
	ImpalaHASettings       *models.ImpalaHASettingsCreateRequest   `tfsdk:"impala_ha_settings"`
	Autoscaling            *models.AutoscalingOptionsCreateRequest `tfsdk:"autoscaling"`
	Config                 *models.ServiceConfigReq                `tfsdk:"config"`
	QueryIsolationOptions  *models.QueryIsolationOptionsRequest    `tfsdk:"query_isolation_options"`
	Tags                   *[]TagRequest                           `tfsdk:"tags"`
	ResourcePool           types.String                            `tfsdk:"resource_pool"`
	HiveAuthenticationMode types.String                            `tfsdk:"hive_authentication_mode"`
	PlatformJwtAuth        types.Bool                              `tfsdk:"platform_jwt_auth"`
	ImpalaQueryLog         types.Bool                              `tfsdk:"impala_query_log"`
	EbsLLAPSpillGB         types.Int64                             `tfsdk:"ebs_llap_spill_gb"`
	HiveServerHaMode       types.String                            `tfsdk:"hive_server_ha_mode"`

	PollingOptions *utils.PollingOptions `tfsdk:"polling_options"`
}

type QueryIsolationOptionsRequest struct {
	MaxQueries       int32 `json:"maxQueries,omitempty"`       // Default: 0, disables query isolation when 0
	MaxNodesPerQuery int32 `json:"maxNodesPerQuery,omitempty"` // Default: 0, disables query isolation when 0
}

type AutoscalingOptionsCreateRequest struct {
	MinClusters                             int32                                `json:"minClusters,omitempty"`
	MaxClusters                             int32                                `json:"maxClusters,omitempty"`
	DisableAutoSuspend                      bool                                 `json:"disableAutoSuspend,omitempty"`
	AutoSuspendTimeoutSeconds               int32                                `json:"autoSuspendTimeoutSeconds,omitempty"`
	EnableUnifiedAnalytics                  *bool                                `json:"enableUnifiedAnalytics,omitempty"` // Deprecated, nullable if not used
	HiveScaleWaitTimeSeconds                int32                                `json:"hiveScaleWaitTimeSeconds,omitempty"`
	HiveDesiredFreeCapacity                 int32                                `json:"hiveDesiredFreeCapacity,omitempty"`
	ImpalaHighAvailabilityMode              string                               `json:"impalaHighAvailabilityMode,omitempty"`
	ImpalaScaleUpDelaySeconds               int32                                `json:"impalaScaleUpDelaySeconds,omitempty"`
	ImpalaScaleDownDelaySeconds             int32                                `json:"impalaScaleDownDelaySeconds,omitempty"`
	ImpalaEnableShutdownOfCoordinator       bool                                 `json:"impalaEnableShutdownOfCoordinator,omitempty"`
	ImpalaShutdownOfCoordinatorDelaySeconds int32                                `json:"impalaShutdownOfCoordinatorDelaySeconds,omitempty"`
	ImpalaNumOfActiveCoordinators           int32                                `json:"impalaNumOfActiveCoordinators,omitempty"`
	ImpalaEnableCatalogHighAvailability     bool                                 `json:"impalaEnableCatalogHighAvailability,omitempty"`
	ImpalaExecutorGroupSets                 ImpalaExecutorGroupSetsCreateRequest `json:"impalaExecutorGroupSets"`
}

func (p *resourceModel) setFromDescribeVwResponse(resp *models.DescribeVwResponse) {
	p.ID = types.StringValue(resp.Vw.ID)
	p.DatabaseCatalogID = types.StringValue(resp.Vw.DbcID)
	p.Name = types.StringValue(resp.Vw.Name)
	p.Status = types.StringValue(resp.Vw.Status)
	p.ImageVersion = types.StringValue(resp.Vw.CdhVersion)
	p.LastUpdated = types.StringValue(time.Now().Format(timeZone))

	if resp.Vw.HiveServerHaMode != nil {
		p.InstanceType = types.StringValue(resp.Vw.InstanceType)
	}
	// Not present in VW
	// p.TShirtSize = types.StringValue(resp.Vw.TshirtSize)

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
		p.ImpalaOptions = (*models.ImpalaOptionsCreateRequest)(resp.Vw.ImpalaOptions)
	}
	if resp.Vw.ImpalaHaSettingsOptions != nil {
		p.ImpalaHASettings = (*models.ImpalaHASettingsCreateRequest)(resp.Vw.ImpalaHaSettingsOptions)
	}
	/* type not compatible
	Cannot use 'resp.Vw.AutoscalingOptions' (type *AutoscalingOptionsResponse) as the type *models.AutoscalingOptionsCreateRequest
	if resp.Vw.AutoscalingOptions != nil {
		p.Autoscaling = resp.Vw.AutoscalingOptions
	}*/

	/*Missing
	if resp.Vw.ServiceConfigReq != nil {
		p.ServiceConfigReq = (*models.ImpalaHASettingsCreateRequest)(resp.Vw.ImpalaHaSettingsOptions)
	}*/

	if resp.Vw.QueryIsolationOptions != nil {
		p.QueryIsolationOptions = (*models.QueryIsolationOptionsRequest)(resp.Vw.QueryIsolationOptions)
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

	// Missing
	// p.PlatformJwtAuth = types.BoolValue(resp.Vw.PlatformJwtAuth)

	p.ImpalaQueryLog = types.BoolValue(resp.Vw.ImpalaQueryLog)

	if resp.Vw.EbsLLAPSpillGB != 0 {
		p.EbsLLAPSpillGB = types.Int64Value(int64(resp.Vw.EbsLLAPSpillGB))
	}

	if resp.Vw.HiveServerHaMode != nil {
		p.HiveServerHaMode = types.StringValue(*(resp.Vw.HiveServerHaMode))
	}

}

func (p *resourceModel) GetPollingOptions() *utils.PollingOptions {
	return p.PollingOptions
}
