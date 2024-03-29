// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetRootCertificateResponse Response object with base64 encoded contents of the public certificate for an environment.
//
// swagger:model GetRootCertificateResponse
type GetRootCertificateResponse struct {

	// Contents of a certificate.
	Contents string `json:"contents,omitempty"`
}

// Validate validates this get root certificate response
func (m *GetRootCertificateResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get root certificate response based on context it is used
func (m *GetRootCertificateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetRootCertificateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetRootCertificateResponse) UnmarshalBinary(b []byte) error {
	var res GetRootCertificateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
