// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GCPCredentialProperties The credential properties that closely related to those that have created on GCP.
//
// swagger:model GCPCredentialProperties
type GCPCredentialProperties struct {

	// The GCP credential key type. Json is the only supported key type.
	KeyType string `json:"keyType,omitempty"`
}

// Validate validates this g c p credential properties
func (m *GCPCredentialProperties) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this g c p credential properties based on context it is used
func (m *GCPCredentialProperties) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GCPCredentialProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GCPCredentialProperties) UnmarshalBinary(b []byte) error {
	var res GCPCredentialProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
