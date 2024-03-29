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

// CollectDiagnosticsResponse Information about diagnostic bundle generation
//
// swagger:model CollectDiagnosticsResponse
type CollectDiagnosticsResponse struct {

	// Details of requested diagnostic bundle collection
	DiagnosticsBundle *DiagnosticsBundle `json:"diagnosticsBundle,omitempty"`
}

// Validate validates this collect diagnostics response
func (m *CollectDiagnosticsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiagnosticsBundle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CollectDiagnosticsResponse) validateDiagnosticsBundle(formats strfmt.Registry) error {
	if swag.IsZero(m.DiagnosticsBundle) { // not required
		return nil
	}

	if m.DiagnosticsBundle != nil {
		if err := m.DiagnosticsBundle.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("diagnosticsBundle")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("diagnosticsBundle")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this collect diagnostics response based on the context it is used
func (m *CollectDiagnosticsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDiagnosticsBundle(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CollectDiagnosticsResponse) contextValidateDiagnosticsBundle(ctx context.Context, formats strfmt.Registry) error {

	if m.DiagnosticsBundle != nil {

		if swag.IsZero(m.DiagnosticsBundle) { // not required
			return nil
		}

		if err := m.DiagnosticsBundle.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("diagnosticsBundle")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("diagnosticsBundle")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CollectDiagnosticsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CollectDiagnosticsResponse) UnmarshalBinary(b []byte) error {
	var res CollectDiagnosticsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
