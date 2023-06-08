// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteSamlProviderRequest Request object for deleting SAML provider request.
//
// swagger:model DeleteSamlProviderRequest
type DeleteSamlProviderRequest struct {

	// The name or CRN of the SAML provider to delete.
	SamlProviderName string `json:"samlProviderName,omitempty"`
}

// Validate validates this delete saml provider request
func (m *DeleteSamlProviderRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete saml provider request based on context it is used
func (m *DeleteSamlProviderRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteSamlProviderRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteSamlProviderRequest) UnmarshalBinary(b []byte) error {
	var res DeleteSamlProviderRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
