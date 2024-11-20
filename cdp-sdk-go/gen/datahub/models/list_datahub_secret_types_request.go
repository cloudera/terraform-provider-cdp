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

// ListDatahubSecretTypesRequest Request for listing possible secret values for Datahub.
//
// swagger:model ListDatahubSecretTypesRequest
type ListDatahubSecretTypesRequest struct {

	// The Datahub CRN where we wish to get the rotatable secrets.
	// Required: true
	Datahub *string `json:"datahub"`
}

// Validate validates this list datahub secret types request
func (m *ListDatahubSecretTypesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatahub(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListDatahubSecretTypesRequest) validateDatahub(formats strfmt.Registry) error {

	if err := validate.Required("datahub", "body", m.Datahub); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list datahub secret types request based on context it is used
func (m *ListDatahubSecretTypesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListDatahubSecretTypesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListDatahubSecretTypesRequest) UnmarshalBinary(b []byte) error {
	var res ListDatahubSecretTypesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
