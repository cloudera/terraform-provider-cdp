// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetVMTypesResponse Response object from the VM type fetch operation.
//
// swagger:model GetVmTypesResponse
type GetVMTypesResponse struct {

	// The supported VM types based on the given parameters.
	VMTypes []string `json:"vmTypes"`
}

// Validate validates this get Vm types response
func (m *GetVMTypesResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get Vm types response based on context it is used
func (m *GetVMTypesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetVMTypesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetVMTypesResponse) UnmarshalBinary(b []byte) error {
	var res GetVMTypesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
