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

// CreateGCPCredentialRequest Request object for a create GCP credential request.
//
// swagger:model CreateGCPCredentialRequest
type CreateGCPCredentialRequest struct {

	// The JSON key for the service account. Please use local path when using the CLI (e.g. file:///absolute/path/to/cred.json) to avoid exposing the keys in the command line history.
	// Required: true
	CredentialKey *string `json:"credentialKey"`

	// The name of the credential.
	// Required: true
	CredentialName *string `json:"credentialName"`

	// A description for the credential.
	Description string `json:"description,omitempty"`
}

// Validate validates this create g c p credential request
func (m *CreateGCPCredentialRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateGCPCredentialRequest) validateCredentialKey(formats strfmt.Registry) error {

	if err := validate.Required("credentialKey", "body", m.CredentialKey); err != nil {
		return err
	}

	return nil
}

func (m *CreateGCPCredentialRequest) validateCredentialName(formats strfmt.Registry) error {

	if err := validate.Required("credentialName", "body", m.CredentialName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create g c p credential request based on context it is used
func (m *CreateGCPCredentialRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateGCPCredentialRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateGCPCredentialRequest) UnmarshalBinary(b []byte) error {
	var res CreateGCPCredentialRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
