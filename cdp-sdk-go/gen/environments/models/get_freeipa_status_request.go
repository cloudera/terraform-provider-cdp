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

// GetFreeipaStatusRequest Request object for getting the status of the FreeIPA servers.
//
// swagger:model GetFreeipaStatusRequest
type GetFreeipaStatusRequest struct {

	// The environment name or CRN of the FreeIPA to repair
	// Required: true
	EnvironmentName *string `json:"environmentName"`
}

// Validate validates this get freeipa status request
func (m *GetFreeipaStatusRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironmentName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetFreeipaStatusRequest) validateEnvironmentName(formats strfmt.Registry) error {

	if err := validate.Required("environmentName", "body", m.EnvironmentName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get freeipa status request based on context it is used
func (m *GetFreeipaStatusRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetFreeipaStatusRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetFreeipaStatusRequest) UnmarshalBinary(b []byte) error {
	var res GetFreeipaStatusRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
