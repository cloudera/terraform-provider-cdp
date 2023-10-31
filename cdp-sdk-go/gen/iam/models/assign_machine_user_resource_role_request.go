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

// AssignMachineUserResourceRoleRequest Request object for an assign machine user resource role request.
//
// swagger:model AssignMachineUserResourceRoleRequest
type AssignMachineUserResourceRoleRequest struct {

	// The machine user to assign the resource role to. Can be the machine user's name or CRN.
	// Required: true
	MachineUserName *string `json:"machineUserName"`

	// The resource for which the resource role rights are granted.
	// Required: true
	ResourceCrn *string `json:"resourceCrn"`

	// The CRN of the resource role to assign to the machine user.
	// Required: true
	ResourceRoleCrn *string `json:"resourceRoleCrn"`
}

// Validate validates this assign machine user resource role request
func (m *AssignMachineUserResourceRoleRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMachineUserName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceRoleCrn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssignMachineUserResourceRoleRequest) validateMachineUserName(formats strfmt.Registry) error {

	if err := validate.Required("machineUserName", "body", m.MachineUserName); err != nil {
		return err
	}

	return nil
}

func (m *AssignMachineUserResourceRoleRequest) validateResourceCrn(formats strfmt.Registry) error {

	if err := validate.Required("resourceCrn", "body", m.ResourceCrn); err != nil {
		return err
	}

	return nil
}

func (m *AssignMachineUserResourceRoleRequest) validateResourceRoleCrn(formats strfmt.Registry) error {

	if err := validate.Required("resourceRoleCrn", "body", m.ResourceRoleCrn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this assign machine user resource role request based on context it is used
func (m *AssignMachineUserResourceRoleRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AssignMachineUserResourceRoleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssignMachineUserResourceRoleRequest) UnmarshalBinary(b []byte) error {
	var res AssignMachineUserResourceRoleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}