// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DisengageAutoAdminResponse A response object for disengage-autoadmin request.
//
// swagger:model DisengageAutoAdminResponse
type DisengageAutoAdminResponse struct {

	// status of disengage-autoadmin request.
	WasDisengaged bool `json:"wasDisengaged,omitempty"`
}

// Validate validates this disengage auto admin response
func (m *DisengageAutoAdminResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this disengage auto admin response based on context it is used
func (m *DisengageAutoAdminResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DisengageAutoAdminResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DisengageAutoAdminResponse) UnmarshalBinary(b []byte) error {
	var res DisengageAutoAdminResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
