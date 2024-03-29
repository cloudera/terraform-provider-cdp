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

// ImpalaExecutorGroupSetsUpdateRequest Re-configure executor group sets for workload aware autoscaling.
//
// swagger:model ImpalaExecutorGroupSetsUpdateRequest
type ImpalaExecutorGroupSetsUpdateRequest struct {

	// Re-configure first optional custom executor group set for workload aware autoscaling.
	Custom1 *ImpalaExecutorGroupSetUpdateRequest `json:"custom1,omitempty"`

	// Re-configure second optional custom executor group set for workload aware autoscaling.
	Custom2 *ImpalaExecutorGroupSetUpdateRequest `json:"custom2,omitempty"`

	// Re-configure third optional custom executor group set for workload aware autoscaling.
	Custom3 *ImpalaExecutorGroupSetUpdateRequest `json:"custom3,omitempty"`

	// Re-configure large executor group set for workload aware autoscaling.
	Large *ImpalaExecutorGroupSetUpdateRequest `json:"large,omitempty"`

	// Re-configure small executor group set for workload aware autoscaling.
	Small *ImpalaExecutorGroupSetUpdateRequest `json:"small,omitempty"`
}

// Validate validates this impala executor group sets update request
func (m *ImpalaExecutorGroupSetsUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustom1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustom2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustom3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLarge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSmall(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImpalaExecutorGroupSetsUpdateRequest) validateCustom1(formats strfmt.Registry) error {
	if swag.IsZero(m.Custom1) { // not required
		return nil
	}

	if m.Custom1 != nil {
		if err := m.Custom1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("custom1")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("custom1")
			}
			return err
		}
	}

	return nil
}

func (m *ImpalaExecutorGroupSetsUpdateRequest) validateCustom2(formats strfmt.Registry) error {
	if swag.IsZero(m.Custom2) { // not required
		return nil
	}

	if m.Custom2 != nil {
		if err := m.Custom2.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("custom2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("custom2")
			}
			return err
		}
	}

	return nil
}

func (m *ImpalaExecutorGroupSetsUpdateRequest) validateCustom3(formats strfmt.Registry) error {
	if swag.IsZero(m.Custom3) { // not required
		return nil
	}

	if m.Custom3 != nil {
		if err := m.Custom3.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("custom3")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("custom3")
			}
			return err
		}
	}

	return nil
}

func (m *ImpalaExecutorGroupSetsUpdateRequest) validateLarge(formats strfmt.Registry) error {
	if swag.IsZero(m.Large) { // not required
		return nil
	}

	if m.Large != nil {
		if err := m.Large.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("large")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("large")
			}
			return err
		}
	}

	return nil
}

func (m *ImpalaExecutorGroupSetsUpdateRequest) validateSmall(formats strfmt.Registry) error {
	if swag.IsZero(m.Small) { // not required
		return nil
	}

	if m.Small != nil {
		if err := m.Small.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("small")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("small")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this impala executor group sets update request based on the context it is used
func (m *ImpalaExecutorGroupSetsUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustom1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustom2(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustom3(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLarge(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSmall(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImpalaExecutorGroupSetsUpdateRequest) contextValidateCustom1(ctx context.Context, formats strfmt.Registry) error {

	if m.Custom1 != nil {

		if swag.IsZero(m.Custom1) { // not required
			return nil
		}

		if err := m.Custom1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("custom1")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("custom1")
			}
			return err
		}
	}

	return nil
}

func (m *ImpalaExecutorGroupSetsUpdateRequest) contextValidateCustom2(ctx context.Context, formats strfmt.Registry) error {

	if m.Custom2 != nil {

		if swag.IsZero(m.Custom2) { // not required
			return nil
		}

		if err := m.Custom2.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("custom2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("custom2")
			}
			return err
		}
	}

	return nil
}

func (m *ImpalaExecutorGroupSetsUpdateRequest) contextValidateCustom3(ctx context.Context, formats strfmt.Registry) error {

	if m.Custom3 != nil {

		if swag.IsZero(m.Custom3) { // not required
			return nil
		}

		if err := m.Custom3.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("custom3")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("custom3")
			}
			return err
		}
	}

	return nil
}

func (m *ImpalaExecutorGroupSetsUpdateRequest) contextValidateLarge(ctx context.Context, formats strfmt.Registry) error {

	if m.Large != nil {

		if swag.IsZero(m.Large) { // not required
			return nil
		}

		if err := m.Large.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("large")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("large")
			}
			return err
		}
	}

	return nil
}

func (m *ImpalaExecutorGroupSetsUpdateRequest) contextValidateSmall(ctx context.Context, formats strfmt.Registry) error {

	if m.Small != nil {

		if swag.IsZero(m.Small) { // not required
			return nil
		}

		if err := m.Small.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("small")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("small")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImpalaExecutorGroupSetsUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImpalaExecutorGroupSetsUpdateRequest) UnmarshalBinary(b []byte) error {
	var res ImpalaExecutorGroupSetsUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
