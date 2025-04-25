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

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
	"github.com/cloudera/terraform-provider-cdp/utils"
)

const timeZone = time.RFC850

type TagRequest struct {
	Key   string `tfsdk:"key"`   // The tag's name
	Value string `tfsdk:"value"` // The associated value of the tag
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
	ImpalaOptions          types.Object `tfsdk:"aws_options"`
	ImpalaHASettings       types.Object `tfsdk:"ha_settings"`
	// TODO Prateek Add this validation either use node_count or autoscaling options, not both
	Autoscaling types.Object `tfsdk:"autoscaling"`
	EnableSSO   types.Bool   `tfsdk:"enable_sso"`
	// TODO Prateek Make this work in a way that this setting is only accepted if enable_unified_analytics is true
	QueryIsolationOptions types.Object `tfsdk:"query_isolation_options"`
	Tags                  types.List   `tfsdk:"tags"`
	PlatformJwtAuth       types.Bool   `tfsdk:"platform_jwt_auth"`
	ImpalaQueryLog        types.Bool   `tfsdk:"query_log"`

	PollingOptions *utils.PollingOptions `tfsdk:"polling_options"`
}

func (p *resourceModel) setFromDescribeVwResponse(resp *models.DescribeVwResponse, ctx context.Context) {
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

	setInt32IfPositive(&p.NodeCount, resp.Vw.NodeCount)
	setStringIfNotEmpty(&p.InstanceType, resp.Vw.InstanceType)
	setStringIfNotEmpty(&p.AvailabilityZone, resp.Vw.AvailabilityZone)

	p.EnableUnifiedAnalytics = types.BoolValue(resp.Vw.EnableUnifiedAnalytics)
	p.ImpalaQueryLog = types.BoolValue(resp.Vw.ImpalaQueryLog)

	if resp.Vw.ImpalaOptions != nil {
		p.ImpalaOptions = convertFromAPIImpalaOptions(resp.Vw.ImpalaOptions)
	}
	if resp.Vw.ImpalaHaSettingsOptions != nil {
		p.ImpalaHASettings = convertFromAPIImpalaHASettings(resp.Vw.ImpalaHaSettingsOptions, ctx)
	}
	if resp.Vw.AutoscalingOptions != nil {
		p.Autoscaling = convertFromAPIAutoscaling(resp.Vw.AutoscalingOptions)
	}
	if resp.Vw.QueryIsolationOptions != nil {
		p.QueryIsolationOptions = convertFromAPIQueryIsolationOptions(resp.Vw.QueryIsolationOptions)
	}
	if resp.Vw.SupportedAuthMethods != nil {
		config := convertFromAPIServiceConfigReq(resp.Vw.SupportedAuthMethods)
		if enableSSOAttr, ok := config.Attributes()["enable_sso"].(types.Bool); ok {
			p.EnableSSO = enableSSOAttr
		} else {
			p.EnableSSO = types.BoolValue(false)
		}
	}

	if len(resp.Vw.Tags) != 0 {
		p.Tags = convertFromAPITagRequests(resp.Vw.Tags)
	}
}

func (p *resourceModel) GetPollingOptions() *utils.PollingOptions {
	return p.PollingOptions
}

