// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateApplicationResourcesLocalStorageSize Storage related information.
//
// swagger:model CreateApplicationResourcesLocalStorageSize
type CreateApplicationResourcesLocalStorageSize struct {

	// Local disk space used for writing cache data.
	Cache string `json:"cache,omitempty"`

	// Local disk space used for writing other temporary data, tools, etc.
	Overhead string `json:"overhead,omitempty"`

	// Local disk space used for writing scratch data.
	Scratch string `json:"scratch,omitempty"`
}

// Validate validates this create application resources local storage size
func (m *CreateApplicationResourcesLocalStorageSize) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create application resources local storage size based on context it is used
func (m *CreateApplicationResourcesLocalStorageSize) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateApplicationResourcesLocalStorageSize) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateApplicationResourcesLocalStorageSize) UnmarshalBinary(b []byte) error {
	var res CreateApplicationResourcesLocalStorageSize
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
