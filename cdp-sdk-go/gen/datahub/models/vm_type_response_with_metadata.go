// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VMTypeResponseWithMetadata Response object containing vm types and its metadata.
//
// swagger:model VmTypeResponseWithMetadata
type VMTypeResponseWithMetadata struct {

	// Name of the vm type.
	Name string `json:"name,omitempty"`

	// JSON string with metadata as key value pairs.
	Properties string `json:"properties,omitempty"`
}

// Validate validates this Vm type response with metadata
func (m *VMTypeResponseWithMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Vm type response with metadata based on context it is used
func (m *VMTypeResponseWithMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMTypeResponseWithMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMTypeResponseWithMetadata) UnmarshalBinary(b []byte) error {
	var res VMTypeResponseWithMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
