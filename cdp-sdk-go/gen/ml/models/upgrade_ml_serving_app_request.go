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

// UpgradeMlServingAppRequest Request object for the UpgradeMlServingApp method.
//
// swagger:model UpgradeMlServingAppRequest
type UpgradeMlServingAppRequest struct {

	// The serving app CRN.
	// Required: true
	AppCrn *string `json:"appCrn"`
}

// Validate validates this upgrade ml serving app request
func (m *UpgradeMlServingAppRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppCrn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpgradeMlServingAppRequest) validateAppCrn(formats strfmt.Registry) error {

	if err := validate.Required("appCrn", "body", m.AppCrn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this upgrade ml serving app request based on context it is used
func (m *UpgradeMlServingAppRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpgradeMlServingAppRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpgradeMlServingAppRequest) UnmarshalBinary(b []byte) error {
	var res UpgradeMlServingAppRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