func convertToAPIImpalaHASettings(model types.Object, ctx context.Context) (*models.ImpalaHASettingsCreateRequest, error) {
	attributes := model.Attributes()

	req := &models.ImpalaHASettingsCreateRequest{}

	if val, ok := attributes["enable_catalog_high_availability"]; ok {
		if value, err := ExtractBoolFromAttribute(ctx, val.(basetypes.BoolValue)); err != nil {
			tflog.Error(ctx, "Error extracting bool for EnableCatalogHighAvailability", map[string]interface{}{"error": err})
			return req, fmt.Errorf("error extracting bool for EnableCatalogHighAvailability: %w", err)
		} else {
			req.EnableCatalogHighAvailability = value
		}
	}

	if val, ok := attributes["enable_shutdown_of_coordinator"]; ok {
		if value, err := ExtractBoolFromAttribute(ctx, val.(basetypes.BoolValue)); err != nil {
			tflog.Error(ctx, "Error extracting bool for EnableShutdownOfCoordinator", map[string]interface{}{"error": err})
			return req, fmt.Errorf("error extracting bool for EnableShutdownOfCoordinator: %w", err)
		} else {
			req.EnableShutdownOfCoordinator = value
		}
	}

	if val, ok := attributes["enable_statestore_high_availability"]; ok {
		if value, err := ExtractBoolFromAttribute(ctx, val.(basetypes.BoolValue)); err != nil {
			tflog.Error(ctx, "Error extracting bool for EnableStatestoreHighAvailability", map[string]interface{}{"error": err})
			return req, fmt.Errorf("error extracting bool for EnableStatestoreHighAvailability: %w", err)
		} else {
			req.EnableStatestoreHighAvailability = value
		}
	}

	if val, ok := attributes["high_availability_mode"]; ok {
		if value, err := ExtractStringFromAttribute(ctx, val.(basetypes.StringValue)); err != nil {
			tflog.Error(ctx, "Error extracting string for HighAvailabilityMode", map[string]interface{}{"error": err})
			return req, fmt.Errorf("error extracting string for HighAvailabilityMode: %w", err)
		} else {
			req.HighAvailabilityMode = models.ImpalaHighAvailabilityMode(value)
		}
	}

	if val, ok := attributes["num_of_active_coordinators"]; ok {
		if value, err := val.(basetypes.Int32Value).ToInt32Value(ctx); err != nil {
			tflog.Error(ctx, "Error extracting int32 for NumOfActiveCoordinators", map[string]interface{}{"error": err})
			return req, fmt.Errorf("error extracting int32 for NumOfActiveCoordinators: %v", err)
		} else {
			req.NumOfActiveCoordinators = value.ValueInt32()
		}
	}

	if val, ok := attributes["shutdown_of_coordinator_delay_secs"]; ok {
		if value, err := val.(basetypes.Int32Value).ToInt32Value(ctx); err != nil {
			tflog.Error(ctx, "Error extracting int32 for ShutdownOfCoordinatorDelaySeconds", map[string]interface{}{"error": err})
			return req, fmt.Errorf("error extracting int32 for ShutdownOfCoordinatorDelaySeconds: %v", err)
		} else {
			req.ShutdownOfCoordinatorDelaySeconds = value.ValueInt32()
		}
	}
	return req, nil
}

func convertFromAPIImpalaHASettings(apiModel *models.ImpalaHASettingsOptionsResponse, ctx context.Context) types.Object {

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
		tflog.Error(ctx, "Error creating ObjectValue", map[string]interface{}{"error": err})
	}
	return ret
}

