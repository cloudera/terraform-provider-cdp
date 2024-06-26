// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteVcResponse Response object for DeleteVc method.
//
// swagger:model DeleteVcResponse
type DeleteVcResponse struct {

	// status of virtual cluster deletion.
	Status string `json:"status,omitempty"`
}

// Validate validates this delete vc response
func (m *DeleteVcResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete vc response based on context it is used
func (m *DeleteVcResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteVcResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteVcResponse) UnmarshalBinary(b []byte) error {
	var res DeleteVcResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
