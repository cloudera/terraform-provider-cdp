// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AutoscalingOptionsCreateRequest Auto-scaling configuration for a Virtual Warehouse.
//
// swagger:model AutoscalingOptionsCreateRequest
type AutoscalingOptionsCreateRequest struct {

	// Auto suspend threshold for Virtual Warehouse.
	AutoSuspendTimeoutSeconds int32 `json:"autoSuspendTimeoutSeconds,omitempty"`

	// Turn off auto suspend for Virtual Warehouse.
	DisableAutoSuspend bool `json:"disableAutoSuspend,omitempty"`

	// DEPRECATED in favor of the top level enableUnifiedAnalytics flag. Enable Unified Analytics. In case of Hive Virtual Warehouses this cannot be provided, because this value is inferred. In case of Impala this can be set. Passing --query-isolation-options will be considered if this flag is set to true. If Unified Analytics enabled then the "impalaEnableShutdownOfCoordinator" explicitly disabled and should not be provided, furthermore the "impalaHighAvailabilityMode" cannot be set to ACTIVE_ACTIVE.
	EnableUnifiedAnalytics bool `json:"enableUnifiedAnalytics,omitempty"`

	// Set Desired free capacity. Either "hiveScaleWaitTimeSeconds" or "hiveDesiredFreeCapacity" can be provided.
	HiveDesiredFreeCapacity int32 `json:"hiveDesiredFreeCapacity,omitempty"`

	// Set wait time before a scale event happens. Either "hiveScaleWaitTimeSeconds" or "hiveDesiredFreeCapacity" can be provided.
	HiveScaleWaitTimeSeconds int32 `json:"hiveScaleWaitTimeSeconds,omitempty"`

	// DEPRECATED in favor of the top level impalaHASettings object. Enables a backup instance for Impala catalog to ensure high availability.
	ImpalaEnableCatalogHighAvailability bool `json:"impalaEnableCatalogHighAvailability,omitempty"`

	// DEPRECATED in favor of the top level impalaHASettings object. Enables a shutdown of the coordinator. If Unified Analytics enabled then this setting explicitly disabled and should not be provided.
	ImpalaEnableShutdownOfCoordinator bool `json:"impalaEnableShutdownOfCoordinator,omitempty"`

	// Configures executor group sets for workload aware autoscaling.
	ImpalaExecutorGroupSets *ImpalaExecutorGroupSetsCreateRequest `json:"impalaExecutorGroupSets,omitempty"`

	// DEPRECATED in favor of the top level impalaHASettings object. Set High Availability mode. If not provided the default will apply. This value is disregarded for Hive.
	// Enum: [ACTIVE_PASSIVE ACTIVE_ACTIVE DISABLED]
	ImpalaHighAvailabilityMode string `json:"impalaHighAvailabilityMode,omitempty"`

	// DEPRECATED in favor of the top level impalaHASettings object. Number of the active coordinators.
	ImpalaNumOfActiveCoordinators int32 `json:"impalaNumOfActiveCoordinators,omitempty"`

	// Scale down threshold in seconds. If not provided defaults will apply.
	ImpalaScaleDownDelaySeconds int32 `json:"impalaScaleDownDelaySeconds,omitempty"`

	// Scale up the scaling up threshold in seconds. If not provided defaults will apply.
	ImpalaScaleUpDelaySeconds int32 `json:"impalaScaleUpDelaySeconds,omitempty"`

	// DEPRECATED in favor of the top level impalaHASettings object. Delay in seconds before the shutdown of coordinator event happens.
	ImpalaShutdownOfCoordinatorDelaySeconds int32 `json:"impalaShutdownOfCoordinatorDelaySeconds,omitempty"`

	// Maximum number of available compute groups.
	MaxClusters *int32 `json:"maxClusters,omitempty"`

	// Minimum number of available compute groups.
	MinClusters *int32 `json:"minClusters,omitempty"`

	// Name of the pod configuration.
	PodConfigName string `json:"podConfigName,omitempty"`
}

// Validate validates this autoscaling options create request
func (m *AutoscalingOptionsCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImpalaExecutorGroupSets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImpalaHighAvailabilityMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoscalingOptionsCreateRequest) validateImpalaExecutorGroupSets(formats strfmt.Registry) error {
	if swag.IsZero(m.ImpalaExecutorGroupSets) { // not required
		return nil
	}

	if m.ImpalaExecutorGroupSets != nil {
		if err := m.ImpalaExecutorGroupSets.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("impalaExecutorGroupSets")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("impalaExecutorGroupSets")
			}
			return err
		}
	}

	return nil
}

var autoscalingOptionsCreateRequestTypeImpalaHighAvailabilityModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE_PASSIVE","ACTIVE_ACTIVE","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		autoscalingOptionsCreateRequestTypeImpalaHighAvailabilityModePropEnum = append(autoscalingOptionsCreateRequestTypeImpalaHighAvailabilityModePropEnum, v)
	}
}

const (

	// AutoscalingOptionsCreateRequestImpalaHighAvailabilityModeACTIVEPASSIVE captures enum value "ACTIVE_PASSIVE"
	AutoscalingOptionsCreateRequestImpalaHighAvailabilityModeACTIVEPASSIVE string = "ACTIVE_PASSIVE"

	// AutoscalingOptionsCreateRequestImpalaHighAvailabilityModeACTIVEACTIVE captures enum value "ACTIVE_ACTIVE"
	AutoscalingOptionsCreateRequestImpalaHighAvailabilityModeACTIVEACTIVE string = "ACTIVE_ACTIVE"

	// AutoscalingOptionsCreateRequestImpalaHighAvailabilityModeDISABLED captures enum value "DISABLED"
	AutoscalingOptionsCreateRequestImpalaHighAvailabilityModeDISABLED string = "DISABLED"
)

// prop value enum
func (m *AutoscalingOptionsCreateRequest) validateImpalaHighAvailabilityModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, autoscalingOptionsCreateRequestTypeImpalaHighAvailabilityModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AutoscalingOptionsCreateRequest) validateImpalaHighAvailabilityMode(formats strfmt.Registry) error {
	if swag.IsZero(m.ImpalaHighAvailabilityMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateImpalaHighAvailabilityModeEnum("impalaHighAvailabilityMode", "body", m.ImpalaHighAvailabilityMode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this autoscaling options create request based on the context it is used
func (m *AutoscalingOptionsCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImpalaExecutorGroupSets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoscalingOptionsCreateRequest) contextValidateImpalaExecutorGroupSets(ctx context.Context, formats strfmt.Registry) error {

	if m.ImpalaExecutorGroupSets != nil {

		if swag.IsZero(m.ImpalaExecutorGroupSets) { // not required
			return nil
		}

		if err := m.ImpalaExecutorGroupSets.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("impalaExecutorGroupSets")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("impalaExecutorGroupSets")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutoscalingOptionsCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoscalingOptionsCreateRequest) UnmarshalBinary(b []byte) error {
	var res AutoscalingOptionsCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
