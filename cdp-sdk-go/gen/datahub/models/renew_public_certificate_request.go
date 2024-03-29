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

// RenewPublicCertificateRequest Request object for renewing the Datahub public certificate.
//
// swagger:model RenewPublicCertificateRequest
type RenewPublicCertificateRequest struct {

	// The name or CRN of the datahub.
	// Required: true
	Datahub *string `json:"datahub"`
}

// Validate validates this renew public certificate request
func (m *RenewPublicCertificateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatahub(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RenewPublicCertificateRequest) validateDatahub(formats strfmt.Registry) error {

	if err := validate.Required("datahub", "body", m.Datahub); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this renew public certificate request based on context it is used
func (m *RenewPublicCertificateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RenewPublicCertificateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RenewPublicCertificateRequest) UnmarshalBinary(b []byte) error {
	var res RenewPublicCertificateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
