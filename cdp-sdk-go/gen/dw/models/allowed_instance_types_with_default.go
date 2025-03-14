// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AllowedInstanceTypesWithDefault Allowed Virtual Warehouse compute instance types and their defaults.
//
// swagger:model AllowedInstanceTypesWithDefault
type AllowedInstanceTypesWithDefault struct {

	// Allowed values for the compute instance type usage. One of these values can be used in the 'create-vw' command's 'instance-type' field .
	Allowed []string `json:"allowed"`

	// Default value for the compute instance type usage. This instance type will be used in the 'create-vw' command's 'instance-type' field in case it has to be customized. The default value also depends on the cloud platform of the Cluster (AWS/Azure).
	Default string `json:"default,omitempty"`
}

// Validate validates this allowed instance types with default
func (m *AllowedInstanceTypesWithDefault) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this allowed instance types with default based on context it is used
func (m *AllowedInstanceTypesWithDefault) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AllowedInstanceTypesWithDefault) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AllowedInstanceTypesWithDefault) UnmarshalBinary(b []byte) error {
	var res AllowedInstanceTypesWithDefault
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
