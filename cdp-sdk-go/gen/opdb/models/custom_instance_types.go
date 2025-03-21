// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomInstanceTypes Custom VM Instance Types.
//
// swagger:model CustomInstanceTypes
type CustomInstanceTypes struct {

	// Compute VM Instance Type.
	ComputeType string `json:"computeType,omitempty"`

	// Edge VM Instance Type.
	EdgeType string `json:"edgeType,omitempty"`

	// Gateway VM Instance Type.
	GatewayType string `json:"gatewayType,omitempty"`

	// Leader VM Instance Type.
	LeaderType string `json:"leaderType,omitempty"`

	// Master VM Instance Type.
	MasterType string `json:"masterType,omitempty"`

	// Worker VM Instance Type.
	WorkerType string `json:"workerType,omitempty"`
}

// Validate validates this custom instance types
func (m *CustomInstanceTypes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this custom instance types based on context it is used
func (m *CustomInstanceTypes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomInstanceTypes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomInstanceTypes) UnmarshalBinary(b []byte) error {
	var res CustomInstanceTypes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
