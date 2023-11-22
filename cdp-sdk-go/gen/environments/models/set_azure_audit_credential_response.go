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

// SetAzureAuditCredentialResponse Response object for a set Azure audit credential request.
//
// swagger:model SetAzureAuditCredentialResponse
type SetAzureAuditCredentialResponse struct {

	// The credential object.
	// Required: true
	Credential *Credential `json:"credential"`
}

// Validate validates this set azure audit credential response
func (m *SetAzureAuditCredentialResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredential(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetAzureAuditCredentialResponse) validateCredential(formats strfmt.Registry) error {

	if err := validate.Required("credential", "body", m.Credential); err != nil {
		return err
	}

	if m.Credential != nil {
		if err := m.Credential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credential")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credential")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this set azure audit credential response based on the context it is used
func (m *SetAzureAuditCredentialResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredential(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetAzureAuditCredentialResponse) contextValidateCredential(ctx context.Context, formats strfmt.Registry) error {

	if m.Credential != nil {

		if err := m.Credential.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credential")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credential")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SetAzureAuditCredentialResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetAzureAuditCredentialResponse) UnmarshalBinary(b []byte) error {
	var res SetAzureAuditCredentialResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}