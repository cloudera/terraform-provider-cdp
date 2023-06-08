// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateVwRequest Request object for the updateVw method.
//
// swagger:model UpdateVwRequest
type UpdateVwRequest struct {

	// Autoscaling settings for the Virtual Warehouse.
	Autoscaling *AutoscalingOptionsUpdateRequest `json:"autoscaling,omitempty"`

	// ID of the Virtual Warehouse's cluster.
	// Required: true
	ClusterID *string `json:"clusterId"`

	// The service configuration to update the VW with. This will be applied on top of the existing configuration so there's no need to list configurations that stay the same.
	Config *ServiceConfigReq `json:"config,omitempty"`

	// High Availability settings update for the Impala Virtual Warehouse.
	ImpalaHaSettings *ImpalaHASettingsUpdateRequest `json:"impalaHaSettings,omitempty"`

	// Value of 'true' automatically configures the Virtual Warehouse to support JWTs issues by the CDP JWT token provider.  Value of 'false' does not enable JWT auth on the Virtual Warehouse.  If this field is not specified, it defaults to 'false'.
	PlatformJwtAuth *bool `json:"platformJwtAuth,omitempty"`

	// Query isolation settings for Hive Virtual Warehouses.
	QueryIsolationOptions *QueryIsolationOptionsRequest `json:"queryIsolationOptions,omitempty"`

	// ID of the Virtual Warehouse.
	// Required: true
	VwID *string `json:"vwId"`
}

// Validate validates this update vw request
func (m *UpdateVwRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoscaling(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImpalaHaSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryIsolationOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVwID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateVwRequest) validateAutoscaling(formats strfmt.Registry) error {
	if swag.IsZero(m.Autoscaling) { // not required
		return nil
	}

	if m.Autoscaling != nil {
		if err := m.Autoscaling.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoscaling")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoscaling")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateVwRequest) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("clusterId", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *UpdateVwRequest) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateVwRequest) validateImpalaHaSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.ImpalaHaSettings) { // not required
		return nil
	}

	if m.ImpalaHaSettings != nil {
		if err := m.ImpalaHaSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("impalaHaSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("impalaHaSettings")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateVwRequest) validateQueryIsolationOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.QueryIsolationOptions) { // not required
		return nil
	}

	if m.QueryIsolationOptions != nil {
		if err := m.QueryIsolationOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queryIsolationOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queryIsolationOptions")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateVwRequest) validateVwID(formats strfmt.Registry) error {

	if err := validate.Required("vwId", "body", m.VwID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update vw request based on the context it is used
func (m *UpdateVwRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutoscaling(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImpalaHaSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueryIsolationOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateVwRequest) contextValidateAutoscaling(ctx context.Context, formats strfmt.Registry) error {

	if m.Autoscaling != nil {
		if err := m.Autoscaling.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoscaling")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoscaling")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateVwRequest) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {
		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateVwRequest) contextValidateImpalaHaSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.ImpalaHaSettings != nil {
		if err := m.ImpalaHaSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("impalaHaSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("impalaHaSettings")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateVwRequest) contextValidateQueryIsolationOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.QueryIsolationOptions != nil {
		if err := m.QueryIsolationOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queryIsolationOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queryIsolationOptions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateVwRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateVwRequest) UnmarshalBinary(b []byte) error {
	var res UpdateVwRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
