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

// CreateMachineUserRequest Request object for create machine user request.
//
// swagger:model CreateMachineUserRequest
type CreateMachineUserRequest struct {

	// The name to use for the new machine user. The name must be an alpha numeric string, including '-' and '_', cannot start with '__' (double underscore) and cannot be longer than 128 characters. Only one machine user with this name can exist in an account at a given time.
	// Required: true
	MachineUserName *string `json:"machineUserName"`
}

// Validate validates this create machine user request
func (m *CreateMachineUserRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMachineUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateMachineUserRequest) validateMachineUserName(formats strfmt.Registry) error {

	if err := validate.Required("machineUserName", "body", m.MachineUserName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create machine user request based on context it is used
func (m *CreateMachineUserRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateMachineUserRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateMachineUserRequest) UnmarshalBinary(b []byte) error {
	var res CreateMachineUserRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
