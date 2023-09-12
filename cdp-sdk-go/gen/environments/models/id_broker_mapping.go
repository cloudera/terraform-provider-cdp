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

// IDBrokerMapping A mapping of an actor or group to a cloud provider role.
//
// swagger:model IdBrokerMapping
type IDBrokerMapping struct {

	// The CRN of the actor or group.
	// Required: true
	AccessorCrn *string `json:"accessorCrn"`

	// The cloud provider role (e.g., ARN in AWS, Resource ID in Azure) to which the actor or group is mapped.
	// Required: true
	Role *string `json:"role"`
}

// Validate validates this Id broker mapping
func (m *IDBrokerMapping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessorCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IDBrokerMapping) validateAccessorCrn(formats strfmt.Registry) error {

	if err := validate.Required("accessorCrn", "body", m.AccessorCrn); err != nil {
		return err
	}

	return nil
}

func (m *IDBrokerMapping) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Id broker mapping based on context it is used
func (m *IDBrokerMapping) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IDBrokerMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IDBrokerMapping) UnmarshalBinary(b []byte) error {
	var res IDBrokerMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}