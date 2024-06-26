// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpgradeModelRegistryResponse Response for upgrading model registry.
//
// swagger:model UpgradeModelRegistryResponse
type UpgradeModelRegistryResponse struct {

	// The CRN of the model registry after upgrade.
	Crn string `json:"crn,omitempty"`
}

// Validate validates this upgrade model registry response
func (m *UpgradeModelRegistryResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this upgrade model registry response based on context it is used
func (m *UpgradeModelRegistryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpgradeModelRegistryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpgradeModelRegistryResponse) UnmarshalBinary(b []byte) error {
	var res UpgradeModelRegistryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
