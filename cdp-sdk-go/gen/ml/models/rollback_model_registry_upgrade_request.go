// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RollbackModelRegistryUpgradeRequest Request for rollback model registry upgrade.
//
// swagger:model RollbackModelRegistryUpgradeRequest
type RollbackModelRegistryUpgradeRequest struct {

	// The CRN of the model registry.
	Crn string `json:"crn,omitempty"`
}

// Validate validates this rollback model registry upgrade request
func (m *RollbackModelRegistryUpgradeRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this rollback model registry upgrade request based on context it is used
func (m *RollbackModelRegistryUpgradeRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RollbackModelRegistryUpgradeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RollbackModelRegistryUpgradeRequest) UnmarshalBinary(b []byte) error {
	var res RollbackModelRegistryUpgradeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
