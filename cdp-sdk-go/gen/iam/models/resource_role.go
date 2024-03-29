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

// ResourceRole Information about a resource role. A resource role is a role that grants a collection of rights to a user on resources.
//
// swagger:model ResourceRole
type ResourceRole struct {

	// The CRN of the resource role.
	// Required: true
	Crn *string `json:"crn"`

	// The rights granted by this role.
	// Required: true
	Rights []string `json:"rights"`
}

// Validate validates this resource role
func (m *ResourceRole) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRights(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceRole) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

func (m *ResourceRole) validateRights(formats strfmt.Registry) error {

	if err := validate.Required("rights", "body", m.Rights); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this resource role based on context it is used
func (m *ResourceRole) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResourceRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceRole) UnmarshalBinary(b []byte) error {
	var res ResourceRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
