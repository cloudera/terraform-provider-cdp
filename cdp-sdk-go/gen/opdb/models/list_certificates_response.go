// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListCertificatesResponse The response of listing fingerprints of certificates in Global Trust Store
//
// swagger:model ListCertificatesResponse
type ListCertificatesResponse struct {

	// List of certificate SHA-1 fingerprints in Global Trust Store
	Fingerprints []string `json:"fingerprints"`
}

// Validate validates this list certificates response
func (m *ListCertificatesResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list certificates response based on context it is used
func (m *ListCertificatesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListCertificatesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListCertificatesResponse) UnmarshalBinary(b []byte) error {
	var res ListCertificatesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
