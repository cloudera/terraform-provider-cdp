// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetK8sCertPEMResponse The response object for the getK8sCertPEM method.
//
// swagger:model GetK8sCertPEMResponse
type GetK8sCertPEMResponse struct {

	// The Kubernetes certificate in PEM format.
	Pem string `json:"pem,omitempty"`
}

// Validate validates this get k8s cert p e m response
func (m *GetK8sCertPEMResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get k8s cert p e m response based on context it is used
func (m *GetK8sCertPEMResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetK8sCertPEMResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetK8sCertPEMResponse) UnmarshalBinary(b []byte) error {
	var res GetK8sCertPEMResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
