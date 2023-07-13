// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImpalaHASettingsCreateRequest High Availability settings for the Impala Virtual Warehouse. The values are disregarded for Hive.
//
// swagger:model ImpalaHASettingsCreateRequest
type ImpalaHASettingsCreateRequest struct {

	// Enables a backup instance for Impala catalog to ensure high availability.
	EnableCatalogHighAvailability bool `json:"enableCatalogHighAvailability,omitempty"`

	// Enables a shutdown of the coordinator. If Unified Analytics is enabled, then this setting is explicitly disabled (ignored) and should not be provided.
	EnableShutdownOfCoordinator bool `json:"enableShutdownOfCoordinator,omitempty"`

	// Set High Availability mode. If not provided, the default will apply. DISABLED - Disables Impala coordinator and Database Catalog high availability. ACTIVE_PASSIVE - Runs multiple coordinators (one active, one passive) and Database Catalogs (one active, one passive). ACTIVE_ACTIVE - Runs multiple coordinators (both active) and Database Catalogs (one active, one passive). If Unified Analytics is enabled, then this cannot be set to ACTIVE_ACTIVE.
	HighAvailabilityMode ImpalaHighAvailabilityMode `json:"highAvailabilityMode,omitempty"`

	// The number of active coordinators.
	NumOfActiveCoordinators int32 `json:"numOfActiveCoordinators,omitempty"`

	// Delay in seconds before the shutdown of coordinator event happens.
	ShutdownOfCoordinatorDelaySeconds int32 `json:"shutdownOfCoordinatorDelaySeconds,omitempty"`
}

// Validate validates this impala h a settings create request
func (m *ImpalaHASettingsCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHighAvailabilityMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImpalaHASettingsCreateRequest) validateHighAvailabilityMode(formats strfmt.Registry) error {
	if swag.IsZero(m.HighAvailabilityMode) { // not required
		return nil
	}

	if err := m.HighAvailabilityMode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("highAvailabilityMode")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("highAvailabilityMode")
		}
		return err
	}

	return nil
}

// ContextValidate validate this impala h a settings create request based on the context it is used
func (m *ImpalaHASettingsCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHighAvailabilityMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImpalaHASettingsCreateRequest) contextValidateHighAvailabilityMode(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.HighAvailabilityMode) { // not required
		return nil
	}

	if err := m.HighAvailabilityMode.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("highAvailabilityMode")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("highAvailabilityMode")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImpalaHASettingsCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImpalaHASettingsCreateRequest) UnmarshalBinary(b []byte) error {
	var res ImpalaHASettingsCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
