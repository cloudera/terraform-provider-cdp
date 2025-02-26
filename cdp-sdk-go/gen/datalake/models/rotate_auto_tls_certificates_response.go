// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RotateAutoTLSCertificatesResponse Response object to rotate autotls certificates on datalake's hosts, deprecated.
//
// swagger:model RotateAutoTlsCertificatesResponse
type RotateAutoTLSCertificatesResponse struct {

	// Unique operation ID assigned to this command execution. Use this identifier with 'get-operation' to track status and retrieve detailed results.
	OperationID string `json:"operationId,omitempty"`
}

// Validate validates this rotate auto Tls certificates response
func (m *RotateAutoTLSCertificatesResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this rotate auto Tls certificates response based on context it is used
func (m *RotateAutoTLSCertificatesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RotateAutoTLSCertificatesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RotateAutoTLSCertificatesResponse) UnmarshalBinary(b []byte) error {
	var res RotateAutoTLSCertificatesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
