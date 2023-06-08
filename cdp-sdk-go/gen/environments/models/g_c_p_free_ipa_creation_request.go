// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GCPFreeIpaCreationRequest Request object for creating FreeIPA in the environment.
//
// swagger:model GCPFreeIpaCreationRequest
type GCPFreeIpaCreationRequest struct {

	// The number of FreeIPA instances to create per group when creating FreeIPA in the environment
	InstanceCountByGroup int32 `json:"instanceCountByGroup,omitempty"`

	// Custom instance type of FreeIPA instances.
	InstanceType string `json:"instanceType,omitempty"`

	// The recipes for the FreeIPA cluster.
	Recipes []string `json:"recipes"`
}

// Validate validates this g c p free ipa creation request
func (m *GCPFreeIpaCreationRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this g c p free ipa creation request based on context it is used
func (m *GCPFreeIpaCreationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GCPFreeIpaCreationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GCPFreeIpaCreationRequest) UnmarshalBinary(b []byte) error {
	var res GCPFreeIpaCreationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
