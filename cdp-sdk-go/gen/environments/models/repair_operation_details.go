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

// RepairOperationDetails Details object of the repair operation for success or falure.
//
// swagger:model RepairOperationDetails
type RepairOperationDetails struct {

	// environment crn.
	// Required: true
	EnvironmentCrn *string `json:"environmentCrn"`

	// The detail of the success or failure.
	Message string `json:"message,omitempty"`
}

// Validate validates this repair operation details
func (m *RepairOperationDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironmentCrn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RepairOperationDetails) validateEnvironmentCrn(formats strfmt.Registry) error {

	if err := validate.Required("environmentCrn", "body", m.EnvironmentCrn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this repair operation details based on context it is used
func (m *RepairOperationDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RepairOperationDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RepairOperationDetails) UnmarshalBinary(b []byte) error {
	var res RepairOperationDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}