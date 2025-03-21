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

// RotateSecretsRequest Request object for starting secret rotation for datalake.
//
// swagger:model RotateSecretsRequest
type RotateSecretsRequest struct {

	// The datalake name or CRN where we wish to rotate secrets.
	// Required: true
	Datalake *string `json:"datalake"`

	// The list of secrets that need replacement.
	// Required: true
	SecretTypes []string `json:"secretTypes"`
}

// Validate validates this rotate secrets request
func (m *RotateSecretsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatalake(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RotateSecretsRequest) validateDatalake(formats strfmt.Registry) error {

	if err := validate.Required("datalake", "body", m.Datalake); err != nil {
		return err
	}

	return nil
}

func (m *RotateSecretsRequest) validateSecretTypes(formats strfmt.Registry) error {

	if err := validate.Required("secretTypes", "body", m.SecretTypes); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rotate secrets request based on context it is used
func (m *RotateSecretsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RotateSecretsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RotateSecretsRequest) UnmarshalBinary(b []byte) error {
	var res RotateSecretsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
