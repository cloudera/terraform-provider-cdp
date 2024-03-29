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

// ListDatahubDiagnosticsRequest Request object for listing recent Datahub diagnostics collections.
//
// swagger:model ListDatahubDiagnosticsRequest
type ListDatahubDiagnosticsRequest struct {

	// CRN of the DataHub cluster.
	// Required: true
	Crn *string `json:"crn"`
}

// Validate validates this list datahub diagnostics request
func (m *ListDatahubDiagnosticsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListDatahubDiagnosticsRequest) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list datahub diagnostics request based on context it is used
func (m *ListDatahubDiagnosticsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListDatahubDiagnosticsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListDatahubDiagnosticsRequest) UnmarshalBinary(b []byte) error {
	var res ListDatahubDiagnosticsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
