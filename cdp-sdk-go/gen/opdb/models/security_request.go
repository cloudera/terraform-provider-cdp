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

// SecurityRequest Security related configurations for the clusters.
//
// swagger:model SecurityRequest
type SecurityRequest struct {

	// SELinux enforcement policy, can be PERMISSIVE or ENFORCING
	SeLinux SELinux `json:"seLinux,omitempty"`
}

// Validate validates this security request
func (m *SecurityRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSeLinux(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityRequest) validateSeLinux(formats strfmt.Registry) error {
	if swag.IsZero(m.SeLinux) { // not required
		return nil
	}

	if err := m.SeLinux.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("seLinux")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("seLinux")
		}
		return err
	}

	return nil
}

// ContextValidate validate this security request based on the context it is used
func (m *SecurityRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSeLinux(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityRequest) contextValidateSeLinux(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.SeLinux) { // not required
		return nil
	}

	if err := m.SeLinux.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("seLinux")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("seLinux")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityRequest) UnmarshalBinary(b []byte) error {
	var res SecurityRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
