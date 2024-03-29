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

// UnassignGroupResourceRoleRequest Request object for an unassign group resource role request.
//
// swagger:model UnassignGroupResourceRoleRequest
type UnassignGroupResourceRoleRequest struct {

	// The group to unassign the resource role from.
	// Required: true
	GroupName *string `json:"groupName"`

	// The CRN of the resource for which the resource role rights will be unassigned.
	// Required: true
	ResourceCrn *string `json:"resourceCrn"`

	// The CRN of the resource role to unassign from the group.
	// Required: true
	ResourceRoleCrn *string `json:"resourceRoleCrn"`
}

// Validate validates this unassign group resource role request
func (m *UnassignGroupResourceRoleRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupName(formats); err != nil {
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

func (m *UnassignGroupResourceRoleRequest) validateGroupName(formats strfmt.Registry) error {

	if err := validate.Required("groupName", "body", m.GroupName); err != nil {
		return err
	}

	return nil
}

func (m *UnassignGroupResourceRoleRequest) validateResourceCrn(formats strfmt.Registry) error {

	if err := validate.Required("resourceCrn", "body", m.ResourceCrn); err != nil {
		return err
	}

	return nil
}

func (m *UnassignGroupResourceRoleRequest) validateResourceRoleCrn(formats strfmt.Registry) error {

	if err := validate.Required("resourceRoleCrn", "body", m.ResourceRoleCrn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this unassign group resource role request based on context it is used
func (m *UnassignGroupResourceRoleRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UnassignGroupResourceRoleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnassignGroupResourceRoleRequest) UnmarshalBinary(b []byte) error {
	var res UnassignGroupResourceRoleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
