// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ScaleHorizontallyResponse Data Lake horizontal scaling response
//
// swagger:model ScaleHorizontallyResponse
type ScaleHorizontallyResponse struct {

	// The name of the datalake.
	DatalakeName string `json:"datalakeName,omitempty"`

	// Unique operation ID assigned to this command execution. Use this identifier with 'get-operation' to track status and retrieve detailed results.
	OperationID string `json:"operationId,omitempty"`
}

// Validate validates this scale horizontally response
func (m *ScaleHorizontallyResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this scale horizontally response based on context it is used
func (m *ScaleHorizontallyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScaleHorizontallyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScaleHorizontallyResponse) UnmarshalBinary(b []byte) error {
	var res ScaleHorizontallyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
