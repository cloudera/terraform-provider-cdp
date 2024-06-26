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

// CustomDockerRegistryRequest The desired custom docker registry for data services to be used.
//
// swagger:model CustomDockerRegistryRequest
type CustomDockerRegistryRequest struct {

	// The CRN of the desired custom docker registry for data services to be used.
	// Required: true
	Crn *string `json:"crn"`
}

// Validate validates this custom docker registry request
func (m *CustomDockerRegistryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomDockerRegistryRequest) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this custom docker registry request based on context it is used
func (m *CustomDockerRegistryRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomDockerRegistryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomDockerRegistryRequest) UnmarshalBinary(b []byte) error {
	var res CustomDockerRegistryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