func convertToAPIImpalaOptions(model types.Object, ctx context.Context) (*models.ImpalaOptionsCreateRequest, error) {
	attributes := model.Attributes()

	scratchSpaceLimit, hasScratchSpace := attributes["scratch_space_limit"]
	spillToS3URI, hasSpillToS3URI := attributes["spill_to_s3_uri"]

	req := &models.ImpalaOptionsCreateRequest{}

	if hasScratchSpace && !scratchSpaceLimit.IsUnknown() {
		intValue, err := scratchSpaceLimit.(types.Int32).ToInt32Value(ctx)
		if err != nil {
			tflog.Error(ctx, "Error extracting int32 for ScratchSpaceLimit", map[string]interface{}{"error": err})
			return req, fmt.Errorf("error extracting int32 for ScratchSpaceLimit: %v", err)
		} else {
			tflog.Debug(ctx, "Assigned value to ScratchSpaceLimit", map[string]interface{}{"ScratchSpaceLimit": req.ScratchSpaceLimit})
			req.ScratchSpaceLimit = intValue.ValueInt32()
		}
		// If this is set, don't set spillToS3URI
		return req, nil
	}

	if hasSpillToS3URI && !spillToS3URI.IsUnknown() {
		valueString, err := ExtractStringFromAttribute(context.Background(), spillToS3URI)
		if err != nil {
			tflog.Error(ctx, "Error extracting string for SpillToS3URI", map[string]interface{}{"error": err})
		} else {
			tflog.Debug(ctx, "Assigned value to SpillToS3URI", map[string]interface{}{"SpillToS3URI": req.SpillToS3URI})
			req.SpillToS3URI = valueString
		}
	}

	return req, nil
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

func convertFromAPIServiceConfigReq(apiModel *models.VwSummarySupportedAuthMethods) types.Object {
	if apiModel == nil {
		return types.ObjectNull(map[string]attr.Type{
			"enable_sso": types.BoolType,
		})
	}

	attributeTypes := map[string]attr.Type{
		"enable_sso": types.BoolType,
	}

	attributeValues := map[string]attr.Value{}

	// Handle nil pointer safely
	if apiModel.Sso != nil {
		boolValue, err := ExtractBoolFromAttribute(context.Background(), types.BoolValue(*apiModel.Sso))
		if err != nil {
			fmt.Printf("Error extracting bool: %v\n", err)
		} else {
			attributeValues["enable_sso"] = types.BoolValue(boolValue)
			fmt.Printf("Assigned value to EnableSSO: %t\n", boolValue)
		}
	} else {
		attributeValues["enable_sso"] = types.BoolNull()
	}

	ret, _ := types.ObjectValue(attributeTypes, attributeValues)
	return ret
}

func convertToAPIQueryIsolationOptions(model types.Object, ctx context.Context) *models.QueryIsolationOptionsRequest {
	attributes := model.Attributes()

	req := &models.QueryIsolationOptionsRequest{}

	if val, ok := attributes["max_nodes_per_query"]; ok {
		if value, err := val.(basetypes.Int32Value).ToInt32Value(ctx); err == nil {
			req.MaxNodesPerQuery = value.ValueInt32()
		}
	}

	if val, ok := attributes["max_queries"]; ok {
		if value, err := val.(basetypes.Int32Value).ToInt32Value(ctx); err == nil {
			req.MaxQueries = value.ValueInt32()
		}
	}

	return req
}

func convertFromAPIQueryIsolationOptions(apiModel *models.QueryIsolationOptionsResponse) types.Object {

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

func convertToAPIAutoscaling(model types.Object, ctx context.Context) (*models.AutoscalingOptionsCreateRequest, error) {
	attributes := model.Attributes()

	req := &models.AutoscalingOptionsCreateRequest{}

	if val, ok := attributes["auto_suspend_timeout_seconds"]; ok {
		intValue, err := val.(types.Int32).ToInt32Value(ctx)

		if err != nil {
			tflog.Error(ctx, "Error extracting int32 for auto_suspend_timeout_seconds", map[string]interface{}{"error": err})
			return req, fmt.Errorf("error extracting int32 for auto_suspend_timeout_seconds: %v", err)
		} else {
			req.AutoSuspendTimeoutSeconds = intValue.ValueInt32()
		}
	}

	if val, ok := attributes["disable_auto_suspend"]; ok {
		boolValue, err := ExtractBoolFromAttribute(context.Background(), val)
		if err != nil {
			tflog.Debug(ctx, "Error extracting bool", map[string]interface{}{"error": err})
			return req, fmt.Errorf("error extracting bool: %w", err)
		} else {
			req.DisableAutoSuspend = boolValue
			tflog.Debug(ctx, "Assigned value to DisableAutoSuspend", map[string]interface{}{"value": req.DisableAutoSuspend})
		}
	}

	if val, ok := attributes["scale_down_delay_seconds"]; ok {
		intValue, err := val.(types.Int32).ToInt32Value(ctx)
		if err != nil {
			tflog.Error(ctx, "Error extracting int32 for scale_down_delay_seconds", map[string]interface{}{"error": err})
			return req, fmt.Errorf("error extracting int32 for scale_down_delay_seconds: %v", err)
		} else {
			req.ImpalaScaleDownDelaySeconds = intValue.ValueInt32()
			tflog.Debug(ctx, "Assigned value to ImpalaScaleDownDelaySeconds", map[string]interface{}{"value": req.ImpalaScaleDownDelaySeconds})
		}
	}

	if val, ok := attributes["scale_up_delay_seconds"]; ok {
		intValue, err := (val.(types.Int32)).ToInt32Value(context.Background())
		if err != nil {
			tflog.Error(ctx, "Error extracting int32 for scale_up_delay_seconds", map[string]interface{}{"error": err})
			return req, fmt.Errorf("error extracting int32 for scale_up_delay_seconds: %v", err)
		} else {
			req.ImpalaScaleUpDelaySeconds = intValue.ValueInt32()
			tflog.Debug(ctx, "Assigned value to ImpalaScaleUpDelaySeconds", map[string]interface{}{"value": req.ImpalaScaleUpDelaySeconds})
		}
	}

	if val, ok := attributes["max_clusters"]; ok {
		intValue, err := (val.(types.Int32)).ToInt32Value(context.Background())
		if err != nil {
			tflog.Error(ctx, "Error extracting int32 for max_clusters", map[string]interface{}{"error": err})
			return req, fmt.Errorf("error extracting int32 for max_clusters: %v", err)

		} else {
			req.MaxClusters = intValue.ValueInt32Pointer()
			tflog.Debug(ctx, "Assigned value to MaxClusters", map[string]interface{}{"value": *req.MaxClusters})
		}
	}

	if val, ok := attributes["min_clusters"]; ok {
		intValue, err := val.(types.Int32).ToInt32Value(ctx)
		if err != nil {
			tflog.Error(ctx, "Error extracting int32 for min_clusters", map[string]interface{}{"error": err})
			return req, fmt.Errorf("error extracting int32 for min_clusters: %v", err)
		} else {
			req.MinClusters = intValue.ValueInt32Pointer()
			tflog.Debug(ctx, "Assigned value to MinClusters", map[string]interface{}{"value": *req.MinClusters})
		}
	}
	return req, nil
}

func convertFromAPIAutoscaling(apiModel *models.AutoscalingOptionsResponse) types.Object {
	if apiModel == nil {
		return types.ObjectNull(map[string]attr.Type{
			"auto_suspend_timeout_seconds": types.Int32Type,
			"disable_auto_suspend":         types.BoolType,
			"scale_down_delay_seconds":     types.Int32Type,
			"scale_up_delay_seconds":       types.Int32Type,
			"max_clusters":                 types.Int32Type,
			"min_clusters":                 types.Int32Type,
		})
	}

	attributeTypes := map[string]attr.Type{
		"auto_suspend_timeout_seconds": types.Int32Type,
		"disable_auto_suspend":         types.BoolType,
		"scale_down_delay_seconds":     types.Int32Type,
		"scale_up_delay_seconds":       types.Int32Type,
		"max_clusters":                 types.Int32Type,
		"min_clusters":                 types.Int32Type,
	}

	attributeValues := map[string]attr.Value{
		"auto_suspend_timeout_seconds": types.Int32Null(),
		"disable_auto_suspend":         types.BoolNull(),
		"scale_down_delay_seconds":     types.Int32Null(),
		"scale_up_delay_seconds":       types.Int32Null(),
		"max_clusters":                 types.Int32Null(),
		"min_clusters":                 types.Int32Null(),
	}

	if apiModel.AutoSuspendTimeoutSeconds != 0 {
		attributeValues["auto_suspend_timeout_seconds"] = types.Int32Value(apiModel.AutoSuspendTimeoutSeconds)
	}
	attributeValues["disable_auto_suspend"] = types.BoolValue(apiModel.DisableAutoSuspend)
	if apiModel.ImpalaScaleDownDelaySeconds != 0 {
		attributeValues["scale_down_delay_seconds"] = types.Int32Value(apiModel.ImpalaScaleDownDelaySeconds)
	}
	if apiModel.ImpalaScaleUpDelaySeconds != 0 {
		attributeValues["scale_up_delay_seconds"] = types.Int32Value(apiModel.ImpalaScaleUpDelaySeconds)
	}
	if apiModel.MaxClusters != 0 {
		attributeValues["max_clusters"] = types.Int32Value(apiModel.MaxClusters)
	}
	if apiModel.MinClusters != 0 {
		attributeValues["min_clusters"] = types.Int32Value(apiModel.MinClusters)
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

	if len(apiTags) == 0 {
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
