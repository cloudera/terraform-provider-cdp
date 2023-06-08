// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListDatalakesRequest Request object for list datalakes request.
//
// swagger:model ListDatalakesRequest
type ListDatalakesRequest struct {

	// The name or CRN of the datalake for which details are requested.
	DatalakeName string `json:"datalakeName,omitempty"`

	// The name or CRN of the environment for which the datalakes will be listed.
	EnvironmentName string `json:"environmentName,omitempty"`
}

// Validate validates this list datalakes request
func (m *ListDatalakesRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list datalakes request based on context it is used
func (m *ListDatalakesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListDatalakesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListDatalakesRequest) UnmarshalBinary(b []byte) error {
	var res ListDatalakesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
